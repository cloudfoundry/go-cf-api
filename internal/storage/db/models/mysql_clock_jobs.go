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

type ClockJobUpserter interface {
	Upsert(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLClockJobUniqueColumns = []string{
	"id",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q clockJobQuery) Upsert(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clock_jobs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(clockJobColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLClockJobUniqueColumns, o)

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

	clockJobUpsertCacheMut.RLock()
	cache, cached := clockJobUpsertCache[key]
	clockJobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			clockJobAllColumns,
			clockJobColumnsWithDefault,
			clockJobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			clockJobAllColumns,
			clockJobPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert clock_jobs, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`clock_jobs`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `clock_jobs` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(clockJobType, clockJobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(clockJobType, clockJobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for clock_jobs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clockJobMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(clockJobType, clockJobMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for clock_jobs")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for clock_jobs")
	}

CacheNoHooks:
	if !cached {
		clockJobUpsertCacheMut.Lock()
		clockJobUpsertCache[key] = cache
		clockJobUpsertCacheMut.Unlock()
	}

	return nil
}

// ClockJob is an object representing the database table.
type ClockJob struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name            string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	LastStartedAt   null.Time `boil:"last_started_at" json:"last_started_at,omitempty" toml:"last_started_at" yaml:"last_started_at,omitempty"`
	LastCompletedAt null.Time `boil:"last_completed_at" json:"last_completed_at,omitempty" toml:"last_completed_at" yaml:"last_completed_at,omitempty"`

	R *clockJobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L clockJobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ClockJobColumns = struct {
	ID              string
	Name            string
	LastStartedAt   string
	LastCompletedAt string
}{
	ID:              "id",
	Name:            "name",
	LastStartedAt:   "last_started_at",
	LastCompletedAt: "last_completed_at",
}

var ClockJobTableColumns = struct {
	ID              string
	Name            string
	LastStartedAt   string
	LastCompletedAt string
}{
	ID:              "clock_jobs.id",
	Name:            "clock_jobs.name",
	LastStartedAt:   "clock_jobs.last_started_at",
	LastCompletedAt: "clock_jobs.last_completed_at",
}

// Generated where

var ClockJobWhere = struct {
	ID              whereHelperint
	Name            whereHelperstring
	LastStartedAt   whereHelpernull_Time
	LastCompletedAt whereHelpernull_Time
}{
	ID:              whereHelperint{field: "`clock_jobs`.`id`"},
	Name:            whereHelperstring{field: "`clock_jobs`.`name`"},
	LastStartedAt:   whereHelpernull_Time{field: "`clock_jobs`.`last_started_at`"},
	LastCompletedAt: whereHelpernull_Time{field: "`clock_jobs`.`last_completed_at`"},
}

// ClockJobRels is where relationship names are stored.
var ClockJobRels = struct {
}{}

// clockJobR is where relationships are stored.
type clockJobR struct {
}

// NewStruct creates a new relationship struct
func (*clockJobR) NewStruct() *clockJobR {
	return &clockJobR{}
}

// clockJobL is where Load methods for each relationship are stored.
type clockJobL struct{}

var (
	clockJobAllColumns            = []string{"id", "name", "last_started_at", "last_completed_at"}
	clockJobColumnsWithoutDefault = []string{"name", "last_started_at", "last_completed_at"}
	clockJobColumnsWithDefault    = []string{"id"}
	clockJobPrimaryKeyColumns     = []string{"id"}
)

type (
	// ClockJobSlice is an alias for a slice of pointers to ClockJob.
	// This should almost always be used instead of []ClockJob.
	ClockJobSlice []*ClockJob

	clockJobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	clockJobType                 = reflect.TypeOf(&ClockJob{})
	clockJobMapping              = queries.MakeStructMapping(clockJobType)
	clockJobPrimaryKeyMapping, _ = queries.BindMapping(clockJobType, clockJobMapping, clockJobPrimaryKeyColumns)
	clockJobInsertCacheMut       sync.RWMutex
	clockJobInsertCache          = make(map[string]insertCache)
	clockJobUpdateCacheMut       sync.RWMutex
	clockJobUpdateCache          = make(map[string]updateCache)
	clockJobUpsertCacheMut       sync.RWMutex
	clockJobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ClockJobFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ClockJob, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ClockJobSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single clockJob record from the query.
func (q clockJobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ClockJob, error) {
	o := &ClockJob{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for clock_jobs")
	}

	return o, nil
}

// All returns all ClockJob records from the query.
func (q clockJobQuery) All(ctx context.Context, exec boil.ContextExecutor) (ClockJobSlice, error) {
	var o []*ClockJob

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ClockJob slice")
	}

	return o, nil
}

// Count returns the count of all ClockJob records in the query.
func (q clockJobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count clock_jobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q clockJobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if clock_jobs exists")
	}

	return count > 0, nil
}

// ClockJobs retrieves all the records using an executor.
func ClockJobs(mods ...qm.QueryMod) clockJobQuery {
	mods = append(mods, qm.From("`clock_jobs`"))
	return clockJobQuery{NewQuery(mods...)}
}

type ClockJobFinder interface {
	FindClockJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ClockJob, error)
}

// FindClockJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindClockJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ClockJob, error) {
	clockJobObj := &ClockJob{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `clock_jobs` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, clockJobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from clock_jobs")
	}

	return clockJobObj, nil
}

