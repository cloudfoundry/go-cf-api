package logging

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func NewTimingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			startTime := time.Now().UTC()
			c.Response().Before(func() {
				runtime := time.Now().UTC().Sub(startTime).Seconds()
				c.Response().Header().Set("X-Runtime", strconv.FormatFloat(runtime, 'f', 6, 64)) // nolint:gomnd
			})
			return next(c)
		}
	}
}
