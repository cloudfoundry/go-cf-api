package logging

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func NewTimingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			startTime := time.Now().UTC()
			ctx.Response().Before(func() {
				runtime := time.Now().UTC().Sub(startTime).Seconds()
				ctx.Response().Header().Set("X-Runtime", strconv.FormatFloat(runtime, 'f', 6, 64)) // nolint:gomnd
			})
			return next(ctx)
		}
	}
}
