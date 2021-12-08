package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/cloudfoundry/go-cf-api/internal/config"
)

func NewV3InfoEndpoint(config *config.CfAPIConfig) echo.HandlerFunc {
	return func(ctx echo.Context) error {
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
		return ctx.JSON(http.StatusOK, v3Info)
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
