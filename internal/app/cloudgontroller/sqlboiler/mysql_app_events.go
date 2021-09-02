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

// AppEvent is an object representing the database table.
type AppEvent struct {
	ID              int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID            string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt       time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt       null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	AppID           int         `boil:"app_id" json:"app_id" toml:"app_id" yaml:"app_id"`
	InstanceGUID    string      `boil:"instance_guid" json:"instance_guid" toml:"instance_guid" yaml:"instance_guid"`
	InstanceIndex   int         `boil:"instance_index" json:"instance_index" toml:"instance_index" yaml:"instance_index"`
	ExitStatus      int         `boil:"exit_status" json:"exit_status" toml:"exit_status" yaml:"exit_status"`
	Timestamp       time.Time   `boil:"timestamp" json:"timestamp" toml:"timestamp" yaml:"timestamp"`
	ExitDescription null.String `boil:"exit_description" json:"exit_description,omitempty" toml:"exit_description" yaml:"exit_description,omitempty"`

	R *appEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L appEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AppEventColumns = struct {
	ID              string
	GUID            string
	CreatedAt       string
	UpdatedAt       string
	AppID           string
	InstanceGUID    string
	InstanceIndex   string
	ExitStatus      string
	Timestamp       string
	ExitDescription string
}{
	ID:              "id",
	GUID:            "guid",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	AppID:           "app_id",
	InstanceGUID:    "instance_guid",
	InstanceIndex:   "instance_index",
	ExitStatus:      "exit_status",
	Timestamp:       "timestamp",
	ExitDescription: "exit_description",
}

var AppEventTableColumns = struct {
	ID              string
	GUID            string
	CreatedAt       string
	UpdatedAt       string
	AppID           string
	InstanceGUID    string
	InstanceIndex   string
	ExitStatus      string
	Timestamp       string
	ExitDescription string
}{
	ID:              "app_events.id",
	GUID:            "app_events.guid",
	CreatedAt:       "app_events.created_at",
	UpdatedAt:       "app_events.updated_at",
	AppID:           "app_events.app_id",
	InstanceGUID:    "app_events.instance_guid",
	InstanceIndex:   "app_events.instance_index",
	ExitStatus:      "app_events.exit_status",
	Timestamp:       "app_events.timestamp",
	ExitDescription: "app_events.exit_description",
}

// Generated where

var AppEventWhere = struct {
	ID              whereHelperint
	GUID            whereHelperstring
	CreatedAt       whereHelpertime_Time
	UpdatedAt       whereHelpernull_Time
	AppID           whereHelperint
	InstanceGUID    whereHelperstring
	InstanceIndex   whereHelperint
	ExitStatus      whereHelperint
	Timestamp       whereHelpertime_Time
	ExitDescription whereHelpernull_String
}{
	ID:              whereHelperint{field: "`app_events`.`id`"},
	GUID:            whereHelperstring{field: "`app_events`.`guid`"},
	CreatedAt:       whereHelpertime_Time{field: "`app_events`.`created_at`"},
	UpdatedAt:       whereHelpernull_Time{field: "`app_events`.`updated_at`"},
	AppID:           whereHelperint{field: "`app_events`.`app_id`"},
	InstanceGUID:    whereHelperstring{field: "`app_events`.`instance_guid`"},
	InstanceIndex:   whereHelperint{field: "`app_events`.`instance_index`"},
	ExitStatus:      whereHelperint{field: "`app_events`.`exit_status`"},
	Timestamp:       whereHelpertime_Time{field: "`app_events`.`timestamp`"},
	ExitDescription: whereHelpernull_String{field: "`app_events`.`exit_description`"},
}

// AppEventRels is where relationship names are stored.
var AppEventRels = struct {
	App string
}{
	App: "App",
}

// appEventR is where relationships are stored.
type appEventR struct {
	App *Process `boil:"App" json:"App" toml:"App" yaml:"App"`
}

// NewStruct creates a new relationship struct
func (*appEventR) NewStruct() *appEventR {
	return &appEventR{}
}

// appEventL is where Load methods for each relationship are stored.
type appEventL struct{}

var (
	appEventAllColumns            = []string{"id", "guid", "created_at", "updated_at", "app_id", "instance_guid", "instance_index", "exit_status", "timestamp", "exit_description"}
	appEventColumnsWithoutDefault = []string{"guid", "updated_at", "app_id", "instance_guid", "instance_index", "exit_status", "timestamp", "exit_description"}
	appEventColumnsWithDefault    = []string{"id", "created_at"}
	appEventPrimaryKeyColumns     = []string{"id"}
)

type (
	// AppEventSlice is an alias for a slice of pointers to AppEvent.
	// This should almost always be used instead of []AppEvent.
	AppEventSlice []*AppEvent

	AppEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	appEventType                 = reflect.TypeOf(&AppEvent{})
	appEventMapping              = queries.MakeStructMapping(appEventType)
	appEventPrimaryKeyMapping, _ = queries.BindMapping(appEventType, appEventMapping, appEventPrimaryKeyColumns)
	appEventInsertCacheMut       sync.RWMutex
	appEventInsertCache          = make(map[string]insertCache)
	appEventUpdateCacheMut       sync.RWMutex
	appEventUpdateCache          = make(map[string]updateCache)
	appEventUpsertCacheMut       sync.RWMutex
	appEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type AppEventFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*AppEvent, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (AppEventSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single appEvent record from the query.
func (q AppEventQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AppEvent, error) {
	o := &AppEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for app_events")
	}

	return o, nil
}

// All returns all AppEvent records from the query.
func (q AppEventQuery) All(ctx context.Context, exec boil.ContextExecutor) (AppEventSlice, error) {
	var o []*AppEvent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AppEvent slice")
	}

	return o, nil
}

// Count returns the count of all AppEvent records in the query.
func (q AppEventQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count app_events rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q AppEventQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if app_events exists")
	}

	return count > 0, nil
}

