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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

type StagingSecurityGroupsSpaceUpserter interface {
	Upsert(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLStagingSecurityGroupsSpaceUniqueColumns = []string{
	"staging_security_groups_spaces_pk",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q stagingSecurityGroupsSpaceQuery) Upsert(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no staging_security_groups_spaces provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(stagingSecurityGroupsSpaceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLStagingSecurityGroupsSpaceUniqueColumns, o)

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

	stagingSecurityGroupsSpaceUpsertCacheMut.RLock()
	cache, cached := stagingSecurityGroupsSpaceUpsertCache[key]
	stagingSecurityGroupsSpaceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stagingSecurityGroupsSpaceAllColumns,
			stagingSecurityGroupsSpaceColumnsWithDefault,
			stagingSecurityGroupsSpaceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stagingSecurityGroupsSpaceAllColumns,
			stagingSecurityGroupsSpacePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert staging_security_groups_spaces, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`staging_security_groups_spaces`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `staging_security_groups_spaces` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for staging_security_groups_spaces")
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

	o.StagingSecurityGroupsSpacesPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stagingSecurityGroupsSpaceMapping["staging_security_groups_spaces_pk"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for staging_security_groups_spaces")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for staging_security_groups_spaces")
	}

CacheNoHooks:
	if !cached {
		stagingSecurityGroupsSpaceUpsertCacheMut.Lock()
		stagingSecurityGroupsSpaceUpsertCache[key] = cache
		stagingSecurityGroupsSpaceUpsertCacheMut.Unlock()
	}

	return nil
}

// StagingSecurityGroupsSpace is an object representing the database table.
type StagingSecurityGroupsSpace struct {
	StagingSecurityGroupID        int `boil:"staging_security_group_id" json:"staging_security_group_id" toml:"staging_security_group_id" yaml:"staging_security_group_id"`
	StagingSpaceID                int `boil:"staging_space_id" json:"staging_space_id" toml:"staging_space_id" yaml:"staging_space_id"`
	StagingSecurityGroupsSpacesPK int `boil:"staging_security_groups_spaces_pk" json:"staging_security_groups_spaces_pk" toml:"staging_security_groups_spaces_pk" yaml:"staging_security_groups_spaces_pk"`

	R *stagingSecurityGroupsSpaceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stagingSecurityGroupsSpaceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StagingSecurityGroupsSpaceColumns = struct {
	StagingSecurityGroupID        string
	StagingSpaceID                string
	StagingSecurityGroupsSpacesPK string
}{
	StagingSecurityGroupID:        "staging_security_group_id",
	StagingSpaceID:                "staging_space_id",
	StagingSecurityGroupsSpacesPK: "staging_security_groups_spaces_pk",
}

var StagingSecurityGroupsSpaceTableColumns = struct {
	StagingSecurityGroupID        string
	StagingSpaceID                string
	StagingSecurityGroupsSpacesPK string
}{
	StagingSecurityGroupID:        "staging_security_groups_spaces.staging_security_group_id",
	StagingSpaceID:                "staging_security_groups_spaces.staging_space_id",
	StagingSecurityGroupsSpacesPK: "staging_security_groups_spaces.staging_security_groups_spaces_pk",
}

// Generated where

var StagingSecurityGroupsSpaceWhere = struct {
	StagingSecurityGroupID        whereHelperint
	StagingSpaceID                whereHelperint
	StagingSecurityGroupsSpacesPK whereHelperint
}{
	StagingSecurityGroupID:        whereHelperint{field: "`staging_security_groups_spaces`.`staging_security_group_id`"},
	StagingSpaceID:                whereHelperint{field: "`staging_security_groups_spaces`.`staging_space_id`"},
	StagingSecurityGroupsSpacesPK: whereHelperint{field: "`staging_security_groups_spaces`.`staging_security_groups_spaces_pk`"},
}

// StagingSecurityGroupsSpaceRels is where relationship names are stored.
var StagingSecurityGroupsSpaceRels = struct {
	StagingSecurityGroup string
	StagingSpace         string
}{
	StagingSecurityGroup: "StagingSecurityGroup",
	StagingSpace:         "StagingSpace",
}

// stagingSecurityGroupsSpaceR is where relationships are stored.
type stagingSecurityGroupsSpaceR struct {
	StagingSecurityGroup *SecurityGroup `boil:"StagingSecurityGroup" json:"StagingSecurityGroup" toml:"StagingSecurityGroup" yaml:"StagingSecurityGroup"`
	StagingSpace         *Space         `boil:"StagingSpace" json:"StagingSpace" toml:"StagingSpace" yaml:"StagingSpace"`
}

// NewStruct creates a new relationship struct
func (*stagingSecurityGroupsSpaceR) NewStruct() *stagingSecurityGroupsSpaceR {
	return &stagingSecurityGroupsSpaceR{}
}

// stagingSecurityGroupsSpaceL is where Load methods for each relationship are stored.
type stagingSecurityGroupsSpaceL struct{}

var (
	stagingSecurityGroupsSpaceAllColumns            = []string{"staging_security_group_id", "staging_space_id", "staging_security_groups_spaces_pk"}
	stagingSecurityGroupsSpaceColumnsWithoutDefault = []string{"staging_security_group_id", "staging_space_id"}
	stagingSecurityGroupsSpaceColumnsWithDefault    = []string{"staging_security_groups_spaces_pk"}
	stagingSecurityGroupsSpacePrimaryKeyColumns     = []string{"staging_security_groups_spaces_pk"}
)

type (
	// StagingSecurityGroupsSpaceSlice is an alias for a slice of pointers to StagingSecurityGroupsSpace.
	// This should almost always be used instead of []StagingSecurityGroupsSpace.
	StagingSecurityGroupsSpaceSlice []*StagingSecurityGroupsSpace

	stagingSecurityGroupsSpaceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stagingSecurityGroupsSpaceType                 = reflect.TypeOf(&StagingSecurityGroupsSpace{})
	stagingSecurityGroupsSpaceMapping              = queries.MakeStructMapping(stagingSecurityGroupsSpaceType)
	stagingSecurityGroupsSpacePrimaryKeyMapping, _ = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, stagingSecurityGroupsSpacePrimaryKeyColumns)
	stagingSecurityGroupsSpaceInsertCacheMut       sync.RWMutex
	stagingSecurityGroupsSpaceInsertCache          = make(map[string]insertCache)
	stagingSecurityGroupsSpaceUpdateCacheMut       sync.RWMutex
	stagingSecurityGroupsSpaceUpdateCache          = make(map[string]updateCache)
	stagingSecurityGroupsSpaceUpsertCacheMut       sync.RWMutex
	stagingSecurityGroupsSpaceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type StagingSecurityGroupsSpaceFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*StagingSecurityGroupsSpace, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (StagingSecurityGroupsSpaceSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single stagingSecurityGroupsSpace record from the query.
func (q stagingSecurityGroupsSpaceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*StagingSecurityGroupsSpace, error) {
	o := &StagingSecurityGroupsSpace{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for staging_security_groups_spaces")
	}

	return o, nil
}

// All returns all StagingSecurityGroupsSpace records from the query.
func (q stagingSecurityGroupsSpaceQuery) All(ctx context.Context, exec boil.ContextExecutor) (StagingSecurityGroupsSpaceSlice, error) {
	var o []*StagingSecurityGroupsSpace

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to StagingSecurityGroupsSpace slice")
	}

	return o, nil
}

// Count returns the count of all StagingSecurityGroupsSpace records in the query.
func (q stagingSecurityGroupsSpaceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count staging_security_groups_spaces rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stagingSecurityGroupsSpaceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if staging_security_groups_spaces exists")
	}

	return count > 0, nil
}

// StagingSecurityGroup pointed to by the foreign key.
func (q stagingSecurityGroupsSpaceQuery) StagingSecurityGroup(o *StagingSecurityGroupsSpace, mods ...qm.QueryMod) securityGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.StagingSecurityGroupID),
	}

	queryMods = append(queryMods, mods...)

	query := SecurityGroups(queryMods...)
	queries.SetFrom(query.Query, "`security_groups`")

	return query
}

