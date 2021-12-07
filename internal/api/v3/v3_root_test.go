//go:build unit
// +build unit

package v3_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
)

func TestV3Root(t *testing.T) {
	expectedResponse := `
	{
		"links": {
			"app_usage_events": {
				"href": "https://api.external.domain/v3/app_usage_events"
			},
			"apps": {
				"href": "https://api.external.domain/v3/apps"
			},
			"audit_events": {
				"href": "https://api.external.domain/v3/audit_events"
			},
			"buildpacks": {
				"href": "https://api.external.domain/v3/buildpacks"
			},
			"builds": {
				"href": "https://api.external.domain/v3/builds"
			},
			"deployments": {
				"href": "https://api.external.domain/v3/deployments"
			},
			"domains": {
				"href": "https://api.external.domain/v3/domains"
			},
			"droplets": {
				"href": "https://api.external.domain/v3/droplets"
			},
			"environment_variable_groups": {
				"href": "https://api.external.domain/v3/environment_variable_groups"
			},
			"feature_flags": {
				"href": "https://api.external.domain/v3/feature_flags"
			},
			"info": {
				"href": "https://api.external.domain/v3/info"
			},
			"isolation_segments": {
				"href": "https://api.external.domain/v3/isolation_segments"
			},
			"organization_quotas": {
				"href": "https://api.external.domain/v3/organization_quotas"
			},
			"organizations": {
				"href": "https://api.external.domain/v3/organizations"
			},
			"packages": {
				"href": "https://api.external.domain/v3/packages"
			},
			"processes": {
				"href": "https://api.external.domain/v3/processes"
			},
			"resource_matches": {
				"href": "https://api.external.domain/v3/resource_matches"
			},
			"roles": {
				"href": "https://api.external.domain/v3/roles"
			},
			"routes": {
				"href": "https://api.external.domain/v3/routes"
			},
			"security_groups": {
				"href": "https://api.external.domain/v3/security_groups"
			},
			"self": {
				"href": "https://api.external.domain/v3"
			},
			"service_brokers": {
				"href": "https://api.external.domain/v3/service_brokers"
			},
			"service_credential_bindings": {
				"href": "https://api.external.domain/v3/service_credential_bindings"
			},
			"service_instances": {
				"href": "https://api.external.domain/v3/service_instances"
			},
			"service_offerings": {
				"href": "https://api.external.domain/v3/service_offerings"
			},
			"service_plans": {
				"href": "https://api.external.domain/v3/service_plans"
			},
			"service_route_bindings": {
				"href": "https://api.external.domain/v3/service_route_bindings"
			},
			"service_usage_events": {
				"href": "https://api.external.domain/v3/service_usage_events"
			},
			"space_quotas": {
				"href": "https://api.external.domain/v3/space_quotas"
			},
			"spaces": {
				"href": "https://api.external.domain/v3/spaces"
			},
			"stacks": {
				"href": "https://api.external.domain/v3/stacks"
			},
			"tasks": {
				"href": "https://api.external.domain/v3/tasks"
			},
			"users": {
				"href": "https://api.external.domain/v3/users"
			}
		}
	}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	handleFunc := v3.NewV3RootEndpoint(&config.CfAPIConfig{
		ExternalDomain:   "api.external.domain",
		ExternalProtocol: "https",
	})
	err := handleFunc(ctx)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}
