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

type OrganizationsAuditorUpserter interface {
	Upsert(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLOrganizationsAuditorUniqueColumns = []string{
	"organizations_auditors_pk",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q organizationsAuditorQuery) Upsert(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_auditors provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsAuditorColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOrganizationsAuditorUniqueColumns, o)

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

	organizationsAuditorUpsertCacheMut.RLock()
	cache, cached := organizationsAuditorUpsertCache[key]
	organizationsAuditorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsAuditorAllColumns,
			organizationsAuditorColumnsWithDefault,
			organizationsAuditorColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationsAuditorAllColumns,
			organizationsAuditorPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_auditors, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`organizations_auditors`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `organizations_auditors` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for organizations_auditors")
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

	o.OrganizationsAuditorsPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsAuditorMapping["organizations_auditors_pk"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for organizations_auditors")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_auditors")
	}

CacheNoHooks:
	if !cached {
		organizationsAuditorUpsertCacheMut.Lock()
		organizationsAuditorUpsertCache[key] = cache
		organizationsAuditorUpsertCacheMut.Unlock()
	}

	return nil
}

// OrganizationsAuditor is an object representing the database table.
type OrganizationsAuditor struct {
	OrganizationID          int         `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	UserID                  int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	OrganizationsAuditorsPK int         `boil:"organizations_auditors_pk" json:"organizations_auditors_pk" toml:"organizations_auditors_pk" yaml:"organizations_auditors_pk"`
	RoleGUID                null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt               time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt               time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *organizationsAuditorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsAuditorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsAuditorColumns = struct {
	OrganizationID          string
	UserID                  string
	OrganizationsAuditorsPK string
	RoleGUID                string
	CreatedAt               string
	UpdatedAt               string
}{
	OrganizationID:          "organization_id",
	UserID:                  "user_id",
	OrganizationsAuditorsPK: "organizations_auditors_pk",
	RoleGUID:                "role_guid",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
}

var OrganizationsAuditorTableColumns = struct {
	OrganizationID          string
	UserID                  string
	OrganizationsAuditorsPK string
	RoleGUID                string
	CreatedAt               string
	UpdatedAt               string
}{
	OrganizationID:          "organizations_auditors.organization_id",
	UserID:                  "organizations_auditors.user_id",
	OrganizationsAuditorsPK: "organizations_auditors.organizations_auditors_pk",
	RoleGUID:                "organizations_auditors.role_guid",
	CreatedAt:               "organizations_auditors.created_at",
	UpdatedAt:               "organizations_auditors.updated_at",
}

// Generated where

var OrganizationsAuditorWhere = struct {
	OrganizationID          whereHelperint
	UserID                  whereHelperint
	OrganizationsAuditorsPK whereHelperint
	RoleGUID                whereHelpernull_String
	CreatedAt               whereHelpertime_Time
	UpdatedAt               whereHelpertime_Time
}{
	OrganizationID:          whereHelperint{field: "`organizations_auditors`.`organization_id`"},
	UserID:                  whereHelperint{field: "`organizations_auditors`.`user_id`"},
	OrganizationsAuditorsPK: whereHelperint{field: "`organizations_auditors`.`organizations_auditors_pk`"},
	RoleGUID:                whereHelpernull_String{field: "`organizations_auditors`.`role_guid`"},
	CreatedAt:               whereHelpertime_Time{field: "`organizations_auditors`.`created_at`"},
	UpdatedAt:               whereHelpertime_Time{field: "`organizations_auditors`.`updated_at`"},
}

// OrganizationsAuditorRels is where relationship names are stored.
var OrganizationsAuditorRels = struct {
	Organization string
	User         string
}{
	Organization: "Organization",
	User:         "User",
}

// organizationsAuditorR is where relationships are stored.
type organizationsAuditorR struct {
	Organization *Organization `boil:"Organization" json:"Organization" toml:"Organization" yaml:"Organization"`
	User         *User         `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*organizationsAuditorR) NewStruct() *organizationsAuditorR {
	return &organizationsAuditorR{}
}

// organizationsAuditorL is where Load methods for each relationship are stored.
type organizationsAuditorL struct{}

