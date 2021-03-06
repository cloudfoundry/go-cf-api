//go:build psql
// +build psql

//go:generate sh -c "echo '\x2bbuild unit' > ../../../../buildtags.txt && mockgen -source=$GOFILE -destination=mocks/stacks.go -copyright_file=../../../../buildtags.txt && rm -f ../../../../buildtags.txt"
// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

type StackUpserter interface {
	Upsert(o *Stack, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q stackQuery) Upsert(o *Stack, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stacks provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(stackColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stackUpsertCacheMut.RLock()
	cache, cached := stackUpsertCache[key]
	stackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stackAllColumns,
			stackColumnsWithDefault,
			stackColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stackAllColumns,
			stackPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stacks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stackPrimaryKeyColumns))
			copy(conflict, stackPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"stacks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stackType, stackMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert stacks")
	}

	if !cached {
		stackUpsertCacheMut.Lock()
		stackUpsertCache[key] = cache
		stackUpsertCacheMut.Unlock()
	}

	return nil
}

// Stack is an object representing the database table.
type Stack struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID        string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name        string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`

	R *stackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StackColumns = struct {
	ID          string
	GUID        string
	CreatedAt   string
	UpdatedAt   string
	Name        string
	Description string
}{
	ID:          "id",
	GUID:        "guid",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	Name:        "name",
	Description: "description",
}

var StackTableColumns = struct {
	ID          string
	GUID        string
	CreatedAt   string
	UpdatedAt   string
	Name        string
	Description string
}{
	ID:          "stacks.id",
	GUID:        "stacks.guid",
	CreatedAt:   "stacks.created_at",
	UpdatedAt:   "stacks.updated_at",
	Name:        "stacks.name",
	Description: "stacks.description",
}

// Generated where

var StackWhere = struct {
	ID          whereHelperint
	GUID        whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpernull_Time
	Name        whereHelperstring
	Description whereHelpernull_String
}{
	ID:          whereHelperint{field: "\"stacks\".\"id\""},
	GUID:        whereHelperstring{field: "\"stacks\".\"guid\""},
	CreatedAt:   whereHelpertime_Time{field: "\"stacks\".\"created_at\""},
	UpdatedAt:   whereHelpernull_Time{field: "\"stacks\".\"updated_at\""},
	Name:        whereHelperstring{field: "\"stacks\".\"name\""},
	Description: whereHelpernull_String{field: "\"stacks\".\"description\""},
}

// StackRels is where relationship names are stored.
var StackRels = struct {
	ResourceStackAnnotations string
	ResourceStackLabels      string
}{
	ResourceStackAnnotations: "ResourceStackAnnotations",
	ResourceStackLabels:      "ResourceStackLabels",
}

// stackR is where relationships are stored.
type stackR struct {
	ResourceStackAnnotations StackAnnotationSlice `boil:"ResourceStackAnnotations" json:"ResourceStackAnnotations" toml:"ResourceStackAnnotations" yaml:"ResourceStackAnnotations"`
	ResourceStackLabels      StackLabelSlice      `boil:"ResourceStackLabels" json:"ResourceStackLabels" toml:"ResourceStackLabels" yaml:"ResourceStackLabels"`
}

// NewStruct creates a new relationship struct
func (*stackR) NewStruct() *stackR {
	return &stackR{}
}

// stackL is where Load methods for each relationship are stored.
type stackL struct{}

var (
	stackAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "description"}
	stackColumnsWithoutDefault = []string{"guid", "updated_at", "name", "description"}
	stackColumnsWithDefault    = []string{"id", "created_at"}
	stackPrimaryKeyColumns     = []string{"id"}
)

type (
	// StackSlice is an alias for a slice of pointers to Stack.
	// This should almost always be used instead of []Stack.
	StackSlice []*Stack

	stackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stackType                 = reflect.TypeOf(&Stack{})
	stackMapping              = queries.MakeStructMapping(stackType)
	stackPrimaryKeyMapping, _ = queries.BindMapping(stackType, stackMapping, stackPrimaryKeyColumns)
	stackInsertCacheMut       sync.RWMutex
	stackInsertCache          = make(map[string]insertCache)
	stackUpdateCacheMut       sync.RWMutex
	stackUpdateCache          = make(map[string]updateCache)
	stackUpsertCacheMut       sync.RWMutex
	stackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type StackFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*Stack, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (StackSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single stack record from the query.
func (q stackQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stack, error) {
	o := &Stack{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stacks")
	}

	return o, nil
}

// All returns all Stack records from the query.
func (q stackQuery) All(ctx context.Context, exec boil.ContextExecutor) (StackSlice, error) {
	var o []*Stack

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stack slice")
	}

	return o, nil
}

// Count returns the count of all Stack records in the query.
func (q stackQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stacks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stackQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stacks exists")
	}

	return count > 0, nil
}

// ResourceStackAnnotations retrieves all the stack_annotation's StackAnnotations with an executor via resource_guid column.
func (q stackQuery) ResourceStackAnnotations(o *Stack, mods ...qm.QueryMod) stackAnnotationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"stack_annotations\".\"resource_guid\"=?", o.GUID),
	)

	query := StackAnnotations(queryMods...)
	queries.SetFrom(query.Query, "\"stack_annotations\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"stack_annotations\".*"})
	}

	return query
}

// ResourceStackLabels retrieves all the stack_label's StackLabels with an executor via resource_guid column.
func (q stackQuery) ResourceStackLabels(o *Stack, mods ...qm.QueryMod) stackLabelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"stack_labels\".\"resource_guid\"=?", o.GUID),
	)

	query := StackLabels(queryMods...)
	queries.SetFrom(query.Query, "\"stack_labels\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"stack_labels\".*"})
	}

	return query
}

// LoadResourceStackAnnotations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stackL) LoadResourceStackAnnotations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStack interface{}, mods queries.Applicator) error {
	var slice []*Stack
	var object *Stack

	if singular {
		object = maybeStack.(*Stack)
	} else {
		slice = *maybeStack.(*[]*Stack)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stackR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stackR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.GUID) {
					continue Outer
				}
			}

			args = append(args, obj.GUID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`stack_annotations`),
		qm.WhereIn(`stack_annotations.resource_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stack_annotations")
	}

	var resultSlice []*StackAnnotation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stack_annotations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on stack_annotations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stack_annotations")
	}

	if singular {
		object.R.ResourceStackAnnotations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &stackAnnotationR{}
			}
			foreign.R.Resource = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.GUID, foreign.ResourceGUID) {
				local.R.ResourceStackAnnotations = append(local.R.ResourceStackAnnotations, foreign)
				if foreign.R == nil {
					foreign.R = &stackAnnotationR{}
				}
				foreign.R.Resource = local
				break
			}
		}
	}

	return nil
}

