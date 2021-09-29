package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	// Needed for swagger
	_ "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/swagger"
	buildpacks "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/buildpacks"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/metadata"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/securitygroups"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/permissions"
)

func RegisterHealthHandler(e *echo.Echo) {
	// Health
	e.GET("healthz", getHealth)
}

func RegisterV3Handlers(
	prefix string,
	e *echo.Echo,
	db *sql.DB,
	jwtMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CloudgontrollerConfig,
) {
	v3Root := e.Group(prefix, jwtMiddleware)
	if conf.RateLimit.Enabled {
		v3Root.Use(rateLimitMiddleware)
	}
	requiresRead := v3Root.Group("", auth.NewRequiresReadMiddleware())
	requiresWrite := v3Root.Group("", auth.NewRequiresWriteMiddleware())

	// Buildpacks
	buildpacksController := buildpacks.Controller{DB: db, LabelSelectorParser: metadata.NewLabelSelectorParser()}
	requiresRead.GET("/buildpacks", buildpacksController.List)
	requiresRead.GET(fmt.Sprintf("/buildpacks/:%s", buildpacks.GUIDParam), buildpacksController.Get)
	requiresWrite.POST("/buildpacks", buildpacksController.Post)

	// Security Groups
	securityGroupsController := securitygroups.Controller{DB: db, Presenter: securitygroups.NewPresenter(), Permissions: permissions.NewQuerier()}
	requiresRead.GET(fmt.Sprintf("/security_groups/:%s", securitygroups.GUIDParam), securityGroupsController.Get)
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
