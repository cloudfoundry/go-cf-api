package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/handlers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

func RegisterHandlers(
	e *echo.Echo,
	db *sql.DB,
	jwtMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CloudgontrollerConfig,
) {
	handlers.RegisterHealthHandler(e)
	handlers.RegisterInfoHandlers(e, conf)
	handlers.RegisterV3Handlers("v3", e, db, jwtMiddleware, rateLimitMiddleware, conf)
	handlers.RegisterV3DocumentationHandlers("docs/v3", e)
}
