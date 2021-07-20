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

// RevisionSidecarProcessType is an object representing the database table.
type RevisionSidecarProcessType struct {
	ID                  int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                string    `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt           time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt           null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Type                string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	RevisionSidecarGUID string    `boil:"revision_sidecar_guid" json:"revision_sidecar_guid" toml:"revision_sidecar_guid" yaml:"revision_sidecar_guid"`

	R *revisionSidecarProcessTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L revisionSidecarProcessTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RevisionSidecarProcessTypeColumns = struct {
	ID                  string
	GUID                string
	CreatedAt           string
	UpdatedAt           string
	Type                string
	RevisionSidecarGUID string
}{
	ID:                  "id",
	GUID:                "guid",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	Type:                "type",
	RevisionSidecarGUID: "revision_sidecar_guid",
}

var RevisionSidecarProcessTypeTableColumns = struct {
	ID                  string
	GUID                string
	CreatedAt           string
	UpdatedAt           string
	Type                string
	RevisionSidecarGUID string
}{
	ID:                  "revision_sidecar_process_types.id",
	GUID:                "revision_sidecar_process_types.guid",
	CreatedAt:           "revision_sidecar_process_types.created_at",
	UpdatedAt:           "revision_sidecar_process_types.updated_at",
	Type:                "revision_sidecar_process_types.type",
	RevisionSidecarGUID: "revision_sidecar_process_types.revision_sidecar_guid",
}

// Generated where

var RevisionSidecarProcessTypeWhere = struct {
	ID                  whereHelperint
	GUID                whereHelperstring
	CreatedAt           whereHelpertime_Time
	UpdatedAt           whereHelpernull_Time
	Type                whereHelperstring
	RevisionSidecarGUID whereHelperstring
}{
	ID:                  whereHelperint{field: "\"revision_sidecar_process_types\".\"id\""},
	GUID:                whereHelperstring{field: "\"revision_sidecar_process_types\".\"guid\""},
	CreatedAt:           whereHelpertime_Time{field: "\"revision_sidecar_process_types\".\"created_at\""},
	UpdatedAt:           whereHelpernull_Time{field: "\"revision_sidecar_process_types\".\"updated_at\""},
	Type:                whereHelperstring{field: "\"revision_sidecar_process_types\".\"type\""},
	RevisionSidecarGUID: whereHelperstring{field: "\"revision_sidecar_process_types\".\"revision_sidecar_guid\""},
}

// RevisionSidecarProcessTypeRels is where relationship names are stored.
var RevisionSidecarProcessTypeRels = struct {
	RevisionSidecar string
}{
	RevisionSidecar: "RevisionSidecar",
}

// revisionSidecarProcessTypeR is where relationships are stored.
type revisionSidecarProcessTypeR struct {
	RevisionSidecar *RevisionSidecar `boil:"RevisionSidecar" json:"RevisionSidecar" toml:"RevisionSidecar" yaml:"RevisionSidecar"`
}

// NewStruct creates a new relationship struct
func (*revisionSidecarProcessTypeR) NewStruct() *revisionSidecarProcessTypeR {
	return &revisionSidecarProcessTypeR{}
}

// revisionSidecarProcessTypeL is where Load methods for each relationship are stored.
type revisionSidecarProcessTypeL struct{}

var (
	revisionSidecarProcessTypeAllColumns            = []string{"id", "guid", "created_at", "updated_at", "type", "revision_sidecar_guid"}
	revisionSidecarProcessTypeColumnsWithoutDefault = []string{"guid", "updated_at", "type", "revision_sidecar_guid"}
	revisionSidecarProcessTypeColumnsWithDefault    = []string{"id", "created_at"}
	revisionSidecarProcessTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// RevisionSidecarProcessTypeSlice is an alias for a slice of pointers to RevisionSidecarProcessType.
	// This should almost always be used instead of []RevisionSidecarProcessType.
	RevisionSidecarProcessTypeSlice []*RevisionSidecarProcessType
	// RevisionSidecarProcessTypeHook is the signature for custom RevisionSidecarProcessType hook methods
	RevisionSidecarProcessTypeHook func(context.Context, boil.ContextExecutor, *RevisionSidecarProcessType) error

	revisionSidecarProcessTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	revisionSidecarProcessTypeType                 = reflect.TypeOf(&RevisionSidecarProcessType{})
	revisionSidecarProcessTypeMapping              = queries.MakeStructMapping(revisionSidecarProcessTypeType)
	revisionSidecarProcessTypePrimaryKeyMapping, _ = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, revisionSidecarProcessTypePrimaryKeyColumns)
	revisionSidecarProcessTypeInsertCacheMut       sync.RWMutex
	revisionSidecarProcessTypeInsertCache          = make(map[string]insertCache)
	revisionSidecarProcessTypeUpdateCacheMut       sync.RWMutex
	revisionSidecarProcessTypeUpdateCache          = make(map[string]updateCache)
	revisionSidecarProcessTypeUpsertCacheMut       sync.RWMutex
	revisionSidecarProcessTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var revisionSidecarProcessTypeBeforeInsertHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeBeforeUpdateHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeBeforeDeleteHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeBeforeUpsertHooks []RevisionSidecarProcessTypeHook

var revisionSidecarProcessTypeAfterInsertHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeAfterSelectHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeAfterUpdateHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeAfterDeleteHooks []RevisionSidecarProcessTypeHook
var revisionSidecarProcessTypeAfterUpsertHooks []RevisionSidecarProcessTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RevisionSidecarProcessType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RevisionSidecarProcessType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RevisionSidecarProcessType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RevisionSidecarProcessType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RevisionSidecarProcessType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RevisionSidecarProcessType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RevisionSidecarProcessType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RevisionSidecarProcessType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RevisionSidecarProcessType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range revisionSidecarProcessTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRevisionSidecarProcessTypeHook registers your hook function for all future operations.
func AddRevisionSidecarProcessTypeHook(hookPoint boil.HookPoint, revisionSidecarProcessTypeHook RevisionSidecarProcessTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		revisionSidecarProcessTypeBeforeInsertHooks = append(revisionSidecarProcessTypeBeforeInsertHooks, revisionSidecarProcessTypeHook)
	case boil.BeforeUpdateHook:
		revisionSidecarProcessTypeBeforeUpdateHooks = append(revisionSidecarProcessTypeBeforeUpdateHooks, revisionSidecarProcessTypeHook)
	case boil.BeforeDeleteHook:
		revisionSidecarProcessTypeBeforeDeleteHooks = append(revisionSidecarProcessTypeBeforeDeleteHooks, revisionSidecarProcessTypeHook)
	case boil.BeforeUpsertHook:
		revisionSidecarProcessTypeBeforeUpsertHooks = append(revisionSidecarProcessTypeBeforeUpsertHooks, revisionSidecarProcessTypeHook)
	case boil.AfterInsertHook:
		revisionSidecarProcessTypeAfterInsertHooks = append(revisionSidecarProcessTypeAfterInsertHooks, revisionSidecarProcessTypeHook)
	case boil.AfterSelectHook:
		revisionSidecarProcessTypeAfterSelectHooks = append(revisionSidecarProcessTypeAfterSelectHooks, revisionSidecarProcessTypeHook)
	case boil.AfterUpdateHook:
		revisionSidecarProcessTypeAfterUpdateHooks = append(revisionSidecarProcessTypeAfterUpdateHooks, revisionSidecarProcessTypeHook)
	case boil.AfterDeleteHook:
		revisionSidecarProcessTypeAfterDeleteHooks = append(revisionSidecarProcessTypeAfterDeleteHooks, revisionSidecarProcessTypeHook)
	case boil.AfterUpsertHook:
		revisionSidecarProcessTypeAfterUpsertHooks = append(revisionSidecarProcessTypeAfterUpsertHooks, revisionSidecarProcessTypeHook)
	}
}

// One returns a single revisionSidecarProcessType record from the query.
func (q revisionSidecarProcessTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RevisionSidecarProcessType, error) {
	o := &RevisionSidecarProcessType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for revision_sidecar_process_types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RevisionSidecarProcessType records from the query.
func (q revisionSidecarProcessTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (RevisionSidecarProcessTypeSlice, error) {
	var o []*RevisionSidecarProcessType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RevisionSidecarProcessType slice")
	}

	if len(revisionSidecarProcessTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RevisionSidecarProcessType records in the query.
func (q revisionSidecarProcessTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count revision_sidecar_process_types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q revisionSidecarProcessTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if revision_sidecar_process_types exists")
	}

	return count > 0, nil
}

// RevisionSidecar pointed to by the foreign key.
func (o *RevisionSidecarProcessType) RevisionSidecar(mods ...qm.QueryMod) revisionSidecarQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.RevisionSidecarGUID),
	}

	queryMods = append(queryMods, mods...)

	query := RevisionSidecars(queryMods...)
	queries.SetFrom(query.Query, "\"revision_sidecars\"")

	return query
}

// LoadRevisionSidecar allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (revisionSidecarProcessTypeL) LoadRevisionSidecar(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRevisionSidecarProcessType interface{}, mods queries.Applicator) error {
	var slice []*RevisionSidecarProcessType
	var object *RevisionSidecarProcessType

	if singular {
		object = maybeRevisionSidecarProcessType.(*RevisionSidecarProcessType)
	} else {
		slice = *maybeRevisionSidecarProcessType.(*[]*RevisionSidecarProcessType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &revisionSidecarProcessTypeR{}
		}
		args = append(args, object.RevisionSidecarGUID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &revisionSidecarProcessTypeR{}
			}

			for _, a := range args {
				if a == obj.RevisionSidecarGUID {
					continue Outer
				}
			}

			args = append(args, obj.RevisionSidecarGUID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`revision_sidecars`),
		qm.WhereIn(`revision_sidecars.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RevisionSidecar")
	}

	var resultSlice []*RevisionSidecar
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RevisionSidecar")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for revision_sidecars")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for revision_sidecars")
	}

	if len(revisionSidecarProcessTypeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.RevisionSidecar = foreign
		if foreign.R == nil {
			foreign.R = &revisionSidecarR{}
		}
		foreign.R.RevisionSidecarProcessTypes = append(foreign.R.RevisionSidecarProcessTypes, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RevisionSidecarGUID == foreign.GUID {
				local.R.RevisionSidecar = foreign
				if foreign.R == nil {
					foreign.R = &revisionSidecarR{}
				}
				foreign.R.RevisionSidecarProcessTypes = append(foreign.R.RevisionSidecarProcessTypes, local)
				break
			}
		}
	}

	return nil
}

