package buildpacks

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
)

const GUIDParam = "guid"

var (
	buildpackQuerier                           = func(qm ...qm.QueryMod) models.BuildpackFinisher { return models.Buildpacks(qm...) }
	buildpackInserter models.BuildpackInserter = models.Buildpacks()
)

type Controller struct {
	DB *sql.DB
}

// List godoc
// @Summary Buildpacks List buildpacks
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Accept json
// @Produce json
// @Success 200 {object} []Response
// @Success 404 {object} interface{}
// @Failure 400 {object} []interface{}
// @Failure 500 {object} CloudControllerError
// @Router /buildpacks [get]
func (cont *Controller) List(c echo.Context) error {
	logger := logging.FromContext(c)
	pagination := pagination.Default()
	filters := DefaultFilters()

	createdAts, updatedAts, err := v3.ParseTimeFilters(c)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		return v3.BadQueryParameter(err)
	}
	if errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); errFilters != nil {
		return v3.BadQueryParameter(errFilters)
	}

	err = validator.New().Struct(pagination)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	errFilter := validator.New().Struct(filters)
	if errFilter != nil {
		return v3.BadQueryParameter(errFilter)
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	_, err = buildpackQuerier().Count(ctx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}

	var mods []qm.QueryMod
	mods = append(mods, qm.Limit(int(pagination.PerPage)))
	mods = append(mods, qm.Offset((pagination.Page-1)*int(pagination.PerPage)))
	// Append Filters to the query
	mods = append(mods, buildFilters(filters, createdAts, updatedAts)...)

	buildpacks, err := buildpackQuerier(
		mods...,
	).All(ctx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", err))
	}
	if buildpacks == nil {
		return c.JSON(http.StatusNotFound, []Response{})
	}

	return c.JSON(http.StatusOK, ListResponseObject(buildpacks, pagination, v3.GetResourcePath(c)))
}

// Get godoc
// @Summary Show a buildpack
// @Description Retrieve all buildpacks the user has access to.
// @Tags Buildpacks
// @Param guid path string true "Buildpack GUID"
// @Success 200 {object} Response
// @Success 404 {object} interface{}
// @Failure 400 {object} CloudControllerError
// @Failure 500 {object} CloudControllerError
// @Router /buildpacks/{guid} [get]
func (cont *Controller) Get(c echo.Context) error {
	guid := c.Param(GUIDParam)
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	buildpack, err := buildpackQuerier(qm.Where("guid=?", guid)).One(ctx, cont.DB)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
		if errors.Cause(err) != sql.ErrNoRows {
			return v3.UnknownError(err)
		}
	}
	if buildpack == nil {
		return v3.ResourceNotFound("buildpack", err)
	}

	return c.JSON(http.StatusOK, ResponseObject(buildpack, v3.GetResourcePath(c)))
}

// PostBuildpack godoc
// @Summary Create a buildpack
// @Description Create a new buildpack
// @Tags Buildpacks
// @accept json
// @produce json
// @Success 201 {object} Response
// @Success 404 {object} interface{}
// @Failure 400 {object} CloudControllerError
// @Failure 500 {object} CloudControllerError
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
	return c.JSON(http.StatusOK, ResponseObject(buildpackToInsert, v3.GetResourcePath(c)))
}

func buildFilters(filters FilterParams, createdAts, updatedAts []v3.TimeFilter) []qm.QueryMod {
	filterMods := []qm.QueryMod{}

	direction := "ASC"
	if strings.HasPrefix(filters.OrderBy, "-") {
		direction = "DESC"
	}
	filterMods = append(filterMods, qm.OrderBy(fmt.Sprintf("%s %s", strings.TrimPrefix(filters.OrderBy, "-"), direction)))

	names := strings.FieldsFunc(filters.Names, splitWithoutEmptyString)
	if len(names) > 0 {
		filterMods = append(filterMods, whereIn(models.BuildpackColumns.Name, names))
	}

	stacks := strings.FieldsFunc(filters.Stacks, splitWithoutEmptyString)
	if len(stacks) > 0 {
		filterMods = append(filterMods, whereIn(models.BuildpackColumns.Stack, stacks))
	}

	for _, createdAt := range createdAts {
		operator := v3.Operators[createdAt.Operator]
		filterMods = append(filterMods, qm.Where(fmt.Sprintf("%s %s ?", models.BuildpackColumns.CreatedAt, operator), createdAt.Timestamp.Format(time.RFC3339)))
	}

	for _, updatedAt := range updatedAts {
		operator := v3.Operators[updatedAt.Operator]
		filterMods = append(filterMods, qm.Where(fmt.Sprintf("%s %s ?", models.BuildpackColumns.UpdatedAt, operator), updatedAt.Timestamp.Format(time.RFC3339)))
	}

	return filterMods
}

func whereIn(field string, slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", field), values...)
}

func splitWithoutEmptyString(c rune) bool { return c == ',' }

type FilterParams struct {
	Names         string `query:"names"`
	Stacks        string `query:"stacks"`
	OrderBy       string `query:"order_by" validate:"oneof=created_at -created_at updated_at -updated_at position -position"`
	LabelSelector string `query:"label_selector"`
}

func DefaultFilters() FilterParams {
	return FilterParams{OrderBy: "position"}
}
