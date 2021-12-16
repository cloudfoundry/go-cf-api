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

type OrganizationsBillingManagerUpserter interface {
	Upsert(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLOrganizationsBillingManagerUniqueColumns = []string{
	"organizations_billing_managers_pk",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q organizationsBillingManagerQuery) Upsert(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_billing_managers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsBillingManagerColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrganizationsBillingManagerUniqueColumns, o)

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

	organizationsBillingManagerUpsertCacheMut.RLock()
	cache, cached := organizationsBillingManagerUpsertCache[key]
	organizationsBillingManagerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsBillingManagerAllColumns,
			organizationsBillingManagerColumnsWithDefault,
			organizationsBillingManagerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationsBillingManagerAllColumns,
			organizationsBillingManagerPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_billing_managers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`organizations_billing_managers`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `organizations_billing_managers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organizations_billing_managers")
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

	o.OrganizationsBillingManagersPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsBillingManagerMapping["organizations_billing_managers_pk"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for organizations_billing_managers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_billing_managers")
	}

CacheNoHooks:
	if !cached {
		organizationsBillingManagerUpsertCacheMut.Lock()
		organizationsBillingManagerUpsertCache[key] = cache
		organizationsBillingManagerUpsertCacheMut.Unlock()
	}

	return nil
}

// OrganizationsBillingManager is an object representing the database table.
type OrganizationsBillingManager struct {
	OrganizationID                 int         `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	UserID                         int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	OrganizationsBillingManagersPK int         `boil:"organizations_billing_managers_pk" json:"organizations_billing_managers_pk" toml:"organizations_billing_managers_pk" yaml:"organizations_billing_managers_pk"`
	RoleGUID                       null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt                      time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                      time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *organizationsBillingManagerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsBillingManagerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsBillingManagerColumns = struct {
	OrganizationID                 string
	UserID                         string
	OrganizationsBillingManagersPK string
	RoleGUID                       string
	CreatedAt                      string
	UpdatedAt                      string
}{
	OrganizationID:                 "organization_id",
	UserID:                         "user_id",
	OrganizationsBillingManagersPK: "organizations_billing_managers_pk",
	RoleGUID:                       "role_guid",
	CreatedAt:                      "created_at",
	UpdatedAt:                      "updated_at",
}

var OrganizationsBillingManagerTableColumns = struct {
	OrganizationID                 string
	UserID                         string
	OrganizationsBillingManagersPK string
	RoleGUID                       string
	CreatedAt                      string
	UpdatedAt                      string
}{
	OrganizationID:                 "organizations_billing_managers.organization_id",
	UserID:                         "organizations_billing_managers.user_id",
	OrganizationsBillingManagersPK: "organizations_billing_managers.organizations_billing_managers_pk",
	RoleGUID:                       "organizations_billing_managers.role_guid",
	CreatedAt:                      "organizations_billing_managers.created_at",
	UpdatedAt:                      "organizations_billing_managers.updated_at",
}

// Generated where

var OrganizationsBillingManagerWhere = struct {
	OrganizationID                 whereHelperint
	UserID                         whereHelperint
	OrganizationsBillingManagersPK whereHelperint
	RoleGUID                       whereHelpernull_String
	CreatedAt                      whereHelpertime_Time
	UpdatedAt                      whereHelpertime_Time
}{
	OrganizationID:                 whereHelperint{field: "`organizations_billing_managers`.`organization_id`"},
	UserID:                         whereHelperint{field: "`organizations_billing_managers`.`user_id`"},
	OrganizationsBillingManagersPK: whereHelperint{field: "`organizations_billing_managers`.`organizations_billing_managers_pk`"},
	RoleGUID:                       whereHelpernull_String{field: "`organizations_billing_managers`.`role_guid`"},
	CreatedAt:                      whereHelpertime_Time{field: "`organizations_billing_managers`.`created_at`"},
	UpdatedAt:                      whereHelpertime_Time{field: "`organizations_billing_managers`.`updated_at`"},
}

// OrganizationsBillingManagerRels is where relationship names are stored.
var OrganizationsBillingManagerRels = struct {
	Organization string
	User         string
}{
	Organization: "Organization",
	User:         "User",
}

// organizationsBillingManagerR is where relationships are stored.
type organizationsBillingManagerR struct {
	Organization *Organization `boil:"Organization" json:"Organization" toml:"Organization" yaml:"Organization"`
	User         *User         `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*organizationsBillingManagerR) NewStruct() *organizationsBillingManagerR {
	return &organizationsBillingManagerR{}
}

// organizationsBillingManagerL is where Load methods for each relationship are stored.
type organizationsBillingManagerL struct{}

var (
	organizationsBillingManagerAllColumns            = []string{"organization_id", "user_id", "organizations_billing_managers_pk", "role_guid", "created_at", "updated_at"}
	organizationsBillingManagerColumnsWithoutDefault = []string{"organization_id", "user_id", "role_guid"}
	organizationsBillingManagerColumnsWithDefault    = []string{"organizations_billing_managers_pk", "created_at", "updated_at"}
	organizationsBillingManagerPrimaryKeyColumns     = []string{"organizations_billing_managers_pk"}
)

type (
	// OrganizationsBillingManagerSlice is an alias for a slice of pointers to OrganizationsBillingManager.
	// This should almost always be used instead of []OrganizationsBillingManager.
	OrganizationsBillingManagerSlice []*OrganizationsBillingManager

	organizationsBillingManagerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsBillingManagerType                 = reflect.TypeOf(&OrganizationsBillingManager{})
	organizationsBillingManagerMapping              = queries.MakeStructMapping(organizationsBillingManagerType)
	organizationsBillingManagerPrimaryKeyMapping, _ = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, organizationsBillingManagerPrimaryKeyColumns)
	organizationsBillingManagerInsertCacheMut       sync.RWMutex
	organizationsBillingManagerInsertCache          = make(map[string]insertCache)
	organizationsBillingManagerUpdateCacheMut       sync.RWMutex
	organizationsBillingManagerUpdateCache          = make(map[string]updateCache)
	organizationsBillingManagerUpsertCacheMut       sync.RWMutex
	organizationsBillingManagerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrganizationsBillingManagerFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsBillingManager, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsBillingManagerSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single organizationsBillingManager record from the query.
func (q organizationsBillingManagerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsBillingManager, error) {
	o := &OrganizationsBillingManager{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_billing_managers")
	}

	return o, nil
}

// All returns all OrganizationsBillingManager records from the query.
func (q organizationsBillingManagerQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsBillingManagerSlice, error) {
	var o []*OrganizationsBillingManager

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsBillingManager slice")
	}

	return o, nil
}

// Count returns the count of all OrganizationsBillingManager records in the query.
func (q organizationsBillingManagerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_billing_managers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsBillingManagerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_billing_managers exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (q organizationsBillingManagerQuery) Organization(o *OrganizationsBillingManager, mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "`organizations`")

	return query
}

// User pointed to by the foreign key.
func (q organizationsBillingManagerQuery) User(o *OrganizationsBillingManager, mods ...qm.QueryMod) userQuery {
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
func (organizationsBillingManagerL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsBillingManager interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsBillingManager
	var object *OrganizationsBillingManager

	if singular {
		object = maybeOrganizationsBillingManager.(*OrganizationsBillingManager)
	} else {
		slice = *maybeOrganizationsBillingManager.(*[]*OrganizationsBillingManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsBillingManagerR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsBillingManagerR{}
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
		foreign.R.OrganizationsBillingManagers = append(foreign.R.OrganizationsBillingManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.OrganizationsBillingManagers = append(foreign.R.OrganizationsBillingManagers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsBillingManagerL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsBillingManager interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsBillingManager
	var object *OrganizationsBillingManager

	if singular {
		object = maybeOrganizationsBillingManager.(*OrganizationsBillingManager)
	} else {
		slice = *maybeOrganizationsBillingManager.(*[]*OrganizationsBillingManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsBillingManagerR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsBillingManagerR{}
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
		foreign.R.OrganizationsBillingManagers = append(foreign.R.OrganizationsBillingManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OrganizationsBillingManagers = append(foreign.R.OrganizationsBillingManagers, local)
				break
			}
		}
	}

	return nil
}

// SetOrganization of the organizationsBillingManager to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.OrganizationsBillingManagers.
func (q organizationsBillingManagerQuery) SetOrganization(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = Organizations().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_billing_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"organization_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsBillingManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsBillingManagersPK}

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
		o.R = &organizationsBillingManagerR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			OrganizationsBillingManagers: OrganizationsBillingManagerSlice{o},
		}
	} else {
		related.R.OrganizationsBillingManagers = append(related.R.OrganizationsBillingManagers, o)
	}

	return nil
}

// SetUser of the organizationsBillingManager to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OrganizationsBillingManagers.
func (q organizationsBillingManagerQuery) SetUser(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = Users().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_billing_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsBillingManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsBillingManagersPK}

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
		o.R = &organizationsBillingManagerR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OrganizationsBillingManagers: OrganizationsBillingManagerSlice{o},
		}
	} else {
		related.R.OrganizationsBillingManagers = append(related.R.OrganizationsBillingManagers, o)
	}

	return nil
}

// OrganizationsBillingManagers retrieves all the records using an executor.
func OrganizationsBillingManagers(mods ...qm.QueryMod) organizationsBillingManagerQuery {
	mods = append(mods, qm.From("`organizations_billing_managers`"))
	return organizationsBillingManagerQuery{NewQuery(mods...)}
}

type OrganizationsBillingManagerFinder interface {
	FindOrganizationsBillingManager(ctx context.Context, exec boil.ContextExecutor, organizationsBillingManagersPK int, selectCols ...string) (*OrganizationsBillingManager, error)
}

// FindOrganizationsBillingManager retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsBillingManager(ctx context.Context, exec boil.ContextExecutor, organizationsBillingManagersPK int, selectCols ...string) (*OrganizationsBillingManager, error) {
	organizationsBillingManagerObj := &OrganizationsBillingManager{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `organizations_billing_managers` where `organizations_billing_managers_pk`=?", sel,
	)

	q := queries.Raw(query, organizationsBillingManagersPK)

	err := q.Bind(ctx, exec, organizationsBillingManagerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_billing_managers")
	}

	return organizationsBillingManagerObj, nil
}

type OrganizationsBillingManagerInserter interface {
	Insert(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q organizationsBillingManagerQuery) Insert(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_billing_managers provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(organizationsBillingManagerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsBillingManagerInsertCacheMut.RLock()
	cache, cached := organizationsBillingManagerInsertCache[key]
	organizationsBillingManagerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsBillingManagerAllColumns,
			organizationsBillingManagerColumnsWithDefault,
			organizationsBillingManagerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `organizations_billing_managers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `organizations_billing_managers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `organizations_billing_managers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, organizationsBillingManagerPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into organizations_billing_managers")
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

	o.OrganizationsBillingManagersPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsBillingManagerMapping["organizations_billing_managers_pk"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OrganizationsBillingManagersPK,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_billing_managers")
	}

CacheNoHooks:
	if !cached {
		organizationsBillingManagerInsertCacheMut.Lock()
		organizationsBillingManagerInsertCache[key] = cache
		organizationsBillingManagerInsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsBillingManagerUpdater interface {
	Update(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrganizationsBillingManager.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q organizationsBillingManagerQuery) Update(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	organizationsBillingManagerUpdateCacheMut.RLock()
	cache, cached := organizationsBillingManagerUpdateCache[key]
	organizationsBillingManagerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsBillingManagerAllColumns,
			organizationsBillingManagerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_billing_managers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `organizations_billing_managers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, organizationsBillingManagerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsBillingManagerType, organizationsBillingManagerMapping, append(wl, organizationsBillingManagerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_billing_managers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_billing_managers")
	}

	if !cached {
		organizationsBillingManagerUpdateCacheMut.Lock()
		organizationsBillingManagerUpdateCache[key] = cache
		organizationsBillingManagerUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsBillingManagerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_billing_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_billing_managers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q organizationsBillingManagerQuery) UpdateAllSlice(o OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsBillingManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `organizations_billing_managers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsBillingManagerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsBillingManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsBillingManager")
	}
	return rowsAff, nil
}

type OrganizationsBillingManagerDeleter interface {
	Delete(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrganizationsBillingManager record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q organizationsBillingManagerQuery) Delete(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsBillingManager provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsBillingManagerPrimaryKeyMapping)
	sql := "DELETE FROM `organizations_billing_managers` WHERE `organizations_billing_managers_pk`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_billing_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_billing_managers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsBillingManagerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsBillingManagerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_billing_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_billing_managers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q organizationsBillingManagerQuery) DeleteAllSlice(o OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsBillingManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `organizations_billing_managers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsBillingManagerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsBillingManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_billing_managers")
	}

	return rowsAff, nil
}

type OrganizationsBillingManagerReloader interface {
	Reload(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q organizationsBillingManagerQuery) Reload(o *OrganizationsBillingManager, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsBillingManager(ctx, exec, o.OrganizationsBillingManagersPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q organizationsBillingManagerQuery) ReloadAll(o *OrganizationsBillingManagerSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsBillingManagerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsBillingManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `organizations_billing_managers`.* FROM `organizations_billing_managers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsBillingManagerPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsBillingManagerSlice")
	}

	*o = slice

	return nil
}

// OrganizationsBillingManagerExists checks if the OrganizationsBillingManager row exists.
func OrganizationsBillingManagerExists(ctx context.Context, exec boil.ContextExecutor, organizationsBillingManagersPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `organizations_billing_managers` where `organizations_billing_managers_pk`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, organizationsBillingManagersPK)
	}
	row := exec.QueryRowContext(ctx, sql, organizationsBillingManagersPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_billing_managers exists")
	}

	return exists, nil
}