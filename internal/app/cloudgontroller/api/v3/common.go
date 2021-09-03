package v3

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetResourcePath returns the full url that was called by the web request
func GetResourcePath(c echo.Context) string {
	return fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, c.Request().URL.Path)
}

type Metadata struct {
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
}
