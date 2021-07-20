package api

import (
	"github.com/labstack/echo/v4"
	_ "github.tools.sap/cloudfoundry/cloudgontroller/docs/swagger"
	v1 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/api/v3"
)

func RegisterHandlers(e *echo.Echo) {
	v1.RegisterV3Handlers("api/v3", e)
	v1.RegisterV3DocumentationHandlers("docs/v3", e)
	// TODO Authentication
}
