// +build integration

package permissions_test

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
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
	models.TableNames.Spaces,
	models.TableNames.OrganizationsManagers,
	models.TableNames.OrganizationsBillingManagers,
	models.TableNames.OrganizationsAuditors,
	models.TableNames.Organizations,
	models.TableNames.QuotaDefinitions,
	models.TableNames.Users,
}

type AllowedSpaceIDsIntegrationTestSuite struct {
	testutils.DBIntegrationTestSuite

	querier Querier

	user  *models.User
	quota *models.QuotaDefinition
	org   *models.Organization
}

func TestAllowedSpaceIDsIntegrationTest(t *testing.T) {
	suite.Run(t, new(AllowedSpaceIDsIntegrationTestSuite))
}

func (suite *AllowedSpaceIDsIntegrationTestSuite) SetupSuite() {
	conf := config.DBConfig{
		Type:             os.Getenv("CC_DB_TYPE"),
		ConnectionString: os.Getenv("CC_DB_CONNECTION_STRING"),
	}
	suite.DB, _ = dbconfig.NewConnection(conf, true)
	suite.ClearTables(tablesToClear)
	logger := zaptest.NewLogger(suite.T())
	suite.Random = rand.New(rand.NewSource(time.Now().UTC().Unix()))
	suite.DBCtx = boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	suite.querier = NewQuerier()
}

func (suite *AllowedSpaceIDsIntegrationTestSuite) SetupTest() {
	suite.ClearTables(tablesToClear)
	suite.user = suite.CreateUser()
	suite.quota = suite.CreateQuota()
	suite.org = suite.CreateOrg(suite.quota.ID)
}

func (suite *AllowedSpaceIDsIntegrationTestSuite) TestSpaceRoles() {
	// Setup spaces and roles for tests
	developerSpace := suite.CreateSpace(suite.org.ID)
	suite.AssignSpaceDeveloper(suite.user.ID, developerSpace.ID)
	auditorSpace := suite.CreateSpace(suite.org.ID)
	suite.AssignSpaceAuditor(suite.user.ID, auditorSpace.ID)
	supporterSpace := suite.CreateSpace(suite.org.ID)
	suite.AssignSpaceSupporter(suite.user.ID, supporterSpace.ID)
	managerSpace := suite.CreateSpace(suite.org.ID)
	suite.AssignSpaceManager(suite.user.ID, managerSpace.ID)
	// Create extra space to check that spaces with no role aren't returned
	suite.CreateSpace(suite.org.ID)

	cases := map[string]struct {
		roles []Role
		// Just use IDs here since we can't compare *models.Space objects due to timestamps
		expectedSpaceIDs func() []int
	}{
		"space developer only": {
			roles:            []Role{SpaceDeveloper},
			expectedSpaceIDs: func() []int { return []int{developerSpace.ID} },
		},
		"space auditor only": {
			roles:            []Role{SpaceAuditor},
			expectedSpaceIDs: func() []int { return []int{auditorSpace.ID} },
		},
		"space supporter only": {
			roles:            []Role{SpaceSupporter},
			expectedSpaceIDs: func() []int { return []int{supporterSpace.ID} },
		},
		"space manager only": {
			roles:            []Role{SpaceManager},
			expectedSpaceIDs: func() []int { return []int{managerSpace.ID} },
		},
		"multiple roles": {
			roles:            []Role{SpaceDeveloper, SpaceSupporter},
			expectedSpaceIDs: func() []int { return []int{developerSpace.ID, supporterSpace.ID} },
		},
		"all space roles": {
			roles:            AllSpaceRoles,
			expectedSpaceIDs: func() []int { return []int{developerSpace.ID, auditorSpace.ID, supporterSpace.ID, managerSpace.ID} },
		},
	}
	for name, tc := range cases {
		suite.Run(name, func() {
			allowedSpaceIDs, err := suite.querier.AllowedSpaceIDsForUser(suite.user.GUID, tc.roles...)
			suite.NoError(err)
			suite.NotNil(allowedSpaceIDs)

			qms := []qm.QueryMod{qm.Select(models.SpaceTableColumns.ID)}
			qms = append(qms, allowedSpaceIDs.With()...)
			qms = append(qms, allowedSpaceIDs.Contains(models.SpaceTableColumns.ID))

			spaces, err := models.Spaces(qms...).All(suite.DBCtx, suite.DB)
			suite.NoError(err)
			for _, space := range spaces {
				suite.Contains(tc.expectedSpaceIDs(), space.ID)
			}
		})
	}
}

