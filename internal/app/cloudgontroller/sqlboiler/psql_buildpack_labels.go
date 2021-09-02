// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/buildpack_labels.go
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

// BuildpackLabel is an object representing the database table.
type BuildpackLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *buildpackLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildpackLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildpackLabelColumns = struct {
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

var BuildpackLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "buildpack_labels.id",
	GUID:         "buildpack_labels.guid",
	CreatedAt:    "buildpack_labels.created_at",
	UpdatedAt:    "buildpack_labels.updated_at",
	ResourceGUID: "buildpack_labels.resource_guid",
	KeyPrefix:    "buildpack_labels.key_prefix",
	KeyName:      "buildpack_labels.key_name",
	Value:        "buildpack_labels.value",
}

// Generated where

var BuildpackLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"buildpack_labels\".\"id\""},
	GUID:         whereHelperstring{field: "\"buildpack_labels\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"buildpack_labels\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"buildpack_labels\".\"updated_at\""},
	ResourceGUID: whereHelpernull_String{field: "\"buildpack_labels\".\"resource_guid\""},
	KeyPrefix:    whereHelpernull_String{field: "\"buildpack_labels\".\"key_prefix\""},
	KeyName:      whereHelpernull_String{field: "\"buildpack_labels\".\"key_name\""},
	Value:        whereHelpernull_String{field: "\"buildpack_labels\".\"value\""},
}

// BuildpackLabelRels is where relationship names are stored.
var BuildpackLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// buildpackLabelR is where relationships are stored.
type buildpackLabelR struct {
	Resource *Buildpack `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*buildpackLabelR) NewStruct() *buildpackLabelR {
	return &buildpackLabelR{}
}

// buildpackLabelL is where Load methods for each relationship are stored.
type buildpackLabelL struct{}

var (
	buildpackLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	buildpackLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	buildpackLabelColumnsWithDefault    = []string{"id", "created_at"}
	buildpackLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// BuildpackLabelSlice is an alias for a slice of pointers to BuildpackLabel.
	// This should almost always be used instead of []BuildpackLabel.
	BuildpackLabelSlice []*BuildpackLabel

	BuildpackLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildpackLabelType                 = reflect.TypeOf(&BuildpackLabel{})
	buildpackLabelMapping              = queries.MakeStructMapping(buildpackLabelType)
	buildpackLabelPrimaryKeyMapping, _ = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, buildpackLabelPrimaryKeyColumns)
	buildpackLabelInsertCacheMut       sync.RWMutex
	buildpackLabelInsertCache          = make(map[string]insertCache)
	buildpackLabelUpdateCacheMut       sync.RWMutex
	buildpackLabelUpdateCache          = make(map[string]updateCache)
	buildpackLabelUpsertCacheMut       sync.RWMutex
	buildpackLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type BuildpackLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single buildpackLabel record from the query.
func (q BuildpackLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLabel, error) {
	o := &BuildpackLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buildpack_labels")
	}

	return o, nil
}

// All returns all BuildpackLabel records from the query.
func (q BuildpackLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLabelSlice, error) {
	var o []*BuildpackLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BuildpackLabel slice")
	}

	return o, nil
}

// Count returns the count of all BuildpackLabel records in the query.
func (q BuildpackLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buildpack_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q BuildpackLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buildpack_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q BuildpackLabelQuery) Resource(o *BuildpackLabel, mods ...qm.QueryMod) BuildpackQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Buildpacks(queryMods...)
	queries.SetFrom(query.Query, "\"buildpacks\"")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (buildpackLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildpackLabel interface{}, mods queries.Applicator) error {
	var slice []*BuildpackLabel
	var object *BuildpackLabel

	if singular {
		object = maybeBuildpackLabel.(*BuildpackLabel)
	} else {
		slice = *maybeBuildpackLabel.(*[]*BuildpackLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildpackLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildpackLabelR{}
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
		qm.From(`buildpacks`),
		qm.WhereIn(`buildpacks.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Buildpack")
	}

	var resultSlice []*Buildpack
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Buildpack")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for buildpacks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for buildpacks")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &buildpackR{}
		}
		foreign.R.ResourceBuildpackLabels = append(foreign.R.ResourceBuildpackLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &buildpackR{}
				}
				foreign.R.ResourceBuildpackLabels = append(foreign.R.ResourceBuildpackLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the buildpackLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceBuildpackLabels.
func (q BuildpackLabelQuery) SetResource(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Buildpack) error {
	var err error
	if insert {
		if err = Buildpacks().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"buildpack_labels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
		strmangle.WhereClause("\"", "\"", 2, buildpackLabelPrimaryKeyColumns),
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
		o.R = &buildpackLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &buildpackR{
			ResourceBuildpackLabels: BuildpackLabelSlice{o},
		}
	} else {
		related.R.ResourceBuildpackLabels = append(related.R.ResourceBuildpackLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q BuildpackLabelQuery) RemoveResource(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, related *Buildpack) error {
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

	for i, ri := range related.R.ResourceBuildpackLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceBuildpackLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceBuildpackLabels[i] = related.R.ResourceBuildpackLabels[ln-1]
		}
		related.R.ResourceBuildpackLabels = related.R.ResourceBuildpackLabels[:ln-1]
		break
	}
	return nil
}

// BuildpackLabels retrieves all the records using an executor.
func BuildpackLabels(mods ...qm.QueryMod) BuildpackLabelQuery {
	mods = append(mods, qm.From("\"buildpack_labels\""))
	return BuildpackLabelQuery{NewQuery(mods...)}
}

type BuildpackLabelFinder interface {
	FindBuildpackLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLabel, error)
}

// FindBuildpackLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuildpackLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLabel, error) {
	buildpackLabelObj := &BuildpackLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"buildpack_labels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildpackLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buildpack_labels")
	}

	return buildpackLabelObj, nil
}

type BuildpackLabelInserter interface {
	Insert(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q BuildpackLabelQuery) Insert(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(buildpackLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildpackLabelInsertCacheMut.RLock()
	cache, cached := buildpackLabelInsertCache[key]
	buildpackLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildpackLabelAllColumns,
			buildpackLabelColumnsWithDefault,
			buildpackLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"buildpack_labels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"buildpack_labels\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into buildpack_labels")
	}

	if !cached {
		buildpackLabelInsertCacheMut.Lock()
		buildpackLabelInsertCache[key] = cache
		buildpackLabelInsertCacheMut.Unlock()
	}

	return nil
}

type BuildpackLabelUpdater interface {
	Update(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the BuildpackLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q BuildpackLabelQuery) Update(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	buildpackLabelUpdateCacheMut.RLock()
	cache, cached := buildpackLabelUpdateCache[key]
	buildpackLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildpackLabelAllColumns,
			buildpackLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buildpack_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"buildpack_labels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, buildpackLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, append(wl, buildpackLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update buildpack_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buildpack_labels")
	}

	if !cached {
		buildpackLabelUpdateCacheMut.Lock()
		buildpackLabelUpdateCache[key] = cache
		buildpackLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q BuildpackLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buildpack_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buildpack_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q BuildpackLabelQuery) UpdateAllSlice(o BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"buildpack_labels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, buildpackLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in buildpackLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all buildpackLabel")
	}
	return rowsAff, nil
}

type BuildpackLabelDeleter interface {
	Delete(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single BuildpackLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q BuildpackLabelQuery) Delete(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BuildpackLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildpackLabelPrimaryKeyMapping)
	sql := "DELETE FROM \"buildpack_labels\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buildpack_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buildpack_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q BuildpackLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildpackLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpack_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q BuildpackLabelQuery) DeleteAllSlice(o BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"buildpack_labels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpackLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_labels")
	}

	return rowsAff, nil
}

type BuildpackLabelReloader interface {
	Reload(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q BuildpackLabelQuery) Reload(o *BuildpackLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuildpackLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q BuildpackLabelQuery) ReloadAll(o *BuildpackLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildpackLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"buildpack_labels\".* FROM \"buildpack_labels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildpackLabelSlice")
	}

	*o = slice

	return nil
}

// BuildpackLabelExists checks if the BuildpackLabel row exists.
func BuildpackLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"buildpack_labels\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buildpack_labels exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BuildpackLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(buildpackLabelColumnsWithDefault, o)

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

	buildpackLabelUpsertCacheMut.RLock()
	cache, cached := buildpackLabelUpsertCache[key]
	buildpackLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			buildpackLabelAllColumns,
			buildpackLabelColumnsWithDefault,
			buildpackLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			buildpackLabelAllColumns,
			buildpackLabelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert buildpack_labels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(buildpackLabelPrimaryKeyColumns))
			copy(conflict, buildpackLabelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"buildpack_labels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildpackLabelType, buildpackLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert buildpack_labels")
	}

	if !cached {
		buildpackLabelUpsertCacheMut.Lock()
		buildpackLabelUpsertCache[key] = cache
		buildpackLabelUpsertCacheMut.Unlock()
	}

	return nil
}
