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

type RouteMappingUpserter interface {
	Upsert(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLRouteMappingUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q routeMappingQuery) Upsert(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_mappings provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(routeMappingColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRouteMappingUniqueColumns, o)

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

	routeMappingUpsertCacheMut.RLock()
	cache, cached := routeMappingUpsertCache[key]
	routeMappingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			routeMappingAllColumns,
			routeMappingColumnsWithDefault,
			routeMappingColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			routeMappingAllColumns,
			routeMappingPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert route_mappings, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`route_mappings`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `route_mappings` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(routeMappingType, routeMappingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(routeMappingType, routeMappingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for route_mappings")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == routeMappingMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(routeMappingType, routeMappingMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for route_mappings")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for route_mappings")
	}

CacheNoHooks:
	if !cached {
		routeMappingUpsertCacheMut.Lock()
		routeMappingUpsertCache[key] = cache
		routeMappingUpsertCacheMut.Unlock()
	}

	return nil
}

// RouteMapping is an object representing the database table.
type RouteMapping struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	AppPort     null.Int    `boil:"app_port" json:"app_port,omitempty" toml:"app_port" yaml:"app_port,omitempty"`
	GUID        string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	AppGUID     string      `boil:"app_guid" json:"app_guid" toml:"app_guid" yaml:"app_guid"`
	RouteGUID   string      `boil:"route_guid" json:"route_guid" toml:"route_guid" yaml:"route_guid"`
	ProcessType null.String `boil:"process_type" json:"process_type,omitempty" toml:"process_type" yaml:"process_type,omitempty"`
	Weight      null.Int    `boil:"weight" json:"weight,omitempty" toml:"weight" yaml:"weight,omitempty"`

	R *routeMappingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L routeMappingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RouteMappingColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	AppPort     string
	GUID        string
	AppGUID     string
	RouteGUID   string
	ProcessType string
	Weight      string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	AppPort:     "app_port",
	GUID:        "guid",
	AppGUID:     "app_guid",
	RouteGUID:   "route_guid",
	ProcessType: "process_type",
	Weight:      "weight",
}

var RouteMappingTableColumns = struct {
	ID          string
	CreatedAt   string
	UpdatedAt   string
	AppPort     string
	GUID        string
	AppGUID     string
	RouteGUID   string
	ProcessType string
	Weight      string
}{
	ID:          "route_mappings.id",
	CreatedAt:   "route_mappings.created_at",
	UpdatedAt:   "route_mappings.updated_at",
	AppPort:     "route_mappings.app_port",
	GUID:        "route_mappings.guid",
	AppGUID:     "route_mappings.app_guid",
	RouteGUID:   "route_mappings.route_guid",
	ProcessType: "route_mappings.process_type",
	Weight:      "route_mappings.weight",
}

// Generated where

var RouteMappingWhere = struct {
	ID          whereHelperint
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpernull_Time
	AppPort     whereHelpernull_Int
	GUID        whereHelperstring
	AppGUID     whereHelperstring
	RouteGUID   whereHelperstring
	ProcessType whereHelpernull_String
	Weight      whereHelpernull_Int
}{
	ID:          whereHelperint{field: "`route_mappings`.`id`"},
	CreatedAt:   whereHelpertime_Time{field: "`route_mappings`.`created_at`"},
	UpdatedAt:   whereHelpernull_Time{field: "`route_mappings`.`updated_at`"},
	AppPort:     whereHelpernull_Int{field: "`route_mappings`.`app_port`"},
	GUID:        whereHelperstring{field: "`route_mappings`.`guid`"},
	AppGUID:     whereHelperstring{field: "`route_mappings`.`app_guid`"},
	RouteGUID:   whereHelperstring{field: "`route_mappings`.`route_guid`"},
	ProcessType: whereHelpernull_String{field: "`route_mappings`.`process_type`"},
	Weight:      whereHelpernull_Int{field: "`route_mappings`.`weight`"},
}

// RouteMappingRels is where relationship names are stored.
var RouteMappingRels = struct {
	App   string
	Route string
}{
	App:   "App",
	Route: "Route",
}

// routeMappingR is where relationships are stored.
type routeMappingR struct {
	App   *App   `boil:"App" json:"App" toml:"App" yaml:"App"`
	Route *Route `boil:"Route" json:"Route" toml:"Route" yaml:"Route"`
}

// NewStruct creates a new relationship struct
func (*routeMappingR) NewStruct() *routeMappingR {
	return &routeMappingR{}
}

