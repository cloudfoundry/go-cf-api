package securitygroups //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

type GetSecurityGroupTestSuite struct {
	SecurityGroupControllerTestSuite
}

func TestGetSecurityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GetSecurityGroupTestSuite))
}

func (suite *GetSecurityGroupTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/security_groups")
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
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	dbErr := sql.ErrNoRows
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, dbErr)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.ResourceNotFound("security_group", dbErr), err)
	suite.presenter.AssertNotCalled(suite.T(), "ResponseObject")
}

func (suite *GetSecurityGroupTestSuite) TestOtherDBError() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	dbErr := errors.New("something went wrong")
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, dbErr)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.UnknownError(dbErr), err)
	suite.presenter.AssertNotCalled(suite.T(), "ResponseObject")
}

func (suite *GetSecurityGroupTestSuite) TestPresenterError() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	presenterErr := errors.New("something went wrong")
	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{}, presenterErr)
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(v3.UnknownError(fmt.Errorf("could not construct response: %w", presenterErr)), err)
}

func (suite *GetSecurityGroupTestSuite) TestQueryMods() {
	expectedGUID := "123"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	suite.NoError(suite.controller.Get(suite.ctx))
	suite.Contains(suite.queryMods, models.SecurityGroupWhere.GUID.EQ(expectedGUID))
	suite.Contains(suite.queryMods, qm.Load(qm.Rels(
		models.SecurityGroupRels.SecurityGroupsSpaces,
		models.SecurityGroupsSpaceRels.Space,
	)))
	suite.Contains(suite.queryMods, qm.Load(qm.Rels(
		models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces,
		models.StagingSecurityGroupsSpaceRels.StagingSpace,
	)))
}
