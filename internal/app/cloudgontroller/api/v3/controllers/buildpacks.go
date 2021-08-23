package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers/common"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/presenter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
)

const GUIDParam = "guid"

type BuildpackController struct {
	DB *sql.DB
}

// GetBuildpacks godoc
// @Summary Buildpacks List buildpacks
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Accept json
// @Produce json
// @Success 200 {object} []presenter.BuildpackResponse
// @Success 404 {object} interface{}
// @Failure 400 {object} []interface{}
// @Failure 500 {object} CloudControllerError
// @Router /buildpacks [get]
func (cont *BuildpackController) GetBuildpacks(c echo.Context) error {
	logger := logging.FromContext(c)
	pagination := common.DefaultPagination()
	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		return BadQueryParameter(err)
	}

	err := validator.New().Struct(pagination)
	if err != nil {
		return BadQueryParameter(err)
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	_, err = models.Buildpacks().Count(ctx, cont.DB)
	if err != nil {
		return UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}
	buildpacks, err := models.Buildpacks(
		qm.OrderBy("position"),
		qm.Limit(int(pagination.PerPage)),
		qm.Offset((pagination.Page-1)*int(pagination.PerPage)),
	).All(ctx, cont.DB)
	if err != nil {
		return UnknownError(fmt.Errorf("could not Select: %w", err))
	}
	if buildpacks == nil {
		return c.JSON(http.StatusNotFound, []presenter.BuildpackResponse{})
	}

	return c.JSON(http.StatusOK, presenter.BuildpacksResponseObject(buildpacks, pagination, GetResourcePath(c)))
}

// GetBuildpack godoc
// @Summary Show a buildpack
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Param guid path string true "Buildpack GUID"
// @Success 200 {object} presenter.BuildpackResponse
// @Success 404 {object} interface{}
// @Failure 400 {object} CloudControllerError
// @Failure 500 {object} CloudControllerError
// @Router /buildpacks/{guid} [get]
func (cont *BuildpackController) GetBuildpack(c echo.Context) error {
	guid := c.Param(GUIDParam)
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	buildpack, err := models.Buildpacks(qm.Where("guid=?", guid)).One(ctx, cont.DB)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
		if errors.Cause(err) != sql.ErrNoRows {
			return UnknownError(err)
		}
	}
	if buildpack == nil {
		return ResourceNotFound("buildpack", err)
	}

	return c.JSON(http.StatusOK, presenter.BuildpackResponseObject(buildpack, GetResourcePath(c)))
}
