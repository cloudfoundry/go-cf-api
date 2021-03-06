//go:build psql
// +build psql

//go:generate sh -c "echo '\x2bbuild unit' > ../../../../buildtags.txt && mockgen -source=$GOFILE -destination=mocks/spaces_auditors.go -copyright_file=../../../../buildtags.txt && rm -f ../../../../buildtags.txt"
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

type SpacesAuditorUpserter interface {
	Upsert(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q spacesAuditorQuery) Upsert(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_auditors provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(spacesAuditorColumnsWithDefault, o)

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

	spacesAuditorUpsertCacheMut.RLock()
	cache, cached := spacesAuditorUpsertCache[key]
	spacesAuditorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spacesAuditorAllColumns,
			spacesAuditorColumnsWithDefault,
			spacesAuditorColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			spacesAuditorAllColumns,
			spacesAuditorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert spaces_auditors, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(spacesAuditorPrimaryKeyColumns))
			copy(conflict, spacesAuditorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"spaces_auditors\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert spaces_auditors")
	}

	if !cached {
		spacesAuditorUpsertCacheMut.Lock()
		spacesAuditorUpsertCache[key] = cache
		spacesAuditorUpsertCacheMut.Unlock()
	}

	return nil
}

// SpacesAuditor is an object representing the database table.
type SpacesAuditor struct {
	SpaceID          int         `boil:"space_id" json:"space_id" toml:"space_id" yaml:"space_id"`
	UserID           int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	SpacesAuditorsPK int         `boil:"spaces_auditors_pk" json:"spaces_auditors_pk" toml:"spaces_auditors_pk" yaml:"spaces_auditors_pk"`
	RoleGUID         null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt        time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *spacesAuditorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spacesAuditorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpacesAuditorColumns = struct {
	SpaceID          string
	UserID           string
	SpacesAuditorsPK string
	RoleGUID         string
	CreatedAt        string
	UpdatedAt        string
}{
	SpaceID:          "space_id",
	UserID:           "user_id",
	SpacesAuditorsPK: "spaces_auditors_pk",
	RoleGUID:         "role_guid",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

var SpacesAuditorTableColumns = struct {
	SpaceID          string
	UserID           string
	SpacesAuditorsPK string
	RoleGUID         string
	CreatedAt        string
	UpdatedAt        string
}{
	SpaceID:          "spaces_auditors.space_id",
	UserID:           "spaces_auditors.user_id",
	SpacesAuditorsPK: "spaces_auditors.spaces_auditors_pk",
	RoleGUID:         "spaces_auditors.role_guid",
	CreatedAt:        "spaces_auditors.created_at",
	UpdatedAt:        "spaces_auditors.updated_at",
}

// Generated where

var SpacesAuditorWhere = struct {
	SpaceID          whereHelperint
	UserID           whereHelperint
	SpacesAuditorsPK whereHelperint
	RoleGUID         whereHelpernull_String
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	SpaceID:          whereHelperint{field: "\"spaces_auditors\".\"space_id\""},
	UserID:           whereHelperint{field: "\"spaces_auditors\".\"user_id\""},
	SpacesAuditorsPK: whereHelperint{field: "\"spaces_auditors\".\"spaces_auditors_pk\""},
	RoleGUID:         whereHelpernull_String{field: "\"spaces_auditors\".\"role_guid\""},
	CreatedAt:        whereHelpertime_Time{field: "\"spaces_auditors\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"spaces_auditors\".\"updated_at\""},
}

// SpacesAuditorRels is where relationship names are stored.
var SpacesAuditorRels = struct {
	Space string
	User  string
}{
	Space: "Space",
	User:  "User",
}

// spacesAuditorR is where relationships are stored.
type spacesAuditorR struct {
	Space *Space `boil:"Space" json:"Space" toml:"Space" yaml:"Space"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*spacesAuditorR) NewStruct() *spacesAuditorR {
	return &spacesAuditorR{}
}

// spacesAuditorL is where Load methods for each relationship are stored.
type spacesAuditorL struct{}

var (
	spacesAuditorAllColumns            = []string{"space_id", "user_id", "spaces_auditors_pk", "role_guid", "created_at", "updated_at"}
	spacesAuditorColumnsWithoutDefault = []string{"space_id", "user_id", "role_guid"}
	spacesAuditorColumnsWithDefault    = []string{"spaces_auditors_pk", "created_at", "updated_at"}
	spacesAuditorPrimaryKeyColumns     = []string{"spaces_auditors_pk"}
)

type (
	// SpacesAuditorSlice is an alias for a slice of pointers to SpacesAuditor.
	// This should almost always be used instead of []SpacesAuditor.
	SpacesAuditorSlice []*SpacesAuditor

	spacesAuditorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spacesAuditorType                 = reflect.TypeOf(&SpacesAuditor{})
	spacesAuditorMapping              = queries.MakeStructMapping(spacesAuditorType)
	spacesAuditorPrimaryKeyMapping, _ = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, spacesAuditorPrimaryKeyColumns)
	spacesAuditorInsertCacheMut       sync.RWMutex
	spacesAuditorInsertCache          = make(map[string]insertCache)
	spacesAuditorUpdateCacheMut       sync.RWMutex
	spacesAuditorUpdateCache          = make(map[string]updateCache)
	spacesAuditorUpsertCacheMut       sync.RWMutex
	spacesAuditorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type SpacesAuditorFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*SpacesAuditor, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (SpacesAuditorSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single spacesAuditor record from the query.
func (q spacesAuditorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SpacesAuditor, error) {
	o := &SpacesAuditor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for spaces_auditors")
	}

	return o, nil
}

// All returns all SpacesAuditor records from the query.
func (q spacesAuditorQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpacesAuditorSlice, error) {
	var o []*SpacesAuditor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SpacesAuditor slice")
	}

	return o, nil
}

// Count returns the count of all SpacesAuditor records in the query.
func (q spacesAuditorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count spaces_auditors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q spacesAuditorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if spaces_auditors exists")
	}

	return count > 0, nil
}

// Space pointed to by the foreign key.
func (q spacesAuditorQuery) Space(o *SpacesAuditor, mods ...qm.QueryMod) spaceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SpaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Spaces(queryMods...)
	queries.SetFrom(query.Query, "\"spaces\"")

	return query
}

// User pointed to by the foreign key.
func (q spacesAuditorQuery) User(o *SpacesAuditor, mods ...qm.QueryMod) userQuery {
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
func (spacesAuditorL) LoadSpace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesAuditor interface{}, mods queries.Applicator) error {
	var slice []*SpacesAuditor
	var object *SpacesAuditor

	if singular {
		object = maybeSpacesAuditor.(*SpacesAuditor)
	} else {
		slice = *maybeSpacesAuditor.(*[]*SpacesAuditor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesAuditorR{}
		}
		args = append(args, object.SpaceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesAuditorR{}
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
		foreign.R.SpacesAuditors = append(foreign.R.SpacesAuditors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SpaceID == foreign.ID {
				local.R.Space = foreign
				if foreign.R == nil {
					foreign.R = &spaceR{}
				}
				foreign.R.SpacesAuditors = append(foreign.R.SpacesAuditors, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spacesAuditorL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesAuditor interface{}, mods queries.Applicator) error {
	var slice []*SpacesAuditor
	var object *SpacesAuditor

	if singular {
		object = maybeSpacesAuditor.(*SpacesAuditor)
	} else {
		slice = *maybeSpacesAuditor.(*[]*SpacesAuditor)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesAuditorR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesAuditorR{}
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
		foreign.R.SpacesAuditors = append(foreign.R.SpacesAuditors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.SpacesAuditors = append(foreign.R.SpacesAuditors, local)
				break
			}
		}
	}

	return nil
}

// SetSpace of the spacesAuditor to the related item.
// Sets o.R.Space to related.
// Adds o to related.R.SpacesAuditors.
func (q spacesAuditorQuery) SetSpace(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Space) error {
	var err error
	if insert {
		if err = Spaces().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_auditors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"space_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesAuditorPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesAuditorsPK}

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
		o.R = &spacesAuditorR{
			Space: related,
		}
	} else {
		o.R.Space = related
	}

	if related.R == nil {
		related.R = &spaceR{
			SpacesAuditors: SpacesAuditorSlice{o},
		}
	} else {
		related.R.SpacesAuditors = append(related.R.SpacesAuditors, o)
	}

	return nil
}

// SetUser of the spacesAuditor to the related item.
// Sets o.R.User to related.
// Adds o to related.R.SpacesAuditors.
func (q spacesAuditorQuery) SetUser(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = Users().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_auditors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesAuditorPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesAuditorsPK}

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
		o.R = &spacesAuditorR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			SpacesAuditors: SpacesAuditorSlice{o},
		}
	} else {
		related.R.SpacesAuditors = append(related.R.SpacesAuditors, o)
	}

	return nil
}

// SpacesAuditors retrieves all the records using an executor.
func SpacesAuditors(mods ...qm.QueryMod) spacesAuditorQuery {
	mods = append(mods, qm.From("\"spaces_auditors\""))
	return spacesAuditorQuery{NewQuery(mods...)}
}

type SpacesAuditorFinder interface {
	FindSpacesAuditor(ctx context.Context, exec boil.ContextExecutor, spacesAuditorsPK int, selectCols ...string) (*SpacesAuditor, error)
}

// FindSpacesAuditor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpacesAuditor(ctx context.Context, exec boil.ContextExecutor, spacesAuditorsPK int, selectCols ...string) (*SpacesAuditor, error) {
	spacesAuditorObj := &SpacesAuditor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"spaces_auditors\" where \"spaces_auditors_pk\"=$1", sel,
	)

	q := queries.Raw(query, spacesAuditorsPK)

	err := q.Bind(ctx, exec, spacesAuditorObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from spaces_auditors")
	}

	return spacesAuditorObj, nil
}

type SpacesAuditorInserter interface {
	Insert(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q spacesAuditorQuery) Insert(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_auditors provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(spacesAuditorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spacesAuditorInsertCacheMut.RLock()
	cache, cached := spacesAuditorInsertCache[key]
	spacesAuditorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spacesAuditorAllColumns,
			spacesAuditorColumnsWithDefault,
			spacesAuditorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"spaces_auditors\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"spaces_auditors\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into spaces_auditors")
	}

	if !cached {
		spacesAuditorInsertCacheMut.Lock()
		spacesAuditorInsertCache[key] = cache
		spacesAuditorInsertCacheMut.Unlock()
	}

	return nil
}

type SpacesAuditorUpdater interface {
	Update(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the SpacesAuditor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q spacesAuditorQuery) Update(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	spacesAuditorUpdateCacheMut.RLock()
	cache, cached := spacesAuditorUpdateCache[key]
	spacesAuditorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spacesAuditorAllColumns,
			spacesAuditorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update spaces_auditors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"spaces_auditors\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, spacesAuditorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spacesAuditorType, spacesAuditorMapping, append(wl, spacesAuditorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update spaces_auditors row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for spaces_auditors")
	}

	if !cached {
		spacesAuditorUpdateCacheMut.Lock()
		spacesAuditorUpdateCache[key] = cache
		spacesAuditorUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q spacesAuditorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for spaces_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for spaces_auditors")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q spacesAuditorQuery) UpdateAllSlice(o SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"spaces_auditors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, spacesAuditorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in spacesAuditor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all spacesAuditor")
	}
	return rowsAff, nil
}

type SpacesAuditorDeleter interface {
	Delete(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single SpacesAuditor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q spacesAuditorQuery) Delete(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SpacesAuditor provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spacesAuditorPrimaryKeyMapping)
	sql := "DELETE FROM \"spaces_auditors\" WHERE \"spaces_auditors_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from spaces_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for spaces_auditors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q spacesAuditorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no spacesAuditorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spaces_auditors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_auditors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q spacesAuditorQuery) DeleteAllSlice(o SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"spaces_auditors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesAuditorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spacesAuditor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_auditors")
	}

	return rowsAff, nil
}

type SpacesAuditorReloader interface {
	Reload(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q spacesAuditorQuery) Reload(o *SpacesAuditor, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpacesAuditor(ctx, exec, o.SpacesAuditorsPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q spacesAuditorQuery) ReloadAll(o *SpacesAuditorSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpacesAuditorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesAuditorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"spaces_auditors\".* FROM \"spaces_auditors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesAuditorPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpacesAuditorSlice")
	}

	*o = slice

	return nil
}

// SpacesAuditorExists checks if the SpacesAuditor row exists.
func SpacesAuditorExists(ctx context.Context, exec boil.ContextExecutor, spacesAuditorsPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"spaces_auditors\" where \"spaces_auditors_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, spacesAuditorsPK)
	}
	row := exec.QueryRowContext(ctx, sql, spacesAuditorsPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if spaces_auditors exists")
	}

	return exists, nil
}
