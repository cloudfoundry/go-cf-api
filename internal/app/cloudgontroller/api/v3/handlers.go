package v3

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", controllers.GetHealth)
}

func RegisterV3Handlers(
	prefix string,
	e *echo.Echo,
	db *sql.DB,
	authMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CloudgontrollerConfig,
) {
	// Restricted group
	restrictedGroup := e.Group(prefix)
	restrictedGroup.Use(authMiddleware)
	if conf.RateLimit.Enabled {
		restrictedGroup.Use(rateLimitMiddleware)
	}
	buildpacksController := controllers.BuildpackController{DB: db}
	{
		// Buildpacks
		restrictedGroup.GET("/buildpacks", buildpacksController.GetBuildpacks)
		restrictedGroup.GET(fmt.Sprintf("/buildpacks/:%s", controllers.GUIDParam), buildpacksController.GetBuildpack)
		restrictedGroup.POST("/buildpacks", buildpacksController.PostBuildpack)
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
