package buildpacks

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/api/v3/timefilters"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/sqlhelpers"
)

type FilterParams struct {
	Names         *string `query:"names"`
	Stacks        *string `query:"stacks"`
	OrderBy       string  `query:"order_by" validate:"oneof=created_at -created_at updated_at -updated_at position -position"`
	LabelSelector string  `query:"label_selector"`
}

func DefaultFilters() FilterParams {
	return FilterParams{OrderBy: "position"}
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
// @Failure 500 {object} v3.CfApiError
// @Router /buildpacks [get]

func (cont *Controller) List(c echo.Context) error {
	logger := logging.FromContext(c)
	pagination := pagination.Default()
	filterParams := DefaultFilters()

	createdAts, updatedAts, err := timefilters.ParseTimeFilters(c)
	if err != nil {
		return v3.BadQueryParameter(err)
	}
	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		return v3.BadQueryParameter(err)
	}
	if errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filterParams); errFilters != nil {
		return v3.BadQueryParameter(errFilters)
	}
	err = validator.New().Struct(pagination)
	if err != nil {
		return v3.BadQueryParameter(err)
	}
	errFilter := validator.New().Struct(filterParams)
	if errFilter != nil {
		return v3.BadQueryParameter(errFilter)
	}
	labelSelectors, err := cont.LabelSelectorParser.Parse(filterParams.LabelSelector)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	var mods []qm.QueryMod
	mods = append(mods, filters(filterParams)...)
	mods = append(mods, timefilters.Filters(createdAts, updatedAts, models.BuildpackTableColumns.CreatedAt, models.BuildpackTableColumns.UpdatedAt)...)
	mods = append(mods, labelSelectors.Filters(models.TableNames.Buildpacks, models.TableNames.BuildpackLabels)...)

	totalResults, err := buildpackQuerier(mods...).Count(ctx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}

	bplCols, bpaCols := models.BuildpackLabelColumns, models.BuildpackAnnotationColumns
	mods = append(mods,
		sqlhelpers.OrderBy(filterParams.OrderBy),
		qm.Limit(int(pagination.PerPage)),
		qm.Offset((pagination.Page-1)*int(pagination.PerPage)),
		qm.Load(models.BuildpackRels.ResourceBuildpackLabels, qm.Select(bplCols.KeyName, bplCols.Value, bplCols.ResourceGUID)),
		qm.Load(models.BuildpackRels.ResourceBuildpackAnnotations, qm.Select(bpaCols.Key, bpaCols.Value, bpaCols.ResourceGUID)),
	)

	buildpacks, err := buildpackQuerier(mods...).All(ctx, cont.DB)
	if err != nil && err != sql.ErrNoRows {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", err))
	}

	response, err := cont.Presenter.ListResponseObject(buildpacks, totalResults, pagination, v3.GetResourcePath(c))
	if err != nil {
		return v3.UnknownError(err)
	}
	return c.JSON(http.StatusOK, response)
}

func filters(filters FilterParams) []qm.QueryMod {
	filterMods := []qm.QueryMod{}

	if filters.Names != nil {
		names := sqlhelpers.Split(*filters.Names)
		filterMods = append(filterMods, sqlhelpers.WhereIn(models.BuildpackTableColumns.Name, names))
	}

	if filters.Stacks != nil {
		stacks := sqlhelpers.Split(*filters.Stacks)
		filterMods = append(filterMods, sqlhelpers.WhereIn(models.BuildpackTableColumns.Stack, stacks))
	}

	return filterMods
}
