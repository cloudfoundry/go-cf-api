package security_groups

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
)

type GetSecurityGroupTestSuite struct {
	suite.Suite
	Ctx                     echo.Context
	Rec                     httptest.ResponseRecorder
	securityGroupController Controller
	queryMods               []qm.QueryMod
	querier                 *mock_models.MockSecurityGroupFinisher
}

func TestGetSecurityGroupTestSuite(t *testing.T) {
	suite.Run(t, new(GetSecurityGroupTestSuite))
}

func (suite *GetSecurityGroupTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/security_groups", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	securityGroupController := Controller{DB: nil, Presenter: &MockPresenter{}}

	suite.Ctx = c
	suite.Rec = *rec
	suite.securityGroupController = securityGroupController
	ctrl := gomock.NewController(suite.T())
	suite.querier = mock_models.NewMockSecurityGroupFinisher(ctrl)
	securityGroupQuerier = func(qm ...qm.QueryMod) models.SecurityGroupFinisher {
		suite.queryMods = qm
		return suite.querier
	}
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error) {
	response := &Response{GUID: "123"}

	return response, nil
}
func (suite *GetSecurityGroupTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.SecurityGroup{GUID: expectedGUID, Rules: null.StringFrom("[]")}, nil)

	if suite.NoError(suite.securityGroupController.Get(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.Rec.Code)
	}
}

func (suite *GetSecurityGroupTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, nil)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.securityGroupController.Get(suite.Ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetSecurityGroupTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, errors.New("something went wrong"))

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.securityGroupController.Get(suite.Ctx), &err)
	suite.Equal(http.StatusInternalServerError, err.HTTPStatus)
}
