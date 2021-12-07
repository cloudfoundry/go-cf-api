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

type OrganizationAnnotationUpserter interface {
	Upsert(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLOrganizationAnnotationUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q organizationAnnotationQuery) Upsert(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organization_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationAnnotationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrganizationAnnotationUniqueColumns, o)

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

	organizationAnnotationUpsertCacheMut.RLock()
	cache, cached := organizationAnnotationUpsertCache[key]
	organizationAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationAnnotationAllColumns,
			organizationAnnotationColumnsWithDefault,
			organizationAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationAnnotationAllColumns,
			organizationAnnotationPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert organization_annotations, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`organization_annotations`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `organization_annotations` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organization_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationAnnotationMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for organization_annotations")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organization_annotations")
	}

CacheNoHooks:
	if !cached {
		organizationAnnotationUpsertCacheMut.Lock()
		organizationAnnotationUpsertCache[key] = cache
		organizationAnnotationUpsertCacheMut.Unlock()
	}

	return nil
}

// OrganizationAnnotation is an object representing the database table.
type OrganizationAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *organizationAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationAnnotationColumns = struct {
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

var OrganizationAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "organization_annotations.id",
	GUID:         "organization_annotations.guid",
	CreatedAt:    "organization_annotations.created_at",
	UpdatedAt:    "organization_annotations.updated_at",
	ResourceGUID: "organization_annotations.resource_guid",
	KeyPrefix:    "organization_annotations.key_prefix",
	Key:          "organization_annotations.key",
	Value:        "organization_annotations.value",
}

// Generated where

var OrganizationAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`organization_annotations`.`id`"},
	GUID:         whereHelperstring{field: "`organization_annotations`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`organization_annotations`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`organization_annotations`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`organization_annotations`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`organization_annotations`.`key_prefix`"},
	Key:          whereHelpernull_String{field: "`organization_annotations`.`key`"},
	Value:        whereHelpernull_String{field: "`organization_annotations`.`value`"},
}

// OrganizationAnnotationRels is where relationship names are stored.
var OrganizationAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// organizationAnnotationR is where relationships are stored.
type organizationAnnotationR struct {
	Resource *Organization `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*organizationAnnotationR) NewStruct() *organizationAnnotationR {
	return &organizationAnnotationR{}
}

// organizationAnnotationL is where Load methods for each relationship are stored.
type organizationAnnotationL struct{}

var (
	organizationAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	organizationAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	organizationAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	organizationAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrganizationAnnotationSlice is an alias for a slice of pointers to OrganizationAnnotation.
	// This should almost always be used instead of []OrganizationAnnotation.
	OrganizationAnnotationSlice []*OrganizationAnnotation

	organizationAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationAnnotationType                 = reflect.TypeOf(&OrganizationAnnotation{})
	organizationAnnotationMapping              = queries.MakeStructMapping(organizationAnnotationType)
	organizationAnnotationPrimaryKeyMapping, _ = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, organizationAnnotationPrimaryKeyColumns)
	organizationAnnotationInsertCacheMut       sync.RWMutex
	organizationAnnotationInsertCache          = make(map[string]insertCache)
	organizationAnnotationUpdateCacheMut       sync.RWMutex
	organizationAnnotationUpdateCache          = make(map[string]updateCache)
	organizationAnnotationUpsertCacheMut       sync.RWMutex
	organizationAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrganizationAnnotationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationAnnotation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrganizationAnnotationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single organizationAnnotation record from the query.
func (q organizationAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationAnnotation, error) {
	o := &OrganizationAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organization_annotations")
	}

	return o, nil
}

// All returns all OrganizationAnnotation records from the query.
func (q organizationAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationAnnotationSlice, error) {
	var o []*OrganizationAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationAnnotation slice")
	}

	return o, nil
}

// Count returns the count of all OrganizationAnnotation records in the query.
func (q organizationAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organization_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organization_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q organizationAnnotationQuery) Resource(o *OrganizationAnnotation, mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "`organizations`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationAnnotation interface{}, mods queries.Applicator) error {
	var slice []*OrganizationAnnotation
	var object *OrganizationAnnotation

	if singular {
		object = maybeOrganizationAnnotation.(*OrganizationAnnotation)
	} else {
		slice = *maybeOrganizationAnnotation.(*[]*OrganizationAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationAnnotationR{}
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
		qm.From(`organizations`),
		qm.WhereIn(`organizations.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Organization")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Organization")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.ResourceOrganizationAnnotations = append(foreign.R.ResourceOrganizationAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.ResourceOrganizationAnnotations = append(foreign.R.ResourceOrganizationAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the organizationAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceOrganizationAnnotations.
func (q organizationAnnotationQuery) SetResource(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = Organizations().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organization_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, organizationAnnotationPrimaryKeyColumns),
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
		o.R = &organizationAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &organizationR{
			ResourceOrganizationAnnotations: OrganizationAnnotationSlice{o},
		}
	} else {
		related.R.ResourceOrganizationAnnotations = append(related.R.ResourceOrganizationAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q organizationAnnotationQuery) RemoveResource(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, related *Organization) error {
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

	for i, ri := range related.R.ResourceOrganizationAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceOrganizationAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceOrganizationAnnotations[i] = related.R.ResourceOrganizationAnnotations[ln-1]
		}
		related.R.ResourceOrganizationAnnotations = related.R.ResourceOrganizationAnnotations[:ln-1]
		break
	}
	return nil
}

