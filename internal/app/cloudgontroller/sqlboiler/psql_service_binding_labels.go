// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/service_binding_labels.go
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

// ServiceBindingLabel is an object representing the database table.
type ServiceBindingLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *serviceBindingLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceBindingLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceBindingLabelColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ResourceGUID: "resource_guid",
	KeyPrefix:    "key_prefix",
	KeyName:      "key_name",
	Value:        "value",
}

var ServiceBindingLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "service_binding_labels.id",
	GUID:         "service_binding_labels.guid",
	CreatedAt:    "service_binding_labels.created_at",
	UpdatedAt:    "service_binding_labels.updated_at",
	ResourceGUID: "service_binding_labels.resource_guid",
	KeyPrefix:    "service_binding_labels.key_prefix",
	KeyName:      "service_binding_labels.key_name",
	Value:        "service_binding_labels.value",
}

// Generated where

var ServiceBindingLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"service_binding_labels\".\"id\""},
	GUID:         whereHelperstring{field: "\"service_binding_labels\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"service_binding_labels\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"service_binding_labels\".\"updated_at\""},
	ResourceGUID: whereHelpernull_String{field: "\"service_binding_labels\".\"resource_guid\""},
	KeyPrefix:    whereHelpernull_String{field: "\"service_binding_labels\".\"key_prefix\""},
	KeyName:      whereHelpernull_String{field: "\"service_binding_labels\".\"key_name\""},
	Value:        whereHelpernull_String{field: "\"service_binding_labels\".\"value\""},
}

