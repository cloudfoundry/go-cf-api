package sqlhelpers

import (
	"fmt"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
)

func WhereIn(field string, slice []string) qm.QueryMod {
	if len(slice) == 0 {
		return qmhelper.WhereIsNull(field)
	}
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", field), values...)
}

func OrderBy(orderBy string) qm.QueryMod {
	direction := "ASC"
	if strings.HasPrefix(orderBy, "-") {
		direction = "DESC"
	}
	return qm.OrderBy(fmt.Sprintf("%s %s", strings.TrimPrefix(orderBy, "-"), direction))
}

// Splits a string on comma without including empty strings
func Split(s string) []string {
	return strings.FieldsFunc(s, split)
}

func split(c rune) bool { return c == ',' }
