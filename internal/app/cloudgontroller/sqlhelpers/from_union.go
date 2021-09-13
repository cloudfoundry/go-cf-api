package sqlhelpers

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	models "github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/sqlboiler"
)

func FromUnion(unionTables []UnionTable) qm.QueryMod {
	return &fromUnion{
		unionTables: unionTables,
	}
}

type UnionTable struct {
	Column string
	Table  string
}

type fromUnion struct {
	unionTables []UnionTable
}

func (fu *fromUnion) Apply(q *queries.Query) {
	if len(fu.unionTables) == 0 {
		queries.SetSQL(q, "SELECT -1")
		return
	}
	from := models.Quote(fu.unionTables[0].Table)
	for _, table := range fu.unionTables[1:] {
		from += fmt.Sprintf(" UNION SELECT %s FROM %s", models.Quote(table.Column), models.Quote(table.Table))
	}
	queries.SetSelect(q, []string{fu.unionTables[0].Column})
	queries.SetFrom(q, from)
}
