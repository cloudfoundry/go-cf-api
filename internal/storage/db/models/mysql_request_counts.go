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

type RequestCountUpserter interface {
	Upsert(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLRequestCountUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q requestCountQuery) Upsert(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no request_counts provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(requestCountColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRequestCountUniqueColumns, o)

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

	requestCountUpsertCacheMut.RLock()
	cache, cached := requestCountUpsertCache[key]
	requestCountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			requestCountAllColumns,
			requestCountColumnsWithDefault,
			requestCountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			requestCountAllColumns,
			requestCountPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert request_counts, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`request_counts`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `request_counts` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(requestCountType, requestCountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(requestCountType, requestCountMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for request_counts")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == requestCountMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(requestCountType, requestCountMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for request_counts")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for request_counts")
	}

CacheNoHooks:
	if !cached {
		requestCountUpsertCacheMut.Lock()
		requestCountUpsertCache[key] = cache
		requestCountUpsertCacheMut.Unlock()
	}

	return nil
}

// RequestCount is an object representing the database table.
type RequestCount struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserGUID   null.String `boil:"user_guid" json:"user_guid,omitempty" toml:"user_guid" yaml:"user_guid,omitempty"`
	Count      null.Int    `boil:"count" json:"count,omitempty" toml:"count" yaml:"count,omitempty"`
	ValidUntil null.Time   `boil:"valid_until" json:"valid_until,omitempty" toml:"valid_until" yaml:"valid_until,omitempty"`

	R *requestCountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L requestCountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RequestCountColumns = struct {
	ID         string
	UserGUID   string
	Count      string
	ValidUntil string
}{
	ID:         "id",
	UserGUID:   "user_guid",
	Count:      "count",
	ValidUntil: "valid_until",
}

var RequestCountTableColumns = struct {
	ID         string
	UserGUID   string
	Count      string
	ValidUntil string
}{
	ID:         "request_counts.id",
	UserGUID:   "request_counts.user_guid",
	Count:      "request_counts.count",
	ValidUntil: "request_counts.valid_until",
}

// Generated where

var RequestCountWhere = struct {
	ID         whereHelperint
	UserGUID   whereHelpernull_String
	Count      whereHelpernull_Int
	ValidUntil whereHelpernull_Time
}{
	ID:         whereHelperint{field: "`request_counts`.`id`"},
	UserGUID:   whereHelpernull_String{field: "`request_counts`.`user_guid`"},
	Count:      whereHelpernull_Int{field: "`request_counts`.`count`"},
	ValidUntil: whereHelpernull_Time{field: "`request_counts`.`valid_until`"},
}

// RequestCountRels is where relationship names are stored.
var RequestCountRels = struct {
}{}

// requestCountR is where relationships are stored.
type requestCountR struct {
}

// NewStruct creates a new relationship struct
func (*requestCountR) NewStruct() *requestCountR {
	return &requestCountR{}
}

// requestCountL is where Load methods for each relationship are stored.
type requestCountL struct{}

var (
	requestCountAllColumns            = []string{"id", "user_guid", "count", "valid_until"}
	requestCountColumnsWithoutDefault = []string{"user_guid", "valid_until"}
	requestCountColumnsWithDefault    = []string{"id", "count"}
	requestCountPrimaryKeyColumns     = []string{"id"}
)

type (
	// RequestCountSlice is an alias for a slice of pointers to RequestCount.
	// This should almost always be used instead of []RequestCount.
	RequestCountSlice []*RequestCount

	requestCountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	requestCountType                 = reflect.TypeOf(&RequestCount{})
	requestCountMapping              = queries.MakeStructMapping(requestCountType)
	requestCountPrimaryKeyMapping, _ = queries.BindMapping(requestCountType, requestCountMapping, requestCountPrimaryKeyColumns)
	requestCountInsertCacheMut       sync.RWMutex
	requestCountInsertCache          = make(map[string]insertCache)
	requestCountUpdateCacheMut       sync.RWMutex
	requestCountUpdateCache          = make(map[string]updateCache)
	requestCountUpsertCacheMut       sync.RWMutex
	requestCountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type RequestCountFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*RequestCount, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (RequestCountSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single requestCount record from the query.
func (q requestCountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RequestCount, error) {
	o := &RequestCount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for request_counts")
	}

	return o, nil
}

// All returns all RequestCount records from the query.
func (q requestCountQuery) All(ctx context.Context, exec boil.ContextExecutor) (RequestCountSlice, error) {
	var o []*RequestCount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RequestCount slice")
	}

	return o, nil
}

// Count returns the count of all RequestCount records in the query.
func (q requestCountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count request_counts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q requestCountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if request_counts exists")
	}

	return count > 0, nil
}

