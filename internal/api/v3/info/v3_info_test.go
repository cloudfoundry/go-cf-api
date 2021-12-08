//go:build unit
// +build unit

package info_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudfoundry/go-cf-api/internal/api/v3/info"
	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
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

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/info", nil)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			handleFunc := info.NewV3InfoEndpoint(&config.CfAPIConfig{
				ExternalDomain:   testCase.externalDomain,
				ExternalProtocol: testCase.externalProtocol,
				Info:             testCase.infoConfig,
			})
			err := handleFunc(ctx)

			assert.NoError(t, err)
			assert.JSONEq(t, testCase.expectedResponse, rec.Body.String())
		})
	}
}
