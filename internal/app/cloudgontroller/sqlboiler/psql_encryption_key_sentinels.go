//go:generate mockgen -source=$GOFILE -destination=mocks/encryption_key_sentinels.go
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

// EncryptionKeySentinel is an object representing the database table.
type EncryptionKeySentinel struct {
	ID                   int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                 string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt            time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt            null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ExpectedValue        null.String `boil:"expected_value" json:"expected_value,omitempty" toml:"expected_value" yaml:"expected_value,omitempty"`
	EncryptedValue       null.String `boil:"encrypted_value" json:"encrypted_value,omitempty" toml:"encrypted_value" yaml:"encrypted_value,omitempty"`
	EncryptionKeyLabel   null.String `boil:"encryption_key_label" json:"encryption_key_label,omitempty" toml:"encryption_key_label" yaml:"encryption_key_label,omitempty"`
	Salt                 null.String `boil:"salt" json:"salt,omitempty" toml:"salt" yaml:"salt,omitempty"`
	EncryptionIterations int         `boil:"encryption_iterations" json:"encryption_iterations" toml:"encryption_iterations" yaml:"encryption_iterations"`

	R *encryptionKeySentinelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L encryptionKeySentinelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EncryptionKeySentinelColumns = struct {
	ID                   string
	GUID                 string
	CreatedAt            string
	UpdatedAt            string
	ExpectedValue        string
	EncryptedValue       string
	EncryptionKeyLabel   string
	Salt                 string
	EncryptionIterations string
}{
	ID:                   "id",
	GUID:                 "guid",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	ExpectedValue:        "expected_value",
	EncryptedValue:       "encrypted_value",
	EncryptionKeyLabel:   "encryption_key_label",
	Salt:                 "salt",
	EncryptionIterations: "encryption_iterations",
}

var EncryptionKeySentinelTableColumns = struct {
	ID                   string
	GUID                 string
	CreatedAt            string
	UpdatedAt            string
	ExpectedValue        string
	EncryptedValue       string
	EncryptionKeyLabel   string
	Salt                 string
	EncryptionIterations string
}{
	ID:                   "encryption_key_sentinels.id",
	GUID:                 "encryption_key_sentinels.guid",
	CreatedAt:            "encryption_key_sentinels.created_at",
	UpdatedAt:            "encryption_key_sentinels.updated_at",
	ExpectedValue:        "encryption_key_sentinels.expected_value",
	EncryptedValue:       "encryption_key_sentinels.encrypted_value",
	EncryptionKeyLabel:   "encryption_key_sentinels.encryption_key_label",
	Salt:                 "encryption_key_sentinels.salt",
	EncryptionIterations: "encryption_key_sentinels.encryption_iterations",
}

// Generated where

var EncryptionKeySentinelWhere = struct {
	ID                   whereHelperint
	GUID                 whereHelperstring
	CreatedAt            whereHelpertime_Time
	UpdatedAt            whereHelpernull_Time
	ExpectedValue        whereHelpernull_String
	EncryptedValue       whereHelpernull_String
	EncryptionKeyLabel   whereHelpernull_String
	Salt                 whereHelpernull_String
	EncryptionIterations whereHelperint
}{
	ID:                   whereHelperint{field: "\"encryption_key_sentinels\".\"id\""},
	GUID:                 whereHelperstring{field: "\"encryption_key_sentinels\".\"guid\""},
	CreatedAt:            whereHelpertime_Time{field: "\"encryption_key_sentinels\".\"created_at\""},
	UpdatedAt:            whereHelpernull_Time{field: "\"encryption_key_sentinels\".\"updated_at\""},
	ExpectedValue:        whereHelpernull_String{field: "\"encryption_key_sentinels\".\"expected_value\""},
	EncryptedValue:       whereHelpernull_String{field: "\"encryption_key_sentinels\".\"encrypted_value\""},
	EncryptionKeyLabel:   whereHelpernull_String{field: "\"encryption_key_sentinels\".\"encryption_key_label\""},
	Salt:                 whereHelpernull_String{field: "\"encryption_key_sentinels\".\"salt\""},
	EncryptionIterations: whereHelperint{field: "\"encryption_key_sentinels\".\"encryption_iterations\""},
}

// EncryptionKeySentinelRels is where relationship names are stored.
var EncryptionKeySentinelRels = struct {
}{}

// encryptionKeySentinelR is where relationships are stored.
type encryptionKeySentinelR struct {
}

// NewStruct creates a new relationship struct
func (*encryptionKeySentinelR) NewStruct() *encryptionKeySentinelR {
	return &encryptionKeySentinelR{}
}

// encryptionKeySentinelL is where Load methods for each relationship are stored.
type encryptionKeySentinelL struct{}

var (
	encryptionKeySentinelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "expected_value", "encrypted_value", "encryption_key_label", "salt", "encryption_iterations"}
	encryptionKeySentinelColumnsWithoutDefault = []string{"guid", "updated_at", "expected_value", "encrypted_value", "encryption_key_label", "salt"}
	encryptionKeySentinelColumnsWithDefault    = []string{"id", "created_at", "encryption_iterations"}
	encryptionKeySentinelPrimaryKeyColumns     = []string{"id"}
)

