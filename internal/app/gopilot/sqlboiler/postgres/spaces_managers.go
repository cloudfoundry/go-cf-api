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

// SpacesManager is an object representing the database table.
type SpacesManager struct {
	SpaceID          int         `boil:"space_id" json:"space_id" toml:"space_id" yaml:"space_id"`
	UserID           int         `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	SpacesManagersPK int         `boil:"spaces_managers_pk" json:"spaces_managers_pk" toml:"spaces_managers_pk" yaml:"spaces_managers_pk"`
	RoleGUID         null.String `boil:"role_guid" json:"role_guid,omitempty" toml:"role_guid" yaml:"role_guid,omitempty"`
	CreatedAt        time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *spacesManagerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L spacesManagerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SpacesManagerColumns = struct {
	SpaceID          string
	UserID           string
	SpacesManagersPK string
	RoleGUID         string
	CreatedAt        string
	UpdatedAt        string
}{
	SpaceID:          "space_id",
	UserID:           "user_id",
	SpacesManagersPK: "spaces_managers_pk",
	RoleGUID:         "role_guid",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

var SpacesManagerTableColumns = struct {
	SpaceID          string
	UserID           string
	SpacesManagersPK string
	RoleGUID         string
	CreatedAt        string
	UpdatedAt        string
}{
	SpaceID:          "spaces_managers.space_id",
	UserID:           "spaces_managers.user_id",
	SpacesManagersPK: "spaces_managers.spaces_managers_pk",
	RoleGUID:         "spaces_managers.role_guid",
	CreatedAt:        "spaces_managers.created_at",
	UpdatedAt:        "spaces_managers.updated_at",
}

// Generated where

var SpacesManagerWhere = struct {
	SpaceID          whereHelperint
	UserID           whereHelperint
	SpacesManagersPK whereHelperint
	RoleGUID         whereHelpernull_String
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
}{
	SpaceID:          whereHelperint{field: "\"spaces_managers\".\"space_id\""},
	UserID:           whereHelperint{field: "\"spaces_managers\".\"user_id\""},
	SpacesManagersPK: whereHelperint{field: "\"spaces_managers\".\"spaces_managers_pk\""},
	RoleGUID:         whereHelpernull_String{field: "\"spaces_managers\".\"role_guid\""},
	CreatedAt:        whereHelpertime_Time{field: "\"spaces_managers\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"spaces_managers\".\"updated_at\""},
}

// SpacesManagerRels is where relationship names are stored.
var SpacesManagerRels = struct {
	Space string
	User  string
}{
	Space: "Space",
	User:  "User",
}

// spacesManagerR is where relationships are stored.
type spacesManagerR struct {
	Space *Space `boil:"Space" json:"Space" toml:"Space" yaml:"Space"`
	User  *User  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*spacesManagerR) NewStruct() *spacesManagerR {
	return &spacesManagerR{}
}

// spacesManagerL is where Load methods for each relationship are stored.
type spacesManagerL struct{}

var (
	spacesManagerAllColumns            = []string{"space_id", "user_id", "spaces_managers_pk", "role_guid", "created_at", "updated_at"}
	spacesManagerColumnsWithoutDefault = []string{"space_id", "user_id", "role_guid"}
	spacesManagerColumnsWithDefault    = []string{"spaces_managers_pk", "created_at", "updated_at"}
	spacesManagerPrimaryKeyColumns     = []string{"spaces_managers_pk"}
)

type (
	// SpacesManagerSlice is an alias for a slice of pointers to SpacesManager.
	// This should almost always be used instead of []SpacesManager.
	SpacesManagerSlice []*SpacesManager
	// SpacesManagerHook is the signature for custom SpacesManager hook methods
	SpacesManagerHook func(context.Context, boil.ContextExecutor, *SpacesManager) error

	spacesManagerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	spacesManagerType                 = reflect.TypeOf(&SpacesManager{})
	spacesManagerMapping              = queries.MakeStructMapping(spacesManagerType)
	spacesManagerPrimaryKeyMapping, _ = queries.BindMapping(spacesManagerType, spacesManagerMapping, spacesManagerPrimaryKeyColumns)
	spacesManagerInsertCacheMut       sync.RWMutex
	spacesManagerInsertCache          = make(map[string]insertCache)
	spacesManagerUpdateCacheMut       sync.RWMutex
	spacesManagerUpdateCache          = make(map[string]updateCache)
	spacesManagerUpsertCacheMut       sync.RWMutex
	spacesManagerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var spacesManagerBeforeInsertHooks []SpacesManagerHook
var spacesManagerBeforeUpdateHooks []SpacesManagerHook
var spacesManagerBeforeDeleteHooks []SpacesManagerHook
var spacesManagerBeforeUpsertHooks []SpacesManagerHook

var spacesManagerAfterInsertHooks []SpacesManagerHook
var spacesManagerAfterSelectHooks []SpacesManagerHook
var spacesManagerAfterUpdateHooks []SpacesManagerHook
var spacesManagerAfterDeleteHooks []SpacesManagerHook
var spacesManagerAfterUpsertHooks []SpacesManagerHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SpacesManager) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SpacesManager) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SpacesManager) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SpacesManager) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SpacesManager) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SpacesManager) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SpacesManager) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SpacesManager) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SpacesManager) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range spacesManagerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSpacesManagerHook registers your hook function for all future operations.
func AddSpacesManagerHook(hookPoint boil.HookPoint, spacesManagerHook SpacesManagerHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		spacesManagerBeforeInsertHooks = append(spacesManagerBeforeInsertHooks, spacesManagerHook)
	case boil.BeforeUpdateHook:
		spacesManagerBeforeUpdateHooks = append(spacesManagerBeforeUpdateHooks, spacesManagerHook)
	case boil.BeforeDeleteHook:
		spacesManagerBeforeDeleteHooks = append(spacesManagerBeforeDeleteHooks, spacesManagerHook)
	case boil.BeforeUpsertHook:
		spacesManagerBeforeUpsertHooks = append(spacesManagerBeforeUpsertHooks, spacesManagerHook)
	case boil.AfterInsertHook:
		spacesManagerAfterInsertHooks = append(spacesManagerAfterInsertHooks, spacesManagerHook)
	case boil.AfterSelectHook:
		spacesManagerAfterSelectHooks = append(spacesManagerAfterSelectHooks, spacesManagerHook)
	case boil.AfterUpdateHook:
		spacesManagerAfterUpdateHooks = append(spacesManagerAfterUpdateHooks, spacesManagerHook)
	case boil.AfterDeleteHook:
		spacesManagerAfterDeleteHooks = append(spacesManagerAfterDeleteHooks, spacesManagerHook)
	case boil.AfterUpsertHook:
		spacesManagerAfterUpsertHooks = append(spacesManagerAfterUpsertHooks, spacesManagerHook)
	}
}

// One returns a single spacesManager record from the query.
func (q spacesManagerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SpacesManager, error) {
	o := &SpacesManager{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for spaces_managers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all SpacesManager records from the query.
func (q spacesManagerQuery) All(ctx context.Context, exec boil.ContextExecutor) (SpacesManagerSlice, error) {
	var o []*SpacesManager

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SpacesManager slice")
	}

	if len(spacesManagerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all SpacesManager records in the query.
func (q spacesManagerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count spaces_managers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q spacesManagerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if spaces_managers exists")
	}

	return count > 0, nil
}

// Space pointed to by the foreign key.
func (o *SpacesManager) Space(mods ...qm.QueryMod) spaceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SpaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Spaces(queryMods...)
	queries.SetFrom(query.Query, "\"spaces\"")

	return query
}

// User pointed to by the foreign key.
func (o *SpacesManager) User(mods ...qm.QueryMod) userQuery {
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
func (spacesManagerL) LoadSpace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesManager interface{}, mods queries.Applicator) error {
	var slice []*SpacesManager
	var object *SpacesManager

	if singular {
		object = maybeSpacesManager.(*SpacesManager)
	} else {
		slice = *maybeSpacesManager.(*[]*SpacesManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesManagerR{}
		}
		args = append(args, object.SpaceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesManagerR{}
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

	if len(spacesManagerAfterSelectHooks) != 0 {
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
		object.R.Space = foreign
		if foreign.R == nil {
			foreign.R = &spaceR{}
		}
		foreign.R.SpacesManagers = append(foreign.R.SpacesManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SpaceID == foreign.ID {
				local.R.Space = foreign
				if foreign.R == nil {
					foreign.R = &spaceR{}
				}
				foreign.R.SpacesManagers = append(foreign.R.SpacesManagers, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (spacesManagerL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSpacesManager interface{}, mods queries.Applicator) error {
	var slice []*SpacesManager
	var object *SpacesManager

	if singular {
		object = maybeSpacesManager.(*SpacesManager)
	} else {
		slice = *maybeSpacesManager.(*[]*SpacesManager)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &spacesManagerR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &spacesManagerR{}
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

	if len(spacesManagerAfterSelectHooks) != 0 {
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
		foreign.R.SpacesManagers = append(foreign.R.SpacesManagers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.SpacesManagers = append(foreign.R.SpacesManagers, local)
				break
			}
		}
	}

	return nil
}

// SetSpace of the spacesManager to the related item.
// Sets o.R.Space to related.
// Adds o to related.R.SpacesManagers.
func (o *SpacesManager) SetSpace(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Space) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_managers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"space_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesManagersPK}

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
		o.R = &spacesManagerR{
			Space: related,
		}
	} else {
		o.R.Space = related
	}

	if related.R == nil {
		related.R = &spaceR{
			SpacesManagers: SpacesManagerSlice{o},
		}
	} else {
		related.R.SpacesManagers = append(related.R.SpacesManagers, o)
	}

	return nil
}

// SetUser of the spacesManager to the related item.
// Sets o.R.User to related.
// Adds o to related.R.SpacesManagers.
func (o *SpacesManager) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"spaces_managers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, spacesManagerPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SpacesManagersPK}

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
		o.R = &spacesManagerR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			SpacesManagers: SpacesManagerSlice{o},
		}
	} else {
		related.R.SpacesManagers = append(related.R.SpacesManagers, o)
	}

	return nil
}

// SpacesManagers retrieves all the records using an executor.
func SpacesManagers(mods ...qm.QueryMod) spacesManagerQuery {
	mods = append(mods, qm.From("\"spaces_managers\""))
	return spacesManagerQuery{NewQuery(mods...)}
}

// FindSpacesManager retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSpacesManager(ctx context.Context, exec boil.ContextExecutor, spacesManagersPK int, selectCols ...string) (*SpacesManager, error) {
	spacesManagerObj := &SpacesManager{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"spaces_managers\" where \"spaces_managers_pk\"=$1", sel,
	)

	q := queries.Raw(query, spacesManagersPK)

	err := q.Bind(ctx, exec, spacesManagerObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from spaces_managers")
	}

	if err = spacesManagerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return spacesManagerObj, err
	}

	return spacesManagerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *SpacesManager) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_managers provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(spacesManagerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	spacesManagerInsertCacheMut.RLock()
	cache, cached := spacesManagerInsertCache[key]
	spacesManagerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			spacesManagerAllColumns,
			spacesManagerColumnsWithDefault,
			spacesManagerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(spacesManagerType, spacesManagerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(spacesManagerType, spacesManagerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"spaces_managers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"spaces_managers\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into spaces_managers")
	}

	if !cached {
		spacesManagerInsertCacheMut.Lock()
		spacesManagerInsertCache[key] = cache
		spacesManagerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the SpacesManager.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *SpacesManager) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	spacesManagerUpdateCacheMut.RLock()
	cache, cached := spacesManagerUpdateCache[key]
	spacesManagerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			spacesManagerAllColumns,
			spacesManagerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update spaces_managers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"spaces_managers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, spacesManagerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(spacesManagerType, spacesManagerMapping, append(wl, spacesManagerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update spaces_managers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for spaces_managers")
	}

	if !cached {
		spacesManagerUpdateCacheMut.Lock()
		spacesManagerUpdateCache[key] = cache
		spacesManagerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q spacesManagerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for spaces_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for spaces_managers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SpacesManagerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"spaces_managers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, spacesManagerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in spacesManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all spacesManager")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SpacesManager) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no spaces_managers provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(spacesManagerColumnsWithDefault, o)

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

	spacesManagerUpsertCacheMut.RLock()
	cache, cached := spacesManagerUpsertCache[key]
	spacesManagerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			spacesManagerAllColumns,
			spacesManagerColumnsWithDefault,
			spacesManagerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			spacesManagerAllColumns,
			spacesManagerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert spaces_managers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(spacesManagerPrimaryKeyColumns))
			copy(conflict, spacesManagerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"spaces_managers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(spacesManagerType, spacesManagerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(spacesManagerType, spacesManagerMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert spaces_managers")
	}

	if !cached {
		spacesManagerUpsertCacheMut.Lock()
		spacesManagerUpsertCache[key] = cache
		spacesManagerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single SpacesManager record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *SpacesManager) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SpacesManager provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), spacesManagerPrimaryKeyMapping)
	sql := "DELETE FROM \"spaces_managers\" WHERE \"spaces_managers_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from spaces_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for spaces_managers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q spacesManagerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no spacesManagerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spaces_managers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_managers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SpacesManagerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(spacesManagerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"spaces_managers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesManagerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from spacesManager slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for spaces_managers")
	}

	if len(spacesManagerAfterDeleteHooks) != 0 {
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
func (o *SpacesManager) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSpacesManager(ctx, exec, o.SpacesManagersPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SpacesManagerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SpacesManagerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), spacesManagerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"spaces_managers\".* FROM \"spaces_managers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, spacesManagerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SpacesManagerSlice")
	}

	*o = slice

	return nil
}

// SpacesManagerExists checks if the SpacesManager row exists.
func SpacesManagerExists(ctx context.Context, exec boil.ContextExecutor, spacesManagersPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"spaces_managers\" where \"spaces_managers_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, spacesManagersPK)
	}
	row := exec.QueryRowContext(ctx, sql, spacesManagersPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if spaces_managers exists")
	}

	return exists, nil
}
