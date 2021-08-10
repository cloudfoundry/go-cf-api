package v3

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
)

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", controllers.GetHealth)
}

func RegisterV3Handlers(prefix string, e *echo.Echo, authMiddleware echo.MiddlewareFunc) {
	// Restricted group
	restrictedGroup := e.Group(prefix)
	restrictedGroup.Use(authMiddleware)
	{
		// Buildpacks
		restrictedGroup.GET("/buildpacks", controllers.GetBuildpacks)
		restrictedGroup.GET("/buildpacks/:guid", controllers.GetBuildpack)
	}
}

func RegisterV3DocumentationHandlers(prefix string, e *echo.Echo) {
	// Serve Swagger-UI API Docs
	e.GET(prefix, func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
	})
	e.GET(fmt.Sprintf("%s/", prefix), func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
	})
	e.GET(fmt.Sprintf("%s/*", prefix), echoSwagger.WrapHandler)
}
