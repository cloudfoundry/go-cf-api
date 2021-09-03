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

// IsolationSegmentLabel is an object representing the database table.
type IsolationSegmentLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *isolationSegmentLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L isolationSegmentLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IsolationSegmentLabelColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ResourceGUID: "resource_guid",
	KeyPrefix:    "key_prefix",
	KeyName:      "key_name",
	Value:        "value",
}

var IsolationSegmentLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "isolation_segment_labels.id",
	GUID:         "isolation_segment_labels.guid",
	CreatedAt:    "isolation_segment_labels.created_at",
	UpdatedAt:    "isolation_segment_labels.updated_at",
	ResourceGUID: "isolation_segment_labels.resource_guid",
	KeyPrefix:    "isolation_segment_labels.key_prefix",
	KeyName:      "isolation_segment_labels.key_name",
	Value:        "isolation_segment_labels.value",
}

// Generated where

var IsolationSegmentLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`isolation_segment_labels`.`id`"},
	GUID:         whereHelperstring{field: "`isolation_segment_labels`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`isolation_segment_labels`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`isolation_segment_labels`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`isolation_segment_labels`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`isolation_segment_labels`.`key_prefix`"},
	KeyName:      whereHelpernull_String{field: "`isolation_segment_labels`.`key_name`"},
	Value:        whereHelpernull_String{field: "`isolation_segment_labels`.`value`"},
}