type (
	// EncryptionKeySentinelSlice is an alias for a slice of pointers to EncryptionKeySentinel.
	// This should almost always be used instead of []EncryptionKeySentinel.
	EncryptionKeySentinelSlice []*EncryptionKeySentinel

	encryptionKeySentinelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	encryptionKeySentinelType                 = reflect.TypeOf(&EncryptionKeySentinel{})
	encryptionKeySentinelMapping              = queries.MakeStructMapping(encryptionKeySentinelType)
	encryptionKeySentinelPrimaryKeyMapping, _ = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, encryptionKeySentinelPrimaryKeyColumns)
	encryptionKeySentinelInsertCacheMut       sync.RWMutex
	encryptionKeySentinelInsertCache          = make(map[string]insertCache)
	encryptionKeySentinelUpdateCacheMut       sync.RWMutex
	encryptionKeySentinelUpdateCache          = make(map[string]updateCache)
	encryptionKeySentinelUpsertCacheMut       sync.RWMutex
	encryptionKeySentinelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type EncryptionKeySentinelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*EncryptionKeySentinel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (EncryptionKeySentinelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single encryptionKeySentinel record from the query.
func (q encryptionKeySentinelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EncryptionKeySentinel, error) {
	o := &EncryptionKeySentinel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for encryption_key_sentinels")
	}

	return o, nil
}

// All returns all EncryptionKeySentinel records from the query.
func (q encryptionKeySentinelQuery) All(ctx context.Context, exec boil.ContextExecutor) (EncryptionKeySentinelSlice, error) {
	var o []*EncryptionKeySentinel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EncryptionKeySentinel slice")
	}

	return o, nil
}

// Count returns the count of all EncryptionKeySentinel records in the query.
func (q encryptionKeySentinelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count encryption_key_sentinels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q encryptionKeySentinelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if encryption_key_sentinels exists")
	}

	return count > 0, nil
}

// EncryptionKeySentinels retrieves all the records using an executor.
func EncryptionKeySentinels(mods ...qm.QueryMod) encryptionKeySentinelQuery {
	mods = append(mods, qm.From("\"encryption_key_sentinels\""))
	return encryptionKeySentinelQuery{NewQuery(mods...)}
}

type EncryptionKeySentinelFinder interface {
	FindEncryptionKeySentinel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*EncryptionKeySentinel, error)
}

// FindEncryptionKeySentinel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEncryptionKeySentinel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*EncryptionKeySentinel, error) {
	encryptionKeySentinelObj := &EncryptionKeySentinel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"encryption_key_sentinels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, encryptionKeySentinelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from encryption_key_sentinels")
	}

	return encryptionKeySentinelObj, nil
}

