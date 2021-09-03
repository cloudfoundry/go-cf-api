//go:generate mockgen -source=$GOFILE -destination=mocks/service_offering_annotations.go
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

// ServiceOfferingAnnotation is an object representing the database table.
type ServiceOfferingAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *serviceOfferingAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceOfferingAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceOfferingAnnotationColumns = struct {
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

var ServiceOfferingAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "service_offering_annotations.id",
	GUID:         "service_offering_annotations.guid",
	CreatedAt:    "service_offering_annotations.created_at",
	UpdatedAt:    "service_offering_annotations.updated_at",
	ResourceGUID: "service_offering_annotations.resource_guid",
	KeyPrefix:    "service_offering_annotations.key_prefix",
	Key:          "service_offering_annotations.key",
	Value:        "service_offering_annotations.value",
}

// Generated where

var ServiceOfferingAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"service_offering_annotations\".\"id\""},
	GUID:         whereHelperstring{field: "\"service_offering_annotations\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"service_offering_annotations\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"service_offering_annotations\".\"updated_at\""},
	ResourceGUID: whereHelpernull_String{field: "\"service_offering_annotations\".\"resource_guid\""},
	KeyPrefix:    whereHelpernull_String{field: "\"service_offering_annotations\".\"key_prefix\""},
	Key:          whereHelpernull_String{field: "\"service_offering_annotations\".\"key\""},
	Value:        whereHelpernull_String{field: "\"service_offering_annotations\".\"value\""},
}

// ServiceOfferingAnnotationRels is where relationship names are stored.
var ServiceOfferingAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// serviceOfferingAnnotationR is where relationships are stored.
type serviceOfferingAnnotationR struct {
	Resource *Service `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*serviceOfferingAnnotationR) NewStruct() *serviceOfferingAnnotationR {
	return &serviceOfferingAnnotationR{}
}

// serviceOfferingAnnotationL is where Load methods for each relationship are stored.
type serviceOfferingAnnotationL struct{}

var (
	serviceOfferingAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	serviceOfferingAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	serviceOfferingAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	serviceOfferingAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceOfferingAnnotationSlice is an alias for a slice of pointers to ServiceOfferingAnnotation.
	// This should almost always be used instead of []ServiceOfferingAnnotation.
	ServiceOfferingAnnotationSlice []*ServiceOfferingAnnotation

	serviceOfferingAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceOfferingAnnotationType                 = reflect.TypeOf(&ServiceOfferingAnnotation{})
	serviceOfferingAnnotationMapping              = queries.MakeStructMapping(serviceOfferingAnnotationType)
	serviceOfferingAnnotationPrimaryKeyMapping, _ = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, serviceOfferingAnnotationPrimaryKeyColumns)
	serviceOfferingAnnotationInsertCacheMut       sync.RWMutex
	serviceOfferingAnnotationInsertCache          = make(map[string]insertCache)
	serviceOfferingAnnotationUpdateCacheMut       sync.RWMutex
	serviceOfferingAnnotationUpdateCache          = make(map[string]updateCache)
	serviceOfferingAnnotationUpsertCacheMut       sync.RWMutex
	serviceOfferingAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceOfferingAnnotationFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceOfferingAnnotation, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceOfferingAnnotationSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceOfferingAnnotation record from the query.
func (q serviceOfferingAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceOfferingAnnotation, error) {
	o := &ServiceOfferingAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_offering_annotations")
	}

	return o, nil
}

// All returns all ServiceOfferingAnnotation records from the query.
func (q serviceOfferingAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceOfferingAnnotationSlice, error) {
	var o []*ServiceOfferingAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceOfferingAnnotation slice")
	}

	return o, nil
}

// Count returns the count of all ServiceOfferingAnnotation records in the query.
func (q serviceOfferingAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_offering_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceOfferingAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_offering_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q serviceOfferingAnnotationQuery) Resource(o *ServiceOfferingAnnotation, mods ...qm.QueryMod) serviceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Services(queryMods...)
	queries.SetFrom(query.Query, "\"services\"")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceOfferingAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceOfferingAnnotation interface{}, mods queries.Applicator) error {
	var slice []*ServiceOfferingAnnotation
	var object *ServiceOfferingAnnotation

	if singular {
		object = maybeServiceOfferingAnnotation.(*ServiceOfferingAnnotation)
	} else {
		slice = *maybeServiceOfferingAnnotation.(*[]*ServiceOfferingAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceOfferingAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceOfferingAnnotationR{}
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
		qm.From(`services`),
		qm.WhereIn(`services.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Service")
	}

	var resultSlice []*Service
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Service")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for services")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for services")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &serviceR{}
		}
		foreign.R.ResourceServiceOfferingAnnotations = append(foreign.R.ResourceServiceOfferingAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &serviceR{}
				}
				foreign.R.ResourceServiceOfferingAnnotations = append(foreign.R.ResourceServiceOfferingAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the serviceOfferingAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceServiceOfferingAnnotations.
func (q serviceOfferingAnnotationQuery) SetResource(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Service) error {
	var err error
	if insert {
		if err = Services().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"service_offering_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
		strmangle.WhereClause("\"", "\"", 2, serviceOfferingAnnotationPrimaryKeyColumns),
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
		o.R = &serviceOfferingAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &serviceR{
			ResourceServiceOfferingAnnotations: ServiceOfferingAnnotationSlice{o},
		}
	} else {
		related.R.ResourceServiceOfferingAnnotations = append(related.R.ResourceServiceOfferingAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q serviceOfferingAnnotationQuery) RemoveResource(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, related *Service) error {
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

	for i, ri := range related.R.ResourceServiceOfferingAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceServiceOfferingAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceServiceOfferingAnnotations[i] = related.R.ResourceServiceOfferingAnnotations[ln-1]
		}
		related.R.ResourceServiceOfferingAnnotations = related.R.ResourceServiceOfferingAnnotations[:ln-1]
		break
	}
	return nil
}

// ServiceOfferingAnnotations retrieves all the records using an executor.
func ServiceOfferingAnnotations(mods ...qm.QueryMod) serviceOfferingAnnotationQuery {
	mods = append(mods, qm.From("\"service_offering_annotations\""))
	return serviceOfferingAnnotationQuery{NewQuery(mods...)}
}

type ServiceOfferingAnnotationFinder interface {
	FindServiceOfferingAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceOfferingAnnotation, error)
}

// FindServiceOfferingAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceOfferingAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceOfferingAnnotation, error) {
	serviceOfferingAnnotationObj := &ServiceOfferingAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_offering_annotations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceOfferingAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_offering_annotations")
	}

	return serviceOfferingAnnotationObj, nil
}

type ServiceOfferingAnnotationInserter interface {
	Insert(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q serviceOfferingAnnotationQuery) Insert(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_offering_annotations provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceOfferingAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceOfferingAnnotationInsertCacheMut.RLock()
	cache, cached := serviceOfferingAnnotationInsertCache[key]
	serviceOfferingAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceOfferingAnnotationAllColumns,
			serviceOfferingAnnotationColumnsWithDefault,
			serviceOfferingAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_offering_annotations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_offering_annotations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_offering_annotations")
	}

	if !cached {
		serviceOfferingAnnotationInsertCacheMut.Lock()
		serviceOfferingAnnotationInsertCache[key] = cache
		serviceOfferingAnnotationInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceOfferingAnnotationUpdater interface {
	Update(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceOfferingAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q serviceOfferingAnnotationQuery) Update(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceOfferingAnnotationUpdateCacheMut.RLock()
	cache, cached := serviceOfferingAnnotationUpdateCache[key]
	serviceOfferingAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceOfferingAnnotationAllColumns,
			serviceOfferingAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_offering_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_offering_annotations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceOfferingAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, append(wl, serviceOfferingAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_offering_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_offering_annotations")
	}

	if !cached {
		serviceOfferingAnnotationUpdateCacheMut.Lock()
		serviceOfferingAnnotationUpdateCache[key] = cache
		serviceOfferingAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serviceOfferingAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_offering_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_offering_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q serviceOfferingAnnotationQuery) UpdateAllSlice(o ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_offering_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceOfferingAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceOfferingAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceOfferingAnnotation")
	}
	return rowsAff, nil
}

type ServiceOfferingAnnotationDeleter interface {
	Delete(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceOfferingAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q serviceOfferingAnnotationQuery) Delete(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceOfferingAnnotation provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceOfferingAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM \"service_offering_annotations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_offering_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_offering_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceOfferingAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceOfferingAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_offering_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_offering_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q serviceOfferingAnnotationQuery) DeleteAllSlice(o ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_offering_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceOfferingAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceOfferingAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_offering_annotations")
	}

	return rowsAff, nil
}

type ServiceOfferingAnnotationReloader interface {
	Reload(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q serviceOfferingAnnotationQuery) Reload(o *ServiceOfferingAnnotation, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceOfferingAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q serviceOfferingAnnotationQuery) ReloadAll(o *ServiceOfferingAnnotationSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceOfferingAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_offering_annotations\".* FROM \"service_offering_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceOfferingAnnotationPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceOfferingAnnotationSlice")
	}

	*o = slice

	return nil
}

// ServiceOfferingAnnotationExists checks if the ServiceOfferingAnnotation row exists.
func ServiceOfferingAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_offering_annotations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_offering_annotations exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceOfferingAnnotation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_offering_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceOfferingAnnotationColumnsWithDefault, o)

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

	serviceOfferingAnnotationUpsertCacheMut.RLock()
	cache, cached := serviceOfferingAnnotationUpsertCache[key]
	serviceOfferingAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceOfferingAnnotationAllColumns,
			serviceOfferingAnnotationColumnsWithDefault,
			serviceOfferingAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceOfferingAnnotationAllColumns,
			serviceOfferingAnnotationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_offering_annotations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceOfferingAnnotationPrimaryKeyColumns))
			copy(conflict, serviceOfferingAnnotationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_offering_annotations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceOfferingAnnotationType, serviceOfferingAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_offering_annotations")
	}

	if !cached {
		serviceOfferingAnnotationUpsertCacheMut.Lock()
		serviceOfferingAnnotationUpsertCache[key] = cache
		serviceOfferingAnnotationUpsertCacheMut.Unlock()
	}

	return nil
}