// IsolationSegmentLabelRels is where relationship names are stored.
var IsolationSegmentLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// isolationSegmentLabelR is where relationships are stored.
type isolationSegmentLabelR struct {
	Resource *IsolationSegment `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*isolationSegmentLabelR) NewStruct() *isolationSegmentLabelR {
	return &isolationSegmentLabelR{}
}

// isolationSegmentLabelL is where Load methods for each relationship are stored.
type isolationSegmentLabelL struct{}

var (
	isolationSegmentLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	isolationSegmentLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	isolationSegmentLabelColumnsWithDefault    = []string{"id", "created_at"}
	isolationSegmentLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// IsolationSegmentLabelSlice is an alias for a slice of pointers to IsolationSegmentLabel.
	// This should almost always be used instead of []IsolationSegmentLabel.
	IsolationSegmentLabelSlice []*IsolationSegmentLabel

	isolationSegmentLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	isolationSegmentLabelType                 = reflect.TypeOf(&IsolationSegmentLabel{})
	isolationSegmentLabelMapping              = queries.MakeStructMapping(isolationSegmentLabelType)
	isolationSegmentLabelPrimaryKeyMapping, _ = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, isolationSegmentLabelPrimaryKeyColumns)
	isolationSegmentLabelInsertCacheMut       sync.RWMutex
	isolationSegmentLabelInsertCache          = make(map[string]insertCache)
	isolationSegmentLabelUpdateCacheMut       sync.RWMutex
	isolationSegmentLabelUpdateCache          = make(map[string]updateCache)
	isolationSegmentLabelUpsertCacheMut       sync.RWMutex
	isolationSegmentLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type IsolationSegmentLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*IsolationSegmentLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (IsolationSegmentLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single isolationSegmentLabel record from the query.
func (q isolationSegmentLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*IsolationSegmentLabel, error) {
	o := &IsolationSegmentLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for isolation_segment_labels")
	}

	return o, nil
}

// All returns all IsolationSegmentLabel records from the query.
func (q isolationSegmentLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (IsolationSegmentLabelSlice, error) {
	var o []*IsolationSegmentLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to IsolationSegmentLabel slice")
	}

	return o, nil
}

// Count returns the count of all IsolationSegmentLabel records in the query.
func (q isolationSegmentLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count isolation_segment_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q isolationSegmentLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if isolation_segment_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q isolationSegmentLabelQuery) Resource(o *IsolationSegmentLabel, mods ...qm.QueryMod) isolationSegmentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := IsolationSegments(queryMods...)
	queries.SetFrom(query.Query, "`isolation_segments`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (isolationSegmentLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeIsolationSegmentLabel interface{}, mods queries.Applicator) error {
	var slice []*IsolationSegmentLabel
	var object *IsolationSegmentLabel

	if singular {
		object = maybeIsolationSegmentLabel.(*IsolationSegmentLabel)
	} else {
		slice = *maybeIsolationSegmentLabel.(*[]*IsolationSegmentLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &isolationSegmentLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &isolationSegmentLabelR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ResourceGUID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ResourceGUID) {
				args = append(args, obj.ResourceGUID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`isolation_segments`),
		qm.WhereIn(`isolation_segments.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load IsolationSegment")
	}

	var resultSlice []*IsolationSegment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice IsolationSegment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for isolation_segments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for isolation_segments")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &isolationSegmentR{}
		}
		foreign.R.ResourceIsolationSegmentLabels = append(foreign.R.ResourceIsolationSegmentLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &isolationSegmentR{}
				}
				foreign.R.ResourceIsolationSegmentLabels = append(foreign.R.ResourceIsolationSegmentLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the isolationSegmentLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceIsolationSegmentLabels.
func (q isolationSegmentLabelQuery) SetResource(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *IsolationSegment) error {
	var err error
	if insert {
		if err = IsolationSegments().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `isolation_segment_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, isolationSegmentLabelPrimaryKeyColumns),
	)
	values := []interface{}{related.GUID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ResourceGUID, related.GUID)
	if o.R == nil {
		o.R = &isolationSegmentLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &isolationSegmentR{
			ResourceIsolationSegmentLabels: IsolationSegmentLabelSlice{o},
		}
	} else {
		related.R.ResourceIsolationSegmentLabels = append(related.R.ResourceIsolationSegmentLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q isolationSegmentLabelQuery) RemoveResource(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, related *IsolationSegment) error {
	var err error

	queries.SetScanner(&o.ResourceGUID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Resource = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ResourceIsolationSegmentLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceIsolationSegmentLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceIsolationSegmentLabels[i] = related.R.ResourceIsolationSegmentLabels[ln-1]
		}
		related.R.ResourceIsolationSegmentLabels = related.R.ResourceIsolationSegmentLabels[:ln-1]
		break
	}
	return nil
}

// IsolationSegmentLabels retrieves all the records using an executor.
func IsolationSegmentLabels(mods ...qm.QueryMod) isolationSegmentLabelQuery {
	mods = append(mods, qm.From("`isolation_segment_labels`"))
	return isolationSegmentLabelQuery{NewQuery(mods...)}
}

type IsolationSegmentLabelFinder interface {
	FindIsolationSegmentLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*IsolationSegmentLabel, error)
}

// FindIsolationSegmentLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIsolationSegmentLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*IsolationSegmentLabel, error) {
	isolationSegmentLabelObj := &IsolationSegmentLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `isolation_segment_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, isolationSegmentLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from isolation_segment_labels")
	}

	return isolationSegmentLabelObj, nil
}

type IsolationSegmentLabelInserter interface {
	Insert(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q isolationSegmentLabelQuery) Insert(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no isolation_segment_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(isolationSegmentLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	isolationSegmentLabelInsertCacheMut.RLock()
	cache, cached := isolationSegmentLabelInsertCache[key]
	isolationSegmentLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			isolationSegmentLabelAllColumns,
			isolationSegmentLabelColumnsWithDefault,
			isolationSegmentLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `isolation_segment_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `isolation_segment_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `isolation_segment_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, isolationSegmentLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into isolation_segment_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == isolationSegmentLabelMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for isolation_segment_labels")
	}

CacheNoHooks:
	if !cached {
		isolationSegmentLabelInsertCacheMut.Lock()
		isolationSegmentLabelInsertCache[key] = cache
		isolationSegmentLabelInsertCacheMut.Unlock()
	}

	return nil
}

type IsolationSegmentLabelUpdater interface {
	Update(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the IsolationSegmentLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q isolationSegmentLabelQuery) Update(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	isolationSegmentLabelUpdateCacheMut.RLock()
	cache, cached := isolationSegmentLabelUpdateCache[key]
	isolationSegmentLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			isolationSegmentLabelAllColumns,
			isolationSegmentLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update isolation_segment_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `isolation_segment_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, isolationSegmentLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, append(wl, isolationSegmentLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update isolation_segment_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for isolation_segment_labels")
	}

	if !cached {
		isolationSegmentLabelUpdateCacheMut.Lock()
		isolationSegmentLabelUpdateCache[key] = cache
		isolationSegmentLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q isolationSegmentLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for isolation_segment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for isolation_segment_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q isolationSegmentLabelQuery) UpdateAllSlice(o IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), isolationSegmentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `isolation_segment_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, isolationSegmentLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in isolationSegmentLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all isolationSegmentLabel")
	}
	return rowsAff, nil
}

type IsolationSegmentLabelUpserter interface {
	Upsert(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLIsolationSegmentLabelUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q isolationSegmentLabelQuery) Upsert(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no isolation_segment_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(isolationSegmentLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLIsolationSegmentLabelUniqueColumns, o)

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

	isolationSegmentLabelUpsertCacheMut.RLock()
	cache, cached := isolationSegmentLabelUpsertCache[key]
	isolationSegmentLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			isolationSegmentLabelAllColumns,
			isolationSegmentLabelColumnsWithDefault,
			isolationSegmentLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			isolationSegmentLabelAllColumns,
			isolationSegmentLabelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert isolation_segment_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`isolation_segment_labels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `isolation_segment_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for isolation_segment_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == isolationSegmentLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(isolationSegmentLabelType, isolationSegmentLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for isolation_segment_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for isolation_segment_labels")
	}

CacheNoHooks:
	if !cached {
		isolationSegmentLabelUpsertCacheMut.Lock()
		isolationSegmentLabelUpsertCache[key] = cache
		isolationSegmentLabelUpsertCacheMut.Unlock()
	}

	return nil
}

type IsolationSegmentLabelDeleter interface {
	Delete(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single IsolationSegmentLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q isolationSegmentLabelQuery) Delete(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no IsolationSegmentLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), isolationSegmentLabelPrimaryKeyMapping)
	sql := "DELETE FROM `isolation_segment_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from isolation_segment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for isolation_segment_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q isolationSegmentLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no isolationSegmentLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from isolation_segment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for isolation_segment_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q isolationSegmentLabelQuery) DeleteAllSlice(o IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), isolationSegmentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `isolation_segment_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, isolationSegmentLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from isolationSegmentLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for isolation_segment_labels")
	}

	return rowsAff, nil
}

type IsolationSegmentLabelReloader interface {
	Reload(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q isolationSegmentLabelQuery) Reload(o *IsolationSegmentLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIsolationSegmentLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q isolationSegmentLabelQuery) ReloadAll(o *IsolationSegmentLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IsolationSegmentLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), isolationSegmentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `isolation_segment_labels`.* FROM `isolation_segment_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, isolationSegmentLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in IsolationSegmentLabelSlice")
	}

	*o = slice

	return nil
}

// IsolationSegmentLabelExists checks if the IsolationSegmentLabel row exists.
func IsolationSegmentLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `isolation_segment_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if isolation_segment_labels exists")
	}

	return exists, nil
}
