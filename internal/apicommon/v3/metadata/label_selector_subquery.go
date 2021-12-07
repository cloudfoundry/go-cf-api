package metadata

import (
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/storage/db/models"
)

//go:generate mockgen -source=$GOFILE -destination=mocks/$GOFILE
type LabelSelectorFilters interface {
	Filters(resourceTable, labelsTable string) []qm.QueryMod
}

//nolint:gochecknoglobals  // cannot be const because of function calls
var (
	// All label tables should be the same so just pick a resource to get the column names from
	guidColumn         = models.Quote(models.AppColumns.GUID)
	resourceGUIDColumn = models.Quote(models.AppLabelColumns.ResourceGUID)
	keyColumn          = models.Quote(models.AppLabelColumns.KeyName)
	valueColumn        = models.Quote(models.AppLabelColumns.Value)
)

type (
	exists    []Key
	notExists []Key
	inEquals  struct {
		in []LabelSelector
		eq []LabelSelector
	}
	notInEquals struct {
		notIn []LabelSelector
		notEq []LabelSelector
	}
)

func (selectors LabelSelectorSlice) Filters(resourceTable, labelsTable string) []qm.QueryMod {
	var mods []qm.QueryMod

	// Return empty mods if no labels are provided
	if len(selectors) == 0 {
		return mods
	}

	var exists exists
	var notExists notExists
	var inEquals inEquals
	var notInEquals notInEquals

	// Group selectors into operators that can be combined so we only need <=4 subqueries
	for _, labelSelector := range selectors {
		switch labelSelector.Operator {
		case Exists:
			exists = append(exists, labelSelector.Key)
		case NotExists:
			notExists = append(notExists, labelSelector.Key)
		case Eq, EqAlt:
			inEquals.eq = append(inEquals.eq, labelSelector)
		case In:
			inEquals.in = append(inEquals.in, labelSelector)
		case NotEq:
			notInEquals.notEq = append(notInEquals.notEq, labelSelector)
		case NotIn:
			notInEquals.notIn = append(notInEquals.notIn, labelSelector)
		}
	}

	if len(exists) > 0 {
		mods = append(mods, exists.Filter(resourceTable, labelsTable))
	}
	if len(notExists) > 0 {
		mods = append(mods, notExists.Filter(resourceTable, labelsTable))
	}
	if inEquals.any() {
		mods = append(mods, inEquals.Filter(resourceTable, labelsTable))
	}
	if notInEquals.any() {
		mods = append(mods, notInEquals.Filter(resourceTable, labelsTable))
	}
	return mods
}

func (s exists) Filter(resourceTable, labelsTable string) qm.QueryMod {
	qms := subqueryBase(resourceTable, labelsTable)
	// Add AND (key_name = 'key') for each exists
	for _, key := range s {
		qms = append(qms, qm.Expr(qmhelper.Where(keyColumn, qmhelper.EQ, key)))
	}
	subquery, args := models.NewSubquery(qms...).SQL()
	return qm.Where(fmt.Sprintf("EXISTS (%s)", subquery), args...)
}

func (s notExists) Filter(resourceTable, labelsTable string) qm.QueryMod {
	qms := subqueryBase(resourceTable, labelsTable)
	qms = append(qms, qm.Expr(whereIn(keyColumn, s)))
	subquery, args := models.NewSubquery(qms...).SQL()
	return qm.Where(fmt.Sprintf("NOT EXISTS (%s)", subquery), args...)
}

func (s inEquals) any() bool {
	return len(s.in) > 0 || len(s.eq) > 0
}

func (s inEquals) Filter(resourceTable, labelsTable string) qm.QueryMod {
	qms := subqueryBase(resourceTable, labelsTable)
	for _, selector := range s.eq {
		// Add AND (key_name = 'key' AND value = 'value') for each ==
		qms = append(qms, qm.Expr(
			qm.Where(fmt.Sprintf("%s = ?", keyColumn), selector.Key),
			qm.And(fmt.Sprintf("%s = ?", valueColumn), selector.Values[0]),
		))
	}
	for _, selector := range s.in {
		// Add AND (key_name = 'key' AND value IN ('value1','value2')) for each in
		qms = append(qms, qm.Expr(
			qm.Where(fmt.Sprintf("%s = ?", keyColumn), selector.Key),
			andIn(valueColumn, selector.Values),
		))
	}
	subquery, args := models.NewSubquery(qms...).SQL()
	return qm.Where(fmt.Sprintf("EXISTS (%s)", subquery), args...)
}

func (s notInEquals) any() bool {
	return len(s.notIn) > 0 || len(s.notEq) > 0
}

func (s notInEquals) Filter(resourceTable, labelsTable string) qm.QueryMod {
	ors := []qm.QueryMod{}
	for _, selector := range s.notEq {
		// Add OR (key_name = 'key' AND value = 'value') for each !=
		ors = append(ors, qm.Or2(qm.Expr(
			qm.Where(fmt.Sprintf("%s = ?", keyColumn), selector.Key),
			qm.And(fmt.Sprintf("%s = ?", valueColumn), selector.Values[0]),
		)))
	}
	for _, selector := range s.notIn {
		// Add OR (key_name = 'key' AND value IN ('value1','value2')) for each notin
		ors = append(ors, qm.Or2(qm.Expr(
			qm.Where(fmt.Sprintf("%s = ?", keyColumn), selector.Key),
			andIn(valueColumn, selector.Values),
		)))
	}
	qms := subqueryBase(resourceTable, labelsTable)
	qms = append(qms, qm.Expr(ors...))
	subquery, args := models.NewSubquery(qms...).SQL()
	return qm.Where(fmt.Sprintf("NOT EXISTS (%s)", subquery), args...)
}

func subqueryBase(resourceTable, labelsTable string) []qm.QueryMod {
	return []qm.QueryMod{
		qm.Select("1"),
		qm.From(labelsTable),
		qm.Expr(qm.Where(fmt.Sprintf("%s = %s.%s", resourceGUIDColumn, models.Quote(resourceTable), guidColumn))),
	}
}

func whereIn(field string, slice []Key) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", field), values...)
}

func andIn(field string, slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", field), values...)
}
