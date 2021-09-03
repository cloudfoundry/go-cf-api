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

// QuotaDefinition is an object representing the database table.
type QuotaDefinition struct {
	ID                      int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                    string    `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt               time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name                    string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	NonBasicServicesAllowed bool      `boil:"non_basic_services_allowed" json:"non_basic_services_allowed" toml:"non_basic_services_allowed" yaml:"non_basic_services_allowed"`
	TotalServices           int       `boil:"total_services" json:"total_services" toml:"total_services" yaml:"total_services"`
	MemoryLimit             int       `boil:"memory_limit" json:"memory_limit" toml:"memory_limit" yaml:"memory_limit"`
	TotalRoutes             int       `boil:"total_routes" json:"total_routes" toml:"total_routes" yaml:"total_routes"`
	InstanceMemoryLimit     int       `boil:"instance_memory_limit" json:"instance_memory_limit" toml:"instance_memory_limit" yaml:"instance_memory_limit"`
	TotalPrivateDomains     int       `boil:"total_private_domains" json:"total_private_domains" toml:"total_private_domains" yaml:"total_private_domains"`
	AppInstanceLimit        null.Int  `boil:"app_instance_limit" json:"app_instance_limit,omitempty" toml:"app_instance_limit" yaml:"app_instance_limit,omitempty"`
	AppTaskLimit            null.Int  `boil:"app_task_limit" json:"app_task_limit,omitempty" toml:"app_task_limit" yaml:"app_task_limit,omitempty"`
	TotalServiceKeys        null.Int  `boil:"total_service_keys" json:"total_service_keys,omitempty" toml:"total_service_keys" yaml:"total_service_keys,omitempty"`
	TotalReservedRoutePorts null.Int  `boil:"total_reserved_route_ports" json:"total_reserved_route_ports,omitempty" toml:"total_reserved_route_ports" yaml:"total_reserved_route_ports,omitempty"`

	R *quotaDefinitionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L quotaDefinitionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var QuotaDefinitionColumns = struct {
	ID                      string
	GUID                    string
	CreatedAt               string
	UpdatedAt               string
	Name                    string
	NonBasicServicesAllowed string
	TotalServices           string
	MemoryLimit             string
	TotalRoutes             string
	InstanceMemoryLimit     string
	TotalPrivateDomains     string
	AppInstanceLimit        string
	AppTaskLimit            string
	TotalServiceKeys        string
	TotalReservedRoutePorts string
}{
	ID:                      "id",
	GUID:                    "guid",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
	Name:                    "name",
	NonBasicServicesAllowed: "non_basic_services_allowed",
	TotalServices:           "total_services",
	MemoryLimit:             "memory_limit",
	TotalRoutes:             "total_routes",
	InstanceMemoryLimit:     "instance_memory_limit",
	TotalPrivateDomains:     "total_private_domains",
	AppInstanceLimit:        "app_instance_limit",
	AppTaskLimit:            "app_task_limit",
	TotalServiceKeys:        "total_service_keys",
	TotalReservedRoutePorts: "total_reserved_route_ports",
}

var QuotaDefinitionTableColumns = struct {
	ID                      string
	GUID                    string
	CreatedAt               string
	UpdatedAt               string
	Name                    string
	NonBasicServicesAllowed string
	TotalServices           string
	MemoryLimit             string
	TotalRoutes             string
	InstanceMemoryLimit     string
	TotalPrivateDomains     string
	AppInstanceLimit        string
	AppTaskLimit            string
	TotalServiceKeys        string
	TotalReservedRoutePorts string
}{
	ID:                      "quota_definitions.id",
	GUID:                    "quota_definitions.guid",
	CreatedAt:               "quota_definitions.created_at",
	UpdatedAt:               "quota_definitions.updated_at",
	Name:                    "quota_definitions.name",
	NonBasicServicesAllowed: "quota_definitions.non_basic_services_allowed",
	TotalServices:           "quota_definitions.total_services",
	MemoryLimit:             "quota_definitions.memory_limit",
	TotalRoutes:             "quota_definitions.total_routes",
	InstanceMemoryLimit:     "quota_definitions.instance_memory_limit",
	TotalPrivateDomains:     "quota_definitions.total_private_domains",
	AppInstanceLimit:        "quota_definitions.app_instance_limit",
	AppTaskLimit:            "quota_definitions.app_task_limit",
	TotalServiceKeys:        "quota_definitions.total_service_keys",
	TotalReservedRoutePorts: "quota_definitions.total_reserved_route_ports",
}

// Generated where

var QuotaDefinitionWhere = struct {
	ID                      whereHelperint
	GUID                    whereHelperstring
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpernull_Time
	Name                    whereHelperstring
	NonBasicServicesAllowed whereHelperbool
	TotalServices           whereHelperint
	MemoryLimit             whereHelperint
	TotalRoutes             whereHelperint
	InstanceMemoryLimit     whereHelperint
	TotalPrivateDomains     whereHelperint
	AppInstanceLimit        whereHelpernull_Int
	AppTaskLimit            whereHelpernull_Int
	TotalServiceKeys        whereHelpernull_Int
	TotalReservedRoutePorts whereHelpernull_Int
}{
	ID:                      whereHelperint{field: "`quota_definitions`.`id`"},
	GUID:                    whereHelperstring{field: "`quota_definitions`.`guid`"},
	CreatedAt:               whereHelpertime_Time{field: "`quota_definitions`.`created_at`"},
	UpdatedAt:               whereHelpernull_Time{field: "`quota_definitions`.`updated_at`"},
	Name:                    whereHelperstring{field: "`quota_definitions`.`name`"},
	NonBasicServicesAllowed: whereHelperbool{field: "`quota_definitions`.`non_basic_services_allowed`"},
	TotalServices:           whereHelperint{field: "`quota_definitions`.`total_services`"},
	MemoryLimit:             whereHelperint{field: "`quota_definitions`.`memory_limit`"},
	TotalRoutes:             whereHelperint{field: "`quota_definitions`.`total_routes`"},
	InstanceMemoryLimit:     whereHelperint{field: "`quota_definitions`.`instance_memory_limit`"},
	TotalPrivateDomains:     whereHelperint{field: "`quota_definitions`.`total_private_domains`"},
	AppInstanceLimit:        whereHelpernull_Int{field: "`quota_definitions`.`app_instance_limit`"},
	AppTaskLimit:            whereHelpernull_Int{field: "`quota_definitions`.`app_task_limit`"},
	TotalServiceKeys:        whereHelpernull_Int{field: "`quota_definitions`.`total_service_keys`"},
	TotalReservedRoutePorts: whereHelpernull_Int{field: "`quota_definitions`.`total_reserved_route_ports`"},
}

// QuotaDefinitionRels is where relationship names are stored.
var QuotaDefinitionRels = struct {
	Organizations string
}{
	Organizations: "Organizations",
}

// quotaDefinitionR is where relationships are stored.
type quotaDefinitionR struct {
	Organizations OrganizationSlice `boil:"Organizations" json:"Organizations" toml:"Organizations" yaml:"Organizations"`
}

// NewStruct creates a new relationship struct
func (*quotaDefinitionR) NewStruct() *quotaDefinitionR {
	return &quotaDefinitionR{}
}

// quotaDefinitionL is where Load methods for each relationship are stored.
type quotaDefinitionL struct{}

var (
	quotaDefinitionAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "non_basic_services_allowed", "total_services", "memory_limit", "total_routes", "instance_memory_limit", "total_private_domains", "app_instance_limit", "app_task_limit", "total_service_keys", "total_reserved_route_ports"}
	quotaDefinitionColumnsWithoutDefault = []string{"guid", "updated_at", "name", "non_basic_services_allowed", "total_services", "memory_limit", "total_routes"}
	quotaDefinitionColumnsWithDefault    = []string{"id", "created_at", "instance_memory_limit", "total_private_domains", "app_instance_limit", "app_task_limit", "total_service_keys", "total_reserved_route_ports"}
	quotaDefinitionPrimaryKeyColumns     = []string{"id"}
)

type (
	// QuotaDefinitionSlice is an alias for a slice of pointers to QuotaDefinition.
	// This should almost always be used instead of []QuotaDefinition.
	QuotaDefinitionSlice []*QuotaDefinition

	quotaDefinitionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	quotaDefinitionType                 = reflect.TypeOf(&QuotaDefinition{})
	quotaDefinitionMapping              = queries.MakeStructMapping(quotaDefinitionType)
	quotaDefinitionPrimaryKeyMapping, _ = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, quotaDefinitionPrimaryKeyColumns)
	quotaDefinitionInsertCacheMut       sync.RWMutex
	quotaDefinitionInsertCache          = make(map[string]insertCache)
	quotaDefinitionUpdateCacheMut       sync.RWMutex
	quotaDefinitionUpdateCache          = make(map[string]updateCache)
	quotaDefinitionUpsertCacheMut       sync.RWMutex
	quotaDefinitionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type QuotaDefinitionFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*QuotaDefinition, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (QuotaDefinitionSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single quotaDefinition record from the query.
func (q quotaDefinitionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*QuotaDefinition, error) {
	o := &QuotaDefinition{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for quota_definitions")
	}

	return o, nil
}

// All returns all QuotaDefinition records from the query.
func (q quotaDefinitionQuery) All(ctx context.Context, exec boil.ContextExecutor) (QuotaDefinitionSlice, error) {
	var o []*QuotaDefinition

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to QuotaDefinition slice")
	}

	return o, nil
}

// Count returns the count of all QuotaDefinition records in the query.
func (q quotaDefinitionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count quota_definitions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q quotaDefinitionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if quota_definitions exists")
	}

	return count > 0, nil
}

// Organizations retrieves all the organization's Organizations with an executor.
func (q quotaDefinitionQuery) Organizations(o *QuotaDefinition, mods ...qm.QueryMod) organizationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`organizations`.`quota_definition_id`=?", o.ID),
	)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "`organizations`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`organizations`.*"})
	}

	return query
}

// LoadOrganizations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (quotaDefinitionL) LoadOrganizations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeQuotaDefinition interface{}, mods queries.Applicator) error {
	var slice []*QuotaDefinition
	var object *QuotaDefinition

	if singular {
		object = maybeQuotaDefinition.(*QuotaDefinition)
	} else {
		slice = *maybeQuotaDefinition.(*[]*QuotaDefinition)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &quotaDefinitionR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &quotaDefinitionR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`organizations`),
		qm.WhereIn(`organizations.quota_definition_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load organizations")
	}

	var resultSlice []*Organization
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice organizations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on organizations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for organizations")
	}

	if singular {
		object.R.Organizations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &organizationR{}
			}
			foreign.R.QuotaDefinition = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.QuotaDefinitionID {
				local.R.Organizations = append(local.R.Organizations, foreign)
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.QuotaDefinition = local
				break
			}
		}
	}

	return nil
}

