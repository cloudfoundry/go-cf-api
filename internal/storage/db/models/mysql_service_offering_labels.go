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

type ServiceOfferingLabelUpserter interface {
	Upsert(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLServiceOfferingLabelUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q serviceOfferingLabelQuery) Upsert(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_offering_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(serviceOfferingLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLServiceOfferingLabelUniqueColumns, o)

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

	serviceOfferingLabelUpsertCacheMut.RLock()
	cache, cached := serviceOfferingLabelUpsertCache[key]
	serviceOfferingLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceOfferingLabelAllColumns,
			serviceOfferingLabelColumnsWithDefault,
			serviceOfferingLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceOfferingLabelAllColumns,
			serviceOfferingLabelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert service_offering_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`service_offering_labels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `service_offering_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for service_offering_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceOfferingLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for service_offering_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for service_offering_labels")
	}

CacheNoHooks:
	if !cached {
		serviceOfferingLabelUpsertCacheMut.Lock()
		serviceOfferingLabelUpsertCache[key] = cache
		serviceOfferingLabelUpsertCacheMut.Unlock()
	}

	return nil
}

// ServiceOfferingLabel is an object representing the database table.
type ServiceOfferingLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *serviceOfferingLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceOfferingLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceOfferingLabelColumns = struct {
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

var ServiceOfferingLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "service_offering_labels.id",
	GUID:         "service_offering_labels.guid",
	CreatedAt:    "service_offering_labels.created_at",
	UpdatedAt:    "service_offering_labels.updated_at",
	ResourceGUID: "service_offering_labels.resource_guid",
	KeyPrefix:    "service_offering_labels.key_prefix",
	KeyName:      "service_offering_labels.key_name",
	Value:        "service_offering_labels.value",
}

// Generated where

var ServiceOfferingLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`service_offering_labels`.`id`"},
	GUID:         whereHelperstring{field: "`service_offering_labels`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`service_offering_labels`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`service_offering_labels`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`service_offering_labels`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`service_offering_labels`.`key_prefix`"},
	KeyName:      whereHelpernull_String{field: "`service_offering_labels`.`key_name`"},
	Value:        whereHelpernull_String{field: "`service_offering_labels`.`value`"},
}

// ServiceOfferingLabelRels is where relationship names are stored.
var ServiceOfferingLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// serviceOfferingLabelR is where relationships are stored.
type serviceOfferingLabelR struct {
	Resource *Service `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*serviceOfferingLabelR) NewStruct() *serviceOfferingLabelR {
	return &serviceOfferingLabelR{}
}

// serviceOfferingLabelL is where Load methods for each relationship are stored.
type serviceOfferingLabelL struct{}

var (
	serviceOfferingLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceOfferingLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceOfferingLabelColumnsWithDefault    = []string{"id", "created_at"}
	serviceOfferingLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceOfferingLabelSlice is an alias for a slice of pointers to ServiceOfferingLabel.
	// This should almost always be used instead of []ServiceOfferingLabel.
	ServiceOfferingLabelSlice []*ServiceOfferingLabel

	serviceOfferingLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceOfferingLabelType                 = reflect.TypeOf(&ServiceOfferingLabel{})
	serviceOfferingLabelMapping              = queries.MakeStructMapping(serviceOfferingLabelType)
	serviceOfferingLabelPrimaryKeyMapping, _ = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, serviceOfferingLabelPrimaryKeyColumns)
	serviceOfferingLabelInsertCacheMut       sync.RWMutex
	serviceOfferingLabelInsertCache          = make(map[string]insertCache)
	serviceOfferingLabelUpdateCacheMut       sync.RWMutex
	serviceOfferingLabelUpdateCache          = make(map[string]updateCache)
	serviceOfferingLabelUpsertCacheMut       sync.RWMutex
	serviceOfferingLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type ServiceOfferingLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*ServiceOfferingLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (ServiceOfferingLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single serviceOfferingLabel record from the query.
func (q serviceOfferingLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceOfferingLabel, error) {
	o := &ServiceOfferingLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_offering_labels")
	}

	return o, nil
}

// All returns all ServiceOfferingLabel records from the query.
func (q serviceOfferingLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceOfferingLabelSlice, error) {
	var o []*ServiceOfferingLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceOfferingLabel slice")
	}

	return o, nil
}

// Count returns the count of all ServiceOfferingLabel records in the query.
func (q serviceOfferingLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_offering_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceOfferingLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_offering_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q serviceOfferingLabelQuery) Resource(o *ServiceOfferingLabel, mods ...qm.QueryMod) serviceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Services(queryMods...)
	queries.SetFrom(query.Query, "`services`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceOfferingLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceOfferingLabel interface{}, mods queries.Applicator) error {
	var slice []*ServiceOfferingLabel
	var object *ServiceOfferingLabel

	if singular {
		object = maybeServiceOfferingLabel.(*ServiceOfferingLabel)
	} else {
		slice = *maybeServiceOfferingLabel.(*[]*ServiceOfferingLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceOfferingLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceOfferingLabelR{}
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
		foreign.R.ResourceServiceOfferingLabels = append(foreign.R.ResourceServiceOfferingLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &serviceR{}
				}
				foreign.R.ResourceServiceOfferingLabels = append(foreign.R.ResourceServiceOfferingLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the serviceOfferingLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceServiceOfferingLabels.
func (q serviceOfferingLabelQuery) SetResource(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Service) error {
	var err error
	if insert {
		if err = Services().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `service_offering_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, serviceOfferingLabelPrimaryKeyColumns),
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
		o.R = &serviceOfferingLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &serviceR{
			ResourceServiceOfferingLabels: ServiceOfferingLabelSlice{o},
		}
	} else {
		related.R.ResourceServiceOfferingLabels = append(related.R.ResourceServiceOfferingLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q serviceOfferingLabelQuery) RemoveResource(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, related *Service) error {
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

	for i, ri := range related.R.ResourceServiceOfferingLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceServiceOfferingLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceServiceOfferingLabels[i] = related.R.ResourceServiceOfferingLabels[ln-1]
		}
		related.R.ResourceServiceOfferingLabels = related.R.ResourceServiceOfferingLabels[:ln-1]
		break
	}
	return nil
}

