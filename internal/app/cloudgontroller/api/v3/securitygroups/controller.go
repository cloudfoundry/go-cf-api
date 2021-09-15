package securitygroups

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
	"go.uber.org/zap"
)

const GUIDParam = "guid"

//nolint:gochecknoglobals // here to be overridden in tests
var securityGroupQuerier = func(qm ...qm.QueryMod) models.SecurityGroupFinisher { return models.SecurityGroups(qm...) }

type Controller struct {
	DB        *sql.DB
	Presenter Presenter
}

func (cont *Controller) Get(c echo.Context) error {
	guid := c.Param(GUIDParam)
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	securityGroup, err := securityGroupQuerier(
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
		models.SecurityGroupWhere.GUID.EQ(guid),
	).One(ctx, cont.DB)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
		if errors.Cause(err) != sql.ErrNoRows {
			return v3.UnknownError(err)
		}
	}
	if securityGroup == nil {
		return v3.ResourceNotFound("security_group", err)
	}
	responseObject, err := cont.Presenter.ResponseObject(securityGroup, v3.GetResourcePath(c))
	if err != nil {
		return v3.UnknownError(fmt.Errorf("could not construct response: %w", err))
	}

	return c.JSON(http.StatusOK, responseObject)
}