// RequestCounts retrieves all the records using an executor.
func RequestCounts(mods ...qm.QueryMod) requestCountQuery {
	mods = append(mods, qm.From("`request_counts`"))
	return requestCountQuery{NewQuery(mods...)}
}

type RequestCountFinder interface {
	FindRequestCount(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RequestCount, error)
}

// FindRequestCount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRequestCount(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RequestCount, error) {
	requestCountObj := &RequestCount{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `request_counts` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, requestCountObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from request_counts")
	}

	return requestCountObj, nil
}

type RequestCountInserter interface {
	Insert(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q requestCountQuery) Insert(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no request_counts provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(requestCountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	requestCountInsertCacheMut.RLock()
	cache, cached := requestCountInsertCache[key]
	requestCountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			requestCountAllColumns,
			requestCountColumnsWithDefault,
			requestCountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(requestCountType, requestCountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(requestCountType, requestCountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `request_counts` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `request_counts` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `request_counts` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, requestCountPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into request_counts")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == requestCountMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for request_counts")
	}

CacheNoHooks:
	if !cached {
		requestCountInsertCacheMut.Lock()
		requestCountInsertCache[key] = cache
		requestCountInsertCacheMut.Unlock()
	}

	return nil
}

type RequestCountUpdater interface {
	Update(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o RequestCountSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the RequestCount.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q requestCountQuery) Update(o *RequestCount, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	requestCountUpdateCacheMut.RLock()
	cache, cached := requestCountUpdateCache[key]
	requestCountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			requestCountAllColumns,
			requestCountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update request_counts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `request_counts` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, requestCountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(requestCountType, requestCountMapping, append(wl, requestCountPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update request_counts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for request_counts")
	}

	if !cached {
		requestCountUpdateCacheMut.Lock()
		requestCountUpdateCache[key] = cache
		requestCountUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q requestCountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for request_counts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for request_counts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q requestCountQuery) UpdateAllSlice(o RequestCountSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `request_counts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestCountPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in requestCount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all requestCount")
	}
	return rowsAff, nil
}

type RequestCountDeleter interface {
	Delete(o *RequestCount, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single RequestCount record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q requestCountQuery) Delete(o *RequestCount, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RequestCount provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), requestCountPrimaryKeyMapping)
	sql := "DELETE FROM `request_counts` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from request_counts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for request_counts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q requestCountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no requestCountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from request_counts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for request_counts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q requestCountQuery) DeleteAllSlice(o RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `request_counts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestCountPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from requestCount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for request_counts")
	}

	return rowsAff, nil
}

type RequestCountReloader interface {
	Reload(o *RequestCount, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q requestCountQuery) Reload(o *RequestCount, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRequestCount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q requestCountQuery) ReloadAll(o *RequestCountSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RequestCountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestCountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `request_counts`.* FROM `request_counts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, requestCountPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RequestCountSlice")
	}

	*o = slice

	return nil
}

// RequestCountExists checks if the RequestCount row exists.
func RequestCountExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `request_counts` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if request_counts exists")
	}

	return exists, nil
}
