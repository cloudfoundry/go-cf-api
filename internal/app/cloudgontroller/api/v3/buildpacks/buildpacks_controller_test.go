// +build unit

<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
package controllers //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions
=======
package buildpacks_test
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	mock_models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/mocks"
=======
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/buildpacks"
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

//nolint:gochecknoglobals // convenient helper in tests
var bpCols = models.BuildpackColumns

type GetMultipleBuildpacksTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	buildpackController BuildpackController
	queryMods           []qm.QueryMod
	querier             *mock_models.MockBuildpackFinisher
=======
	SQLMock             sqlmock.Sqlmock
<<<<<<< HEAD
	buildpackController buildpacks.BuildpackController
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController buildpacks.Controller
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	logger              *zap.Logger
	ObservedLogs        *observer.ObservedLogs
}

func (suite *GetMultipleBuildpacksTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	buildpackController := BuildpackController{DB: nil}
=======
	db, mock, _ := sqlmock.New()
<<<<<<< HEAD
	buildpackController := buildpacks.BuildpackController{DB: db}
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController := buildpacks.Controller{DB: db}
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)

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
<<<<<<< HEAD
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"},
	}, nil)
	if assert.NoError(suite.T(), suite.buildpackController.GetBuildpacks(suite.Ctx)) {
		assert.Contains(suite.T(), suite.Rec.Body.String(), "first-guid")
		assert.Contains(suite.T(), suite.Rec.Body.String(), "second-guid")
		assert.Equal(suite.T(), http.StatusOK, suite.Ctx.Response().Status)
=======
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(*) FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow("50"))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks" ORDER BY position ASC LIMIT 50;`)).
		WillReturnRows(sqlmock.NewRows([]string{"guid"}).AddRow("first-guid").AddRow("second-guid"))
	if suite.NoError(suite.buildpackController.List(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), "first-guid")
		suite.Contains(suite.Rec.Body.String(), "second-guid")
		suite.Equal(http.StatusOK, suite.Ctx.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
	suite.Contains(suite.queryMods, qm.Limit(50))
	suite.Contains(suite.queryMods, qm.Offset(0))
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusNotFound() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(nil, nil)

	suite.NoError(suite.buildpackController.List(suite.Ctx))
	suite.Equal(http.StatusNotFound, suite.Ctx.Response().Status)
}

func (suite *GetMultipleBuildpacksTestSuite) TestInternalServerError() {
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), errors.New("something went wrong"))

	suite.Error(v3.UnknownError(nil), suite.buildpackController.List(suite.Ctx))
}

func (suite *GetMultipleBuildpacksTestSuite) TestPaginationParameters() {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"}, {GUID: "third-guid"},
	}, nil)

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Limit(2))
	suite.Contains(suite.queryMods, qm.Offset(4))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `"total_pages":2`)
		assert.Contains(suite.T(), rec.Body.String(), `"total_results":3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `"total_pages":2`)
		suite.Contains(rec.Body.String(), `"total_results":3`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
}

func TestGetMultipleBuildpacksTestSuite(t *testing.T) {
	suite.Run(t, new(GetMultipleBuildpacksTestSuite))
}

type GetBuildpackTestSuite struct {
	suite.Suite
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	buildpackController BuildpackController
	queryMods           []qm.QueryMod
	querier             *mock_models.MockBuildpackFinisher
=======
	SQLMock             sqlmock.Sqlmock
<<<<<<< HEAD
	buildpackController buildpacks.BuildpackController
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController buildpacks.Controller
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
}

func (suite *GetBuildpackTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	buildpackController := BuildpackController{DB: nil}
=======
	db, mock, _ := sqlmock.New()
<<<<<<< HEAD
	buildpackController := buildpacks.BuildpackController{db}
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController := buildpacks.Controller{db}
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)

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

func (suite *GetBuildpackTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.Ctx.SetParamNames(buildpacks.GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.Buildpack{GUID: expectedGUID}, nil)

	if suite.NoError(suite.buildpackController.Get(suite.Ctx)) {
		suite.Contains(suite.Rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.Rec.Code)
	}
}

func (suite *GetBuildpackTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(buildpacks.GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, nil)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.buildpackController.Get(suite.Ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetBuildpackTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.Ctx.SetParamNames(buildpacks.GUIDParam)
	suite.Ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, errors.New("something went wrong"))

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.buildpackController.Get(suite.Ctx), &err)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", bpCols.Position)))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack", "go_buildpack"))
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3"))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.UpdatedAt), timeNow))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack", "go_buildpack", "php_buildpack"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `go_buildpack`)
		assert.Contains(suite.T(), rec.Body.String(), `php_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Contains(rec.Body.String(), `go_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Name), "java_buildpack"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

	err := suite.buildpackController.List(context)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3", "testStack", "testStack2"))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack2`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Contains(rec.Body.String(), `testStack`)
		suite.Contains(rec.Body.String(), `testStack2`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.WhereIn(fmt.Sprintf("%s IN ?", bpCols.Stack), "cflinuxfs3"))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `cflinuxfs3`)
		assert.Contains(suite.T(), rec.Body.String(), `testStack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `cflixnuxfs3`)
		suite.Contains(rec.Body.String(), `testStack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))

=======
	err := suite.buildpackController.List(context)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s DESC", bpCols.Position)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.CreatedAt)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.UpdatedAt)))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), `java_buildpack`)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), `java_buildpack`)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterOrderByUnknownFilter() {
	req := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080/v3/buildpacks?order_by=foo",
		nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.buildpackController.List(context), &err)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s = ?", bpCols.UpdatedAt), timeNow))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s > ?", bpCols.UpdatedAt), timeNow))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s < ?", bpCols.UpdatedAt), timeNow))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.CreatedAt), timeNow))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.UpdatedAt), timeNow))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), timeNow)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), timeNow)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	err := suite.buildpackController.GetBuildpacks(context)
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s >= ?", bpCols.CreatedAt), startTimeFormatted))
	suite.Contains(suite.queryMods, qm.Where(fmt.Sprintf("%s <= ?", bpCols.CreatedAt), endTimeFormatted))

	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), startTimeFormatted)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
