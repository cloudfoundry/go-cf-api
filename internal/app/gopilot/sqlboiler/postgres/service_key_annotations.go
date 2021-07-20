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

// ServiceKeyAnnotation is an object representing the database table.
type ServiceKeyAnnotation struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	Key          null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *serviceKeyAnnotationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceKeyAnnotationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceKeyAnnotationColumns = struct {
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

var ServiceKeyAnnotationTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	Key          string
	Value        string
}{
	ID:           "service_key_annotations.id",
	GUID:         "service_key_annotations.guid",
	CreatedAt:    "service_key_annotations.created_at",
	UpdatedAt:    "service_key_annotations.updated_at",
	ResourceGUID: "service_key_annotations.resource_guid",
	KeyPrefix:    "service_key_annotations.key_prefix",
	Key:          "service_key_annotations.key",
	Value:        "service_key_annotations.value",
}

// Generated where

var ServiceKeyAnnotationWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	Key          whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "\"service_key_annotations\".\"id\""},
	GUID:         whereHelperstring{field: "\"service_key_annotations\".\"guid\""},
	CreatedAt:    whereHelpertime_Time{field: "\"service_key_annotations\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"service_key_annotations\".\"updated_at\""},
	ResourceGUID: whereHelpernull_String{field: "\"service_key_annotations\".\"resource_guid\""},
	KeyPrefix:    whereHelpernull_String{field: "\"service_key_annotations\".\"key_prefix\""},
	Key:          whereHelpernull_String{field: "\"service_key_annotations\".\"key\""},
	Value:        whereHelpernull_String{field: "\"service_key_annotations\".\"value\""},
}

