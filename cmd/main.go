package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudfoundry/go-cf-api/internal/api"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon"
	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/auth"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/ratelimiter"
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/helpers"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/metrics"
	"github.com/cloudfoundry/go-cf-api/internal/metrics/custommetrics"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db"
	"github.com/cloudfoundry/go-cf-api/internal/uaa"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/spf13/cobra"
	_ "github.com/volatiletech/null/v8"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func RootFunc(cmd *cobra.Command, args []string) error { //nolint:funlen // length is not a problem for now
	// Parse Config
	var conf *config.CfAPIConfig
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
	echoInstance := echo.New()
	echoInstance.JSONSerializer = &apicommon.JSONSerializer{DefaultJSONSerializer: echo.DefaultJSONSerializer{}}
	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.IPExtractor = echo.ExtractIPFromXFFHeader()
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(logging.NewTimingMiddleware())
	echoInstance.Use(logging.NewVcapRequestID())
	echoInstance.Use(logging.NewEchoZapLogger(zap.L()))
	metrics.EchoPrometheusMiddleware().Use(echoInstance)
	echoInstance.HTTPErrorHandler = func(err error, ctx echo.Context) {
		var ccErr *v3.CfAPIError
		if errors.As(err, &ccErr) {
			errResponse := ctx.JSON(ccErr.HTTPStatus, v3.AsErrors(*ccErr))
			if errResponse != nil {
				_ = ctx.String(http.StatusInternalServerError, "Unknown error occurred")
				zap.L().Error(err.Error())
			}
			zap.L().Error(errors.Unwrap(err).Error())
		} else {
			echoInstance.DefaultHTTPErrorHandler(err, ctx)
		}
	}

	// Add Validator for Structs
	echoInstance.Validator = &apicommon.Validator{Validator: validator.New()}

	// Initialize DB
	database, _, err := db.NewConnection(conf.DB, true)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	// Configure auth middleware
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ukf, err := uaa.NewKeyFetcher(ctx, conf.Uaa)
	helpers.CheckErrFatal(err)
	jwtMiddleware := auth.NewJWTMiddleware(ukf.Fetch)

	userHashResetInterval := ratelimiter.NewUserHashResetInterval(conf.RateLimit.ResetInterval, time.Now)
	generalRateLimiter := ratelimiter.NewRateLimiter(
		conf.RateLimit.GlobalGeneralLimit,
		conf.RateLimit.PerProcessGeneralLimit,
		userHashResetInterval,
	)
	unauthenticatedRateLimiter := ratelimiter.NewRateLimiter(
		conf.RateLimit.GlobalUnauthenticatedLimit,
		conf.RateLimit.PerProcessUnauthenticatedLimit,
		userHashResetInterval,
	)
	rateLimitMiddleware := ratelimiter.NewRateLimiterMiddleware(generalRateLimiter, unauthenticatedRateLimiter)

	// Register API Handlers
	api.RegisterHandlers(echoInstance, database, jwtMiddleware, rateLimitMiddleware, conf)
	prometheus.MustRegister(collectors.NewDBStatsCollector(database, conf.DB.Type), custommetrics.NewCustomCollector(time.Now().UTC()))

	// Start to Serve
	lock := make(chan error)
	go func(lock chan error) { lock <- echoInstance.Start(fmt.Sprintf("%v", conf.Listen)) }(lock)

	time.Sleep(1 * time.Millisecond)
	zap.L().Warn("Application started without ssl/tls enabled")

	if err := <-lock; err != nil {
		zap.L().Panic("failed to start application", zap.Error(err))
	}
	return nil
}

// @title CF API
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

	exe, err := os.Executable()
	if err != nil {
		panic(err)
	}

	// Cobra Flags.
	rootCmd := &cobra.Command{
		Use:     "CF API [ConfigPath]",
		Short:   "POC implemetation of a CAPI V3 API compatible golang webserver",
		Example: fmt.Sprintf("%s/%s config_psql.yaml", filepath.Dir(exe), filepath.Base(os.Args[0])),
		Args:    cobra.ExactArgs(1),
		RunE:    RootFunc,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
		os.Exit(1)
	}
}
