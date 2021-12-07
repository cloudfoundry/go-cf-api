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

type TaskLabelUpserter interface {
	Upsert(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLTaskLabelUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q taskLabelQuery) Upsert(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no task_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(taskLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTaskLabelUniqueColumns, o)

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

	taskLabelUpsertCacheMut.RLock()
	cache, cached := taskLabelUpsertCache[key]
	taskLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			taskLabelAllColumns,
			taskLabelColumnsWithDefault,
			taskLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			taskLabelAllColumns,
			taskLabelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert task_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`task_labels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `task_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(taskLabelType, taskLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(taskLabelType, taskLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for task_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == taskLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(taskLabelType, taskLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for task_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for task_labels")
	}

CacheNoHooks:
	if !cached {
		taskLabelUpsertCacheMut.Lock()
		taskLabelUpsertCache[key] = cache
		taskLabelUpsertCacheMut.Unlock()
	}

	return nil
}

// TaskLabel is an object representing the database table.
type TaskLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *taskLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L taskLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TaskLabelColumns = struct {
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

var TaskLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "task_labels.id",
	GUID:         "task_labels.guid",
	CreatedAt:    "task_labels.created_at",
	UpdatedAt:    "task_labels.updated_at",
	ResourceGUID: "task_labels.resource_guid",
	KeyPrefix:    "task_labels.key_prefix",
	KeyName:      "task_labels.key_name",
	Value:        "task_labels.value",
}

// Generated where

var TaskLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`task_labels`.`id`"},
	GUID:         whereHelperstring{field: "`task_labels`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`task_labels`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`task_labels`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`task_labels`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`task_labels`.`key_prefix`"},
	KeyName:      whereHelpernull_String{field: "`task_labels`.`key_name`"},
	Value:        whereHelpernull_String{field: "`task_labels`.`value`"},
}

// TaskLabelRels is where relationship names are stored.
var TaskLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// taskLabelR is where relationships are stored.
type taskLabelR struct {
	Resource *Task `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*taskLabelR) NewStruct() *taskLabelR {
	return &taskLabelR{}
}

// taskLabelL is where Load methods for each relationship are stored.
type taskLabelL struct{}

var (
	taskLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	taskLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	taskLabelColumnsWithDefault    = []string{"id", "created_at"}
	taskLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// TaskLabelSlice is an alias for a slice of pointers to TaskLabel.
	// This should almost always be used instead of []TaskLabel.
	TaskLabelSlice []*TaskLabel

	taskLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	taskLabelType                 = reflect.TypeOf(&TaskLabel{})
	taskLabelMapping              = queries.MakeStructMapping(taskLabelType)
	taskLabelPrimaryKeyMapping, _ = queries.BindMapping(taskLabelType, taskLabelMapping, taskLabelPrimaryKeyColumns)
	taskLabelInsertCacheMut       sync.RWMutex
	taskLabelInsertCache          = make(map[string]insertCache)
	taskLabelUpdateCacheMut       sync.RWMutex
	taskLabelUpdateCache          = make(map[string]updateCache)
	taskLabelUpsertCacheMut       sync.RWMutex
	taskLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type TaskLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*TaskLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (TaskLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single taskLabel record from the query.
func (q taskLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TaskLabel, error) {
	o := &TaskLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for task_labels")
	}

	return o, nil
}

// All returns all TaskLabel records from the query.
func (q taskLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (TaskLabelSlice, error) {
	var o []*TaskLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TaskLabel slice")
	}

	return o, nil
}

// Count returns the count of all TaskLabel records in the query.
func (q taskLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count task_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q taskLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if task_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q taskLabelQuery) Resource(o *TaskLabel, mods ...qm.QueryMod) taskQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Tasks(queryMods...)
	queries.SetFrom(query.Query, "`tasks`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (taskLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTaskLabel interface{}, mods queries.Applicator) error {
	var slice []*TaskLabel
	var object *TaskLabel

	if singular {
		object = maybeTaskLabel.(*TaskLabel)
	} else {
		slice = *maybeTaskLabel.(*[]*TaskLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &taskLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &taskLabelR{}
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
		qm.From(`tasks`),
		qm.WhereIn(`tasks.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Task")
	}

	var resultSlice []*Task
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Task")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for tasks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for tasks")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &taskR{}
		}
		foreign.R.ResourceTaskLabels = append(foreign.R.ResourceTaskLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &taskR{}
				}
				foreign.R.ResourceTaskLabels = append(foreign.R.ResourceTaskLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the taskLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceTaskLabels.
func (q taskLabelQuery) SetResource(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Task) error {
	var err error
	if insert {
		if err = Tasks().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `task_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, taskLabelPrimaryKeyColumns),
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
		o.R = &taskLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &taskR{
			ResourceTaskLabels: TaskLabelSlice{o},
		}
	} else {
		related.R.ResourceTaskLabels = append(related.R.ResourceTaskLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q taskLabelQuery) RemoveResource(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, related *Task) error {
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

	for i, ri := range related.R.ResourceTaskLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceTaskLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceTaskLabels[i] = related.R.ResourceTaskLabels[ln-1]
		}
		related.R.ResourceTaskLabels = related.R.ResourceTaskLabels[:ln-1]
		break
	}
	return nil
}

// TaskLabels retrieves all the records using an executor.
func TaskLabels(mods ...qm.QueryMod) taskLabelQuery {
	mods = append(mods, qm.From("`task_labels`"))
	return taskLabelQuery{NewQuery(mods...)}
}

type TaskLabelFinder interface {
	FindTaskLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TaskLabel, error)
}

// FindTaskLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTaskLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*TaskLabel, error) {
	taskLabelObj := &TaskLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `task_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, taskLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from task_labels")
	}

	return taskLabelObj, nil
}

type TaskLabelInserter interface {
	Insert(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q taskLabelQuery) Insert(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no task_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(taskLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	taskLabelInsertCacheMut.RLock()
	cache, cached := taskLabelInsertCache[key]
	taskLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			taskLabelAllColumns,
			taskLabelColumnsWithDefault,
			taskLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(taskLabelType, taskLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(taskLabelType, taskLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `task_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `task_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `task_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, taskLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into task_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == taskLabelMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for task_labels")
	}

CacheNoHooks:
	if !cached {
		taskLabelInsertCacheMut.Lock()
		taskLabelInsertCache[key] = cache
		taskLabelInsertCacheMut.Unlock()
	}

	return nil
}

type TaskLabelUpdater interface {
	Update(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the TaskLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q taskLabelQuery) Update(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	taskLabelUpdateCacheMut.RLock()
	cache, cached := taskLabelUpdateCache[key]
	taskLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			taskLabelAllColumns,
			taskLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update task_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `task_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, taskLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(taskLabelType, taskLabelMapping, append(wl, taskLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update task_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for task_labels")
	}

	if !cached {
		taskLabelUpdateCacheMut.Lock()
		taskLabelUpdateCache[key] = cache
		taskLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q taskLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for task_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for task_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q taskLabelQuery) UpdateAllSlice(o TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `task_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in taskLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all taskLabel")
	}
	return rowsAff, nil
}

type TaskLabelDeleter interface {
	Delete(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single TaskLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q taskLabelQuery) Delete(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TaskLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), taskLabelPrimaryKeyMapping)
	sql := "DELETE FROM `task_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from task_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for task_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q taskLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no taskLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from task_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for task_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q taskLabelQuery) DeleteAllSlice(o TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `task_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from taskLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for task_labels")
	}

	return rowsAff, nil
}

type TaskLabelReloader interface {
	Reload(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q taskLabelQuery) Reload(o *TaskLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTaskLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q taskLabelQuery) ReloadAll(o *TaskLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TaskLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), taskLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `task_labels`.* FROM `task_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, taskLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TaskLabelSlice")
	}

	*o = slice

	return nil
}

// TaskLabelExists checks if the TaskLabel row exists.
func TaskLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `task_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if task_labels exists")
	}

	return exists, nil
}
