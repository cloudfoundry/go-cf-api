package buildpacks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type PostRequest struct {
	Name     string  `json:"name" validate:"required,min=1,max=250"`
	Stack    *string `json:"stack" validate:"omitempty,max=250"`
	Position int     `json:"position" validate:"gte=1"`
	Enabled  bool    `json:"enabled"`
	Locked   bool    `json:"locked"`
}

func DefaultPostRequest() PostRequest {
	return PostRequest{
		Position: 1,
		Enabled:  true,
		Locked:   false,
	}
}

// PostBuildpack godoc
// @Summary Create a buildpack
// @Description Create a new buildpack
// @Tags Buildpacks
// @accept json
// @produce json
// @Success 201 {object} Response
// @Success 404 {object} interface{}
// @Failure 400 {object} v3.CfAPIError
// @Failure 500 {object} v3.CfAPIError
// @Router /buildpacks [post]
//nolint:funlen // length is not a problem for now
func (cont *Controller) Post(echoCtx echo.Context) error {
	logger := logging.FromContext(echoCtx)
	requestBody := DefaultPostRequest()

	boilCtx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))

	if err := json.NewDecoder(echoCtx.Request().Body).Decode(&requestBody); err != nil {
		logger.Error("Could not parse JSON provided in the body")
		return v3.UnprocessableEntity("Could not parse JSON provided in the body", err)
	}

	err := validator.New().Struct(requestBody)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	dbTx, err := cont.DB.BeginTx(boilCtx, nil)
	if err != nil {
		return err
	}
	defer dbTx.Rollback() //nolint:errcheck  //No error handling needed in defer

	buildpackToInsert := &models.Buildpack{
		GUID:     uuid.New().String(),
		Name:     requestBody.Name,
		Position: requestBody.Position,
		Enabled:  null.BoolFrom(requestBody.Enabled),
		Locked:   null.BoolFrom(requestBody.Locked),
		Stack:    null.StringFromPtr(requestBody.Stack),
	}

	buildpacksInDB, errDB := buildpackQuerier(qm.OrderBy(models.BuildpackTableColumns.Position)).All(boilCtx, cont.DB)
	if errDB != nil {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", errDB))
	}

	maxPos := buildpacksInDB[len(buildpacksInDB)-1].Position
	if maxPos+1 < buildpackToInsert.Position {
		buildpackToInsert.Position = maxPos + 1
	} else {
		for _, bp := range buildpacksInDB {
			if bp.Position >= buildpackToInsert.Position {
				bp.Position++
				_, err = buildpackUpdater.Update(bp, boilCtx, dbTx, boil.Whitelist(models.BuildpackColumns.Position))
				if err != nil {
					logger.Error("Buildpack position cannot be updated", zap.String("buildpack", bp.Name))
					return v3.UnknownError(err)
				}
			}
		}
	}

	err = buildpackInserter.Insert(buildpackToInsert, boilCtx, dbTx, boil.Infer())
	if err != nil {
		logger.Error("Buildpack could not be inserted", zap.String("buildpack", buildpackToInsert.Name))
		return v3.UnknownError(err)
	}

	err = dbTx.Commit()
	if err != nil {
		logger.Error("DB transaction failed")
		return v3.UnknownError(err)
	}

	response, err := cont.Presenter.ResponseObject(buildpackToInsert, v3.GetResourcePath(echoCtx))
	if err != nil {
		return v3.UnknownError(err)
	}
	return echoCtx.JSON(http.StatusOK, response)
}
