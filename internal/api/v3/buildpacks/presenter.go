package buildpacks

import (
	"fmt"
	"time"

	"github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/metadata"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
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

type presenter struct{}

func NewPresenter() Presenter {
	return &presenter{}
}

type Presenter interface {
	ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error)
	ListResponseObject(
		buildpacks models.BuildpackSlice,
		totalResults int64,
		paginationParams pagination.Params,
		resourcePath string) (*ListResponse, error)
}

func (p *presenter) ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error) {
	meta := metadata.Metadata{}
	var err error
	if buildpack.R == nil {
		zap.L().Warn(fmt.Sprintf("Buildpack with guid %s does not contain metadata", buildpack.GUID))
	} else {
		meta, err = metadata.Get(buildpack.R.ResourceBuildpackAnnotations, buildpack.R.ResourceBuildpackLabels)
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
		Metadata:  meta,
	}
	response.Links.Self = pagination.GetResourcePathLink(resourcePath)
	response.Links.Upload = pagination.GetResourcePathLinkWithMethod(fmt.Sprintf("%s/%s", resourcePath, "upload"), "POST")
	return response, nil
}

func (p *presenter) ListResponseObject(
	buildpacks models.BuildpackSlice,
	totalResults int64,
	paginationParams pagination.Params,
	resourcePath string) (*ListResponse, error) {
	out := []*Response{}
	for _, buildpack := range buildpacks {
		buildpackresp, err := p.ResponseObject(buildpack, fmt.Sprintf("%s/%s", resourcePath, buildpack.GUID))
		if err != nil {
			return nil, err
		}
		out = append(out, buildpackresp)
	}

	return &ListResponse{
		Pagination: pagination.NewPagination(totalResults, paginationParams, resourcePath),
		Resources:  out,
	}, nil
}

func GetBuildpackState(buildpack *models.Buildpack) string {
	if buildpack.Filename.IsZero() {
		return StateAwaitingUpload
	}
	return StateReady
}
