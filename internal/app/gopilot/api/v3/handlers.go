package v3

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/api/v3/controllers"
)

func RegisterV3Handlers(prefix string, e *echo.Echo) {
	// Pilots
	e.GET(fmt.Sprintf("%s/buildpacks", prefix), controllers.GetBuildpacks)
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
