package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"time"

	_ "github.com/FloThinksPi/golang-vuejs-bootstrap/docs/swagger"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/api"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/config"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/logging"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/metrics"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/storage/db"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/time/rate"
)

var (
	//go:embed frontend
	embededFrontendFiles embed.FS

	// Cobra Flags
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "gopilot",
		Short: "A example webserver that can be used as a bootstrap project",
		// Long: `GoPilot is a example webserver that can be useds as a bootstrap project.
		// 		It provides many patterns out of the box like automatic api documentation, vuejs frontend, fast echo webserver,
		// 		patterns for fast structured logging and prometheus metrics, rate limiting, config management,
		// 		sqlboiler generated models from db structure, db schema creation and migration,
		// 		interactive access to the models with a RPEL console, mage commands for developers and operators and more ...`,
		Example: "gopilot config.yaml",
		Args:    cobra.MaximumNArgs(1),
		Run:     RootFunc,
	}
)

func RootFunc(cmd *cobra.Command, args []string) {
	// Parse Config
	var conf *config.GopilotConfig
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

	// Register Frontend Handler
	frontendFS, err := fs.Sub(embededFrontendFiles, "frontend")
	if err != nil {
		panic(err)
	}
	frontendHandler := http.FileServer(http.FS(frontendFS))
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", frontendHandler)))

	// Register API Handlers
	api.RegisterHandlers(e)

	// Start to Serve
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(fmt.Sprintf(":%v", conf.Port)) }(lock)

	time.Sleep(1 * time.Millisecond)
	zap.L().Warn("application started without ssl/tls enabled")

	err = <-lock
	if err != nil {
		zap.L().Panic("failed to start application")
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

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

	// Init Cobra and Viper
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
}