// ServiceBindingLabelRels is where relationship names are stored.
var ServiceBindingLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// serviceBindingLabelR is where relationships are stored.
type serviceBindingLabelR struct {
	Resource *ServiceBinding `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*serviceBindingLabelR) NewStruct() *serviceBindingLabelR {
	return &serviceBindingLabelR{}
}

// serviceBindingLabelL is where Load methods for each relationship are stored.
type serviceBindingLabelL struct{}

var (
	serviceBindingLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceBindingLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceBindingLabelColumnsWithDefault    = []string{"id", "created_at"}
	serviceBindingLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceBindingLabelSlice is an alias for a slice of pointers to ServiceBindingLabel.
	// This should almost always be used instead of []ServiceBindingLabel.
	ServiceBindingLabelSlice []*ServiceBindingLabel

	ServiceBindingLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceBindingLabelType                 = reflect.TypeOf(&ServiceBindingLabel{})
	serviceBindingLabelMapping              = queries.MakeStructMapping(serviceBindingLabelType)
	serviceBindingLabelPrimaryKeyMapping, _ = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, serviceBindingLabelPrimaryKeyColumns)
	serviceBindingLabelInsertCacheMut       sync.RWMutex
	serviceBindingLabelInsertCache          = make(map[string]insertCache)
	serviceBindingLabelUpdateCacheMut       sync.RWMutex
	serviceBindingLabelUpdateCache          = make(map[string]updateCache)
	serviceBindingLabelUpsertCacheMut       sync.RWMutex
	serviceBindingLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceBindingLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceBindingLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceBindingLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceBindingLabel record from the query.
func (q ServiceBindingLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceBindingLabel, error) {
	o := &ServiceBindingLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_binding_labels")
	}

	return o, nil
}

// All returns all ServiceBindingLabel records from the query.
func (q ServiceBindingLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceBindingLabelSlice, error) {
	var o []*ServiceBindingLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceBindingLabel slice")
	}

	return o, nil
}

// Count returns the count of all ServiceBindingLabel records in the query.
func (q ServiceBindingLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_binding_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ServiceBindingLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_binding_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q ServiceBindingLabelQuery) Resource(o *ServiceBindingLabel, mods ...qm.QueryMod) ServiceBindingQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceBindings(queryMods...)
	queries.SetFrom(query.Query, "\"service_bindings\"")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceBindingLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceBindingLabel interface{}, mods queries.Applicator) error {
	var slice []*ServiceBindingLabel
	var object *ServiceBindingLabel

	if singular {
		object = maybeServiceBindingLabel.(*ServiceBindingLabel)
	} else {
		slice = *maybeServiceBindingLabel.(*[]*ServiceBindingLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceBindingLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceBindingLabelR{}
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
		qm.From(`service_bindings`),
		qm.WhereIn(`service_bindings.guid in ?`, args...),
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

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &serviceBindingR{}
		}
		foreign.R.ResourceServiceBindingLabels = append(foreign.R.ResourceServiceBindingLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &serviceBindingR{}
				}
				foreign.R.ResourceServiceBindingLabels = append(foreign.R.ResourceServiceBindingLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the serviceBindingLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceServiceBindingLabels.
func (q ServiceBindingLabelQuery) SetResource(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceBinding) error {
	var err error
	if insert {
		if err = ServiceBindings().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"service_binding_labels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
		strmangle.WhereClause("\"", "\"", 2, serviceBindingLabelPrimaryKeyColumns),
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
		o.R = &serviceBindingLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &serviceBindingR{
			ResourceServiceBindingLabels: ServiceBindingLabelSlice{o},
		}
	} else {
		related.R.ResourceServiceBindingLabels = append(related.R.ResourceServiceBindingLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q ServiceBindingLabelQuery) RemoveResource(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, related *ServiceBinding) error {
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

	for i, ri := range related.R.ResourceServiceBindingLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceServiceBindingLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceServiceBindingLabels[i] = related.R.ResourceServiceBindingLabels[ln-1]
		}
		related.R.ResourceServiceBindingLabels = related.R.ResourceServiceBindingLabels[:ln-1]
		break
	}
	return nil
}

// ServiceBindingLabels retrieves all the records using an executor.
func ServiceBindingLabels(mods ...qm.QueryMod) ServiceBindingLabelQuery {
	mods = append(mods, qm.From("\"service_binding_labels\""))
	return ServiceBindingLabelQuery{NewQuery(mods...)}
}

type ServiceBindingLabelFinder interface {
	FindServiceBindingLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceBindingLabel, error)
}

// FindServiceBindingLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceBindingLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceBindingLabel, error) {
	serviceBindingLabelObj := &ServiceBindingLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_binding_labels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceBindingLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_binding_labels")
	}

	return serviceBindingLabelObj, nil
}

type ServiceBindingLabelInserter interface {
	Insert(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q ServiceBindingLabelQuery) Insert(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_binding_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceBindingLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceBindingLabelInsertCacheMut.RLock()
	cache, cached := serviceBindingLabelInsertCache[key]
	serviceBindingLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceBindingLabelAllColumns,
			serviceBindingLabelColumnsWithDefault,
			serviceBindingLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_binding_labels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_binding_labels\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_binding_labels")
	}

	if !cached {
		serviceBindingLabelInsertCacheMut.Lock()
		serviceBindingLabelInsertCache[key] = cache
		serviceBindingLabelInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceBindingLabelUpdater interface {
	Update(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceBindingLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q ServiceBindingLabelQuery) Update(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceBindingLabelUpdateCacheMut.RLock()
	cache, cached := serviceBindingLabelUpdateCache[key]
	serviceBindingLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceBindingLabelAllColumns,
			serviceBindingLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_binding_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_binding_labels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceBindingLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, append(wl, serviceBindingLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_binding_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_binding_labels")
	}

	if !cached {
		serviceBindingLabelUpdateCacheMut.Lock()
		serviceBindingLabelUpdateCache[key] = cache
		serviceBindingLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q ServiceBindingLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_binding_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_binding_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q ServiceBindingLabelQuery) UpdateAllSlice(o ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_binding_labels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceBindingLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceBindingLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceBindingLabel")
	}
	return rowsAff, nil
}

type ServiceBindingLabelDeleter interface {
	Delete(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceBindingLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q ServiceBindingLabelQuery) Delete(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceBindingLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceBindingLabelPrimaryKeyMapping)
	sql := "DELETE FROM \"service_binding_labels\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_binding_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_binding_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ServiceBindingLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceBindingLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_binding_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_binding_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q ServiceBindingLabelQuery) DeleteAllSlice(o ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_binding_labels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceBindingLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceBindingLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_binding_labels")
	}

	return rowsAff, nil
}

type ServiceBindingLabelReloader interface {
	Reload(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q ServiceBindingLabelQuery) Reload(o *ServiceBindingLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceBindingLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q ServiceBindingLabelQuery) ReloadAll(o *ServiceBindingLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceBindingLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBindingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_binding_labels\".* FROM \"service_binding_labels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceBindingLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceBindingLabelSlice")
	}

	*o = slice

	return nil
}

// ServiceBindingLabelExists checks if the ServiceBindingLabel row exists.
func ServiceBindingLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_binding_labels\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_binding_labels exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceBindingLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_binding_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceBindingLabelColumnsWithDefault, o)

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

	serviceBindingLabelUpsertCacheMut.RLock()
	cache, cached := serviceBindingLabelUpsertCache[key]
	serviceBindingLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceBindingLabelAllColumns,
			serviceBindingLabelColumnsWithDefault,
			serviceBindingLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceBindingLabelAllColumns,
			serviceBindingLabelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_binding_labels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceBindingLabelPrimaryKeyColumns))
			copy(conflict, serviceBindingLabelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_binding_labels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceBindingLabelType, serviceBindingLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_binding_labels")
	}

	if !cached {
		serviceBindingLabelUpsertCacheMut.Lock()
		serviceBindingLabelUpsertCache[key] = cache
		serviceBindingLabelUpsertCacheMut.Unlock()
	}

	return nil
}
