package info

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

const indent = " "

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

func (cont *Controller) GetV3Info(c echo.Context) error {
	v3Info := V3Info{
		Name:        cont.Info.Name,
		Description: cont.Info.Description,
		Build:       cont.Info.Build,
		Version:     cont.Info.Version,
		CLIVersion: CLIVersion{
			Minimum:     cont.Info.MinCliVersion,
			Recommended: cont.Info.MinRecommendedCliVersion,
		},
		Links: V3Links{
			Self:    Link{HREF: cont.externalURL("/v3/info")},
			Support: Link{HREF: cont.Info.SupportAddress},
		},
	}
	return c.JSONPretty(http.StatusOK, v3Info, indent)
}

type Controller struct {
	Info             config.InfoConfig
	URLs             config.URLs
	AppSSH           config.AppSSH
	ExternalDomain   string
	ExternalProtocol string
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

type V3Info struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Build       string     `json:"build"`
	Version     int        `json:"version"`
	CLIVersion  CLIVersion `json:"cli_version"`
	Links       V3Links    `json:"links"`
}

type CLIVersion struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type V3Links struct {
	Self    Link `json:"self"`
	Support Link `json:"support"`
}

type Link struct {
	HREF string            `json:"href"`
	Meta map[string]string `json:"meta,omitempty"`
}

func (cont *Controller) externalURL(endpoint string) string {
	url := url.URL{
		Scheme: cont.ExternalProtocol,
		Host:   cont.ExternalDomain,
		Path:   endpoint,
	}
	return url.String()
}

func (cont *Controller) routing() *Link {
	var link *Link
	if cont.URLs.RoutingAPI != nil {
		link = &Link{HREF: *cont.URLs.RoutingAPI}
	}
	return link
}
