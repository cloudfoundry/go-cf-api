package securitygroups

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/permissions"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
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
