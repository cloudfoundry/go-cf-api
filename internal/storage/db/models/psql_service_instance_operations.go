//go:build psql
// +build psql

//go:generate sh -c "echo '\x2bbuild unit' > ../../../../buildtags.txt && mockgen -source=$GOFILE -destination=mocks/service_instance_operations.go -copyright_file=../../../../buildtags.txt && rm -f ../../../../buildtags.txt"
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

type ServiceInstanceOperationUpserter interface {
	Upsert(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q serviceInstanceOperationQuery) Upsert(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_instance_operations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceInstanceOperationColumnsWithDefault, o)

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

	serviceInstanceOperationUpsertCacheMut.RLock()
	cache, cached := serviceInstanceOperationUpsertCache[key]
	serviceInstanceOperationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceInstanceOperationAllColumns,
			serviceInstanceOperationColumnsWithDefault,
			serviceInstanceOperationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceInstanceOperationAllColumns,
			serviceInstanceOperationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_instance_operations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceInstanceOperationPrimaryKeyColumns))
			copy(conflict, serviceInstanceOperationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_instance_operations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_instance_operations")
	}

	if !cached {
		serviceInstanceOperationUpsertCacheMut.Lock()
		serviceInstanceOperationUpsertCache[key] = cache
		serviceInstanceOperationUpsertCacheMut.Unlock()
	}

	return nil
}

