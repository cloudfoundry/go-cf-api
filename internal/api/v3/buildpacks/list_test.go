//go:build unit
// +build unit

//nolint:forcetypeassert // Casting from mock calls doesn't need to be checked
package buildpacks //nolint:testpackage // we have to assign package level vars due to models using static functions

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	mock_metadata "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/metadata/mocks"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/pagination"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type GetMultipleBuildpacksTestSuite struct {
	BuildpackControllerTestSuite
	observedLogs         *observer.ObservedLogs
	labelSelectorParser  *mock_metadata.MockLabelSelectorParser
	labelSelectorFilters *mock_metadata.MockLabelSelectorFilters
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
	suite.labelSelectorFilters = mock_metadata.NewMockLabelSelectorFilters(suite.ctrl)
	suite.labelSelectorParser = mock_metadata.NewMockLabelSelectorParser(suite.ctrl)
	// Default mocks for label selectors
	suite.labelSelectorFilters.EXPECT().Filters(models.TableNames.Buildpacks, models.TableNames.BuildpackLabels).AnyTimes().Return([]qm.QueryMod{})
	suite.labelSelectorParser.EXPECT().Parse("").AnyTimes().Return(suite.labelSelectorFilters, nil)
	suite.controller.LabelSelectorParser = suite.labelSelectorParser
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusOKWhenBuildpacksExist() {
	buildpacks := models.BuildpackSlice{
		{GUID: "first-guid"},
		{GUID: "second-guid"},
	}
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(buildpacks, nil)
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	if assert.NoError(suite.T(), suite.controller.List(suite.ctx)) {
		suite.presenter.AssertCalled(suite.T(), "ListResponseObject", buildpacks, int64(50), pagination.Params{Page: 1, PerPage: 50}, mock.Anything)
		assert.Equal(suite.T(), http.StatusOK, suite.ctx.Response().Status)
	}

	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Empty(countQueryMods)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)
	suite.Contains(allQueryMods, qm.Limit(50))
	suite.Contains(allQueryMods, qm.Offset(0))
	suite.Contains(allQueryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestStatusOKWhenNoBuildpacks() {
	var buildpacks models.BuildpackSlice
	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(0), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(buildpacks, sql.ErrNoRows)
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

	if assert.NoError(suite.T(), suite.controller.List(suite.ctx)) {
		suite.presenter.AssertCalled(suite.T(), "ListResponseObject", buildpacks, int64(0), pagination.Params{Page: 1, PerPage: 50}, mock.Anything)
		assert.Equal(suite.T(), http.StatusOK, suite.ctx.Response().Status)
	}

	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Empty(countQueryMods)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)
	suite.Contains(allQueryMods, qm.Limit(50))
	suite.Contains(allQueryMods, qm.Offset(0))
	suite.Contains(allQueryMods, qm.OrderBy(fmt.Sprintf("%s ASC", bpCols.Position)))
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
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	assert.NoError(suite.T(), suite.controller.List(suite.ctx))
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Empty(countQueryMods)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)

	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	suite.Contains(allQueryMods, qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)))
	suite.Contains(allQueryMods, qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)))
}

func (suite *GetMultipleBuildpacksTestSuite) TestPaginationParameters() {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?per_page=2&page=3", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(3), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{}, nil)
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

	err := suite.controller.List(context)
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Empty(countQueryMods)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)

	suite.Contains(allQueryMods, qm.Limit(2))
	suite.Contains(allQueryMods, qm.Offset(4))
	if assert.NoError(suite.T(), err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.presenter.AssertCalled(suite.T(), "ListResponseObject", models.BuildpackSlice{}, int64(3), pagination.Params{Page: 3, PerPage: 2}, mock.Anything)
		assert.Equal(suite.T(), http.StatusOK, context.Response().Status)
	}
}

