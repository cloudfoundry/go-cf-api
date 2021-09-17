// +build unit

package buildpacks //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
)

//nolint:gochecknoglobals // convenient helper in tests
var bpCols = models.BuildpackColumns

type BuildpackControllerTestSuite struct {
	suite.Suite
	ctx        echo.Context
	req        *http.Request
	rec        *httptest.ResponseRecorder
	controller Controller
	ctrl       *gomock.Controller
	querier    *mock_models.MockBuildpackFinisher
	inserter   *mock_models.MockBuildpackInserter
	queryMods  []qm.QueryMod
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
	suite.inserter = mock_models.NewMockBuildpackInserter(suite.ctrl)
	buildpackQuerier = func(qm ...qm.QueryMod) models.BuildpackFinisher {
		suite.queryMods = qm
		return suite.querier
	}
	buildpackInserter = suite.inserter

	suite.controller = Controller{DB: nil}
}
