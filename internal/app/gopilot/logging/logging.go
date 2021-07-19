package logging

import (
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/config"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/helpers"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/logging/enums"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Setup(config *config.GopilotConfig) error {
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
	return nil
}