func (suite *GetMultipleBuildpacksTestSuite) TestLabelSelectorsParameters() {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/v3/buildpacks?label_selector=key1,!key2", nil)
	rec := httptest.NewRecorder()
	context := echo.New().NewContext(req, rec)

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{GUID: "first-guid"}, {GUID: "second-guid"}, {GUID: "third-guid"},
	}, nil)

	var labelSelectorArg string
	suite.labelSelectorParser.EXPECT().Parse(gomock.Any()).DoAndReturn(
		func(arg string) (interface{}, error) {
			labelSelectorArg = arg
			return suite.labelSelectorFilters, nil
		})
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

	err := suite.controller.List(context)
	suite.NoError(err)
	suite.Equal("key1,!key2", labelSelectorArg)
}

func (suite *GetMultipleBuildpacksTestSuite) TestLabelSelectorsFilter() {
	labelSelectorFilters := mock_metadata.NewMockLabelSelectorFilters(suite.ctrl)
	labelSelectorParser := mock_metadata.NewMockLabelSelectorParser(suite.ctrl)
	suite.controller.LabelSelectorParser = labelSelectorParser

	filterQueryMod := qm.Comment("QueryMod should contain label selector")

	labelSelectorParser.EXPECT().Parse(gomock.Any()).AnyTimes().Return(labelSelectorFilters, nil)
	labelSelectorFilters.EXPECT().Filters(models.TableNames.Buildpacks, models.TableNames.BuildpackLabels).Return([]qm.QueryMod{filterQueryMod})

	suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).Return(int64(50), nil)
	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{}, nil)
	suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)
	suite.NoError(suite.controller.List(suite.ctx))
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)
	countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
	suite.Contains(countQueryMods, filterQueryMod)
	allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)

	suite.Contains(allQueryMods, filterQueryMod)
}