// StagingSpace pointed to by the foreign key.
func (q stagingSecurityGroupsSpaceQuery) StagingSpace(o *StagingSecurityGroupsSpace, mods ...qm.QueryMod) spaceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.StagingSpaceID),
	}

	queryMods = append(queryMods, mods...)

	query := Spaces(queryMods...)
	queries.SetFrom(query.Query, "`spaces`")

	return query
}

// LoadStagingSecurityGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stagingSecurityGroupsSpaceL) LoadStagingSecurityGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStagingSecurityGroupsSpace interface{}, mods queries.Applicator) error {
	var slice []*StagingSecurityGroupsSpace
	var object *StagingSecurityGroupsSpace

	if singular {
		object = maybeStagingSecurityGroupsSpace.(*StagingSecurityGroupsSpace)
	} else {
		slice = *maybeStagingSecurityGroupsSpace.(*[]*StagingSecurityGroupsSpace)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stagingSecurityGroupsSpaceR{}
		}
		args = append(args, object.StagingSecurityGroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stagingSecurityGroupsSpaceR{}
			}

			for _, a := range args {
				if a == obj.StagingSecurityGroupID {
					continue Outer
				}
			}

			args = append(args, obj.StagingSecurityGroupID)

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
		object.R.StagingSecurityGroup = foreign
		if foreign.R == nil {
			foreign.R = &securityGroupR{}
		}
		foreign.R.StagingSecurityGroupStagingSecurityGroupsSpaces = append(foreign.R.StagingSecurityGroupStagingSecurityGroupsSpaces, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StagingSecurityGroupID == foreign.ID {
				local.R.StagingSecurityGroup = foreign
				if foreign.R == nil {
					foreign.R = &securityGroupR{}
				}
				foreign.R.StagingSecurityGroupStagingSecurityGroupsSpaces = append(foreign.R.StagingSecurityGroupStagingSecurityGroupsSpaces, local)
				break
			}
		}
	}

	return nil
}

