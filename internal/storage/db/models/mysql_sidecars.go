//go:build mysql
// +build mysql

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

type SidecarUpserter interface {
	Upsert(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLSidecarUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q sidecarQuery) Upsert(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sidecars provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(sidecarColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSidecarUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	sidecarUpsertCacheMut.RLock()
	cache, cached := sidecarUpsertCache[key]
	sidecarUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			sidecarAllColumns,
			sidecarColumnsWithDefault,
			sidecarColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			sidecarAllColumns,
			sidecarPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert sidecars, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`sidecars`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `sidecars` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(sidecarType, sidecarMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(sidecarType, sidecarMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for sidecars")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == sidecarMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(sidecarType, sidecarMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for sidecars")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sidecars")
	}

CacheNoHooks:
	if !cached {
		sidecarUpsertCacheMut.Lock()
		sidecarUpsertCache[key] = cache
		sidecarUpsertCacheMut.Unlock()
	}

	return nil
}

// Sidecar is an object representing the database table.
type Sidecar struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID      string    `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Command   string    `boil:"command" json:"command" toml:"command" yaml:"command"`
	AppGUID   string    `boil:"app_guid" json:"app_guid" toml:"app_guid" yaml:"app_guid"`
	Memory    null.Int  `boil:"memory" json:"memory,omitempty" toml:"memory" yaml:"memory,omitempty"`
	Origin    string    `boil:"origin" json:"origin" toml:"origin" yaml:"origin"`

	R *sidecarR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L sidecarL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SidecarColumns = struct {
	ID        string
	GUID      string
	CreatedAt string
	UpdatedAt string
	Name      string
	Command   string
	AppGUID   string
	Memory    string
	Origin    string
}{
	ID:        "id",
	GUID:      "guid",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Name:      "name",
	Command:   "command",
	AppGUID:   "app_guid",
	Memory:    "memory",
	Origin:    "origin",
}

var SidecarTableColumns = struct {
	ID        string
	GUID      string
	CreatedAt string
	UpdatedAt string
	Name      string
	Command   string
	AppGUID   string
	Memory    string
	Origin    string
}{
	ID:        "sidecars.id",
	GUID:      "sidecars.guid",
	CreatedAt: "sidecars.created_at",
	UpdatedAt: "sidecars.updated_at",
	Name:      "sidecars.name",
	Command:   "sidecars.command",
	AppGUID:   "sidecars.app_guid",
	Memory:    "sidecars.memory",
	Origin:    "sidecars.origin",
}

// Generated where

var SidecarWhere = struct {
	ID        whereHelperint
	GUID      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
	Name      whereHelperstring
	Command   whereHelperstring
	AppGUID   whereHelperstring
	Memory    whereHelpernull_Int
	Origin    whereHelperstring
}{
	ID:        whereHelperint{field: "`sidecars`.`id`"},
	GUID:      whereHelperstring{field: "`sidecars`.`guid`"},
	CreatedAt: whereHelpertime_Time{field: "`sidecars`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`sidecars`.`updated_at`"},
	Name:      whereHelperstring{field: "`sidecars`.`name`"},
	Command:   whereHelperstring{field: "`sidecars`.`command`"},
	AppGUID:   whereHelperstring{field: "`sidecars`.`app_guid`"},
	Memory:    whereHelpernull_Int{field: "`sidecars`.`memory`"},
	Origin:    whereHelperstring{field: "`sidecars`.`origin`"},
}

// SidecarRels is where relationship names are stored.
var SidecarRels = struct {
	App                 string
	SidecarProcessTypes string
}{
	App:                 "App",
	SidecarProcessTypes: "SidecarProcessTypes",
}

// sidecarR is where relationships are stored.
type sidecarR struct {
	App                 *App                    `boil:"App" json:"App" toml:"App" yaml:"App"`
	SidecarProcessTypes SidecarProcessTypeSlice `boil:"SidecarProcessTypes" json:"SidecarProcessTypes" toml:"SidecarProcessTypes" yaml:"SidecarProcessTypes"`
}

// NewStruct creates a new relationship struct
func (*sidecarR) NewStruct() *sidecarR {
	return &sidecarR{}
}

// sidecarL is where Load methods for each relationship are stored.
type sidecarL struct{}

var (
	sidecarAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "command", "app_guid", "memory", "origin"}
	sidecarColumnsWithoutDefault = []string{"guid", "updated_at", "name", "command", "app_guid", "memory"}
	sidecarColumnsWithDefault    = []string{"id", "created_at", "origin"}
	sidecarPrimaryKeyColumns     = []string{"id"}
)

type (
	// SidecarSlice is an alias for a slice of pointers to Sidecar.
	// This should almost always be used instead of []Sidecar.
	SidecarSlice []*Sidecar

	sidecarQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	sidecarType                 = reflect.TypeOf(&Sidecar{})
	sidecarMapping              = queries.MakeStructMapping(sidecarType)
	sidecarPrimaryKeyMapping, _ = queries.BindMapping(sidecarType, sidecarMapping, sidecarPrimaryKeyColumns)
	sidecarInsertCacheMut       sync.RWMutex
	sidecarInsertCache          = make(map[string]insertCache)
	sidecarUpdateCacheMut       sync.RWMutex
	sidecarUpdateCache          = make(map[string]updateCache)
	sidecarUpsertCacheMut       sync.RWMutex
	sidecarUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type SidecarFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*Sidecar, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (SidecarSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single sidecar record from the query.
func (q sidecarQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Sidecar, error) {
	o := &Sidecar{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for sidecars")
	}

	return o, nil
}

// All returns all Sidecar records from the query.
func (q sidecarQuery) All(ctx context.Context, exec boil.ContextExecutor) (SidecarSlice, error) {
	var o []*Sidecar

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Sidecar slice")
	}

	return o, nil
}

// Count returns the count of all Sidecar records in the query.
func (q sidecarQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count sidecars rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q sidecarQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if sidecars exists")
	}

	return count > 0, nil
}

// App pointed to by the foreign key.
func (q sidecarQuery) App(o *Sidecar, mods ...qm.QueryMod) appQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.AppGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Apps(queryMods...)
	queries.SetFrom(query.Query, "`apps`")

	return query
}

// SidecarProcessTypes retrieves all the sidecar_process_type's SidecarProcessTypes with an executor.
func (q sidecarQuery) SidecarProcessTypes(o *Sidecar, mods ...qm.QueryMod) sidecarProcessTypeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`sidecar_process_types`.`sidecar_guid`=?", o.GUID),
	)

	query := SidecarProcessTypes(queryMods...)
	queries.SetFrom(query.Query, "`sidecar_process_types`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`sidecar_process_types`.*"})
	}

	return query
}

// LoadApp allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (sidecarL) LoadApp(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSidecar interface{}, mods queries.Applicator) error {
	var slice []*Sidecar
	var object *Sidecar

	if singular {
		object = maybeSidecar.(*Sidecar)
	} else {
		slice = *maybeSidecar.(*[]*Sidecar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &sidecarR{}
		}
		args = append(args, object.AppGUID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &sidecarR{}
			}

			for _, a := range args {
				if a == obj.AppGUID {
					continue Outer
				}
			}

			args = append(args, obj.AppGUID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`apps`),
		qm.WhereIn(`apps.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load App")
	}

	var resultSlice []*App
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice App")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for apps")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for apps")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.App = foreign
		if foreign.R == nil {
			foreign.R = &appR{}
		}
		foreign.R.Sidecars = append(foreign.R.Sidecars, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AppGUID == foreign.GUID {
				local.R.App = foreign
				if foreign.R == nil {
					foreign.R = &appR{}
				}
				foreign.R.Sidecars = append(foreign.R.Sidecars, local)
				break
			}
		}
	}

	return nil
}

