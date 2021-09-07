package buildpacks

import (
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/metadata"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
)

const (
	StateAwaitingUpload = "AWAITING_UPLOAD"
	StateReady          = "READY"
)

type Response struct {
	GUID      string            `json:"guid"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
	Name      string            `json:"name"`
	Stack     null.String       `json:"stack"`
	State     string            `json:"state"`
	Filename  null.String       `json:"filename"`
	Position  int               `json:"position"`
	Enabled   null.Bool         `json:"enabled"`
	Locked    null.Bool         `json:"locked"`
	Metadata  metadata.Metadata `json:"metadata"`
	Links     struct {
		Self   pagination.Link `json:"self"`
		Upload pagination.Link `json:"upload"`
	} `json:"links"`
}

type ListResponse struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Resources  []*Response            `json:"resources"`
}

func ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error) {
	md := metadata.Metadata{}
	var err error
	if buildpack.R == nil {
		zap.L().Warn(fmt.Sprintf("Buildpack with guid %s does not contain metadata", buildpack.GUID))
	} else {
		md, err = metadata.Get(buildpack.R.ResourceBuildpackAnnotations, buildpack.R.ResourceBuildpackLabels)
		if err != nil {
			return nil, fmt.Errorf("cannot build response object: %s", err)
		}
	}
	response := &Response{
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
		Metadata:  md,
	}
	response.Links.Self = GetResourcePathLink(resourcePath)
	response.Links.Upload = GetResourcePathLinkWithMethod(fmt.Sprintf("%s/%s", resourcePath, "upload"), "POST")
	return response, nil
}

func ListResponseObject(
	buildpacks models.BuildpackSlice,
	paginationParams pagination.Params,
	resourcePath string) (*ListResponse, error) {
	out := []*Response{}
	for _, buildpack := range buildpacks {
		buildpackresp, err := ResponseObject(buildpack, fmt.Sprintf("%s/%s", resourcePath, buildpack.GUID))
		if err != nil {
			return nil, err
		}
		out = append(out, buildpackresp)
	}

	return &ListResponse{
		Pagination: pagination.NewPagination(len(buildpacks), paginationParams, resourcePath),
		Resources:  out,
	}, nil
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