func (suite *GetMultipleBuildpacksTestSuite) TestFilters() {
	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	baseQueryMods := []qm.QueryMod{
		qm.Limit(50),
		qm.Offset(0),
		qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)),
		qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)),
	}
	now := time.Now().UTC().Format(time.RFC3339)
	cases := map[string]struct {
		query                string
		expectedCountFilters []qm.QueryMod
		expectedAllFilters   []qm.QueryMod
		expectedErr          *v3.CfAPIError
	}{
		"no name": {
			query:                "names=",
			expectedCountFilters: []qm.QueryMod{qmhelper.WhereIsNull("buildpacks.name")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qmhelper.WhereIsNull("buildpacks.name"),
			},
		},
		"single name": {
			query:                "names=java_buildpack",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("buildpacks.name IN ?", "java_buildpack")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.WhereIn("buildpacks.name IN ?", "java_buildpack"),
			},
		},
		"multiple names": {
			query:                "names=java_buildpack,go_buildpack,php_buildpack",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("buildpacks.name IN ?", "java_buildpack", "go_buildpack", "php_buildpack")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.WhereIn("buildpacks.name IN ?", "java_buildpack", "go_buildpack", "php_buildpack"),
			},
		},
		"no stack": {
			query:                "stacks=",
			expectedCountFilters: []qm.QueryMod{qmhelper.WhereIsNull("buildpacks.stack")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qmhelper.WhereIsNull("buildpacks.stack"),
			},
		},
		"single stack": {
			query:                "stacks=cflinuxfs3",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3"),
			},
		},
		"multiple stacks": {
			query:                "stacks=cflinuxfs3,cflinuxfs2",
			expectedCountFilters: []qm.QueryMod{qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3", "cflinuxfs2")},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3", "cflinuxfs2"),
			},
		},
		"order by position": {
			query:                "order_by=position",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("position ASC")},
		},
		"order by -position": {
			query:                "order_by=-position",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("position DESC")},
		},
		"order by created_at": {
			query:                "order_by=created_at",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("created_at ASC")},
		},
		"order by -created_at": {
			query:                "order_by=-created_at",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("created_at DESC")},
		},
		"order by updated_at": {
			query:                "order_by=updated_at",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("updated_at ASC")},
		},
		"order by -updated_at": {
			query:                "order_by=-updated_at",
			expectedCountFilters: []qm.QueryMod{},
			expectedAllFilters:   []qm.QueryMod{qm.OrderBy("updated_at DESC")},
		},
		"order by unknown filter": {
			query:       "order_by=foo",
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by time": {
			query: fmt.Sprintf("created_ats=%s&updated_ats=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("buildpacks.created_at = ?", now),
				qm.Where("buildpacks.updated_at = ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.Where("buildpacks.created_at = ?", now),
				qm.Where("buildpacks.updated_at = ?", now),
			},
		},
		"filter by time with lt/gt suffix": {
			query: fmt.Sprintf("created_ats[lt]=%s&updated_ats[gt]=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("buildpacks.created_at < ?", now),
				qm.Where("buildpacks.updated_at > ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.Where("buildpacks.created_at < ?", now),
				qm.Where("buildpacks.updated_at > ?", now),
			},
		},
		"filter by time with lte/gte suffix": {
			query: fmt.Sprintf("created_ats[gte]=%s&updated_ats[lte]=%s", now, now),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("buildpacks.created_at >= ?", now),
				qm.Where("buildpacks.updated_at <= ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.Where("buildpacks.created_at >= ?", now),
				qm.Where("buildpacks.updated_at <= ?", now),
			},
		},
		"filter by time between timestamps": {
			query: fmt.Sprintf(
				"created_ats[gte]=%s&created_ats[lte]=%s",
				time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339),
				time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339),
			),
			expectedCountFilters: []qm.QueryMod{
				qm.Where("buildpacks.created_at >= ?", time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339)),
				qm.Where("buildpacks.created_at <= ?", time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339)),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.OrderBy("position ASC"),
				qm.Where("buildpacks.created_at >= ?", time.Now().UTC().Add(-1*time.Hour).Format(time.RFC3339)),
				qm.Where("buildpacks.created_at <= ?", time.Now().UTC().Add(1*time.Hour).Format(time.RFC3339)),
			},
		},
		"filter by time with invalid param": {
			query:       fmt.Sprintf("created_ats(lte)=%s", now),
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by time with invalid comparison operator": {
			query:       fmt.Sprintf("created_ats[foo]=%s", now),
			expectedErr: v3.BadQueryParameter(errors.New("validation error")),
		},
		"filter by everything": {
			query: fmt.Sprintf(
				"names=java_buildpack,go_buildpack&stacks=cflinuxfs3&order_by=-position&created_ats[gt]=%s&updated_ats[lte]=%s",
				now, now,
			),
			expectedCountFilters: []qm.QueryMod{
				qm.WhereIn("buildpacks.name IN ?", "java_buildpack", "go_buildpack"),
				qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3"),
				qm.Where("buildpacks.created_at > ?", now),
				qm.Where("buildpacks.updated_at <= ?", now),
			},
			expectedAllFilters: []qm.QueryMod{
				qm.WhereIn("buildpacks.name IN ?", "java_buildpack", "go_buildpack"),
				qm.WhereIn("buildpacks.stack IN ?", "cflinuxfs3"),
				qm.OrderBy("position DESC"),
				qm.Where("buildpacks.created_at > ?", now),
				qm.Where("buildpacks.updated_at <= ?", now),
			},
		},
	}
	for testCaseName, testCase := range cases {
		suite.Run(testCaseName, func() {
			suite.SetupTest() // needed to ensure mocks are fresh for every test case
			suite.req.URL.RawQuery = testCase.query
			suite.querier.EXPECT().Count(gomock.Any(), gomock.Any()).AnyTimes().Return(int64(50), nil)
			suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).AnyTimes().Return(models.BuildpackSlice{{Name: "test_buildpack"}}, nil)
			suite.presenter.On("ListResponseObject", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&ListResponse{}, nil)

			err := suite.controller.List(suite.ctx)
			if testCase.expectedErr != nil {
				var ccErr *v3.CfAPIError
				suite.ErrorAs(err, &ccErr)
				suite.Equal(testCase.expectedErr.HTTPStatus, ccErr.HTTPStatus)
				return
			}
			suite.NoError(err)
			suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 2)

			countQueryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)
			suite.ElementsMatch(testCase.expectedCountFilters, countQueryMods)

			allQueryMods := suite.querierFunc.Calls[1].Arguments.Get(0).([]qm.QueryMod)
			expectedAllQueryMods := append(baseQueryMods, testCase.expectedAllFilters...) //nolint:gocritic // Deliberately appending to a different slice
			suite.ElementsMatch(expectedAllQueryMods, allQueryMods)
		})
	}
}