// LoadStagingSpace allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stagingSecurityGroupsSpaceL) LoadStagingSpace(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStagingSecurityGroupsSpace interface{}, mods queries.Applicator) error {
	var slice []*StagingSecurityGroupsSpace
	var object *StagingSecurityGroupsSpace

	if singular {
		object = maybeStagingSecurityGroupsSpace.(*StagingSecurityGroupsSpace)
	} else {
		slice = *maybeStagingSecurityGroupsSpace.(*[]*StagingSecurityGroupsSpace)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stagingSecurityGroupsSpaceR{}
		}
		args = append(args, object.StagingSpaceID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stagingSecurityGroupsSpaceR{}
			}

			for _, a := range args {
				if a == obj.StagingSpaceID {
					continue Outer
				}
			}

			args = append(args, obj.StagingSpaceID)

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
		object.R.StagingSpace = foreign
		if foreign.R == nil {
			foreign.R = &spaceR{}
		}
		foreign.R.StagingSpaceStagingSecurityGroupsSpaces = append(foreign.R.StagingSpaceStagingSecurityGroupsSpaces, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.StagingSpaceID == foreign.ID {
				local.R.StagingSpace = foreign
				if foreign.R == nil {
					foreign.R = &spaceR{}
				}
				foreign.R.StagingSpaceStagingSecurityGroupsSpaces = append(foreign.R.StagingSpaceStagingSecurityGroupsSpaces, local)
				break
			}
		}
	}

	return nil
}

