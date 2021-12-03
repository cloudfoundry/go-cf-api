package api

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/health"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/info"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/buildpacks"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/metadata"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/securitygroups"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/permissions"
	"net/http"
)

func RegisterHandlers(
	e *echo.Echo,
	db *sql.DB,
	jwtMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CloudgontrollerConfig,
) {
	// Health Endoint
	e.GET("healthz", health.GetHealth)

	// Info Endpoint
	infoController := info.Controller{
		Info:             conf.Info,
		URLs:             conf.URLs,
		AppSSH:           conf.AppSSH,
		ExternalDomain:   conf.ExternalDomain,
		ExternalProtocol: conf.ExternalProtocol,
	}
	e.GET("/", infoController.GetRoot)
	e.GET("/v3", infoController.GetV3Root)
	e.GET("/v3/info", infoController.GetV3Info)

	// V3 API
	registerV3Handlers(e, db, jwtMiddleware, rateLimitMiddleware, conf)

	// Serve Swagger-UI API Docs
	e.GET("docs/v3", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", "docs/v3"))
	})
	e.GET(fmt.Sprintf("%s/", "docs/v3"), func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", "docs/v3"))
	})
	e.GET(fmt.Sprintf("%s/*", "docs/v3"), echoSwagger.WrapHandler)
}

func registerV3Handlers(e *echo.Echo, db *sql.DB, jwtMiddleware echo.MiddlewareFunc, rateLimitMiddleware echo.MiddlewareFunc, conf *config.CloudgontrollerConfig) {
	v3Root := e.Group("v3/", jwtMiddleware)
	if conf.RateLimit.Enabled {
		v3Root.Use(rateLimitMiddleware)
	}
	requiresRead := v3Root.Group("", auth.NewRequiresReadMiddleware())
	requiresWrite := v3Root.Group("", auth.NewRequiresWriteMiddleware())

	// Buildpacks
	buildpacksController := buildpacks.Controller{DB: db, Presenter: buildpacks.NewPresenter(), LabelSelectorParser: metadata.NewLabelSelectorParser()}
	requiresRead.GET("buildpacks", buildpacksController.List)
	requiresRead.GET(fmt.Sprintf("buildpacks/:%s", buildpacks.GUIDParam), buildpacksController.Get)
	requiresWrite.POST("buildpacks", buildpacksController.Post)

	// Security Groups
	securityGroupsController := securitygroups.Controller{DB: db, Presenter: securitygroups.NewPresenter(), Permissions: permissions.NewQuerier()}
	requiresRead.GET(fmt.Sprintf("security_groups/:%s", securitygroups.GUIDParam), securityGroupsController.Get)
	requiresRead.GET("security_groups", securityGroupsController.List)
}
