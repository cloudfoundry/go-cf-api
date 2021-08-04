package presenter

import (
	"fmt"
	"github.com/volatiletech/null/v8"
	"time"

	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

const STATE_AWAITING_UPLOAD = "AWAITING_UPLOAD"
const STATE_READY = "READY"

type BuildpackResponse struct {
	Guid       string      `json:"guid"`
	Created_at time.Time   `json:"created_at"`
	Updated_at null.Time   `json:"updated_at"`
	Name       string      `json:"name"`
	State      string      `json:"state"`
	Filename   null.String `json:"filename"`
	Stack      null.String `json:"stack"`
	Position   int         `json:"position"`
	Enabled    null.Bool   `json:"enabled"`
	Locked     null.Bool   `json:"locked"`
	Metadata   Metadata    `json:"metadata"`
	Links      struct {
		Self   Link `json:"self"`
		Upload Link `json:"upload"`
	} `json:"links"`
}

type BuildpacksResponse struct {
	Pagination *Pagination          `json:"pagination"`
	Resources  []*BuildpackResponse `json:"resources"`
}

func BuildpackResponseObject(buildpack *models.Buildpack, resourcePath string) *BuildpackResponse {
	response := &BuildpackResponse{
		Guid:       buildpack.GUID,
		Created_at: buildpack.CreatedAt,
		Updated_at: buildpack.UpdatedAt,
		Name:       buildpack.Name,
		State:      GetBuildpackState(buildpack),
		Filename:   buildpack.Filename,
		Stack:      buildpack.Stack,
		Position:   buildpack.Position,
		Enabled:    buildpack.Enabled,
		Locked:     buildpack.Locked,
	}
	response.Links.Self = GetResourcePathLink(resourcePath)
	response.Links.Upload = GetResourcePathLinkWithMethod(fmt.Sprintf("%s/%s", resourcePath, "upload"), "POST")
	return response
}

func BuildpacksResponseObject(buildpacks models.BuildpackSlice, resourcePath string) *BuildpacksResponse {
	out := []*BuildpackResponse{}
	for _, buildpack := range buildpacks {
		buildpackresp := BuildpackResponseObject(buildpack, fmt.Sprintf("%s/%s", resourcePath, buildpack.GUID))
		out = append(out, buildpackresp)
	}
	return &BuildpacksResponse{
		Pagination: &Pagination{
			TotalResults: len(out),
			TotalPages:   1,
			First:        GetResourcePathLink(resourcePath),
			Last:         GetResourcePathLink(resourcePath),
			Next:         GetResourcePathLink(resourcePath),
			Previous:     GetResourcePathLink(resourcePath),
		},
		Resources: out,
	}
}

func GetResourcePathLink(resourcePath string) Link {
	return Link{
		Href: resourcePath,
	}
}

func GetResourcePathLinkWithMethod(resourcePath string, method string) Link {
	link := GetResourcePathLink(resourcePath)
	link.Method = method
	return link
}

func GetBuildpackState(buildpack *models.Buildpack) string {
	if buildpack.Filename.IsZero() {
		return STATE_AWAITING_UPLOAD
	}
	return STATE_READY
}
