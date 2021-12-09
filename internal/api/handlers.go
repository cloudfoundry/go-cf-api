package api

import (
	"database/sql"
	"fmt"
	"net/http"

	// Blank import needed for swagger doc
	_ "github.com/cloudfoundry/go-cf-api/internal/api/docs"
	"github.com/cloudfoundry/go-cf-api/internal/api/health"
	v3 "github.com/cloudfoundry/go-cf-api/internal/api/v3"
	"github.com/cloudfoundry/go-cf-api/internal/api/v3/buildpacks"
	"github.com/cloudfoundry/go-cf-api/internal/api/v3/info"
	"github.com/cloudfoundry/go-cf-api/internal/api/v3/securitygroups"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/auth"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/metadata"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/permissions"
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterHandlers(
	e *echo.Echo,
	database *sql.DB,
	jwtMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CfAPIConfig,
) {
	// Health Endpoint
	e.GET("healthz", health.GetHealth)

	// Info Endpoint
	e.GET("/", NewRootEndpoint(conf))

	// V3 API
	e.GET("/v3", v3.NewV3RootEndpoint(conf))
	registerV3Handlers(e, database, jwtMiddleware, rateLimitMiddleware, conf)

	// Serve V3 Swagger-UI API Docs
	e.GET("docs/v3", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", "docs/v3"))
	})
	e.GET(fmt.Sprintf("%s/", "docs/v3"), func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/%s/index.html", "docs/v3"))
	})
	e.GET(fmt.Sprintf("%s/*", "docs/v3"), echoSwagger.WrapHandler)
}

func registerV3Handlers(echoInstance *echo.Echo,
	database *sql.DB,
	jwtMiddleware echo.MiddlewareFunc,
	rateLimitMiddleware echo.MiddlewareFunc,
	conf *config.CfAPIConfig) {
	v3Root := echoInstance.Group("v3/", jwtMiddleware)
	if conf.RateLimit.Enabled {
		v3Root.Use(rateLimitMiddleware)
	}
	requiresRead := v3Root.Group("", auth.NewRequiresReadMiddleware())
	requiresWrite := v3Root.Group("", auth.NewRequiresWriteMiddleware())

	// Info
	echoInstance.GET("/v3/info", info.NewV3InfoEndpoint(conf))

	// Buildpacks
	buildpacksController := buildpacks.Controller{DB: database, Presenter: buildpacks.NewPresenter(), LabelSelectorParser: metadata.NewLabelSelectorParser()}
	requiresRead.GET("buildpacks", buildpacksController.List)
	requiresRead.GET(fmt.Sprintf("buildpacks/:%s", buildpacks.GUIDParam), buildpacksController.Get)
	requiresWrite.POST("buildpacks", buildpacksController.Post)

	// Security Groups
	securityGroupsController := securitygroups.Controller{DB: database, Presenter: securitygroups.NewPresenter(), Permissions: permissions.NewQuerier()}
	requiresRead.GET(fmt.Sprintf("security_groups/:%s", securitygroups.GUIDParam), securityGroupsController.Get)
	requiresRead.GET("security_groups", securityGroupsController.List)
}
