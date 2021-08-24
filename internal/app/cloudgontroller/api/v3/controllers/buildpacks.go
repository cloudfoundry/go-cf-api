package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/controllers/common"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/presenter"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
	"net/http"
	"strings"
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
	filters := DefaultFilters()
	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &pagination); err != nil {
		return BadQueryParameter(err)
	}

	if errFilters := (&echo.DefaultBinder{}).BindQueryParams(c, &filters); errFilters != nil {
		return BadQueryParameter(errFilters)
	}

	err := validator.New().Struct(pagination)
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
	mods = append(mods, BuildFilters(filters)...)

	buildpacks, err := models.Buildpacks(
		mods...,
	).All(ctx, cont.DB)
	if err != nil {
		return UnknownError(fmt.Errorf("could not Select: %w", err))
	}
	if buildpacks == nil {
		return c.JSON(http.StatusNotFound, []presenter.BuildpackResponse{})
	}
	// If the order is reversed by the order_by parameter starting with a -
	if strings.HasPrefix(filters.OrderBy, "-"){
		for i, j := 0, len(buildpacks)-1; i < j; i, j = i+1, j-1 {
			buildpacks[i], buildpacks[j] = buildpacks[j], buildpacks[i]
		}
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


func BuildFilters (filters FilterParams) []qm.QueryMod {
	var filterMods []qm.QueryMod

	// Make sure that the sorting is done by one of the supported values.
	possibleSortValues := []string{"created_at", "updated_at", "position"}
	if strings.HasPrefix(filters.OrderBy, "-") {
		filters.OrderBy = strings.TrimPrefix(filters.OrderBy, "-")
	}
	for _,value  := range possibleSortValues {
		if value == filters.OrderBy {
			filterMods = append(filterMods, qm.OrderBy(filters.OrderBy))
		}
	}

	names := strings.Split(filters.Names, ",")
	for index, name := range names {
		if name != "" {
			if index == 0{
				filterMods = append(filterMods, qm.And("name=?", name))
			} else {
				filterMods = append(filterMods, qm.Or("name=?", name))
			}
		}
	}

	stacks := strings.Split(filters.Stacks, ",")
	for index, stack := range stacks {
		if stack != "" {
			if index == 0 {
				filterMods = append(filterMods, qm.And("stack=?", stack))
			} else {
				filterMods = append(filterMods, qm.Or("stack=?", stack))
			}
		}
	}


	return filterMods
}

type FilterParams struct {
	Names   		string `query:"names"`
	Stacks 			string `query:"stacks"`
	OrderBy 		string `query:"order_by"`
	LabelSelector 	string `query:"label_selector"`
	CreatedAts 		null.Time `query:"created_ats"`
	UpdatedAts 		null.Time `query:"updated_ats"`

}

func DefaultFilters() FilterParams {
	return FilterParams{OrderBy: "position"} //nolint:gomnd // Default values
}