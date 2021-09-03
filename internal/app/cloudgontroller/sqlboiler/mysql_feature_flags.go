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

// FeatureFlag is an object representing the database table.
type FeatureFlag struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name         string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Enabled      bool        `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	ErrorMessage null.String `boil:"error_message" json:"error_message,omitempty" toml:"error_message" yaml:"error_message,omitempty"`

	R *featureFlagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L featureFlagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FeatureFlagColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	Name         string
	Enabled      string
	ErrorMessage string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	Name:         "name",
	Enabled:      "enabled",
	ErrorMessage: "error_message",
}

var FeatureFlagTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	Name         string
	Enabled      string
	ErrorMessage string
}{
	ID:           "feature_flags.id",
	GUID:         "feature_flags.guid",
	CreatedAt:    "feature_flags.created_at",
	UpdatedAt:    "feature_flags.updated_at",
	Name:         "feature_flags.name",
	Enabled:      "feature_flags.enabled",
	ErrorMessage: "feature_flags.error_message",
}

// Generated where

var FeatureFlagWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	Name         whereHelperstring
	Enabled      whereHelperbool
	ErrorMessage whereHelpernull_String
}{
	ID:           whereHelperint{field: "`feature_flags`.`id`"},
	GUID:         whereHelperstring{field: "`feature_flags`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`feature_flags`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`feature_flags`.`updated_at`"},
	Name:         whereHelperstring{field: "`feature_flags`.`name`"},
	Enabled:      whereHelperbool{field: "`feature_flags`.`enabled`"},
	ErrorMessage: whereHelpernull_String{field: "`feature_flags`.`error_message`"},
}

// FeatureFlagRels is where relationship names are stored.
var FeatureFlagRels = struct {
}{}

// featureFlagR is where relationships are stored.
type featureFlagR struct {
}

// NewStruct creates a new relationship struct
func (*featureFlagR) NewStruct() *featureFlagR {
	return &featureFlagR{}
}

// featureFlagL is where Load methods for each relationship are stored.
type featureFlagL struct{}

var (
	featureFlagAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "enabled", "error_message"}
	featureFlagColumnsWithoutDefault = []string{"guid", "updated_at", "name", "enabled", "error_message"}
	featureFlagColumnsWithDefault    = []string{"id", "created_at"}
	featureFlagPrimaryKeyColumns     = []string{"id"}
)

type (
	// FeatureFlagSlice is an alias for a slice of pointers to FeatureFlag.
	// This should almost always be used instead of []FeatureFlag.
	FeatureFlagSlice []*FeatureFlag

	featureFlagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	featureFlagType                 = reflect.TypeOf(&FeatureFlag{})
	featureFlagMapping              = queries.MakeStructMapping(featureFlagType)
	featureFlagPrimaryKeyMapping, _ = queries.BindMapping(featureFlagType, featureFlagMapping, featureFlagPrimaryKeyColumns)
	featureFlagInsertCacheMut       sync.RWMutex
	featureFlagInsertCache          = make(map[string]insertCache)
	featureFlagUpdateCacheMut       sync.RWMutex
	featureFlagUpdateCache          = make(map[string]updateCache)
	featureFlagUpsertCacheMut       sync.RWMutex
	featureFlagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type FeatureFlagFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*FeatureFlag, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (FeatureFlagSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single featureFlag record from the query.
func (q featureFlagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FeatureFlag, error) {
	o := &FeatureFlag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for feature_flags")
	}

	return o, nil
}

// All returns all FeatureFlag records from the query.
func (q featureFlagQuery) All(ctx context.Context, exec boil.ContextExecutor) (FeatureFlagSlice, error) {
	var o []*FeatureFlag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FeatureFlag slice")
	}

	return o, nil
}

// Count returns the count of all FeatureFlag records in the query.
func (q featureFlagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count feature_flags rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q featureFlagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if feature_flags exists")
	}

	return count > 0, nil
}

// FeatureFlags retrieves all the records using an executor.
func FeatureFlags(mods ...qm.QueryMod) featureFlagQuery {
	mods = append(mods, qm.From("`feature_flags`"))
	return featureFlagQuery{NewQuery(mods...)}
}

type FeatureFlagFinder interface {
	FindFeatureFlag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*FeatureFlag, error)
}

// FindFeatureFlag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFeatureFlag(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*FeatureFlag, error) {
	featureFlagObj := &FeatureFlag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `feature_flags` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, featureFlagObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from feature_flags")
	}

	return featureFlagObj, nil
}

type FeatureFlagInserter interface {
	Insert(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q featureFlagQuery) Insert(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feature_flags provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(featureFlagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	featureFlagInsertCacheMut.RLock()
	cache, cached := featureFlagInsertCache[key]
	featureFlagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			featureFlagAllColumns,
			featureFlagColumnsWithDefault,
			featureFlagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(featureFlagType, featureFlagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(featureFlagType, featureFlagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `feature_flags` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `feature_flags` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `feature_flags` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, featureFlagPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into feature_flags")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == featureFlagMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for feature_flags")
	}

CacheNoHooks:
	if !cached {
		featureFlagInsertCacheMut.Lock()
		featureFlagInsertCache[key] = cache
		featureFlagInsertCacheMut.Unlock()
	}

	return nil
}

type FeatureFlagUpdater interface {
	Update(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the FeatureFlag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q featureFlagQuery) Update(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	featureFlagUpdateCacheMut.RLock()
	cache, cached := featureFlagUpdateCache[key]
	featureFlagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			featureFlagAllColumns,
			featureFlagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update feature_flags, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `feature_flags` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, featureFlagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(featureFlagType, featureFlagMapping, append(wl, featureFlagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update feature_flags row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for feature_flags")
	}

	if !cached {
		featureFlagUpdateCacheMut.Lock()
		featureFlagUpdateCache[key] = cache
		featureFlagUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q featureFlagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for feature_flags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for feature_flags")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q featureFlagQuery) UpdateAllSlice(o FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureFlagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `feature_flags` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, featureFlagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in featureFlag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all featureFlag")
	}
	return rowsAff, nil
}

type FeatureFlagUpserter interface {
	Upsert(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLFeatureFlagUniqueColumns = []string{
	"id",
	"guid",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q featureFlagQuery) Upsert(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no feature_flags provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(featureFlagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFeatureFlagUniqueColumns, o)

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

	featureFlagUpsertCacheMut.RLock()
	cache, cached := featureFlagUpsertCache[key]
	featureFlagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			featureFlagAllColumns,
			featureFlagColumnsWithDefault,
			featureFlagColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			featureFlagAllColumns,
			featureFlagPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert feature_flags, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`feature_flags`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `feature_flags` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(featureFlagType, featureFlagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(featureFlagType, featureFlagMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for feature_flags")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == featureFlagMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(featureFlagType, featureFlagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for feature_flags")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for feature_flags")
	}

CacheNoHooks:
	if !cached {
		featureFlagUpsertCacheMut.Lock()
		featureFlagUpsertCache[key] = cache
		featureFlagUpsertCacheMut.Unlock()
	}

	return nil
}

type FeatureFlagDeleter interface {
	Delete(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single FeatureFlag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q featureFlagQuery) Delete(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FeatureFlag provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), featureFlagPrimaryKeyMapping)
	sql := "DELETE FROM `feature_flags` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from feature_flags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for feature_flags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q featureFlagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no featureFlagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from feature_flags")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feature_flags")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q featureFlagQuery) DeleteAllSlice(o FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureFlagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `feature_flags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, featureFlagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from featureFlag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for feature_flags")
	}

	return rowsAff, nil
}

type FeatureFlagReloader interface {
	Reload(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q featureFlagQuery) Reload(o *FeatureFlag, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFeatureFlag(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q featureFlagQuery) ReloadAll(o *FeatureFlagSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FeatureFlagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), featureFlagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `feature_flags`.* FROM `feature_flags` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, featureFlagPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FeatureFlagSlice")
	}

	*o = slice

	return nil
}

// FeatureFlagExists checks if the FeatureFlag row exists.
func FeatureFlagExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `feature_flags` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if feature_flags exists")
	}

	return exists, nil
}
