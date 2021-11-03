// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/revision_sidecars.go -copyright_file=../../../../buildtags.txt
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

type RevisionSidecarUpserter interface {
	Upsert(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q revisionSidecarQuery) Upsert(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no revision_sidecars provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(revisionSidecarColumnsWithDefault, o)

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

	revisionSidecarUpsertCacheMut.RLock()
	cache, cached := revisionSidecarUpsertCache[key]
	revisionSidecarUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			revisionSidecarAllColumns,
			revisionSidecarColumnsWithDefault,
			revisionSidecarColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			revisionSidecarAllColumns,
			revisionSidecarPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert revision_sidecars, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(revisionSidecarPrimaryKeyColumns))
			copy(conflict, revisionSidecarPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"revision_sidecars\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert revision_sidecars")
	}

	if !cached {
		revisionSidecarUpsertCacheMut.Lock()
		revisionSidecarUpsertCache[key] = cache
		revisionSidecarUpsertCacheMut.Unlock()
	}

	return nil
}

// RevisionSidecar is an object representing the database table.
type RevisionSidecar struct {
	ID           int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string    `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name         string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Command      string    `boil:"command" json:"command" toml:"command" yaml:"command"`
	RevisionGUID string    `boil:"revision_guid" json:"revision_guid" toml:"revision_guid" yaml:"revision_guid"`
	Memory       null.Int  `boil:"memory" json:"memory,omitempty" toml:"memory" yaml:"memory,omitempty"`

	R *revisionSidecarR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L revisionSidecarL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RevisionSidecarColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	Name         string
	Command      string
	RevisionGUID string
	Memory       string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Name:         "name",
	Command:      "command",
	RevisionGUID: "revision_guid",
	Memory:       "memory",
}

var RevisionSidecarTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	Name         string
	Command      string
	RevisionGUID string
	Memory       string
}{
	ID:           "revision_sidecars.id",
	GUID:         "revision_sidecars.guid",
	CreatedAt:    "revision_sidecars.created_at",
	UpdatedAt:    "revision_sidecars.updated_at",
	Name:         "revision_sidecars.name",
	Command:      "revision_sidecars.command",
	RevisionGUID: "revision_sidecars.revision_guid",
	Memory:       "revision_sidecars.memory",
}

// Generated where

var RevisionSidecarWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	Name         whereHelperstring
	Command      whereHelperstring
	RevisionGUID whereHelperstring
	Memory       whereHelpernull_Int
}{
	ID:           whereHelperint{field: "\"revision_sidecars\".\"id\""},
	GUID:         whereHelperstring{field: "\"revision_sidecars\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"revision_sidecars\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"revision_sidecars\".\"updated_at\""},
	Name:         whereHelperstring{field: "\"revision_sidecars\".\"name\""},
	Command:      whereHelperstring{field: "\"revision_sidecars\".\"command\""},
	RevisionGUID: whereHelperstring{field: "\"revision_sidecars\".\"revision_guid\""},
	Memory:       whereHelpernull_Int{field: "\"revision_sidecars\".\"memory\""},
}

// RevisionSidecarRels is where relationship names are stored.
var RevisionSidecarRels = struct {
	Revision                    string
	RevisionSidecarProcessTypes string
}{
	Revision:                    "Revision",
	RevisionSidecarProcessTypes: "RevisionSidecarProcessTypes",
}

// revisionSidecarR is where relationships are stored.
type revisionSidecarR struct {
	Revision                    *Revision                       `boil:"Revision" json:"Revision" toml:"Revision" yaml:"Revision"`
	RevisionSidecarProcessTypes RevisionSidecarProcessTypeSlice `boil:"RevisionSidecarProcessTypes" json:"RevisionSidecarProcessTypes" toml:"RevisionSidecarProcessTypes" yaml:"RevisionSidecarProcessTypes"`
}

// NewStruct creates a new relationship struct
func (*revisionSidecarR) NewStruct() *revisionSidecarR {
	return &revisionSidecarR{}
}

// revisionSidecarL is where Load methods for each relationship are stored.
type revisionSidecarL struct{}

var (
	revisionSidecarAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "command", "revision_guid", "memory"}
	revisionSidecarColumnsWithoutDefault = []string{"guid", "updated_at", "name", "command", "revision_guid", "memory"}
	revisionSidecarColumnsWithDefault    = []string{"id", "created_at"}
	revisionSidecarPrimaryKeyColumns     = []string{"id"}
)

type (
	// RevisionSidecarSlice is an alias for a slice of pointers to RevisionSidecar.
	// This should almost always be used instead of []RevisionSidecar.
	RevisionSidecarSlice []*RevisionSidecar

	revisionSidecarQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	revisionSidecarType                 = reflect.TypeOf(&RevisionSidecar{})
	revisionSidecarMapping              = queries.MakeStructMapping(revisionSidecarType)
	revisionSidecarPrimaryKeyMapping, _ = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, revisionSidecarPrimaryKeyColumns)
	revisionSidecarInsertCacheMut       sync.RWMutex
	revisionSidecarInsertCache          = make(map[string]insertCache)
	revisionSidecarUpdateCacheMut       sync.RWMutex
	revisionSidecarUpdateCache          = make(map[string]updateCache)
	revisionSidecarUpsertCacheMut       sync.RWMutex
	revisionSidecarUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type RevisionSidecarFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*RevisionSidecar, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (RevisionSidecarSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single revisionSidecar record from the query.
func (q revisionSidecarQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RevisionSidecar, error) {
	o := &RevisionSidecar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for revision_sidecars")
	}

	return o, nil
}

// All returns all RevisionSidecar records from the query.
func (q revisionSidecarQuery) All(ctx context.Context, exec boil.ContextExecutor) (RevisionSidecarSlice, error) {
	var o []*RevisionSidecar

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RevisionSidecar slice")
	}

	return o, nil
}

// Count returns the count of all RevisionSidecar records in the query.
func (q revisionSidecarQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count revision_sidecars rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q revisionSidecarQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if revision_sidecars exists")
	}

	return count > 0, nil
}

// Revision pointed to by the foreign key.
func (q revisionSidecarQuery) Revision(o *RevisionSidecar, mods ...qm.QueryMod) revisionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.RevisionGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Revisions(queryMods...)
	queries.SetFrom(query.Query, "\"revisions\"")

	return query
}

// RevisionSidecarProcessTypes retrieves all the revision_sidecar_process_type's RevisionSidecarProcessTypes with an executor.
func (q revisionSidecarQuery) RevisionSidecarProcessTypes(o *RevisionSidecar, mods ...qm.QueryMod) revisionSidecarProcessTypeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"revision_sidecar_process_types\".\"revision_sidecar_guid\"=?", o.GUID),
	)

	query := RevisionSidecarProcessTypes(queryMods...)
	queries.SetFrom(query.Query, "\"revision_sidecar_process_types\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"revision_sidecar_process_types\".*"})
	}

	return query
}

// LoadRevision allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (revisionSidecarL) LoadRevision(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRevisionSidecar interface{}, mods queries.Applicator) error {
	var slice []*RevisionSidecar
	var object *RevisionSidecar

	if singular {
		object = maybeRevisionSidecar.(*RevisionSidecar)
	} else {
		slice = *maybeRevisionSidecar.(*[]*RevisionSidecar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &revisionSidecarR{}
		}
		args = append(args, object.RevisionGUID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &revisionSidecarR{}
			}

			for _, a := range args {
				if a == obj.RevisionGUID {
					continue Outer
				}
			}

			args = append(args, obj.RevisionGUID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`revisions`),
		qm.WhereIn(`revisions.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Revision")
	}

	var resultSlice []*Revision
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Revision")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for revisions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for revisions")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Revision = foreign
		if foreign.R == nil {
			foreign.R = &revisionR{}
		}
		foreign.R.RevisionSidecars = append(foreign.R.RevisionSidecars, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RevisionGUID == foreign.GUID {
				local.R.Revision = foreign
				if foreign.R == nil {
					foreign.R = &revisionR{}
				}
				foreign.R.RevisionSidecars = append(foreign.R.RevisionSidecars, local)
				break
			}
		}
	}

	return nil
}

