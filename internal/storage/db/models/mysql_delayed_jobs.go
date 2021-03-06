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

type DelayedJobUpserter interface {
	Upsert(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLDelayedJobUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q delayedJobQuery) Upsert(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no delayed_jobs provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(delayedJobColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDelayedJobUniqueColumns, o)

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

	delayedJobUpsertCacheMut.RLock()
	cache, cached := delayedJobUpsertCache[key]
	delayedJobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			delayedJobAllColumns,
			delayedJobColumnsWithDefault,
			delayedJobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			delayedJobAllColumns,
			delayedJobPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert delayed_jobs, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`delayed_jobs`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `delayed_jobs` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(delayedJobType, delayedJobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(delayedJobType, delayedJobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for delayed_jobs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == delayedJobMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(delayedJobType, delayedJobMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for delayed_jobs")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for delayed_jobs")
	}

CacheNoHooks:
	if !cached {
		delayedJobUpsertCacheMut.Lock()
		delayedJobUpsertCache[key] = cache
		delayedJobUpsertCacheMut.Unlock()
	}

	return nil
}

// DelayedJob is an object representing the database table.
type DelayedJob struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID       string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Priority   null.Int    `boil:"priority" json:"priority,omitempty" toml:"priority" yaml:"priority,omitempty"`
	Attempts   null.Int    `boil:"attempts" json:"attempts,omitempty" toml:"attempts" yaml:"attempts,omitempty"`
	Handler    null.String `boil:"handler" json:"handler,omitempty" toml:"handler" yaml:"handler,omitempty"`
	LastError  null.String `boil:"last_error" json:"last_error,omitempty" toml:"last_error" yaml:"last_error,omitempty"`
	RunAt      null.Time   `boil:"run_at" json:"run_at,omitempty" toml:"run_at" yaml:"run_at,omitempty"`
	LockedAt   null.Time   `boil:"locked_at" json:"locked_at,omitempty" toml:"locked_at" yaml:"locked_at,omitempty"`
	FailedAt   null.Time   `boil:"failed_at" json:"failed_at,omitempty" toml:"failed_at" yaml:"failed_at,omitempty"`
	LockedBy   null.String `boil:"locked_by" json:"locked_by,omitempty" toml:"locked_by" yaml:"locked_by,omitempty"`
	Queue      null.String `boil:"queue" json:"queue,omitempty" toml:"queue" yaml:"queue,omitempty"`
	CFAPIError null.String `boil:"cf_api_error" json:"cf_api_error,omitempty" toml:"cf_api_error" yaml:"cf_api_error,omitempty"`

	R *delayedJobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L delayedJobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DelayedJobColumns = struct {
	ID         string
	GUID       string
	CreatedAt  string
	UpdatedAt  string
	Priority   string
	Attempts   string
	Handler    string
	LastError  string
	RunAt      string
	LockedAt   string
	FailedAt   string
	LockedBy   string
	Queue      string
	CFAPIError string
}{
	ID:         "id",
	GUID:       "guid",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Priority:   "priority",
	Attempts:   "attempts",
	Handler:    "handler",
	LastError:  "last_error",
	RunAt:      "run_at",
	LockedAt:   "locked_at",
	FailedAt:   "failed_at",
	LockedBy:   "locked_by",
	Queue:      "queue",
	CFAPIError: "cf_api_error",
}

var DelayedJobTableColumns = struct {
	ID         string
	GUID       string
	CreatedAt  string
	UpdatedAt  string
	Priority   string
	Attempts   string
	Handler    string
	LastError  string
	RunAt      string
	LockedAt   string
	FailedAt   string
	LockedBy   string
	Queue      string
	CFAPIError string
}{
	ID:         "delayed_jobs.id",
	GUID:       "delayed_jobs.guid",
	CreatedAt:  "delayed_jobs.created_at",
	UpdatedAt:  "delayed_jobs.updated_at",
	Priority:   "delayed_jobs.priority",
	Attempts:   "delayed_jobs.attempts",
	Handler:    "delayed_jobs.handler",
	LastError:  "delayed_jobs.last_error",
	RunAt:      "delayed_jobs.run_at",
	LockedAt:   "delayed_jobs.locked_at",
	FailedAt:   "delayed_jobs.failed_at",
	LockedBy:   "delayed_jobs.locked_by",
	Queue:      "delayed_jobs.queue",
	CFAPIError: "delayed_jobs.cf_api_error",
}

// Generated where

var DelayedJobWhere = struct {
	ID         whereHelperint
	GUID       whereHelperstring
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpernull_Time
	Priority   whereHelpernull_Int
	Attempts   whereHelpernull_Int
	Handler    whereHelpernull_String
	LastError  whereHelpernull_String
	RunAt      whereHelpernull_Time
	LockedAt   whereHelpernull_Time
	FailedAt   whereHelpernull_Time
	LockedBy   whereHelpernull_String
	Queue      whereHelpernull_String
	CFAPIError whereHelpernull_String
}{
	ID:         whereHelperint{field: "`delayed_jobs`.`id`"},
	GUID:       whereHelperstring{field: "`delayed_jobs`.`guid`"},
	CreatedAt:  whereHelpertime_Time{field: "`delayed_jobs`.`created_at`"},
	UpdatedAt:  whereHelpernull_Time{field: "`delayed_jobs`.`updated_at`"},
	Priority:   whereHelpernull_Int{field: "`delayed_jobs`.`priority`"},
	Attempts:   whereHelpernull_Int{field: "`delayed_jobs`.`attempts`"},
	Handler:    whereHelpernull_String{field: "`delayed_jobs`.`handler`"},
	LastError:  whereHelpernull_String{field: "`delayed_jobs`.`last_error`"},
	RunAt:      whereHelpernull_Time{field: "`delayed_jobs`.`run_at`"},
	LockedAt:   whereHelpernull_Time{field: "`delayed_jobs`.`locked_at`"},
	FailedAt:   whereHelpernull_Time{field: "`delayed_jobs`.`failed_at`"},
	LockedBy:   whereHelpernull_String{field: "`delayed_jobs`.`locked_by`"},
	Queue:      whereHelpernull_String{field: "`delayed_jobs`.`queue`"},
	CFAPIError: whereHelpernull_String{field: "`delayed_jobs`.`cf_api_error`"},
}

// DelayedJobRels is where relationship names are stored.
var DelayedJobRels = struct {
}{}

// delayedJobR is where relationships are stored.
type delayedJobR struct {
}

// NewStruct creates a new relationship struct
func (*delayedJobR) NewStruct() *delayedJobR {
	return &delayedJobR{}
}

// delayedJobL is where Load methods for each relationship are stored.
type delayedJobL struct{}

var (
	delayedJobAllColumns            = []string{"id", "guid", "created_at", "updated_at", "priority", "attempts", "handler", "last_error", "run_at", "locked_at", "failed_at", "locked_by", "queue", "cf_api_error"}
	delayedJobColumnsWithoutDefault = []string{"guid", "updated_at", "handler", "last_error", "run_at", "locked_at", "failed_at", "locked_by", "queue", "cf_api_error"}
	delayedJobColumnsWithDefault    = []string{"id", "created_at", "priority", "attempts"}
	delayedJobPrimaryKeyColumns     = []string{"id"}
)

type (
	// DelayedJobSlice is an alias for a slice of pointers to DelayedJob.
	// This should almost always be used instead of []DelayedJob.
	DelayedJobSlice []*DelayedJob

	delayedJobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	delayedJobType                 = reflect.TypeOf(&DelayedJob{})
	delayedJobMapping              = queries.MakeStructMapping(delayedJobType)
	delayedJobPrimaryKeyMapping, _ = queries.BindMapping(delayedJobType, delayedJobMapping, delayedJobPrimaryKeyColumns)
	delayedJobInsertCacheMut       sync.RWMutex
	delayedJobInsertCache          = make(map[string]insertCache)
	delayedJobUpdateCacheMut       sync.RWMutex
	delayedJobUpdateCache          = make(map[string]updateCache)
	delayedJobUpsertCacheMut       sync.RWMutex
	delayedJobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type DelayedJobFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*DelayedJob, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (DelayedJobSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single delayedJob record from the query.
func (q delayedJobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DelayedJob, error) {
	o := &DelayedJob{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for delayed_jobs")
	}

	return o, nil
}

// All returns all DelayedJob records from the query.
func (q delayedJobQuery) All(ctx context.Context, exec boil.ContextExecutor) (DelayedJobSlice, error) {
	var o []*DelayedJob

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DelayedJob slice")
	}

	return o, nil
}

// Count returns the count of all DelayedJob records in the query.
func (q delayedJobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count delayed_jobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q delayedJobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if delayed_jobs exists")
	}

	return count > 0, nil
}

// DelayedJobs retrieves all the records using an executor.
func DelayedJobs(mods ...qm.QueryMod) delayedJobQuery {
	mods = append(mods, qm.From("`delayed_jobs`"))
	return delayedJobQuery{NewQuery(mods...)}
}

type DelayedJobFinder interface {
	FindDelayedJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DelayedJob, error)
}

// FindDelayedJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDelayedJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DelayedJob, error) {
	delayedJobObj := &DelayedJob{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `delayed_jobs` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, delayedJobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from delayed_jobs")
	}

	return delayedJobObj, nil
}

