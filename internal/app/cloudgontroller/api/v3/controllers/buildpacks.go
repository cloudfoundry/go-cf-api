package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/presenter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"go.uber.org/zap"
)

// GetBuildpacks godoc
// @Summary Buildpacks List buildpacks
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Accept json
// @Produce json
// @Success 200 {object} []presenter.BuildpackResponse
// @Success 404 {object} interface{}
// @Failure 400 {object} []interface{}
// @Failure 500 {object} HTTPError
// @Router /buildpacks [get]
func GetBuildpacks(c echo.Context) error {
	db := db.GetConnection()
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	buildpacks, err := models.Buildpacks(qm.Limit(50)).All(ctx, db) //nolint:gomnd // This won't be hardcoded when we finish this endpoint
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
	}
	if buildpacks == nil {
		return c.JSON(http.StatusNotFound, []presenter.BuildpackResponse{})
	}

	return c.JSON(http.StatusOK, presenter.BuildpacksResponseObject(buildpacks, GetResourcePath(c)))
}

// GetBuildpack godoc
// @Summary Show a buildpack
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Param guid path string true "Buildpack GUID"
// @Success 200 {object} presenter.BuildpackResponse
// @Success 404 {object} interface{}
// @Failure 400 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /buildpacks/{guid} [get]
func GetBuildpack(c echo.Context) error {
	guid := c.Param("guid")
	db := db.GetConnection()
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	buildpack, err := models.Buildpacks(qm.Where("guid=?", guid)).One(ctx, db)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
	}
	if buildpack == nil {
		return c.JSON(http.StatusNotFound, HTTPError{Code: http.StatusNotFound, Message: "Buildpack not Found"})
	}

	return c.JSON(http.StatusOK, presenter.BuildpackResponseObject(buildpack, GetResourcePath(c)))
}
