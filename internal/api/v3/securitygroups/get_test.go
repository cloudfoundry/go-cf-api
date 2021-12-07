//go:build unit
// +build unit

//nolint:forcetypeassert // Casting from mock calls doesn't need to be checked
package securitygroups //nolint:testpackage // we have to assign package level vars due to models using static functions

import (
	"database/sql"
	"errors"
	"fmt"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/auth"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/permissions"
	mock_permissions "github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/permissions/mocks"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

type GetSecurityGroupTestSuite struct {
	SecurityGroupControllerTestSuite
}

func TestGetSecurityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GetSecurityGroupTestSuite))
}

func (suite *GetSecurityGroupTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/security_groups")
	suite.ctx.Set(auth.Username, "user-guid")

	allowedSpaceIDs := mock_permissions.NewMockAllowedSpaceIDs(suite.ctrl)
	allowedSpaceIDs.EXPECT().With().Return([]qm.QueryMod{qm.Comment("Mocked WITH statement")})
	allowedSpaceIDs.EXPECT().Contains(models.SecurityGroupsSpaceTableColumns.SpaceID).Return(qm.Comment("Mocked WHERE statement"))
	allowedSpaceIDs.EXPECT().Contains(models.StagingSecurityGroupsSpaceTableColumns.StagingSpaceID).Return(qm.Comment("Mocked WHERE statement"))
	permissionsQuerier := mock_permissions.NewMockQuerier(suite.ctrl)
	permissionsQuerier.EXPECT().AllowedSpaceIDsForUser(
		"user-guid",
		permissions.SpaceDeveloper,
		permissions.SpaceSupporter,
		permissions.SpaceManager,
		permissions.SpaceAuditor,
		permissions.OrgManager,
	).Return(allowedSpaceIDs, nil)
	suite.controller.Permissions = permissionsQuerier
}

func (suite *GetSecurityGroupTestSuite) TestSecurityGroupFound() {
	expectedGUID := "123"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{GUID: expectedGUID}, nil)
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	suite.NoError(suite.controller.Get(suite.ctx))
	suite.Equal(http.StatusOK, suite.rec.Code)
	suite.Contains(suite.rec.Body.String(), expectedGUID)
	suite.presenter.AssertNumberOfCalls(suite.T(), "ResponseObject", 1)
}

func (suite *GetSecurityGroupTestSuite) TestDBNoRowsError() {
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues("non-existing-guid")

	dbErr := sql.ErrNoRows
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, dbErr)

	var err *v3.CfApiError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.ResourceNotFound("security_group", dbErr), err)
	suite.presenter.AssertNotCalled(suite.T(), "ResponseObject")
}

func (suite *GetSecurityGroupTestSuite) TestOtherDBError() {
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues("other-db-error")

	dbErr := errors.New("something went wrong")
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, dbErr)

	var err *v3.CfApiError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.UnknownError(dbErr), err)
	suite.presenter.AssertNotCalled(suite.T(), "ResponseObject")
}

func (suite *GetSecurityGroupTestSuite) TestPresenterError() {
	expectedGUID := "presenter-error"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	presenterErr := errors.New("something went wrong")
	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{}, presenterErr)
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	var err *v3.CfApiError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.UnknownError(fmt.Errorf("could not construct response: %w", presenterErr)), err)
}

func (suite *GetSecurityGroupTestSuite) TestQueryMods() {
	expectedGUID := "123"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{GUID: expectedGUID}, nil)
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	suite.NoError(suite.controller.Get(suite.ctx))
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 1)
	queryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)

	suite.Contains(queryMods, models.SecurityGroupWhere.GUID.EQ(expectedGUID))
	suite.Contains(queryMods, qm.Load(qm.Rels(
		models.SecurityGroupRels.SecurityGroupsSpaces,
		models.SecurityGroupsSpaceRels.Space,
	)))
	suite.Contains(queryMods, qm.Load(qm.Rels(
		models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces,
		models.StagingSecurityGroupsSpaceRels.StagingSpace,
	)))
}

type GetSecurityGroupAdminTestSuite struct {
	SecurityGroupControllerTestSuite
}

func TestGetSecurityGroupAdminTestSuite(t *testing.T) {
	suite.Run(t, new(GetSecurityGroupAdminTestSuite))
}

func (suite *GetSecurityGroupAdminTestSuite) TestPermissionsNotCalledForAdminRoles() {
	for _, scope := range []auth.Scope{auth.Admin, auth.AdminReadOnly, auth.GlobalAuditor} {
		suite.Run(fmt.Sprintf("test permissions not called for %s", scope), func() {
			suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/security_groups")
			suite.ctx.Set(auth.Username, "user-guid")
			suite.ctx.SetParamNames(GUIDParam)
			suite.ctx.SetParamValues("123")

			suite.ctx.Set(auth.Scopes, []string{string(scope)})

			// No EXPECT so mock shouldn't be called
			permissions := mock_permissions.NewMockQuerier(suite.ctrl)
			suite.controller.Permissions = permissions

			suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{GUID: "123"}, nil)
			suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: "123", Rules: null.StringFrom("[]")}, nil)

			suite.NoError(suite.controller.Get(suite.ctx))
		})
	}
}