var (
	organizationsAuditorAllColumns            = []string{"organization_id", "user_id", "organizations_auditors_pk", "role_guid", "created_at", "updated_at"}
	organizationsAuditorColumnsWithoutDefault = []string{"organization_id", "user_id", "role_guid"}
	organizationsAuditorColumnsWithDefault    = []string{"organizations_auditors_pk", "created_at", "updated_at"}
	organizationsAuditorPrimaryKeyColumns     = []string{"organizations_auditors_pk"}
)

type (
	// OrganizationsAuditorSlice is an alias for a slice of pointers to OrganizationsAuditor.
	// This should almost always be used instead of []OrganizationsAuditor.
	OrganizationsAuditorSlice []*OrganizationsAuditor

	organizationsAuditorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsAuditorType                 = reflect.TypeOf(&OrganizationsAuditor{})
	organizationsAuditorMapping              = queries.MakeStructMapping(organizationsAuditorType)
	organizationsAuditorPrimaryKeyMapping, _ = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, organizationsAuditorPrimaryKeyColumns)
	organizationsAuditorInsertCacheMut       sync.RWMutex
	organizationsAuditorInsertCache          = make(map[string]insertCache)
	organizationsAuditorUpdateCacheMut       sync.RWMutex
	organizationsAuditorUpdateCache          = make(map[string]updateCache)
	organizationsAuditorUpsertCacheMut       sync.RWMutex
	organizationsAuditorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrganizationsAuditorFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsAuditor, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsAuditorSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single organizationsAuditor record from the query.
func (q organizationsAuditorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsAuditor, error) {
	o := &OrganizationsAuditor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_auditors")
	}

	return o, nil
}

// All returns all OrganizationsAuditor records from the query.
func (q organizationsAuditorQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsAuditorSlice, error) {
	var o []*OrganizationsAuditor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsAuditor slice")
	}

	return o, nil
}

// Count returns the count of all OrganizationsAuditor records in the query.
func (q organizationsAuditorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_auditors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsAuditorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_auditors exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (q organizationsAuditorQuery) Organization(o *OrganizationsAuditor, mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "`organizations`")

	return query
}

