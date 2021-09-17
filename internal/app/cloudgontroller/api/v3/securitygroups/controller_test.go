package securitygroups //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
)

type SecurityGroupControllerTestSuite struct {
	suite.Suite
	ctx         echo.Context
	req         *http.Request
	rec         *httptest.ResponseRecorder
	controller  Controller
	querier     *mock_models.MockSecurityGroupFinisher
	querierFunc *QuerierFunc
	presenter   *MockPresenter
}

func (suite *SecurityGroupControllerTestSuite) SetupTestSuite(method, endpoint string) {
	e := echo.New()
	req := httptest.NewRequest(method, endpoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	suite.ctx = c
	suite.req = req
	suite.rec = rec
	ctrl := gomock.NewController(suite.T())
	suite.querier = mock_models.NewMockSecurityGroupFinisher(ctrl)
	suite.querierFunc = &QuerierFunc{querier: suite.querier}
	suite.querierFunc.On("Get", mock.Anything).Return(suite.querier)
	securityGroupQuerier = suite.querierFunc.Get
	suite.presenter = &MockPresenter{}
	suite.controller = Controller{DB: nil, Presenter: suite.presenter}
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) ResponseObject(securityGroup *models.SecurityGroup, resourcePath string) (*Response, error) {
	args := m.Called(securityGroup, resourcePath)
	return args.Get(0).(*Response), args.Error(1)
}

type QuerierFunc struct {
	mock.Mock
	querier *mock_models.MockSecurityGroupFinisher
}

func (m *QuerierFunc) Get(mods ...qm.QueryMod) models.SecurityGroupFinisher {
	args := m.Called(mods)
	return args.Get(0).(models.SecurityGroupFinisher)
}