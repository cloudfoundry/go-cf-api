// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/orphaned_blobs.go
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

// OrphanedBlob is an object representing the database table.
type OrphanedBlob struct {
	ID            int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID          string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt     null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	BlobKey       null.String `boil:"blob_key" json:"blob_key,omitempty" toml:"blob_key" yaml:"blob_key,omitempty"`
	DirtyCount    null.Int    `boil:"dirty_count" json:"dirty_count,omitempty" toml:"dirty_count" yaml:"dirty_count,omitempty"`
	BlobstoreType null.String `boil:"blobstore_type" json:"blobstore_type,omitempty" toml:"blobstore_type" yaml:"blobstore_type,omitempty"`

	R *orphanedBlobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orphanedBlobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrphanedBlobColumns = struct {
	ID            string
	GUID          string
	CreatedAt     string
	UpdatedAt     string
	BlobKey       string
	DirtyCount    string
	BlobstoreType string
}{
	ID:            "id",
	GUID:          "guid",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	BlobKey:       "blob_key",
	DirtyCount:    "dirty_count",
	BlobstoreType: "blobstore_type",
}

var OrphanedBlobTableColumns = struct {
	ID            string
	GUID          string
	CreatedAt     string
	UpdatedAt     string
	BlobKey       string
	DirtyCount    string
	BlobstoreType string
}{
	ID:            "orphaned_blobs.id",
	GUID:          "orphaned_blobs.guid",
	CreatedAt:     "orphaned_blobs.created_at",
	UpdatedAt:     "orphaned_blobs.updated_at",
	BlobKey:       "orphaned_blobs.blob_key",
	DirtyCount:    "orphaned_blobs.dirty_count",
	BlobstoreType: "orphaned_blobs.blobstore_type",
}

// Generated where

var OrphanedBlobWhere = struct {
	ID            whereHelperint
	GUID          whereHelperstring
	CreatedAt     whereHelpertime_Time
	UpdatedAt     whereHelpernull_Time
	BlobKey       whereHelpernull_String
	DirtyCount    whereHelpernull_Int
	BlobstoreType whereHelpernull_String
}{
	ID:            whereHelperint{field: "\"orphaned_blobs\".\"id\""},
	GUID:          whereHelperstring{field: "\"orphaned_blobs\".\"guid\""},
	CreatedAt:     whereHelpertime_Time{field: "\"orphaned_blobs\".\"created_at\""},
	UpdatedAt:     whereHelpernull_Time{field: "\"orphaned_blobs\".\"updated_at\""},
	BlobKey:       whereHelpernull_String{field: "\"orphaned_blobs\".\"blob_key\""},
	DirtyCount:    whereHelpernull_Int{field: "\"orphaned_blobs\".\"dirty_count\""},
	BlobstoreType: whereHelpernull_String{field: "\"orphaned_blobs\".\"blobstore_type\""},
}

// OrphanedBlobRels is where relationship names are stored.
var OrphanedBlobRels = struct {
}{}

// orphanedBlobR is where relationships are stored.
type orphanedBlobR struct {
}

// NewStruct creates a new relationship struct
func (*orphanedBlobR) NewStruct() *orphanedBlobR {
	return &orphanedBlobR{}
}

// orphanedBlobL is where Load methods for each relationship are stored.
type orphanedBlobL struct{}

var (
	orphanedBlobAllColumns            = []string{"id", "guid", "created_at", "updated_at", "blob_key", "dirty_count", "blobstore_type"}
	orphanedBlobColumnsWithoutDefault = []string{"guid", "updated_at", "blob_key", "dirty_count", "blobstore_type"}
	orphanedBlobColumnsWithDefault    = []string{"id", "created_at"}
	orphanedBlobPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrphanedBlobSlice is an alias for a slice of pointers to OrphanedBlob.
	// This should almost always be used instead of []OrphanedBlob.
	OrphanedBlobSlice []*OrphanedBlob

	orphanedBlobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orphanedBlobType                 = reflect.TypeOf(&OrphanedBlob{})
	orphanedBlobMapping              = queries.MakeStructMapping(orphanedBlobType)
	orphanedBlobPrimaryKeyMapping, _ = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, orphanedBlobPrimaryKeyColumns)
	orphanedBlobInsertCacheMut       sync.RWMutex
	orphanedBlobInsertCache          = make(map[string]insertCache)
	orphanedBlobUpdateCacheMut       sync.RWMutex
	orphanedBlobUpdateCache          = make(map[string]updateCache)
	orphanedBlobUpsertCacheMut       sync.RWMutex
	orphanedBlobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrphanedBlobFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrphanedBlob, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrphanedBlobSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single orphanedBlob record from the query.
func (q orphanedBlobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrphanedBlob, error) {
	o := &OrphanedBlob{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for orphaned_blobs")
	}

	return o, nil
}

// All returns all OrphanedBlob records from the query.
func (q orphanedBlobQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrphanedBlobSlice, error) {
	var o []*OrphanedBlob

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrphanedBlob slice")
	}

	return o, nil
}

