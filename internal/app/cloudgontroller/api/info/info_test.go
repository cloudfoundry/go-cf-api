// +build unit

package info_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/info"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

func TestV3Info(t *testing.T) {
	cases := map[string]struct {
		externalDomain   string
		externalProtocol string
		infoConfig       config.InfoConfig
		expectedResponse string
	}{
		"only required config": {
			externalDomain:   "example.org",
			externalProtocol: "http",
			infoConfig: config.InfoConfig{
				Name:           "foo",
				Description:    "test cc",
				Build:          "test",
				Version:        1234,
				SupportAddress: "i@need.help",
			},
			expectedResponse: `{
				"name": "foo",
				"description": "test cc",
				"build": "test",
				"version": 1234,
				"cli_version": {
					"minimum": "",
					"recommended": ""
				},
				"links": {
					"self": {
						"href": "http://example.org/v3/info"
					},
					"support": {
						"href": "i@need.help"
					}
				}
			}`,
		},
		"with optional config": {
			externalDomain:   "example.org",
			externalProtocol: "http",
			infoConfig: config.InfoConfig{
				Name:                     "foo",
				Description:              "test cc",
				Build:                    "test",
				Version:                  1234,
				MinCliVersion:            "1.0.0",
				MinRecommendedCliVersion: "1.2.3",
				SupportAddress:           "i@need.help",
			},
			expectedResponse: `{
				"name": "foo",
				"description": "test cc",
				"build": "test",
				"version": 1234,
				"cli_version": {
					"minimum": "1.0.0",
					"recommended": "1.2.3"
				},
				"links": {
					"self": {
						"href": "http://example.org/v3/info"
					},
					"support": {
						"href": "i@need.help"
					}
				}
			}`,
		},
		"different external domain and protocol": {
			externalDomain:   "example.org:2222",
			externalProtocol: "ssh",
			infoConfig: config.InfoConfig{
				Name:           "foo",
				Description:    "test cc",
				Build:          "test",
				Version:        1234,
				SupportAddress: "i@need.help",
			},
			expectedResponse: `{
				"name": "foo",
				"description": "test cc",
				"build": "test",
				"version": 1234,
				"cli_version": {
					"minimum": "",
					"recommended": ""
				},
				"links": {
					"self": {
						"href": "ssh://example.org:2222/v3/info"
					},
					"support": {
						"href": "i@need.help"
					}
				}
			}`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/info", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{
				ExternalDomain:   tc.externalDomain,
				ExternalProtocol: tc.externalProtocol,
				Info:             tc.infoConfig,
			}
			err := controller.GetV3Info(c)

			assert.NoError(t, err)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}

func TestRoot(t *testing.T) {
	cases := map[string]struct {
		externalDomain   string
		externalProtocol string
		urls             config.URLs
		appSSH           config.AppSSH
		expectedResponse string
	}{
		"only required config": {
			externalDomain:   "example.org",
			externalProtocol: "https",
			urls: config.URLs{
				LogCache:   "https://log-cache.logging.system",
				LogStream:  "https://log-stream.logging.system",
				Doppler:    "wss://doppler.logging.system:443",
				Login:      "https://login.identity.system",
				UAA:        "https://uaa.identity.system",
				RoutingAPI: nil,
			},
			appSSH: config.AppSSH{
				Endpoint:    "ssh://ssh.apps.system:2222",
				OAuthClient: "ssh-proxy",
			},
			expectedResponse: `{
				"links": {
					"app_ssh": {
						"href": "ssh://ssh.apps.system:2222",
						"meta": {
							"host_key_fingerprint": "",
							"oauth_client": "ssh-proxy"
						}
					},
					"bits_service": null,
					"cloud_controller_v2": {
						"href": "https://example.org/v2",
						"meta": {
							"version": "2.171.0"
						}
					},
					"cloud_controller_v3": {
						"href": "https://example.org/v3",
						"meta": {
							"version": "3.106.0"
						}
					},
					"credhub": null,
					"log_cache": {
						"href": "https://log-cache.logging.system"
					},
					"log_stream": {
						"href": "https://log-stream.logging.system"
					},
					"logging": {
						"href": "wss://doppler.logging.system:443"
					},
					"login": {
						"href": "https://login.identity.system"
					},
					"network_policy_v0": {
						"href": "https://example.org/networking/v0/external"
					},
					"network_policy_v1": {
						"href": "https://example.org/networking/v1/external"
					},
					"routing": null,
					"self": {
						"href": "https://example.org"
					},
					"uaa": {
						"href": "https://uaa.identity.system"
					}
				}
			}`,
		},
		"with optional config": {
			externalDomain:   "example.org",
			externalProtocol: "https",
			urls: config.URLs{
				LogCache:   "https://log-cache.logging.system",
				LogStream:  "https://log-stream.logging.system",
				Doppler:    "wss://doppler.logging.system:443",
				Login:      "https://login.identity.system",
				UAA:        "https://uaa.identity.system",
				RoutingAPI: toPtr("https://api.routing.system"),
			},
			appSSH: config.AppSSH{
				Endpoint:    "ssh://ssh.apps.system:2222",
				OAuthClient: "ssh-proxy",
			},
			expectedResponse: `{
				"links": {
					"app_ssh": {
						"href": "ssh://ssh.apps.system:2222",
						"meta": {
							"host_key_fingerprint": "",
							"oauth_client": "ssh-proxy"
						}
					},
					"bits_service": null,
					"cloud_controller_v2": {
						"href": "https://example.org/v2",
						"meta": {
							"version": "2.171.0"
						}
					},
					"cloud_controller_v3": {
						"href": "https://example.org/v3",
						"meta": {
							"version": "3.106.0"
						}
					},
					"credhub": null,
					"log_cache": {
						"href": "https://log-cache.logging.system"
					},
					"log_stream": {
						"href": "https://log-stream.logging.system"
					},
					"logging": {
						"href": "wss://doppler.logging.system:443"
					},
					"login": {
						"href": "https://login.identity.system"
					},
					"network_policy_v0": {
						"href": "https://example.org/networking/v0/external"
					},
					"network_policy_v1": {
						"href": "https://example.org/networking/v1/external"
					},
					"routing": {
						"href": "https://api.routing.system"
					},
					"self": {
						"href": "https://example.org"
					},
					"uaa": {
						"href": "https://uaa.identity.system"
					}
				}
			}`,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := Controller{
				ExternalDomain:   tc.externalDomain,
				ExternalProtocol: tc.externalProtocol,
				URLs:             tc.urls,
				AppSSH:           tc.appSSH,
			}
			err := controller.GetRoot(c)

			assert.NoError(t, err)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}

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
	c := e.NewContext(req, rec)

	controller := Controller{ExternalDomain: "api.external.domain", ExternalProtocol: "https"}
	err := controller.GetV3Root(c)

	assert.NoError(t, err)
	assert.JSONEq(t, expectedResponse, rec.Body.String())
}

func toPtr(s string) *string {
	return &s
}
