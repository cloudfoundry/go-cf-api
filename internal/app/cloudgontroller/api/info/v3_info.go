package info

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
	return c.JSON(http.StatusOK, v3Info)
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
