package securitygroups

import (
	"encoding/json"
	"fmt"
	"time"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

type Response struct {
	GUID            string                      `json:"guid"`
	CreatedAt       string                      `json:"created_at"`
	UpdatedAt       string                      `json:"updated_at"`
	Name            string                      `json:"name"`
	Rules           []Rule                      `json:"rules"`
	GloballyEnabled map[string]bool             `json:"globally_enabled"`
	Links           Links                       `json:"links"`
	Relationships   map[string]RelationshipData `json:"relationships"`
}

type Links struct {
	Self pagination.Link `json:"self"`
}

type RelationshipData struct {
	Data []SecurityGroupSpace `json:"data"`
}

type SecurityGroupSpace struct {
	GUID string `json:"guid"`
}

type ListResponse struct {
	Pagination *pagination.Pagination `json:"pagination"`
	Resources  []*Response            `json:"resources"`
}

type Presenter interface {
	ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error)
	ListResponseObject(
		securityGroups models.SecurityGroupSlice,
		totalResults int64,
		paginationParams pagination.Params,
		resourcePath string) (*ListResponse, error)
}

type presenter struct{}

func (p *presenter) ListResponseObject(securityGroups models.SecurityGroupSlice,
	totalResults int64,
	paginationParams pagination.Params,
	resourcePath string) (*ListResponse, error) {
	out := []*Response{}
	for _, securitygroup := range securityGroups {
		securityGroupResp, err := p.ResponseObject(securitygroup, fmt.Sprintf("%s/%s", resourcePath, securitygroup.GUID))
		if err != nil {
			return nil, err
		}
		out = append(out, securityGroupResp)
	}

	return &ListResponse{
		Pagination: pagination.NewPagination(totalResults, paginationParams, resourcePath),
		Resources:  out,
	}, nil
}

func NewPresenter() Presenter {
	return &presenter{}
}

type Rule struct {
	Protocol    string `json:"protocol"`
	Destination string `json:"destination"`
	Ports       string `json:"ports,omitempty"`
	Type        int    `json:"type,omitempty"`
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Log         bool   `json:"log,omitempty"`
}

func (p *presenter) ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error) {
	var rules []Rule
	if err := json.Unmarshal([]byte(securityGroup.Rules.String), &rules); err != nil {
		return nil, err
	}

	stagingSpaces := []SecurityGroupSpace{}
	for _, space := range securityGroup.R.StagingSecurityGroupStagingSecurityGroupsSpaces {
		spaceData := SecurityGroupSpace{
			GUID: space.R.StagingSpace.GUID,
		}
		stagingSpaces = append(stagingSpaces, spaceData)
	}

	runningSpaces := []SecurityGroupSpace{}
	for _, space := range securityGroup.R.SecurityGroupsSpaces {
		spaceData := SecurityGroupSpace{
			GUID: space.R.Space.GUID,
		}
		runningSpaces = append(runningSpaces, spaceData)
	}

	response := &Response{
		GUID:      securityGroup.GUID,
		CreatedAt: securityGroup.CreatedAt.Format(time.RFC3339),
		UpdatedAt: securityGroup.UpdatedAt.Time.Format(time.RFC3339),
		Name:      securityGroup.Name,
		Rules:     rules,
		GloballyEnabled: map[string]bool{
			"running": securityGroup.RunningDefault.Bool,
			"staging": securityGroup.StagingDefault.Bool,
		},
		Relationships: map[string]RelationshipData{
			"staging_spaces": {
				Data: stagingSpaces,
			},
			"running_spaces": {
				Data: runningSpaces,
			},
		},
	}
	response.Links.Self = pagination.GetResourcePathLink(resourcePath)
	return response, nil
}
