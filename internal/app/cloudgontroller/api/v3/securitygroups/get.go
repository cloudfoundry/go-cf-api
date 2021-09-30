package securitygroups

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	v3 "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/api/v3"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/logging"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/permissions"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
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

	// Only add subqueries and joins when the user isn't an admin
	if !(auth.IsAdmin(c) || auth.IsAdminReadOnly(c) || auth.IsGlobalAuditor(c)) {
		userGUID, ok := c.Get(auth.Username).(string)
		if !ok {
			return v3.UnknownError(fmt.Errorf("could not get username from context"))
		}
		roles := append(permissions.AllSpaceRoles, permissions.OrgManager) //nolint:gocritic
		allowedSpaceIDs, err := cont.Permissions.AllowedSpaceIDsForUser(userGUID, roles...)
		if err != nil {
			return v3.UnknownError(fmt.Errorf("error getting allowed space IDs for user %s: %w", userGUID, err))
		}

		mods = append(mods, qm.Expr(
			qm.Or2(models.SecurityGroupWhere.RunningDefault.EQ(null.BoolFrom(true))),
			qm.Or2(models.SecurityGroupWhere.StagingDefault.EQ(null.BoolFrom(true))),
			allowedSpaceIDs.Contains(models.SecurityGroupsSpaceTableColumns.SpaceID),
			allowedSpaceIDs.Contains(models.StagingSecurityGroupsSpaceTableColumns.StagingSpaceID),
		))
		mods = append(mods, qm.LeftOuterJoin(fmt.Sprintf(
			"%s ON %s = %s",
			models.Quote(models.TableNames.SecurityGroupsSpaces),
			models.Quote(models.SecurityGroupTableColumns.ID),
			models.Quote(models.SecurityGroupsSpaceTableColumns.SecurityGroupID)),
		))
		mods = append(mods, qm.LeftOuterJoin(fmt.Sprintf(
			"%s ON %s = %s",
			models.Quote(models.TableNames.StagingSecurityGroupsSpaces),
			models.Quote(models.SecurityGroupTableColumns.ID),
			models.Quote(models.StagingSecurityGroupsSpaceTableColumns.StagingSecurityGroupID)),
		))
		mods = append(mods, allowedSpaceIDs.With()...)
	}

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
