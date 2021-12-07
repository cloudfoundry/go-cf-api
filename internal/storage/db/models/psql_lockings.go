//go:build psql
// +build psql

//go:generate sh -c "echo '\x2bbuild unit' > ../../../../buildtags.txt && mockgen -source=$GOFILE -destination=mocks/lockings.go -copyright_file=../../../../buildtags.txt && rm -f ../../../../buildtags.txt"
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

type LockingUpserter interface {
	Upsert(o *Locking, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q lockingQuery) Upsert(o *Locking, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lockings provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(lockingColumnsWithDefault, o)

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

	lockingUpsertCacheMut.RLock()
	cache, cached := lockingUpsertCache[key]
	lockingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lockingAllColumns,
			lockingColumnsWithDefault,
			lockingColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			lockingAllColumns,
			lockingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert lockings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(lockingPrimaryKeyColumns))
			copy(conflict, lockingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"lockings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(lockingType, lockingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lockingType, lockingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert lockings")
	}

	if !cached {
		lockingUpsertCacheMut.Lock()
		lockingUpsertCache[key] = cache
		lockingUpsertCacheMut.Unlock()
	}

	return nil
}

// Locking is an object representing the database table.
type Locking struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *lockingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lockingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LockingColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var LockingTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "lockings.id",
	Name: "lockings.name",
}

// Generated where

var LockingWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
}{
	ID:   whereHelperint{field: "\"lockings\".\"id\""},
	Name: whereHelperstring{field: "\"lockings\".\"name\""},
}

// LockingRels is where relationship names are stored.
var LockingRels = struct {
}{}

// lockingR is where relationships are stored.
type lockingR struct {
}

// NewStruct creates a new relationship struct
func (*lockingR) NewStruct() *lockingR {
	return &lockingR{}
}

// lockingL is where Load methods for each relationship are stored.
type lockingL struct{}

var (
	lockingAllColumns            = []string{"id", "name"}
	lockingColumnsWithoutDefault = []string{"name"}
	lockingColumnsWithDefault    = []string{"id"}
	lockingPrimaryKeyColumns     = []string{"id"}
)

type (
	// LockingSlice is an alias for a slice of pointers to Locking.
	// This should almost always be used instead of []Locking.
	LockingSlice []*Locking

	lockingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lockingType                 = reflect.TypeOf(&Locking{})
	lockingMapping              = queries.MakeStructMapping(lockingType)
	lockingPrimaryKeyMapping, _ = queries.BindMapping(lockingType, lockingMapping, lockingPrimaryKeyColumns)
	lockingInsertCacheMut       sync.RWMutex
	lockingInsertCache          = make(map[string]insertCache)
	lockingUpdateCacheMut       sync.RWMutex
	lockingUpdateCache          = make(map[string]updateCache)
	lockingUpsertCacheMut       sync.RWMutex
	lockingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type LockingFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*Locking, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (LockingSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single locking record from the query.
func (q lockingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Locking, error) {
	o := &Locking{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for lockings")
	}

	return o, nil
}

// All returns all Locking records from the query.
func (q lockingQuery) All(ctx context.Context, exec boil.ContextExecutor) (LockingSlice, error) {
	var o []*Locking

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Locking slice")
	}

	return o, nil
}

// Count returns the count of all Locking records in the query.
func (q lockingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count lockings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lockingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if lockings exists")
	}

	return count > 0, nil
}

// Lockings retrieves all the records using an executor.
func Lockings(mods ...qm.QueryMod) lockingQuery {
	mods = append(mods, qm.From("\"lockings\""))
	return lockingQuery{NewQuery(mods...)}
}

type LockingFinder interface {
	FindLocking(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Locking, error)
}

// FindLocking retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLocking(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Locking, error) {
	lockingObj := &Locking{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"lockings\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lockingObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from lockings")
	}

	return lockingObj, nil
}

type LockingInserter interface {
	Insert(o *Locking, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q lockingQuery) Insert(o *Locking, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no lockings provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(lockingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lockingInsertCacheMut.RLock()
	cache, cached := lockingInsertCache[key]
	lockingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lockingAllColumns,
			lockingColumnsWithDefault,
			lockingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lockingType, lockingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lockingType, lockingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"lockings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"lockings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into lockings")
	}

	if !cached {
		lockingInsertCacheMut.Lock()
		lockingInsertCache[key] = cache
		lockingInsertCacheMut.Unlock()
	}

	return nil
}

type LockingUpdater interface {
	Update(o *Locking, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o LockingSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the Locking.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q lockingQuery) Update(o *Locking, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	lockingUpdateCacheMut.RLock()
	cache, cached := lockingUpdateCache[key]
	lockingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lockingAllColumns,
			lockingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update lockings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"lockings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, lockingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lockingType, lockingMapping, append(wl, lockingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update lockings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for lockings")
	}

	if !cached {
		lockingUpdateCacheMut.Lock()
		lockingUpdateCache[key] = cache
		lockingUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q lockingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for lockings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for lockings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q lockingQuery) UpdateAllSlice(o LockingSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lockingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"lockings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, lockingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in locking slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all locking")
	}
	return rowsAff, nil
}

type LockingDeleter interface {
	Delete(o *Locking, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o LockingSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single Locking record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q lockingQuery) Delete(o *Locking, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Locking provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lockingPrimaryKeyMapping)
	sql := "DELETE FROM \"lockings\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from lockings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for lockings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lockingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no lockingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from lockings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lockings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q lockingQuery) DeleteAllSlice(o LockingSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lockingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"lockings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lockingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from locking slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for lockings")
	}

	return rowsAff, nil
}

type LockingReloader interface {
	Reload(o *Locking, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *LockingSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q lockingQuery) Reload(o *Locking, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLocking(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q lockingQuery) ReloadAll(o *LockingSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LockingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lockingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"lockings\".* FROM \"lockings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, lockingPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LockingSlice")
	}

	*o = slice

	return nil
}

// LockingExists checks if the Locking row exists.
func LockingExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"lockings\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if lockings exists")
	}

	return exists, nil
}
