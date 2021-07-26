package logging

import (
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/helpers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging/enums"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var once sync.Once

func Setup(config *config.CloudgontrollerConfig) error {
	once.Do(func() {
		// Logging
		var zapconf zap.Config
		if config.Log.Production {
			zapconf = zap.NewProductionConfig()
		} else {
			zapconf = zap.NewDevelopmentConfig()
		}
		zapconf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapconf.DisableStacktrace = true
		zapLogLevel, err := enums.StringToLogLevel(config.Log.Level)
		helpers.CheckErrFatal(err)
		zapconf.Level.SetLevel(zapLogLevel)
		logger, _ := zapconf.Build()
		zap.ReplaceGlobals(logger)
		logger.Info("Logger Initialized", zap.Bool("Production", config.Log.Production), zap.String("LogLevel", string(config.Log.Level)))
	})
	return nil
}
