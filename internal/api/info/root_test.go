// +build unit

package info_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/api/info"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
)

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

func toPtr(s string) *string {
	return &s
}