// Count returns the count of all OrphanedBlob records in the query.
func (q orphanedBlobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count orphaned_blobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orphanedBlobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if orphaned_blobs exists")
	}

	return count > 0, nil
}

// OrphanedBlobs retrieves all the records using an executor.
func OrphanedBlobs(mods ...qm.QueryMod) orphanedBlobQuery {
	mods = append(mods, qm.From("\"orphaned_blobs\""))
	return orphanedBlobQuery{NewQuery(mods...)}
}

type OrphanedBlobFinder interface {
	FindOrphanedBlob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrphanedBlob, error)
}

// FindOrphanedBlob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrphanedBlob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrphanedBlob, error) {
	orphanedBlobObj := &OrphanedBlob{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"orphaned_blobs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orphanedBlobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from orphaned_blobs")
	}

	return orphanedBlobObj, nil
}

type OrphanedBlobInserter interface {
	Insert(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q orphanedBlobQuery) Insert(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orphaned_blobs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(orphanedBlobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orphanedBlobInsertCacheMut.RLock()
	cache, cached := orphanedBlobInsertCache[key]
	orphanedBlobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orphanedBlobAllColumns,
			orphanedBlobColumnsWithDefault,
			orphanedBlobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"orphaned_blobs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"orphaned_blobs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into orphaned_blobs")
	}

	if !cached {
		orphanedBlobInsertCacheMut.Lock()
		orphanedBlobInsertCache[key] = cache
		orphanedBlobInsertCacheMut.Unlock()
	}

	return nil
}

type OrphanedBlobUpdater interface {
	Update(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrphanedBlob.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q orphanedBlobQuery) Update(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	orphanedBlobUpdateCacheMut.RLock()
	cache, cached := orphanedBlobUpdateCache[key]
	orphanedBlobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orphanedBlobAllColumns,
			orphanedBlobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update orphaned_blobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"orphaned_blobs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orphanedBlobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, append(wl, orphanedBlobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update orphaned_blobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for orphaned_blobs")
	}

	if !cached {
		orphanedBlobUpdateCacheMut.Lock()
		orphanedBlobUpdateCache[key] = cache
		orphanedBlobUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q orphanedBlobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for orphaned_blobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for orphaned_blobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q orphanedBlobQuery) UpdateAllSlice(o OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orphanedBlobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"orphaned_blobs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orphanedBlobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in orphanedBlob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all orphanedBlob")
	}
	return rowsAff, nil
}

type OrphanedBlobDeleter interface {
	Delete(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrphanedBlob record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q orphanedBlobQuery) Delete(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrphanedBlob provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orphanedBlobPrimaryKeyMapping)
	sql := "DELETE FROM \"orphaned_blobs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from orphaned_blobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for orphaned_blobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orphanedBlobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orphanedBlobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orphaned_blobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orphaned_blobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q orphanedBlobQuery) DeleteAllSlice(o OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orphanedBlobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"orphaned_blobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orphanedBlobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orphanedBlob slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orphaned_blobs")
	}

	return rowsAff, nil
}

type OrphanedBlobReloader interface {
	Reload(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q orphanedBlobQuery) Reload(o *OrphanedBlob, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrphanedBlob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q orphanedBlobQuery) ReloadAll(o *OrphanedBlobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrphanedBlobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orphanedBlobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"orphaned_blobs\".* FROM \"orphaned_blobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orphanedBlobPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrphanedBlobSlice")
	}

	*o = slice

	return nil
}

// OrphanedBlobExists checks if the OrphanedBlob row exists.
func OrphanedBlobExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"orphaned_blobs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if orphaned_blobs exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrphanedBlob) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orphaned_blobs provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(orphanedBlobColumnsWithDefault, o)

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

	orphanedBlobUpsertCacheMut.RLock()
	cache, cached := orphanedBlobUpsertCache[key]
	orphanedBlobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orphanedBlobAllColumns,
			orphanedBlobColumnsWithDefault,
			orphanedBlobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orphanedBlobAllColumns,
			orphanedBlobPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert orphaned_blobs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orphanedBlobPrimaryKeyColumns))
			copy(conflict, orphanedBlobPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"orphaned_blobs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orphanedBlobType, orphanedBlobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert orphaned_blobs")
	}

	if !cached {
		orphanedBlobUpsertCacheMut.Lock()
		orphanedBlobUpsertCache[key] = cache
		orphanedBlobUpsertCacheMut.Unlock()
	}

	return nil
}
