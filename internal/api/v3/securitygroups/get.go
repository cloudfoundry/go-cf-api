package securitygroups

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	v3 "github.com/cloudfoundry/go-cf-api/internal/apicommon/v3"
	"github.com/cloudfoundry/go-cf-api/internal/logging"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

func (cont *Controller) Get(echoCtx echo.Context) error {
	guid := echoCtx.Param(GUIDParam)
	logger := logging.FromContext(echoCtx)

	boilCtx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(false, logger))
	mods := []qm.QueryMod{
		models.SecurityGroupWhere.GUID.EQ(guid),
		qm.Load(qm.Rels(models.SecurityGroupRels.SecurityGroupsSpaces, models.SecurityGroupsSpaceRels.Space)),
		qm.Load(qm.Rels(models.SecurityGroupRels.StagingSecurityGroupStagingSecurityGroupsSpaces, models.StagingSecurityGroupsSpaceRels.StagingSpace)),
	}

	permissionsMods, err := cont.readPermissionMods(echoCtx)
	if err != nil {
		return v3.UnknownError(err)
	}

	mods = append(mods, permissionsMods...)

	securityGroup, err := securityGroupQuerier(mods...).One(boilCtx, cont.DB)
	if err != nil {
		logger.Error("Couldn't select", zap.Error(err))
		if errors.Cause(err) != sql.ErrNoRows {
			return v3.UnknownError(err)
		}
	}
	if securityGroup == nil {
		return v3.ResourceNotFound("security_group", err)
	}
	responseObject, err := cont.Presenter.ResponseObject(securityGroup, v3.GetResourcePath(echoCtx))
	if err != nil {
		return v3.UnknownError(fmt.Errorf("could not construct response: %w", err))
	}

	return echoCtx.JSON(http.StatusOK, responseObject)
}
