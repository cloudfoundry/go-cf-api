package securitygroups //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
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

func (suite *GetSecurityGroupTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	if suite.NoError(suite.controller.Get(suite.ctx)) {
		suite.Contains(suite.rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.rec.Code)
	}
}

func (suite *GetSecurityGroupTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, nil)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetSecurityGroupTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, errors.New("something went wrong"))

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(http.StatusInternalServerError, err.HTTPStatus)
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
