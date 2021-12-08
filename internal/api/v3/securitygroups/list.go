package securitygroups

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/pagination"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/timefilters"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/sqlhelpers"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type FilterParams struct {
	Names   *string `query:"names" validate:"omitempty"`
	OrderBy *string `query:"order_by" validate:"omitempty,oneof=-created_at created_at -updated_at updated_at"`
}

// List godoc
// @Summary Security Groups List security groups
// @Description Retrieve all security groups
// @Tags SecurityGroups
// @Accept json
// @Produce json
// @Success 200 {object} []Response
// @Success 404 {object} interface{}
// @Failure 400 {object} []interface{}
// @Failure 500 {object} v3.CfAPIError
// @Router /securitygroups [get]

func (cont *Controller) List(echoCtx echo.Context) error {
	logger := logging.FromContext(echoCtx)
	paginationParams := pagination.Default()
	filterParams := FilterParams{}

	createdAts, updatedAts, err := timefilters.ParseTimeFilters(echoCtx)
	if err != nil {
		return v3.BadQueryParameter(err)
	}
	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(echoCtx, &paginationParams); err != nil {
		return v3.BadQueryParameter(err)
	}
	if errFilters := (&echo.DefaultBinder{}).BindQueryParams(echoCtx, &filterParams); errFilters != nil {
		return v3.BadQueryParameter(errFilters)
	}
	err = validator.New().Struct(paginationParams)
	if err != nil {
		return v3.BadQueryParameter(err)
	}
	errFilter := validator.New().Struct(filterParams)
	if errFilter != nil {
		return v3.BadQueryParameter(errFilter)
	}

	boilCtx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	var mods []qm.QueryMod
	mods = append(mods, timefilters.Filters(createdAts, updatedAts, models.SecurityGroupTableColumns.CreatedAt, models.SecurityGroupTableColumns.UpdatedAt)...)

	permissionsMods, err := cont.readPermissionMods(echoCtx)
	if err != nil {
		return v3.UnknownError(err)
	}
	mods = append(mods, permissionsMods...)
	mods = append(mods, filters(filterParams)...)

	totalResults, err := securityGroupQuerier(mods...).Count(boilCtx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}

	mods = append(mods,
		qm.Limit(int(paginationParams.PerPage)),
		qm.Offset((paginationParams.Page-1)*int(paginationParams.PerPage)),
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
	)

	securityGroups, err := securityGroupQuerier(mods...).All(boilCtx, cont.DB)
	if err != nil && err != sql.ErrNoRows {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", err))
	}

	response, err := cont.Presenter.ListResponseObject(securityGroups, totalResults, paginationParams, v3.GetResourcePath(echoCtx))
	if err != nil {
		return v3.UnknownError(err)
	}

	return echoCtx.JSON(http.StatusOK, response)
}

func filters(filters FilterParams) []qm.QueryMod {
	filterMods := []qm.QueryMod{}

	if filters.Names != nil {
		names := sqlhelpers.Split(*filters.Names)
		filterMods = append(filterMods, sqlhelpers.WhereIn(models.SecurityGroupTableColumns.Name, names))
	}

	if filters.OrderBy != nil {
		filterMods = append(filterMods, sqlhelpers.OrderBy(*filters.OrderBy))
	}

	return filterMods
}
