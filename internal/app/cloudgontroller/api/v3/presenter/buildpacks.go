package presenter

import (
	"time"

	psqlModels "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/postgres"
)

type BuildpackResponse struct {
	Guid       string    `json:"guid"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	State      string    `json:"state"`
	Filename   string    `json:"filename"`
	Stack      string    `json:"stack"`
	Position   int       `json:"position"`
	Enabled    bool      `json:"enabled"`
	Locked     bool      `json:"locked"`
	Metadata   Metadata  `json:"metadata"`
	Links      struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Upload struct {
			Href   string `json:"href"`
			Method string `json:"method"`
		} `json:"upload"`
	} `json:"links"`
}

func BuildpackResponseObject(buildpack *psqlModels.Buildpack) *BuildpackResponse {
	return &BuildpackResponse{
		Guid:       buildpack.GUID,
		Created_at: buildpack.CreatedAt,
		Name:       buildpack.Name,
		Position:   buildpack.Position,
	}
}

func BuildpacksResponseObject(buildpacks psqlModels.BuildpackSlice) []*BuildpackResponse {
	out := []*BuildpackResponse{}
	for _, buildpack := range buildpacks {
		buildpackresp := BuildpackResponseObject(buildpack)
		out = append(out, buildpackresp)
	}
	return out
}
