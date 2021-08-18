// +build unit

package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/ccerrors"
)

type GetMultipleBuildpacksTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	SQLMock             sqlmock.Sqlmock
	buildpackController BuildpackController
}

func (suite *GetMultipleBuildpacksTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, mock, _ := sqlmock.New()
	buildpackController := BuildpackController{DB: db}

	suite.Ctx = c
	suite.SQLMock = mock
	suite.Rec = *rec
	suite.buildpackController = buildpackController
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusOk() {
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow("first-guid").AddRow("second-guid"))
	if assert.NoError(suite.T(), suite.buildpackController.GetBuildpacks(suite.Ctx)) {
		assert.Contains(suite.T(), suite.Rec.Body.String(), "first-guid")
		assert.Contains(suite.T(), suite.Rec.Body.String(), "second-guid")
		assert.Equal(suite.T(), http.StatusOK, suite.Ctx.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusNotFound() {
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}))

	assert.NoError(suite.T(), suite.buildpackController.GetBuildpacks(suite.Ctx))
	assert.Equal(suite.T(), http.StatusNotFound, suite.Ctx.Response().Status)
}

func (suite *GetMultipleBuildpacksTestSuite) TestInternalServerError() {
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" LIMIT 50;`)).
		WillReturnError(errors.New("bommel"))

	assert.Error(suite.T(), ccerrors.UnknownError(), suite.buildpackController.GetBuildpacks(suite.Ctx))
}

func TestGetMultipleBuildpacksTestSuite(t *testing.T) {
	suite.Run(t, new(GetMultipleBuildpacksTestSuite))
}

type GetBuildpackTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	SQLMock             sqlmock.Sqlmock
	buildpackController BuildpackController
}

func (suite *GetBuildpackTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, mock, _ := sqlmock.New()
	buildpackController := BuildpackController{db}

	suite.Ctx = c
	suite.Rec = *rec
	suite.SQLMock = mock
	suite.buildpackController = buildpackController
}

func (suite *GetBuildpackTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (guid=$1) LIMIT 1;`)).
		WithArgs(expectedGUID).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow(expectedGUID))

	if assert.NoError(suite.T(), suite.buildpackController.GetBuildpack(suite.Ctx)) {
		assert.Contains(suite.T(), suite.Rec.Body.String(), expectedGUID)
		assert.Equal(suite.T(), http.StatusOK, suite.Rec.Code)
	}
}

func (suite *GetBuildpackTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (guid=$1) LIMIT 1;`)).
		WithArgs(expectedGUID).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}))

	var err *ccerrors.CloudControllerError
	assert.ErrorAs(suite.T(), suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	assert.Equal(suite.T(), http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetBuildpackTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (guid=$1) LIMIT 1;`)).
		WithArgs(expectedGUID).
		WillReturnError(errors.New("bommel"))

	var err *ccerrors.CloudControllerError
	assert.ErrorAs(suite.T(), suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	assert.Equal(suite.T(), http.StatusInternalServerError, err.HTTPStatus)
}

func TestGetBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(GetBuildpackTestSuite))
}