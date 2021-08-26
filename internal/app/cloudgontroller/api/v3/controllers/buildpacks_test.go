// +build unit

package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type GetMultipleBuildpacksTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	SQLMock             sqlmock.Sqlmock
	buildpackController BuildpackController
	logger              *zap.Logger
	ObservedLogs        *observer.ObservedLogs
}

func (suite *GetMultipleBuildpacksTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, mock, _ := sqlmock.New()
	buildpackController := BuildpackController{DB: db}

	core, recorded := observer.New(zapcore.InfoLevel)
	suite.logger = zap.New(core)
	zap.ReplaceGlobals(suite.logger)
	suite.ObservedLogs = recorded

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

	assert.Error(suite.T(), UnknownError(nil), suite.buildpackController.GetBuildpacks(suite.Ctx))
}

func (suite *GetMultipleBuildpacksTestSuite) TestPaginationParameters() {
	// Mock Call
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	// Mock SQL Queries
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 2 OFFSET 4;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow("first-guid").AddRow("second-guid").AddRow("third-guid"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `"total_pages":2`)
		assert.Contains(suite.T(), rec.Body.String(), `"total_results":3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
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

	var err *CloudControllerError
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

	var err *CloudControllerError
	assert.ErrorAs(suite.T(), suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	assert.Equal(suite.T(), http.StatusInternalServerError, err.HTTPStatus)
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterEverything() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?names=java_buildpack,go_buildpack&"+
			"stacks=cflinuxfs3&order_by=-position&created_ats[gt]=%s&updated_ats[lte]=%s", timeNow, timeNow),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ((name=$1) OR (name=$2)) AND (stack=$3) AND 
(created_at>$4) AND (updated_at<=$5) ORDER BY position LIMIT 50;`)).
		WithArgs("java_buildpack", "go_buildpack", "cflinuxfs3", timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"name", "stack", "created_at", "updated_at"}).
			AddRow("java_buildpack", "cflinuxfs3", timeAsTime, timeAsTime).
			AddRow("go_buildpack", "cflinuxfs3", timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterMultipleNames() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?names=java_buildpack,go_buildpack,php_buildpack",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ((name=$1) OR (name=$2) OR (name=$3)) ORDER BY position LIMIT 50;`)).
		WithArgs("java_buildpack", "go_buildpack", "php_buildpack").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("java_buildpack").
			AddRow("go_buildpack").
			AddRow("php_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleName() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?names=java_buildpack",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (name=$1) ORDER BY position LIMIT 50;`)).
		WithArgs("java_buildpack").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterEmptyNames() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?names=",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack").AddRow("go_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterMultipleStacks() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=cflixnuxfs3,testStack,testStack2",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ((stack=$1) OR (stack=$2) OR (stack=$3)) ORDER BY position LIMIT 50;`)).
		WithArgs("cflixnuxfs3", "testStack", "testStack2").
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).
			AddRow("cflixnuxfs3").
			AddRow("testStack").
			AddRow("testStack2"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflixnuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack2`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleStack() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=cflixnuxfs3",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (stack=$1) ORDER BY position LIMIT 50;`)).
		WithArgs("cflixnuxfs3").
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).AddRow("cflixnuxfs3"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflixnuxfs3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterEmptyStacks() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).AddRow("cflixnuxfs3").AddRow("testStack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflixnuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByPosition() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=-position",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByCreated() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=created_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY created_at LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUpdated() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=updated_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY updated_at LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTime() { //nolint:dupl // mistakenly gets taken as duplicate
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"http://localhost:8080/v3/buildpacks?created_ats=%s&updated_ats=%s",
			timeNow, timeNow),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at=$1) AND (updated_at=$2) ORDER BY position LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithSuffix() { //nolint:dupl // mistakenly gets taken as duplicate
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"http://localhost:8080/v3/buildpacks?created_ats[lt]=%s&updated_ats[gt]=%s",
			timeNow, timeNow),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at<$1) AND (updated_at>$2) ORDER BY position LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithOtherSuffix() { //nolint:dupl // mistakenly gets taken as duplicate
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"http://localhost:8080/v3/buildpacks?created_ats[gte]=%s&updated_ats[lt]=%s",
			timeNow, timeNow),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at>=$1) AND (updated_at<$2) ORDER BY position LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithSuffixEquals() { //nolint:dupl // mistakenly gets taken as duplicate
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"http://localhost:8080/v3/buildpacks?created_ats[lte]=%s&updated_ats[gte]=%s",
			timeNow, timeNow),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at<=$1) AND (updated_at>=$2) ORDER BY position LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func TestGetBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(GetBuildpackTestSuite))
}