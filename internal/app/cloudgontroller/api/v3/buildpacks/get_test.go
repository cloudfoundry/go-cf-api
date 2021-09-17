// +build unit

package buildpacks //nolint:testpackage // we have to assign package level vars due to sqlboiler using static functions

import (
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

type GetBuildpackTestSuite struct {
	BuildpackControllerTestSuite
}

func (suite *GetBuildpackTestSuite) SetupTest() {
	suite.SetupTestSuite(http.MethodGet, "http://localhost:8080/v3/buildpacks")
}

func TestGetBuildpackTestSuite(t *testing.T) {
	suite.Run(t, new(GetBuildpackTestSuite))
}

func (suite *GetBuildpackTestSuite) TestStatusOk() {
	expectedGUID := "123"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.Buildpack{GUID: expectedGUID}, nil)

	if suite.NoError(suite.controller.Get(suite.ctx)) {
		suite.Contains(suite.rec.Body.String(), expectedGUID)
		suite.Equal(http.StatusOK, suite.rec.Code)
	}
}

func (suite *GetBuildpackTestSuite) TestStatusNotFound() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, nil)

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(http.StatusNotFound, err.HTTPStatus)
}

func (suite *GetBuildpackTestSuite) TestInternalServerError() {
	expectedGUID := "non-existing-guid"
	suite.ctx.SetParamNames(GUIDParam)
	suite.ctx.SetParamValues(expectedGUID)

	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(nil, errors.New("something went wrong"))

	var err *v3.CloudControllerError
	suite.ErrorAs(suite.controller.Get(suite.ctx), &err)
	suite.Equal(http.StatusInternalServerError, err.HTTPStatus)
}

func (suite *GetBuildpackTestSuite) TestMetadataIsEagerLoaded() {
	suite.querier.EXPECT().One(gomock.Any(), gomock.Any()).Return(&models.Buildpack{GUID: "first-guid"}, nil)

	assert.NoError(suite.T(), suite.controller.Get(suite.ctx))
	suite.querierFunc.AssertNumberOfCalls(suite.T(), "Get", 1)
	queryMods := suite.querierFunc.Calls[0].Arguments.Get(0).([]qm.QueryMod)

	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	suite.Contains(queryMods, qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)))
	suite.Contains(queryMods, qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)))
}