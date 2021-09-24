// +build integration

package securitygroups_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/securitygroups"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/permissions"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/permissions"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	dbconfig "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/testutils"
	"go.uber.org/zap/zaptest"
)

// Order matters to prevent foreign key errors
var tablesToClear = []string{
	models.TableNames.SpacesDevelopers,
	models.TableNames.SpacesAuditors,
	models.TableNames.SpacesApplicationSupporters,
	models.TableNames.SpacesManagers,
	models.TableNames.SecurityGroupsSpaces,
	models.TableNames.StagingSecurityGroupsSpaces,
	models.TableNames.Spaces,
	models.TableNames.SecurityGroups,
	models.TableNames.OrganizationsManagers,
	models.TableNames.Organizations,
	models.TableNames.QuotaDefinitions,
	models.TableNames.Users,
}

type SecurityGroupIntegrationTestSuite struct {
	testutils.DBIntegrationTestSuite

	controller *Controller
	rec        *httptest.ResponseRecorder
	c          echo.Context

	unassignedSecurityGroup      *models.SecurityGroup
	globallyRunningSecurityGroup *models.SecurityGroup
	globallyStagingSecurityGroup *models.SecurityGroup
	spaceRunningSecurityGroup    *models.SecurityGroup
	spaceStagingSecurityGroup    *models.SecurityGroup
	user                         *models.User
	quota                        *models.QuotaDefinition
	org                          *models.Organization
	runningSpace                 *models.Space
	stagingSpace                 *models.Space
	otherSpace                   *models.Space
}

func TestSecurityGroupIntegrationTest(t *testing.T) {
	suite.Run(t, new(SecurityGroupIntegrationTestSuite))
}

func (suite *SecurityGroupIntegrationTestSuite) SetupSuite() {
	conf := config.DBConfig{
		Type:             os.Getenv("CC_DB_TYPE"),
		ConnectionString: os.Getenv("CC_DB_CONNECTION_STRING"),
	}
	var err error
	suite.DB, _, err = dbconfig.NewConnection(conf, true)
	suite.Require().NoError(err)
	suite.ClearTables(tablesToClear)
	logger := zaptest.NewLogger(suite.T())
	suite.Random = rand.New(rand.NewSource(time.Now().UTC().Unix()))
	suite.DBCtx = boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	suite.controller = &Controller{
		DB:          suite.DB,
		Presenter:   NewPresenter(),
		Permissions: permissions.NewQuerier(),
	}
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/security_groups", nil)
	suite.rec = httptest.NewRecorder()
	suite.c = e.NewContext(req, suite.rec)
}

func (suite *SecurityGroupIntegrationTestSuite) SetupTest() {
	suite.ClearTables(tablesToClear)
	suite.user = suite.CreateUser()
	suite.quota = suite.CreateQuota()
	suite.org = suite.CreateOrg(suite.quota.ID)
	suite.runningSpace = suite.CreateSpace(suite.org.ID)
	suite.stagingSpace = suite.CreateSpace(suite.org.ID)
	suite.otherSpace = suite.CreateSpace(suite.org.ID)
	suite.unassignedSecurityGroup = suite.CreateSecurityGroup("unassigned", false, false)
	suite.globallyRunningSecurityGroup = suite.CreateSecurityGroup("globally-running", true, false)
	suite.globallyStagingSecurityGroup = suite.CreateSecurityGroup("globally-staging", false, true)
	suite.spaceRunningSecurityGroup = suite.CreateSecurityGroup("space-running", false, false)
	suite.AssignToRunningSpace(suite.runningSpace.ID, suite.spaceRunningSecurityGroup.ID)
	suite.spaceStagingSecurityGroup = suite.CreateSecurityGroup("space-staging", false, false)
	suite.AssignToStagingSpace(suite.stagingSpace.ID, suite.spaceStagingSecurityGroup.ID)

	suite.c.Set("username", suite.user.GUID)
}

