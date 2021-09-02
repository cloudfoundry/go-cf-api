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

// SpacesApplicationSupporter is an object representing the database table.
type SpacesApplicationSupporter struct {
	SpacesApplicationSupportersPK int         `boil:"spaces_application_supporters_pk" json:"spaces_application_supporters_pk" toml:"spaces_application_supporters_pk" yaml:"spaces_application_supporters_pk"`
	RoleGUID                      null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	SpaceID                       int         `boil:"space_id" json:"space_id" toml:"space_id" yaml:"space_id"`
	UserID                        int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt                     time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                     time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *spacesApplicationSupporterR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spacesApplicationSupporterL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpacesApplicationSupporterColumns = struct {
	SpacesApplicationSupportersPK string
	RoleGUID                      string
	SpaceID                       string
	UserID                        string
	CreatedAt                     string
	UpdatedAt                     string
}{
	SpacesApplicationSupportersPK: "spaces_application_supporters_pk",
	RoleGUID:                      "role_guid",
	SpaceID:                       "space_id",
	UserID:                        "user_id",
	CreatedAt:                     "created_at",
	UpdatedAt:                     "updated_at",
}

var SpacesApplicationSupporterTableColumns = struct {
	SpacesApplicationSupportersPK string
	RoleGUID                      string
	SpaceID                       string
	UserID                        string
	CreatedAt                     string
	UpdatedAt                     string
}{
	SpacesApplicationSupportersPK: "spaces_application_supporters.spaces_application_supporters_pk",
	RoleGUID:                      "spaces_application_supporters.role_guid",
	SpaceID:                       "spaces_application_supporters.space_id",
	UserID:                        "spaces_application_supporters.user_id",
	CreatedAt:                     "spaces_application_supporters.created_at",
	UpdatedAt:                     "spaces_application_supporters.updated_at",
}

// Generated where

var SpacesApplicationSupporterWhere = struct {
	SpacesApplicationSupportersPK whereHelperint
	RoleGUID                      whereHelpernull_String
	SpaceID                       whereHelperint
	UserID                        whereHelperint
	CreatedAt                     whereHelpertime_Time
	UpdatedAt                     whereHelpertime_Time
}{
	SpacesApplicationSupportersPK: whereHelperint{field: "\"spaces_application_supporters\".\"spaces_application_supporters_pk\""},
	RoleGUID:                      whereHelpernull_String{field: "\"spaces_application_supporters\".\"role_guid\""},
	SpaceID:                       whereHelperint{field: "\"spaces_application_supporters\".\"space_id\""},
	UserID:                        whereHelperint{field: "\"spaces_application_supporters\".\"user_id\""},
	CreatedAt:                     whereHelpertime_Time{field: "\"spaces_application_supporters\".\"created_at\""},
	UpdatedAt:                     whereHelpertime_Time{field: "\"spaces_application_supporters\".\"updated_at\""},
}

// SpacesApplicationSupporterRels is where relationship names are stored.
var SpacesApplicationSupporterRels = struct {
	Space string
	User  string
}{
	Space: "Space",
	User:  "User",
}

// spacesApplicationSupporterR is where relationships are stored.
type spacesApplicationSupporterR struct {
	Space *Space `boil:"Space" json:"Space" toml:"Space" yaml:"Space"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*spacesApplicationSupporterR) NewStruct() *spacesApplicationSupporterR {
	return &spacesApplicationSupporterR{}
}

// spacesApplicationSupporterL is where Load methods for each relationship are stored.
type spacesApplicationSupporterL struct{}

var (
	spacesApplicationSupporterAllColumns            = []string{"spaces_application_supporters_pk", "role_guid", "space_id", "user_id", "created_at", "updated_at"}
	spacesApplicationSupporterColumnsWithoutDefault = []string{"role_guid", "space_id", "user_id"}
	spacesApplicationSupporterColumnsWithDefault    = []string{"spaces_application_supporters_pk", "created_at", "updated_at"}
	spacesApplicationSupporterPrimaryKeyColumns     = []string{"spaces_application_supporters_pk"}
)

type (
	// SpacesApplicationSupporterSlice is an alias for a slice of pointers to SpacesApplicationSupporter.
	// This should almost always be used instead of []SpacesApplicationSupporter.
	SpacesApplicationSupporterSlice []*SpacesApplicationSupporter

	SpacesApplicationSupporterQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spacesApplicationSupporterType                 = reflect.TypeOf(&SpacesApplicationSupporter{})
	spacesApplicationSupporterMapping              = queries.MakeStructMapping(spacesApplicationSupporterType)
	spacesApplicationSupporterPrimaryKeyMapping, _ = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, spacesApplicationSupporterPrimaryKeyColumns)
	spacesApplicationSupporterInsertCacheMut       sync.RWMutex
	spacesApplicationSupporterInsertCache          = make(map[string]insertCache)
	spacesApplicationSupporterUpdateCacheMut       sync.RWMutex
	spacesApplicationSupporterUpdateCache          = make(map[string]updateCache)
	spacesApplicationSupporterUpsertCacheMut       sync.RWMutex
	spacesApplicationSupporterUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type SpacesApplicationSupporterFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*SpacesApplicationSupporter, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (SpacesApplicationSupporterSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single spacesApplicationSupporter record from the query.
func (q SpacesApplicationSupporterQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SpacesApplicationSupporter, error) {
	o := &SpacesApplicationSupporter{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for spaces_application_supporters")
	}

	return o, nil
}

// All returns all SpacesApplicationSupporter records from the query.
func (q SpacesApplicationSupporterQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpacesApplicationSupporterSlice, error) {
	var o []*SpacesApplicationSupporter

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SpacesApplicationSupporter slice")
	}

	return o, nil
}

// Count returns the count of all SpacesApplicationSupporter records in the query.
func (q SpacesApplicationSupporterQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count spaces_application_supporters rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q SpacesApplicationSupporterQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if spaces_application_supporters exists")
	}

	return count > 0, nil
}

// Space pointed to by the foreign key.
func (q SpacesApplicationSupporterQuery) Space(o *SpacesApplicationSupporter, mods ...qm.QueryMod) SpaceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SpaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Spaces(queryMods...)
	queries.SetFrom(query.Query, "\"spaces\"")

	return query
}

// User pointed to by the foreign key.
func (q SpacesApplicationSupporterQuery) User(o *SpacesApplicationSupporter, mods ...qm.QueryMod) UserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadSpace allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spacesApplicationSupporterL) LoadSpace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesApplicationSupporter interface{}, mods queries.Applicator) error {
	var slice []*SpacesApplicationSupporter
	var object *SpacesApplicationSupporter

	if singular {
		object = maybeSpacesApplicationSupporter.(*SpacesApplicationSupporter)
	} else {
		slice = *maybeSpacesApplicationSupporter.(*[]*SpacesApplicationSupporter)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesApplicationSupporterR{}
		}
		args = append(args, object.SpaceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesApplicationSupporterR{}
			}

			for _, a := range args {
				if a == obj.SpaceID {
					continue Outer
				}
			}

			args = append(args, obj.SpaceID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`spaces`),
		qm.WhereIn(`spaces.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Space")
	}

	var resultSlice []*Space
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Space")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for spaces")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for spaces")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Space = foreign
		if foreign.R == nil {
			foreign.R = &spaceR{}
		}
		foreign.R.SpacesApplicationSupporters = append(foreign.R.SpacesApplicationSupporters, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SpaceID == foreign.ID {
				local.R.Space = foreign
				if foreign.R == nil {
					foreign.R = &spaceR{}
				}
				foreign.R.SpacesApplicationSupporters = append(foreign.R.SpacesApplicationSupporters, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spacesApplicationSupporterL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesApplicationSupporter interface{}, mods queries.Applicator) error {
	var slice []*SpacesApplicationSupporter
	var object *SpacesApplicationSupporter

	if singular {
		object = maybeSpacesApplicationSupporter.(*SpacesApplicationSupporter)
	} else {
		slice = *maybeSpacesApplicationSupporter.(*[]*SpacesApplicationSupporter)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesApplicationSupporterR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesApplicationSupporterR{}
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
		foreign.R.SpacesApplicationSupporters = append(foreign.R.SpacesApplicationSupporters, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.SpacesApplicationSupporters = append(foreign.R.SpacesApplicationSupporters, local)
				break
			}
		}
	}

	return nil
}

// SetSpace of the spacesApplicationSupporter to the related item.
// Sets o.R.Space to related.
// Adds o to related.R.SpacesApplicationSupporters.
func (q SpacesApplicationSupporterQuery) SetSpace(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Space) error {
	var err error
	if insert {
		if err = Spaces().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_application_supporters\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"space_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesApplicationSupporterPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesApplicationSupportersPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SpaceID = related.ID
	if o.R == nil {
		o.R = &spacesApplicationSupporterR{
			Space: related,
		}
	} else {
		o.R.Space = related
	}

	if related.R == nil {
		related.R = &spaceR{
			SpacesApplicationSupporters: SpacesApplicationSupporterSlice{o},
		}
	} else {
		related.R.SpacesApplicationSupporters = append(related.R.SpacesApplicationSupporters, o)
	}

	return nil
}

// SetUser of the spacesApplicationSupporter to the related item.
// Sets o.R.User to related.
// Adds o to related.R.SpacesApplicationSupporters.
func (q SpacesApplicationSupporterQuery) SetUser(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = Users().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_application_supporters\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesApplicationSupporterPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesApplicationSupportersPK}

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
		o.R = &spacesApplicationSupporterR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			SpacesApplicationSupporters: SpacesApplicationSupporterSlice{o},
		}
	} else {
		related.R.SpacesApplicationSupporters = append(related.R.SpacesApplicationSupporters, o)
	}

	return nil
}

