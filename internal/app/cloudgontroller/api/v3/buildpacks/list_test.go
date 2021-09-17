// +build unit

package buildpacks //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type GetMultipleBuildpacksTestSuite struct {
	BuildpackControllerTestSuite
	observedLogs *observer.ObservedLogs
}

func TestGetMultipleBuildpacksTestSuite(t *testing.T) {
	suite.Run(t, new(GetMultipleBuildpacksTestSuite))
}

func (suite *GetMultipleBuildpacksTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/buildpacks")

	core, recorded := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	suite.observedLogs = recorded
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusOk() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"},
	}, nil)
	if assert.NoError(suite.T(), suite.controller.List(suite.ctx)) {
		assert.Contains(suite.T(), suite.rec.Body.String(), "first-guid")
		assert.Contains(suite.T(), suite.rec.Body.String(), "second-guid")
		assert.Equal(suite.T(), http.StatusOK, suite.ctx.Response().Status)
	}
	suite.Contains(suite.queryMods, qm.Limit(50))
	suite.Contains(suite.queryMods, qm.Offset(0))
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusNotFound() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(nil, nil)

	suite.NoError(suite.controller.List(suite.ctx))
	suite.Equal(http.StatusNotFound, suite.ctx.Response().Status)
}

func (suite *GetMultipleBuildpacksTestSuite) TestInternalServerError() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("something went wrong"))

	suite.Error(v3.UnknownError(nil), suite.controller.List(suite.ctx))
}

func (suite *GetMultipleBuildpacksTestSuite) TestMetadataIsEagerLoaded() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"},
	}, nil)
	assert.NoError(suite.T(), suite.controller.List(suite.ctx))
	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	suite.Contains(suite.queryMods, qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)))
	suite.Contains(suite.queryMods, qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestPaginationParameters() {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"}, {GUID: "third-guid"},
	}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Limit(2))
	suite.Contains(suite.queryMods, qm.Offset(4))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `"total_pages":2`)
		assert.Contains(suite.T(), rec.Body.String(), `"total_results":3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", bpCols.Position)))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack", "go_buildpack"))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3"))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.UpdatedAt), timeNow))
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack", "go_buildpack", "php_buildpack"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `php_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleName() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?names=java_buildpack",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack"))

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

	err := suite.controller.List(context)
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3", "testStack", "testStack2"))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack2`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterSingleStack() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?stacks=cflinuxfs3",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Stack: null.StringFrom("cflinuxfs3")}}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3"))

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

	err := suite.controller.List(context)
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByPositionDescending() { //nolint:dupl
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=-position",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", bpCols.Position)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByCreated() { //nolint:dupl
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=created_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.CreatedAt)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUpdated() { //nolint:dupl
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=updated_at",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: "java_buildpack"}}, nil)

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.UpdatedAt)))

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

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.List(context), &err)
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", bpCols.UpdatedAt), timeNow))

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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", bpCols.UpdatedAt), timeNow))

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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", bpCols.UpdatedAt), timeNow))
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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.UpdatedAt), timeNow))

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

	err := suite.controller.List(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.CreatedAt), startTimeFormatted))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.CreatedAt), endTimeFormatted))

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

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.List(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithInvalidComparisonOperator() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?created_ats[foo]=%s", timeNow), nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.List(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}
