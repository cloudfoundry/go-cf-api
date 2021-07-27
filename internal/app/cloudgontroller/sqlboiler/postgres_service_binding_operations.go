// +build postgres
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

// ServiceBindingOperation is an object representing the database table.
type ServiceBindingOperation struct {
	ID                      int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ServiceBindingID        null.Int    `boil:"service_binding_id" json:"service_binding_id,omitempty" toml:"service_binding_id" yaml:"service_binding_id,omitempty"`
	State                   string      `boil:"state" json:"state" toml:"state" yaml:"state"`
	Type                    string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Description             null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	BrokerProvidedOperation null.String `boil:"broker_provided_operation" json:"broker_provided_operation,omitempty" toml:"broker_provided_operation" yaml:"broker_provided_operation,omitempty"`

	R *serviceBindingOperationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceBindingOperationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceBindingOperationColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	ServiceBindingID        string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "id",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	ServiceBindingID:        "service_binding_id",
	State:                   "state",
	Type:                    "type",
	Description:             "description",
	BrokerProvidedOperation: "broker_provided_operation",
}

var ServiceBindingOperationTableColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	ServiceBindingID        string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "service_binding_operations.id",
	CreatedAt:               "service_binding_operations.created_at",
	UpdatedAt:               "service_binding_operations.updated_at",
	ServiceBindingID:        "service_binding_operations.service_binding_id",
	State:                   "service_binding_operations.state",
	Type:                    "service_binding_operations.type",
	Description:             "service_binding_operations.description",
	BrokerProvidedOperation: "service_binding_operations.broker_provided_operation",
}

// Generated where

var ServiceBindingOperationWhere = struct {
	ID                      whereHelperint
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpernull_Time
	ServiceBindingID        whereHelpernull_Int
	State                   whereHelperstring
	Type                    whereHelperstring
	Description             whereHelpernull_String
	BrokerProvidedOperation whereHelpernull_String
}{
	ID:                      whereHelperint{field: "\"service_binding_operations\".\"id\""},
	CreatedAt:               whereHelpertime_Time{field: "\"service_binding_operations\".\"created_at\""},
	UpdatedAt:               whereHelpernull_Time{field: "\"service_binding_operations\".\"updated_at\""},
	ServiceBindingID:        whereHelpernull_Int{field: "\"service_binding_operations\".\"service_binding_id\""},
	State:                   whereHelperstring{field: "\"service_binding_operations\".\"state\""},
	Type:                    whereHelperstring{field: "\"service_binding_operations\".\"type\""},
	Description:             whereHelpernull_String{field: "\"service_binding_operations\".\"description\""},
	BrokerProvidedOperation: whereHelpernull_String{field: "\"service_binding_operations\".\"broker_provided_operation\""},
}

// ServiceBindingOperationRels is where relationship names are stored.
var ServiceBindingOperationRels = struct {
	ServiceBinding string
}{
	ServiceBinding: "ServiceBinding",
}

// serviceBindingOperationR is where relationships are stored.
type serviceBindingOperationR struct {
	ServiceBinding *ServiceBinding `boil:"ServiceBinding" json:"ServiceBinding" toml:"ServiceBinding" yaml:"ServiceBinding"`
}

// NewStruct creates a new relationship struct
func (*serviceBindingOperationR) NewStruct() *serviceBindingOperationR {
	return &serviceBindingOperationR{}
}

// serviceBindingOperationL is where Load methods for each relationship are stored.
type serviceBindingOperationL struct{}

var (
	serviceBindingOperationAllColumns            = []string{"id", "created_at", "updated_at", "service_binding_id", "state", "type", "description", "broker_provided_operation"}
	serviceBindingOperationColumnsWithoutDefault = []string{"updated_at", "service_binding_id", "state", "type", "description", "broker_provided_operation"}
	serviceBindingOperationColumnsWithDefault    = []string{"id", "created_at"}
	serviceBindingOperationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceBindingOperationSlice is an alias for a slice of pointers to ServiceBindingOperation.
	// This should almost always be used instead of []ServiceBindingOperation.
	ServiceBindingOperationSlice []*ServiceBindingOperation
	// ServiceBindingOperationHook is the signature for custom ServiceBindingOperation hook methods
	ServiceBindingOperationHook func(context.Context, boil.ContextExecutor, *ServiceBindingOperation) error

	serviceBindingOperationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceBindingOperationType                 = reflect.TypeOf(&ServiceBindingOperation{})
	serviceBindingOperationMapping              = queries.MakeStructMapping(serviceBindingOperationType)
	serviceBindingOperationPrimaryKeyMapping, _ = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, serviceBindingOperationPrimaryKeyColumns)
	serviceBindingOperationInsertCacheMut       sync.RWMutex
	serviceBindingOperationInsertCache          = make(map[string]insertCache)
	serviceBindingOperationUpdateCacheMut       sync.RWMutex
	serviceBindingOperationUpdateCache          = make(map[string]updateCache)
	serviceBindingOperationUpsertCacheMut       sync.RWMutex
	serviceBindingOperationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var serviceBindingOperationBeforeInsertHooks []ServiceBindingOperationHook
