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

// ServiceBrokerLabel is an object representing the database table.
type ServiceBrokerLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *serviceBrokerLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L serviceBrokerLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ServiceBrokerLabelColumns = struct {
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

var ServiceBrokerLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "service_broker_labels.id",
	GUID:         "service_broker_labels.guid",
	CreatedAt:    "service_broker_labels.created_at",
	UpdatedAt:    "service_broker_labels.updated_at",
	ResourceGUID: "service_broker_labels.resource_guid",
	KeyPrefix:    "service_broker_labels.key_prefix",
	KeyName:      "service_broker_labels.key_name",
	Value:        "service_broker_labels.value",
}

// Generated where

var ServiceBrokerLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`service_broker_labels`.`id`"},
	GUID:         whereHelperstring{field: "`service_broker_labels`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`service_broker_labels`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`service_broker_labels`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`service_broker_labels`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`service_broker_labels`.`key_prefix`"},
	KeyName:      whereHelpernull_String{field: "`service_broker_labels`.`key_name`"},
	Value:        whereHelpernull_String{field: "`service_broker_labels`.`value`"},
}

// ServiceBrokerLabelRels is where relationship names are stored.
var ServiceBrokerLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// serviceBrokerLabelR is where relationships are stored.
type serviceBrokerLabelR struct {
	Resource *ServiceBroker `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*serviceBrokerLabelR) NewStruct() *serviceBrokerLabelR {
	return &serviceBrokerLabelR{}
}

// serviceBrokerLabelL is where Load methods for each relationship are stored.
type serviceBrokerLabelL struct{}

var (
	serviceBrokerLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceBrokerLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	serviceBrokerLabelColumnsWithDefault    = []string{"id", "created_at"}
	serviceBrokerLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// ServiceBrokerLabelSlice is an alias for a slice of pointers to ServiceBrokerLabel.
	// This should almost always be used instead of []ServiceBrokerLabel.
	ServiceBrokerLabelSlice []*ServiceBrokerLabel
	// ServiceBrokerLabelHook is the signature for custom ServiceBrokerLabel hook methods
	ServiceBrokerLabelHook func(context.Context, boil.ContextExecutor, *ServiceBrokerLabel) error

	serviceBrokerLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	serviceBrokerLabelType                 = reflect.TypeOf(&ServiceBrokerLabel{})
	serviceBrokerLabelMapping              = queries.MakeStructMapping(serviceBrokerLabelType)
	serviceBrokerLabelPrimaryKeyMapping, _ = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, serviceBrokerLabelPrimaryKeyColumns)
	serviceBrokerLabelInsertCacheMut       sync.RWMutex
	serviceBrokerLabelInsertCache          = make(map[string]insertCache)
	serviceBrokerLabelUpdateCacheMut       sync.RWMutex
	serviceBrokerLabelUpdateCache          = make(map[string]updateCache)
	serviceBrokerLabelUpsertCacheMut       sync.RWMutex
	serviceBrokerLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var serviceBrokerLabelBeforeInsertHooks []ServiceBrokerLabelHook
var serviceBrokerLabelBeforeUpdateHooks []ServiceBrokerLabelHook
var serviceBrokerLabelBeforeDeleteHooks []ServiceBrokerLabelHook
var serviceBrokerLabelBeforeUpsertHooks []ServiceBrokerLabelHook

var serviceBrokerLabelAfterInsertHooks []ServiceBrokerLabelHook
var serviceBrokerLabelAfterSelectHooks []ServiceBrokerLabelHook
var serviceBrokerLabelAfterUpdateHooks []ServiceBrokerLabelHook
var serviceBrokerLabelAfterDeleteHooks []ServiceBrokerLabelHook
var serviceBrokerLabelAfterUpsertHooks []ServiceBrokerLabelHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ServiceBrokerLabel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ServiceBrokerLabel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ServiceBrokerLabel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ServiceBrokerLabel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ServiceBrokerLabel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ServiceBrokerLabel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ServiceBrokerLabel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ServiceBrokerLabel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ServiceBrokerLabel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range serviceBrokerLabelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddServiceBrokerLabelHook registers your hook function for all future operations.
func AddServiceBrokerLabelHook(hookPoint boil.HookPoint, serviceBrokerLabelHook ServiceBrokerLabelHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		serviceBrokerLabelBeforeInsertHooks = append(serviceBrokerLabelBeforeInsertHooks, serviceBrokerLabelHook)
	case boil.BeforeUpdateHook:
		serviceBrokerLabelBeforeUpdateHooks = append(serviceBrokerLabelBeforeUpdateHooks, serviceBrokerLabelHook)
	case boil.BeforeDeleteHook:
		serviceBrokerLabelBeforeDeleteHooks = append(serviceBrokerLabelBeforeDeleteHooks, serviceBrokerLabelHook)
	case boil.BeforeUpsertHook:
		serviceBrokerLabelBeforeUpsertHooks = append(serviceBrokerLabelBeforeUpsertHooks, serviceBrokerLabelHook)
	case boil.AfterInsertHook:
		serviceBrokerLabelAfterInsertHooks = append(serviceBrokerLabelAfterInsertHooks, serviceBrokerLabelHook)
	case boil.AfterSelectHook:
		serviceBrokerLabelAfterSelectHooks = append(serviceBrokerLabelAfterSelectHooks, serviceBrokerLabelHook)
	case boil.AfterUpdateHook:
		serviceBrokerLabelAfterUpdateHooks = append(serviceBrokerLabelAfterUpdateHooks, serviceBrokerLabelHook)
	case boil.AfterDeleteHook:
		serviceBrokerLabelAfterDeleteHooks = append(serviceBrokerLabelAfterDeleteHooks, serviceBrokerLabelHook)
	case boil.AfterUpsertHook:
		serviceBrokerLabelAfterUpsertHooks = append(serviceBrokerLabelAfterUpsertHooks, serviceBrokerLabelHook)
	}
}

// One returns a single serviceBrokerLabel record from the query.
func (q serviceBrokerLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ServiceBrokerLabel, error) {
	o := &ServiceBrokerLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for service_broker_labels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ServiceBrokerLabel records from the query.
func (q serviceBrokerLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ServiceBrokerLabelSlice, error) {
	var o []*ServiceBrokerLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ServiceBrokerLabel slice")
	}

	if len(serviceBrokerLabelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ServiceBrokerLabel records in the query.
func (q serviceBrokerLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count service_broker_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q serviceBrokerLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if service_broker_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (o *ServiceBrokerLabel) Resource(mods ...qm.QueryMod) serviceBrokerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := ServiceBrokers(queryMods...)
	queries.SetFrom(query.Query, "`service_brokers`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (serviceBrokerLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeServiceBrokerLabel interface{}, mods queries.Applicator) error {
	var slice []*ServiceBrokerLabel
	var object *ServiceBrokerLabel

	if singular {
		object = maybeServiceBrokerLabel.(*ServiceBrokerLabel)
	} else {
		slice = *maybeServiceBrokerLabel.(*[]*ServiceBrokerLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &serviceBrokerLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &serviceBrokerLabelR{}
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
		qm.From(`service_brokers`),
		qm.WhereIn(`service_brokers.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ServiceBroker")
	}

	var resultSlice []*ServiceBroker
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ServiceBroker")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for service_brokers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for service_brokers")
	}

	if len(serviceBrokerLabelAfterSelectHooks) != 0 {
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
			foreign.R = &serviceBrokerR{}
		}
		foreign.R.ResourceServiceBrokerLabels = append(foreign.R.ResourceServiceBrokerLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &serviceBrokerR{}
				}
				foreign.R.ResourceServiceBrokerLabels = append(foreign.R.ResourceServiceBrokerLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the serviceBrokerLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceServiceBrokerLabels.
func (o *ServiceBrokerLabel) SetResource(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ServiceBroker) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `service_broker_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, serviceBrokerLabelPrimaryKeyColumns),
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
		o.R = &serviceBrokerLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &serviceBrokerR{
			ResourceServiceBrokerLabels: ServiceBrokerLabelSlice{o},
		}
	} else {
		related.R.ResourceServiceBrokerLabels = append(related.R.ResourceServiceBrokerLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *ServiceBrokerLabel) RemoveResource(ctx context.Context, exec boil.ContextExecutor, related *ServiceBroker) error {
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

	for i, ri := range related.R.ResourceServiceBrokerLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceServiceBrokerLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceServiceBrokerLabels[i] = related.R.ResourceServiceBrokerLabels[ln-1]
		}
		related.R.ResourceServiceBrokerLabels = related.R.ResourceServiceBrokerLabels[:ln-1]
		break
	}
	return nil
}

// ServiceBrokerLabels retrieves all the records using an executor.
func ServiceBrokerLabels(mods ...qm.QueryMod) serviceBrokerLabelQuery {
	mods = append(mods, qm.From("`service_broker_labels`"))
	return serviceBrokerLabelQuery{NewQuery(mods...)}
}

// FindServiceBrokerLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindServiceBrokerLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*ServiceBrokerLabel, error) {
	serviceBrokerLabelObj := &ServiceBrokerLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `service_broker_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, serviceBrokerLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from service_broker_labels")
	}

	if err = serviceBrokerLabelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return serviceBrokerLabelObj, err
	}

	return serviceBrokerLabelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ServiceBrokerLabel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_broker_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceBrokerLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	serviceBrokerLabelInsertCacheMut.RLock()
	cache, cached := serviceBrokerLabelInsertCache[key]
	serviceBrokerLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			serviceBrokerLabelAllColumns,
			serviceBrokerLabelColumnsWithDefault,
			serviceBrokerLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `service_broker_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `service_broker_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `service_broker_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, serviceBrokerLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into service_broker_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceBrokerLabelMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for service_broker_labels")
	}

CacheNoHooks:
	if !cached {
		serviceBrokerLabelInsertCacheMut.Lock()
		serviceBrokerLabelInsertCache[key] = cache
		serviceBrokerLabelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ServiceBrokerLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ServiceBrokerLabel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	serviceBrokerLabelUpdateCacheMut.RLock()
	cache, cached := serviceBrokerLabelUpdateCache[key]
	serviceBrokerLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			serviceBrokerLabelAllColumns,
			serviceBrokerLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update service_broker_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `service_broker_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, serviceBrokerLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, append(wl, serviceBrokerLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update service_broker_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for service_broker_labels")
	}

	if !cached {
		serviceBrokerLabelUpdateCacheMut.Lock()
		serviceBrokerLabelUpdateCache[key] = cache
		serviceBrokerLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q serviceBrokerLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for service_broker_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for service_broker_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ServiceBrokerLabelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBrokerLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `service_broker_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceBrokerLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in serviceBrokerLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all serviceBrokerLabel")
	}
	return rowsAff, nil
}

var mySQLServiceBrokerLabelUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ServiceBrokerLabel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no service_broker_labels provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(serviceBrokerLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLServiceBrokerLabelUniqueColumns, o)

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

	serviceBrokerLabelUpsertCacheMut.RLock()
	cache, cached := serviceBrokerLabelUpsertCache[key]
	serviceBrokerLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			serviceBrokerLabelAllColumns,
			serviceBrokerLabelColumnsWithDefault,
			serviceBrokerLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			serviceBrokerLabelAllColumns,
			serviceBrokerLabelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert service_broker_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`service_broker_labels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `service_broker_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for service_broker_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == serviceBrokerLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(serviceBrokerLabelType, serviceBrokerLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for service_broker_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for service_broker_labels")
	}

CacheNoHooks:
	if !cached {
		serviceBrokerLabelUpsertCacheMut.Lock()
		serviceBrokerLabelUpsertCache[key] = cache
		serviceBrokerLabelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ServiceBrokerLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ServiceBrokerLabel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ServiceBrokerLabel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), serviceBrokerLabelPrimaryKeyMapping)
	sql := "DELETE FROM `service_broker_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from service_broker_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for service_broker_labels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q serviceBrokerLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no serviceBrokerLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from service_broker_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_broker_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ServiceBrokerLabelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(serviceBrokerLabelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBrokerLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `service_broker_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceBrokerLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from serviceBrokerLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for service_broker_labels")
	}

	if len(serviceBrokerLabelAfterDeleteHooks) != 0 {
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
func (o *ServiceBrokerLabel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindServiceBrokerLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ServiceBrokerLabelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ServiceBrokerLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), serviceBrokerLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `service_broker_labels`.* FROM `service_broker_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, serviceBrokerLabelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ServiceBrokerLabelSlice")
	}

	*o = slice

	return nil
}

// ServiceBrokerLabelExists checks if the ServiceBrokerLabel row exists.
func ServiceBrokerLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `service_broker_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if service_broker_labels exists")
	}

	return exists, nil
}