// ServiceOfferingLabels retrieves all the records using an executor.
func ServiceOfferingLabels(mods ...qm.QueryMod) serviceOfferingLabelQuery {
	mods = append(mods, qm.From("`service_offering_labels`"))
	return serviceOfferingLabelQuery{NewQuery(mods...)}
}

type ServiceOfferingLabelFinder interface {
	FindServiceOfferingLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceOfferingLabel, error)
}

// FindServiceOfferingLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceOfferingLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceOfferingLabel, error) {
	serviceOfferingLabelObj := &ServiceOfferingLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `service_offering_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceOfferingLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_offering_labels")
	}

	return serviceOfferingLabelObj, nil
}

type ServiceOfferingLabelInserter interface {
	Insert(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q serviceOfferingLabelQuery) Insert(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_offering_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceOfferingLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceOfferingLabelInsertCacheMut.RLock()
	cache, cached := serviceOfferingLabelInsertCache[key]
	serviceOfferingLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceOfferingLabelAllColumns,
			serviceOfferingLabelColumnsWithDefault,
			serviceOfferingLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `service_offering_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `service_offering_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `service_offering_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, serviceOfferingLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into service_offering_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceOfferingLabelMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for service_offering_labels")
	}

CacheNoHooks:
	if !cached {
		serviceOfferingLabelInsertCacheMut.Lock()
		serviceOfferingLabelInsertCache[key] = cache
		serviceOfferingLabelInsertCacheMut.Unlock()
	}

	return nil
}

type ServiceOfferingLabelUpdater interface {
	Update(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the ServiceOfferingLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q serviceOfferingLabelQuery) Update(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	serviceOfferingLabelUpdateCacheMut.RLock()
	cache, cached := serviceOfferingLabelUpdateCache[key]
	serviceOfferingLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceOfferingLabelAllColumns,
			serviceOfferingLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_offering_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `service_offering_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, serviceOfferingLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceOfferingLabelType, serviceOfferingLabelMapping, append(wl, serviceOfferingLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_offering_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_offering_labels")
	}

	if !cached {
		serviceOfferingLabelUpdateCacheMut.Lock()
		serviceOfferingLabelUpdateCache[key] = cache
		serviceOfferingLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q serviceOfferingLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_offering_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_offering_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q serviceOfferingLabelQuery) UpdateAllSlice(o ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `service_offering_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceOfferingLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceOfferingLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceOfferingLabel")
	}
	return rowsAff, nil
}

type ServiceOfferingLabelDeleter interface {
	Delete(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single ServiceOfferingLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q serviceOfferingLabelQuery) Delete(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceOfferingLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceOfferingLabelPrimaryKeyMapping)
	sql := "DELETE FROM `service_offering_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_offering_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_offering_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceOfferingLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceOfferingLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_offering_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_offering_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q serviceOfferingLabelQuery) DeleteAllSlice(o ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `service_offering_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceOfferingLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceOfferingLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_offering_labels")
	}

	return rowsAff, nil
}

type ServiceOfferingLabelReloader interface {
	Reload(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q serviceOfferingLabelQuery) Reload(o *ServiceOfferingLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceOfferingLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q serviceOfferingLabelQuery) ReloadAll(o *ServiceOfferingLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceOfferingLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceOfferingLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `service_offering_labels`.* FROM `service_offering_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceOfferingLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceOfferingLabelSlice")
	}

	*o = slice

	return nil
}

// ServiceOfferingLabelExists checks if the ServiceOfferingLabel row exists.
func ServiceOfferingLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `service_offering_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_offering_labels exists")
	}

	return exists, nil
}