// AddOrganizations adds the given related objects to the existing relationships
// of the quota_definition, optionally inserting them as new records.
// Appends related to o.R.Organizations.
// Sets related.R.QuotaDefinition appropriately.
func (q quotaDefinitionQuery) AddOrganizations(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Organization) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.QuotaDefinitionID = o.ID
			if err = Organizations().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `organizations` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"quota_definition_id"}),
				strmangle.WhereClause("`", "`", 0, organizationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.QuotaDefinitionID = o.ID
		}
	}

	if o.R == nil {
		o.R = &quotaDefinitionR{
			Organizations: related,
		}
	} else {
		o.R.Organizations = append(o.R.Organizations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &organizationR{
				QuotaDefinition: o,
			}
		} else {
			rel.R.QuotaDefinition = o
		}
	}
	return nil
}

// QuotaDefinitions retrieves all the records using an executor.
func QuotaDefinitions(mods ...qm.QueryMod) quotaDefinitionQuery {
	mods = append(mods, qm.From("`quota_definitions`"))
	return quotaDefinitionQuery{NewQuery(mods...)}
}

type QuotaDefinitionFinder interface {
	FindQuotaDefinition(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*QuotaDefinition, error)
}

// FindQuotaDefinition retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindQuotaDefinition(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*QuotaDefinition, error) {
	quotaDefinitionObj := &QuotaDefinition{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `quota_definitions` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, quotaDefinitionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from quota_definitions")
	}

	return quotaDefinitionObj, nil
}

