package info

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetV3Root(c echo.Context) error {
	links := map[string]Link{"self": {HREF: cont.externalURL("v3")}}
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
		links[endpoint] = Link{HREF: cont.externalURL(fmt.Sprintf("v3/%s", endpoint))}
	}
	return c.JSONPretty(http.StatusOK, map[string]interface{}{"links": links}, indent)
}
