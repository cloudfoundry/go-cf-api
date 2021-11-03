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

type ServiceKeyOperationUpserter interface {
	Upsert(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLServiceKeyOperationUniqueColumns = []string{
	"id",
	"service_key_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q serviceKeyOperationQuery) Upsert(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_key_operations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceKeyOperationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLServiceKeyOperationUniqueColumns, o)

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

	serviceKeyOperationUpsertCacheMut.RLock()
	cache, cached := serviceKeyOperationUpsertCache[key]
	serviceKeyOperationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceKeyOperationAllColumns,
			serviceKeyOperationColumnsWithDefault,
			serviceKeyOperationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceKeyOperationAllColumns,
			serviceKeyOperationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert service_key_operations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`service_key_operations`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `service_key_operations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for service_key_operations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceKeyOperationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for service_key_operations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for service_key_operations")
	}

CacheNoHooks:
	if !cached {
		serviceKeyOperationUpsertCacheMut.Lock()
		serviceKeyOperationUpsertCache[key] = cache
		serviceKeyOperationUpsertCacheMut.Unlock()
	}

	return nil
}

// ServiceKeyOperation is an object representing the database table.
type ServiceKeyOperation struct {
	ID                      int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ServiceKeyID            null.Int    `boil:"service_key_id" json:"service_key_id,omitempty" toml:"service_key_id" yaml:"service_key_id,omitempty"`
	State                   string      `boil:"state" json:"state" toml:"state" yaml:"state"`
	Type                    string      `boil:"type" json:"type" toml:"type" yaml:"type"`
	Description             null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	BrokerProvidedOperation null.String `boil:"broker_provided_operation" json:"broker_provided_operation,omitempty" toml:"broker_provided_operation" yaml:"broker_provided_operation,omitempty"`

	R *serviceKeyOperationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceKeyOperationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceKeyOperationColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	ServiceKeyID            string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "id",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	ServiceKeyID:            "service_key_id",
	State:                   "state",
	Type:                    "type",
	Description:             "description",
	BrokerProvidedOperation: "broker_provided_operation",
}

var ServiceKeyOperationTableColumns = struct {
	ID                      string
	CreatedAt               string
	UpdatedAt               string
	ServiceKeyID            string
	State                   string
	Type                    string
	Description             string
	BrokerProvidedOperation string
}{
	ID:                      "service_key_operations.id",
	CreatedAt:               "service_key_operations.created_at",
	UpdatedAt:               "service_key_operations.updated_at",
	ServiceKeyID:            "service_key_operations.service_key_id",
	State:                   "service_key_operations.state",
	Type:                    "service_key_operations.type",
	Description:             "service_key_operations.description",
	BrokerProvidedOperation: "service_key_operations.broker_provided_operation",
}

// Generated where

var ServiceKeyOperationWhere = struct {
	ID                      whereHelperint
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpernull_Time
	ServiceKeyID            whereHelpernull_Int
	State                   whereHelperstring
	Type                    whereHelperstring
	Description             whereHelpernull_String
	BrokerProvidedOperation whereHelpernull_String
}{
	ID:                      whereHelperint{field: "`service_key_operations`.`id`"},
	CreatedAt:               whereHelpertime_Time{field: "`service_key_operations`.`created_at`"},
	UpdatedAt:               whereHelpernull_Time{field: "`service_key_operations`.`updated_at`"},
	ServiceKeyID:            whereHelpernull_Int{field: "`service_key_operations`.`service_key_id`"},
	State:                   whereHelperstring{field: "`service_key_operations`.`state`"},
	Type:                    whereHelperstring{field: "`service_key_operations`.`type`"},
	Description:             whereHelpernull_String{field: "`service_key_operations`.`description`"},
	BrokerProvidedOperation: whereHelpernull_String{field: "`service_key_operations`.`broker_provided_operation`"},
}

// ServiceKeyOperationRels is where relationship names are stored.
var ServiceKeyOperationRels = struct {
	ServiceKey string
}{
	ServiceKey: "ServiceKey",
}

// serviceKeyOperationR is where relationships are stored.
type serviceKeyOperationR struct {
	ServiceKey *ServiceKey `boil:"ServiceKey" json:"ServiceKey" toml:"ServiceKey" yaml:"ServiceKey"`
}

// NewStruct creates a new relationship struct
func (*serviceKeyOperationR) NewStruct() *serviceKeyOperationR {
	return &serviceKeyOperationR{}
}

// serviceKeyOperationL is where Load methods for each relationship are stored.
type serviceKeyOperationL struct{}

var (
	serviceKeyOperationAllColumns            = []string{"id", "created_at", "updated_at", "service_key_id", "state", "type", "description", "broker_provided_operation"}
	serviceKeyOperationColumnsWithoutDefault = []string{"updated_at", "service_key_id", "state", "type", "description", "broker_provided_operation"}
	serviceKeyOperationColumnsWithDefault    = []string{"id", "created_at"}
	serviceKeyOperationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceKeyOperationSlice is an alias for a slice of pointers to ServiceKeyOperation.
	// This should almost always be used instead of []ServiceKeyOperation.
	ServiceKeyOperationSlice []*ServiceKeyOperation

	serviceKeyOperationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceKeyOperationType                 = reflect.TypeOf(&ServiceKeyOperation{})
	serviceKeyOperationMapping              = queries.MakeStructMapping(serviceKeyOperationType)
	serviceKeyOperationPrimaryKeyMapping, _ = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, serviceKeyOperationPrimaryKeyColumns)
	serviceKeyOperationInsertCacheMut       sync.RWMutex
	serviceKeyOperationInsertCache          = make(map[string]insertCache)
	serviceKeyOperationUpdateCacheMut       sync.RWMutex
	serviceKeyOperationUpdateCache          = make(map[string]updateCache)
	serviceKeyOperationUpsertCacheMut       sync.RWMutex
	serviceKeyOperationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceKeyOperationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceKeyOperation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceKeyOperationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceKeyOperation record from the query.
func (q serviceKeyOperationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceKeyOperation, error) {
	o := &ServiceKeyOperation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_key_operations")
	}

	return o, nil
}

// All returns all ServiceKeyOperation records from the query.
func (q serviceKeyOperationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceKeyOperationSlice, error) {
	var o []*ServiceKeyOperation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceKeyOperation slice")
	}

	return o, nil
}

// Count returns the count of all ServiceKeyOperation records in the query.
func (q serviceKeyOperationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_key_operations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceKeyOperationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_key_operations exists")
	}

	return count > 0, nil
}

// ServiceKey pointed to by the foreign key.
func (q serviceKeyOperationQuery) ServiceKey(o *ServiceKeyOperation, mods ...qm.QueryMod) serviceKeyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.ServiceKeyID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceKeys(queryMods...)
	queries.SetFrom(query.Query, "`service_keys`")

	return query
}

// LoadServiceKey allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceKeyOperationL) LoadServiceKey(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceKeyOperation interface{}, mods queries.Applicator) error {
	var slice []*ServiceKeyOperation
	var object *ServiceKeyOperation

	if singular {
		object = maybeServiceKeyOperation.(*ServiceKeyOperation)
	} else {
		slice = *maybeServiceKeyOperation.(*[]*ServiceKeyOperation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceKeyOperationR{}
		}
		if !queries.IsNil(object.ServiceKeyID) {
			args = append(args, object.ServiceKeyID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceKeyOperationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ServiceKeyID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ServiceKeyID) {
				args = append(args, obj.ServiceKeyID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`service_keys`),
		qm.WhereIn(`service_keys.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServiceKey")
	}

	var resultSlice []*ServiceKey
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServiceKey")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for service_keys")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for service_keys")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ServiceKey = foreign
		if foreign.R == nil {
			foreign.R = &serviceKeyR{}
		}
		foreign.R.ServiceKeyOperation = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ServiceKeyID, foreign.ID) {
				local.R.ServiceKey = foreign
				if foreign.R == nil {
					foreign.R = &serviceKeyR{}
				}
				foreign.R.ServiceKeyOperation = local
				break
			}
		}
	}

	return nil
}