=======
	err := suite.buildpackController.List(context)
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(rec.Body.String(), startTimeFormatted)
		suite.Equal(http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithInvalidCreatedAtsParam() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?created_ats(lte)=%s", timeNow), nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.buildpackController.List(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilterByTimeWithInvalidComparisonOperator() {
	timeAsTime := time.Now().UTC()
	timeNow := timeAsTime.Format(time.RFC3339)
	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8080/v3/buildpacks?created_ats[foo]=%s", timeNow), nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.buildpackController.List(context), &err)
	suite.Equal(http.StatusBadRequest, err.HTTPStatus)
}

func TestGetBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(GetBuildpackTestSuite))
}

type PostBuildpackTestSuite struct {
	suite.Suite
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	ctx                 echo.Context
	req                 *http.Request
	rec                 *httptest.ResponseRecorder
	buildpackController BuildpackController
	finisher            *mock_models.MockBuildpackFinisher
	inserter            *mock_models.MockBuildpackInserter
=======
	Ctx                 echo.Context
	Rec                 httptest.ResponseRecorder
	SQLMock             sqlmock.Sqlmock
<<<<<<< HEAD
	buildpackController buildpacks.BuildpackController
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController buildpacks.Controller
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
}

func (suite *PostBuildpackTestSuite) SetupTest() {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/v3/buildpacks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
<<<<<<< HEAD:internal/app/cloudgontroller/api/v3/controllers/buildpacks_test.go
	buildpackController := BuildpackController{DB: nil}
=======
	db, mock, _ := sqlmock.New()
<<<<<<< HEAD
	buildpackController := buildpacks.BuildpackController{db}
>>>>>>> 0051e52 (Create packages for different resources):internal/app/cloudgontroller/api/v3/buildpacks/buildpacks_controller_test.go
=======
	buildpackController := buildpacks.Controller{db}
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)

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

<<<<<<< HEAD
	suite.finisher.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: buildpackName}}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})
=======
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), 1, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	err := suite.buildpackController.Post(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)

	err := suite.buildpackController.PostBuildpacks(suite.ctx)

<<<<<<< HEAD
	suite.Equal(buildpackName, got.Name)
	suite.Equal(1, got.Position)
	suite.Equal(null.String{}, got.Filename)
	suite.Equal(null.String{}, got.Sha256Checksum)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(suite.rec.Body.String(), buildpackName)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
=======
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id","enabled","locked"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), position, nil, nil, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "enabled", "locked"}).
			AddRow(1, true, false))

	err := suite.buildpackController.Post(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
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

<<<<<<< HEAD
	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.Contains(suite.rec.Body.String(), buildpackName)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
=======
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "buildpacks";`)).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(buildpackName))
	suite.SQLMock.
		ExpectQuery(regexp.QuoteMeta(`INSERT INTO "buildpacks" ("guid","created_at","updated_at","name","key","position","enabled","locked","filename","sha256_checksum","stack")
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), buildpackName, sqlmock.AnyArg(), 1, true, false, nil, nil, stack).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))

	err := suite.buildpackController.Post(context)
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		assert.Contains(suite.T(), rec.Body.String(), buildpackName)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithInvalidJson() {
	reader := strings.NewReader(`{"name" : "}`)
	suite.req.Body = ioutil.NopCloser(reader)

<<<<<<< HEAD
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
=======
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

	err := suite.buildpackController.Post(context)
	assert.Error(suite.T(), v3.UnprocessableEntity("Could not parse JSON provided in the body", err), suite.buildpackController.Post(suite.Ctx))
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithExistedPosition() {
	buildpackName1, position1 := "test_buildpack1", 1
	buildpackName2, position2 := "test_buildpack2", 1
	reader1 := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %v}`, buildpackName1, position1))
	req1 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader1)
	rec1 := httptest.NewRecorder()
	context1 := echo.New().NewContext(req1, rec1)
	reader2 := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %v}`, buildpackName2, position2))
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

	_ = suite.buildpackController.Post(context1)
	err2 := suite.buildpackController.Post(context2)

	assert.Error(suite.T(), v3.UnprocessableEntity("Position already exists", err2), suite.buildpackController.Post(suite.Ctx))
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithoutExistedPosition() {
	buildpackName1, position1 := "test_buildpack1", 1
	buildpackName2, position2 := "test_buildpack2", 0
	reader1 := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %v}`, buildpackName1, position1))
	req1 := httptest.NewRequest(
		http.MethodPost, "http://localhost:8080/v3/buildpacks", reader1)
	rec1 := httptest.NewRecorder()
	context1 := echo.New().NewContext(req1, rec1)
	reader2 := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %v}`, buildpackName2, position2))
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

	_ = suite.buildpackController.Post(context1)
	err2 := suite.buildpackController.Post(context2)

	if assert.NoError(suite.T(), err2, fmt.Errorf("%w", errors.Unwrap(err2)).Error()) {
		assert.Contains(suite.T(), rec2.Body.String(), buildpackName2)
		assert.Equal(suite.T(), http.StatusOK, context2.Response().Status)
>>>>>>> 470b8b1 (Cut unnecessary 'buildpack' naming duplication)
	}
}

func TestPostBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(PostBuildpackTestSuite))
}
