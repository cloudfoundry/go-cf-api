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

// OrganizationsUser is an object representing the database table.
type OrganizationsUser struct {
	OrganizationID       int         `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	UserID               int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	OrganizationsUsersPK int         `boil:"organizations_users_pk" json:"organizations_users_pk" toml:"organizations_users_pk" yaml:"organizations_users_pk"`
	RoleGUID             null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt            time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt            time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *organizationsUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsUserColumns = struct {
	OrganizationID       string
	UserID               string
	OrganizationsUsersPK string
	RoleGUID             string
	CreatedAt            string
	UpdatedAt            string
}{
	OrganizationID:       "organization_id",
	UserID:               "user_id",
	OrganizationsUsersPK: "organizations_users_pk",
	RoleGUID:             "role_guid",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

var OrganizationsUserTableColumns = struct {
	OrganizationID       string
	UserID               string
	OrganizationsUsersPK string
	RoleGUID             string
	CreatedAt            string
	UpdatedAt            string
}{
	OrganizationID:       "organizations_users.organization_id",
	UserID:               "organizations_users.user_id",
	OrganizationsUsersPK: "organizations_users.organizations_users_pk",
	RoleGUID:             "organizations_users.role_guid",
	CreatedAt:            "organizations_users.created_at",
	UpdatedAt:            "organizations_users.updated_at",
}

// Generated where

var OrganizationsUserWhere = struct {
	OrganizationID       whereHelperint
	UserID               whereHelperint
	OrganizationsUsersPK whereHelperint
	RoleGUID             whereHelpernull_String
	CreatedAt            whereHelpertime_Time
	UpdatedAt            whereHelpertime_Time
}{
	OrganizationID:       whereHelperint{field: "\"organizations_users\".\"organization_id\""},
	UserID:               whereHelperint{field: "\"organizations_users\".\"user_id\""},
	OrganizationsUsersPK: whereHelperint{field: "\"organizations_users\".\"organizations_users_pk\""},
	RoleGUID:             whereHelpernull_String{field: "\"organizations_users\".\"role_guid\""},
	CreatedAt:            whereHelpertime_Time{field: "\"organizations_users\".\"created_at\""},
	UpdatedAt:            whereHelpertime_Time{field: "\"organizations_users\".\"updated_at\""},
}

// OrganizationsUserRels is where relationship names are stored.
var OrganizationsUserRels = struct {
	User string
}{
	User: "User",
}

// organizationsUserR is where relationships are stored.
type organizationsUserR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*organizationsUserR) NewStruct() *organizationsUserR {
	return &organizationsUserR{}
}

// organizationsUserL is where Load methods for each relationship are stored.
type organizationsUserL struct{}

var (
	organizationsUserAllColumns            = []string{"organization_id", "user_id", "organizations_users_pk", "role_guid", "created_at", "updated_at"}
	organizationsUserColumnsWithoutDefault = []string{"organization_id", "user_id", "role_guid"}
	organizationsUserColumnsWithDefault    = []string{"organizations_users_pk", "created_at", "updated_at"}
	organizationsUserPrimaryKeyColumns     = []string{"organizations_users_pk"}
)

type (
	// OrganizationsUserSlice is an alias for a slice of pointers to OrganizationsUser.
	// This should almost always be used instead of []OrganizationsUser.
	OrganizationsUserSlice []*OrganizationsUser
	// OrganizationsUserHook is the signature for custom OrganizationsUser hook methods
	OrganizationsUserHook func(context.Context, boil.ContextExecutor, *OrganizationsUser) error

	organizationsUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsUserType                 = reflect.TypeOf(&OrganizationsUser{})
	organizationsUserMapping              = queries.MakeStructMapping(organizationsUserType)
	organizationsUserPrimaryKeyMapping, _ = queries.BindMapping(organizationsUserType, organizationsUserMapping, organizationsUserPrimaryKeyColumns)
	organizationsUserInsertCacheMut       sync.RWMutex
	organizationsUserInsertCache          = make(map[string]insertCache)
	organizationsUserUpdateCacheMut       sync.RWMutex
	organizationsUserUpdateCache          = make(map[string]updateCache)
	organizationsUserUpsertCacheMut       sync.RWMutex
	organizationsUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var organizationsUserBeforeInsertHooks []OrganizationsUserHook
var organizationsUserBeforeUpdateHooks []OrganizationsUserHook
var organizationsUserBeforeDeleteHooks []OrganizationsUserHook
var organizationsUserBeforeUpsertHooks []OrganizationsUserHook

var organizationsUserAfterInsertHooks []OrganizationsUserHook
var organizationsUserAfterSelectHooks []OrganizationsUserHook
var organizationsUserAfterUpdateHooks []OrganizationsUserHook
var organizationsUserAfterDeleteHooks []OrganizationsUserHook
var organizationsUserAfterUpsertHooks []OrganizationsUserHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrganizationsUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrganizationsUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrganizationsUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrganizationsUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrganizationsUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrganizationsUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrganizationsUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrganizationsUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrganizationsUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range organizationsUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrganizationsUserHook registers your hook function for all future operations.
func AddOrganizationsUserHook(hookPoint boil.HookPoint, organizationsUserHook OrganizationsUserHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		organizationsUserBeforeInsertHooks = append(organizationsUserBeforeInsertHooks, organizationsUserHook)
	case boil.BeforeUpdateHook:
		organizationsUserBeforeUpdateHooks = append(organizationsUserBeforeUpdateHooks, organizationsUserHook)
	case boil.BeforeDeleteHook:
		organizationsUserBeforeDeleteHooks = append(organizationsUserBeforeDeleteHooks, organizationsUserHook)
	case boil.BeforeUpsertHook:
		organizationsUserBeforeUpsertHooks = append(organizationsUserBeforeUpsertHooks, organizationsUserHook)
	case boil.AfterInsertHook:
		organizationsUserAfterInsertHooks = append(organizationsUserAfterInsertHooks, organizationsUserHook)
	case boil.AfterSelectHook:
		organizationsUserAfterSelectHooks = append(organizationsUserAfterSelectHooks, organizationsUserHook)
	case boil.AfterUpdateHook:
		organizationsUserAfterUpdateHooks = append(organizationsUserAfterUpdateHooks, organizationsUserHook)
	case boil.AfterDeleteHook:
		organizationsUserAfterDeleteHooks = append(organizationsUserAfterDeleteHooks, organizationsUserHook)
	case boil.AfterUpsertHook:
		organizationsUserAfterUpsertHooks = append(organizationsUserAfterUpsertHooks, organizationsUserHook)
	}
}

// One returns a single organizationsUser record from the query.
func (q organizationsUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsUser, error) {
	o := &OrganizationsUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrganizationsUser records from the query.
func (q organizationsUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsUserSlice, error) {
	var o []*OrganizationsUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsUser slice")
	}

	if len(organizationsUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrganizationsUser records in the query.
func (q organizationsUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_users exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *OrganizationsUser) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsUser interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsUser
	var object *OrganizationsUser

	if singular {
		object = maybeOrganizationsUser.(*OrganizationsUser)
	} else {
		slice = *maybeOrganizationsUser.(*[]*OrganizationsUser)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsUserR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsUserR{}
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

	if len(organizationsUserAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.OrganizationsUsers = append(foreign.R.OrganizationsUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OrganizationsUsers = append(foreign.R.OrganizationsUsers, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the organizationsUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OrganizationsUsers.
func (o *OrganizationsUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organizations_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, organizationsUserPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsUsersPK}

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
		o.R = &organizationsUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OrganizationsUsers: OrganizationsUserSlice{o},
		}
	} else {
		related.R.OrganizationsUsers = append(related.R.OrganizationsUsers, o)
	}

	return nil
}

// OrganizationsUsers retrieves all the records using an executor.
func OrganizationsUsers(mods ...qm.QueryMod) organizationsUserQuery {
	mods = append(mods, qm.From("\"organizations_users\""))
	return organizationsUserQuery{NewQuery(mods...)}
}

// FindOrganizationsUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsUser(ctx context.Context, exec boil.ContextExecutor, organizationsUsersPK int, selectCols ...string) (*OrganizationsUser, error) {
	organizationsUserObj := &OrganizationsUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organizations_users\" where \"organizations_users_pk\"=$1", sel,
	)

	q := queries.Raw(query, organizationsUsersPK)

	err := q.Bind(ctx, exec, organizationsUserObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_users")
	}

	if err = organizationsUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return organizationsUserObj, err
	}

	return organizationsUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrganizationsUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_users provided for insertion")
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

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsUserInsertCacheMut.RLock()
	cache, cached := organizationsUserInsertCache[key]
	organizationsUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsUserAllColumns,
			organizationsUserColumnsWithDefault,
			organizationsUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsUserType, organizationsUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsUserType, organizationsUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"organizations_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"organizations_users\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into organizations_users")
	}

	if !cached {
		organizationsUserInsertCacheMut.Lock()
		organizationsUserInsertCache[key] = cache
		organizationsUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrganizationsUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrganizationsUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	organizationsUserUpdateCacheMut.RLock()
	cache, cached := organizationsUserUpdateCache[key]
	organizationsUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsUserAllColumns,
			organizationsUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organizations_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organizationsUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsUserType, organizationsUserMapping, append(wl, organizationsUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_users")
	}

	if !cached {
		organizationsUserUpdateCacheMut.Lock()
		organizationsUserUpdateCache[key] = cache
		organizationsUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrganizationsUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"organizations_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, organizationsUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrganizationsUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_users provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsUserColumnsWithDefault, o)

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

	organizationsUserUpsertCacheMut.RLock()
	cache, cached := organizationsUserUpsertCache[key]
	organizationsUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsUserAllColumns,
			organizationsUserColumnsWithDefault,
			organizationsUserColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationsUserAllColumns,
			organizationsUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organizationsUserPrimaryKeyColumns))
			copy(conflict, organizationsUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"organizations_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(organizationsUserType, organizationsUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsUserType, organizationsUserMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert organizations_users")
	}

	if !cached {
		organizationsUserUpsertCacheMut.Lock()
		organizationsUserUpsertCache[key] = cache
		organizationsUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrganizationsUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrganizationsUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsUserPrimaryKeyMapping)
	sql := "DELETE FROM \"organizations_users\" WHERE \"organizations_users_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrganizationsUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(organizationsUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"organizations_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, organizationsUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_users")
	}

	if len(organizationsUserAfterDeleteHooks) != 0 {
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
func (o *OrganizationsUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsUser(ctx, exec, o.OrganizationsUsersPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrganizationsUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"organizations_users\".* FROM \"organizations_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, organizationsUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsUserSlice")
	}

	*o = slice

	return nil
}

// OrganizationsUserExists checks if the OrganizationsUser row exists.
func OrganizationsUserExists(ctx context.Context, exec boil.ContextExecutor, organizationsUsersPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"organizations_users\" where \"organizations_users_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, organizationsUsersPK)
	}
	row := exec.QueryRowContext(ctx, sql, organizationsUsersPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_users exists")
	}

	return exists, nil
}
