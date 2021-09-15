// +build unit

//nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions
package security_groups

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

func TestSecurityGroupResponseObject(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		securityGroup    models.SecurityGroup
		stagingSpaceIDs  []int
		runningSpaceIDs  []int
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
			stagingSpaceIDs: []int{123},
			runningSpaceIDs: []int{456},
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
					"staging_spaces": {Data: []SecurityGroupSpace{{Guid: "123"}}},
					"running_spaces": {Data: []SecurityGroupSpace{{Guid: "456"}}},
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
			stagingSpaceIDs: []int{},
			runningSpaceIDs: []int{},
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
	var p Presenter = &presenter{}

	for name, tc := range cases {
		r := tc.securityGroup.R.NewStruct()
		for _, stagingSpaceID := range tc.stagingSpaceIDs {
			r.StagingSecurityGroupStagingSecurityGroupsSpaces = append(
				r.StagingSecurityGroupStagingSecurityGroupsSpaces,
				&models.StagingSecurityGroupsSpace{StagingSpaceID: stagingSpaceID},
			)
		}
		for _, runningSpaceID := range tc.runningSpaceIDs {
			r.SecurityGroupsSpaces = append(
				r.SecurityGroupsSpaces,
				&models.SecurityGroupsSpace{SpaceID: runningSpaceID},
			)
		}
		tc.securityGroup.R = r
		t.Run(name, func(t *testing.T) {
			response, err := p.ResponseObject(&tc.securityGroup, "v3/security_groups/123")
			require.NoError(t, err)
			require.Equal(t, tc.expectedResponse, response)
		})
	}
}
