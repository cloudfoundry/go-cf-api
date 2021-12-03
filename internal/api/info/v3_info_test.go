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