type ClockJobInserter interface {
	Insert(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q clockJobQuery) Insert(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no clock_jobs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(clockJobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	clockJobInsertCacheMut.RLock()
	cache, cached := clockJobInsertCache[key]
	clockJobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			clockJobAllColumns,
			clockJobColumnsWithDefault,
			clockJobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(clockJobType, clockJobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(clockJobType, clockJobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `clock_jobs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `clock_jobs` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `clock_jobs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, clockJobPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into clock_jobs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == clockJobMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for clock_jobs")
	}

CacheNoHooks:
	if !cached {
		clockJobInsertCacheMut.Lock()
		clockJobInsertCache[key] = cache
		clockJobInsertCacheMut.Unlock()
	}

	return nil
}

type ClockJobUpdater interface {
	Update(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ClockJobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ClockJob.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q clockJobQuery) Update(o *ClockJob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	clockJobUpdateCacheMut.RLock()
	cache, cached := clockJobUpdateCache[key]
	clockJobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			clockJobAllColumns,
			clockJobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update clock_jobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `clock_jobs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, clockJobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(clockJobType, clockJobMapping, append(wl, clockJobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update clock_jobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for clock_jobs")
	}

	if !cached {
		clockJobUpdateCacheMut.Lock()
		clockJobUpdateCache[key] = cache
		clockJobUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q clockJobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for clock_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for clock_jobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q clockJobQuery) UpdateAllSlice(o ClockJobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clockJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `clock_jobs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clockJobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in clockJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all clockJob")
	}
	return rowsAff, nil
}

type ClockJobDeleter interface {
	Delete(o *ClockJob, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ClockJob record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q clockJobQuery) Delete(o *ClockJob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ClockJob provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), clockJobPrimaryKeyMapping)
	sql := "DELETE FROM `clock_jobs` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from clock_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for clock_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q clockJobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no clockJobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clock_jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clock_jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q clockJobQuery) DeleteAllSlice(o ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clockJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `clock_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clockJobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from clockJob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for clock_jobs")
	}

	return rowsAff, nil
}

type ClockJobReloader interface {
	Reload(o *ClockJob, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q clockJobQuery) Reload(o *ClockJob, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindClockJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q clockJobQuery) ReloadAll(o *ClockJobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ClockJobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), clockJobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `clock_jobs`.* FROM `clock_jobs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, clockJobPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ClockJobSlice")
	}

	*o = slice

	return nil
}

// ClockJobExists checks if the ClockJob row exists.
func ClockJobExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `clock_jobs` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if clock_jobs exists")
	}

	return exists, nil
}
