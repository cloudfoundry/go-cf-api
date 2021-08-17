package presenter

import (
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

const (
	StateAwaitingUpload = "AWAITING_UPLOAD"
	StateReady          = "READY"
)

type BuildpackResponse struct {
	GUID      string      `json:"guid"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at"`
	Name      string      `json:"name"`
	Stack     null.String `json:"stack"`
	State     string      `json:"state"`
	Filename  null.String `json:"filename"`
	Position  int         `json:"position"`
	Enabled   null.Bool   `json:"enabled"`
	Locked    null.Bool   `json:"locked"`
	Metadata  Metadata    `json:"metadata"`
	Links     struct {
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
		GUID:      buildpack.GUID,
		CreatedAt: buildpack.CreatedAt.Format(time.RFC3339),
		UpdatedAt: buildpack.UpdatedAt.Time.Format(time.RFC3339),
		Name:      buildpack.Name,
		State:     GetBuildpackState(buildpack),
		Filename:  buildpack.Filename,
		Stack:     buildpack.Stack,
		Position:  buildpack.Position,
		Enabled:   buildpack.Enabled,
		Locked:    buildpack.Locked,
		// Workaround until metadata is implemented
		Metadata: Metadata{Annotations: map[string]string{}, Labels: map[string]string{}},
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
		return StateAwaitingUpload
	}
	return StateReady
}