// SpacesApplicationSupporters retrieves all the records using an executor.
func SpacesApplicationSupporters(mods ...qm.QueryMod) SpacesApplicationSupporterQuery {
	mods = append(mods, qm.From("\"spaces_application_supporters\""))
	return SpacesApplicationSupporterQuery{NewQuery(mods...)}
}

type SpacesApplicationSupporterFinder interface {
	FindSpacesApplicationSupporter(ctx context.Context, exec boil.ContextExecutor, spacesApplicationSupportersPK int, selectCols ...string) (*SpacesApplicationSupporter, error)
}

// FindSpacesApplicationSupporter retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpacesApplicationSupporter(ctx context.Context, exec boil.ContextExecutor, spacesApplicationSupportersPK int, selectCols ...string) (*SpacesApplicationSupporter, error) {
	spacesApplicationSupporterObj := &SpacesApplicationSupporter{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"spaces_application_supporters\" where \"spaces_application_supporters_pk\"=$1", sel,
	)

	q := queries.Raw(query, spacesApplicationSupportersPK)

	err := q.Bind(ctx, exec, spacesApplicationSupporterObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from spaces_application_supporters")
	}

	return spacesApplicationSupporterObj, nil
}

type SpacesApplicationSupporterInserter interface {
	Insert(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q SpacesApplicationSupporterQuery) Insert(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_application_supporters provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(spacesApplicationSupporterColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spacesApplicationSupporterInsertCacheMut.RLock()
	cache, cached := spacesApplicationSupporterInsertCache[key]
	spacesApplicationSupporterInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spacesApplicationSupporterAllColumns,
			spacesApplicationSupporterColumnsWithDefault,
			spacesApplicationSupporterColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"spaces_application_supporters\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"spaces_application_supporters\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into spaces_application_supporters")
	}

	if !cached {
		spacesApplicationSupporterInsertCacheMut.Lock()
		spacesApplicationSupporterInsertCache[key] = cache
		spacesApplicationSupporterInsertCacheMut.Unlock()
	}

	return nil
}

type SpacesApplicationSupporterUpdater interface {
	Update(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the SpacesApplicationSupporter.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q SpacesApplicationSupporterQuery) Update(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	spacesApplicationSupporterUpdateCacheMut.RLock()
	cache, cached := spacesApplicationSupporterUpdateCache[key]
	spacesApplicationSupporterUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spacesApplicationSupporterAllColumns,
			spacesApplicationSupporterPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update spaces_application_supporters, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"spaces_application_supporters\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, spacesApplicationSupporterPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, append(wl, spacesApplicationSupporterPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update spaces_application_supporters row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for spaces_application_supporters")
	}

	if !cached {
		spacesApplicationSupporterUpdateCacheMut.Lock()
		spacesApplicationSupporterUpdateCache[key] = cache
		spacesApplicationSupporterUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q SpacesApplicationSupporterQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for spaces_application_supporters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for spaces_application_supporters")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q SpacesApplicationSupporterQuery) UpdateAllSlice(o SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesApplicationSupporterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"spaces_application_supporters\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, spacesApplicationSupporterPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in spacesApplicationSupporter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all spacesApplicationSupporter")
	}
	return rowsAff, nil
}

