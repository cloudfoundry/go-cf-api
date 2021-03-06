//go:build unit
// +build unit

package buildpacks //nolint:testpackage // we have to assign package level vars due to models using static functions

import (
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/pagination"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	mock_models "github.com/cloudfoundry/go-cf-api/internal/storage/db/models/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

//nolint:gochecknoglobals // convenient helper in tests
var bpCols = models.BuildpackColumns

type BuildpackControllerTestSuite struct {
	suite.Suite
	ctx         echo.Context
	req         *http.Request
	rec         *httptest.ResponseRecorder
	controller  Controller
	ctrl        *gomock.Controller
	querier     *mock_models.MockBuildpackFinisher
	querierFunc *QuerierFunc
	inserter    *mock_models.MockBuildpackInserter
	updater     *mock_models.MockBuildpackUpdater
	presenter   *MockPresenter
	mockDB      sqlmock.Sqlmock
}

func (suite *BuildpackControllerTestSuite) SetupTestSuite(method, endpoint string) {
	e := echo.New()
	req := httptest.NewRequest(method, endpoint, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	suite.ctx = c
	suite.req = req
	suite.rec = rec
	suite.ctrl = gomock.NewController(suite.T())
	suite.querier = mock_models.NewMockBuildpackFinisher(suite.ctrl)
	suite.querierFunc = &QuerierFunc{}
	suite.querierFunc.On("Get", mock.Anything).Return(suite.querier)
	suite.inserter = mock_models.NewMockBuildpackInserter(suite.ctrl)
	suite.updater = mock_models.NewMockBuildpackUpdater(suite.ctrl)
	buildpackQuerier = suite.querierFunc.Get
	buildpackInserter = suite.inserter
	buildpackUpdater = suite.updater
	suite.presenter = &MockPresenter{}
	db, mockDB, _ := sqlmock.New()
	suite.mockDB = mockDB
	suite.controller = Controller{DB: db, Presenter: suite.presenter, LabelSelectorParser: nil}
}

type MockPresenter struct {
	mock.Mock
}

func (m *MockPresenter) ResponseObject(buildpack *models.Buildpack, resourcePath string) (*Response, error) {
	args := m.Called(buildpack, resourcePath)
	return args.Get(0).(*Response), args.Error(1)
}

func (m *MockPresenter) ListResponseObject(buildpacks models.BuildpackSlice,
	totalResults int64,
	paginationParams pagination.Params,
	resourcePath string) (*ListResponse, error) {
	args := m.Called(buildpacks, totalResults, paginationParams, resourcePath)
	return args.Get(0).(*ListResponse), args.Error(1)
}

type QuerierFunc struct {
	mock.Mock
}

func (m *QuerierFunc) Get(mods ...qm.QueryMod) models.BuildpackFinisher {
	args := m.Called(mods)
	return args.Get(0).(models.BuildpackFinisher)
}
