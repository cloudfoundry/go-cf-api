package controllers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/api/v3/presenter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/logging"
	psqlModels "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/sqlboiler/postgres"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/gopilot/storage/db"
	"go.uber.org/zap"
)

// GetBuildpacks godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router api/v3/buildpacks [get]
func GetBuildpacks(c echo.Context) error {
	db := db.GetConnection()

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true))
	buildpacks, err := psqlModels.Buildpacks(qm.Limit(50)).All(ctx, db)
	if err != nil {
		zap.L().Error("Couldn't select", zap.Error(err))
	}
	return c.JSON(http.StatusOK, presenter.BuildpacksResponseObject(buildpacks))
}
