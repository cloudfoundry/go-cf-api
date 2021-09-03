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

// RouteBindingAnnotation is an object representing the database table.
type RouteBindingAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *routeBindingAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L routeBindingAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RouteBindingAnnotationColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ResourceGUID: "resource_guid",
	KeyPrefix:    "key_prefix",
	Key:          "key",
	Value:        "value",
}

var RouteBindingAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "route_binding_annotations.id",
	GUID:         "route_binding_annotations.guid",
	CreatedAt:    "route_binding_annotations.created_at",
	UpdatedAt:    "route_binding_annotations.updated_at",
	ResourceGUID: "route_binding_annotations.resource_guid",
	KeyPrefix:    "route_binding_annotations.key_prefix",
	Key:          "route_binding_annotations.key",
	Value:        "route_binding_annotations.value",
}

// Generated where

var RouteBindingAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`route_binding_annotations`.`id`"},
	GUID:         whereHelperstring{field: "`route_binding_annotations`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`route_binding_annotations`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`route_binding_annotations`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`route_binding_annotations`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`route_binding_annotations`.`key_prefix`"},
	Key:          whereHelpernull_String{field: "`route_binding_annotations`.`key`"},
	Value:        whereHelpernull_String{field: "`route_binding_annotations`.`value`"},
}

