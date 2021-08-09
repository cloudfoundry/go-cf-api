package api

import (
	"github.com/labstack/echo/v4"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

func RegisterHandlers(e *echo.Echo, conf *config.CloudgontrollerConfig) {
	v3.RegisterHealthHandler(e)
	v3.RegisterV3Handlers("v3", e, conf)
	v3.RegisterV3DocumentationHandlers("docs/v3", e)
}