// routeMappingL is where Load methods for each relationship are stored.
type routeMappingL struct{}

var (
	routeMappingAllColumns            = []string{"id", "created_at", "updated_at", "app_port", "guid", "app_guid", "route_guid", "process_type", "weight"}
	routeMappingColumnsWithoutDefault = []string{"updated_at", "guid", "app_guid", "route_guid", "process_type", "weight"}
	routeMappingColumnsWithDefault    = []string{"id", "created_at", "app_port"}
	routeMappingPrimaryKeyColumns     = []string{"id"}
)

type (
	// RouteMappingSlice is an alias for a slice of pointers to RouteMapping.
	// This should almost always be used instead of []RouteMapping.
	RouteMappingSlice []*RouteMapping

	routeMappingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	routeMappingType                 = reflect.TypeOf(&RouteMapping{})
	routeMappingMapping              = queries.MakeStructMapping(routeMappingType)
	routeMappingPrimaryKeyMapping, _ = queries.BindMapping(routeMappingType, routeMappingMapping, routeMappingPrimaryKeyColumns)
	routeMappingInsertCacheMut       sync.RWMutex
	routeMappingInsertCache          = make(map[string]insertCache)
	routeMappingUpdateCacheMut       sync.RWMutex
	routeMappingUpdateCache          = make(map[string]updateCache)
	routeMappingUpsertCacheMut       sync.RWMutex
	routeMappingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type RouteMappingFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*RouteMapping, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (RouteMappingSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single routeMapping record from the query.
func (q routeMappingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RouteMapping, error) {
	o := &RouteMapping{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for route_mappings")
	}

	return o, nil
}

// All returns all RouteMapping records from the query.
func (q routeMappingQuery) All(ctx context.Context, exec boil.ContextExecutor) (RouteMappingSlice, error) {
	var o []*RouteMapping

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RouteMapping slice")
	}

	return o, nil
}

// Count returns the count of all RouteMapping records in the query.
func (q routeMappingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count route_mappings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q routeMappingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if route_mappings exists")
	}

	return count > 0, nil
}

// App pointed to by the foreign key.
func (q routeMappingQuery) App(o *RouteMapping, mods ...qm.QueryMod) appQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.AppGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Apps(queryMods...)
	queries.SetFrom(query.Query, "`apps`")

	return query
}

// Route pointed to by the foreign key.
func (q routeMappingQuery) Route(o *RouteMapping, mods ...qm.QueryMod) routeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.RouteGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Routes(queryMods...)
	queries.SetFrom(query.Query, "`routes`")

	return query
}

// LoadApp allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (routeMappingL) LoadApp(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRouteMapping interface{}, mods queries.Applicator) error {
	var slice []*RouteMapping
	var object *RouteMapping

	if singular {
		object = maybeRouteMapping.(*RouteMapping)
	} else {
		slice = *maybeRouteMapping.(*[]*RouteMapping)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &routeMappingR{}
		}
		args = append(args, object.AppGUID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &routeMappingR{}
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
		foreign.R.RouteMappings = append(foreign.R.RouteMappings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AppGUID == foreign.GUID {
				local.R.App = foreign
				if foreign.R == nil {
					foreign.R = &appR{}
				}
				foreign.R.RouteMappings = append(foreign.R.RouteMappings, local)
				break
			}
		}
	}

	return nil
}

// LoadRoute allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (routeMappingL) LoadRoute(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRouteMapping interface{}, mods queries.Applicator) error {
	var slice []*RouteMapping
	var object *RouteMapping

	if singular {
		object = maybeRouteMapping.(*RouteMapping)
	} else {
		slice = *maybeRouteMapping.(*[]*RouteMapping)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &routeMappingR{}
		}
		args = append(args, object.RouteGUID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &routeMappingR{}
			}

			for _, a := range args {
				if a == obj.RouteGUID {
					continue Outer
				}
			}

			args = append(args, obj.RouteGUID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`routes`),
		qm.WhereIn(`routes.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Route")
	}

	var resultSlice []*Route
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Route")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for routes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for routes")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Route = foreign
		if foreign.R == nil {
			foreign.R = &routeR{}
		}
		foreign.R.RouteMappings = append(foreign.R.RouteMappings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RouteGUID == foreign.GUID {
				local.R.Route = foreign
				if foreign.R == nil {
					foreign.R = &routeR{}
				}
				foreign.R.RouteMappings = append(foreign.R.RouteMappings, local)
				break
			}
		}
	}

	return nil
}

// SetApp of the routeMapping to the related item.
// Sets o.R.App to related.
// Adds o to related.R.RouteMappings.
func (q routeMappingQuery) SetApp(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, insert bool, related *App) error {
	var err error
	if insert {
		if err = Apps().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `route_mappings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"app_guid"}),
		strmangle.WhereClause("`", "`", 0, routeMappingPrimaryKeyColumns),
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
		o.R = &routeMappingR{
			App: related,
		}
	} else {
		o.R.App = related
	}

	if related.R == nil {
		related.R = &appR{
			RouteMappings: RouteMappingSlice{o},
		}
	} else {
		related.R.RouteMappings = append(related.R.RouteMappings, o)
	}

	return nil
}