// ServiceInstanceOperation is an object representing the database table.
type ServiceInstanceOperation struct {
	ID                      int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                    string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ServiceInstanceID       null.Int    `boil:"service_instance_id" json:"service_instance_id,omitempty" toml:"service_instance_id" yaml:"service_instance_id,omitempty"`
	Type                    null.String `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	State                   null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Description             null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	ProposedChanges         string      `boil:"proposed_changes" json:"proposed_changes" toml:"proposed_changes" yaml:"proposed_changes"`
	BrokerProvidedOperation null.String `boil:"broker_provided_operation" json:"broker_provided_operation,omitempty" toml:"broker_provided_operation" yaml:"broker_provided_operation,omitempty"`

	R *serviceInstanceOperationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceInstanceOperationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceInstanceOperationColumns = struct {
	ID                      string
	GUID                    string
	CreatedAt               string
	UpdatedAt               string
	ServiceInstanceID       string
	Type                    string
	State                   string
	Description             string
	ProposedChanges         string
	BrokerProvidedOperation string
}{
	ID:                      "id",
	GUID:                    "guid",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	ServiceInstanceID:       "service_instance_id",
	Type:                    "type",
	State:                   "state",
	Description:             "description",
	ProposedChanges:         "proposed_changes",
	BrokerProvidedOperation: "broker_provided_operation",
}

var ServiceInstanceOperationTableColumns = struct {
	ID                      string
	GUID                    string
	CreatedAt               string
	UpdatedAt               string
	ServiceInstanceID       string
	Type                    string
	State                   string
	Description             string
	ProposedChanges         string
	BrokerProvidedOperation string
}{
	ID:                      "service_instance_operations.id",
	GUID:                    "service_instance_operations.guid",
	CreatedAt:               "service_instance_operations.created_at",
	UpdatedAt:               "service_instance_operations.updated_at",
	ServiceInstanceID:       "service_instance_operations.service_instance_id",
	Type:                    "service_instance_operations.type",
	State:                   "service_instance_operations.state",
	Description:             "service_instance_operations.description",
	ProposedChanges:         "service_instance_operations.proposed_changes",
	BrokerProvidedOperation: "service_instance_operations.broker_provided_operation",
}

// Generated where

var ServiceInstanceOperationWhere = struct {
	ID                      whereHelperint
	GUID                    whereHelperstring
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpernull_Time
	ServiceInstanceID       whereHelpernull_Int
	Type                    whereHelpernull_String
	State                   whereHelpernull_String
	Description             whereHelpernull_String
	ProposedChanges         whereHelperstring
	BrokerProvidedOperation whereHelpernull_String
}{
	ID:                      whereHelperint{field: "\"service_instance_operations\".\"id\""},
	GUID:                    whereHelperstring{field: "\"service_instance_operations\".\"guid\""},
	CreatedAt:               whereHelpertime_Time{field: "\"service_instance_operations\".\"created_at\""},
	UpdatedAt:               whereHelpernull_Time{field: "\"service_instance_operations\".\"updated_at\""},
	ServiceInstanceID:       whereHelpernull_Int{field: "\"service_instance_operations\".\"service_instance_id\""},
	Type:                    whereHelpernull_String{field: "\"service_instance_operations\".\"type\""},
	State:                   whereHelpernull_String{field: "\"service_instance_operations\".\"state\""},
	Description:             whereHelpernull_String{field: "\"service_instance_operations\".\"description\""},
	ProposedChanges:         whereHelperstring{field: "\"service_instance_operations\".\"proposed_changes\""},
	BrokerProvidedOperation: whereHelpernull_String{field: "\"service_instance_operations\".\"broker_provided_operation\""},
}

// ServiceInstanceOperationRels is where relationship names are stored.
var ServiceInstanceOperationRels = struct {
	ServiceInstance string
}{
	ServiceInstance: "ServiceInstance",
}

// serviceInstanceOperationR is where relationships are stored.
type serviceInstanceOperationR struct {
	ServiceInstance *ServiceInstance `boil:"ServiceInstance" json:"ServiceInstance" toml:"ServiceInstance" yaml:"ServiceInstance"`
}

// NewStruct creates a new relationship struct
func (*serviceInstanceOperationR) NewStruct() *serviceInstanceOperationR {
	return &serviceInstanceOperationR{}
}

// serviceInstanceOperationL is where Load methods for each relationship are stored.
type serviceInstanceOperationL struct{}

var (
	serviceInstanceOperationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "service_instance_id", "type", "state", "description", "proposed_changes", "broker_provided_operation"}
	serviceInstanceOperationColumnsWithoutDefault = []string{"guid", "updated_at", "service_instance_id", "type", "state", "description", "broker_provided_operation"}
	serviceInstanceOperationColumnsWithDefault    = []string{"id", "created_at", "proposed_changes"}
	serviceInstanceOperationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceInstanceOperationSlice is an alias for a slice of pointers to ServiceInstanceOperation.
	// This should almost always be used instead of []ServiceInstanceOperation.
	ServiceInstanceOperationSlice []*ServiceInstanceOperation

	serviceInstanceOperationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceInstanceOperationType                 = reflect.TypeOf(&ServiceInstanceOperation{})
	serviceInstanceOperationMapping              = queries.MakeStructMapping(serviceInstanceOperationType)
	serviceInstanceOperationPrimaryKeyMapping, _ = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, serviceInstanceOperationPrimaryKeyColumns)
	serviceInstanceOperationInsertCacheMut       sync.RWMutex
	serviceInstanceOperationInsertCache          = make(map[string]insertCache)
	serviceInstanceOperationUpdateCacheMut       sync.RWMutex
	serviceInstanceOperationUpdateCache          = make(map[string]updateCache)
	serviceInstanceOperationUpsertCacheMut       sync.RWMutex
	serviceInstanceOperationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceInstanceOperationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceInstanceOperation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceInstanceOperationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceInstanceOperation record from the query.
func (q serviceInstanceOperationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceInstanceOperation, error) {
	o := &ServiceInstanceOperation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_instance_operations")
	}

	return o, nil
}

// All returns all ServiceInstanceOperation records from the query.
func (q serviceInstanceOperationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceInstanceOperationSlice, error) {
	var o []*ServiceInstanceOperation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceInstanceOperation slice")
	}

	return o, nil
}

// Count returns the count of all ServiceInstanceOperation records in the query.
func (q serviceInstanceOperationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_instance_operations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceInstanceOperationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_instance_operations exists")
	}

	return count > 0, nil
}

// ServiceInstance pointed to by the foreign key.
func (q serviceInstanceOperationQuery) ServiceInstance(o *ServiceInstanceOperation, mods ...qm.QueryMod) serviceInstanceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ServiceInstanceID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceInstances(queryMods...)
	queries.SetFrom(query.Query, "\"service_instances\"")

	return query
}

// LoadServiceInstance allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceInstanceOperationL) LoadServiceInstance(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceInstanceOperation interface{}, mods queries.Applicator) error {
	var slice []*ServiceInstanceOperation
	var object *ServiceInstanceOperation

	if singular {
		object = maybeServiceInstanceOperation.(*ServiceInstanceOperation)
	} else {
		slice = *maybeServiceInstanceOperation.(*[]*ServiceInstanceOperation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceInstanceOperationR{}
		}
		if !queries.IsNil(object.ServiceInstanceID) {
			args = append(args, object.ServiceInstanceID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceInstanceOperationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ServiceInstanceID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ServiceInstanceID) {
				args = append(args, obj.ServiceInstanceID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`service_instances`),
		qm.WhereIn(`service_instances.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServiceInstance")
	}

	var resultSlice []*ServiceInstance
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServiceInstance")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for service_instances")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for service_instances")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ServiceInstance = foreign
		if foreign.R == nil {
			foreign.R = &serviceInstanceR{}
		}
		foreign.R.ServiceInstanceOperations = append(foreign.R.ServiceInstanceOperations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ServiceInstanceID, foreign.ID) {
				local.R.ServiceInstance = foreign
				if foreign.R == nil {
					foreign.R = &serviceInstanceR{}
				}
				foreign.R.ServiceInstanceOperations = append(foreign.R.ServiceInstanceOperations, local)
				break
			}
		}
	}

	return nil
}