// App pointed to by the foreign key.
func (q AppEventQuery) App(o *AppEvent, mods ...qm.QueryMod) ProcessQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.AppID),
	}

	queryMods = append(queryMods, mods...)

	query := Processes(queryMods...)
	queries.SetFrom(query.Query, "`processes`")

	return query
}

// LoadApp allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (appEventL) LoadApp(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAppEvent interface{}, mods queries.Applicator) error {
	var slice []*AppEvent
	var object *AppEvent

	if singular {
		object = maybeAppEvent.(*AppEvent)
	} else {
		slice = *maybeAppEvent.(*[]*AppEvent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &appEventR{}
		}
		args = append(args, object.AppID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &appEventR{}
			}

			for _, a := range args {
				if a == obj.AppID {
					continue Outer
				}
			}

			args = append(args, obj.AppID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`processes`),
		qm.WhereIn(`processes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Process")
	}

	var resultSlice []*Process
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Process")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for processes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for processes")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.App = foreign
		if foreign.R == nil {
			foreign.R = &processR{}
		}
		foreign.R.AppAppEvents = append(foreign.R.AppAppEvents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AppID == foreign.ID {
				local.R.App = foreign
				if foreign.R == nil {
					foreign.R = &processR{}
				}
				foreign.R.AppAppEvents = append(foreign.R.AppAppEvents, local)
				break
			}
		}
	}

	return nil
}

// SetApp of the appEvent to the related item.
// Sets o.R.App to related.
// Adds o to related.R.AppAppEvents.
func (q AppEventQuery) SetApp(o *AppEvent, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Process) error {
	var err error
	if insert {
		if err = Processes().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `app_events` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"app_id"}),
		strmangle.WhereClause("`", "`", 0, appEventPrimaryKeyColumns),
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

	o.AppID = related.ID
	if o.R == nil {
		o.R = &appEventR{
			App: related,
		}
	} else {
		o.R.App = related
	}

	if related.R == nil {
		related.R = &processR{
			AppAppEvents: AppEventSlice{o},
		}
	} else {
		related.R.AppAppEvents = append(related.R.AppAppEvents, o)
	}

	return nil
}

// AppEvents retrieves all the records using an executor.
func AppEvents(mods ...qm.QueryMod) AppEventQuery {
	mods = append(mods, qm.From("`app_events`"))
	return AppEventQuery{NewQuery(mods...)}
}

type AppEventFinder interface {
	FindAppEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AppEvent, error)
}

// FindAppEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAppEvent(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AppEvent, error) {
	appEventObj := &AppEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `app_events` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, appEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from app_events")
	}

	return appEventObj, nil
}

type AppEventInserter interface {
	Insert(o *AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q AppEventQuery) Insert(o *AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_events provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(appEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	appEventInsertCacheMut.RLock()
	cache, cached := appEventInsertCache[key]
	appEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			appEventAllColumns,
			appEventColumnsWithDefault,
			appEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(appEventType, appEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(appEventType, appEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `app_events` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `app_events` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `app_events` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, appEventPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into app_events")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == appEventMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for app_events")
	}

CacheNoHooks:
	if !cached {
		appEventInsertCacheMut.Lock()
		appEventInsertCache[key] = cache
		appEventInsertCacheMut.Unlock()
	}

	return nil
}

type AppEventUpdater interface {
	Update(o *AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o AppEventSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the AppEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q AppEventQuery) Update(o *AppEvent, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	appEventUpdateCacheMut.RLock()
	cache, cached := appEventUpdateCache[key]
	appEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			appEventAllColumns,
			appEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update app_events, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `app_events` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, appEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(appEventType, appEventMapping, append(wl, appEventPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update app_events row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for app_events")
	}

	if !cached {
		appEventUpdateCacheMut.Lock()
		appEventUpdateCache[key] = cache
		appEventUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q AppEventQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for app_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for app_events")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q AppEventQuery) UpdateAllSlice(o AppEventSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `app_events` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appEventPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in appEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all appEvent")
	}
	return rowsAff, nil
}

type AppEventDeleter interface {
	Delete(o *AppEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o AppEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single AppEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q AppEventQuery) Delete(o *AppEvent, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AppEvent provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), appEventPrimaryKeyMapping)
	sql := "DELETE FROM `app_events` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from app_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for app_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q AppEventQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no appEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from app_events")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_events")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q AppEventQuery) DeleteAllSlice(o AppEventSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `app_events` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appEventPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from appEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for app_events")
	}

	return rowsAff, nil
}

type AppEventReloader interface {
	Reload(o *AppEvent, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *AppEventSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q AppEventQuery) Reload(o *AppEvent, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAppEvent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q AppEventQuery) ReloadAll(o *AppEventSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AppEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `app_events`.* FROM `app_events` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, appEventPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AppEventSlice")
	}

	*o = slice

	return nil
}

// AppEventExists checks if the AppEvent row exists.
func AppEventExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `app_events` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if app_events exists")
	}

	return exists, nil
}

var mySQLAppEventUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AppEvent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no app_events provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(appEventColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAppEventUniqueColumns, o)

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

	appEventUpsertCacheMut.RLock()
	cache, cached := appEventUpsertCache[key]
	appEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			appEventAllColumns,
			appEventColumnsWithDefault,
			appEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			appEventAllColumns,
			appEventPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert app_events, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`app_events`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `app_events` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(appEventType, appEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(appEventType, appEventMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for app_events")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == appEventMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(appEventType, appEventMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for app_events")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for app_events")
	}

CacheNoHooks:
	if !cached {
		appEventUpsertCacheMut.Lock()
		appEventUpsertCache[key] = cache
		appEventUpsertCacheMut.Unlock()
	}

	return nil
}