type SpacesApplicationSupporterDeleter interface {
	Delete(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single SpacesApplicationSupporter record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q SpacesApplicationSupporterQuery) Delete(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SpacesApplicationSupporter provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spacesApplicationSupporterPrimaryKeyMapping)
	sql := "DELETE FROM \"spaces_application_supporters\" WHERE \"spaces_application_supporters_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from spaces_application_supporters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for spaces_application_supporters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q SpacesApplicationSupporterQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no spacesApplicationSupporterQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spaces_application_supporters")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_application_supporters")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q SpacesApplicationSupporterQuery) DeleteAllSlice(o SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesApplicationSupporterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"spaces_application_supporters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesApplicationSupporterPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spacesApplicationSupporter slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_application_supporters")
	}

	return rowsAff, nil
}

type SpacesApplicationSupporterReloader interface {
	Reload(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q SpacesApplicationSupporterQuery) Reload(o *SpacesApplicationSupporter, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpacesApplicationSupporter(ctx, exec, o.SpacesApplicationSupportersPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q SpacesApplicationSupporterQuery) ReloadAll(o *SpacesApplicationSupporterSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpacesApplicationSupporterSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesApplicationSupporterPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"spaces_application_supporters\".* FROM \"spaces_application_supporters\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesApplicationSupporterPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpacesApplicationSupporterSlice")
	}

	*o = slice

	return nil
}

