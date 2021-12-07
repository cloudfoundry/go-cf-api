package securitygroups

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/auth"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/apicommon/v3/permissions"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

const GUIDParam = "guid"

//nolint:gochecknoglobals // here to be overridden in tests
var (
	securityGroupQuerier = func(qm ...qm.QueryMod) models.SecurityGroupFinisher { return models.SecurityGroups(qm...) }
)

type Controller struct {
	DB          *sql.DB
	Presenter   Presenter
	Permissions permissions.Querier
}

func (cont *Controller) readPermissionMods(c echo.Context) ([]qm.QueryMod, error) {
	var mods []qm.QueryMod
	if !(auth.IsAdmin(c) || auth.IsAdminReadOnly(c) || auth.IsGlobalAuditor(c)) {
		userGUID, ok := c.Get(auth.Username).(string)
		if !ok {
			return mods, fmt.Errorf("could not get username from context")
		}
		roles := append(permissions.AllSpaceRoles, permissions.OrgManager) //nolint:gocritic
		allowedSpaceIDs, err := cont.Permissions.AllowedSpaceIDsForUser(userGUID, roles...)
		if err != nil {
			return mods, fmt.Errorf("error getting allowed space IDs for user %s: %w", userGUID, err)
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
	return mods, nil
}
