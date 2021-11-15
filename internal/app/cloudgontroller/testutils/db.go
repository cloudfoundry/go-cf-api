// +build integration

package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"go.uber.org/zap/zaptest"
)

type DBIntegrationTestSuite struct {
	suite.Suite

	DB     *sql.DB
	DBCtx  context.Context
	Random *rand.Rand
}

var configFile string

func init() {
	configFile = os.Args[len(os.Args)-1]
}

func (s *DBIntegrationTestSuite) Setup() {
	ccConfig, err := config.Get(configFile)
	s.Require().NoError(err)
	db, info, err := db.NewConnection(ccConfig.DB, true)
	s.Require().NoError(err)
	logger := zaptest.NewLogger(s.T())
	logger.Info(fmt.Sprintf("Using DB type: %s at %s:%s", info.Type, info.Host, info.Port))

	s.DB = db
	s.Random = rand.New(rand.NewSource(time.Now().UTC().Unix()))
	s.DBCtx = boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
}

func (s *DBIntegrationTestSuite) ClearTables(tables []string) {
	for _, table := range tables {
		_, err := models.NewQuery(qm.SQL(fmt.Sprintf("DELETE FROM %s", models.Quote(table)))).Exec(s.DB)
		s.Require().NoError(err)
	}
}

func (s *DBIntegrationTestSuite) CreateUser() *models.User {
	user := &models.User{
		GUID: fmt.Sprintf("user-guid-%d", s.Random.Int()),
	}
	err := models.Users().Insert(user, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return user
}

func (s *DBIntegrationTestSuite) CreateQuota() *models.QuotaDefinition {
	id := s.Random.Int()
	quota := &models.QuotaDefinition{
		GUID: fmt.Sprintf("quota-guid-%d", id),
		Name: fmt.Sprintf("quota-%d", id),
	}
	err := models.QuotaDefinitions().Insert(quota, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return quota
}

func (s *DBIntegrationTestSuite) CreateOrg(quotaID int) *models.Organization {
	id := s.Random.Int()
	org := &models.Organization{
		GUID:              fmt.Sprintf("org-guid-%d", id),
		Name:              fmt.Sprintf("org-%d", id),
		QuotaDefinitionID: quotaID,
	}
	err := models.Organizations().Insert(org, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return org
}

func (s *DBIntegrationTestSuite) CreateSpace(orgID int) *models.Space {
	id := s.Random.Int()
	space := &models.Space{
		GUID:           fmt.Sprintf("space-guid-%d", id),
		Name:           fmt.Sprintf("space-%d", id),
		OrganizationID: orgID,
	}
	err := models.Spaces().Insert(space, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return space
}

func (s *DBIntegrationTestSuite) AssignSpaceDeveloper(userID, spaceID int) {
	developer := &models.SpacesDeveloper{
		UserID:  userID,
		SpaceID: spaceID,
	}
	err := models.SpacesDevelopers().Insert(developer, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignSpaceAuditor(userID, spaceID int) {
	auditor := &models.SpacesAuditor{
		UserID:  userID,
		SpaceID: spaceID,
	}
	err := models.SpacesAuditors().Insert(auditor, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignSpaceSupporter(userID, spaceID int) {
	developer := &models.SpacesApplicationSupporter{
		UserID:  userID,
		SpaceID: spaceID,
	}
	err := models.SpacesApplicationSupporters().Insert(developer, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignSpaceManager(userID, spaceID int) {
	manager := &models.SpacesManager{
		UserID:  userID,
		SpaceID: spaceID,
	}
	err := models.SpacesManagers().Insert(manager, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignOrgManager(userID, orgID int) {
	manager := &models.OrganizationsManager{
		UserID:         userID,
		OrganizationID: orgID,
	}
	err := models.OrganizationsManagers().Insert(manager, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignOrgBillingManager(userID, orgID int) {
	billingManager := &models.OrganizationsBillingManager{
		UserID:         userID,
		OrganizationID: orgID,
	}
	err := models.OrganizationsBillingManagers().Insert(billingManager, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}

func (s *DBIntegrationTestSuite) AssignOrgAuditor(userID, orgID int) {
	auditor := &models.OrganizationsAuditor{
		UserID:         userID,
		OrganizationID: orgID,
	}
	err := models.OrganizationsAuditors().Insert(auditor, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
}
