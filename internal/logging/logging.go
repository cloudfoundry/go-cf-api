package logging

import (
	"github.com/labstack/echo/v4"
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/helpers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(config *config.CfAPIConfig) error {
	// Logging
	var zapconf zap.Config
	if config.Log.Production {
		zapconf = zap.NewProductionConfig()
	} else {
		zapconf = zap.NewDevelopmentConfig()
	}
	zapconf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapconf.DisableStacktrace = true
	var zapLogLevel zapcore.Level
	err := zapLogLevel.UnmarshalText([]byte(config.Log.Level))
	helpers.CheckErrFatal(err)
	zapconf.Level.SetLevel(zapLogLevel)
	logger, _ := zapconf.Build()
	zap.ReplaceGlobals(logger)
	logger.Info("Logger Initialized", zap.Bool("Production", config.Log.Production), zap.String("LogLevel", config.Log.Level))
	return nil
}

func FromContext(c echo.Context) *zap.Logger {
	logger := zap.L().With(zap.String(RequestIDField, c.Response().Header().Get(HeaderVcapRequestID)))
	return logger
}
