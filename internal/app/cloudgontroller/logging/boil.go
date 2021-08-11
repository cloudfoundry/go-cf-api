package logging

import (
	"errors"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// nolint:godox
// TODO Maybe Switch to https://github.com/simukti/sqldb-logger and log on the driver level

type BoilLogger struct {
	RedactParams             bool
	query, params, requestID string
}

func NewBoilLogger(redactParams bool, c echo.Context) *BoilLogger {
	return &BoilLogger{RedactParams: redactParams, requestID: c.Response().Header().Get(HeaderVcapRequestID)}
}

func (b *BoilLogger) Write(data []byte) (n int, err error) {
	if len(b.query) == 0 {
		b.query = string(data)
		return 0, nil
	} else if len(b.params) == 0 {
		b.params = string(data)
		if !b.RedactParams {
			b.getLogger().Debug("SQL LOG", zap.String("query", b.query), zap.String("params", b.params))
		} else {
			b.getLogger().Debug("SQL LOG", zap.String("query", b.query), zap.String("params", "REDACTED"))
		}
		b.query = ""
		b.params = ""
		return 0, nil
	}
	return 0, errors.New("we got a sql query parameter but no sql query")
}

func (b *BoilLogger) getLogger() *zap.Logger {
	for i := 3; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, "src/fmt/print.go"):
		case strings.Contains(file, "logging/boil.go"):
		default:
			return zap.L().WithOptions(zap.AddCallerSkip(i), zap.AddCaller())
		}
	}
	return zap.L().With(zap.String("request_id", b.requestID))
}
