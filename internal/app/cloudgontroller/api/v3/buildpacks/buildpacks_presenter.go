package buildpacks

import (
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
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
	Metadata  v3.Metadata `json:"metadata"`
	Links     struct {
		Self   pagination.Link `json:"self"`
		Upload pagination.Link `json:"upload"`
	} `json:"links"`
}

type BuildpacksResponse struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Resources  []*BuildpackResponse   `json:"resources"`
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
		Metadata: v3.Metadata{Annotations: map[string]string{}, Labels: map[string]string{}},
	}
	response.Links.Self = GetResourcePathLink(resourcePath)
	response.Links.Upload = GetResourcePathLinkWithMethod(fmt.Sprintf("%s/%s", resourcePath, "upload"), "POST")
	return response
}

func BuildpacksResponseObject(
	buildpacks models.BuildpackSlice,
	paginationParams pagination.Params,
	resourcePath string) *BuildpacksResponse {
	out := []*BuildpackResponse{}
	for _, buildpack := range buildpacks {
		buildpackresp := BuildpackResponseObject(buildpack, fmt.Sprintf("%s/%s", resourcePath, buildpack.GUID))
		out = append(out, buildpackresp)
	}

	return &BuildpacksResponse{
		Pagination: pagination.NewPagination(len(buildpacks), paginationParams, resourcePath),
		Resources:  out,
	}
}

func GetResourcePathLink(resourcePath string) pagination.Link {
	return pagination.Link{
		Href: resourcePath,
	}
}

func GetResourcePathLinkWithMethod(resourcePath string, method string) pagination.Link {
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
