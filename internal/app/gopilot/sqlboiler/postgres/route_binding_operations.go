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

// RouteBindingOperation is an object representing the database table.
type RouteBindingOperation struct {
	ID                      int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	RouteBindingID          null.Int    `boil:"route_binding_id" json:"route_binding_id,omitempty" toml:"route_binding_id" yaml:"route_binding_id,omitempty"`
	State                   string      `boil:"state" json:"state" toml:"state" yaml:"state"`
	Type                    string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Description             null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	BrokerProvidedOperation null.String `boil:"broker_provided_operation" json:"broker_provided_operation,omitempty" toml:"broker_provided_operation" yaml:"broker_provided_operation,omitempty"`

	R *routeBindingOperationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L routeBindingOperationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RouteBindingOperationColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	RouteBindingID          string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "id",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	RouteBindingID:          "route_binding_id",
	State:                   "state",
	Type:                    "type",
	Description:             "description",
	BrokerProvidedOperation: "broker_provided_operation",
}

var RouteBindingOperationTableColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	RouteBindingID          string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "route_binding_operations.id",
	CreatedAt:               "route_binding_operations.created_at",
	UpdatedAt:               "route_binding_operations.updated_at",
	RouteBindingID:          "route_binding_operations.route_binding_id",
	State:                   "route_binding_operations.state",
	Type:                    "route_binding_operations.type",
	Description:             "route_binding_operations.description",
	BrokerProvidedOperation: "route_binding_operations.broker_provided_operation",
}

// Generated where

var RouteBindingOperationWhere = struct {
	ID                      whereHelperint
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpernull_Time
	RouteBindingID          whereHelpernull_Int
	State                   whereHelperstring
	Type                    whereHelperstring
	Description             whereHelpernull_String
	BrokerProvidedOperation whereHelpernull_String
}{
	ID:                      whereHelperint{field: "\"route_binding_operations\".\"id\""},
	CreatedAt:               whereHelpertime_Time{field: "\"route_binding_operations\".\"created_at\""},
	UpdatedAt:               whereHelpernull_Time{field: "\"route_binding_operations\".\"updated_at\""},
	RouteBindingID:          whereHelpernull_Int{field: "\"route_binding_operations\".\"route_binding_id\""},
	State:                   whereHelperstring{field: "\"route_binding_operations\".\"state\""},
	Type:                    whereHelperstring{field: "\"route_binding_operations\".\"type\""},
	Description:             whereHelpernull_String{field: "\"route_binding_operations\".\"description\""},
	BrokerProvidedOperation: whereHelpernull_String{field: "\"route_binding_operations\".\"broker_provided_operation\""},
}

// RouteBindingOperationRels is where relationship names are stored.
var RouteBindingOperationRels = struct {
	RouteBinding string
}{
	RouteBinding: "RouteBinding",
}

// routeBindingOperationR is where relationships are stored.
type routeBindingOperationR struct {
	RouteBinding *RouteBinding `boil:"RouteBinding" json:"RouteBinding" toml:"RouteBinding" yaml:"RouteBinding"`
}

// NewStruct creates a new relationship struct
func (*routeBindingOperationR) NewStruct() *routeBindingOperationR {
	return &routeBindingOperationR{}
}

// routeBindingOperationL is where Load methods for each relationship are stored.
type routeBindingOperationL struct{}

var (
	routeBindingOperationAllColumns            = []string{"id", "created_at", "updated_at", "route_binding_id", "state", "type", "description", "broker_provided_operation"}
	routeBindingOperationColumnsWithoutDefault = []string{"updated_at", "route_binding_id", "state", "type", "description", "broker_provided_operation"}
	routeBindingOperationColumnsWithDefault    = []string{"id", "created_at"}
	routeBindingOperationPrimaryKeyColumns     = []string{"id"}
)

type (
	// RouteBindingOperationSlice is an alias for a slice of pointers to RouteBindingOperation.
	// This should almost always be used instead of []RouteBindingOperation.
	RouteBindingOperationSlice []*RouteBindingOperation
	// RouteBindingOperationHook is the signature for custom RouteBindingOperation hook methods
	RouteBindingOperationHook func(context.Context, boil.ContextExecutor, *RouteBindingOperation) error

	routeBindingOperationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	routeBindingOperationType                 = reflect.TypeOf(&RouteBindingOperation{})
	routeBindingOperationMapping              = queries.MakeStructMapping(routeBindingOperationType)
	routeBindingOperationPrimaryKeyMapping, _ = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, routeBindingOperationPrimaryKeyColumns)
	routeBindingOperationInsertCacheMut       sync.RWMutex
	routeBindingOperationInsertCache          = make(map[string]insertCache)
	routeBindingOperationUpdateCacheMut       sync.RWMutex
	routeBindingOperationUpdateCache          = make(map[string]updateCache)
	routeBindingOperationUpsertCacheMut       sync.RWMutex
	routeBindingOperationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var routeBindingOperationBeforeInsertHooks []RouteBindingOperationHook