// ServiceKeyAnnotationRels is where relationship names are stored.
var ServiceKeyAnnotationRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// serviceKeyAnnotationR is where relationships are stored.
type serviceKeyAnnotationR struct {
	Resource *ServiceKey `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*serviceKeyAnnotationR) NewStruct() *serviceKeyAnnotationR {
	return &serviceKeyAnnotationR{}
}

// serviceKeyAnnotationL is where Load methods for each relationship are stored.
type serviceKeyAnnotationL struct{}

var (
	serviceKeyAnnotationAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	serviceKeyAnnotationColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key", "value"}
	serviceKeyAnnotationColumnsWithDefault    = []string{"id", "created_at"}
	serviceKeyAnnotationPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceKeyAnnotationSlice is an alias for a slice of pointers to ServiceKeyAnnotation.
	// This should almost always be used instead of []ServiceKeyAnnotation.
	ServiceKeyAnnotationSlice []*ServiceKeyAnnotation
	// ServiceKeyAnnotationHook is the signature for custom ServiceKeyAnnotation hook methods
	ServiceKeyAnnotationHook func(context.Context, boil.ContextExecutor, *ServiceKeyAnnotation) error

	serviceKeyAnnotationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceKeyAnnotationType                 = reflect.TypeOf(&ServiceKeyAnnotation{})
	serviceKeyAnnotationMapping              = queries.MakeStructMapping(serviceKeyAnnotationType)
	serviceKeyAnnotationPrimaryKeyMapping, _ = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, serviceKeyAnnotationPrimaryKeyColumns)
	serviceKeyAnnotationInsertCacheMut       sync.RWMutex
	serviceKeyAnnotationInsertCache          = make(map[string]insertCache)
	serviceKeyAnnotationUpdateCacheMut       sync.RWMutex
	serviceKeyAnnotationUpdateCache          = make(map[string]updateCache)
	serviceKeyAnnotationUpsertCacheMut       sync.RWMutex
	serviceKeyAnnotationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var serviceKeyAnnotationBeforeInsertHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationBeforeUpdateHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationBeforeDeleteHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationBeforeUpsertHooks []ServiceKeyAnnotationHook

var serviceKeyAnnotationAfterInsertHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationAfterSelectHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationAfterUpdateHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationAfterDeleteHooks []ServiceKeyAnnotationHook
var serviceKeyAnnotationAfterUpsertHooks []ServiceKeyAnnotationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ServiceKeyAnnotation) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ServiceKeyAnnotation) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ServiceKeyAnnotation) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ServiceKeyAnnotation) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ServiceKeyAnnotation) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ServiceKeyAnnotation) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ServiceKeyAnnotation) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ServiceKeyAnnotation) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ServiceKeyAnnotation) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceKeyAnnotationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddServiceKeyAnnotationHook registers your hook function for all future operations.
func AddServiceKeyAnnotationHook(hookPoint boil.HookPoint, serviceKeyAnnotationHook ServiceKeyAnnotationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		serviceKeyAnnotationBeforeInsertHooks = append(serviceKeyAnnotationBeforeInsertHooks, serviceKeyAnnotationHook)
	case boil.BeforeUpdateHook:
		serviceKeyAnnotationBeforeUpdateHooks = append(serviceKeyAnnotationBeforeUpdateHooks, serviceKeyAnnotationHook)
	case boil.BeforeDeleteHook:
		serviceKeyAnnotationBeforeDeleteHooks = append(serviceKeyAnnotationBeforeDeleteHooks, serviceKeyAnnotationHook)
	case boil.BeforeUpsertHook:
		serviceKeyAnnotationBeforeUpsertHooks = append(serviceKeyAnnotationBeforeUpsertHooks, serviceKeyAnnotationHook)
	case boil.AfterInsertHook:
		serviceKeyAnnotationAfterInsertHooks = append(serviceKeyAnnotationAfterInsertHooks, serviceKeyAnnotationHook)
	case boil.AfterSelectHook:
		serviceKeyAnnotationAfterSelectHooks = append(serviceKeyAnnotationAfterSelectHooks, serviceKeyAnnotationHook)
	case boil.AfterUpdateHook:
		serviceKeyAnnotationAfterUpdateHooks = append(serviceKeyAnnotationAfterUpdateHooks, serviceKeyAnnotationHook)
	case boil.AfterDeleteHook:
		serviceKeyAnnotationAfterDeleteHooks = append(serviceKeyAnnotationAfterDeleteHooks, serviceKeyAnnotationHook)
	case boil.AfterUpsertHook:
		serviceKeyAnnotationAfterUpsertHooks = append(serviceKeyAnnotationAfterUpsertHooks, serviceKeyAnnotationHook)
	}
}

// One returns a single serviceKeyAnnotation record from the query.
func (q serviceKeyAnnotationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceKeyAnnotation, error) {
	o := &ServiceKeyAnnotation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_key_annotations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ServiceKeyAnnotation records from the query.
func (q serviceKeyAnnotationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceKeyAnnotationSlice, error) {
	var o []*ServiceKeyAnnotation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceKeyAnnotation slice")
	}

	if len(serviceKeyAnnotationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ServiceKeyAnnotation records in the query.
func (q serviceKeyAnnotationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_key_annotations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceKeyAnnotationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_key_annotations exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (o *ServiceKeyAnnotation) Resource(mods ...qm.QueryMod) serviceKeyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guid\" = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceKeys(queryMods...)
	queries.SetFrom(query.Query, "\"service_keys\"")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceKeyAnnotationL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceKeyAnnotation interface{}, mods queries.Applicator) error {
	var slice []*ServiceKeyAnnotation
	var object *ServiceKeyAnnotation

	if singular {
		object = maybeServiceKeyAnnotation.(*ServiceKeyAnnotation)
	} else {
		slice = *maybeServiceKeyAnnotation.(*[]*ServiceKeyAnnotation)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceKeyAnnotationR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceKeyAnnotationR{}
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
		qm.From(`service_keys`),
		qm.WhereIn(`service_keys.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServiceKey")
	}

	var resultSlice []*ServiceKey
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServiceKey")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for service_keys")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for service_keys")
	}

	if len(serviceKeyAnnotationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &serviceKeyR{}
		}
		foreign.R.ResourceServiceKeyAnnotations = append(foreign.R.ResourceServiceKeyAnnotations, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &serviceKeyR{}
				}
				foreign.R.ResourceServiceKeyAnnotations = append(foreign.R.ResourceServiceKeyAnnotations, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the serviceKeyAnnotation to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceServiceKeyAnnotations.
func (o *ServiceKeyAnnotation) SetResource(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceKey) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"service_key_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
		strmangle.WhereClause("\"", "\"", 2, serviceKeyAnnotationPrimaryKeyColumns),
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
		o.R = &serviceKeyAnnotationR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &serviceKeyR{
			ResourceServiceKeyAnnotations: ServiceKeyAnnotationSlice{o},
		}
	} else {
		related.R.ResourceServiceKeyAnnotations = append(related.R.ResourceServiceKeyAnnotations, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *ServiceKeyAnnotation) RemoveResource(ctx context.Context, exec boil.ContextExecutor, related *ServiceKey) error {
	var err error

	queries.SetScanner(&o.ResourceGUID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("resource_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Resource = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ResourceServiceKeyAnnotations {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceServiceKeyAnnotations)
		if ln > 1 && i < ln-1 {
			related.R.ResourceServiceKeyAnnotations[i] = related.R.ResourceServiceKeyAnnotations[ln-1]
		}
		related.R.ResourceServiceKeyAnnotations = related.R.ResourceServiceKeyAnnotations[:ln-1]
		break
	}
	return nil
}

// ServiceKeyAnnotations retrieves all the records using an executor.
func ServiceKeyAnnotations(mods ...qm.QueryMod) serviceKeyAnnotationQuery {
	mods = append(mods, qm.From("\"service_key_annotations\""))
	return serviceKeyAnnotationQuery{NewQuery(mods...)}
}

// FindServiceKeyAnnotation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceKeyAnnotation(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceKeyAnnotation, error) {
	serviceKeyAnnotationObj := &ServiceKeyAnnotation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"service_key_annotations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceKeyAnnotationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_key_annotations")
	}

	if err = serviceKeyAnnotationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return serviceKeyAnnotationObj, err
	}

	return serviceKeyAnnotationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServiceKeyAnnotation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_key_annotations provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceKeyAnnotationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceKeyAnnotationInsertCacheMut.RLock()
	cache, cached := serviceKeyAnnotationInsertCache[key]
	serviceKeyAnnotationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceKeyAnnotationAllColumns,
			serviceKeyAnnotationColumnsWithDefault,
			serviceKeyAnnotationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"service_key_annotations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"service_key_annotations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into service_key_annotations")
	}

	if !cached {
		serviceKeyAnnotationInsertCacheMut.Lock()
		serviceKeyAnnotationInsertCache[key] = cache
		serviceKeyAnnotationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ServiceKeyAnnotation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServiceKeyAnnotation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	serviceKeyAnnotationUpdateCacheMut.RLock()
	cache, cached := serviceKeyAnnotationUpdateCache[key]
	serviceKeyAnnotationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceKeyAnnotationAllColumns,
			serviceKeyAnnotationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_key_annotations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"service_key_annotations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, serviceKeyAnnotationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, append(wl, serviceKeyAnnotationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_key_annotations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_key_annotations")
	}

	if !cached {
		serviceKeyAnnotationUpdateCacheMut.Lock()
		serviceKeyAnnotationUpdateCache[key] = cache
		serviceKeyAnnotationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q serviceKeyAnnotationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_key_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_key_annotations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServiceKeyAnnotationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"service_key_annotations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, serviceKeyAnnotationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceKeyAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceKeyAnnotation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceKeyAnnotation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_key_annotations provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceKeyAnnotationColumnsWithDefault, o)

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

	serviceKeyAnnotationUpsertCacheMut.RLock()
	cache, cached := serviceKeyAnnotationUpsertCache[key]
	serviceKeyAnnotationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceKeyAnnotationAllColumns,
			serviceKeyAnnotationColumnsWithDefault,
			serviceKeyAnnotationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceKeyAnnotationAllColumns,
			serviceKeyAnnotationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert service_key_annotations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(serviceKeyAnnotationPrimaryKeyColumns))
			copy(conflict, serviceKeyAnnotationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"service_key_annotations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceKeyAnnotationType, serviceKeyAnnotationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert service_key_annotations")
	}

	if !cached {
		serviceKeyAnnotationUpsertCacheMut.Lock()
		serviceKeyAnnotationUpsertCache[key] = cache
		serviceKeyAnnotationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ServiceKeyAnnotation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServiceKeyAnnotation) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceKeyAnnotation provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceKeyAnnotationPrimaryKeyMapping)
	sql := "DELETE FROM \"service_key_annotations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_key_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_key_annotations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceKeyAnnotationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceKeyAnnotationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_key_annotations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_key_annotations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServiceKeyAnnotationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(serviceKeyAnnotationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"service_key_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceKeyAnnotationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceKeyAnnotation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_key_annotations")
	}

	if len(serviceKeyAnnotationAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ServiceKeyAnnotation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceKeyAnnotation(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServiceKeyAnnotationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceKeyAnnotationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceKeyAnnotationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"service_key_annotations\".* FROM \"service_key_annotations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, serviceKeyAnnotationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceKeyAnnotationSlice")
	}

	*o = slice

	return nil
}

// ServiceKeyAnnotationExists checks if the ServiceKeyAnnotation row exists.
func ServiceKeyAnnotationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"service_key_annotations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_key_annotations exists")
	}

	return exists, nil
}