// LoadRevisionSidecarProcessTypes allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (revisionSidecarL) LoadRevisionSidecarProcessTypes(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRevisionSidecar interface{}, mods queries.Applicator) error {
	var slice []*RevisionSidecar
	var object *RevisionSidecar

	if singular {
		object = maybeRevisionSidecar.(*RevisionSidecar)
	} else {
		slice = *maybeRevisionSidecar.(*[]*RevisionSidecar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &revisionSidecarR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &revisionSidecarR{}
			}

			for _, a := range args {
				if a == obj.GUID {
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
		qm.From(`revision_sidecar_process_types`),
		qm.WhereIn(`revision_sidecar_process_types.revision_sidecar_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load revision_sidecar_process_types")
	}

	var resultSlice []*RevisionSidecarProcessType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice revision_sidecar_process_types")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on revision_sidecar_process_types")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for revision_sidecar_process_types")
	}

	if singular {
		object.R.RevisionSidecarProcessTypes = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &revisionSidecarProcessTypeR{}
			}
			foreign.R.RevisionSidecar = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GUID == foreign.RevisionSidecarGUID {
				local.R.RevisionSidecarProcessTypes = append(local.R.RevisionSidecarProcessTypes, foreign)
				if foreign.R == nil {
					foreign.R = &revisionSidecarProcessTypeR{}
				}
				foreign.R.RevisionSidecar = local
				break
			}
		}
	}

	return nil
}

// SetRevision of the revisionSidecar to the related item.
// Sets o.R.Revision to related.
// Adds o to related.R.RevisionSidecars.
func (q revisionSidecarQuery) SetRevision(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Revision) error {
	var err error
	if insert {
		if err = Revisions().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"revision_sidecars\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"revision_guid"}),
		strmangle.WhereClause("\"", "\"", 2, revisionSidecarPrimaryKeyColumns),
	)
	values := []interface{}{related.GUID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.RevisionGUID = related.GUID
	if o.R == nil {
		o.R = &revisionSidecarR{
			Revision: related,
		}
	} else {
		o.R.Revision = related
	}

	if related.R == nil {
		related.R = &revisionR{
			RevisionSidecars: RevisionSidecarSlice{o},
		}
	} else {
		related.R.RevisionSidecars = append(related.R.RevisionSidecars, o)
	}

	return nil
}

// AddRevisionSidecarProcessTypes adds the given related objects to the existing relationships
// of the revision_sidecar, optionally inserting them as new records.
// Appends related to o.R.RevisionSidecarProcessTypes.
// Sets related.R.RevisionSidecar appropriately.
func (q revisionSidecarQuery) AddRevisionSidecarProcessTypes(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RevisionSidecarProcessType) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RevisionSidecarGUID = o.GUID
			if err = RevisionSidecarProcessTypes().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"revision_sidecar_process_types\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"revision_sidecar_guid"}),
				strmangle.WhereClause("\"", "\"", 2, revisionSidecarProcessTypePrimaryKeyColumns),
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

			rel.RevisionSidecarGUID = o.GUID
		}
	}

	if o.R == nil {
		o.R = &revisionSidecarR{
			RevisionSidecarProcessTypes: related,
		}
	} else {
		o.R.RevisionSidecarProcessTypes = append(o.R.RevisionSidecarProcessTypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &revisionSidecarProcessTypeR{
				RevisionSidecar: o,
			}
		} else {
			rel.R.RevisionSidecar = o
		}
	}
	return nil
}

// RevisionSidecars retrieves all the records using an executor.
func RevisionSidecars(mods ...qm.QueryMod) revisionSidecarQuery {
	mods = append(mods, qm.From("\"revision_sidecars\""))
	return revisionSidecarQuery{NewQuery(mods...)}
}

type RevisionSidecarFinder interface {
	FindRevisionSidecar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RevisionSidecar, error)
}

// FindRevisionSidecar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRevisionSidecar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RevisionSidecar, error) {
	revisionSidecarObj := &RevisionSidecar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"revision_sidecars\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, revisionSidecarObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from revision_sidecars")
	}

	return revisionSidecarObj, nil
}

type RevisionSidecarInserter interface {
	Insert(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q revisionSidecarQuery) Insert(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no revision_sidecars provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(revisionSidecarColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	revisionSidecarInsertCacheMut.RLock()
	cache, cached := revisionSidecarInsertCache[key]
	revisionSidecarInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			revisionSidecarAllColumns,
			revisionSidecarColumnsWithDefault,
			revisionSidecarColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"revision_sidecars\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"revision_sidecars\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into revision_sidecars")
	}

	if !cached {
		revisionSidecarInsertCacheMut.Lock()
		revisionSidecarInsertCache[key] = cache
		revisionSidecarInsertCacheMut.Unlock()
	}

	return nil
}

type RevisionSidecarUpdater interface {
	Update(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the RevisionSidecar.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q revisionSidecarQuery) Update(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	revisionSidecarUpdateCacheMut.RLock()
	cache, cached := revisionSidecarUpdateCache[key]
	revisionSidecarUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			revisionSidecarAllColumns,
			revisionSidecarPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update revision_sidecars, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"revision_sidecars\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, revisionSidecarPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(revisionSidecarType, revisionSidecarMapping, append(wl, revisionSidecarPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update revision_sidecars row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for revision_sidecars")
	}

	if !cached {
		revisionSidecarUpdateCacheMut.Lock()
		revisionSidecarUpdateCache[key] = cache
		revisionSidecarUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q revisionSidecarQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for revision_sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for revision_sidecars")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q revisionSidecarQuery) UpdateAllSlice(o RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"revision_sidecars\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, revisionSidecarPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in revisionSidecar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all revisionSidecar")
	}
	return rowsAff, nil
}

type RevisionSidecarDeleter interface {
	Delete(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single RevisionSidecar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q revisionSidecarQuery) Delete(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RevisionSidecar provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), revisionSidecarPrimaryKeyMapping)
	sql := "DELETE FROM \"revision_sidecars\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from revision_sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for revision_sidecars")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q revisionSidecarQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no revisionSidecarQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revision_sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revision_sidecars")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q revisionSidecarQuery) DeleteAllSlice(o RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"revision_sidecars\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revisionSidecarPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revisionSidecar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revision_sidecars")
	}

	return rowsAff, nil
}

type RevisionSidecarReloader interface {
	Reload(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q revisionSidecarQuery) Reload(o *RevisionSidecar, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRevisionSidecar(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q revisionSidecarQuery) ReloadAll(o *RevisionSidecarSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RevisionSidecarSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"revision_sidecars\".* FROM \"revision_sidecars\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revisionSidecarPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RevisionSidecarSlice")
	}

	*o = slice

	return nil
}

// RevisionSidecarExists checks if the RevisionSidecar row exists.
func RevisionSidecarExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"revision_sidecars\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if revision_sidecars exists")
	}

	return exists, nil
}
