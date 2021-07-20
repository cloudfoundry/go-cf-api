package presenter

import (
	"time"

	psqlModels "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/sqlboiler/postgres"
)

type BuildpackResponse struct {
	Guid       string
	Created_at time.Time
	Updated_at time.Time
	Name       string
	State      string
	Filename   string
	Stack      string
	Position   int64
	Enabled    bool
	Locked     bool
	Metadata   Metadata
	Links      struct {
		Self struct {
			Href string
		}
		Upload struct {
			Href   string
			Method string
		}
	}
}

func BuildpacksResponseObject(buildpacks psqlModels.BuildpackSlice) []*BuildpackResponse {
	out := []*BuildpackResponse{}
	for _, buildpack := range buildpacks {
		buildpackresp := BuildpackResponse{
			Guid:       buildpack.GUID,
			Created_at: buildpack.CreatedAt,
		}
		out = append(out, &buildpackresp)
	}

	return out
}