type EncryptionKeySentinelInserter interface {
	Insert(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q encryptionKeySentinelQuery) Insert(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no encryption_key_sentinels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(encryptionKeySentinelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	encryptionKeySentinelInsertCacheMut.RLock()
	cache, cached := encryptionKeySentinelInsertCache[key]
	encryptionKeySentinelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			encryptionKeySentinelAllColumns,
			encryptionKeySentinelColumnsWithDefault,
			encryptionKeySentinelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"encryption_key_sentinels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"encryption_key_sentinels\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into encryption_key_sentinels")
	}

	if !cached {
		encryptionKeySentinelInsertCacheMut.Lock()
		encryptionKeySentinelInsertCache[key] = cache
		encryptionKeySentinelInsertCacheMut.Unlock()
	}

	return nil
}

type EncryptionKeySentinelUpdater interface {
	Update(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the EncryptionKeySentinel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q encryptionKeySentinelQuery) Update(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	encryptionKeySentinelUpdateCacheMut.RLock()
	cache, cached := encryptionKeySentinelUpdateCache[key]
	encryptionKeySentinelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			encryptionKeySentinelAllColumns,
			encryptionKeySentinelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update encryption_key_sentinels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"encryption_key_sentinels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, encryptionKeySentinelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, append(wl, encryptionKeySentinelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update encryption_key_sentinels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for encryption_key_sentinels")
	}

	if !cached {
		encryptionKeySentinelUpdateCacheMut.Lock()
		encryptionKeySentinelUpdateCache[key] = cache
		encryptionKeySentinelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q encryptionKeySentinelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for encryption_key_sentinels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for encryption_key_sentinels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q encryptionKeySentinelQuery) UpdateAllSlice(o EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), encryptionKeySentinelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"encryption_key_sentinels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, encryptionKeySentinelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in encryptionKeySentinel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all encryptionKeySentinel")
	}
	return rowsAff, nil
}

type EncryptionKeySentinelUpserter interface {
	Upsert(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q encryptionKeySentinelQuery) Upsert(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no encryption_key_sentinels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(encryptionKeySentinelColumnsWithDefault, o)

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

	encryptionKeySentinelUpsertCacheMut.RLock()
	cache, cached := encryptionKeySentinelUpsertCache[key]
	encryptionKeySentinelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			encryptionKeySentinelAllColumns,
			encryptionKeySentinelColumnsWithDefault,
			encryptionKeySentinelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			encryptionKeySentinelAllColumns,
			encryptionKeySentinelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert encryption_key_sentinels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(encryptionKeySentinelPrimaryKeyColumns))
			copy(conflict, encryptionKeySentinelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"encryption_key_sentinels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(encryptionKeySentinelType, encryptionKeySentinelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert encryption_key_sentinels")
	}

	if !cached {
		encryptionKeySentinelUpsertCacheMut.Lock()
		encryptionKeySentinelUpsertCache[key] = cache
		encryptionKeySentinelUpsertCacheMut.Unlock()
	}

	return nil
}

type EncryptionKeySentinelDeleter interface {
	Delete(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single EncryptionKeySentinel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q encryptionKeySentinelQuery) Delete(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EncryptionKeySentinel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), encryptionKeySentinelPrimaryKeyMapping)
	sql := "DELETE FROM \"encryption_key_sentinels\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from encryption_key_sentinels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for encryption_key_sentinels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q encryptionKeySentinelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no encryptionKeySentinelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from encryption_key_sentinels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for encryption_key_sentinels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q encryptionKeySentinelQuery) DeleteAllSlice(o EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), encryptionKeySentinelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"encryption_key_sentinels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, encryptionKeySentinelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from encryptionKeySentinel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for encryption_key_sentinels")
	}

	return rowsAff, nil
}

type EncryptionKeySentinelReloader interface {
	Reload(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q encryptionKeySentinelQuery) Reload(o *EncryptionKeySentinel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEncryptionKeySentinel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q encryptionKeySentinelQuery) ReloadAll(o *EncryptionKeySentinelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EncryptionKeySentinelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), encryptionKeySentinelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"encryption_key_sentinels\".* FROM \"encryption_key_sentinels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, encryptionKeySentinelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EncryptionKeySentinelSlice")
	}

	*o = slice

	return nil
}

// EncryptionKeySentinelExists checks if the EncryptionKeySentinel row exists.
func EncryptionKeySentinelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"encryption_key_sentinels\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if encryption_key_sentinels exists")
	}

	return exists, nil
}
