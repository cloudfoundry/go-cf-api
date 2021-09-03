// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/service_usage_events.go
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

// ServiceUsageEvent is an object representing the database table.
type ServiceUsageEvent struct {
	ID                  int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt           time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	State               string      `boil:"state" json:"state" toml:"state" yaml:"state"`
	OrgGUID             string      `boil:"org_guid" json:"org_guid" toml:"org_guid" yaml:"org_guid"`
	SpaceGUID           string      `boil:"space_guid" json:"space_guid" toml:"space_guid" yaml:"space_guid"`
	SpaceName           string      `boil:"space_name" json:"space_name" toml:"space_name" yaml:"space_name"`
	ServiceInstanceGUID string      `boil:"service_instance_guid" json:"service_instance_guid" toml:"service_instance_guid" yaml:"service_instance_guid"`
	ServiceInstanceName string      `boil:"service_instance_name" json:"service_instance_name" toml:"service_instance_name" yaml:"service_instance_name"`
	ServiceInstanceType string      `boil:"service_instance_type" json:"service_instance_type" toml:"service_instance_type" yaml:"service_instance_type"`
	ServicePlanGUID     null.String `boil:"service_plan_guid" json:"service_plan_guid,omitempty" toml:"service_plan_guid" yaml:"service_plan_guid,omitempty"`
	ServicePlanName     null.String `boil:"service_plan_name" json:"service_plan_name,omitempty" toml:"service_plan_name" yaml:"service_plan_name,omitempty"`
	ServiceGUID         null.String `boil:"service_guid" json:"service_guid,omitempty" toml:"service_guid" yaml:"service_guid,omitempty"`
	ServiceLabel        null.String `boil:"service_label" json:"service_label,omitempty" toml:"service_label" yaml:"service_label,omitempty"`
	ServiceBrokerName   null.String `boil:"service_broker_name" json:"service_broker_name,omitempty" toml:"service_broker_name" yaml:"service_broker_name,omitempty"`
	ServiceBrokerGUID   null.String `boil:"service_broker_guid" json:"service_broker_guid,omitempty" toml:"service_broker_guid" yaml:"service_broker_guid,omitempty"`

	R *serviceUsageEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceUsageEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceUsageEventColumns = struct {
	ID                  string
	GUID                string
	CreatedAt           string
	State               string
	OrgGUID             string
	SpaceGUID           string
	SpaceName           string
	ServiceInstanceGUID string
	ServiceInstanceName string
	ServiceInstanceType string
	ServicePlanGUID     string
	ServicePlanName     string
	ServiceGUID         string
	ServiceLabel        string
	ServiceBrokerName   string
	ServiceBrokerGUID   string
}{
	ID:                  "id",
	GUID:                "guid",
	CreatedAt:           "created_at",
	State:               "state",
	OrgGUID:             "org_guid",
	SpaceGUID:           "space_guid",
	SpaceName:           "space_name",
	ServiceInstanceGUID: "service_instance_guid",
	ServiceInstanceName: "service_instance_name",
	ServiceInstanceType: "service_instance_type",
	ServicePlanGUID:     "service_plan_guid",
	ServicePlanName:     "service_plan_name",
	ServiceGUID:         "service_guid",
	ServiceLabel:        "service_label",
	ServiceBrokerName:   "service_broker_name",
	ServiceBrokerGUID:   "service_broker_guid",
}

var ServiceUsageEventTableColumns = struct {
	ID                  string
	GUID                string
	CreatedAt           string
	State               string
	OrgGUID             string
	SpaceGUID           string
	SpaceName           string
	ServiceInstanceGUID string
	ServiceInstanceName string
	ServiceInstanceType string
	ServicePlanGUID     string
	ServicePlanName     string
	ServiceGUID         string
	ServiceLabel        string
	ServiceBrokerName   string
	ServiceBrokerGUID   string
}{
	ID:                  "service_usage_events.id",
	GUID:                "service_usage_events.guid",
	CreatedAt:           "service_usage_events.created_at",
	State:               "service_usage_events.state",
	OrgGUID:             "service_usage_events.org_guid",
	SpaceGUID:           "service_usage_events.space_guid",
	SpaceName:           "service_usage_events.space_name",
	ServiceInstanceGUID: "service_usage_events.service_instance_guid",
	ServiceInstanceName: "service_usage_events.service_instance_name",
	ServiceInstanceType: "service_usage_events.service_instance_type",
	ServicePlanGUID:     "service_usage_events.service_plan_guid",
	ServicePlanName:     "service_usage_events.service_plan_name",
	ServiceGUID:         "service_usage_events.service_guid",
	ServiceLabel:        "service_usage_events.service_label",
	ServiceBrokerName:   "service_usage_events.service_broker_name",
	ServiceBrokerGUID:   "service_usage_events.service_broker_guid",
}

// Generated where

var ServiceUsageEventWhere = struct {
	ID                  whereHelperint
	GUID                whereHelperstring
	CreatedAt           whereHelpertime_Time
	State               whereHelperstring
	OrgGUID             whereHelperstring
	SpaceGUID           whereHelperstring
	SpaceName           whereHelperstring
	ServiceInstanceGUID whereHelperstring
	ServiceInstanceName whereHelperstring
	ServiceInstanceType whereHelperstring
	ServicePlanGUID     whereHelpernull_String
	ServicePlanName     whereHelpernull_String
	ServiceGUID         whereHelpernull_String
	ServiceLabel        whereHelpernull_String
	ServiceBrokerName   whereHelpernull_String
	ServiceBrokerGUID   whereHelpernull_String
}{
	ID:                  whereHelperint{field: "\"service_usage_events\".\"id\""},
	GUID:                whereHelperstring{field: "\"service_usage_events\".\"guid\""},
	CreatedAt:           whereHelpertime_Time{field: "\"service_usage_events\".\"created_at\""},
	State:               whereHelperstring{field: "\"service_usage_events\".\"state\""},
	OrgGUID:             whereHelperstring{field: "\"service_usage_events\".\"org_guid\""},
	SpaceGUID:           whereHelperstring{field: "\"service_usage_events\".\"space_guid\""},
	SpaceName:           whereHelperstring{field: "\"service_usage_events\".\"space_name\""},
	ServiceInstanceGUID: whereHelperstring{field: "\"service_usage_events\".\"service_instance_guid\""},
	ServiceInstanceName: whereHelperstring{field: "\"service_usage_events\".\"service_instance_name\""},
	ServiceInstanceType: whereHelperstring{field: "\"service_usage_events\".\"service_instance_type\""},
	ServicePlanGUID:     whereHelpernull_String{field: "\"service_usage_events\".\"service_plan_guid\""},
	ServicePlanName:     whereHelpernull_String{field: "\"service_usage_events\".\"service_plan_name\""},
	ServiceGUID:         whereHelpernull_String{field: "\"service_usage_events\".\"service_guid\""},
	ServiceLabel:        whereHelpernull_String{field: "\"service_usage_events\".\"service_label\""},
	ServiceBrokerName:   whereHelpernull_String{field: "\"service_usage_events\".\"service_broker_name\""},
	ServiceBrokerGUID:   whereHelpernull_String{field: "\"service_usage_events\".\"service_broker_guid\""},
}

// ServiceUsageEventRels is where relationship names are stored.
var ServiceUsageEventRels = struct {
}{}

// serviceUsageEventR is where relationships are stored.
type serviceUsageEventR struct {
}

// NewStruct creates a new relationship struct
func (*serviceUsageEventR) NewStruct() *serviceUsageEventR {
	return &serviceUsageEventR{}
}

// serviceUsageEventL is where Load methods for each relationship are stored.
type serviceUsageEventL struct{}

var (
	serviceUsageEventAllColumns            = []string{"id", "guid", "created_at", "state", "org_guid", "space_guid", "space_name", "service_instance_guid", "service_instance_name", "service_instance_type", "service_plan_guid", "service_plan_name", "service_guid", "service_label", "service_broker_name", "service_broker_guid"}
	serviceUsageEventColumnsWithoutDefault = []string{"guid", "created_at", "state", "org_guid", "space_guid", "space_name", "service_instance_guid", "service_instance_name", "service_instance_type", "service_plan_guid", "service_plan_name", "service_guid", "service_label", "service_broker_name", "service_broker_guid"}
	serviceUsageEventColumnsWithDefault    = []string{"id"}
	serviceUsageEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceUsageEventSlice is an alias for a slice of pointers to ServiceUsageEvent.
	// This should almost always be used instead of []ServiceUsageEvent.
	ServiceUsageEventSlice []*ServiceUsageEvent

	serviceUsageEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceUsageEventType                 = reflect.TypeOf(&ServiceUsageEvent{})
	serviceUsageEventMapping              = queries.MakeStructMapping(serviceUsageEventType)
	serviceUsageEventPrimaryKeyMapping, _ = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, serviceUsageEventPrimaryKeyColumns)
	serviceUsageEventInsertCacheMut       sync.RWMutex
	serviceUsageEventInsertCache          = make(map[string]insertCache)
	serviceUsageEventUpdateCacheMut       sync.RWMutex
	serviceUsageEventUpdateCache          = make(map[string]updateCache)
	serviceUsageEventUpsertCacheMut       sync.RWMutex
	serviceUsageEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceUsageEventFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceUsageEvent, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceUsageEventSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceUsageEvent record from the query.
func (q serviceUsageEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceUsageEvent, error) {
	o := &ServiceUsageEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_usage_events")
	}

	return o, nil
}