// LoadResourceStackLabels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (stackL) LoadResourceStackLabels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStack interface{}, mods queries.Applicator) error {
	var slice []*Stack
	var object *Stack

	if singular {
		object = maybeStack.(*Stack)
	} else {
		slice = *maybeStack.(*[]*Stack)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stackR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stackR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.GUID) {
					continue Outer
				}
			}

			args = append(args, obj.GUID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`stack_labels`),
		qm.WhereIn(`stack_labels.resource_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stack_labels")
	}

	var resultSlice []*StackLabel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stack_labels")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on stack_labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stack_labels")
	}

	if singular {
		object.R.ResourceStackLabels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &stackLabelR{}
			}
			foreign.R.Resource = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.GUID, foreign.ResourceGUID) {
				local.R.ResourceStackLabels = append(local.R.ResourceStackLabels, foreign)
				if foreign.R == nil {
					foreign.R = &stackLabelR{}
				}
				foreign.R.Resource = local
				break
			}
		}
	}

	return nil
}

// AddResourceStackAnnotations adds the given related objects to the existing relationships
// of the stack, optionally inserting them as new records.
// Appends related to o.R.ResourceStackAnnotations.
// Sets related.R.Resource appropriately.
func (q stackQuery) AddResourceStackAnnotations(o *Stack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StackAnnotation) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ResourceGUID, o.GUID)
			if err = StackAnnotations().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"stack_annotations\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
				strmangle.WhereClause("\"", "\"", 2, stackAnnotationPrimaryKeyColumns),
			)
			values := []interface{}{o.GUID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ResourceGUID, o.GUID)
		}
	}

	if o.R == nil {
		o.R = &stackR{
			ResourceStackAnnotations: related,
		}
	} else {
		o.R.ResourceStackAnnotations = append(o.R.ResourceStackAnnotations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stackAnnotationR{
				Resource: o,
			}
		} else {
			rel.R.Resource = o
		}
	}
	return nil
}

// SetResourceStackAnnotations removes all previously related items of the
// stack replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Resource's ResourceStackAnnotations accordingly.
// Replaces o.R.ResourceStackAnnotations with related.
// Sets related.R.Resource's ResourceStackAnnotations accordingly.
func (q stackQuery) SetResourceStackAnnotations(o *Stack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StackAnnotation) error {
	query := "update \"stack_annotations\" set \"resource_guid\" = null where \"resource_guid\" = $1"
	values := []interface{}{o.GUID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ResourceStackAnnotations {
			queries.SetScanner(&rel.ResourceGUID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Resource = nil
		}

		o.R.ResourceStackAnnotations = nil
	}
	return q.AddResourceStackAnnotations(o, ctx, exec, insert, related...)
}

// RemoveResourceStackAnnotations relationships from objects passed in.
// Removes related items from R.ResourceStackAnnotations (uses pointer comparison, removal does not keep order)
// Sets related.R.Resource.
func (q stackQuery) RemoveResourceStackAnnotations(o *Stack, ctx context.Context, exec boil.ContextExecutor, related ...*StackAnnotation) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ResourceGUID, nil)
		if rel.R != nil {
			rel.R.Resource = nil
		}
		if _, err = StackAnnotations().Update(rel, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ResourceStackAnnotations {
			if rel != ri {
				continue
			}

			ln := len(o.R.ResourceStackAnnotations)
			if ln > 1 && i < ln-1 {
				o.R.ResourceStackAnnotations[i] = o.R.ResourceStackAnnotations[ln-1]
			}
			o.R.ResourceStackAnnotations = o.R.ResourceStackAnnotations[:ln-1]
			break
		}
	}

	return nil
}

// AddResourceStackLabels adds the given related objects to the existing relationships
// of the stack, optionally inserting them as new records.
// Appends related to o.R.ResourceStackLabels.
// Sets related.R.Resource appropriately.
func (q stackQuery) AddResourceStackLabels(o *Stack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StackLabel) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ResourceGUID, o.GUID)
			if err = StackLabels().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"stack_labels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
				strmangle.WhereClause("\"", "\"", 2, stackLabelPrimaryKeyColumns),
			)
			values := []interface{}{o.GUID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ResourceGUID, o.GUID)
		}
	}

	if o.R == nil {
		o.R = &stackR{
			ResourceStackLabels: related,
		}
	} else {
		o.R.ResourceStackLabels = append(o.R.ResourceStackLabels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stackLabelR{
				Resource: o,
			}
		} else {
			rel.R.Resource = o
		}
	}
	return nil
}

// SetResourceStackLabels removes all previously related items of the
// stack replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Resource's ResourceStackLabels accordingly.
// Replaces o.R.ResourceStackLabels with related.
// Sets related.R.Resource's ResourceStackLabels accordingly.
func (q stackQuery) SetResourceStackLabels(o *Stack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*StackLabel) error {
	query := "update \"stack_labels\" set \"resource_guid\" = null where \"resource_guid\" = $1"
	values := []interface{}{o.GUID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ResourceStackLabels {
			queries.SetScanner(&rel.ResourceGUID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Resource = nil
		}

		o.R.ResourceStackLabels = nil
	}
	return q.AddResourceStackLabels(o, ctx, exec, insert, related...)
}

// RemoveResourceStackLabels relationships from objects passed in.
// Removes related items from R.ResourceStackLabels (uses pointer comparison, removal does not keep order)
// Sets related.R.Resource.
func (q stackQuery) RemoveResourceStackLabels(o *Stack, ctx context.Context, exec boil.ContextExecutor, related ...*StackLabel) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ResourceGUID, nil)
		if rel.R != nil {
			rel.R.Resource = nil
		}
		if _, err = StackLabels().Update(rel, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ResourceStackLabels {
			if rel != ri {
				continue
			}

			ln := len(o.R.ResourceStackLabels)
			if ln > 1 && i < ln-1 {
				o.R.ResourceStackLabels[i] = o.R.ResourceStackLabels[ln-1]
			}
			o.R.ResourceStackLabels = o.R.ResourceStackLabels[:ln-1]
			break
		}
	}

	return nil
}

// Stacks retrieves all the records using an executor.
func Stacks(mods ...qm.QueryMod) stackQuery {
	mods = append(mods, qm.From("\"stacks\""))
	return stackQuery{NewQuery(mods...)}
}

type StackFinder interface {
	FindStack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Stack, error)
}

// FindStack retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Stack, error) {
	stackObj := &Stack{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stacks\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stacks")
	}

	return stackObj, nil
}

type StackInserter interface {
	Insert(o *Stack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q stackQuery) Insert(o *Stack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stacks provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(stackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stackInsertCacheMut.RLock()
	cache, cached := stackInsertCache[key]
	stackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stackAllColumns,
			stackColumnsWithDefault,
			stackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stackType, stackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"stacks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"stacks\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stacks")
	}

	if !cached {
		stackInsertCacheMut.Lock()
		stackInsertCache[key] = cache
		stackInsertCacheMut.Unlock()
	}

	return nil
}

type StackUpdater interface {
	Update(o *Stack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o StackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the Stack.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q stackQuery) Update(o *Stack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	stackUpdateCacheMut.RLock()
	cache, cached := stackUpdateCache[key]
	stackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stackAllColumns,
			stackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stacks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stacks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stackType, stackMapping, append(wl, stackPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stacks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stacks")
	}

	if !cached {
		stackUpdateCacheMut.Lock()
		stackUpdateCache[key] = cache
		stackUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q stackQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stacks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q stackQuery) UpdateAllSlice(o StackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"stacks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stackPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stack")
	}
	return rowsAff, nil
}

type StackDeleter interface {
	Delete(o *Stack, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o StackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single Stack record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q stackQuery) Delete(o *Stack, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stack provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stackPrimaryKeyMapping)
	sql := "DELETE FROM \"stacks\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stackQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q stackQuery) DeleteAllSlice(o StackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"stacks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stackPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stacks")
	}

	return rowsAff, nil
}

type StackReloader interface {
	Reload(o *Stack, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *StackSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q stackQuery) Reload(o *Stack, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStack(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q stackQuery) ReloadAll(o *StackSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"stacks\".* FROM \"stacks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stackPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StackSlice")
	}

	*o = slice

	return nil
}

// StackExists checks if the Stack row exists.
func StackExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"stacks\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stacks exists")
	}

	return exists, nil
}
