package api

import (
	_ "github.com/FloThinksPi/golang-vuejs-bootstrap/docs/swagger"
	v1 "github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/api/v1"
	"github.com/labstack/echo/v4"
)

func RegisterHandlers(e *echo.Echo) {
	v1.RegisterHandlers("api/v1", e)
	v1.RegisterDocumentationHandlers("docs/v1", e)
	// TODO Authentication
}
