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

// StackAnnotation is an object representing the database table.
type StackAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *stackAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stackAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StackAnnotationColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ResourceGUID: "resource_guid",
	KeyPrefix:    "key_prefix",
	Key:          "key",
	Value:        "value",
}

var StackAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "stack_annotations.id",
	GUID:         "stack_annotations.guid",
	CreatedAt:    "stack_annotations.created_at",
	UpdatedAt:    "stack_annotations.updated_at",
	ResourceGUID: "stack_annotations.resource_guid",
	KeyPrefix:    "stack_annotations.key_prefix",
	Key:          "stack_annotations.key",
	Value:        "stack_annotations.value",
}

// Generated where

var StackAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`stack_annotations`.`id`"},
	GUID:         whereHelperstring{field: "`stack_annotations`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`stack_annotations`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`stack_annotations`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`stack_annotations`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`stack_annotations`.`key_prefix`"},
	Key:          whereHelpernull_String{field: "`stack_annotations`.`key`"},
	Value:        whereHelpernull_String{field: "`stack_annotations`.`value`"},
}

// StackAnnotationRels is where relationship names are stored.
var StackAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// stackAnnotationR is where relationships are stored.
type stackAnnotationR struct {
	Resource *Stack `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*stackAnnotationR) NewStruct() *stackAnnotationR {
	return &stackAnnotationR{}
}

// stackAnnotationL is where Load methods for each relationship are stored.
type stackAnnotationL struct{}

var (
	stackAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	stackAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	stackAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	stackAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// StackAnnotationSlice is an alias for a slice of pointers to StackAnnotation.
	// This should almost always be used instead of []StackAnnotation.
	StackAnnotationSlice []*StackAnnotation

	stackAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stackAnnotationType                 = reflect.TypeOf(&StackAnnotation{})
	stackAnnotationMapping              = queries.MakeStructMapping(stackAnnotationType)
	stackAnnotationPrimaryKeyMapping, _ = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, stackAnnotationPrimaryKeyColumns)
	stackAnnotationInsertCacheMut       sync.RWMutex
	stackAnnotationInsertCache          = make(map[string]insertCache)
	stackAnnotationUpdateCacheMut       sync.RWMutex
	stackAnnotationUpdateCache          = make(map[string]updateCache)
	stackAnnotationUpsertCacheMut       sync.RWMutex
	stackAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type StackAnnotationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*StackAnnotation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (StackAnnotationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single stackAnnotation record from the query.
func (q stackAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StackAnnotation, error) {
	o := &StackAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stack_annotations")
	}

	return o, nil
}

// All returns all StackAnnotation records from the query.
func (q stackAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (StackAnnotationSlice, error) {
	var o []*StackAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StackAnnotation slice")
	}

	return o, nil
}

// Count returns the count of all StackAnnotation records in the query.
func (q stackAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stack_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stackAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stack_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q stackAnnotationQuery) Resource(o *StackAnnotation, mods ...qm.QueryMod) stackQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Stacks(queryMods...)
	queries.SetFrom(query.Query, "`stacks`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stackAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStackAnnotation interface{}, mods queries.Applicator) error {
	var slice []*StackAnnotation
	var object *StackAnnotation

	if singular {
		object = maybeStackAnnotation.(*StackAnnotation)
	} else {
		slice = *maybeStackAnnotation.(*[]*StackAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stackAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stackAnnotationR{}
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
		qm.From(`stacks`),
		qm.WhereIn(`stacks.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Stack")
	}

	var resultSlice []*Stack
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Stack")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for stacks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stacks")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &stackR{}
		}
		foreign.R.ResourceStackAnnotations = append(foreign.R.ResourceStackAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &stackR{}
				}
				foreign.R.ResourceStackAnnotations = append(foreign.R.ResourceStackAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the stackAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceStackAnnotations.
func (q stackAnnotationQuery) SetResource(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Stack) error {
	var err error
	if insert {
		if err = Stacks().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `stack_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, stackAnnotationPrimaryKeyColumns),
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
		o.R = &stackAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &stackR{
			ResourceStackAnnotations: StackAnnotationSlice{o},
		}
	} else {
		related.R.ResourceStackAnnotations = append(related.R.ResourceStackAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q stackAnnotationQuery) RemoveResource(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, related *Stack) error {
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

	for i, ri := range related.R.ResourceStackAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceStackAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceStackAnnotations[i] = related.R.ResourceStackAnnotations[ln-1]
		}
		related.R.ResourceStackAnnotations = related.R.ResourceStackAnnotations[:ln-1]
		break
	}
	return nil
}

// StackAnnotations retrieves all the records using an executor.
func StackAnnotations(mods ...qm.QueryMod) stackAnnotationQuery {
	mods = append(mods, qm.From("`stack_annotations`"))
	return stackAnnotationQuery{NewQuery(mods...)}
}

type StackAnnotationFinder interface {
	FindStackAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*StackAnnotation, error)
}

// FindStackAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStackAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*StackAnnotation, error) {
	stackAnnotationObj := &StackAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `stack_annotations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stackAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stack_annotations")
	}

	return stackAnnotationObj, nil
}

type StackAnnotationInserter interface {
	Insert(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q stackAnnotationQuery) Insert(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stack_annotations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(stackAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stackAnnotationInsertCacheMut.RLock()
	cache, cached := stackAnnotationInsertCache[key]
	stackAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stackAnnotationAllColumns,
			stackAnnotationColumnsWithDefault,
			stackAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `stack_annotations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `stack_annotations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `stack_annotations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, stackAnnotationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into stack_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stackAnnotationMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for stack_annotations")
	}

CacheNoHooks:
	if !cached {
		stackAnnotationInsertCacheMut.Lock()
		stackAnnotationInsertCache[key] = cache
		stackAnnotationInsertCacheMut.Unlock()
	}

	return nil
}

type StackAnnotationUpdater interface {
	Update(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the StackAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q stackAnnotationQuery) Update(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	stackAnnotationUpdateCacheMut.RLock()
	cache, cached := stackAnnotationUpdateCache[key]
	stackAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stackAnnotationAllColumns,
			stackAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stack_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `stack_annotations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, stackAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, append(wl, stackAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update stack_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stack_annotations")
	}

	if !cached {
		stackAnnotationUpdateCacheMut.Lock()
		stackAnnotationUpdateCache[key] = cache
		stackAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q stackAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stack_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stack_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q stackAnnotationQuery) UpdateAllSlice(o StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `stack_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stackAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stackAnnotation")
	}
	return rowsAff, nil
}

type StackAnnotationDeleter interface {
	Delete(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single StackAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q stackAnnotationQuery) Delete(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StackAnnotation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stackAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM `stack_annotations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stack_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stack_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stackAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stackAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stack_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stack_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q stackAnnotationQuery) DeleteAllSlice(o StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `stack_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stackAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stack_annotations")
	}

	return rowsAff, nil
}

type StackAnnotationReloader interface {
	Reload(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q stackAnnotationQuery) Reload(o *StackAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStackAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q stackAnnotationQuery) ReloadAll(o *StackAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StackAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stackAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `stack_annotations`.* FROM `stack_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stackAnnotationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StackAnnotationSlice")
	}

	*o = slice

	return nil
}

// StackAnnotationExists checks if the StackAnnotation row exists.
func StackAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `stack_annotations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stack_annotations exists")
	}

	return exists, nil
}

var mySQLStackAnnotationUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *StackAnnotation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stack_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(stackAnnotationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLStackAnnotationUniqueColumns, o)

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

	stackAnnotationUpsertCacheMut.RLock()
	cache, cached := stackAnnotationUpsertCache[key]
	stackAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stackAnnotationAllColumns,
			stackAnnotationColumnsWithDefault,
			stackAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stackAnnotationAllColumns,
			stackAnnotationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert stack_annotations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`stack_annotations`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `stack_annotations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for stack_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stackAnnotationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(stackAnnotationType, stackAnnotationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for stack_annotations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for stack_annotations")
	}

CacheNoHooks:
	if !cached {
		stackAnnotationUpsertCacheMut.Lock()
		stackAnnotationUpsertCache[key] = cache
		stackAnnotationUpsertCacheMut.Unlock()
	}

	return nil
}
