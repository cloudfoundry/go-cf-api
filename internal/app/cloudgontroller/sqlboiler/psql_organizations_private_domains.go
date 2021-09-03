// +build psql
//go:generate mockgen -source=$GOFILE -destination=mocks/organizations_private_domains.go
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OrganizationsPrivateDomain is an object representing the database table.
type OrganizationsPrivateDomain struct {
	OrganizationID                int `boil:"organization_id" json:"organization_id" toml:"organization_id" yaml:"organization_id"`
	PrivateDomainID               int `boil:"private_domain_id" json:"private_domain_id" toml:"private_domain_id" yaml:"private_domain_id"`
	OrganizationsPrivateDomainsPK int `boil:"organizations_private_domains_pk" json:"organizations_private_domains_pk" toml:"organizations_private_domains_pk" yaml:"organizations_private_domains_pk"`

	R *organizationsPrivateDomainR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L organizationsPrivateDomainL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrganizationsPrivateDomainColumns = struct {
	OrganizationID                string
	PrivateDomainID               string
	OrganizationsPrivateDomainsPK string
}{
	OrganizationID:                "organization_id",
	PrivateDomainID:               "private_domain_id",
	OrganizationsPrivateDomainsPK: "organizations_private_domains_pk",
}

var OrganizationsPrivateDomainTableColumns = struct {
	OrganizationID                string
	PrivateDomainID               string
	OrganizationsPrivateDomainsPK string
}{
	OrganizationID:                "organizations_private_domains.organization_id",
	PrivateDomainID:               "organizations_private_domains.private_domain_id",
	OrganizationsPrivateDomainsPK: "organizations_private_domains.organizations_private_domains_pk",
}

// Generated where

var OrganizationsPrivateDomainWhere = struct {
	OrganizationID                whereHelperint
	PrivateDomainID               whereHelperint
	OrganizationsPrivateDomainsPK whereHelperint
}{
	OrganizationID:                whereHelperint{field: "\"organizations_private_domains\".\"organization_id\""},
	PrivateDomainID:               whereHelperint{field: "\"organizations_private_domains\".\"private_domain_id\""},
	OrganizationsPrivateDomainsPK: whereHelperint{field: "\"organizations_private_domains\".\"organizations_private_domains_pk\""},
}

// OrganizationsPrivateDomainRels is where relationship names are stored.
var OrganizationsPrivateDomainRels = struct {
	Organization  string
	PrivateDomain string
}{
	Organization:  "Organization",
	PrivateDomain: "PrivateDomain",
}

// organizationsPrivateDomainR is where relationships are stored.
type organizationsPrivateDomainR struct {
	Organization  *Organization `boil:"Organization" json:"Organization" toml:"Organization" yaml:"Organization"`
	PrivateDomain *Domain       `boil:"PrivateDomain" json:"PrivateDomain" toml:"PrivateDomain" yaml:"PrivateDomain"`
}

// NewStruct creates a new relationship struct
func (*organizationsPrivateDomainR) NewStruct() *organizationsPrivateDomainR {
	return &organizationsPrivateDomainR{}
}

// organizationsPrivateDomainL is where Load methods for each relationship are stored.
type organizationsPrivateDomainL struct{}

var (
	organizationsPrivateDomainAllColumns            = []string{"organization_id", "private_domain_id", "organizations_private_domains_pk"}
	organizationsPrivateDomainColumnsWithoutDefault = []string{"organization_id", "private_domain_id"}
	organizationsPrivateDomainColumnsWithDefault    = []string{"organizations_private_domains_pk"}
	organizationsPrivateDomainPrimaryKeyColumns     = []string{"organizations_private_domains_pk"}
)

type (
	// OrganizationsPrivateDomainSlice is an alias for a slice of pointers to OrganizationsPrivateDomain.
	// This should almost always be used instead of []OrganizationsPrivateDomain.
	OrganizationsPrivateDomainSlice []*OrganizationsPrivateDomain

	organizationsPrivateDomainQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	organizationsPrivateDomainType                 = reflect.TypeOf(&OrganizationsPrivateDomain{})
	organizationsPrivateDomainMapping              = queries.MakeStructMapping(organizationsPrivateDomainType)
	organizationsPrivateDomainPrimaryKeyMapping, _ = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, organizationsPrivateDomainPrimaryKeyColumns)
	organizationsPrivateDomainInsertCacheMut       sync.RWMutex
	organizationsPrivateDomainInsertCache          = make(map[string]insertCache)
	organizationsPrivateDomainUpdateCacheMut       sync.RWMutex
	organizationsPrivateDomainUpdateCache          = make(map[string]updateCache)
	organizationsPrivateDomainUpsertCacheMut       sync.RWMutex
	organizationsPrivateDomainUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type OrganizationsPrivateDomainFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsPrivateDomain, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsPrivateDomainSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single organizationsPrivateDomain record from the query.
func (q organizationsPrivateDomainQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrganizationsPrivateDomain, error) {
	o := &OrganizationsPrivateDomain{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for organizations_private_domains")
	}

	return o, nil
}

// All returns all OrganizationsPrivateDomain records from the query.
func (q organizationsPrivateDomainQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrganizationsPrivateDomainSlice, error) {
	var o []*OrganizationsPrivateDomain

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrganizationsPrivateDomain slice")
	}

	return o, nil
}

