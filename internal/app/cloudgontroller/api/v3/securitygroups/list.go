package securitygroups

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/pagination"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3/timefilters"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

type FilterParams struct {
	Names   string `query:"names"`
	OrderBy string `query:"order_by" validate:"oneof=-created_at created_at -updated_at updated_at"`
}

func DefaultFilters() FilterParams {
	return FilterParams{OrderBy: "updated_at"}
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
// @Failure 500 {object} v3.CloudControllerError
// @Router /securitygroups [get]

func (cont *Controller) List(c echo.Context) error {
	logger := logging.FromContext(c)
	paginationParams := pagination.Default()
	filterParams := DefaultFilters()

	createdAts, updatedAts, err := timefilters.ParseTimeFilters(c)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	// BindQueryParams will overwrite default values if params were given
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &paginationParams); err != nil {
		return v3.BadQueryParameter(err)
	}

	err = validator.New().Struct(paginationParams)
	if err != nil {
		return v3.BadQueryParameter(err)
	}

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true, logger))
	var mods []qm.QueryMod
	mods = append(mods, timefilters.Filters(createdAts, updatedAts, models.SecurityGroupTableColumns.CreatedAt, models.SecurityGroupTableColumns.UpdatedAt)...)

	permissionsMods, err := cont.readPermissionMods(c)
	if err != nil {
		return v3.UnknownError(err)
	}

	mods = append(mods, permissionsMods...)

	totalResults, err := securityGroupQuerier(mods...).Count(ctx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("couldn't fetch all rows: %w", err))
	}

	mods = append(mods,
		orderBy(filterParams.OrderBy),
		qm.Limit(int(paginationParams.PerPage)),
		qm.Offset((paginationParams.Page-1)*int(paginationParams.PerPage)),
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
	)

	securityGroups, err := securityGroupQuerier(mods...).All(ctx, cont.DB)
	if err != nil {
		return v3.UnknownError(fmt.Errorf("could not Select: %w", err))
	}

	if securityGroups == nil {
		return c.JSON(http.StatusNotFound, []Response{})
	}

	response, err := cont.Presenter.ListResponseObject(securityGroups, totalResults, paginationParams, v3.GetResourcePath(c))
	if err != nil {
		return v3.UnknownError(err)
	}

	return c.JSON(http.StatusOK, response)
}

func orderBy(orderBy string) qm.QueryMod {
	direction := "ASC"
	if strings.HasPrefix(orderBy, "-") {
		direction = "DESC"
	}
	return qm.OrderBy(fmt.Sprintf("%s %s", strings.TrimPrefix(orderBy, "-"), direction))
}
