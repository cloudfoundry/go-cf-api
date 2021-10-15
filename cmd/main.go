package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/spf13/cobra"
	_ "github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/ratelimiter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/metrics"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/metrics/custommetrics"
	dbconfig "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/uaa"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RootFunc(cmd *cobra.Command, args []string) error { //nolint:funlen // length is not a problem for now
	// Parse Config
	var conf *config.CloudgontrollerConfig
	var err error
	if len(args) == 1 {
		// Parse File and Env
		conf, err = config.Get(args[0])
	} else {
		// Just parses Environment Variables
		conf, err = config.Get("")
	}

	if err != nil {
		return err
	}
	// Initialize Logger
	if err := logging.Setup(conf); err != nil {
		return err
	}

	// Initialize Echo Framework
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Use(middleware.Recover())
	e.Use(logging.NewVcapRequestID())
	e.Use(logging.NewEchoZapLogger(zap.L()))
	metrics.EchoPrometheusMiddleware().Use(e)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		var ccErr *v3.CloudControllerError
		if errors.As(err, &ccErr) {
			errResponse := c.JSON(ccErr.HTTPStatus, v3.AsErrors(*ccErr))
			if errResponse != nil {
				_ = c.String(http.StatusInternalServerError, "Unknown error occurred")
				zap.L().Error(err.Error())
			}
			zap.L().Error(errors.Unwrap(err).Error())
		} else {
			e.DefaultHTTPErrorHandler(err, c)
		}
	}

	// Add Validator for Structs
	e.Validator = &api.Validator{Validator: validator.New()}

	// Initialize DB
	db, _, err := dbconfig.NewConnection(conf.DB, true)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	// Configure auth middleware
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ukf, err := uaa.NewKeyFetcher(ctx, conf.Uaa)
	helpers.CheckErrFatal(err)
	jwtMiddleware := auth.NewJWTMiddleware(ukf.Fetch)
	generalRateLimiter := ratelimiter.NewRateLimiter(conf.RateLimit.GeneralLimit, conf.RateLimit.ResetInterval)
	unauthenticatedRateLimiter := ratelimiter.NewRateLimiter(conf.RateLimit.UnauthenticatedLimit, conf.RateLimit.ResetInterval)
	rateLimitMiddleware := ratelimiter.NewRateLimiterMiddleware(generalRateLimiter, unauthenticatedRateLimiter)

	// Register API Handlers
	api.RegisterHandlers(e, db, jwtMiddleware, rateLimitMiddleware, conf)
	prometheus.MustRegister(collectors.NewDBStatsCollector(db, conf.DB.Type), custommetrics.NewCustomCollector(time.Now().UTC()))

	// Start to Serve
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(fmt.Sprintf("%v", conf.Listen)) }(lock)

	time.Sleep(1 * time.Millisecond)
	zap.L().Warn("Application started without ssl/tls enabled")

	if err := <-lock; err != nil {
		zap.L().Panic("failed to start application", zap.Error(err))
	}
	return nil
}

// @title CloudGontroller API
// @version 3.0.0
// @description CAPI V3 Compatible API with blazing fast backend.

// @host localhost:8080
// @BasePath /v3
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	config.Level.SetLevel(zapcore.InfoLevel)
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)

	// Cobra Flags.
	rootCmd := &cobra.Command{
		Use:   "cloudgontroller",
		Short: "POC implemetation of a CAPI V3 compatible golang webserver",
		// Long: `cloudgontroller is a example webserver that can be useds as a bootstrap project.
		// 		It provides many patterns out of the box like automatic api documentation, vuejs frontend, fast echo webserver,
		// 		patterns for fast structured logging and prometheus metrics, rate limiting, config management,
		// 		sqlboiler generated models from db structure, db schema creation and migration,
		// 		interactive access to the models with a RPEL console, mage commands for developers and operators and more ...`,
		Example: "cloudgontroller config_psql.yaml",
		Args:    cobra.MaximumNArgs(1),
		RunE:    RootFunc,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
		os.Exit(1)
	}
}