var serviceBindingOperationBeforeUpdateHooks []ServiceBindingOperationHook
var serviceBindingOperationBeforeDeleteHooks []ServiceBindingOperationHook
var serviceBindingOperationBeforeUpsertHooks []ServiceBindingOperationHook

var serviceBindingOperationAfterInsertHooks []ServiceBindingOperationHook
var serviceBindingOperationAfterSelectHooks []ServiceBindingOperationHook
var serviceBindingOperationAfterUpdateHooks []ServiceBindingOperationHook
var serviceBindingOperationAfterDeleteHooks []ServiceBindingOperationHook
var serviceBindingOperationAfterUpsertHooks []ServiceBindingOperationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ServiceBindingOperation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ServiceBindingOperation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ServiceBindingOperation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ServiceBindingOperation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ServiceBindingOperation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ServiceBindingOperation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ServiceBindingOperation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ServiceBindingOperation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ServiceBindingOperation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBindingOperationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddServiceBindingOperationHook registers your hook function for all future operations.
func AddServiceBindingOperationHook(hookPoint boil.HookPoint, serviceBindingOperationHook ServiceBindingOperationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		serviceBindingOperationBeforeInsertHooks = append(serviceBindingOperationBeforeInsertHooks, serviceBindingOperationHook)
	case boil.BeforeUpdateHook:
		serviceBindingOperationBeforeUpdateHooks = append(serviceBindingOperationBeforeUpdateHooks, serviceBindingOperationHook)
	case boil.BeforeDeleteHook:
		serviceBindingOperationBeforeDeleteHooks = append(serviceBindingOperationBeforeDeleteHooks, serviceBindingOperationHook)
	case boil.BeforeUpsertHook:
		serviceBindingOperationBeforeUpsertHooks = append(serviceBindingOperationBeforeUpsertHooks, serviceBindingOperationHook)
	case boil.AfterInsertHook:
		serviceBindingOperationAfterInsertHooks = append(serviceBindingOperationAfterInsertHooks, serviceBindingOperationHook)
	case boil.AfterSelectHook:
		serviceBindingOperationAfterSelectHooks = append(serviceBindingOperationAfterSelectHooks, serviceBindingOperationHook)
	case boil.AfterUpdateHook:
		serviceBindingOperationAfterUpdateHooks = append(serviceBindingOperationAfterUpdateHooks, serviceBindingOperationHook)
	case boil.AfterDeleteHook:
		serviceBindingOperationAfterDeleteHooks = append(serviceBindingOperationAfterDeleteHooks, serviceBindingOperationHook)
	case boil.AfterUpsertHook:
		serviceBindingOperationAfterUpsertHooks = append(serviceBindingOperationAfterUpsertHooks, serviceBindingOperationHook)
	}
}

// One returns a single serviceBindingOperation record from the query.
func (q serviceBindingOperationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceBindingOperation, error) {
	o := &ServiceBindingOperation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_binding_operations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ServiceBindingOperation records from the query.
func (q serviceBindingOperationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceBindingOperationSlice, error) {
	var o []*ServiceBindingOperation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceBindingOperation slice")
	}

	if len(serviceBindingOperationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ServiceBindingOperation records in the query.
func (q serviceBindingOperationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_binding_operations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceBindingOperationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_binding_operations exists")
	}

	return count > 0, nil
}

// ServiceBinding pointed to by the foreign key.
func (o *ServiceBindingOperation) ServiceBinding(mods ...qm.QueryMod) serviceBindingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ServiceBindingID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceBindings(queryMods...)
	queries.SetFrom(query.Query, "\"service_bindings\"")

	return query
}

