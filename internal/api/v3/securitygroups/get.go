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
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
	"go.uber.org/zap"
)

func (cont *Controller) Get(c echo.Context) error {
	guid := c.Param(GUIDParam)
	logger := logging.FromContext(c)

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	mods := []qm.QueryMod{
		models.SecurityGroupWhere.GUID.EQ(guid),
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
	}

	permissionsMods, err := cont.readPermissionMods(c)
	if err != nil {
		return v3.UnknownError(err)
	}

	mods = append(mods, permissionsMods...)

	securityGroup, err := securityGroupQuerier(mods...).One(ctx, cont.DB)
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
