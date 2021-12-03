package metrics

import (
	"strings"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

// urlSkipper ignores metrics route on some middleware.
func urlSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Path(), "/docs")
}

func EchoPrometheusMiddleware() *prometheus.Prometheus {
	return prometheus.NewPrometheus("echo", urlSkipper)
}
