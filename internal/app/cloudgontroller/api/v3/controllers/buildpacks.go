package controllers

import (
	"context"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/presenter"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	psqlModels "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler/postgres"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/storage/db"
	"go.uber.org/zap"
)

// GetBuildpacks godoc
// @Summary List buildpacks
// @Description Retrieve all buildpacks the user has access to.
// @Produce json
// @Success 200
// @Router /buildpacks [get]
func GetBuildpacks(c echo.Context) error {
	db := db.GetConnection()

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true))
	buildpacks, err := psqlModels.Buildpacks(qm.Limit(50)).All(ctx, db)
	if err != nil {
		zap.L().Error("Couldn't select", zap.Error(err))
	}
	return c.JSON(http.StatusOK, presenter.BuildpacksResponseObject(buildpacks))
}

// GetBuildpack godoc
// @Summary Get a buildpack
// @Param guid path string true "Unique identifier for the buildpack"
// @Produce json
// @Success 200
// @Router /buildpacks/{guid} [get]
func GetBuildpack(c echo.Context) error {
	guid := c.Param("guid")

	db := db.GetConnection()

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true))
	buildpack, err := psqlModels.Buildpacks(qm.Where("guid=?", guid)).One(ctx, db)
	if err != nil {
		zap.L().Error("Couldn't select", zap.Error(err))
	}
	return c.JSON(http.StatusOK, presenter.BuildpackResponseObject(buildpack))
}