// LoadSidecarProcessTypes allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (sidecarL) LoadSidecarProcessTypes(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSidecar interface{}, mods queries.Applicator) error {
	var slice []*Sidecar
	var object *Sidecar

	if singular {
		object = maybeSidecar.(*Sidecar)
	} else {
		slice = *maybeSidecar.(*[]*Sidecar)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &sidecarR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &sidecarR{}
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
		qm.From(`sidecar_process_types`),
		qm.WhereIn(`sidecar_process_types.sidecar_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load sidecar_process_types")
	}

	var resultSlice []*SidecarProcessType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice sidecar_process_types")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on sidecar_process_types")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sidecar_process_types")
	}

	if singular {
		object.R.SidecarProcessTypes = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &sidecarProcessTypeR{}
			}
			foreign.R.Sidecar = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GUID == foreign.SidecarGUID {
				local.R.SidecarProcessTypes = append(local.R.SidecarProcessTypes, foreign)
				if foreign.R == nil {
					foreign.R = &sidecarProcessTypeR{}
				}
				foreign.R.Sidecar = local
				break
			}
		}
	}

	return nil
}

// SetApp of the sidecar to the related item.
// Sets o.R.App to related.
// Adds o to related.R.Sidecars.
func (q sidecarQuery) SetApp(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, insert bool, related *App) error {
	var err error
	if insert {
		if err = Apps().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `sidecars` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"app_guid"}),
		strmangle.WhereClause("`", "`", 0, sidecarPrimaryKeyColumns),
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

	o.AppGUID = related.GUID
	if o.R == nil {
		o.R = &sidecarR{
			App: related,
		}
	} else {
		o.R.App = related
	}

	if related.R == nil {
		related.R = &appR{
			Sidecars: SidecarSlice{o},
		}
	} else {
		related.R.Sidecars = append(related.R.Sidecars, o)
	}

	return nil
}

// AddSidecarProcessTypes adds the given related objects to the existing relationships
// of the sidecar, optionally inserting them as new records.
// Appends related to o.R.SidecarProcessTypes.
// Sets related.R.Sidecar appropriately.
func (q sidecarQuery) AddSidecarProcessTypes(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SidecarProcessType) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SidecarGUID = o.GUID
			if err = SidecarProcessTypes().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `sidecar_process_types` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"sidecar_guid"}),
				strmangle.WhereClause("`", "`", 0, sidecarProcessTypePrimaryKeyColumns),
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

			rel.SidecarGUID = o.GUID
		}
	}

	if o.R == nil {
		o.R = &sidecarR{
			SidecarProcessTypes: related,
		}
	} else {
		o.R.SidecarProcessTypes = append(o.R.SidecarProcessTypes, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &sidecarProcessTypeR{
				Sidecar: o,
			}
		} else {
			rel.R.Sidecar = o
		}
	}
	return nil
}

// Sidecars retrieves all the records using an executor.
func Sidecars(mods ...qm.QueryMod) sidecarQuery {
	mods = append(mods, qm.From("`sidecars`"))
	return sidecarQuery{NewQuery(mods...)}
}

type SidecarFinder interface {
	FindSidecar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Sidecar, error)
}

// FindSidecar retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSidecar(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Sidecar, error) {
	sidecarObj := &Sidecar{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `sidecars` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, sidecarObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from sidecars")
	}

	return sidecarObj, nil
}

type SidecarInserter interface {
	Insert(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q sidecarQuery) Insert(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no sidecars provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(sidecarColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	sidecarInsertCacheMut.RLock()
	cache, cached := sidecarInsertCache[key]
	sidecarInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			sidecarAllColumns,
			sidecarColumnsWithDefault,
			sidecarColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(sidecarType, sidecarMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(sidecarType, sidecarMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `sidecars` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `sidecars` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `sidecars` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, sidecarPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into sidecars")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == sidecarMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for sidecars")
	}

CacheNoHooks:
	if !cached {
		sidecarInsertCacheMut.Lock()
		sidecarInsertCache[key] = cache
		sidecarInsertCacheMut.Unlock()
	}

	return nil
}

type SidecarUpdater interface {
	Update(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o SidecarSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the Sidecar.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q sidecarQuery) Update(o *Sidecar, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	sidecarUpdateCacheMut.RLock()
	cache, cached := sidecarUpdateCache[key]
	sidecarUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			sidecarAllColumns,
			sidecarPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update sidecars, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `sidecars` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, sidecarPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(sidecarType, sidecarMapping, append(wl, sidecarPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update sidecars row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for sidecars")
	}

	if !cached {
		sidecarUpdateCacheMut.Lock()
		sidecarUpdateCache[key] = cache
		sidecarUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q sidecarQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for sidecars")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q sidecarQuery) UpdateAllSlice(o SidecarSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `sidecars` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sidecarPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in sidecar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all sidecar")
	}
	return rowsAff, nil
}

type SidecarDeleter interface {
	Delete(o *Sidecar, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o SidecarSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single Sidecar record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q sidecarQuery) Delete(o *Sidecar, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Sidecar provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), sidecarPrimaryKeyMapping)
	sql := "DELETE FROM `sidecars` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for sidecars")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q sidecarQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no sidecarQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sidecars")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sidecars")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q sidecarQuery) DeleteAllSlice(o SidecarSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `sidecars` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sidecarPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from sidecar slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for sidecars")
	}

	return rowsAff, nil
}

type SidecarReloader interface {
	Reload(o *Sidecar, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *SidecarSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q sidecarQuery) Reload(o *Sidecar, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSidecar(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q sidecarQuery) ReloadAll(o *SidecarSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SidecarSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), sidecarPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `sidecars`.* FROM `sidecars` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, sidecarPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SidecarSlice")
	}

	*o = slice

	return nil
}

// SidecarExists checks if the Sidecar row exists.
func SidecarExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `sidecars` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if sidecars exists")
	}

	return exists, nil
}
