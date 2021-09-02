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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// SecurityGroupsSpace is an object representing the database table.
type SecurityGroupsSpace struct {
	SecurityGroupID        int `boil:"security_group_id" json:"security_group_id" toml:"security_group_id" yaml:"security_group_id"`
	SpaceID                int `boil:"space_id" json:"space_id" toml:"space_id" yaml:"space_id"`
	SecurityGroupsSpacesPK int `boil:"security_groups_spaces_pk" json:"security_groups_spaces_pk" toml:"security_groups_spaces_pk" yaml:"security_groups_spaces_pk"`

	R *securityGroupsSpaceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L securityGroupsSpaceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SecurityGroupsSpaceColumns = struct {
	SecurityGroupID        string
	SpaceID                string
	SecurityGroupsSpacesPK string
}{
	SecurityGroupID:        "security_group_id",
	SpaceID:                "space_id",
	SecurityGroupsSpacesPK: "security_groups_spaces_pk",
}

var SecurityGroupsSpaceTableColumns = struct {
	SecurityGroupID        string
	SpaceID                string
	SecurityGroupsSpacesPK string
}{
	SecurityGroupID:        "security_groups_spaces.security_group_id",
	SpaceID:                "security_groups_spaces.space_id",
	SecurityGroupsSpacesPK: "security_groups_spaces.security_groups_spaces_pk",
}

// Generated where

var SecurityGroupsSpaceWhere = struct {
	SecurityGroupID        whereHelperint
	SpaceID                whereHelperint
	SecurityGroupsSpacesPK whereHelperint
}{
	SecurityGroupID:        whereHelperint{field: "\"security_groups_spaces\".\"security_group_id\""},
	SpaceID:                whereHelperint{field: "\"security_groups_spaces\".\"space_id\""},
	SecurityGroupsSpacesPK: whereHelperint{field: "\"security_groups_spaces\".\"security_groups_spaces_pk\""},
}

// SecurityGroupsSpaceRels is where relationship names are stored.
var SecurityGroupsSpaceRels = struct {
	SecurityGroup string
	Space         string
}{
	SecurityGroup: "SecurityGroup",
	Space:         "Space",
}

// securityGroupsSpaceR is where relationships are stored.
type securityGroupsSpaceR struct {
	SecurityGroup *SecurityGroup `boil:"SecurityGroup" json:"SecurityGroup" toml:"SecurityGroup" yaml:"SecurityGroup"`
	Space         *Space         `boil:"Space" json:"Space" toml:"Space" yaml:"Space"`
}

// NewStruct creates a new relationship struct
func (*securityGroupsSpaceR) NewStruct() *securityGroupsSpaceR {
	return &securityGroupsSpaceR{}
}

// securityGroupsSpaceL is where Load methods for each relationship are stored.
type securityGroupsSpaceL struct{}

var (
	securityGroupsSpaceAllColumns            = []string{"security_group_id", "space_id", "security_groups_spaces_pk"}
	securityGroupsSpaceColumnsWithoutDefault = []string{"security_group_id", "space_id"}
	securityGroupsSpaceColumnsWithDefault    = []string{"security_groups_spaces_pk"}
	securityGroupsSpacePrimaryKeyColumns     = []string{"security_groups_spaces_pk"}
)

type (
	// SecurityGroupsSpaceSlice is an alias for a slice of pointers to SecurityGroupsSpace.
	// This should almost always be used instead of []SecurityGroupsSpace.
	SecurityGroupsSpaceSlice []*SecurityGroupsSpace

	SecurityGroupsSpaceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	securityGroupsSpaceType                 = reflect.TypeOf(&SecurityGroupsSpace{})
	securityGroupsSpaceMapping              = queries.MakeStructMapping(securityGroupsSpaceType)
	securityGroupsSpacePrimaryKeyMapping, _ = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, securityGroupsSpacePrimaryKeyColumns)
	securityGroupsSpaceInsertCacheMut       sync.RWMutex
	securityGroupsSpaceInsertCache          = make(map[string]insertCache)
	securityGroupsSpaceUpdateCacheMut       sync.RWMutex
	securityGroupsSpaceUpdateCache          = make(map[string]updateCache)
	securityGroupsSpaceUpsertCacheMut       sync.RWMutex
	securityGroupsSpaceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type SecurityGroupsSpaceFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*SecurityGroupsSpace, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (SecurityGroupsSpaceSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single securityGroupsSpace record from the query.
func (q SecurityGroupsSpaceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*SecurityGroupsSpace, error) {
	o := &SecurityGroupsSpace{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for security_groups_spaces")
	}

	return o, nil
}

// All returns all SecurityGroupsSpace records from the query.
func (q SecurityGroupsSpaceQuery) All(ctx context.Context, exec boil.ContextExecutor) (SecurityGroupsSpaceSlice, error) {
	var o []*SecurityGroupsSpace

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SecurityGroupsSpace slice")
	}

	return o, nil
}

// Count returns the count of all SecurityGroupsSpace records in the query.
func (q SecurityGroupsSpaceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count security_groups_spaces rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q SecurityGroupsSpaceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if security_groups_spaces exists")
	}

	return count > 0, nil
}

