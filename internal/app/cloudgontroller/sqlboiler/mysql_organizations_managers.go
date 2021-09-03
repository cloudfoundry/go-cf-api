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

// OrganizationsManager is an object representing the database table.
type OrganizationsManager struct {
	OrganizationID          int         `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	UserID                  int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	OrganizationsManagersPK int         `boil:"organizations_managers_pk" json:"organizations_managers_pk" toml:"organizations_managers_pk" yaml:"organizations_managers_pk"`
	RoleGUID                null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *organizationsManagerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsManagerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsManagerColumns = struct {
	OrganizationID          string
	UserID                  string
	OrganizationsManagersPK string
	RoleGUID                string
	CreatedAt               string
	UpdatedAt               string
}{
	OrganizationID:          "organization_id",
	UserID:                  "user_id",
	OrganizationsManagersPK: "organizations_managers_pk",
	RoleGUID:                "role_guid",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
}

var OrganizationsManagerTableColumns = struct {
	OrganizationID          string
	UserID                  string
	OrganizationsManagersPK string
	RoleGUID                string
	CreatedAt               string
	UpdatedAt               string
}{
	OrganizationID:          "organizations_managers.organization_id",
	UserID:                  "organizations_managers.user_id",
	OrganizationsManagersPK: "organizations_managers.organizations_managers_pk",
	RoleGUID:                "organizations_managers.role_guid",
	CreatedAt:               "organizations_managers.created_at",
	UpdatedAt:               "organizations_managers.updated_at",
}

// Generated where

var OrganizationsManagerWhere = struct {
	OrganizationID          whereHelperint
	UserID                  whereHelperint
	OrganizationsManagersPK whereHelperint
	RoleGUID                whereHelpernull_String
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpertime_Time
}{
	OrganizationID:          whereHelperint{field: "`organizations_managers`.`organization_id`"},
	UserID:                  whereHelperint{field: "`organizations_managers`.`user_id`"},
	OrganizationsManagersPK: whereHelperint{field: "`organizations_managers`.`organizations_managers_pk`"},
	RoleGUID:                whereHelpernull_String{field: "`organizations_managers`.`role_guid`"},
	CreatedAt:               whereHelpertime_Time{field: "`organizations_managers`.`created_at`"},
	UpdatedAt:               whereHelpertime_Time{field: "`organizations_managers`.`updated_at`"},
}

// OrganizationsManagerRels is where relationship names are stored.
var OrganizationsManagerRels = struct {
	Organization string
	User         string
}{
	Organization: "Organization",
	User:         "User",
}

// organizationsManagerR is where relationships are stored.
type organizationsManagerR struct {
	Organization *Organization `boil:"Organization" json:"Organization" toml:"Organization" yaml:"Organization"`
	User         *User         `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*organizationsManagerR) NewStruct() *organizationsManagerR {
	return &organizationsManagerR{}
}

// organizationsManagerL is where Load methods for each relationship are stored.
type organizationsManagerL struct{}

var (
	organizationsManagerAllColumns            = []string{"organization_id", "user_id", "organizations_managers_pk", "role_guid", "created_at", "updated_at"}
	organizationsManagerColumnsWithoutDefault = []string{"organization_id", "user_id", "role_guid"}
	organizationsManagerColumnsWithDefault    = []string{"organizations_managers_pk", "created_at", "updated_at"}
	organizationsManagerPrimaryKeyColumns     = []string{"organizations_managers_pk"}
)

type (
	// OrganizationsManagerSlice is an alias for a slice of pointers to OrganizationsManager.
	// This should almost always be used instead of []OrganizationsManager.
	OrganizationsManagerSlice []*OrganizationsManager

	organizationsManagerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsManagerType                 = reflect.TypeOf(&OrganizationsManager{})
	organizationsManagerMapping              = queries.MakeStructMapping(organizationsManagerType)
	organizationsManagerPrimaryKeyMapping, _ = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, organizationsManagerPrimaryKeyColumns)
	organizationsManagerInsertCacheMut       sync.RWMutex
	organizationsManagerInsertCache          = make(map[string]insertCache)
	organizationsManagerUpdateCacheMut       sync.RWMutex
	organizationsManagerUpdateCache          = make(map[string]updateCache)
	organizationsManagerUpsertCacheMut       sync.RWMutex
	organizationsManagerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrganizationsManagerFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsManager, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsManagerSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single organizationsManager record from the query.