// RouteBindingAnnotationRels is where relationship names are stored.
var RouteBindingAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// routeBindingAnnotationR is where relationships are stored.
type routeBindingAnnotationR struct {
	Resource *RouteBinding `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*routeBindingAnnotationR) NewStruct() *routeBindingAnnotationR {
	return &routeBindingAnnotationR{}
}

// routeBindingAnnotationL is where Load methods for each relationship are stored.
type routeBindingAnnotationL struct{}

var (
	routeBindingAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	routeBindingAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	routeBindingAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	routeBindingAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// RouteBindingAnnotationSlice is an alias for a slice of pointers to RouteBindingAnnotation.
	// This should almost always be used instead of []RouteBindingAnnotation.
	RouteBindingAnnotationSlice []*RouteBindingAnnotation

	routeBindingAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	routeBindingAnnotationType                 = reflect.TypeOf(&RouteBindingAnnotation{})
	routeBindingAnnotationMapping              = queries.MakeStructMapping(routeBindingAnnotationType)
	routeBindingAnnotationPrimaryKeyMapping, _ = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, routeBindingAnnotationPrimaryKeyColumns)
	routeBindingAnnotationInsertCacheMut       sync.RWMutex
	routeBindingAnnotationInsertCache          = make(map[string]insertCache)
	routeBindingAnnotationUpdateCacheMut       sync.RWMutex
	routeBindingAnnotationUpdateCache          = make(map[string]updateCache)
	routeBindingAnnotationUpsertCacheMut       sync.RWMutex
	routeBindingAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type RouteBindingAnnotationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*RouteBindingAnnotation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (RouteBindingAnnotationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single routeBindingAnnotation record from the query.
func (q routeBindingAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RouteBindingAnnotation, error) {
	o := &RouteBindingAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for route_binding_annotations")
	}

	return o, nil
}

// All returns all RouteBindingAnnotation records from the query.
func (q routeBindingAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (RouteBindingAnnotationSlice, error) {
	var o []*RouteBindingAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RouteBindingAnnotation slice")
	}

	return o, nil
}

// Count returns the count of all RouteBindingAnnotation records in the query.
func (q routeBindingAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count route_binding_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q routeBindingAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if route_binding_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q routeBindingAnnotationQuery) Resource(o *RouteBindingAnnotation, mods ...qm.QueryMod) routeBindingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := RouteBindings(queryMods...)
	queries.SetFrom(query.Query, "`route_bindings`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (routeBindingAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRouteBindingAnnotation interface{}, mods queries.Applicator) error {
	var slice []*RouteBindingAnnotation
	var object *RouteBindingAnnotation

	if singular {
		object = maybeRouteBindingAnnotation.(*RouteBindingAnnotation)
	} else {
		slice = *maybeRouteBindingAnnotation.(*[]*RouteBindingAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &routeBindingAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &routeBindingAnnotationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ResourceGUID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ResourceGUID) {
				args = append(args, obj.ResourceGUID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`route_bindings`),
		qm.WhereIn(`route_bindings.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RouteBinding")
	}

	var resultSlice []*RouteBinding
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RouteBinding")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for route_bindings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for route_bindings")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &routeBindingR{}
		}
		foreign.R.ResourceRouteBindingAnnotations = append(foreign.R.ResourceRouteBindingAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &routeBindingR{}
				}
				foreign.R.ResourceRouteBindingAnnotations = append(foreign.R.ResourceRouteBindingAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the routeBindingAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceRouteBindingAnnotations.
func (q routeBindingAnnotationQuery) SetResource(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *RouteBinding) error {
	var err error
	if insert {
		if err = RouteBindings().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `route_binding_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, routeBindingAnnotationPrimaryKeyColumns),
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

	queries.Assign(&o.ResourceGUID, related.GUID)
	if o.R == nil {
		o.R = &routeBindingAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &routeBindingR{
			ResourceRouteBindingAnnotations: RouteBindingAnnotationSlice{o},
		}
	} else {
		related.R.ResourceRouteBindingAnnotations = append(related.R.ResourceRouteBindingAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q routeBindingAnnotationQuery) RemoveResource(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, related *RouteBinding) error {
	var err error

	queries.SetScanner(&o.ResourceGUID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Resource = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ResourceRouteBindingAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceRouteBindingAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceRouteBindingAnnotations[i] = related.R.ResourceRouteBindingAnnotations[ln-1]
		}
		related.R.ResourceRouteBindingAnnotations = related.R.ResourceRouteBindingAnnotations[:ln-1]
		break
	}
	return nil
}

// RouteBindingAnnotations retrieves all the records using an executor.
func RouteBindingAnnotations(mods ...qm.QueryMod) routeBindingAnnotationQuery {
	mods = append(mods, qm.From("`route_binding_annotations`"))
	return routeBindingAnnotationQuery{NewQuery(mods...)}
}

type RouteBindingAnnotationFinder interface {
	FindRouteBindingAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RouteBindingAnnotation, error)
}

// FindRouteBindingAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRouteBindingAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RouteBindingAnnotation, error) {
	routeBindingAnnotationObj := &RouteBindingAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `route_binding_annotations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, routeBindingAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from route_binding_annotations")
	}

	return routeBindingAnnotationObj, nil
}

type RouteBindingAnnotationInserter interface {
	Insert(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q routeBindingAnnotationQuery) Insert(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_binding_annotations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(routeBindingAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	routeBindingAnnotationInsertCacheMut.RLock()
	cache, cached := routeBindingAnnotationInsertCache[key]
	routeBindingAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			routeBindingAnnotationAllColumns,
			routeBindingAnnotationColumnsWithDefault,
			routeBindingAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `route_binding_annotations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `route_binding_annotations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `route_binding_annotations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, routeBindingAnnotationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into route_binding_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == routeBindingAnnotationMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for route_binding_annotations")
	}

CacheNoHooks:
	if !cached {
		routeBindingAnnotationInsertCacheMut.Lock()
		routeBindingAnnotationInsertCache[key] = cache
		routeBindingAnnotationInsertCacheMut.Unlock()
	}

	return nil
}

type RouteBindingAnnotationUpdater interface {
	Update(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the RouteBindingAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q routeBindingAnnotationQuery) Update(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	routeBindingAnnotationUpdateCacheMut.RLock()
	cache, cached := routeBindingAnnotationUpdateCache[key]
	routeBindingAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			routeBindingAnnotationAllColumns,
			routeBindingAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update route_binding_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `route_binding_annotations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, routeBindingAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, append(wl, routeBindingAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update route_binding_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for route_binding_annotations")
	}

	if !cached {
		routeBindingAnnotationUpdateCacheMut.Lock()
		routeBindingAnnotationUpdateCache[key] = cache
		routeBindingAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q routeBindingAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for route_binding_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for route_binding_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q routeBindingAnnotationQuery) UpdateAllSlice(o RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `route_binding_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeBindingAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in routeBindingAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all routeBindingAnnotation")
	}
	return rowsAff, nil
}

type RouteBindingAnnotationUpserter interface {
	Upsert(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLRouteBindingAnnotationUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q routeBindingAnnotationQuery) Upsert(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_binding_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(routeBindingAnnotationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRouteBindingAnnotationUniqueColumns, o)

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

	routeBindingAnnotationUpsertCacheMut.RLock()
	cache, cached := routeBindingAnnotationUpsertCache[key]
	routeBindingAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			routeBindingAnnotationAllColumns,
			routeBindingAnnotationColumnsWithDefault,
			routeBindingAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			routeBindingAnnotationAllColumns,
			routeBindingAnnotationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert route_binding_annotations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`route_binding_annotations`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `route_binding_annotations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for route_binding_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == routeBindingAnnotationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(routeBindingAnnotationType, routeBindingAnnotationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for route_binding_annotations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for route_binding_annotations")
	}

CacheNoHooks:
	if !cached {
		routeBindingAnnotationUpsertCacheMut.Lock()
		routeBindingAnnotationUpsertCache[key] = cache
		routeBindingAnnotationUpsertCacheMut.Unlock()
	}

	return nil
}

type RouteBindingAnnotationDeleter interface {
	Delete(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single RouteBindingAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q routeBindingAnnotationQuery) Delete(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RouteBindingAnnotation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), routeBindingAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM `route_binding_annotations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from route_binding_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for route_binding_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q routeBindingAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no routeBindingAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from route_binding_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_binding_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q routeBindingAnnotationQuery) DeleteAllSlice(o RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `route_binding_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeBindingAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from routeBindingAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_binding_annotations")
	}

	return rowsAff, nil
}

type RouteBindingAnnotationReloader interface {
	Reload(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q routeBindingAnnotationQuery) Reload(o *RouteBindingAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRouteBindingAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q routeBindingAnnotationQuery) ReloadAll(o *RouteBindingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RouteBindingAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `route_binding_annotations`.* FROM `route_binding_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeBindingAnnotationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RouteBindingAnnotationSlice")
	}

	*o = slice

	return nil
}

// RouteBindingAnnotationExists checks if the RouteBindingAnnotation row exists.
func RouteBindingAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `route_binding_annotations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if route_binding_annotations exists")
	}

	return exists, nil
}
