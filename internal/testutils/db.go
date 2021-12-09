//go:build integration
// +build integration

package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/cloudfoundry/go-cf-api/internal/config"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap/zaptest"
)

type DBIntegrationTestSuite struct {
	suite.Suite

	DB    *sql.DB
	DBCtx context.Context
}

func (s *DBIntegrationTestSuite) Setup() {
	ccConfig, err := config.Get(os.Args[len(os.Args)-1])
	s.Require().NoError(err)
	db, info, err := db.NewConnection(ccConfig.DB, true)
	s.Require().NoError(err)
	logger := zaptest.NewLogger(s.T())
	logger.Info(fmt.Sprintf("Using DB type: %s at %s:%s", info.Type, info.Host, info.Port))

	s.DB = db
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
		GUID: fmt.Sprintf("user-guid-%s", uuid.New().String()),
	}
	err := models.Users().Insert(user, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return user
}

func (s *DBIntegrationTestSuite) CreateQuota() *models.QuotaDefinition {
	id := uuid.New().String()
	quota := &models.QuotaDefinition{
		GUID: fmt.Sprintf("quota-guid-%s", id),
		Name: fmt.Sprintf("quota-%s", id),
	}
	err := models.QuotaDefinitions().Insert(quota, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return quota
}

func (s *DBIntegrationTestSuite) CreateOrg(quotaID int) *models.Organization {
	id := uuid.New().String()
	org := &models.Organization{
		GUID:              fmt.Sprintf("org-guid-%s", id),
		Name:              fmt.Sprintf("org-%s", id),
		QuotaDefinitionID: quotaID,
	}
	err := models.Organizations().Insert(org, s.DBCtx, s.DB, boil.Infer())
	s.NoError(err)
	return org
}

func (s *DBIntegrationTestSuite) CreateSpace(orgID int) *models.Space {
	id := uuid.New().String()
	space := &models.Space{
		GUID:           fmt.Sprintf("space-guid-%s", id),
		Name:           fmt.Sprintf("space-%s", id),
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