func (q organizationsManagerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsManager, error) {
	o := &OrganizationsManager{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_managers")
	}

	return o, nil
}

// All returns all OrganizationsManager records from the query.
func (q organizationsManagerQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsManagerSlice, error) {
	var o []*OrganizationsManager

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsManager slice")
	}

	return o, nil
}

// Count returns the count of all OrganizationsManager records in the query.
func (q organizationsManagerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_managers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsManagerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_managers exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (q organizationsManagerQuery) Organization(o *OrganizationsManager, mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "`organizations`")

	return query
}

// User pointed to by the foreign key.
func (q organizationsManagerQuery) User(o *OrganizationsManager, mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadOrganization allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsManagerL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsManager interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsManager
	var object *OrganizationsManager

	if singular {
		object = maybeOrganizationsManager.(*OrganizationsManager)
	} else {
		slice = *maybeOrganizationsManager.(*[]*OrganizationsManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsManagerR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsManagerR{}
			}

			for _, a := range args {
				if a == obj.OrganizationID {
					continue Outer
				}
			}

			args = append(args, obj.OrganizationID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`organizations`),
		qm.WhereIn(`organizations.id in ?`, args...),
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
		object.R.Organization = foreign
		if foreign.R == nil {
			foreign.R = &organizationR{}
		}
		foreign.R.OrganizationsManagers = append(foreign.R.OrganizationsManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.OrganizationsManagers = append(foreign.R.OrganizationsManagers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsManagerL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsManager interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsManager
	var object *OrganizationsManager

	if singular {
		object = maybeOrganizationsManager.(*OrganizationsManager)
	} else {
		slice = *maybeOrganizationsManager.(*[]*OrganizationsManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsManagerR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsManagerR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.OrganizationsManagers = append(foreign.R.OrganizationsManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OrganizationsManagers = append(foreign.R.OrganizationsManagers, local)
				break
			}
		}
	}

	return nil
}

// SetOrganization of the organizationsManager to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.OrganizationsManagers.
func (q organizationsManagerQuery) SetOrganization(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = Organizations().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"organization_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsManagersPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrganizationID = related.ID
	if o.R == nil {
		o.R = &organizationsManagerR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			OrganizationsManagers: OrganizationsManagerSlice{o},
		}
	} else {
		related.R.OrganizationsManagers = append(related.R.OrganizationsManagers, o)
	}

	return nil
}

// SetUser of the organizationsManager to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OrganizationsManagers.
func (q organizationsManagerQuery) SetUser(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = Users().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsManagersPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &organizationsManagerR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OrganizationsManagers: OrganizationsManagerSlice{o},
		}
	} else {
		related.R.OrganizationsManagers = append(related.R.OrganizationsManagers, o)
	}

	return nil
}

// OrganizationsManagers retrieves all the records using an executor.
func OrganizationsManagers(mods ...qm.QueryMod) organizationsManagerQuery {
	mods = append(mods, qm.From("`organizations_managers`"))
	return organizationsManagerQuery{NewQuery(mods...)}
}

type OrganizationsManagerFinder interface {
	FindOrganizationsManager(ctx context.Context, exec boil.ContextExecutor, organizationsManagersPK int, selectCols ...string) (*OrganizationsManager, error)
}

// FindOrganizationsManager retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsManager(ctx context.Context, exec boil.ContextExecutor, organizationsManagersPK int, selectCols ...string) (*OrganizationsManager, error) {
	organizationsManagerObj := &OrganizationsManager{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `organizations_managers` where `organizations_managers_pk`=?", sel,
	)

	q := queries.Raw(query, organizationsManagersPK)

	err := q.Bind(ctx, exec, organizationsManagerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_managers")
	}

	return organizationsManagerObj, nil
}