// Count returns the count of all OrganizationsPrivateDomain records in the query.
func (q organizationsPrivateDomainQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count organizations_private_domains rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q organizationsPrivateDomainQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if organizations_private_domains exists")
	}

	return count > 0, nil
}

// Organization pointed to by the foreign key.
func (q organizationsPrivateDomainQuery) Organization(o *OrganizationsPrivateDomain, mods ...qm.QueryMod) organizationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrganizationID),
	}

	queryMods = append(queryMods, mods...)

	query := Organizations(queryMods...)
	queries.SetFrom(query.Query, "\"organizations\"")

	return query
}

// PrivateDomain pointed to by the foreign key.
func (q organizationsPrivateDomainQuery) PrivateDomain(o *OrganizationsPrivateDomain, mods ...qm.QueryMod) domainQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PrivateDomainID),
	}

	queryMods = append(queryMods, mods...)

	query := Domains(queryMods...)
	queries.SetFrom(query.Query, "\"domains\"")

	return query
}

// LoadOrganization allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsPrivateDomainL) LoadOrganization(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsPrivateDomain interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsPrivateDomain
	var object *OrganizationsPrivateDomain

	if singular {
		object = maybeOrganizationsPrivateDomain.(*OrganizationsPrivateDomain)
	} else {
		slice = *maybeOrganizationsPrivateDomain.(*[]*OrganizationsPrivateDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsPrivateDomainR{}
		}
		args = append(args, object.OrganizationID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsPrivateDomainR{}
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
		foreign.R.OrganizationsPrivateDomains = append(foreign.R.OrganizationsPrivateDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrganizationID == foreign.ID {
				local.R.Organization = foreign
				if foreign.R == nil {
					foreign.R = &organizationR{}
				}
				foreign.R.OrganizationsPrivateDomains = append(foreign.R.OrganizationsPrivateDomains, local)
				break
			}
		}
	}

	return nil
}

