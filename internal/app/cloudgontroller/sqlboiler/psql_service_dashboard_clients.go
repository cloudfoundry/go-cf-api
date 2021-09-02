//go:generate mockgen -source=$GOFILE -destination=mocks/service_dashboard_clients.go
// +build psql
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

// ServiceDashboardClient is an object representing the database table.
type ServiceDashboardClient struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt       time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UaaID           string    `boil:"uaa_id" json:"uaa_id" toml:"uaa_id" yaml:"uaa_id"`
	ServiceBrokerID null.Int  `boil:"service_broker_id" json:"service_broker_id,omitempty" toml:"service_broker_id" yaml:"service_broker_id,omitempty"`

	R *serviceDashboardClientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceDashboardClientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceDashboardClientColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	UaaID           string
	ServiceBrokerID string
}{
	ID:              "id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	UaaID:           "uaa_id",
	ServiceBrokerID: "service_broker_id",
}

var ServiceDashboardClientTableColumns = struct {
	ID              string
	CreatedAt       string
	UpdatedAt       string
	UaaID           string
	ServiceBrokerID string
}{
	ID:              "service_dashboard_clients.id",
	CreatedAt:       "service_dashboard_clients.created_at",
	UpdatedAt:       "service_dashboard_clients.updated_at",
	UaaID:           "service_dashboard_clients.uaa_id",
	ServiceBrokerID: "service_dashboard_clients.service_broker_id",
}

// Generated where

var ServiceDashboardClientWhere = struct {
	ID              whereHelperint
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpernull_Time
	UaaID           whereHelperstring
	ServiceBrokerID whereHelpernull_Int
}{
	ID:              whereHelperint{field: "\"service_dashboard_clients\".\"id\""},
	CreatedAt:       whereHelpertime_Time{field: "\"service_dashboard_clients\".\"created_at\""},
	UpdatedAt:       whereHelpernull_Time{field: "\"service_dashboard_clients\".\"updated_at\""},
	UaaID:           whereHelperstring{field: "\"service_dashboard_clients\".\"uaa_id\""},
	ServiceBrokerID: whereHelpernull_Int{field: "\"service_dashboard_clients\".\"service_broker_id\""},
}

// ServiceDashboardClientRels is where relationship names are stored.
var ServiceDashboardClientRels = struct {
}{}

// serviceDashboardClientR is where relationships are stored.
type serviceDashboardClientR struct {
}

// NewStruct creates a new relationship struct
func (*serviceDashboardClientR) NewStruct() *serviceDashboardClientR {
	return &serviceDashboardClientR{}
}

// serviceDashboardClientL is where Load methods for each relationship are stored.
type serviceDashboardClientL struct{}

var (
	serviceDashboardClientAllColumns            = []string{"id", "created_at", "updated_at", "uaa_id", "service_broker_id"}
	serviceDashboardClientColumnsWithoutDefault = []string{"updated_at", "uaa_id", "service_broker_id"}
	serviceDashboardClientColumnsWithDefault    = []string{"id", "created_at"}
	serviceDashboardClientPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceDashboardClientSlice is an alias for a slice of pointers to ServiceDashboardClient.
	// This should almost always be used instead of []ServiceDashboardClient.
	ServiceDashboardClientSlice []*ServiceDashboardClient

	ServiceDashboardClientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceDashboardClientType                 = reflect.TypeOf(&ServiceDashboardClient{})
	serviceDashboardClientMapping              = queries.MakeStructMapping(serviceDashboardClientType)
	serviceDashboardClientPrimaryKeyMapping, _ = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, serviceDashboardClientPrimaryKeyColumns)
	serviceDashboardClientInsertCacheMut       sync.RWMutex
	serviceDashboardClientInsertCache          = make(map[string]insertCache)
	serviceDashboardClientUpdateCacheMut       sync.RWMutex
	serviceDashboardClientUpdateCache          = make(map[string]updateCache)
	serviceDashboardClientUpsertCacheMut       sync.RWMutex
	serviceDashboardClientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceDashboardClientFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceDashboardClient, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceDashboardClientSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceDashboardClient record from the query.
func (q ServiceDashboardClientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceDashboardClient, error) {
	o := &ServiceDashboardClient{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_dashboard_clients")
	}

	return o, nil
}

// All returns all ServiceDashboardClient records from the query.
func (q ServiceDashboardClientQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceDashboardClientSlice, error) {
	var o []*ServiceDashboardClient

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceDashboardClient slice")
	}

	return o, nil
}

// Count returns the count of all ServiceDashboardClient records in the query.
func (q ServiceDashboardClientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_dashboard_clients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ServiceDashboardClientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_dashboard_clients exists")
	}

	return count > 0, nil
}