// All returns all ServiceUsageEvent records from the query.
func (q serviceUsageEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceUsageEventSlice, error) {
	var o []*ServiceUsageEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceUsageEvent slice")
	}

	return o, nil
}

// Count returns the count of all ServiceUsageEvent records in the query.
func (q serviceUsageEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_usage_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceUsageEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_usage_events exists")
	}

	return count > 0, nil
}

// ServiceUsageEvents retrieves all the records using an executor.
func ServiceUsageEvents(mods ...qm.QueryMod) serviceUsageEventQuery {
	mods = append(mods, qm.From("\"service_usage_events\""))
	return serviceUsageEventQuery{NewQuery(mods...)}
}

type ServiceUsageEventFinder interface {
	FindServiceUsageEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceUsageEvent, error)
}

// FindServiceUsageEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceUsageEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceUsageEvent, error) {
	serviceUsageEventObj := &ServiceUsageEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_usage_events\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceUsageEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_usage_events")
	}

	return serviceUsageEventObj, nil
}

type ServiceUsageEventInserter interface {
	Insert(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q serviceUsageEventQuery) Insert(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_usage_events provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceUsageEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceUsageEventInsertCacheMut.RLock()
	cache, cached := serviceUsageEventInsertCache[key]
	serviceUsageEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceUsageEventAllColumns,
			serviceUsageEventColumnsWithDefault,
			serviceUsageEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_usage_events\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_usage_events\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_usage_events")
	}

	if !cached {
		serviceUsageEventInsertCacheMut.Lock()
		serviceUsageEventInsertCache[key] = cache
		serviceUsageEventInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceUsageEventUpdater interface {
	Update(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceUsageEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q serviceUsageEventQuery) Update(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	serviceUsageEventUpdateCacheMut.RLock()
	cache, cached := serviceUsageEventUpdateCache[key]
	serviceUsageEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceUsageEventAllColumns,
			serviceUsageEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_usage_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_usage_events\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceUsageEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, append(wl, serviceUsageEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_usage_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_usage_events")
	}

	if !cached {
		serviceUsageEventUpdateCacheMut.Lock()
		serviceUsageEventUpdateCache[key] = cache
		serviceUsageEventUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serviceUsageEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_usage_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_usage_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q serviceUsageEventQuery) UpdateAllSlice(o ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceUsageEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_usage_events\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceUsageEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceUsageEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceUsageEvent")
	}
	return rowsAff, nil
}

type ServiceUsageEventUpserter interface {
	Upsert(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q serviceUsageEventQuery) Upsert(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_usage_events provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceUsageEventColumnsWithDefault, o)

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

	serviceUsageEventUpsertCacheMut.RLock()
	cache, cached := serviceUsageEventUpsertCache[key]
	serviceUsageEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceUsageEventAllColumns,
			serviceUsageEventColumnsWithDefault,
			serviceUsageEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceUsageEventAllColumns,
			serviceUsageEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_usage_events, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceUsageEventPrimaryKeyColumns))
			copy(conflict, serviceUsageEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_usage_events\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceUsageEventType, serviceUsageEventMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_usage_events")
	}

	if !cached {
		serviceUsageEventUpsertCacheMut.Lock()
		serviceUsageEventUpsertCache[key] = cache
		serviceUsageEventUpsertCacheMut.Unlock()
	}

	return nil
}

type ServiceUsageEventDeleter interface {
	Delete(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceUsageEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q serviceUsageEventQuery) Delete(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceUsageEvent provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceUsageEventPrimaryKeyMapping)
	sql := "DELETE FROM \"service_usage_events\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_usage_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_usage_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceUsageEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceUsageEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_usage_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_usage_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q serviceUsageEventQuery) DeleteAllSlice(o ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceUsageEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_usage_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceUsageEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceUsageEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_usage_events")
	}

	return rowsAff, nil
}

type ServiceUsageEventReloader interface {
	Reload(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q serviceUsageEventQuery) Reload(o *ServiceUsageEvent, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceUsageEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q serviceUsageEventQuery) ReloadAll(o *ServiceUsageEventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceUsageEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceUsageEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_usage_events\".* FROM \"service_usage_events\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceUsageEventPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceUsageEventSlice")
	}

	*o = slice

	return nil
}

// ServiceUsageEventExists checks if the ServiceUsageEvent row exists.
func ServiceUsageEventExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_usage_events\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_usage_events exists")
	}

	return exists, nil
}