// LoadPrivateDomain allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (organizationsPrivateDomainL) LoadPrivateDomain(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrganizationsPrivateDomain interface{}, mods queries.Applicator) error {
	var slice []*OrganizationsPrivateDomain
	var object *OrganizationsPrivateDomain

	if singular {
		object = maybeOrganizationsPrivateDomain.(*OrganizationsPrivateDomain)
	} else {
		slice = *maybeOrganizationsPrivateDomain.(*[]*OrganizationsPrivateDomain)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &organizationsPrivateDomainR{}
		}
		args = append(args, object.PrivateDomainID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &organizationsPrivateDomainR{}
			}

			for _, a := range args {
				if a == obj.PrivateDomainID {
					continue Outer
				}
			}

			args = append(args, obj.PrivateDomainID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`domains`),
		qm.WhereIn(`domains.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Domain")
	}

	var resultSlice []*Domain
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Domain")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for domains")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for domains")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.PrivateDomain = foreign
		if foreign.R == nil {
			foreign.R = &domainR{}
		}
		foreign.R.PrivateDomainOrganizationsPrivateDomains = append(foreign.R.PrivateDomainOrganizationsPrivateDomains, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PrivateDomainID == foreign.ID {
				local.R.PrivateDomain = foreign
				if foreign.R == nil {
					foreign.R = &domainR{}
				}
				foreign.R.PrivateDomainOrganizationsPrivateDomains = append(foreign.R.PrivateDomainOrganizationsPrivateDomains, local)
				break
			}
		}
	}

	return nil
}

// SetOrganization of the organizationsPrivateDomain to the related item.
// Sets o.R.Organization to related.
// Adds o to related.R.OrganizationsPrivateDomains.
func (q organizationsPrivateDomainQuery) SetOrganization(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Organization) error {
	var err error
	if insert {
		if err = Organizations().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organizations_private_domains\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"organization_id"}),
		strmangle.WhereClause("\"", "\"", 2, organizationsPrivateDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsPrivateDomainsPK}

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
		o.R = &organizationsPrivateDomainR{
			Organization: related,
		}
	} else {
		o.R.Organization = related
	}

	if related.R == nil {
		related.R = &organizationR{
			OrganizationsPrivateDomains: OrganizationsPrivateDomainSlice{o},
		}
	} else {
		related.R.OrganizationsPrivateDomains = append(related.R.OrganizationsPrivateDomains, o)
	}

	return nil
}

// SetPrivateDomain of the organizationsPrivateDomain to the related item.
// Sets o.R.PrivateDomain to related.
// Adds o to related.R.PrivateDomainOrganizationsPrivateDomains.
func (q organizationsPrivateDomainQuery) SetPrivateDomain(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Domain) error {
	var err error
	if insert {
		if err = Domains().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"organizations_private_domains\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"private_domain_id"}),
		strmangle.WhereClause("\"", "\"", 2, organizationsPrivateDomainPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.OrganizationsPrivateDomainsPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PrivateDomainID = related.ID
	if o.R == nil {
		o.R = &organizationsPrivateDomainR{
			PrivateDomain: related,
		}
	} else {
		o.R.PrivateDomain = related
	}

	if related.R == nil {
		related.R = &domainR{
			PrivateDomainOrganizationsPrivateDomains: OrganizationsPrivateDomainSlice{o},
		}
	} else {
		related.R.PrivateDomainOrganizationsPrivateDomains = append(related.R.PrivateDomainOrganizationsPrivateDomains, o)
	}

	return nil
}

// OrganizationsPrivateDomains retrieves all the records using an executor.
func OrganizationsPrivateDomains(mods ...qm.QueryMod) organizationsPrivateDomainQuery {
	mods = append(mods, qm.From("\"organizations_private_domains\""))
	return organizationsPrivateDomainQuery{NewQuery(mods...)}
}

type OrganizationsPrivateDomainFinder interface {
	FindOrganizationsPrivateDomain(ctx context.Context, exec boil.ContextExecutor, organizationsPrivateDomainsPK int, selectCols ...string) (*OrganizationsPrivateDomain, error)
}

// FindOrganizationsPrivateDomain retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrganizationsPrivateDomain(ctx context.Context, exec boil.ContextExecutor, organizationsPrivateDomainsPK int, selectCols ...string) (*OrganizationsPrivateDomain, error) {
	organizationsPrivateDomainObj := &OrganizationsPrivateDomain{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"organizations_private_domains\" where \"organizations_private_domains_pk\"=$1", sel,
	)

	q := queries.Raw(query, organizationsPrivateDomainsPK)

	err := q.Bind(ctx, exec, organizationsPrivateDomainObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from organizations_private_domains")
	}

	return organizationsPrivateDomainObj, nil
}

type OrganizationsPrivateDomainInserter interface {
	Insert(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q organizationsPrivateDomainQuery) Insert(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_private_domains provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(organizationsPrivateDomainColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	organizationsPrivateDomainInsertCacheMut.RLock()
	cache, cached := organizationsPrivateDomainInsertCache[key]
	organizationsPrivateDomainInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			organizationsPrivateDomainAllColumns,
			organizationsPrivateDomainColumnsWithDefault,
			organizationsPrivateDomainColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"organizations_private_domains\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"organizations_private_domains\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into organizations_private_domains")
	}

	if !cached {
		organizationsPrivateDomainInsertCacheMut.Lock()
		organizationsPrivateDomainInsertCache[key] = cache
		organizationsPrivateDomainInsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsPrivateDomainUpdater interface {
	Update(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the OrganizationsPrivateDomain.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q organizationsPrivateDomainQuery) Update(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	organizationsPrivateDomainUpdateCacheMut.RLock()
	cache, cached := organizationsPrivateDomainUpdateCache[key]
	organizationsPrivateDomainUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			organizationsPrivateDomainAllColumns,
			organizationsPrivateDomainPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update organizations_private_domains, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"organizations_private_domains\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, organizationsPrivateDomainPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, append(wl, organizationsPrivateDomainPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update organizations_private_domains row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for organizations_private_domains")
	}

	if !cached {
		organizationsPrivateDomainUpdateCacheMut.Lock()
		organizationsPrivateDomainUpdateCache[key] = cache
		organizationsPrivateDomainUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q organizationsPrivateDomainQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for organizations_private_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for organizations_private_domains")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q organizationsPrivateDomainQuery) UpdateAllSlice(o OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsPrivateDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"organizations_private_domains\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, organizationsPrivateDomainPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in organizationsPrivateDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all organizationsPrivateDomain")
	}
	return rowsAff, nil
}

type OrganizationsPrivateDomainUpserter interface {
	Upsert(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q organizationsPrivateDomainQuery) Upsert(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no organizations_private_domains provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(organizationsPrivateDomainColumnsWithDefault, o)

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

	organizationsPrivateDomainUpsertCacheMut.RLock()
	cache, cached := organizationsPrivateDomainUpsertCache[key]
	organizationsPrivateDomainUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			organizationsPrivateDomainAllColumns,
			organizationsPrivateDomainColumnsWithDefault,
			organizationsPrivateDomainColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			organizationsPrivateDomainAllColumns,
			organizationsPrivateDomainPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert organizations_private_domains, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(organizationsPrivateDomainPrimaryKeyColumns))
			copy(conflict, organizationsPrivateDomainPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"organizations_private_domains\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(organizationsPrivateDomainType, organizationsPrivateDomainMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert organizations_private_domains")
	}

	if !cached {
		organizationsPrivateDomainUpsertCacheMut.Lock()
		organizationsPrivateDomainUpsertCache[key] = cache
		organizationsPrivateDomainUpsertCacheMut.Unlock()
	}

	return nil
}

type OrganizationsPrivateDomainDeleter interface {
	Delete(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single OrganizationsPrivateDomain record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q organizationsPrivateDomainQuery) Delete(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrganizationsPrivateDomain provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), organizationsPrivateDomainPrimaryKeyMapping)
	sql := "DELETE FROM \"organizations_private_domains\" WHERE \"organizations_private_domains_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from organizations_private_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for organizations_private_domains")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q organizationsPrivateDomainQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no organizationsPrivateDomainQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizations_private_domains")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_private_domains")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q organizationsPrivateDomainQuery) DeleteAllSlice(o OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsPrivateDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"organizations_private_domains\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, organizationsPrivateDomainPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from organizationsPrivateDomain slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for organizations_private_domains")
	}

	return rowsAff, nil
}

type OrganizationsPrivateDomainReloader interface {
	Reload(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q organizationsPrivateDomainQuery) Reload(o *OrganizationsPrivateDomain, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrganizationsPrivateDomain(ctx, exec, o.OrganizationsPrivateDomainsPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q organizationsPrivateDomainQuery) ReloadAll(o *OrganizationsPrivateDomainSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrganizationsPrivateDomainSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), organizationsPrivateDomainPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"organizations_private_domains\".* FROM \"organizations_private_domains\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, organizationsPrivateDomainPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrganizationsPrivateDomainSlice")
	}

	*o = slice

	return nil
}

// OrganizationsPrivateDomainExists checks if the OrganizationsPrivateDomain row exists.
func OrganizationsPrivateDomainExists(ctx context.Context, exec boil.ContextExecutor, organizationsPrivateDomainsPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"organizations_private_domains\" where \"organizations_private_domains_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, organizationsPrivateDomainsPK)
	}
	row := exec.QueryRowContext(ctx, sql, organizationsPrivateDomainsPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if organizations_private_domains exists")
	}

	return exists, nil
}