// ServiceDashboardClients retrieves all the records using an executor.
func ServiceDashboardClients(mods ...qm.QueryMod) ServiceDashboardClientQuery {
	mods = append(mods, qm.From("\"service_dashboard_clients\""))
	return ServiceDashboardClientQuery{NewQuery(mods...)}
}

type ServiceDashboardClientFinder interface {
	FindServiceDashboardClient(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceDashboardClient, error)
}

// FindServiceDashboardClient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceDashboardClient(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceDashboardClient, error) {
	serviceDashboardClientObj := &ServiceDashboardClient{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_dashboard_clients\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceDashboardClientObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_dashboard_clients")
	}

	return serviceDashboardClientObj, nil
}

type ServiceDashboardClientInserter interface {
	Insert(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q ServiceDashboardClientQuery) Insert(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_dashboard_clients provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceDashboardClientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceDashboardClientInsertCacheMut.RLock()
	cache, cached := serviceDashboardClientInsertCache[key]
	serviceDashboardClientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceDashboardClientAllColumns,
			serviceDashboardClientColumnsWithDefault,
			serviceDashboardClientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_dashboard_clients\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_dashboard_clients\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_dashboard_clients")
	}

	if !cached {
		serviceDashboardClientInsertCacheMut.Lock()
		serviceDashboardClientInsertCache[key] = cache
		serviceDashboardClientInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceDashboardClientUpdater interface {
	Update(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceDashboardClient.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q ServiceDashboardClientQuery) Update(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceDashboardClientUpdateCacheMut.RLock()
	cache, cached := serviceDashboardClientUpdateCache[key]
	serviceDashboardClientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceDashboardClientAllColumns,
			serviceDashboardClientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_dashboard_clients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_dashboard_clients\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceDashboardClientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, append(wl, serviceDashboardClientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_dashboard_clients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_dashboard_clients")
	}

	if !cached {
		serviceDashboardClientUpdateCacheMut.Lock()
		serviceDashboardClientUpdateCache[key] = cache
		serviceDashboardClientUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q ServiceDashboardClientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_dashboard_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_dashboard_clients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q ServiceDashboardClientQuery) UpdateAllSlice(o ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceDashboardClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_dashboard_clients\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceDashboardClientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceDashboardClient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceDashboardClient")
	}
	return rowsAff, nil
}

type ServiceDashboardClientDeleter interface {
	Delete(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceDashboardClient record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q ServiceDashboardClientQuery) Delete(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceDashboardClient provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceDashboardClientPrimaryKeyMapping)
	sql := "DELETE FROM \"service_dashboard_clients\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_dashboard_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_dashboard_clients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ServiceDashboardClientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceDashboardClientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_dashboard_clients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_dashboard_clients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q ServiceDashboardClientQuery) DeleteAllSlice(o ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceDashboardClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_dashboard_clients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceDashboardClientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceDashboardClient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_dashboard_clients")
	}

	return rowsAff, nil
}

type ServiceDashboardClientReloader interface {
	Reload(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q ServiceDashboardClientQuery) Reload(o *ServiceDashboardClient, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceDashboardClient(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q ServiceDashboardClientQuery) ReloadAll(o *ServiceDashboardClientSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceDashboardClientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceDashboardClientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_dashboard_clients\".* FROM \"service_dashboard_clients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceDashboardClientPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceDashboardClientSlice")
	}

	*o = slice

	return nil
}

// ServiceDashboardClientExists checks if the ServiceDashboardClient row exists.
func ServiceDashboardClientExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_dashboard_clients\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_dashboard_clients exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceDashboardClient) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_dashboard_clients provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceDashboardClientColumnsWithDefault, o)

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

	serviceDashboardClientUpsertCacheMut.RLock()
	cache, cached := serviceDashboardClientUpsertCache[key]
	serviceDashboardClientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceDashboardClientAllColumns,
			serviceDashboardClientColumnsWithDefault,
			serviceDashboardClientColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceDashboardClientAllColumns,
			serviceDashboardClientPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_dashboard_clients, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceDashboardClientPrimaryKeyColumns))
			copy(conflict, serviceDashboardClientPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_dashboard_clients\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceDashboardClientType, serviceDashboardClientMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_dashboard_clients")
	}

	if !cached {
		serviceDashboardClientUpsertCacheMut.Lock()
		serviceDashboardClientUpsertCache[key] = cache
		serviceDashboardClientUpsertCacheMut.Unlock()
	}

	return nil
}