func (suite *AllowedSpaceIDsIntegrationTestSuite) TestOrgRoles() {
	// Setup orgs, spaces and roles for tests
	managerOrg := suite.CreateOrg(suite.quota.ID)
	managerSpace := suite.CreateSpace(managerOrg.ID)
	suite.AssignOrgManager(suite.user.ID, managerOrg.ID)
	billingManagerOrg := suite.CreateOrg(suite.quota.ID)
	billingManagerSpace := suite.CreateSpace(billingManagerOrg.ID)
	suite.AssignOrgBillingManager(suite.user.ID, billingManagerOrg.ID)
	auditorOrg := suite.CreateOrg(suite.quota.ID)
	auditorSpace := suite.CreateSpace(auditorOrg.ID)
	suite.AssignOrgAuditor(suite.user.ID, auditorOrg.ID)
	// Create extra space to check that spaces linked to org with no role aren't returned
	suite.CreateSpace(suite.org.ID)

	cases := map[string]struct {
		roles []Role
		// Just use IDs here since we can't compare *models.Space objects due to timestamps
		expectedSpaceIDs func() []int
	}{
		"org manager only": {
			roles:            []Role{OrgManager},
			expectedSpaceIDs: func() []int { return []int{managerSpace.ID} },
		},
		"org billing manager only": {
			roles:            []Role{OrgBillingManager},
			expectedSpaceIDs: func() []int { return []int{billingManagerSpace.ID} },
		},
		"org auditor only": {
			roles:            []Role{OrgAuditor},
			expectedSpaceIDs: func() []int { return []int{auditorSpace.ID} },
		},
		"multiple roles": {
			roles:            []Role{OrgManager, OrgAuditor},
			expectedSpaceIDs: func() []int { return []int{managerSpace.ID, auditorSpace.ID} },
		},
		"all org roles": {
			roles:            AllSpaceRoles,
			expectedSpaceIDs: func() []int { return []int{managerSpace.ID, auditorSpace.ID} },
		},
	}
	for name, tc := range cases {
		suite.Run(name, func() {
			allowedSpaceIDs, err := suite.querier.AllowedSpaceIDsForUser(suite.user.GUID, tc.roles...)
			suite.NoError(err)
			suite.NotNil(allowedSpaceIDs)

			qms := []qm.QueryMod{qm.Select(models.SpaceTableColumns.ID)}
			qms = append(qms, allowedSpaceIDs.With()...)
			qms = append(qms, allowedSpaceIDs.Contains(models.SpaceTableColumns.ID))

			spaces, err := models.Spaces(qms...).All(suite.DBCtx, suite.DB)
			suite.NoError(err)
			for _, space := range spaces {
				suite.Contains(tc.expectedSpaceIDs(), space.ID)
			}
		})
	}
}

func (suite *AllowedSpaceIDsIntegrationTestSuite) TestOrgAndSpaceRoles() {
	managerOrg := suite.CreateOrg(suite.quota.ID)
	orgManagerSpace := suite.CreateSpace(managerOrg.ID)
	suite.AssignOrgManager(suite.user.ID, managerOrg.ID)
	developerSpace := suite.CreateSpace(suite.org.ID)
	suite.AssignSpaceDeveloper(suite.user.ID, developerSpace.ID)
	// Create extra space to check that spaces linked with no role aren't returned
	suite.CreateSpace(suite.org.ID)

	// Include SpaceAuditor even though there are none to check it does not break query
	roles := []Role{SpaceDeveloper, SpaceAuditor, OrgManager}
	expectedSpaceIDs := []int{developerSpace.ID, orgManagerSpace.ID}

	allowedSpaceIDs, err := suite.querier.AllowedSpaceIDsForUser(suite.user.GUID, roles...)
	suite.NoError(err)
	suite.NotNil(allowedSpaceIDs)

	qms := []qm.QueryMod{qm.Select(models.SpaceTableColumns.ID)}
	qms = append(qms, allowedSpaceIDs.With()...)
	qms = append(qms, allowedSpaceIDs.Contains(models.SpaceTableColumns.ID))

	spaces, err := models.Spaces(qms...).All(suite.DBCtx, suite.DB)
	suite.NoError(err)
	for _, space := range spaces {
		suite.Contains(expectedSpaceIDs, space.ID)
	}
}
