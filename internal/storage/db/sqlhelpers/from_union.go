package sqlhelpers

import (
	"fmt"

	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (fu *fromUnion) Apply(query *queries.Query) {
	if len(fu.unionTables) == 0 {
		queries.SetSQL(query, "SELECT -1")
		return
	}
	from := models.Quote(fu.unionTables[0].Table)
	for _, table := range fu.unionTables[1:] {
		from += fmt.Sprintf(" UNION SELECT %s FROM %s", models.Quote(table.Column), models.Quote(table.Table))
	}
	queries.SetSelect(query, []string{fu.unionTables[0].Column})
	queries.SetFrom(query, from)
}