// SetRevisionSidecar of the revisionSidecarProcessType to the related item.
// Sets o.R.RevisionSidecar to related.
// Adds o to related.R.RevisionSidecarProcessTypes.
func (o *RevisionSidecarProcessType) SetRevisionSidecar(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RevisionSidecar) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"revision_sidecar_process_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"revision_sidecar_guid"}),
		strmangle.WhereClause("\"", "\"", 2, revisionSidecarProcessTypePrimaryKeyColumns),
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

	o.RevisionSidecarGUID = related.GUID
	if o.R == nil {
		o.R = &revisionSidecarProcessTypeR{
			RevisionSidecar: related,
		}
	} else {
		o.R.RevisionSidecar = related
	}

	if related.R == nil {
		related.R = &revisionSidecarR{
			RevisionSidecarProcessTypes: RevisionSidecarProcessTypeSlice{o},
		}
	} else {
		related.R.RevisionSidecarProcessTypes = append(related.R.RevisionSidecarProcessTypes, o)
	}

	return nil
}

// RevisionSidecarProcessTypes retrieves all the records using an executor.
func RevisionSidecarProcessTypes(mods ...qm.QueryMod) revisionSidecarProcessTypeQuery {
	mods = append(mods, qm.From("\"revision_sidecar_process_types\""))
	return revisionSidecarProcessTypeQuery{NewQuery(mods...)}
}

