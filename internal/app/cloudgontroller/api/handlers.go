package api

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
)

func RegisterHandlers(e *echo.Echo, db *sql.DB, authMiddleware echo.MiddlewareFunc) {
	v3.RegisterHealthHandler(e)
	v3.RegisterV3Handlers("v3", e, db, authMiddleware)
	v3.RegisterV3DocumentationHandlers("docs/v3", e)
}
