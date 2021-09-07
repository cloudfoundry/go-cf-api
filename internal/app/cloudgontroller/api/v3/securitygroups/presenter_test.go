// +build unit

package securitygroups_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/securitygroups"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

func TestSecurityGroupResponseObject(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		securityGroup    models.SecurityGroup
		stagingSpaces    map[int]string
		runningSpaces    map[int]string
		expectedResponse *Response
	}{
		"security_group valid response": {
			securityGroup: models.SecurityGroup{
				GUID:           "123",
				CreatedAt:      time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC),
				UpdatedAt:      null.Time{Time: time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC)},
				Name:           "TestSecurityGroup",
				Rules:          null.StringFrom(`[{"protocol":"tcp","destination":"10.10.10.0/24","ports":"443,80,8080"}]`),
				StagingDefault: null.BoolFrom(false),
				RunningDefault: null.BoolFrom(false),
			},
			stagingSpaces: map[int]string{123: "space-guid-1"},
			runningSpaces: map[int]string{456: "space-guid-2"},
			expectedResponse: &Response{
				GUID:      "123",
				CreatedAt: "2021-08-16T13:14:15Z", UpdatedAt: "2021-08-16T13:14:15Z",
				Name:            "TestSecurityGroup",
				Rules:           []Rule{{Protocol: "tcp", Destination: "10.10.10.0/24", Ports: "443,80,8080"}},
				GloballyEnabled: map[string]bool{"running": false, "staging": false},
				Links: Links{
					Self: pagination.Link{Href: "v3/security_groups/123"},
				},
				Relationships: map[string]RelationshipData{
					"staging_spaces": {Data: []SecurityGroupSpace{{GUID: "space-guid-1"}}},
					"running_spaces": {Data: []SecurityGroupSpace{{GUID: "space-guid-2"}}},
				},
			},
		},
		"no staging or running spaces": {
			securityGroup: models.SecurityGroup{
				GUID:           "123",
				CreatedAt:      time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC),
				UpdatedAt:      null.Time{Time: time.Date(2021, 8, 16, 13, 14, 15, 16, time.UTC)},
				Name:           "TestSecurityGroup",
				Rules:          null.StringFrom(`[{"protocol":"tcp","destination":"10.10.10.0/24","ports":"443,80,8080"}]`),
				StagingDefault: null.BoolFrom(false),
				RunningDefault: null.BoolFrom(false),
			},
			stagingSpaces: map[int]string{},
			runningSpaces: map[int]string{},
			expectedResponse: &Response{
				GUID:      "123",
				CreatedAt: "2021-08-16T13:14:15Z", UpdatedAt: "2021-08-16T13:14:15Z",
				Name:            "TestSecurityGroup",
				Rules:           []Rule{{Protocol: "tcp", Destination: "10.10.10.0/24", Ports: "443,80,8080"}},
				GloballyEnabled: map[string]bool{"running": false, "staging": false},
				Links: Links{
					Self: pagination.Link{Href: "v3/security_groups/123"},
				},
				Relationships: map[string]RelationshipData{
					"staging_spaces": {Data: []SecurityGroupSpace{}},
					"running_spaces": {Data: []SecurityGroupSpace{}},
				},
			},
		},
	}

	p := NewPresenter()
	for name, tc := range cases {
		r := tc.securityGroup.R.NewStruct()
		for spaceID, spaceGUID := range tc.stagingSpaces {
			sr := models.StagingSecurityGroupsSpace{}.R.NewStruct()
			sr.StagingSpace = &models.Space{GUID: spaceGUID}
			ssgs := &models.StagingSecurityGroupsSpace{StagingSpaceID: spaceID, R: sr}
			r.StagingSecurityGroupStagingSecurityGroupsSpaces = append(r.StagingSecurityGroupStagingSecurityGroupsSpaces, ssgs)
		}
		for spaceID, spaceGUID := range tc.runningSpaces {
			sr := models.SecurityGroupsSpace{}.R.NewStruct()
			sr.Space = &models.Space{GUID: spaceGUID}
			sgs := &models.SecurityGroupsSpace{SpaceID: spaceID, R: sr}
			r.SecurityGroupsSpaces = append(r.SecurityGroupsSpaces, sgs)
		}
		tc.securityGroup.R = r
		t.Run(name, func(t *testing.T) {
			response, err := p.ResponseObject(&tc.securityGroup, "v3/security_groups/123")
			require.NoError(t, err)
			require.Equal(t, tc.expectedResponse, response)
		})
	}
}