// SetServiceKey of the serviceKeyOperation to the related item.
// Sets o.R.ServiceKey to related.
// Adds o to related.R.ServiceKeyOperation.
func (q serviceKeyOperationQuery) SetServiceKey(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceKey) error {
	var err error
	if insert {
		if err = ServiceKeys().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `service_key_operations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"service_key_id"}),
		strmangle.WhereClause("`", "`", 0, serviceKeyOperationPrimaryKeyColumns),
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

	queries.Assign(&o.ServiceKeyID, related.ID)
	if o.R == nil {
		o.R = &serviceKeyOperationR{
			ServiceKey: related,
		}
	} else {
		o.R.ServiceKey = related
	}

	if related.R == nil {
		related.R = &serviceKeyR{
			ServiceKeyOperation: o,
		}
	} else {
		related.R.ServiceKeyOperation = o
	}

	return nil
}

// RemoveServiceKey relationship.
// Sets o.R.ServiceKey to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q serviceKeyOperationQuery) RemoveServiceKey(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, related *ServiceKey) error {
	var err error

	queries.SetScanner(&o.ServiceKeyID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("service_key_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.ServiceKey = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.ServiceKeyOperation = nil
	return nil
}

// ServiceKeyOperations retrieves all the records using an executor.
func ServiceKeyOperations(mods ...qm.QueryMod) serviceKeyOperationQuery {
	mods = append(mods, qm.From("`service_key_operations`"))
	return serviceKeyOperationQuery{NewQuery(mods...)}
}

type ServiceKeyOperationFinder interface {
	FindServiceKeyOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceKeyOperation, error)
}

// FindServiceKeyOperation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceKeyOperation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceKeyOperation, error) {
	serviceKeyOperationObj := &ServiceKeyOperation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `service_key_operations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceKeyOperationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_key_operations")
	}

	return serviceKeyOperationObj, nil
}