var routeBindingOperationBeforeUpdateHooks []RouteBindingOperationHook
var routeBindingOperationBeforeDeleteHooks []RouteBindingOperationHook
var routeBindingOperationBeforeUpsertHooks []RouteBindingOperationHook

var routeBindingOperationAfterInsertHooks []RouteBindingOperationHook
var routeBindingOperationAfterSelectHooks []RouteBindingOperationHook
var routeBindingOperationAfterUpdateHooks []RouteBindingOperationHook
var routeBindingOperationAfterDeleteHooks []RouteBindingOperationHook
var routeBindingOperationAfterUpsertHooks []RouteBindingOperationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RouteBindingOperation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RouteBindingOperation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RouteBindingOperation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RouteBindingOperation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RouteBindingOperation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RouteBindingOperation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RouteBindingOperation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RouteBindingOperation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RouteBindingOperation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range routeBindingOperationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRouteBindingOperationHook registers your hook function for all future operations.
func AddRouteBindingOperationHook(hookPoint boil.HookPoint, routeBindingOperationHook RouteBindingOperationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		routeBindingOperationBeforeInsertHooks = append(routeBindingOperationBeforeInsertHooks, routeBindingOperationHook)
	case boil.BeforeUpdateHook:
		routeBindingOperationBeforeUpdateHooks = append(routeBindingOperationBeforeUpdateHooks, routeBindingOperationHook)
	case boil.BeforeDeleteHook:
		routeBindingOperationBeforeDeleteHooks = append(routeBindingOperationBeforeDeleteHooks, routeBindingOperationHook)
	case boil.BeforeUpsertHook:
		routeBindingOperationBeforeUpsertHooks = append(routeBindingOperationBeforeUpsertHooks, routeBindingOperationHook)
	case boil.AfterInsertHook:
		routeBindingOperationAfterInsertHooks = append(routeBindingOperationAfterInsertHooks, routeBindingOperationHook)
	case boil.AfterSelectHook:
		routeBindingOperationAfterSelectHooks = append(routeBindingOperationAfterSelectHooks, routeBindingOperationHook)
	case boil.AfterUpdateHook:
		routeBindingOperationAfterUpdateHooks = append(routeBindingOperationAfterUpdateHooks, routeBindingOperationHook)
	case boil.AfterDeleteHook:
		routeBindingOperationAfterDeleteHooks = append(routeBindingOperationAfterDeleteHooks, routeBindingOperationHook)
	case boil.AfterUpsertHook:
		routeBindingOperationAfterUpsertHooks = append(routeBindingOperationAfterUpsertHooks, routeBindingOperationHook)
	}
}

// One returns a single routeBindingOperation record from the query.
func (q routeBindingOperationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RouteBindingOperation, error) {
	o := &RouteBindingOperation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for route_binding_operations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RouteBindingOperation records from the query.
func (q routeBindingOperationQuery) All(ctx context.Context, exec boil.ContextExecutor) (RouteBindingOperationSlice, error) {
	var o []*RouteBindingOperation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RouteBindingOperation slice")
	}

	if len(routeBindingOperationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RouteBindingOperation records in the query.
func (q routeBindingOperationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count route_binding_operations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q routeBindingOperationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if route_binding_operations exists")
	}

	return count > 0, nil
}

// RouteBinding pointed to by the foreign key.
func (o *RouteBindingOperation) RouteBinding(mods ...qm.QueryMod) routeBindingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.RouteBindingID),
	}

	queryMods = append(queryMods, mods...)

	query := RouteBindings(queryMods...)
	queries.SetFrom(query.Query, "\"route_bindings\"")

	return query
}

// LoadRouteBinding allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (routeBindingOperationL) LoadRouteBinding(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRouteBindingOperation interface{}, mods queries.Applicator) error {
	var slice []*RouteBindingOperation
	var object *RouteBindingOperation

	if singular {
		object = maybeRouteBindingOperation.(*RouteBindingOperation)
	} else {
		slice = *maybeRouteBindingOperation.(*[]*RouteBindingOperation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &routeBindingOperationR{}
		}
		if !queries.IsNil(object.RouteBindingID) {
			args = append(args, object.RouteBindingID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &routeBindingOperationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.RouteBindingID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.RouteBindingID) {
				args = append(args, obj.RouteBindingID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`route_bindings`),
		qm.WhereIn(`route_bindings.id in ?`, args...),
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

	if len(routeBindingOperationAfterSelectHooks) != 0 {
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
		object.R.RouteBinding = foreign
		if foreign.R == nil {
			foreign.R = &routeBindingR{}
		}
		foreign.R.RouteBindingOperation = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.RouteBindingID, foreign.ID) {
				local.R.RouteBinding = foreign
				if foreign.R == nil {
					foreign.R = &routeBindingR{}
				}
				foreign.R.RouteBindingOperation = local
				break
			}
		}
	}

	return nil
}

