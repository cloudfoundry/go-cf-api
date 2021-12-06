package api

import (
	"github.com/labstack/echo/v4"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/info"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
	"net/http"
)

func NewRootEndpoint(config *config.CloudgontrollerConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		root := Root{
			Links: RootLinks{
				AppSSH: &info.Link{
					HREF: config.AppSSH.Endpoint,
					Meta: map[string]string{
						"host_key_fingerprint": config.AppSSH.HostKeyFingerprint,
						"oauth_client":         config.AppSSH.OAuthClient,
					},
				},
				// Bits Service is deprecated. It is included here for API compatibility
				BitsService: nil,
				CloudControllerV2: &info.Link{
					HREF: info.ExternalURL("/v2", config),
					Meta: map[string]string{"version": "2.171.0"},
				},
				CloudControllerV3: &info.Link{
					HREF: info.ExternalURL("/v3", config),
					Meta: map[string]string{"version": "3.106.0"},
				},
				// We do not currently integrate with CredHub
				Credhub:         nil,
				LogCache:        &info.Link{HREF: config.URLs.LogCache},
				LogStream:       &info.Link{HREF: config.URLs.LogStream},
				Logging:         &info.Link{HREF: config.URLs.Doppler},
				Login:           &info.Link{HREF: config.URLs.Login},
				NetworkPolicyV0: &info.Link{HREF: info.ExternalURL("/networking/v0/external", config)},
				NetworkPolicyV1: &info.Link{HREF: info.ExternalURL("/networking/v1/external", config)},
				Routing:         routingLink(config),
				Self:            &info.Link{HREF: info.ExternalURL("", config)},
				UAA:             &info.Link{HREF: config.URLs.UAA},
			},
		}
		return c.JSON(http.StatusOK, root)
	}
}

func routingLink(config *config.CloudgontrollerConfig) *info.Link {
	var link *info.Link
	if config.URLs.RoutingAPI != nil {
		link = &info.Link{HREF: *config.URLs.RoutingAPI}
	}
	return link
}

type Root struct {
	Links RootLinks `json:"links"`
}

type RootLinks struct {
	AppSSH            *info.Link `json:"app_ssh"`
	BitsService       *info.Link `json:"bits_service"`
	CloudControllerV2 *info.Link `json:"cloud_controller_v2"`
	CloudControllerV3 *info.Link `json:"cloud_controller_v3"`
	Credhub           *info.Link `json:"credhub"`
	LogCache          *info.Link `json:"log_cache"`
	LogStream         *info.Link `json:"log_stream"`
	Logging           *info.Link `json:"logging"`
	Login             *info.Link `json:"login"`
	NetworkPolicyV0   *info.Link `json:"network_policy_v0"`
	NetworkPolicyV1   *info.Link `json:"network_policy_v1"`
	Routing           *info.Link `json:"routing"`
	Self              *info.Link `json:"self"`
	UAA               *info.Link `json:"uaa"`
}
