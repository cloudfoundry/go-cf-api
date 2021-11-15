package timefilters

import (
	"fmt"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func Filters(createdAts, updatedAts []TimeFilter, createdAtColumn string, updatedAtColumn string) []qm.QueryMod {
	filterMods := []qm.QueryMod{}

	for _, createdAt := range createdAts {
		operator := Operators[createdAt.Operator]
		filterMods = append(filterMods, qm.Where(fmt.Sprintf("%s %s ?", createdAtColumn, operator), createdAt.Timestamp.Format(time.RFC3339)))
	}

	for _, updatedAt := range updatedAts {
		operator := Operators[updatedAt.Operator]
		filterMods = append(filterMods, qm.Where(fmt.Sprintf("%s %s ?", updatedAtColumn, operator), updatedAt.Timestamp.Format(time.RFC3339)))
	}

	return filterMods
}