// User pointed to by the foreign key.
func (q organizationsAuditorQuery) User(o *OrganizationsAuditor, mods ...qm.QueryMod) userQuery {
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
func (organizationsAuditorL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsAuditor interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsAuditor
	var object *OrganizationsAuditor

	if singular {
		object = maybeOrganizationsAuditor.(*OrganizationsAuditor)
	} else {
		slice = *maybeOrganizationsAuditor.(*[]*OrganizationsAuditor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsAuditorR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsAuditorR{}
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
		foreign.R.OrganizationsAuditors = append(foreign.R.OrganizationsAuditors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.OrganizationsAuditors = append(foreign.R.OrganizationsAuditors, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsAuditorL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsAuditor interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsAuditor
	var object *OrganizationsAuditor

	if singular {
		object = maybeOrganizationsAuditor.(*OrganizationsAuditor)
	} else {
		slice = *maybeOrganizationsAuditor.(*[]*OrganizationsAuditor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsAuditorR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsAuditorR{}
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
		foreign.R.OrganizationsAuditors = append(foreign.R.OrganizationsAuditors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.OrganizationsAuditors = append(foreign.R.OrganizationsAuditors, local)
				break
			}
		}
	}

	return nil
}

// SetOrganization of the organizationsAuditor to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.OrganizationsAuditors.
func (q organizationsAuditorQuery) SetOrganization(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = Organizations().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_auditors` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"organization_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsAuditorPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsAuditorsPK}

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
		o.R = &organizationsAuditorR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			OrganizationsAuditors: OrganizationsAuditorSlice{o},
		}
	} else {
		related.R.OrganizationsAuditors = append(related.R.OrganizationsAuditors, o)
	}

	return nil
}

// SetUser of the organizationsAuditor to the related item.
// Sets o.R.User to related.
// Adds o to related.R.OrganizationsAuditors.
func (q organizationsAuditorQuery) SetUser(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = Users().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `organizations_auditors` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, organizationsAuditorPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsAuditorsPK}

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
		o.R = &organizationsAuditorR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			OrganizationsAuditors: OrganizationsAuditorSlice{o},
		}
	} else {
		related.R.OrganizationsAuditors = append(related.R.OrganizationsAuditors, o)
	}

	return nil
}

// OrganizationsAuditors retrieves all the records using an executor.
func OrganizationsAuditors(mods ...qm.QueryMod) organizationsAuditorQuery {
	mods = append(mods, qm.From("`organizations_auditors`"))
	return organizationsAuditorQuery{NewQuery(mods...)}
}

type OrganizationsAuditorFinder interface {
	FindOrganizationsAuditor(ctx context.Context, exec boil.ContextExecutor, organizationsAuditorsPK int, selectCols ...string) (*OrganizationsAuditor, error)
}

// FindOrganizationsAuditor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsAuditor(ctx context.Context, exec boil.ContextExecutor, organizationsAuditorsPK int, selectCols ...string) (*OrganizationsAuditor, error) {
	organizationsAuditorObj := &OrganizationsAuditor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `organizations_auditors` where `organizations_auditors_pk`=?", sel,
	)

	q := queries.Raw(query, organizationsAuditorsPK)

	err := q.Bind(ctx, exec, organizationsAuditorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_auditors")
	}

	return organizationsAuditorObj, nil
}

type OrganizationsAuditorInserter interface {
	Insert(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q organizationsAuditorQuery) Insert(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_auditors provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(organizationsAuditorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsAuditorInsertCacheMut.RLock()
	cache, cached := organizationsAuditorInsertCache[key]
	organizationsAuditorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsAuditorAllColumns,
			organizationsAuditorColumnsWithDefault,
			organizationsAuditorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `organizations_auditors` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `organizations_auditors` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `organizations_auditors` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, organizationsAuditorPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into organizations_auditors")
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

	o.OrganizationsAuditorsPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == organizationsAuditorMapping["organizations_auditors_pk"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.OrganizationsAuditorsPK,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for organizations_auditors")
	}

CacheNoHooks:
	if !cached {
		organizationsAuditorInsertCacheMut.Lock()
		organizationsAuditorInsertCache[key] = cache
		organizationsAuditorInsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsAuditorUpdater interface {
	Update(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrganizationsAuditor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q organizationsAuditorQuery) Update(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	organizationsAuditorUpdateCacheMut.RLock()
	cache, cached := organizationsAuditorUpdateCache[key]
	organizationsAuditorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsAuditorAllColumns,
			organizationsAuditorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_auditors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `organizations_auditors` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, organizationsAuditorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsAuditorType, organizationsAuditorMapping, append(wl, organizationsAuditorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_auditors row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_auditors")
	}

	if !cached {
		organizationsAuditorUpdateCacheMut.Lock()
		organizationsAuditorUpdateCache[key] = cache
		organizationsAuditorUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsAuditorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_auditors")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q organizationsAuditorQuery) UpdateAllSlice(o OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `organizations_auditors` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsAuditorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsAuditor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsAuditor")
	}
	return rowsAff, nil
}

type OrganizationsAuditorDeleter interface {
	Delete(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrganizationsAuditor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q organizationsAuditorQuery) Delete(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsAuditor provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsAuditorPrimaryKeyMapping)
	sql := "DELETE FROM `organizations_auditors` WHERE `organizations_auditors_pk`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_auditors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsAuditorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsAuditorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_auditors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q organizationsAuditorQuery) DeleteAllSlice(o OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `organizations_auditors` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsAuditorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsAuditor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_auditors")
	}

	return rowsAff, nil
}

type OrganizationsAuditorReloader interface {
	Reload(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q organizationsAuditorQuery) Reload(o *OrganizationsAuditor, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsAuditor(ctx, exec, o.OrganizationsAuditorsPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q organizationsAuditorQuery) ReloadAll(o *OrganizationsAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsAuditorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `organizations_auditors`.* FROM `organizations_auditors` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, organizationsAuditorPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsAuditorSlice")
	}

	*o = slice

	return nil
}

// OrganizationsAuditorExists checks if the OrganizationsAuditor row exists.
func OrganizationsAuditorExists(ctx context.Context, exec boil.ContextExecutor, organizationsAuditorsPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `organizations_auditors` where `organizations_auditors_pk`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, organizationsAuditorsPK)
	}
	row := exec.QueryRowContext(ctx, sql, organizationsAuditorsPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_auditors exists")
	}

	return exists, nil
}
