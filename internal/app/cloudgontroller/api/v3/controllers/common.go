package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func GetResourcePath(c echo.Context) string {
	return fmt.Sprintf("%s://%s%s", c.Scheme(), c.Request().Host, c.Request().URL.Path)
}
