//go:build unit
// +build unit

package securitygroups_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	. "github.com/cloudfoundry/go-cf-api/internal/api/v3/securitygroups"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/pagination"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
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

	presenter := NewPresenter()
	for testCaseName, testCase := range cases {
		sgR := testCase.securityGroup.R.NewStruct()
		for spaceID, spaceGUID := range testCase.stagingSpaces {
			sr := models.StagingSecurityGroupsSpace{}.R.NewStruct()
			sr.StagingSpace = &models.Space{GUID: spaceGUID}
			ssgs := &models.StagingSecurityGroupsSpace{StagingSpaceID: spaceID, R: sr}
			sgR.StagingSecurityGroupStagingSecurityGroupsSpaces = append(sgR.StagingSecurityGroupStagingSecurityGroupsSpaces, ssgs)
		}
		for spaceID, spaceGUID := range testCase.runningSpaces {
			sr := models.SecurityGroupsSpace{}.R.NewStruct()
			sr.Space = &models.Space{GUID: spaceGUID}
			sgs := &models.SecurityGroupsSpace{SpaceID: spaceID, R: sr}
			sgR.SecurityGroupsSpaces = append(sgR.SecurityGroupsSpaces, sgs)
		}
		testCase.securityGroup.R = sgR
		t.Run(testCaseName, func(t *testing.T) {
			response, err := presenter.ResponseObject(&testCase.securityGroup, "v3/security_groups/123")
			require.NoError(t, err)
			require.Equal(t, testCase.expectedResponse, response)
		})
	}
}