// LoadServiceBinding allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceBindingOperationL) LoadServiceBinding(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceBindingOperation interface{}, mods queries.Applicator) error {
	var slice []*ServiceBindingOperation
	var object *ServiceBindingOperation

	if singular {
		object = maybeServiceBindingOperation.(*ServiceBindingOperation)
	} else {
		slice = *maybeServiceBindingOperation.(*[]*ServiceBindingOperation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceBindingOperationR{}
		}
		if !queries.IsNil(object.ServiceBindingID) {
			args = append(args, object.ServiceBindingID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceBindingOperationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ServiceBindingID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ServiceBindingID) {
				args = append(args, obj.ServiceBindingID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`service_bindings`),
		qm.WhereIn(`service_bindings.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServiceBinding")
	}

	var resultSlice []*ServiceBinding
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServiceBinding")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for service_bindings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for service_bindings")
	}

	if len(serviceBindingOperationAfterSelectHooks) != 0 {
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
		object.R.ServiceBinding = foreign
		if foreign.R == nil {
			foreign.R = &serviceBindingR{}
		}
		foreign.R.ServiceBindingOperation = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ServiceBindingID, foreign.ID) {
				local.R.ServiceBinding = foreign
				if foreign.R == nil {
					foreign.R = &serviceBindingR{}
				}
				foreign.R.ServiceBindingOperation = local
				break
			}
		}
	}

	return nil
}

// SetServiceBinding of the serviceBindingOperation to the related item.
// Sets o.R.ServiceBinding to related.
// Adds o to related.R.ServiceBindingOperation.
func (o *ServiceBindingOperation) SetServiceBinding(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceBinding) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"service_binding_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"service_binding_id"}),
		strmangle.WhereClause("\"", "\"", 2, serviceBindingOperationPrimaryKeyColumns),
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

	queries.Assign(&o.ServiceBindingID, related.ID)
	if o.R == nil {
		o.R = &serviceBindingOperationR{
			ServiceBinding: related,
		}
	} else {
		o.R.ServiceBinding = related
	}

	if related.R == nil {
		related.R = &serviceBindingR{
			ServiceBindingOperation: o,
		}
	} else {
		related.R.ServiceBindingOperation = o
	}

	return nil
}

// RemoveServiceBinding relationship.
// Sets o.R.ServiceBinding to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *ServiceBindingOperation) RemoveServiceBinding(ctx context.Context, exec boil.ContextExecutor, related *ServiceBinding) error {
	var err error

	queries.SetScanner(&o.ServiceBindingID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("service_binding_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ServiceBinding = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.ServiceBindingOperation = nil
	return nil
}

// ServiceBindingOperations retrieves all the records using an executor.
func ServiceBindingOperations(mods ...qm.QueryMod) serviceBindingOperationQuery {
	mods = append(mods, qm.From("\"service_binding_operations\""))
	return serviceBindingOperationQuery{NewQuery(mods...)}
}

// FindServiceBindingOperation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceBindingOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceBindingOperation, error) {
	serviceBindingOperationObj := &ServiceBindingOperation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_binding_operations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceBindingOperationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_binding_operations")
	}

	if err = serviceBindingOperationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return serviceBindingOperationObj, err
	}

	return serviceBindingOperationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServiceBindingOperation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_binding_operations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceBindingOperationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceBindingOperationInsertCacheMut.RLock()
	cache, cached := serviceBindingOperationInsertCache[key]
	serviceBindingOperationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceBindingOperationAllColumns,
			serviceBindingOperationColumnsWithDefault,
			serviceBindingOperationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_binding_operations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_binding_operations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_binding_operations")
	}

	if !cached {
		serviceBindingOperationInsertCacheMut.Lock()
		serviceBindingOperationInsertCache[key] = cache
		serviceBindingOperationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ServiceBindingOperation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServiceBindingOperation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	serviceBindingOperationUpdateCacheMut.RLock()
	cache, cached := serviceBindingOperationUpdateCache[key]
	serviceBindingOperationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceBindingOperationAllColumns,
			serviceBindingOperationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_binding_operations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_binding_operations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceBindingOperationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, append(wl, serviceBindingOperationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_binding_operations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_binding_operations")
	}

	if !cached {
		serviceBindingOperationUpdateCacheMut.Lock()
		serviceBindingOperationUpdateCache[key] = cache
		serviceBindingOperationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q serviceBindingOperationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_binding_operations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServiceBindingOperationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_binding_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceBindingOperationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceBindingOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceBindingOperation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceBindingOperation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_binding_operations provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceBindingOperationColumnsWithDefault, o)

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

	serviceBindingOperationUpsertCacheMut.RLock()
	cache, cached := serviceBindingOperationUpsertCache[key]
	serviceBindingOperationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceBindingOperationAllColumns,
			serviceBindingOperationColumnsWithDefault,
			serviceBindingOperationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceBindingOperationAllColumns,
			serviceBindingOperationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_binding_operations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceBindingOperationPrimaryKeyColumns))
			copy(conflict, serviceBindingOperationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_binding_operations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceBindingOperationType, serviceBindingOperationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_binding_operations")
	}

	if !cached {
		serviceBindingOperationUpsertCacheMut.Lock()
		serviceBindingOperationUpsertCache[key] = cache
		serviceBindingOperationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ServiceBindingOperation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServiceBindingOperation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceBindingOperation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceBindingOperationPrimaryKeyMapping)
	sql := "DELETE FROM \"service_binding_operations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_binding_operations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceBindingOperationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceBindingOperationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_binding_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_binding_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServiceBindingOperationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(serviceBindingOperationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_binding_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceBindingOperationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceBindingOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_binding_operations")
	}

	if len(serviceBindingOperationAfterDeleteHooks) != 0 {
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
func (o *ServiceBindingOperation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceBindingOperation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServiceBindingOperationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceBindingOperationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_binding_operations\".* FROM \"service_binding_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceBindingOperationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceBindingOperationSlice")
	}

	*o = slice

	return nil
}

// ServiceBindingOperationExists checks if the ServiceBindingOperation row exists.
func ServiceBindingOperationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_binding_operations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_binding_operations exists")
	}

	return exists, nil
}