// FindRevisionSidecarProcessType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRevisionSidecarProcessType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RevisionSidecarProcessType, error) {
	revisionSidecarProcessTypeObj := &RevisionSidecarProcessType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"revision_sidecar_process_types\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, revisionSidecarProcessTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from revision_sidecar_process_types")
	}

	if err = revisionSidecarProcessTypeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return revisionSidecarProcessTypeObj, err
	}

	return revisionSidecarProcessTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RevisionSidecarProcessType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no revision_sidecar_process_types provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(revisionSidecarProcessTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	revisionSidecarProcessTypeInsertCacheMut.RLock()
	cache, cached := revisionSidecarProcessTypeInsertCache[key]
	revisionSidecarProcessTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			revisionSidecarProcessTypeAllColumns,
			revisionSidecarProcessTypeColumnsWithDefault,
			revisionSidecarProcessTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"revision_sidecar_process_types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"revision_sidecar_process_types\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into revision_sidecar_process_types")
	}

	if !cached {
		revisionSidecarProcessTypeInsertCacheMut.Lock()
		revisionSidecarProcessTypeInsertCache[key] = cache
		revisionSidecarProcessTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RevisionSidecarProcessType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RevisionSidecarProcessType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	revisionSidecarProcessTypeUpdateCacheMut.RLock()
	cache, cached := revisionSidecarProcessTypeUpdateCache[key]
	revisionSidecarProcessTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			revisionSidecarProcessTypeAllColumns,
			revisionSidecarProcessTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update revision_sidecar_process_types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"revision_sidecar_process_types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, revisionSidecarProcessTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, append(wl, revisionSidecarProcessTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update revision_sidecar_process_types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for revision_sidecar_process_types")
	}

	if !cached {
		revisionSidecarProcessTypeUpdateCacheMut.Lock()
		revisionSidecarProcessTypeUpdateCache[key] = cache
		revisionSidecarProcessTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q revisionSidecarProcessTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for revision_sidecar_process_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for revision_sidecar_process_types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RevisionSidecarProcessTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarProcessTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"revision_sidecar_process_types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, revisionSidecarProcessTypePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in revisionSidecarProcessType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all revisionSidecarProcessType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RevisionSidecarProcessType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no revision_sidecar_process_types provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(revisionSidecarProcessTypeColumnsWithDefault, o)

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

	revisionSidecarProcessTypeUpsertCacheMut.RLock()
	cache, cached := revisionSidecarProcessTypeUpsertCache[key]
	revisionSidecarProcessTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			revisionSidecarProcessTypeAllColumns,
			revisionSidecarProcessTypeColumnsWithDefault,
			revisionSidecarProcessTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			revisionSidecarProcessTypeAllColumns,
			revisionSidecarProcessTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert revision_sidecar_process_types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(revisionSidecarProcessTypePrimaryKeyColumns))
			copy(conflict, revisionSidecarProcessTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"revision_sidecar_process_types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(revisionSidecarProcessTypeType, revisionSidecarProcessTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert revision_sidecar_process_types")
	}

	if !cached {
		revisionSidecarProcessTypeUpsertCacheMut.Lock()
		revisionSidecarProcessTypeUpsertCache[key] = cache
		revisionSidecarProcessTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RevisionSidecarProcessType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RevisionSidecarProcessType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RevisionSidecarProcessType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), revisionSidecarProcessTypePrimaryKeyMapping)
	sql := "DELETE FROM \"revision_sidecar_process_types\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from revision_sidecar_process_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for revision_sidecar_process_types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q revisionSidecarProcessTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no revisionSidecarProcessTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revision_sidecar_process_types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revision_sidecar_process_types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RevisionSidecarProcessTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(revisionSidecarProcessTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarProcessTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"revision_sidecar_process_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revisionSidecarProcessTypePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from revisionSidecarProcessType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for revision_sidecar_process_types")
	}

	if len(revisionSidecarProcessTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *RevisionSidecarProcessType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRevisionSidecarProcessType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RevisionSidecarProcessTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RevisionSidecarProcessTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), revisionSidecarProcessTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"revision_sidecar_process_types\".* FROM \"revision_sidecar_process_types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, revisionSidecarProcessTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RevisionSidecarProcessTypeSlice")
	}

	*o = slice

	return nil
}

// RevisionSidecarProcessTypeExists checks if the RevisionSidecarProcessType row exists.
func RevisionSidecarProcessTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"revision_sidecar_process_types\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if revision_sidecar_process_types exists")
	}

	return exists, nil
}