// SpacesApplicationSupporterExists checks if the SpacesApplicationSupporter row exists.
func SpacesApplicationSupporterExists(ctx context.Context, exec boil.ContextExecutor, spacesApplicationSupportersPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"spaces_application_supporters\" where \"spaces_application_supporters_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, spacesApplicationSupportersPK)
	}
	row := exec.QueryRowContext(ctx, sql, spacesApplicationSupportersPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if spaces_application_supporters exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SpacesApplicationSupporter) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_application_supporters provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(spacesApplicationSupporterColumnsWithDefault, o)

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

	spacesApplicationSupporterUpsertCacheMut.RLock()
	cache, cached := spacesApplicationSupporterUpsertCache[key]
	spacesApplicationSupporterUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spacesApplicationSupporterAllColumns,
			spacesApplicationSupporterColumnsWithDefault,
			spacesApplicationSupporterColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			spacesApplicationSupporterAllColumns,
			spacesApplicationSupporterPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert spaces_application_supporters, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(spacesApplicationSupporterPrimaryKeyColumns))
			copy(conflict, spacesApplicationSupporterPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"spaces_application_supporters\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spacesApplicationSupporterType, spacesApplicationSupporterMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert spaces_application_supporters")
	}

	if !cached {
		spacesApplicationSupporterUpsertCacheMut.Lock()
		spacesApplicationSupporterUpsertCache[key] = cache
		spacesApplicationSupporterUpsertCacheMut.Unlock()
	}

	return nil
}
