package info

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

type Controller struct {
	ExternalDomain   string
	ExternalProtocol string
	Config           config.InfoConfig
}

type V3Info struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Build       string     `json:"build"`
	Version     int        `json:"version"`
	CLIVersion  CLIVersion `json:"cli_version"`
	Links       Links      `json:"links"`
}

type CLIVersion struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type Links struct {
	Self    Link `json:"self"`
	Support Link `json:"support"`
}

type Link struct {
	HREF string `json:"href"`
}

func (cont *Controller) GetV3Info(c echo.Context) error {
	v3Info := V3Info{
		Name:        cont.Config.Name,
		Description: cont.Config.Description,
		Build:       cont.Config.Build,
		Version:     cont.Config.Version,
		CLIVersion: CLIVersion{
			Minimum:     cont.Config.MinCliVersion,
			Recommended: cont.Config.MinRecommendedCliVersion,
		},
		Links: Links{
			Self:    cont.link("/v3/info"),
			Support: Link{HREF: cont.Config.SupportAddress},
		},
	}
	return c.JSONPretty(http.StatusOK, v3Info, "  ")
}

func (cont *Controller) link(endpoint string) Link {
	url := url.URL{
		Scheme: cont.ExternalProtocol,
		Host:   cont.ExternalDomain,
		Path:   endpoint,
	}
	return Link{HREF: url.String()}
}