// SetRoute of the routeMapping to the related item.
// Sets o.R.Route to related.
// Adds o to related.R.RouteMappings.
func (q routeMappingQuery) SetRoute(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Route) error {
	var err error
	if insert {
		if err = Routes().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `route_mappings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"route_guid"}),
		strmangle.WhereClause("`", "`", 0, routeMappingPrimaryKeyColumns),
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

	o.RouteGUID = related.GUID
	if o.R == nil {
		o.R = &routeMappingR{
			Route: related,
		}
	} else {
		o.R.Route = related
	}

	if related.R == nil {
		related.R = &routeR{
			RouteMappings: RouteMappingSlice{o},
		}
	} else {
		related.R.RouteMappings = append(related.R.RouteMappings, o)
	}

	return nil
}

// RouteMappings retrieves all the records using an executor.
func RouteMappings(mods ...qm.QueryMod) routeMappingQuery {
	mods = append(mods, qm.From("`route_mappings`"))
	return routeMappingQuery{NewQuery(mods...)}
}

type RouteMappingFinder interface {
	FindRouteMapping(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RouteMapping, error)
}

// FindRouteMapping retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRouteMapping(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RouteMapping, error) {
	routeMappingObj := &RouteMapping{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `route_mappings` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, routeMappingObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from route_mappings")
	}

	return routeMappingObj, nil
}

type RouteMappingInserter interface {
	Insert(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q routeMappingQuery) Insert(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_mappings provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(routeMappingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	routeMappingInsertCacheMut.RLock()
	cache, cached := routeMappingInsertCache[key]
	routeMappingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			routeMappingAllColumns,
			routeMappingColumnsWithDefault,
			routeMappingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(routeMappingType, routeMappingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(routeMappingType, routeMappingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `route_mappings` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `route_mappings` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `route_mappings` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, routeMappingPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into route_mappings")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == routeMappingMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for route_mappings")
	}

CacheNoHooks:
	if !cached {
		routeMappingInsertCacheMut.Lock()
		routeMappingInsertCache[key] = cache
		routeMappingInsertCacheMut.Unlock()
	}

	return nil
}

type RouteMappingUpdater interface {
	Update(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the RouteMapping.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q routeMappingQuery) Update(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	routeMappingUpdateCacheMut.RLock()
	cache, cached := routeMappingUpdateCache[key]
	routeMappingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			routeMappingAllColumns,
			routeMappingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update route_mappings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `route_mappings` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, routeMappingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(routeMappingType, routeMappingMapping, append(wl, routeMappingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update route_mappings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for route_mappings")
	}

	if !cached {
		routeMappingUpdateCacheMut.Lock()
		routeMappingUpdateCache[key] = cache
		routeMappingUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q routeMappingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for route_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for route_mappings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q routeMappingQuery) UpdateAllSlice(o RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `route_mappings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeMappingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in routeMapping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all routeMapping")
	}
	return rowsAff, nil
}

type RouteMappingDeleter interface {
	Delete(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single RouteMapping record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q routeMappingQuery) Delete(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RouteMapping provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), routeMappingPrimaryKeyMapping)
	sql := "DELETE FROM `route_mappings` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from route_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for route_mappings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q routeMappingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no routeMappingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from route_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_mappings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q routeMappingQuery) DeleteAllSlice(o RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `route_mappings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeMappingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from routeMapping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_mappings")
	}

	return rowsAff, nil
}

type RouteMappingReloader interface {
	Reload(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q routeMappingQuery) Reload(o *RouteMapping, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRouteMapping(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q routeMappingQuery) ReloadAll(o *RouteMappingSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RouteMappingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `route_mappings`.* FROM `route_mappings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, routeMappingPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RouteMappingSlice")
	}

	*o = slice

	return nil
}

// RouteMappingExists checks if the RouteMapping row exists.
func RouteMappingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `route_mappings` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if route_mappings exists")
	}

	return exists, nil
}
