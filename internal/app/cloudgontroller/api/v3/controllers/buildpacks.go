package controllers

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

type BuildpackParams struct {
	Name           string      `json:"name"`
	//Position       int         `boil:"position" json:"position" toml:"position" yaml:"position"`
	//Enabled        null.Bool   `boil:"enabled" json:"enabled,omitempty" toml:"enabled" yaml:"enabled,omitempty"`
	//Locked         null.Bool   `boil:"locked" json:"locked,omitempty" toml:"locked" yaml:"locked,omitempty"`
	//Stack          null.String `boil:"stack" json:"stack,omitempty" toml:"stack" yaml:"stack,omitempty"`
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
func (cont *BuildpackController) GetBuildpacks(c echo.Context) error { //nolint:cyclop // function is logically together
	logger := logging.FromContext(c)
	pagination := common.DefaultPagination()
	filters := DefaultFilters()

	createdAts, updatedAts, err := parseTimeFilters(c)
	if err != nil {
		return BadQueryParameter(err)
	}

	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		return BadQueryParameter(err)
	}
	if errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); errFilters != nil {
		return BadQueryParameter(errFilters)
	}

	err = validator.New().Struct(pagination)
	if err != nil {
		return BadQueryParameter(err)
	}

	errFilter := validator.New().Struct(filters)
	if errFilter != nil {
		return BadQueryParameter(errFilter)
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	_, err = models.Buildpacks().Count(ctx, cont.DB)
	if err != nil {
		return UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}

	var mods []qm.QueryMod
	mods = append(mods, qm.Limit(int(pagination.PerPage)))
	mods = append(mods, qm.Offset((pagination.Page-1)*int(pagination.PerPage)))
	// Append Filters to the query
	mods = append(mods, buildFilters(filters, createdAts, updatedAts)...)

	buildpacks, err := models.Buildpacks(
		mods...,
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
// PostBuildpack godoc
// @Summary Create a buildpack
// @Description Create a new buildpack
// @Tags Buildpacks
// @accespt json
// @produce json
// @Param guid path string true "Buildpack GUID"
// @Success 200 {object} presenter.BuildpackResponse
// @Success 404 {object} interface{}
// @Failure 400 {object} CloudControllerError
// @Failure 500 {object} CloudControllerError
// @Router /buildpacks/{guid} [get]
func (cont *BuildpackController) PostBuildpack(c echo.Context) error {

	logger := logging.FromContext(c)
	//body := models.Buildpack{}
	var body *models.Buildpack

	//body := new(BuildpackParams)
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		logger.Error("Bad reqeust 400")
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	//b, _ := json.Marshal(body)
	err := body.Insert(ctx, cont.DB, boil.Infer())
	if err != nil {
		logger.Error("There is no params")
		}

	createdBuildpack, err := models.Buildpacks(qm.Where("name=?", body.Name)).One(ctx, cont.DB)
	return c.JSON(http.StatusOK, presenter.BuildpackResponseObject(createdBuildpack, GetResourcePath(c)))
}

func buildFilters(filters FilterParams, createdAts, updatedAts []TimeFilter) []qm.QueryMod {
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
		operator := operators[createdAt.Operator]
		filterMods = append(filterMods, qm.Where(fmt.Sprintf("%s %s ?", models.BuildpackColumns.CreatedAt, operator), createdAt.Timestamp.Format(time.RFC3339)))
	}

	for _, updatedAt := range updatedAts {
		operator := operators[updatedAt.Operator]
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
