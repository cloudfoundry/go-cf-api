package v3

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/info"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
)

func NewV3RootEndpoint(config *config.CfAPIConfig) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		links := map[string]info.Link{"self": {HREF: info.ExternalURL("v3", config)}}
		endpoints := []string{
			"app_usage_events",
			"apps",
			"audit_events",
			"buildpacks",
			"builds",
			"deployments",
			"domains",
			"droplets",
			"environment_variable_groups",
			"feature_flags",
			"info",
			"isolation_segments",
			"organizations",
			"organization_quotas",
			"packages",
			"processes",
			"resource_matches",
			"roles",
			"routes",
			"security_groups",
			"service_brokers",
			"service_instances",
			"service_credential_bindings",
			"service_offerings",
			"service_plans",
			"service_route_bindings",
			"service_usage_events",
			"spaces",
			"space_quotas",
			"stacks",
			"tasks",
			"users",
		}
		for _, endpoint := range endpoints {
			links[endpoint] = info.Link{HREF: info.ExternalURL(fmt.Sprintf("v3/%s", endpoint), config)}
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{"links": links})
	}
}
