package main

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api"
	_ "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/metrics"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

var (

	// Cobra Flags
	rootCmd = &cobra.Command{
		Use:   "cloudgontroller",
		Short: "POC implemetation of a CAPI V3 compatible golang webserver",
		// Long: `cloudgontroller is a example webserver that can be useds as a bootstrap project.
		// 		It provides many patterns out of the box like automatic api documentation, vuejs frontend, fast echo webserver,
		// 		patterns for fast structured logging and prometheus metrics, rate limiting, config management,
		// 		sqlboiler generated models from db structure, db schema creation and migration,
		// 		interactive access to the models with a RPEL console, mage commands for developers and operators and more ...`,
		Example: "cloudgontroller config_psql.yaml",
		Args:    cobra.MaximumNArgs(1),
		Run:     RootFunc,
	}
)

func RootFunc(cmd *cobra.Command, args []string) {
	// Parse Config
	var conf *config.CloudgontrollerConfig
	if len(args) == 1 {
		// Parse File and Env
		conf = config.Get(args[0])
	} else {
		// Just parses Environment Variables
		conf = config.Get()
	}
	// Initialize Logger
	logging.Setup(conf)

	// Initialize Echo Framework
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(logging.NewEchoZapLogger(zap.L()))
	metrics.EchoPrometheusMiddleware().Use(e)
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(conf.RateLimit.Rate), Burst: conf.RateLimit.Burst, ExpiresIn: conf.RateLimit.ExpiresIn},
	)))

	// Add Validator for Structs
	e.Validator = &api.Validator{Validator: validator.New()}

	// Initialize DB
	db.NewConnection(conf.DB, true)
	if conf.DB.Migrate {
		db.Migrate(db.GetConnection())
	}

	// Register API Handlers
	api.RegisterHandlers(e)

	// Start to Serve
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(fmt.Sprintf("%v", conf.Listen)) }(lock)

	time.Sleep(1 * time.Millisecond)
	zap.L().Warn("Application started without ssl/tls enabled")

	err := <-lock
	if err != nil {
		zap.L().Panic("failed to start application",zap.Error(err))
	}
}

// @title CloudGontroller API
// @version 3.0.0
// @description CAPI V3 Compatible API with blazing fast backend.

// @host localhost:8080
// @BasePath /api/v3
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	rootCmd.Execute()
}

func init() {
	// Initialize pre start Logging so we see e.g. config parse errors
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	config.Level.SetLevel(zapcore.InfoLevel)
	logger, _ := config.Build()
	zap.ReplaceGlobals(logger)
}