// SecurityGroup pointed to by the foreign key.
func (q SecurityGroupsSpaceQuery) SecurityGroup(o *SecurityGroupsSpace, mods ...qm.QueryMod) SecurityGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SecurityGroupID),
	}

	queryMods = append(queryMods, mods...)

	query := SecurityGroups(queryMods...)
	queries.SetFrom(query.Query, "\"security_groups\"")

	return query
}

// Space pointed to by the foreign key.
func (q SecurityGroupsSpaceQuery) Space(o *SecurityGroupsSpace, mods ...qm.QueryMod) SpaceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SpaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Spaces(queryMods...)
	queries.SetFrom(query.Query, "\"spaces\"")

	return query
}

// LoadSecurityGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (securityGroupsSpaceL) LoadSecurityGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSecurityGroupsSpace interface{}, mods queries.Applicator) error {
	var slice []*SecurityGroupsSpace
	var object *SecurityGroupsSpace

	if singular {
		object = maybeSecurityGroupsSpace.(*SecurityGroupsSpace)
	} else {
		slice = *maybeSecurityGroupsSpace.(*[]*SecurityGroupsSpace)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &securityGroupsSpaceR{}
		}
		args = append(args, object.SecurityGroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &securityGroupsSpaceR{}
			}

			for _, a := range args {
				if a == obj.SecurityGroupID {
					continue Outer
				}
			}

			args = append(args, obj.SecurityGroupID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`security_groups`),
		qm.WhereIn(`security_groups.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load SecurityGroup")
	}

	var resultSlice []*SecurityGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice SecurityGroup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for security_groups")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for security_groups")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.SecurityGroup = foreign
		if foreign.R == nil {
			foreign.R = &securityGroupR{}
		}
		foreign.R.SecurityGroupsSpaces = append(foreign.R.SecurityGroupsSpaces, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SecurityGroupID == foreign.ID {
				local.R.SecurityGroup = foreign
				if foreign.R == nil {
					foreign.R = &securityGroupR{}
				}
				foreign.R.SecurityGroupsSpaces = append(foreign.R.SecurityGroupsSpaces, local)
				break
			}
		}
	}

	return nil
}

// LoadSpace allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (securityGroupsSpaceL) LoadSpace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSecurityGroupsSpace interface{}, mods queries.Applicator) error {
	var slice []*SecurityGroupsSpace
	var object *SecurityGroupsSpace

	if singular {
		object = maybeSecurityGroupsSpace.(*SecurityGroupsSpace)
	} else {
		slice = *maybeSecurityGroupsSpace.(*[]*SecurityGroupsSpace)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &securityGroupsSpaceR{}
		}
		args = append(args, object.SpaceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &securityGroupsSpaceR{}
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
		foreign.R.SecurityGroupsSpaces = append(foreign.R.SecurityGroupsSpaces, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SpaceID == foreign.ID {
				local.R.Space = foreign
				if foreign.R == nil {
					foreign.R = &spaceR{}
				}
				foreign.R.SecurityGroupsSpaces = append(foreign.R.SecurityGroupsSpaces, local)
				break
			}
		}
	}

	return nil
}

// SetSecurityGroup of the securityGroupsSpace to the related item.
// Sets o.R.SecurityGroup to related.
// Adds o to related.R.SecurityGroupsSpaces.
func (q SecurityGroupsSpaceQuery) SetSecurityGroup(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, insert bool, related *SecurityGroup) error {
	var err error
	if insert {
		if err = SecurityGroups().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"security_groups_spaces\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"security_group_id"}),
		strmangle.WhereClause("\"", "\"", 2, securityGroupsSpacePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SecurityGroupsSpacesPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SecurityGroupID = related.ID
	if o.R == nil {
		o.R = &securityGroupsSpaceR{
			SecurityGroup: related,
		}
	} else {
		o.R.SecurityGroup = related
	}

	if related.R == nil {
		related.R = &securityGroupR{
			SecurityGroupsSpaces: SecurityGroupsSpaceSlice{o},
		}
	} else {
		related.R.SecurityGroupsSpaces = append(related.R.SecurityGroupsSpaces, o)
	}

	return nil
}

// SetSpace of the securityGroupsSpace to the related item.
// Sets o.R.Space to related.
// Adds o to related.R.SecurityGroupsSpaces.
func (q SecurityGroupsSpaceQuery) SetSpace(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Space) error {
	var err error
	if insert {
		if err = Spaces().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"security_groups_spaces\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"space_id"}),
		strmangle.WhereClause("\"", "\"", 2, securityGroupsSpacePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.SecurityGroupsSpacesPK}

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
		o.R = &securityGroupsSpaceR{
			Space: related,
		}
	} else {
		o.R.Space = related
	}

	if related.R == nil {
		related.R = &spaceR{
			SecurityGroupsSpaces: SecurityGroupsSpaceSlice{o},
		}
	} else {
		related.R.SecurityGroupsSpaces = append(related.R.SecurityGroupsSpaces, o)
	}

	return nil
}

// SecurityGroupsSpaces retrieves all the records using an executor.
func SecurityGroupsSpaces(mods ...qm.QueryMod) SecurityGroupsSpaceQuery {
	mods = append(mods, qm.From("\"security_groups_spaces\""))
	return SecurityGroupsSpaceQuery{NewQuery(mods...)}
}

type SecurityGroupsSpaceFinder interface {
	FindSecurityGroupsSpace(ctx context.Context, exec boil.ContextExecutor, securityGroupsSpacesPK int, selectCols ...string) (*SecurityGroupsSpace, error)
}

// FindSecurityGroupsSpace retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSecurityGroupsSpace(ctx context.Context, exec boil.ContextExecutor, securityGroupsSpacesPK int, selectCols ...string) (*SecurityGroupsSpace, error) {
	securityGroupsSpaceObj := &SecurityGroupsSpace{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"security_groups_spaces\" where \"security_groups_spaces_pk\"=$1", sel,
	)

	q := queries.Raw(query, securityGroupsSpacesPK)

	err := q.Bind(ctx, exec, securityGroupsSpaceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from security_groups_spaces")
	}

	return securityGroupsSpaceObj, nil
}

type SecurityGroupsSpaceInserter interface {
	Insert(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q SecurityGroupsSpaceQuery) Insert(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no security_groups_spaces provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(securityGroupsSpaceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	securityGroupsSpaceInsertCacheMut.RLock()
	cache, cached := securityGroupsSpaceInsertCache[key]
	securityGroupsSpaceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			securityGroupsSpaceAllColumns,
			securityGroupsSpaceColumnsWithDefault,
			securityGroupsSpaceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"security_groups_spaces\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"security_groups_spaces\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into security_groups_spaces")
	}

	if !cached {
		securityGroupsSpaceInsertCacheMut.Lock()
		securityGroupsSpaceInsertCache[key] = cache
		securityGroupsSpaceInsertCacheMut.Unlock()
	}

	return nil
}

type SecurityGroupsSpaceUpdater interface {
	Update(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the SecurityGroupsSpace.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q SecurityGroupsSpaceQuery) Update(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	securityGroupsSpaceUpdateCacheMut.RLock()
	cache, cached := securityGroupsSpaceUpdateCache[key]
	securityGroupsSpaceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			securityGroupsSpaceAllColumns,
			securityGroupsSpacePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update security_groups_spaces, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"security_groups_spaces\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, securityGroupsSpacePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, append(wl, securityGroupsSpacePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update security_groups_spaces row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for security_groups_spaces")
	}

	if !cached {
		securityGroupsSpaceUpdateCacheMut.Lock()
		securityGroupsSpaceUpdateCache[key] = cache
		securityGroupsSpaceUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q SecurityGroupsSpaceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for security_groups_spaces")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q SecurityGroupsSpaceQuery) UpdateAllSlice(o SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), securityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"security_groups_spaces\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, securityGroupsSpacePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in securityGroupsSpace slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all securityGroupsSpace")
	}
	return rowsAff, nil
}

type SecurityGroupsSpaceDeleter interface {
	Delete(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single SecurityGroupsSpace record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q SecurityGroupsSpaceQuery) Delete(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no SecurityGroupsSpace provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), securityGroupsSpacePrimaryKeyMapping)
	sql := "DELETE FROM \"security_groups_spaces\" WHERE \"security_groups_spaces_pk\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for security_groups_spaces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q SecurityGroupsSpaceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no securityGroupsSpaceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for security_groups_spaces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q SecurityGroupsSpaceQuery) DeleteAllSlice(o SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), securityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"security_groups_spaces\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, securityGroupsSpacePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from securityGroupsSpace slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for security_groups_spaces")
	}

	return rowsAff, nil
}

type SecurityGroupsSpaceReloader interface {
	Reload(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q SecurityGroupsSpaceQuery) Reload(o *SecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSecurityGroupsSpace(ctx, exec, o.SecurityGroupsSpacesPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q SecurityGroupsSpaceQuery) ReloadAll(o *SecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SecurityGroupsSpaceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), securityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"security_groups_spaces\".* FROM \"security_groups_spaces\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, securityGroupsSpacePrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SecurityGroupsSpaceSlice")
	}

	*o = slice

	return nil
}

// SecurityGroupsSpaceExists checks if the SecurityGroupsSpace row exists.
func SecurityGroupsSpaceExists(ctx context.Context, exec boil.ContextExecutor, securityGroupsSpacesPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"security_groups_spaces\" where \"security_groups_spaces_pk\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, securityGroupsSpacesPK)
	}
	row := exec.QueryRowContext(ctx, sql, securityGroupsSpacesPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if security_groups_spaces exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *SecurityGroupsSpace) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no security_groups_spaces provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(securityGroupsSpaceColumnsWithDefault, o)

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

	securityGroupsSpaceUpsertCacheMut.RLock()
	cache, cached := securityGroupsSpaceUpsertCache[key]
	securityGroupsSpaceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			securityGroupsSpaceAllColumns,
			securityGroupsSpaceColumnsWithDefault,
			securityGroupsSpaceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			securityGroupsSpaceAllColumns,
			securityGroupsSpacePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert security_groups_spaces, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(securityGroupsSpacePrimaryKeyColumns))
			copy(conflict, securityGroupsSpacePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"security_groups_spaces\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(securityGroupsSpaceType, securityGroupsSpaceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert security_groups_spaces")
	}

	if !cached {
		securityGroupsSpaceUpsertCacheMut.Lock()
		securityGroupsSpaceUpsertCache[key] = cache
		securityGroupsSpaceUpsertCacheMut.Unlock()
	}

	return nil
}
