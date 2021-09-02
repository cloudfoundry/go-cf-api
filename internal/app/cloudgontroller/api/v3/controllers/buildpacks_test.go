// +build unit

package controllers_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow("first-guid").AddRow("second-guid"))
	if suite.NoError(suite.buildpackController.GetBuildpacks(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), "first-guid")
		suite.Contains(suite.Rec.Body.String(), "second-guid")
		suite.Equal(http.StatusOK, suite.Ctx.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusNotFound() {
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}))

	suite.NoError(suite.buildpackController.GetBuildpacks(suite.Ctx))
	suite.Equal(http.StatusNotFound, suite.Ctx.Response().Status)
}

func (suite *GetMultipleBuildpacksTestSuite) TestInternalServerError() {
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" LIMIT 50;`)).
		WillReturnError(errors.New("bommel"))

	suite.Error(UnknownError(nil), suite.buildpackController.GetBuildpacks(suite.Ctx))
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 2 OFFSET 4;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow("first-guid").AddRow("second-guid").AddRow("third-guid"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `"total_pages":2`)
		suite.Contains(rec.Body.String(), `"total_results":3`)
		suite.Equal(http.StatusOK, context.Response().Status)
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

	if suite.NoError(suite.buildpackController.GetBuildpack(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.Rec.Code)
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
	suite.ErrorAs(suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
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
	suite.ErrorAs(suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	suite.Equal(http.StatusInternalServerError, err.HTTPStatus)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ("name" IN ($1,$2)) AND ("stack" IN ($3)) AND
(created_at > $4) AND (updated_at <= $5) ORDER BY position DESC LIMIT 50;`)).
		WithArgs("java_buildpack", "go_buildpack", "cflinuxfs3", timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"name", "stack", "created_at", "updated_at"}).
			AddRow("java_buildpack", "cflinuxfs3", timeAsTime, timeAsTime).
			AddRow("go_buildpack", "cflinuxfs3", timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ("name" IN ($1,$2,$3)) ORDER BY position ASC LIMIT 50;`)).
		WithArgs("java_buildpack", "go_buildpack", "php_buildpack").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).
			AddRow("java_buildpack").
			AddRow("go_buildpack").
			AddRow("php_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ("name" IN ($1)) ORDER BY position ASC LIMIT 50;`)).
		WithArgs("java_buildpack").
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack").AddRow("go_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ("stack" IN ($1,$2,$3)) ORDER BY position ASC LIMIT 50;`)).
		WithArgs("cflixnuxfs3", "testStack", "testStack2").
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).
			AddRow("cflixnuxfs3").
			AddRow("testStack").
			AddRow("testStack2"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Contains(rec.Body.String(), `testStack`)
		suite.Contains(rec.Body.String(), `testStack2`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE ("stack" IN ($1)) ORDER BY position ASC LIMIT 50;`)).
		WithArgs("cflixnuxfs3").
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).AddRow("cflixnuxfs3"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"stack"}).AddRow("cflixnuxfs3").AddRow("testStack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Contains(rec.Body.String(), `testStack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByPosition() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=position",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByPositionDescending() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=-position",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position DESC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByCreated() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=created_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY created_at ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUpdated() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=updated_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY updated_at ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("java_buildpack"))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUnknownFilter() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=foo",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.GetBuildpacks(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at = $1) AND (updated_at = $2) ORDER BY position ASC LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at < $1) AND (updated_at > $2) ORDER BY position ASC LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at >= $1) AND (updated_at < $2) ORDER BY position ASC LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
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
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at <= $1) AND (updated_at >= $2) ORDER BY position ASC LIMIT 50;`)).
		WithArgs(timeNow, timeNow).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(timeAsTime, timeAsTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeBetweenTimestamps() {
	startTime := time.Now().UTC()
	startTimeFormatted := startTime.Format(time.RFC3339)
	endTime := startTime.Add(time.Hour)
	endTimeFormatted := endTime.Format(time.RFC3339)
	req := httptest.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"http://localhost:8080/v3/buildpacks?created_ats[gte]=%s&created_ats[lte]=%s",
			startTimeFormatted, endTimeFormatted),
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" WHERE (created_at >= $1) AND (created_at <= $2) ORDER BY position ASC LIMIT 50;`)).
		WithArgs(startTimeFormatted, endTimeFormatted).
		WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).
			AddRow(startTime, endTime))

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), startTimeFormatted)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithInvalidCreatedAtsParam() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?created_ats(lte)=%s", timeNow), nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.GetBuildpacks(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithInvalidComparisonOperator() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?created_ats[foo]=%s", timeNow), nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.GetBuildpacks(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}

func TestGetBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(GetBuildpackTestSuite))
}

type PostBuildpackTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	SQLMock             sqlmock.Sqlmock
	buildpackController BuildpackController
}

func (suite *PostBuildpackTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, mock, _ := sqlmock.New()
	buildpackController := BuildpackController{db}

	suite.Ctx = c
	suite.Rec = *rec
	suite.SQLMock = mock
	suite.buildpackController = buildpackController
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithName() {
	buildpackName := "test_buildpack" //nolint:goconst // mistakenly gets taken as duplicate
	reader := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\"}", buildpackName))
	req := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), 1, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	err := suite.buildpackController.PostBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithRequiredParams() {
	buildpackName := "test_buildpack"
	position := 1
	reader := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\"}", buildpackName))
	req := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), position, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	err := suite.buildpackController.PostBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}
//nolint: lll
func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithOptionalParams() {
	buildpackName := "test_buildpack"
	stack := "stacks"
	enabled := "true"
	locked := "false"
	reader := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"enabled\" : %s, \"locked\" : %s, \"stack\" : \"%s\"}", buildpackName, enabled, locked, stack))
	req := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","enabled","locked","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), 1, true, false, nil, nil, stack).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))

	err := suite.buildpackController.PostBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithInvalidJson() {
	buildpackName := "test_buildpack"
	stack := "stacks"
	enabled := "true"
	locked := "false"
	reader := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"enabled\" = %s, \"locked\" : %s, \"stack\" : \"%s\"}", buildpackName, enabled, locked, stack))
	req := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position",
"enabled","locked","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), 1, true, false, nil, nil, stack).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))

	err := suite.buildpackController.PostBuildpacks(context)
	assert.Error(suite.T(), UnprocessableEntity("Could not parse JSON provided in the body", err), suite.buildpackController.PostBuildpacks(suite.Ctx))
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithExistedPosition() {
	buildpackName1 := "test_buildpack1"
	position1 := 1
	buildpackName2 := "test_buildpack2"
	position2 := 1
	reader1 := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"position\" : %v}", buildpackName1, position1))
	req1 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader1)
	rec1 := httptest.NewRecorder()
	context1 := echo.New().NewContext(req1, rec1)
	reader2 := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"position\" : %v}", buildpackName2, position2))
	req2 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader2)
	rec2 := httptest.NewRecorder()
	context2 := echo.New().NewContext(req2, rec2)
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName1, sqlmock.AnyArg(), position1, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName2, sqlmock.AnyArg(), position2, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	_ = suite.buildpackController.PostBuildpacks(context1)
	err2 := suite.buildpackController.PostBuildpacks(context2)

	assert.Error(suite.T(), UnprocessableEntity("Position already exists", err2), suite.buildpackController.PostBuildpacks(suite.Ctx))
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithoutExistedPosition() {
	buildpackName1 := "test_buildpack1"
	position1 := 1
	buildpackName2 := "test_buildpack2"
	position2 := 0
	reader1 := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"position\" : %v}", buildpackName1, position1))
	req1 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader1)
	rec1 := httptest.NewRecorder()
	context1 := echo.New().NewContext(req1, rec1)
	reader2 := strings.NewReader(fmt.Sprintf("{\"name\" : \"%s\", \"position\" : %v}", buildpackName2, position2))
	req2 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader2)
	rec2 := httptest.NewRecorder()
	context2 := echo.New().NewContext(req2, rec2)
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName2, sqlmock.AnyArg(), 2, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name", "position"}).AddRow(buildpackName1, position1))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName2, sqlmock.AnyArg(), 2, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	_ = suite.buildpackController.PostBuildpacks(context1)
	err2 := suite.buildpackController.PostBuildpacks(context2)

	if assert.NoError(suite.T(), err2, fmt.Errorf("%w", errors.Unwrap(err2)).Error()) {
		assert.Contains(suite.T(), rec2.Body.String(), buildpackName2)
		assert.Equal(suite.T(), http.StatusOK, context2.Response().Status)
	}
}

func TestPostBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(PostBuildpackTestSuite))
}