// SetStagingSecurityGroup of the stagingSecurityGroupsSpace to the related item.
// Sets o.R.StagingSecurityGroup to related.
// Adds o to related.R.StagingSecurityGroupStagingSecurityGroupsSpaces.
func (q stagingSecurityGroupsSpaceQuery) SetStagingSecurityGroup(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, insert bool, related *SecurityGroup) error {
	var err error
	if insert {
		if err = SecurityGroups().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `staging_security_groups_spaces` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"staging_security_group_id"}),
		strmangle.WhereClause("`", "`", 0, stagingSecurityGroupsSpacePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.StagingSecurityGroupsSpacesPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StagingSecurityGroupID = related.ID
	if o.R == nil {
		o.R = &stagingSecurityGroupsSpaceR{
			StagingSecurityGroup: related,
		}
	} else {
		o.R.StagingSecurityGroup = related
	}

	if related.R == nil {
		related.R = &securityGroupR{
			StagingSecurityGroupStagingSecurityGroupsSpaces: StagingSecurityGroupsSpaceSlice{o},
		}
	} else {
		related.R.StagingSecurityGroupStagingSecurityGroupsSpaces = append(related.R.StagingSecurityGroupStagingSecurityGroupsSpaces, o)
	}

	return nil
}

// SetStagingSpace of the stagingSecurityGroupsSpace to the related item.
// Sets o.R.StagingSpace to related.
// Adds o to related.R.StagingSpaceStagingSecurityGroupsSpaces.
func (q stagingSecurityGroupsSpaceQuery) SetStagingSpace(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Space) error {
	var err error
	if insert {
		if err = Spaces().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `staging_security_groups_spaces` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"staging_space_id"}),
		strmangle.WhereClause("`", "`", 0, stagingSecurityGroupsSpacePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.StagingSecurityGroupsSpacesPK}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.StagingSpaceID = related.ID
	if o.R == nil {
		o.R = &stagingSecurityGroupsSpaceR{
			StagingSpace: related,
		}
	} else {
		o.R.StagingSpace = related
	}

	if related.R == nil {
		related.R = &spaceR{
			StagingSpaceStagingSecurityGroupsSpaces: StagingSecurityGroupsSpaceSlice{o},
		}
	} else {
		related.R.StagingSpaceStagingSecurityGroupsSpaces = append(related.R.StagingSpaceStagingSecurityGroupsSpaces, o)
	}

	return nil
}

// StagingSecurityGroupsSpaces retrieves all the records using an executor.
func StagingSecurityGroupsSpaces(mods ...qm.QueryMod) stagingSecurityGroupsSpaceQuery {
	mods = append(mods, qm.From("`staging_security_groups_spaces`"))
	return stagingSecurityGroupsSpaceQuery{NewQuery(mods...)}
}

type StagingSecurityGroupsSpaceFinder interface {
	FindStagingSecurityGroupsSpace(ctx context.Context, exec boil.ContextExecutor, stagingSecurityGroupsSpacesPK int, selectCols ...string) (*StagingSecurityGroupsSpace, error)
}

// FindStagingSecurityGroupsSpace retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStagingSecurityGroupsSpace(ctx context.Context, exec boil.ContextExecutor, stagingSecurityGroupsSpacesPK int, selectCols ...string) (*StagingSecurityGroupsSpace, error) {
	stagingSecurityGroupsSpaceObj := &StagingSecurityGroupsSpace{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `staging_security_groups_spaces` where `staging_security_groups_spaces_pk`=?", sel,
	)

	q := queries.Raw(query, stagingSecurityGroupsSpacesPK)

	err := q.Bind(ctx, exec, stagingSecurityGroupsSpaceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from staging_security_groups_spaces")
	}

	return stagingSecurityGroupsSpaceObj, nil
}

