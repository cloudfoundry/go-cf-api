//go:build integration
// +build integration

package securitygroups_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/cloudfoundry/go-cf-api/internal/api/v3/securitygroups"
	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/auth"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/permissions"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/cloudfoundry/go-cf-api/internal/testutils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type guids struct {
	unassigned      string
	globallyStaging string
	globallyRunning string
	spaceStaging    string
	spaceRunning    string
}

type SecurityGroupIntegrationTestSuite struct {
	testutils.DBIntegrationTestSuite

	controller *Controller
	e          *echo.Echo

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
	tablesToClear                []string
	guids                        guids
}

func TestSecurityGroupIntegrationTest(t *testing.T) {
	suite.Run(t, new(SecurityGroupIntegrationTestSuite))
}

func (suite *SecurityGroupIntegrationTestSuite) SetupSuite() {
	// Order matters to prevent foreign key errors
	suite.tablesToClear = []string{
		models.TableNames.SpacesDevelopers,
		models.TableNames.SpacesAuditors,
		models.TableNames.SpacesApplicationSupporters,
		models.TableNames.SpacesManagers,
		models.TableNames.SecurityGroupsSpaces,
		models.TableNames.StagingSecurityGroupsSpaces,
		models.TableNames.Spaces,
		models.TableNames.SecurityGroups,
		models.TableNames.OrganizationsManagers,
		models.TableNames.OrganizationsBillingManagers,
		models.TableNames.OrganizationsAuditors,
		models.TableNames.Organizations,
		models.TableNames.QuotaDefinitions,
		models.TableNames.Users,
	}

	suite.guids = guids{
		unassigned:      "unassigned",
		globallyRunning: "globally-running",
		globallyStaging: "globally-staging",
		spaceStaging:    "space-staging",
		spaceRunning:    "space-running",
	}

	suite.Setup()
	suite.ClearTables(suite.tablesToClear)
	suite.controller = &Controller{
		DB:          suite.DB,
		Presenter:   NewPresenter(),
		Permissions: permissions.NewQuerier(),
	}
	suite.e = echo.New()
}

func (suite *SecurityGroupIntegrationTestSuite) TearDownSuite() {
	suite.ClearTables(suite.tablesToClear)
}

func (suite *SecurityGroupIntegrationTestSuite) SetupTest() {
	suite.ClearTables(suite.tablesToClear)

	suite.user = suite.CreateUser()
	suite.quota = suite.CreateQuota()
	suite.org = suite.CreateOrg(suite.quota.ID)
	suite.runningSpace = suite.CreateSpace(suite.org.ID)
	suite.stagingSpace = suite.CreateSpace(suite.org.ID)
	suite.otherSpace = suite.CreateSpace(suite.org.ID)
	suite.unassignedSecurityGroup = suite.CreateSecurityGroup(suite.guids.unassigned, false, false)
	suite.globallyStagingSecurityGroup = suite.CreateSecurityGroup(suite.guids.globallyStaging, false, true)
	suite.globallyRunningSecurityGroup = suite.CreateSecurityGroup(suite.guids.globallyRunning, true, false)
	suite.spaceStagingSecurityGroup = suite.CreateSecurityGroup(suite.guids.spaceStaging, false, false)
	suite.AssignToStagingSpace(suite.stagingSpace.ID, suite.spaceStagingSecurityGroup.ID)
	suite.spaceRunningSecurityGroup = suite.CreateSecurityGroup(suite.guids.spaceRunning, false, false)
	suite.AssignToRunningSpace(suite.runningSpace.ID, suite.spaceRunningSecurityGroup.ID)
}

func (suite *SecurityGroupIntegrationTestSuite) newRequest() (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/security_groups", nil)
	defer req.Body.Close()
	rec := httptest.NewRecorder()
	ctx := suite.e.NewContext(req, rec)

	ctx.Set("username", suite.user.GUID)
	return rec, ctx
}

