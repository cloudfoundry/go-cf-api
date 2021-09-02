// +build unit

package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

var cols = models.BuildpackColumns

type GetMultipleBuildpacksTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	buildpackController BuildpackController
	queryMods           []qm.QueryMod
	querier             *mock_models.MockBuildpackFinisher
	logger              *zap.Logger
	ObservedLogs        *observer.ObservedLogs
}

func (suite *GetMultipleBuildpacksTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	buildpackController := BuildpackController{DB: nil}

	core, recorded := observer.New(zapcore.InfoLevel)
	suite.logger = zap.New(core)
	zap.ReplaceGlobals(suite.logger)
	suite.ObservedLogs = recorded

	suite.Ctx = c
	suite.Rec = *rec
	suite.buildpackController = buildpackController
	ctrl := gomock.NewController(suite.T())
	suite.querier = mock_models.NewMockBuildpackFinisher(ctrl)
	buildpackQuerier = func(qm ...qm.QueryMod) models.BuildpackFinisher {
		suite.queryMods = qm
		return suite.querier
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusOk() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"},
	}, nil)
	if assert.NoError(suite.T(), suite.buildpackController.GetBuildpacks(suite.Ctx)) {
		assert.Contains(suite.T(), suite.Rec.Body.String(), "first-guid")
		assert.Contains(suite.T(), suite.Rec.Body.String(), "second-guid")
		assert.Equal(suite.T(), http.StatusOK, suite.Ctx.Response().Status)
	}
	suite.Contains(suite.queryMods, qm.Limit(50))
	suite.Contains(suite.queryMods, qm.Offset(0))
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", cols.Position)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusNotFound() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(nil, nil)

	suite.NoError(suite.buildpackController.GetBuildpacks(suite.Ctx))
	suite.Equal(http.StatusNotFound, suite.Ctx.Response().Status)
}

func (suite *GetMultipleBuildpacksTestSuite) TestInternalServerError() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("something went wrong"))

	suite.Error(UnknownError(nil), suite.buildpackController.GetBuildpacks(suite.Ctx))
}

func (suite *GetMultipleBuildpacksTestSuite) TestPaginationParameters() {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"}, {GUID: "third-guid"},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Limit(2))
	suite.Contains(suite.queryMods, qm.Offset(4))
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
	queryMods           []qm.QueryMod
	querier             *mock_models.MockBuildpackFinisher
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
	ctrl := gomock.NewController(suite.T())
	suite.querier = mock_models.NewMockBuildpackFinisher(ctrl)
	buildpackQuerier = func(qm ...qm.QueryMod) models.BuildpackFinisher {
		suite.queryMods = qm
		return suite.querier
	}
}

func (suite *GetBuildpackTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.Buildpack{GUID: expectedGUID}, nil)

	if suite.NoError(suite.buildpackController.GetBuildpack(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.Rec.Code)
	}
}

