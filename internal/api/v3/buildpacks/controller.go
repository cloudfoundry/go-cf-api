package buildpacks

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/cloudfoundry/go-cf-api/internal/apicommon/v3/metadata"
	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
)

const GUIDParam = "guid"

//nolint:gochecknoglobals // here to be overridden in tests
var (
	buildpackQuerier                           = func(qm ...qm.QueryMod) models.BuildpackFinisher { return models.Buildpacks(qm...) }
	buildpackInserter models.BuildpackInserter = models.Buildpacks()
)

type Controller struct {
	DB                  *sql.DB
	Presenter           Presenter
	LabelSelectorParser metadata.LabelSelectorParser
}
