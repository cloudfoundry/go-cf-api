package helpers

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func NormalizeAuthScheme(c echo.Context) {
	headerValue := c.Request().Header.Get(echo.HeaderAuthorization)
	if substrings := strings.SplitN(headerValue, " ", 2); len(substrings) > 1 { //nolint:gomnd
		authScheme := substrings[0]
		token := substrings[1]
		headerValue = fmt.Sprintf("%s %s", strings.ToLower(authScheme), token)
		c.Request().Header.Set(echo.HeaderAuthorization, headerValue)
	}
}
