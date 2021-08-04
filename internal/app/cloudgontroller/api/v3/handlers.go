package v3

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
)

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", controllers.GetHealth)
}

func RegisterV3Handlers(prefix string, e *echo.Echo) {
	// Buildpacks
	e.GET(fmt.Sprintf("%s/buildpacks", prefix), controllers.GetBuildpacks)
	e.GET(fmt.Sprintf("%s/buildpacks/:guid", prefix), controllers.GetBuildpack)
}

func RegisterV3DocumentationHandlers(prefix string, e *echo.Echo) {
	// Serve Swagger-UI API Docs
	e.GET(prefix, func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
		return nil
	})
	e.GET(fmt.Sprintf("%s/", prefix), func(c echo.Context) error {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", prefix))
		return nil
	})
	e.GET(fmt.Sprintf("%s/*", prefix), echoSwagger.WrapHandler)
}