func (suite *SecurityGroupIntegrationTestSuite) TestListAndGetPermissions() {
	// We need to use func() so that when the spaces/security groups are updated in SetupTest we get the new objects
	cases := map[string]struct {
		setupContext           func(echo.Context)
		assignRoles            func()
		expectedSecurityGroups []Entity
	}{
		"Admin can see all security groups": {
			setupContext: func(c echo.Context) { c.Set(auth.Scopes, []string{string(auth.Admin)}) },
			assignRoles:  func() {},
			expectedSecurityGroups: []Entity{
				{suite.guids.unassigned},
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
				{suite.guids.spaceRunning},
			},
		},
		"AdminReadOnly can see all security groups": {
			setupContext: func(c echo.Context) { c.Set(auth.Scopes, []string{string(auth.AdminReadOnly)}) },
			assignRoles:  func() {},
			expectedSecurityGroups: []Entity{
				{suite.guids.unassigned},
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
				{suite.guids.spaceRunning},
			},
		},
		"GlobalAuditor can see all security groups": {
			setupContext: func(c echo.Context) { c.Set(auth.Scopes, []string{string(auth.GlobalAuditor)}) },
			assignRoles:  func() {},
			expectedSecurityGroups: []Entity{
				{suite.guids.unassigned},
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
				{suite.guids.spaceRunning},
			},
		},
		"OrgAuditor can see globally–enabled security groups": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignOrgAuditor(suite.user.ID, suite.org.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
			},
		},
		"OrgBillingManager can see globally–enabled security groups": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignOrgBillingManager(suite.user.ID, suite.org.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
			},
		},
		"OrgManager can see globally–enabled security groups or groups associated with a space they can seen": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignOrgManager(suite.user.ID, suite.org.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
				{suite.guids.spaceRunning},
			},
		},
		"SpaceAuditor can see globally–enabled security groups or groups associated with a space they can seen": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignSpaceAuditor(suite.user.ID, suite.stagingSpace.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
			},
		},
		"SpaceDeveloper can see globally–enabled security groups or groups associated with a space they can seen": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignSpaceDeveloper(suite.user.ID, suite.stagingSpace.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
			},
		},
		"SpaceManager can see globally–enabled security groups or groups associated with a space they can seen": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignSpaceManager(suite.user.ID, suite.stagingSpace.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
			},
		},
		"SpaceSupporter can see globally–enabled security groups or groups associated with a space they can seen": {
			setupContext: func(_ echo.Context) {},
			assignRoles:  func() { suite.AssignSpaceSupporter(suite.user.ID, suite.stagingSpace.ID) },
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
			},
		},
		"Multiple roles in different spaces can see security groups from either space": {
			setupContext: func(_ echo.Context) {},
			assignRoles: func() {
				suite.AssignSpaceDeveloper(suite.user.ID, suite.stagingSpace.ID)
				suite.AssignSpaceSupporter(suite.user.ID, suite.runningSpace.ID)
			},
			expectedSecurityGroups: []Entity{
				{suite.guids.globallyStaging},
				{suite.guids.globallyRunning},
				{suite.guids.spaceStaging},
				{suite.guids.spaceRunning},
			},
		},
	}

	for testCaseName, testCase := range cases {
		suite.Run(testCaseName, func() {
			suite.SetupTest()
			testCase.assignRoles()

			suite.Run("LIST", func() {
				rec, ctx := suite.newRequest()
				testCase.setupContext(ctx)
				err := suite.controller.List(ctx)
				suite.NoError(err)
				suite.Equal(http.StatusOK, rec.Result().StatusCode) //nolint:bodyclose

				resp := ResponseObject{}
				err = json.Unmarshal(rec.Body.Bytes(), &resp)
				suite.NoError(err)

				suite.ElementsMatch(testCase.expectedSecurityGroups, resp.Resources)
			})

			// Check that for every security group can see in list, they can get by GUID
			// and every security group they cannot see in list, cannot be retrieved by GUID
			for _, securityGroupGUID := range []string{
				suite.unassignedSecurityGroup.GUID,
				suite.globallyStagingSecurityGroup.GUID,
				suite.globallyRunningSecurityGroup.GUID,
				suite.spaceStagingSecurityGroup.GUID,
				suite.spaceRunningSecurityGroup.GUID,
			} {
				suite.Run(fmt.Sprintf("GET %s", securityGroupGUID), func() {
					rec, ctx := suite.newRequest()
					testCase.setupContext(ctx)
					ctx.SetParamNames(GUIDParam)
					ctx.SetParamValues(securityGroupGUID)

					err := suite.controller.Get(ctx)
					if contains(testCase.expectedSecurityGroups, securityGroupGUID) {
						suite.NoError(err)
						resp := Entity{}
						err = json.Unmarshal(rec.Body.Bytes(), &resp)
						suite.NoError(err)
						suite.Equal(securityGroupGUID, resp.GUID)
					} else {
						var ccErr *v3.CfAPIError
						suite.ErrorAs(err, &ccErr)
						suite.Equal(http.StatusNotFound, ccErr.HTTPStatus)
					}
				})
			}
		})
	}
}

type Entity struct {
	GUID string `json:"guid"`
}

type ResponseObject struct {
	Resources []Entity `json:"resources"`
}

func contains(entities []Entity, guid string) bool {
	for _, entity := range entities {
		if entity.GUID == guid {
			return true
		}
	}
	return false
}

func (suite *SecurityGroupIntegrationTestSuite) CreateSecurityGroup(name string, globallyEnabledRunning, globallyEnabledStaging bool) *models.SecurityGroup {
	securityGroups := &models.SecurityGroup{
		GUID:           name,
		Name:           fmt.Sprintf("%s-%s", name, uuid.New().String()),
		StagingDefault: null.BoolFrom(globallyEnabledStaging),
		RunningDefault: null.BoolFrom(globallyEnabledRunning),
		Rules:          null.StringFrom("[]"),
	}
	err := models.SecurityGroups().Insert(securityGroups, suite.DBCtx, suite.DB, boil.Infer())
	suite.NoError(err)
	return securityGroups
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
