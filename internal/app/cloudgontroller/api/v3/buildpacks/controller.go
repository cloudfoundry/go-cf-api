package buildpacks

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

const GUIDParam = "guid"

//nolint:gochecknoglobals // here to be overridden in tests
var (
	buildpackQuerier                           = func(qm ...qm.QueryMod) models.BuildpackFinisher { return models.Buildpacks(qm...) }
	buildpackInserter models.BuildpackInserter = models.Buildpacks()
)

type Controller struct {
	DB *sql.DB
}