type QuotaDefinitionInserter interface {
	Insert(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q quotaDefinitionQuery) Insert(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no quota_definitions provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(quotaDefinitionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	quotaDefinitionInsertCacheMut.RLock()
	cache, cached := quotaDefinitionInsertCache[key]
	quotaDefinitionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			quotaDefinitionAllColumns,
			quotaDefinitionColumnsWithDefault,
			quotaDefinitionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `quota_definitions` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `quota_definitions` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `quota_definitions` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, quotaDefinitionPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into quota_definitions")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == quotaDefinitionMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for quota_definitions")
	}

CacheNoHooks:
	if !cached {
		quotaDefinitionInsertCacheMut.Lock()
		quotaDefinitionInsertCache[key] = cache
		quotaDefinitionInsertCacheMut.Unlock()
	}

	return nil
}

type QuotaDefinitionUpdater interface {
	Update(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the QuotaDefinition.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q quotaDefinitionQuery) Update(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	quotaDefinitionUpdateCacheMut.RLock()
	cache, cached := quotaDefinitionUpdateCache[key]
	quotaDefinitionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			quotaDefinitionAllColumns,
			quotaDefinitionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update quota_definitions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `quota_definitions` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, quotaDefinitionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, append(wl, quotaDefinitionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update quota_definitions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for quota_definitions")
	}

	if !cached {
		quotaDefinitionUpdateCacheMut.Lock()
		quotaDefinitionUpdateCache[key] = cache
		quotaDefinitionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q quotaDefinitionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for quota_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for quota_definitions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q quotaDefinitionQuery) UpdateAllSlice(o QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotaDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `quota_definitions` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, quotaDefinitionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in quotaDefinition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all quotaDefinition")
	}
	return rowsAff, nil
}

type QuotaDefinitionDeleter interface {
	Delete(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single QuotaDefinition record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q quotaDefinitionQuery) Delete(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no QuotaDefinition provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), quotaDefinitionPrimaryKeyMapping)
	sql := "DELETE FROM `quota_definitions` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from quota_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for quota_definitions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q quotaDefinitionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no quotaDefinitionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from quota_definitions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for quota_definitions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q quotaDefinitionQuery) DeleteAllSlice(o QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotaDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `quota_definitions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, quotaDefinitionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from quotaDefinition slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for quota_definitions")
	}

	return rowsAff, nil
}