// OrganizationAnnotations retrieves all the records using an executor.
func OrganizationAnnotations(mods ...qm.QueryMod) organizationAnnotationQuery {
	mods = append(mods, qm.From("`organization_annotations`"))
	return organizationAnnotationQuery{NewQuery(mods...)}
}

type OrganizationAnnotationFinder interface {
	FindOrganizationAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrganizationAnnotation, error)
}

// FindOrganizationAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*OrganizationAnnotation, error) {
	organizationAnnotationObj := &OrganizationAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `organization_annotations` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, organizationAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organization_annotations")
	}

	return organizationAnnotationObj, nil
}

type OrganizationAnnotationInserter interface {
	Insert(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q organizationAnnotationQuery) Insert(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organization_annotations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(organizationAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationAnnotationInsertCacheMut.RLock()
	cache, cached := organizationAnnotationInsertCache[key]
	organizationAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationAnnotationAllColumns,
			organizationAnnotationColumnsWithDefault,
			organizationAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `organization_annotations` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `organization_annotations` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `organization_annotations` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, organizationAnnotationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into organization_annotations")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationAnnotationMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for organization_annotations")
	}

CacheNoHooks:
	if !cached {
		organizationAnnotationInsertCacheMut.Lock()
		organizationAnnotationInsertCache[key] = cache
		organizationAnnotationInsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationAnnotationUpdater interface {
	Update(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrganizationAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q organizationAnnotationQuery) Update(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	organizationAnnotationUpdateCacheMut.RLock()
	cache, cached := organizationAnnotationUpdateCache[key]
	organizationAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationAnnotationAllColumns,
			organizationAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organization_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `organization_annotations` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, organizationAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationAnnotationType, organizationAnnotationMapping, append(wl, organizationAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organization_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organization_annotations")
	}

	if !cached {
		organizationAnnotationUpdateCacheMut.Lock()
		organizationAnnotationUpdateCache[key] = cache
		organizationAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q organizationAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organization_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organization_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q organizationAnnotationQuery) UpdateAllSlice(o OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `organization_annotations` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationAnnotation")
	}
	return rowsAff, nil
}

type OrganizationAnnotationDeleter interface {
	Delete(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrganizationAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q organizationAnnotationQuery) Delete(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationAnnotation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM `organization_annotations` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organization_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organization_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organization_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organization_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q organizationAnnotationQuery) DeleteAllSlice(o OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `organization_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organization_annotations")
	}

	return rowsAff, nil
}

type OrganizationAnnotationReloader interface {
	Reload(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q organizationAnnotationQuery) Reload(o *OrganizationAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q organizationAnnotationQuery) ReloadAll(o *OrganizationAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `organization_annotations`.* FROM `organization_annotations` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationAnnotationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationAnnotationSlice")
	}

	*o = slice

	return nil
}

// OrganizationAnnotationExists checks if the OrganizationAnnotation row exists.
func OrganizationAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `organization_annotations` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organization_annotations exists")
	}

	return exists, nil
}