func (suite *GetBuildpackTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, nil)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.GetBuildpack(suite.Ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetBuildpackTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, errors.New("something went wrong"))

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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", Stack: null.StringFrom("cflinuxfs3"), CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
		{Name: "go_buildpack", Stack: null.StringFrom("cflinuxfs3"), CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", cols.Position)))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Name), "java_buildpack", "go_buildpack"))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Stack), "cflinuxfs3"))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", cols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", cols.UpdatedAt), timeNow))
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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack"}, {Name: "go_buildpack"}, {Name: "php_buildpack"},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Name), "java_buildpack", "go_buildpack", "php_buildpack"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `php_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleName() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?names=java_buildpack",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Name), "java_buildpack"))

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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack"}, {Name: "go_buildpack"},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterMultipleStacks() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=cflinuxfs3,testStack,testStack2",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Stack: null.StringFrom("cflinuxfs3")},
		{Stack: null.StringFrom("testStack")},
		{Stack: null.StringFrom("testStack2")},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Stack), "cflinuxfs3", "testStack", "testStack2"))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack2`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleStack() { //nolint:dupl // mistakenly gets taken as duplicate
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=cflinuxfs3",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Stack: null.StringFrom("cflinuxfs3")}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", cols.Stack), "cflinuxfs3"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterEmptyStacks() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Stack: null.StringFrom("cflinuxfs3")}, {Stack: null.StringFrom("testStack")},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByPosition() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=position",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", cols.Position)))

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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", cols.Position)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByCreated() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=created_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", cols.CreatedAt)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUpdated() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=updated_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", cols.UpdatedAt)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", cols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", cols.UpdatedAt), timeNow))

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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", cols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", cols.UpdatedAt), timeNow))

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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", cols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", cols.UpdatedAt), timeNow))
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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", CreatedAt: timeAsTime, UpdatedAt: null.TimeFrom(timeAsTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", cols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", cols.UpdatedAt), timeNow))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
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

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "java_buildpack", CreatedAt: startTime, UpdatedAt: null.TimeFrom(startTime)},
	}, nil)

	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", cols.CreatedAt), startTimeFormatted))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", cols.CreatedAt), endTimeFormatted))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), startTimeFormatted)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
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
	ctx                 echo.Context
	req                 *http.Request
	rec                 *httptest.ResponseRecorder
	buildpackController BuildpackController
	finisher            *mock_models.MockBuildpackFinisher
	inserter            *mock_models.MockBuildpackInserter
}

func (suite *PostBuildpackTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	buildpackController := BuildpackController{DB: nil}

	suite.ctx = c
	suite.req = req
	suite.rec = rec
	suite.buildpackController = buildpackController
	ctrl := gomock.NewController(suite.T())
	suite.finisher = mock_models.NewMockBuildpackFinisher(ctrl)
	suite.inserter = mock_models.NewMockBuildpackInserter(ctrl)
	buildpackQuerier = func(qm ...qm.QueryMod) models.BuildpackFinisher {
		return suite.finisher
	}
	buildpackInserter = suite.inserter
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithName() {
	buildpackName := "test_buildpack" //nolint:goconst // mistakenly gets taken as duplicate
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s"}`, buildpackName))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: buildpackName}}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})

	err := suite.buildpackController.PostBuildpacks(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(1, got.Position)
	suite.Equal(null.String{}, got.Filename)
	suite.Equal(null.String{}, got.Sha256Checksum)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(suite.rec.Body.String(), buildpackName)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithOptionalParams() {
	buildpackName := "test_buildpack"
	stack, enabled, locked := "stack", true, false
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "enabled" : %t, "locked" : %t, "stack" : "%s"}`, buildpackName, enabled, locked, stack))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: buildpackName}}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})

	err := suite.buildpackController.PostBuildpacks(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(null.StringFrom("stack"), got.Stack)
	suite.True(got.Enabled.Bool)
	suite.False(got.Locked.Bool)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(suite.rec.Body.String(), buildpackName)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithInvalidJson() {
	reader := strings.NewReader(`{"name" : "}`)
	suite.req.Body = ioutil.NopCloser(reader)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.PostBuildpacks(suite.ctx), &err)
	suite.Equal(http.StatusUnprocessableEntity, err.HTTPStatus)
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithExistedPosition() {
	buildpackName, position := "test_buildpack", 1
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %d}`, buildpackName, position))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "existing_buildpack", Position: 1},
	}, nil)

	var err *CloudControllerError
	suite.ErrorAs(suite.buildpackController.PostBuildpacks(suite.ctx), &err)
	suite.Equal(http.StatusUnprocessableEntity, err.HTTPStatus)
	suite.Contains(err.Detail, "Position already exists")
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithoutExistedPosition() {
	buildpackName, position := "test_buildpack", 2
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %d}`, buildpackName, position))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "existing_buildpack", Position: 1},
	}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})

	err := suite.buildpackController.PostBuildpacks(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(position, got.Position)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(suite.rec.Body.String(), buildpackName)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}

func TestPostBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(PostBuildpackTestSuite))
}
