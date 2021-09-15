package security_groups

import (
	"encoding/json"
	"strconv"
	"time"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
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
	Guid string `json:"guid"`
}
type Presenter interface {
	ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error)
}

type presenter struct{}

func NewPresenter() Presenter {
	return &presenter{}
}

type Rule struct {
	Protocol    string `json:"protocol"`
	Destination string `json:"destination"`
	Ports       string `json:"ports"`
	Type        int    `json:"type"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	Log         bool   `json:"log"`
}

func (p *presenter) ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error) {
	var rules []Rule
	if err := json.Unmarshal([]byte(securityGroup.Rules.String), &rules); err != nil {
		return nil, err
	}

	stagingSpaces := []SecurityGroupSpace{}
	for _, space := range securityGroup.R.StagingSecurityGroupStagingSecurityGroupsSpaces {
		var spaceData SecurityGroupSpace = SecurityGroupSpace{
			Guid: strconv.Itoa(space.StagingSpaceID),
		}
		stagingSpaces = append(stagingSpaces, spaceData)
	}

	runningSpaces := []SecurityGroupSpace{}
	for _, space := range securityGroup.R.SecurityGroupsSpaces {
		var spaceData SecurityGroupSpace = SecurityGroupSpace{
			Guid: strconv.Itoa(space.SpaceID),
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
