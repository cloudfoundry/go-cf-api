// +build unit

//nolint:forcetypeassert // Casting from mock calls doesn't need to be checked
package securitygroups //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ListSecurityGroupTestSuite struct {
	SecurityGroupControllerTestSuite
}

func TestListSecurityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(ListSecurityGroupTestSuite))
}

func (suite *ListSecurityGroupTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/security_groups")
}

func (suite *ListSecurityGroupTestSuite) TestWithNoSecurityGroups() {
	securityGroups := models.SecurityGroupSlice{
		{},
	}
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, nil)

	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

	suite.NoError(suite.controller.List(suite.ctx))
	suite.Equal(http.StatusOK, suite.rec.Code)
	suite.NotEqual("null\n", suite.rec.Body.String())

	suite.presenter.AssertCalled(suite.T(), "ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func (suite *ListSecurityGroupTestSuite) TestWithASecurityGroup() {
	securityGroups := models.SecurityGroupSlice{
		{},
	}

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)

	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, nil)

	suite.NoError(suite.controller.List(suite.ctx))
	suite.Equal(http.StatusOK, suite.rec.Code)
	suite.NotEqual("null\n", suite.rec.Body.String())

	suite.presenter.AssertCalled(suite.T(), "ListResponseObject", securityGroups, mock.Anything, mock.Anything, mock.Anything)
}

func (suite *ListSecurityGroupTestSuite) TestWithPaginationParameters() {
	securityGroups := models.SecurityGroupSlice{
		{},
	}

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(3), nil)
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(securityGroups, nil)

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/security_groups?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.NoError(suite.controller.List(context))
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Empty(countQueryMods)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)

	suite.Contains(allQueryMods, qm.Limit(2))
	suite.Contains(allQueryMods, qm.Offset(4))
	suite.presenter.AssertCalled(suite.T(), "ListResponseObject", securityGroups, int64(3), pagination.Params{Page: 3, PerPage: 2}, mock.Anything)
}

func (suite *ListSecurityGroupTestSuite) TestInternalServerError() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("something went wrong"))
	suite.Error(v3.UnknownError(nil), suite.controller.List(suite.ctx))
}