func (suite *SecurityGroupIntegrationTestSuite) TestSpaceRolePermissions() {
	noSpace := func() *models.Space { return nil }
	assignRoles := map[Role]func(int, int){
		SpaceDeveloper: suite.AssignSpaceDeveloper,
		SpaceAuditor:   suite.AssignSpaceAuditor,
		SpaceSupporter: suite.AssignSpaceManager,
		SpaceManager:   suite.AssignSpaceManager,
	}
	cases := map[string]struct {
		spaceRoles []Role
		// We need to use func() so that when the spaces/security groups are updated in SetupTest we get the new objects
		securityGroup     func() *models.SecurityGroup
		extraSpaceForRole func() *models.Space
		expectedStatus    int
		expectedErr       *v3.CloudControllerError
	}{
		"unassigned security group should not be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.unassignedSecurityGroup },
			extraSpaceForRole: noSpace,
			expectedStatus:    http.StatusNotFound,
			expectedErr:       v3.ResourceNotFound("security_group", sql.ErrNoRows),
		},
		"globally enabled running security group should be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.globallyRunningSecurityGroup },
			extraSpaceForRole: noSpace,
			expectedStatus:    http.StatusOK,
		},
		"globally enabled staging security group should be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.globallyStagingSecurityGroup },
			extraSpaceForRole: noSpace,
			expectedStatus:    http.StatusOK,
		},
		"security group bound to running space they have no role in should not be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.spaceRunningSecurityGroup },
			extraSpaceForRole: noSpace,
			expectedStatus:    http.StatusNotFound,
			expectedErr:       v3.ResourceNotFound("security_group", sql.ErrNoRows),
		},
		"security group bound to staging space they have a role in should not be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.spaceStagingSecurityGroup },
			extraSpaceForRole: noSpace,
			expectedStatus:    http.StatusNotFound,
			expectedErr:       v3.ResourceNotFound("security_group", sql.ErrNoRows),
		},
		"security group bound to running space they have a role in should be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.spaceRunningSecurityGroup },
			extraSpaceForRole: func() *models.Space { return suite.runningSpace },
			expectedStatus:    http.StatusOK,
		},
		"security group bound to staging space they are a developer in should be readable": {
			spaceRoles:        []Role{SpaceDeveloper, SpaceAuditor, SpaceSupporter, SpaceManager},
			securityGroup:     func() *models.SecurityGroup { return suite.spaceStagingSecurityGroup },
			extraSpaceForRole: func() *models.Space { return suite.stagingSpace },
			expectedStatus:    http.StatusOK,
		},
	}
	for name, tc := range cases {
		for _, spaceRole := range tc.spaceRoles {
			suite.Run(fmt.Sprintf("for a user with %s role a %s", spaceRole, name), func() {
				suite.SetupTest() // needed to ensure DB is fresh for every test case
				assignRole := assignRoles[spaceRole]
				assignRole(suite.user.ID, suite.otherSpace.ID)
				if tc.extraSpaceForRole() != nil {
					assignRole(suite.user.ID, tc.extraSpaceForRole().ID)
				}
				suite.c.SetParamNames(GUIDParam)
				suite.c.SetParamValues(tc.securityGroup().GUID)

				err := suite.controller.Get(suite.c)
				if tc.expectedErr != nil {
					var ccErr *v3.CloudControllerError
					suite.ErrorAs(err, &ccErr)
					suite.Equal(tc.expectedErr.HTTPStatus, ccErr.HTTPStatus)
					suite.Equal(tc.expectedErr.Code, ccErr.Code)
					return
				}

				suite.NoError(err)
				suite.Equal(tc.expectedStatus, suite.rec.Result().StatusCode)
				suite.Contains(suite.rec.Body.String(), tc.securityGroup().Name)
			})
		}
	}
}

func (suite *SecurityGroupIntegrationTestSuite) CreateSecurityGroup(name string, globallyEnabledRunning, globallyEnabledStaging bool) *models.SecurityGroup {
	sg := &models.SecurityGroup{
		GUID:           name,
		Name:           fmt.Sprintf("%s-%d", name, suite.Random.Int()),
		StagingDefault: null.BoolFrom(globallyEnabledStaging),
		RunningDefault: null.BoolFrom(globallyEnabledRunning),
		Rules:          null.StringFrom("[]"),
	}
	err := models.SecurityGroups().Insert(sg, suite.DBCtx, suite.DB, boil.Infer())
	suite.NoError(err)
	return sg
}

func (suite *SecurityGroupIntegrationTestSuite) AssignToRunningSpace(spaceID, securityGroupID int) {
	sgSpace := &models.SecurityGroupsSpace{
		SpaceID:         spaceID,
		SecurityGroupID: securityGroupID,
	}
	err := models.SecurityGroupsSpaces().Insert(sgSpace, suite.DBCtx, suite.DB, boil.Infer())
	suite.NoError(err)
}

func (suite *SecurityGroupIntegrationTestSuite) AssignToStagingSpace(spaceID, securityGroupID int) {
	sgSpace := &models.StagingSecurityGroupsSpace{
		StagingSpaceID:         spaceID,
		StagingSecurityGroupID: securityGroupID,
	}
	err := models.StagingSecurityGroupsSpaces().Insert(sgSpace, suite.DBCtx, suite.DB, boil.Infer())
	suite.NoError(err)
}
