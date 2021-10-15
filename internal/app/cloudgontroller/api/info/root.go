package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cont *Controller) GetRoot(c echo.Context) error {
	root := Root{
		Links: RootLinks{
			AppSSH: &Link{
				HREF: cont.AppSSH.Endpoint,
				Meta: map[string]string{
					"host_key_fingerprint": cont.AppSSH.HostKeyFingerprint,
					"oauth_client":         cont.AppSSH.OAuthClient,
				},
			},
			// Bits Service is deprecated. It is included here for API compatibility
			BitsService: nil,
			CloudControllerV2: &Link{
				HREF: cont.externalURL("/v2"),
				Meta: map[string]string{"version": "2.171.0"},
			},
			CloudControllerV3: &Link{
				HREF: cont.externalURL("/v3"),
				Meta: map[string]string{"version": "3.106.0"},
			},
			// We do not currently integrate with CredHub
			Credhub:         nil,
			LogCache:        &Link{HREF: cont.URLs.LogCache},
			LogStream:       &Link{HREF: cont.URLs.LogStream},
			Logging:         &Link{HREF: cont.URLs.Doppler},
			Login:           &Link{HREF: cont.URLs.Login},
			NetworkPolicyV0: &Link{HREF: cont.externalURL("/networking/v0/external")},
			NetworkPolicyV1: &Link{HREF: cont.externalURL("/networking/v1/external")},
			Routing:         cont.routing(),
			Self:            &Link{HREF: cont.externalURL("")},
			UAA:             &Link{HREF: cont.URLs.UAA},
		},
	}
	return c.JSONPretty(http.StatusOK, root, indent)
}

type Root struct {
	Links RootLinks `json:"links"`
}

type RootLinks struct {
	AppSSH            *Link `json:"app_ssh"`
	BitsService       *Link `json:"bits_service"`
	CloudControllerV2 *Link `json:"cloud_controller_v2"`
	CloudControllerV3 *Link `json:"cloud_controller_v3"`
	Credhub           *Link `json:"credhub"`
	LogCache          *Link `json:"log_cache"`
	LogStream         *Link `json:"log_stream"`
	Logging           *Link `json:"logging"`
	Login             *Link `json:"login"`
	NetworkPolicyV0   *Link `json:"network_policy_v0"`
	NetworkPolicyV1   *Link `json:"network_policy_v1"`
	Routing           *Link `json:"routing"`
	Self              *Link `json:"self"`
	UAA               *Link `json:"uaa"`
}

func (cont *Controller) routing() *Link {
	var link *Link
	if cont.URLs.RoutingAPI != nil {
		link = &Link{HREF: *cont.URLs.RoutingAPI}
	}
	return link
}
