package buildpacks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

// PostBuildpack godoc
// @Summary Create a buildpack
// @Description Create a new buildpack
// @Tags Buildpacks
// @accept json
// @produce json
// @Success 201 {object} Response
// @Success 404 {object} interface{}
// @Failure 400 {object} v3.CloudControllerError
// @Failure 500 {object} v3.CloudControllerError
// @Router /buildpacks [post]
func (cont *Controller) Post(c echo.Context) error {
	logger := logging.FromContext(c)
	var buildpackToInsert *models.Buildpack

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))

	if err := json.NewDecoder(c.Request().Body).Decode(&buildpackToInsert); err != nil {
		logger.Error("Could not parse JSON provided in the body")
		return v3.UnprocessableEntity("Could not parse JSON provided in the body", err)
	}

	buildpacksInDB, errDB := buildpackQuerier().All(ctx, cont.DB)
	if errDB != nil {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", errDB))
	}

	if buildpackToInsert.Position == 0 {
		position := 1
		for _, buildpackToEvaluate := range buildpacksInDB {
			if buildpackToEvaluate.Position >= position {
				position = buildpackToEvaluate.Position + 1
			}
		}
		buildpackToInsert.Position = position
	} else {
		for _, buildpackToEvaluate := range buildpacksInDB {
			if buildpackToEvaluate.Position == buildpackToInsert.Position {
				logger.Error("Position already exists")
				return v3.UnprocessableEntity("Position already exists", fmt.Errorf("position already exists"))
			}
		}
	}

	// Add guid to Buildpack
	buildpackToInsert.GUID = uuid.New().String()
	err := buildpackInserter.Insert(buildpackToInsert, ctx, cont.DB, boil.Infer())
	if err != nil {
		logger.Error("There is no buildpack to insert")
		return v3.UnprocessableEntity("There is no buildpack to insert", err)
	}

	response, err := cont.Presenter.ResponseObject(buildpackToInsert, v3.GetResourcePath(c))
	if err != nil {
		return v3.UnknownError(err)
	}

	return c.JSON(http.StatusOK, response)
}