// SetRouteBinding of the routeBindingOperation to the related item.
// Sets o.R.RouteBinding to related.
// Adds o to related.R.RouteBindingOperation.
func (o *RouteBindingOperation) SetRouteBinding(ctx context.Context, exec boil.ContextExecutor, insert bool, related *RouteBinding) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"route_binding_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"route_binding_id"}),
		strmangle.WhereClause("\"", "\"", 2, routeBindingOperationPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.RouteBindingID, related.ID)
	if o.R == nil {
		o.R = &routeBindingOperationR{
			RouteBinding: related,
		}
	} else {
		o.R.RouteBinding = related
	}

	if related.R == nil {
		related.R = &routeBindingR{
			RouteBindingOperation: o,
		}
	} else {
		related.R.RouteBindingOperation = o
	}

	return nil
}

// RemoveRouteBinding relationship.
// Sets o.R.RouteBinding to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *RouteBindingOperation) RemoveRouteBinding(ctx context.Context, exec boil.ContextExecutor, related *RouteBinding) error {
	var err error

	queries.SetScanner(&o.RouteBindingID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("route_binding_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.RouteBinding = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.RouteBindingOperation = nil
	return nil
}

// RouteBindingOperations retrieves all the records using an executor.
func RouteBindingOperations(mods ...qm.QueryMod) routeBindingOperationQuery {
	mods = append(mods, qm.From("\"route_binding_operations\""))
	return routeBindingOperationQuery{NewQuery(mods...)}
}

// FindRouteBindingOperation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRouteBindingOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RouteBindingOperation, error) {
	routeBindingOperationObj := &RouteBindingOperation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"route_binding_operations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, routeBindingOperationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from route_binding_operations")
	}

	if err = routeBindingOperationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return routeBindingOperationObj, err
	}

	return routeBindingOperationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RouteBindingOperation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_binding_operations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(routeBindingOperationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	routeBindingOperationInsertCacheMut.RLock()
	cache, cached := routeBindingOperationInsertCache[key]
	routeBindingOperationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			routeBindingOperationAllColumns,
			routeBindingOperationColumnsWithDefault,
			routeBindingOperationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"route_binding_operations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"route_binding_operations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into route_binding_operations")
	}

	if !cached {
		routeBindingOperationInsertCacheMut.Lock()
		routeBindingOperationInsertCache[key] = cache
		routeBindingOperationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RouteBindingOperation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RouteBindingOperation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	routeBindingOperationUpdateCacheMut.RLock()
	cache, cached := routeBindingOperationUpdateCache[key]
	routeBindingOperationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			routeBindingOperationAllColumns,
			routeBindingOperationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update route_binding_operations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"route_binding_operations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, routeBindingOperationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, append(wl, routeBindingOperationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update route_binding_operations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for route_binding_operations")
	}

	if !cached {
		routeBindingOperationUpdateCacheMut.Lock()
		routeBindingOperationUpdateCache[key] = cache
		routeBindingOperationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q routeBindingOperationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for route_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for route_binding_operations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RouteBindingOperationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"route_binding_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, routeBindingOperationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in routeBindingOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all routeBindingOperation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RouteBindingOperation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no route_binding_operations provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(routeBindingOperationColumnsWithDefault, o)

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

	routeBindingOperationUpsertCacheMut.RLock()
	cache, cached := routeBindingOperationUpsertCache[key]
	routeBindingOperationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			routeBindingOperationAllColumns,
			routeBindingOperationColumnsWithDefault,
			routeBindingOperationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			routeBindingOperationAllColumns,
			routeBindingOperationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert route_binding_operations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(routeBindingOperationPrimaryKeyColumns))
			copy(conflict, routeBindingOperationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"route_binding_operations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(routeBindingOperationType, routeBindingOperationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert route_binding_operations")
	}

	if !cached {
		routeBindingOperationUpsertCacheMut.Lock()
		routeBindingOperationUpsertCache[key] = cache
		routeBindingOperationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RouteBindingOperation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RouteBindingOperation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RouteBindingOperation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), routeBindingOperationPrimaryKeyMapping)
	sql := "DELETE FROM \"route_binding_operations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from route_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for route_binding_operations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q routeBindingOperationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no routeBindingOperationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from route_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_binding_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RouteBindingOperationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(routeBindingOperationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"route_binding_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, routeBindingOperationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from routeBindingOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for route_binding_operations")
	}

	if len(routeBindingOperationAfterDeleteHooks) != 0 {
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
func (o *RouteBindingOperation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRouteBindingOperation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RouteBindingOperationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RouteBindingOperationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), routeBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"route_binding_operations\".* FROM \"route_binding_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, routeBindingOperationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RouteBindingOperationSlice")
	}

	*o = slice

	return nil
}

// RouteBindingOperationExists checks if the RouteBindingOperation row exists.
func RouteBindingOperationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"route_binding_operations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if route_binding_operations exists")
	}

	return exists, nil
}