// SetServiceInstance of the serviceInstanceOperation to the related item.
// Sets o.R.ServiceInstance to related.
// Adds o to related.R.ServiceInstanceOperations.
func (q serviceInstanceOperationQuery) SetServiceInstance(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceInstance) error {
	var err error
	if insert {
		if err = ServiceInstances().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"service_instance_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"service_instance_id"}),
		strmangle.WhereClause("\"", "\"", 2, serviceInstanceOperationPrimaryKeyColumns),
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

	queries.Assign(&o.ServiceInstanceID, related.ID)
	if o.R == nil {
		o.R = &serviceInstanceOperationR{
			ServiceInstance: related,
		}
	} else {
		o.R.ServiceInstance = related
	}

	if related.R == nil {
		related.R = &serviceInstanceR{
			ServiceInstanceOperations: ServiceInstanceOperationSlice{o},
		}
	} else {
		related.R.ServiceInstanceOperations = append(related.R.ServiceInstanceOperations, o)
	}

	return nil
}

// RemoveServiceInstance relationship.
// Sets o.R.ServiceInstance to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q serviceInstanceOperationQuery) RemoveServiceInstance(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, related *ServiceInstance) error {
	var err error

	queries.SetScanner(&o.ServiceInstanceID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("service_instance_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ServiceInstance = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ServiceInstanceOperations {
		if queries.Equal(o.ServiceInstanceID, ri.ServiceInstanceID) {
			continue
		}

		ln := len(related.R.ServiceInstanceOperations)
		if ln > 1 && i < ln-1 {
			related.R.ServiceInstanceOperations[i] = related.R.ServiceInstanceOperations[ln-1]
		}
		related.R.ServiceInstanceOperations = related.R.ServiceInstanceOperations[:ln-1]
		break
	}
	return nil
}

// ServiceInstanceOperations retrieves all the records using an executor.
func ServiceInstanceOperations(mods ...qm.QueryMod) serviceInstanceOperationQuery {
	mods = append(mods, qm.From("\"service_instance_operations\""))
	return serviceInstanceOperationQuery{NewQuery(mods...)}
}

type ServiceInstanceOperationFinder interface {
	FindServiceInstanceOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceInstanceOperation, error)
}

// FindServiceInstanceOperation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceInstanceOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceInstanceOperation, error) {
	serviceInstanceOperationObj := &ServiceInstanceOperation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_instance_operations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceInstanceOperationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_instance_operations")
	}

	return serviceInstanceOperationObj, nil
}

type ServiceInstanceOperationInserter interface {
	Insert(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q serviceInstanceOperationQuery) Insert(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_instance_operations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceInstanceOperationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceInstanceOperationInsertCacheMut.RLock()
	cache, cached := serviceInstanceOperationInsertCache[key]
	serviceInstanceOperationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceInstanceOperationAllColumns,
			serviceInstanceOperationColumnsWithDefault,
			serviceInstanceOperationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_instance_operations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_instance_operations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_instance_operations")
	}

	if !cached {
		serviceInstanceOperationInsertCacheMut.Lock()
		serviceInstanceOperationInsertCache[key] = cache
		serviceInstanceOperationInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceInstanceOperationUpdater interface {
	Update(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceInstanceOperation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q serviceInstanceOperationQuery) Update(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceInstanceOperationUpdateCacheMut.RLock()
	cache, cached := serviceInstanceOperationUpdateCache[key]
	serviceInstanceOperationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceInstanceOperationAllColumns,
			serviceInstanceOperationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_instance_operations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_instance_operations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceInstanceOperationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceInstanceOperationType, serviceInstanceOperationMapping, append(wl, serviceInstanceOperationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_instance_operations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_instance_operations")
	}

	if !cached {
		serviceInstanceOperationUpdateCacheMut.Lock()
		serviceInstanceOperationUpdateCache[key] = cache
		serviceInstanceOperationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serviceInstanceOperationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_instance_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_instance_operations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q serviceInstanceOperationQuery) UpdateAllSlice(o ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceInstanceOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_instance_operations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceInstanceOperationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceInstanceOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceInstanceOperation")
	}
	return rowsAff, nil
}

type ServiceInstanceOperationDeleter interface {
	Delete(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceInstanceOperation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q serviceInstanceOperationQuery) Delete(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceInstanceOperation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceInstanceOperationPrimaryKeyMapping)
	sql := "DELETE FROM \"service_instance_operations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_instance_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_instance_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceInstanceOperationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceInstanceOperationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_instance_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_instance_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q serviceInstanceOperationQuery) DeleteAllSlice(o ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceInstanceOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_instance_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceInstanceOperationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceInstanceOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_instance_operations")
	}

	return rowsAff, nil
}

type ServiceInstanceOperationReloader interface {
	Reload(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q serviceInstanceOperationQuery) Reload(o *ServiceInstanceOperation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceInstanceOperation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q serviceInstanceOperationQuery) ReloadAll(o *ServiceInstanceOperationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceInstanceOperationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceInstanceOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_instance_operations\".* FROM \"service_instance_operations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceInstanceOperationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceInstanceOperationSlice")
	}

	*o = slice

	return nil
}

// ServiceInstanceOperationExists checks if the ServiceInstanceOperation row exists.
func ServiceInstanceOperationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_instance_operations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_instance_operations exists")
	}

	return exists, nil
}