type QuotaDefinitionReloader interface {
	Reload(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q quotaDefinitionQuery) Reload(o *QuotaDefinition, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindQuotaDefinition(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q quotaDefinitionQuery) ReloadAll(o *QuotaDefinitionSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := QuotaDefinitionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), quotaDefinitionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `quota_definitions`.* FROM `quota_definitions` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, quotaDefinitionPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in QuotaDefinitionSlice")
	}

	*o = slice

	return nil
}

// QuotaDefinitionExists checks if the QuotaDefinition row exists.
func QuotaDefinitionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `quota_definitions` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if quota_definitions exists")
	}

	return exists, nil
}

var mySQLQuotaDefinitionUniqueColumns = []string{
	"id",
	"guid",
	"name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *QuotaDefinition) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no quota_definitions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(quotaDefinitionColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLQuotaDefinitionUniqueColumns, o)

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

	quotaDefinitionUpsertCacheMut.RLock()
	cache, cached := quotaDefinitionUpsertCache[key]
	quotaDefinitionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			quotaDefinitionAllColumns,
			quotaDefinitionColumnsWithDefault,
			quotaDefinitionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			quotaDefinitionAllColumns,
			quotaDefinitionPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert quota_definitions, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`quota_definitions`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `quota_definitions` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for quota_definitions")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == quotaDefinitionMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(quotaDefinitionType, quotaDefinitionMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for quota_definitions")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for quota_definitions")
	}

CacheNoHooks:
	if !cached {
		quotaDefinitionUpsertCacheMut.Lock()
		quotaDefinitionUpsertCache[key] = cache
		quotaDefinitionUpsertCacheMut.Unlock()
	}

	return nil
}
