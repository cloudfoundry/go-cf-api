//go:generate mockgen -source=$GOFILE -destination=mocks/droplet_annotations.go -copyright_file=../../../../buildtags.txt
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

// DropletAnnotation is an object representing the database table.
type DropletAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *dropletAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dropletAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DropletAnnotationColumns = struct {
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

var DropletAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "droplet_annotations.id",
	GUID:         "droplet_annotations.guid",
	CreatedAt:    "droplet_annotations.created_at",
	UpdatedAt:    "droplet_annotations.updated_at",
	ResourceGUID: "droplet_annotations.resource_guid",
	KeyPrefix:    "droplet_annotations.key_prefix",
	Key:          "droplet_annotations.key",
	Value:        "droplet_annotations.value",
}

// Generated where

var DropletAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"droplet_annotations\".\"id\""},
	GUID:         whereHelperstring{field: "\"droplet_annotations\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"droplet_annotations\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"droplet_annotations\".\"updated_at\""},
	ResourceGUID: whereHelpernull_String{field: "\"droplet_annotations\".\"resource_guid\""},
	KeyPrefix:    whereHelpernull_String{field: "\"droplet_annotations\".\"key_prefix\""},
	Key:          whereHelpernull_String{field: "\"droplet_annotations\".\"key\""},
	Value:        whereHelpernull_String{field: "\"droplet_annotations\".\"value\""},
}

// DropletAnnotationRels is where relationship names are stored.
var DropletAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// dropletAnnotationR is where relationships are stored.
type dropletAnnotationR struct {
	Resource *Droplet `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*dropletAnnotationR) NewStruct() *dropletAnnotationR {
	return &dropletAnnotationR{}
}

// dropletAnnotationL is where Load methods for each relationship are stored.
type dropletAnnotationL struct{}

var (
	dropletAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	dropletAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	dropletAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	dropletAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// DropletAnnotationSlice is an alias for a slice of pointers to DropletAnnotation.
	// This should almost always be used instead of []DropletAnnotation.
	DropletAnnotationSlice []*DropletAnnotation

	dropletAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dropletAnnotationType                 = reflect.TypeOf(&DropletAnnotation{})
	dropletAnnotationMapping              = queries.MakeStructMapping(dropletAnnotationType)
	dropletAnnotationPrimaryKeyMapping, _ = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, dropletAnnotationPrimaryKeyColumns)
	dropletAnnotationInsertCacheMut       sync.RWMutex
	dropletAnnotationInsertCache          = make(map[string]insertCache)
	dropletAnnotationUpdateCacheMut       sync.RWMutex
	dropletAnnotationUpdateCache          = make(map[string]updateCache)
	dropletAnnotationUpsertCacheMut       sync.RWMutex
	dropletAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type DropletAnnotationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*DropletAnnotation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (DropletAnnotationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single dropletAnnotation record from the query.
func (q dropletAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DropletAnnotation, error) {
	o := &DropletAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for droplet_annotations")
	}

	return o, nil
}

// All returns all DropletAnnotation records from the query.
func (q dropletAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (DropletAnnotationSlice, error) {
	var o []*DropletAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DropletAnnotation slice")
	}

	return o, nil
}

// Count returns the count of all DropletAnnotation records in the query.
func (q dropletAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count droplet_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dropletAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if droplet_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q dropletAnnotationQuery) Resource(o *DropletAnnotation, mods ...qm.QueryMod) dropletQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Droplets(queryMods...)
	queries.SetFrom(query.Query, "\"droplets\"")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dropletAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDropletAnnotation interface{}, mods queries.Applicator) error {
	var slice []*DropletAnnotation
	var object *DropletAnnotation

	if singular {
		object = maybeDropletAnnotation.(*DropletAnnotation)
	} else {
		slice = *maybeDropletAnnotation.(*[]*DropletAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dropletAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dropletAnnotationR{}
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
		qm.From(`droplets`),
		qm.WhereIn(`droplets.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Droplet")
	}

	var resultSlice []*Droplet
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Droplet")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for droplets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for droplets")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &dropletR{}
		}
		foreign.R.ResourceDropletAnnotations = append(foreign.R.ResourceDropletAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &dropletR{}
				}
				foreign.R.ResourceDropletAnnotations = append(foreign.R.ResourceDropletAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the dropletAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceDropletAnnotations.
func (q dropletAnnotationQuery) SetResource(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Droplet) error {
	var err error
	if insert {
		if err = Droplets().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"droplet_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
		strmangle.WhereClause("\"", "\"", 2, dropletAnnotationPrimaryKeyColumns),
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
		o.R = &dropletAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &dropletR{
			ResourceDropletAnnotations: DropletAnnotationSlice{o},
		}
	} else {
		related.R.ResourceDropletAnnotations = append(related.R.ResourceDropletAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q dropletAnnotationQuery) RemoveResource(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, related *Droplet) error {
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

	for i, ri := range related.R.ResourceDropletAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceDropletAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceDropletAnnotations[i] = related.R.ResourceDropletAnnotations[ln-1]
		}
		related.R.ResourceDropletAnnotations = related.R.ResourceDropletAnnotations[:ln-1]
		break
	}
	return nil
}

// DropletAnnotations retrieves all the records using an executor.
func DropletAnnotations(mods ...qm.QueryMod) dropletAnnotationQuery {
	mods = append(mods, qm.From("\"droplet_annotations\""))
	return dropletAnnotationQuery{NewQuery(mods...)}
}

type DropletAnnotationFinder interface {
	FindDropletAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DropletAnnotation, error)
}

// FindDropletAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDropletAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DropletAnnotation, error) {
	dropletAnnotationObj := &DropletAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"droplet_annotations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, dropletAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from droplet_annotations")
	}

	return dropletAnnotationObj, nil
}

type DropletAnnotationInserter interface {
	Insert(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q dropletAnnotationQuery) Insert(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no droplet_annotations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(dropletAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dropletAnnotationInsertCacheMut.RLock()
	cache, cached := dropletAnnotationInsertCache[key]
	dropletAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dropletAnnotationAllColumns,
			dropletAnnotationColumnsWithDefault,
			dropletAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"droplet_annotations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"droplet_annotations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into droplet_annotations")
	}

	if !cached {
		dropletAnnotationInsertCacheMut.Lock()
		dropletAnnotationInsertCache[key] = cache
		dropletAnnotationInsertCacheMut.Unlock()
	}

	return nil
}

type DropletAnnotationUpdater interface {
	Update(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the DropletAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q dropletAnnotationQuery) Update(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	dropletAnnotationUpdateCacheMut.RLock()
	cache, cached := dropletAnnotationUpdateCache[key]
	dropletAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dropletAnnotationAllColumns,
			dropletAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update droplet_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"droplet_annotations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, dropletAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, append(wl, dropletAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update droplet_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for droplet_annotations")
	}

	if !cached {
		dropletAnnotationUpdateCacheMut.Lock()
		dropletAnnotationUpdateCache[key] = cache
		dropletAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q dropletAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for droplet_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for droplet_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q dropletAnnotationQuery) UpdateAllSlice(o DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dropletAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"droplet_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, dropletAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dropletAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dropletAnnotation")
	}
	return rowsAff, nil
}

type DropletAnnotationUpserter interface {
	Upsert(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q dropletAnnotationQuery) Upsert(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no droplet_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(dropletAnnotationColumnsWithDefault, o)

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

	dropletAnnotationUpsertCacheMut.RLock()
	cache, cached := dropletAnnotationUpsertCache[key]
	dropletAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dropletAnnotationAllColumns,
			dropletAnnotationColumnsWithDefault,
			dropletAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			dropletAnnotationAllColumns,
			dropletAnnotationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert droplet_annotations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(dropletAnnotationPrimaryKeyColumns))
			copy(conflict, dropletAnnotationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"droplet_annotations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dropletAnnotationType, dropletAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert droplet_annotations")
	}

	if !cached {
		dropletAnnotationUpsertCacheMut.Lock()
		dropletAnnotationUpsertCache[key] = cache
		dropletAnnotationUpsertCacheMut.Unlock()
	}

	return nil
}

type DropletAnnotationDeleter interface {
	Delete(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single DropletAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q dropletAnnotationQuery) Delete(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DropletAnnotation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dropletAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM \"droplet_annotations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from droplet_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for droplet_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dropletAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dropletAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from droplet_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for droplet_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q dropletAnnotationQuery) DeleteAllSlice(o DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dropletAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"droplet_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dropletAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dropletAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for droplet_annotations")
	}

	return rowsAff, nil
}

type DropletAnnotationReloader interface {
	Reload(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q dropletAnnotationQuery) Reload(o *DropletAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDropletAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q dropletAnnotationQuery) ReloadAll(o *DropletAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DropletAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dropletAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"droplet_annotations\".* FROM \"droplet_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, dropletAnnotationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DropletAnnotationSlice")
	}

	*o = slice

	return nil
}

// DropletAnnotationExists checks if the DropletAnnotation row exists.
func DropletAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"droplet_annotations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if droplet_annotations exists")
	}

	return exists, nil
}
