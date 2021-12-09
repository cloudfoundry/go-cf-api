package buildpacks

import (
	"context"
	"database/sql"
	"net/http"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

// Get godoc
// @Summary Show a buildpack
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Param guid path string true "Buildpack GUID"
// @Success 200 {object} Response
// @Success 404 {object} interface{}
// @Failure 400 {object} v3.CfAPIError
// @Failure 500 {object} v3.CfAPIError
// @Router /buildpacks/{guid} [get]
func (cont *Controller) Get(echoCtx echo.Context) error {
	guid := echoCtx.Param(GUIDParam)
	logger := logging.FromContext(echoCtx)

	boilCtx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	buildpack, err := buildpackQuerier(
		qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)),
		qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)),
		qm.Where("guid=?", guid),
	).One(boilCtx, cont.DB)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
		if errors.Cause(err) != sql.ErrNoRows {
			return v3.UnknownError(err)
		}
	}
	if buildpack == nil {
		return v3.ResourceNotFound("buildpack", err)
	}

	response, err := cont.Presenter.ResponseObject(buildpack, v3.GetResourcePath(echoCtx))
	if err != nil {
		return v3.UnknownError(err)
	}

	return echoCtx.JSON(http.StatusOK, response)
}