type OrganizationsManagerInserter interface {
	Insert(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q organizationsManagerQuery) Insert(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_managers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsManagerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsManagerInsertCacheMut.RLock()
	cache, cached := organizationsManagerInsertCache[key]
	organizationsManagerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsManagerAllColumns,
			organizationsManagerColumnsWithDefault,
			organizationsManagerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `organizations_managers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `organizations_managers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `organizations_managers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, organizationsManagerPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into organizations_managers")
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

	o.OrganizationsManagersPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsManagerMapping["organizations_managers_pk"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OrganizationsManagersPK,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_managers")
	}

CacheNoHooks:
	if !cached {
		organizationsManagerInsertCacheMut.Lock()
		organizationsManagerInsertCache[key] = cache
		organizationsManagerInsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsManagerUpdater interface {
	Update(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrganizationsManager.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q organizationsManagerQuery) Update(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	organizationsManagerUpdateCacheMut.RLock()
	cache, cached := organizationsManagerUpdateCache[key]
	organizationsManagerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsManagerAllColumns,
			organizationsManagerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_managers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `organizations_managers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, organizationsManagerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, append(wl, organizationsManagerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_managers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_managers")
	}

	if !cached {
		organizationsManagerUpdateCacheMut.Lock()
		organizationsManagerUpdateCache[key] = cache
		organizationsManagerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsManagerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_managers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q organizationsManagerQuery) UpdateAllSlice(o OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `organizations_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsManagerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsManager")
	}
	return rowsAff, nil
}

type OrganizationsManagerUpserter interface {
	Upsert(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLOrganizationsManagerUniqueColumns = []string{
	"organizations_managers_pk",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q organizationsManagerQuery) Upsert(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_managers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsManagerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrganizationsManagerUniqueColumns, o)

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

	organizationsManagerUpsertCacheMut.RLock()
	cache, cached := organizationsManagerUpsertCache[key]
	organizationsManagerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsManagerAllColumns,
			organizationsManagerColumnsWithDefault,
			organizationsManagerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationsManagerAllColumns,
			organizationsManagerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_managers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`organizations_managers`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `organizations_managers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organizations_managers")
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

	o.OrganizationsManagersPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsManagerMapping["organizations_managers_pk"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(organizationsManagerType, organizationsManagerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for organizations_managers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_managers")
	}

CacheNoHooks:
	if !cached {
		organizationsManagerUpsertCacheMut.Lock()
		organizationsManagerUpsertCache[key] = cache
		organizationsManagerUpsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsManagerDeleter interface {
	Delete(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrganizationsManager record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q organizationsManagerQuery) Delete(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsManager provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsManagerPrimaryKeyMapping)
	sql := "DELETE FROM `organizations_managers` WHERE `organizations_managers_pk`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_managers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsManagerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsManagerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_managers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q organizationsManagerQuery) DeleteAllSlice(o OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `organizations_managers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsManagerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_managers")
	}

	return rowsAff, nil
}

type OrganizationsManagerReloader interface {
	Reload(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q organizationsManagerQuery) Reload(o *OrganizationsManager, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsManager(ctx, exec, o.OrganizationsManagersPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q organizationsManagerQuery) ReloadAll(o *OrganizationsManagerSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsManagerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `organizations_managers`.* FROM `organizations_managers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsManagerPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsManagerSlice")
	}

	*o = slice

	return nil
}

// OrganizationsManagerExists checks if the OrganizationsManager row exists.
func OrganizationsManagerExists(ctx context.Context, exec boil.ContextExecutor, organizationsManagersPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `organizations_managers` where `organizations_managers_pk`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, organizationsManagersPK)
	}
	row := exec.QueryRowContext(ctx, sql, organizationsManagersPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_managers exists")
	}

	return exists, nil
}
