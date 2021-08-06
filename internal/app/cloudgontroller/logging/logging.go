package logging

import (
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(config *config.CloudgontrollerConfig) error {
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
