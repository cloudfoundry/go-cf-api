package api

import (
	"github.com/labstack/echo/v4"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/api/v3"
	_ "github.tools.sap/cloudfoundry/cloudgontroller/swagger"
)

func RegisterHandlers(e *echo.Echo) {
	v3.RegisterV3Handlers("api/v3", e)
	v3.RegisterV3DocumentationHandlers("docs/v3", e)
	// TODO Authentication
}
