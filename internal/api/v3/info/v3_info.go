package info

import (
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewV3InfoEndpoint(config *config.CloudgontrollerConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		v3Info := V3Info{
			Name:        config.Info.Name,
			Description: config.Info.Description,
			Build:       config.Info.Build,
			Version:     config.Info.Version,
			CLIVersion: CLIVersion{
				Minimum:     config.Info.MinCliVersion,
				Recommended: config.Info.MinRecommendedCliVersion,
			},
			Links: V3Links{
				Self:    Link{HREF: ExternalURL("/v3/info", config)},
				Support: Link{HREF: config.Info.SupportAddress},
			},
		}
		return c.JSON(http.StatusOK, v3Info)
	}
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