type ServiceKeyOperationInserter interface {
	Insert(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q serviceKeyOperationQuery) Insert(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_key_operations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceKeyOperationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceKeyOperationInsertCacheMut.RLock()
	cache, cached := serviceKeyOperationInsertCache[key]
	serviceKeyOperationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceKeyOperationAllColumns,
			serviceKeyOperationColumnsWithDefault,
			serviceKeyOperationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `service_key_operations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `service_key_operations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `service_key_operations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, serviceKeyOperationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into service_key_operations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceKeyOperationMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for service_key_operations")
	}

CacheNoHooks:
	if !cached {
		serviceKeyOperationInsertCacheMut.Lock()
		serviceKeyOperationInsertCache[key] = cache
		serviceKeyOperationInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceKeyOperationUpdater interface {
	Update(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceKeyOperation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q serviceKeyOperationQuery) Update(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceKeyOperationUpdateCacheMut.RLock()
	cache, cached := serviceKeyOperationUpdateCache[key]
	serviceKeyOperationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceKeyOperationAllColumns,
			serviceKeyOperationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_key_operations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `service_key_operations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, serviceKeyOperationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceKeyOperationType, serviceKeyOperationMapping, append(wl, serviceKeyOperationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_key_operations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_key_operations")
	}

	if !cached {
		serviceKeyOperationUpdateCacheMut.Lock()
		serviceKeyOperationUpdateCache[key] = cache
		serviceKeyOperationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serviceKeyOperationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_key_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_key_operations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q serviceKeyOperationQuery) UpdateAllSlice(o ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `service_key_operations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceKeyOperationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceKeyOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceKeyOperation")
	}
	return rowsAff, nil
}

type ServiceKeyOperationDeleter interface {
	Delete(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceKeyOperation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q serviceKeyOperationQuery) Delete(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceKeyOperation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceKeyOperationPrimaryKeyMapping)
	sql := "DELETE FROM `service_key_operations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_key_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_key_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceKeyOperationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceKeyOperationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_key_operations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_key_operations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q serviceKeyOperationQuery) DeleteAllSlice(o ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `service_key_operations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceKeyOperationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceKeyOperation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_key_operations")
	}

	return rowsAff, nil
}

type ServiceKeyOperationReloader interface {
	Reload(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q serviceKeyOperationQuery) Reload(o *ServiceKeyOperation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceKeyOperation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q serviceKeyOperationQuery) ReloadAll(o *ServiceKeyOperationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceKeyOperationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyOperationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `service_key_operations`.* FROM `service_key_operations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceKeyOperationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceKeyOperationSlice")
	}

	*o = slice

	return nil
}

// ServiceKeyOperationExists checks if the ServiceKeyOperation row exists.
func ServiceKeyOperationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `service_key_operations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_key_operations exists")
	}

	return exists, nil
}