type StagingSecurityGroupsSpaceInserter interface {
	Insert(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q stagingSecurityGroupsSpaceQuery) Insert(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no staging_security_groups_spaces provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(stagingSecurityGroupsSpaceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stagingSecurityGroupsSpaceInsertCacheMut.RLock()
	cache, cached := stagingSecurityGroupsSpaceInsertCache[key]
	stagingSecurityGroupsSpaceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stagingSecurityGroupsSpaceAllColumns,
			stagingSecurityGroupsSpaceColumnsWithDefault,
			stagingSecurityGroupsSpaceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `staging_security_groups_spaces` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `staging_security_groups_spaces` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `staging_security_groups_spaces` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, stagingSecurityGroupsSpacePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into staging_security_groups_spaces")
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

	o.StagingSecurityGroupsSpacesPK = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == stagingSecurityGroupsSpaceMapping["staging_security_groups_spaces_pk"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.StagingSecurityGroupsSpacesPK,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for staging_security_groups_spaces")
	}

CacheNoHooks:
	if !cached {
		stagingSecurityGroupsSpaceInsertCacheMut.Lock()
		stagingSecurityGroupsSpaceInsertCache[key] = cache
		stagingSecurityGroupsSpaceInsertCacheMut.Unlock()
	}

	return nil
}

type StagingSecurityGroupsSpaceUpdater interface {
	Update(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the StagingSecurityGroupsSpace.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q stagingSecurityGroupsSpaceQuery) Update(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	stagingSecurityGroupsSpaceUpdateCacheMut.RLock()
	cache, cached := stagingSecurityGroupsSpaceUpdateCache[key]
	stagingSecurityGroupsSpaceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stagingSecurityGroupsSpaceAllColumns,
			stagingSecurityGroupsSpacePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update staging_security_groups_spaces, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `staging_security_groups_spaces` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, stagingSecurityGroupsSpacePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stagingSecurityGroupsSpaceType, stagingSecurityGroupsSpaceMapping, append(wl, stagingSecurityGroupsSpacePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update staging_security_groups_spaces row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for staging_security_groups_spaces")
	}

	if !cached {
		stagingSecurityGroupsSpaceUpdateCacheMut.Lock()
		stagingSecurityGroupsSpaceUpdateCache[key] = cache
		stagingSecurityGroupsSpaceUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q stagingSecurityGroupsSpaceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for staging_security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for staging_security_groups_spaces")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q stagingSecurityGroupsSpaceQuery) UpdateAllSlice(o StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stagingSecurityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `staging_security_groups_spaces` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stagingSecurityGroupsSpacePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stagingSecurityGroupsSpace slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stagingSecurityGroupsSpace")
	}
	return rowsAff, nil
}

type StagingSecurityGroupsSpaceDeleter interface {
	Delete(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single StagingSecurityGroupsSpace record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q stagingSecurityGroupsSpaceQuery) Delete(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no StagingSecurityGroupsSpace provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stagingSecurityGroupsSpacePrimaryKeyMapping)
	sql := "DELETE FROM `staging_security_groups_spaces` WHERE `staging_security_groups_spaces_pk`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from staging_security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for staging_security_groups_spaces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stagingSecurityGroupsSpaceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stagingSecurityGroupsSpaceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from staging_security_groups_spaces")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for staging_security_groups_spaces")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q stagingSecurityGroupsSpaceQuery) DeleteAllSlice(o StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stagingSecurityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `staging_security_groups_spaces` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stagingSecurityGroupsSpacePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stagingSecurityGroupsSpace slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for staging_security_groups_spaces")
	}

	return rowsAff, nil
}

type StagingSecurityGroupsSpaceReloader interface {
	Reload(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q stagingSecurityGroupsSpaceQuery) Reload(o *StagingSecurityGroupsSpace, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStagingSecurityGroupsSpace(ctx, exec, o.StagingSecurityGroupsSpacesPK)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q stagingSecurityGroupsSpaceQuery) ReloadAll(o *StagingSecurityGroupsSpaceSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StagingSecurityGroupsSpaceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stagingSecurityGroupsSpacePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `staging_security_groups_spaces`.* FROM `staging_security_groups_spaces` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, stagingSecurityGroupsSpacePrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StagingSecurityGroupsSpaceSlice")
	}

	*o = slice

	return nil
}

// StagingSecurityGroupsSpaceExists checks if the StagingSecurityGroupsSpace row exists.
func StagingSecurityGroupsSpaceExists(ctx context.Context, exec boil.ContextExecutor, stagingSecurityGroupsSpacesPK int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `staging_security_groups_spaces` where `staging_security_groups_spaces_pk`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, stagingSecurityGroupsSpacesPK)
	}
	row := exec.QueryRowContext(ctx, sql, stagingSecurityGroupsSpacesPK)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if staging_security_groups_spaces exists")
	}

	return exists, nil
}
