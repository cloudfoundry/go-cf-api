//go:build unit
// +build unit

package sqlhelpers_test

import (
	"fmt"
	"testing"

	"github.com/cloudfoundry/go-cf-api/internal/storage/db/models"
	. "github.com/cloudfoundry/go-cf-api/internal/storage/db/sqlhelpers"
	"github.com/stretchr/testify/assert"
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

	for testCaseName, testCase := range cases {
		t.Run(testCaseName, func(t *testing.T) {
			fu := FromUnion(testCase.unionTables)
			sq := models.NewSubquery()
			fu.Apply(sq.Query)
			query, args := sq.SQL()
			assert.Empty(t, args)
			assert.Equal(t, query, testCase.expectedQuery)
		})
	}
}
