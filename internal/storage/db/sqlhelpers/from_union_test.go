//go:build unit
// +build unit

package sqlhelpers_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/sqlhelpers"
)

func TestFromUnion(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		unionTables   []UnionTable
		expectedQuery string
	}{
		"1 table": {
			unionTables:   []UnionTable{{Table: "table", Column: "col"}},
			expectedQuery: fmt.Sprintf("SELECT %s FROM %s", models.Quote("col"), models.Quote("table")),
		},
		"2 tables": {
			unionTables: []UnionTable{{Table: "table1", Column: "col"}, {Table: "table2", Column: "col"}},
			expectedQuery: fmt.Sprintf("SELECT %s FROM %s UNION SELECT %s FROM %s",
				models.Quote("col"), models.Quote("table1"),
				models.Quote("col"), models.Quote("table2"),
			),
		},
		"3 tables": {
			unionTables: []UnionTable{{Table: "table1", Column: "col"}, {Table: "table2", Column: "col"}, {Table: "table3", Column: "col"}},
			expectedQuery: fmt.Sprintf("SELECT %s FROM %s UNION SELECT %s FROM %s UNION SELECT %s FROM %s",
				models.Quote("col"), models.Quote("table1"),
				models.Quote("col"), models.Quote("table2"),
				models.Quote("col"), models.Quote("table3"),
			),
		},
		"no tables": {
			unionTables:   []UnionTable{},
			expectedQuery: "SELECT -1",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			fu := FromUnion(tc.unionTables)
			sq := models.NewSubquery()
			fu.Apply(sq.Query)
			query, args := sq.SQL()
			assert.Empty(t, args)
			assert.Equal(t, query, tc.expectedQuery)
		})
	}
}