type DelayedJobInserter interface {
	Insert(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q delayedJobQuery) Insert(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no delayed_jobs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(delayedJobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	delayedJobInsertCacheMut.RLock()
	cache, cached := delayedJobInsertCache[key]
	delayedJobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			delayedJobAllColumns,
			delayedJobColumnsWithDefault,
			delayedJobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(delayedJobType, delayedJobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(delayedJobType, delayedJobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `delayed_jobs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `delayed_jobs` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `delayed_jobs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, delayedJobPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into delayed_jobs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == delayedJobMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for delayed_jobs")
	}

CacheNoHooks:
	if !cached {
		delayedJobInsertCacheMut.Lock()
		delayedJobInsertCache[key] = cache
		delayedJobInsertCacheMut.Unlock()
	}

	return nil
}

type DelayedJobUpdater interface {
	Update(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the DelayedJob.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q delayedJobQuery) Update(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	delayedJobUpdateCacheMut.RLock()
	cache, cached := delayedJobUpdateCache[key]
	delayedJobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			delayedJobAllColumns,
			delayedJobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update delayed_jobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `delayed_jobs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, delayedJobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(delayedJobType, delayedJobMapping, append(wl, delayedJobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update delayed_jobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for delayed_jobs")
	}

	if !cached {
		delayedJobUpdateCacheMut.Lock()
		delayedJobUpdateCache[key] = cache
		delayedJobUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q delayedJobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for delayed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for delayed_jobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q delayedJobQuery) UpdateAllSlice(o DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), delayedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `delayed_jobs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, delayedJobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in delayedJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all delayedJob")
	}
	return rowsAff, nil
}

type DelayedJobDeleter interface {
	Delete(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single DelayedJob record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q delayedJobQuery) Delete(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DelayedJob provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), delayedJobPrimaryKeyMapping)
	sql := "DELETE FROM `delayed_jobs` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from delayed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for delayed_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q delayedJobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no delayedJobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from delayed_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for delayed_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q delayedJobQuery) DeleteAllSlice(o DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), delayedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `delayed_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, delayedJobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from delayedJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for delayed_jobs")
	}

	return rowsAff, nil
}

type DelayedJobReloader interface {
	Reload(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q delayedJobQuery) Reload(o *DelayedJob, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDelayedJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q delayedJobQuery) ReloadAll(o *DelayedJobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DelayedJobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), delayedJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `delayed_jobs`.* FROM `delayed_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, delayedJobPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DelayedJobSlice")
	}

	*o = slice

	return nil
}

// DelayedJobExists checks if the DelayedJob row exists.
func DelayedJobExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `delayed_jobs` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if delayed_jobs exists")
	}

	return exists, nil
}
