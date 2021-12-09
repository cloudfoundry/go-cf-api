//go:build unit
// +build unit

package buildpacks //nolint:testpackage // we have to assign package level vars due to models using static functions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type PostBuildpackTestSuite struct {
	BuildpackControllerTestSuite
}

func (suite *PostBuildpackTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodPost, "http://localhost:8080/v3/buildpacks")
}

func TestPostBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(PostBuildpackTestSuite))
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithName() {
	buildpackName := "test_buildpack" //nolint:goconst // mistakenly gets taken as duplicate
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s"}`, buildpackName))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: buildpackName}}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})
	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{}, nil)

	err := suite.controller.Post(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(1, got.Position)
	suite.Equal(null.String{}, got.Filename)
	suite.Equal(null.String{}, got.Sha256Checksum)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.presenter.AssertCalled(suite.T(), "ResponseObject", got, mock.Anything)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithOptionalParams() {
	buildpackName := "test_buildpack"
	stack, enabled, locked := "stack", true, false
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "enabled" : %t, "locked" : %t, "stack" : "%s"}`, buildpackName, enabled, locked, stack))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{{Name: buildpackName}}, nil)
	var got *models.Buildpack
	suite.inserter.
		EXPECT().
		Insert(gomock.AssignableToTypeOf(&models.Buildpack{}), gomock.Any(), gomock.Any(), boil.Infer()).
		DoAndReturn(func(o *models.Buildpack, _, _, _ interface{}) error {
			got = o
			return nil
		})
	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{}, nil)

	err := suite.controller.Post(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(null.StringFrom("stack"), got.Stack)
	suite.True(got.Enabled.Bool)
	suite.False(got.Locked.Bool)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.presenter.AssertCalled(suite.T(), "ResponseObject", got, mock.Anything)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackswithInvalidJson() {
	reader := strings.NewReader(`{"name" : "}`)
	suite.req.Body = ioutil.NopCloser(reader)

	var err *v3.CfAPIError
	suite.ErrorAs(suite.controller.Post(suite.ctx), &err)
	suite.Equal(http.StatusUnprocessableEntity, err.HTTPStatus)
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithExistedPosition() {
	buildpackName, position := "test_buildpack", 1
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %d}`, buildpackName, position))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
		{Name: "existing_buildpack", Position: 1},
	}, nil)

	var err *v3.CfAPIError
	suite.ErrorAs(suite.controller.Post(suite.ctx), &err)
	suite.Equal(http.StatusUnprocessableEntity, err.HTTPStatus)
	suite.Contains(err.Detail, "Position already exists")
}

func (suite *PostBuildpackTestSuite) TestInsertBuildpackWithoutExistedPosition() {
	buildpackName, position := "test_buildpack", 2
	reader := strings.NewReader(fmt.Sprintf(`{"name" : "%s", "position" : %d}`, buildpackName, position))
	suite.req.Body = ioutil.NopCloser(reader)

	suite.querier.EXPECT().All(gomock.Any(), gomock.Any()).Return(models.BuildpackSlice{
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
	suite.presenter.On("ResponseObject", mock.Anything, mock.Anything).Return(&Response{}, nil)

	err := suite.controller.Post(suite.ctx)

	suite.Equal(buildpackName, got.Name)
	suite.Equal(position, got.Position)

	if suite.NoError(err, fmt.Errorf("%w", errors.Unwrap(err)).Error()) {
		suite.presenter.AssertCalled(suite.T(), "ResponseObject", got, mock.Anything)
		suite.Equal(http.StatusOK, suite.ctx.Response().Status)
	}
}
