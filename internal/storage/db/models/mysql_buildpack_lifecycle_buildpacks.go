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

type BuildpackLifecycleBuildpackUpserter interface {
	Upsert(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLBuildpackLifecycleBuildpackUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q buildpackLifecycleBuildpackQuery) Upsert(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_lifecycle_buildpacks provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(buildpackLifecycleBuildpackColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLBuildpackLifecycleBuildpackUniqueColumns, o)

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

	buildpackLifecycleBuildpackUpsertCacheMut.RLock()
	cache, cached := buildpackLifecycleBuildpackUpsertCache[key]
	buildpackLifecycleBuildpackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			buildpackLifecycleBuildpackAllColumns,
			buildpackLifecycleBuildpackColumnsWithDefault,
			buildpackLifecycleBuildpackColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			buildpackLifecycleBuildpackAllColumns,
			buildpackLifecycleBuildpackPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert buildpack_lifecycle_buildpacks, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`buildpack_lifecycle_buildpacks`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `buildpack_lifecycle_buildpacks` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for buildpack_lifecycle_buildpacks")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == buildpackLifecycleBuildpackMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for buildpack_lifecycle_buildpacks")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for buildpack_lifecycle_buildpacks")
	}

CacheNoHooks:
	if !cached {
		buildpackLifecycleBuildpackUpsertCacheMut.Lock()
		buildpackLifecycleBuildpackUpsertCache[key] = cache
		buildpackLifecycleBuildpackUpsertCacheMut.Unlock()
	}

	return nil
}

// BuildpackLifecycleBuildpack is an object representing the database table.
type BuildpackLifecycleBuildpack struct {
	ID                         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                       string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt                  time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	AdminBuildpackName         null.String `boil:"admin_buildpack_name" json:"admin_buildpack_name,omitempty" toml:"admin_buildpack_name" yaml:"admin_buildpack_name,omitempty"`
	EncryptedBuildpackURL      null.String `boil:"encrypted_buildpack_url" json:"encrypted_buildpack_url,omitempty" toml:"encrypted_buildpack_url" yaml:"encrypted_buildpack_url,omitempty"`
	EncryptedBuildpackURLSalt  null.String `boil:"encrypted_buildpack_url_salt" json:"encrypted_buildpack_url_salt,omitempty" toml:"encrypted_buildpack_url_salt" yaml:"encrypted_buildpack_url_salt,omitempty"`
	BuildpackLifecycleDataGUID null.String `boil:"buildpack_lifecycle_data_guid" json:"buildpack_lifecycle_data_guid,omitempty" toml:"buildpack_lifecycle_data_guid" yaml:"buildpack_lifecycle_data_guid,omitempty"`
	EncryptionKeyLabel         null.String `boil:"encryption_key_label" json:"encryption_key_label,omitempty" toml:"encryption_key_label" yaml:"encryption_key_label,omitempty"`
	Version                    null.String `boil:"version" json:"version,omitempty" toml:"version" yaml:"version,omitempty"`
	BuildpackName              null.String `boil:"buildpack_name" json:"buildpack_name,omitempty" toml:"buildpack_name" yaml:"buildpack_name,omitempty"`
	EncryptionIterations       int         `boil:"encryption_iterations" json:"encryption_iterations" toml:"encryption_iterations" yaml:"encryption_iterations"`

	R *buildpackLifecycleBuildpackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildpackLifecycleBuildpackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildpackLifecycleBuildpackColumns = struct {
	ID                         string
	GUID                       string
	CreatedAt                  string
	UpdatedAt                  string
	AdminBuildpackName         string
	EncryptedBuildpackURL      string
	EncryptedBuildpackURLSalt  string
	BuildpackLifecycleDataGUID string
	EncryptionKeyLabel         string
	Version                    string
	BuildpackName              string
	EncryptionIterations       string
}{
	ID:                         "id",
	GUID:                       "guid",
	CreatedAt:                  "created_at",
	UpdatedAt:                  "updated_at",
	AdminBuildpackName:         "admin_buildpack_name",
	EncryptedBuildpackURL:      "encrypted_buildpack_url",
	EncryptedBuildpackURLSalt:  "encrypted_buildpack_url_salt",
	BuildpackLifecycleDataGUID: "buildpack_lifecycle_data_guid",
	EncryptionKeyLabel:         "encryption_key_label",
	Version:                    "version",
	BuildpackName:              "buildpack_name",
	EncryptionIterations:       "encryption_iterations",
}

var BuildpackLifecycleBuildpackTableColumns = struct {
	ID                         string
	GUID                       string
	CreatedAt                  string
	UpdatedAt                  string
	AdminBuildpackName         string
	EncryptedBuildpackURL      string
	EncryptedBuildpackURLSalt  string
	BuildpackLifecycleDataGUID string
	EncryptionKeyLabel         string
	Version                    string
	BuildpackName              string
	EncryptionIterations       string
}{
	ID:                         "buildpack_lifecycle_buildpacks.id",
	GUID:                       "buildpack_lifecycle_buildpacks.guid",
	CreatedAt:                  "buildpack_lifecycle_buildpacks.created_at",
	UpdatedAt:                  "buildpack_lifecycle_buildpacks.updated_at",
	AdminBuildpackName:         "buildpack_lifecycle_buildpacks.admin_buildpack_name",
	EncryptedBuildpackURL:      "buildpack_lifecycle_buildpacks.encrypted_buildpack_url",
	EncryptedBuildpackURLSalt:  "buildpack_lifecycle_buildpacks.encrypted_buildpack_url_salt",
	BuildpackLifecycleDataGUID: "buildpack_lifecycle_buildpacks.buildpack_lifecycle_data_guid",
	EncryptionKeyLabel:         "buildpack_lifecycle_buildpacks.encryption_key_label",
	Version:                    "buildpack_lifecycle_buildpacks.version",
	BuildpackName:              "buildpack_lifecycle_buildpacks.buildpack_name",
	EncryptionIterations:       "buildpack_lifecycle_buildpacks.encryption_iterations",
}

// Generated where

var BuildpackLifecycleBuildpackWhere = struct {
	ID                         whereHelperint
	GUID                       whereHelperstring
	CreatedAt                  whereHelpertime_Time
	UpdatedAt                  whereHelpernull_Time
	AdminBuildpackName         whereHelpernull_String
	EncryptedBuildpackURL      whereHelpernull_String
	EncryptedBuildpackURLSalt  whereHelpernull_String
	BuildpackLifecycleDataGUID whereHelpernull_String
	EncryptionKeyLabel         whereHelpernull_String
	Version                    whereHelpernull_String
	BuildpackName              whereHelpernull_String
	EncryptionIterations       whereHelperint
}{
	ID:                         whereHelperint{field: "`buildpack_lifecycle_buildpacks`.`id`"},
	GUID:                       whereHelperstring{field: "`buildpack_lifecycle_buildpacks`.`guid`"},
	CreatedAt:                  whereHelpertime_Time{field: "`buildpack_lifecycle_buildpacks`.`created_at`"},
	UpdatedAt:                  whereHelpernull_Time{field: "`buildpack_lifecycle_buildpacks`.`updated_at`"},
	AdminBuildpackName:         whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`admin_buildpack_name`"},
	EncryptedBuildpackURL:      whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`encrypted_buildpack_url`"},
	EncryptedBuildpackURLSalt:  whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`encrypted_buildpack_url_salt`"},
	BuildpackLifecycleDataGUID: whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`buildpack_lifecycle_data_guid`"},
	EncryptionKeyLabel:         whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`encryption_key_label`"},
	Version:                    whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`version`"},
	BuildpackName:              whereHelpernull_String{field: "`buildpack_lifecycle_buildpacks`.`buildpack_name`"},
	EncryptionIterations:       whereHelperint{field: "`buildpack_lifecycle_buildpacks`.`encryption_iterations`"},
}

// BuildpackLifecycleBuildpackRels is where relationship names are stored.
var BuildpackLifecycleBuildpackRels = struct {
	BuildpackLifecycleDatum string
}{
	BuildpackLifecycleDatum: "BuildpackLifecycleDatum",
}

// buildpackLifecycleBuildpackR is where relationships are stored.
type buildpackLifecycleBuildpackR struct {
	BuildpackLifecycleDatum *BuildpackLifecycleDatum `boil:"BuildpackLifecycleDatum" json:"BuildpackLifecycleDatum" toml:"BuildpackLifecycleDatum" yaml:"BuildpackLifecycleDatum"`
}

// NewStruct creates a new relationship struct
func (*buildpackLifecycleBuildpackR) NewStruct() *buildpackLifecycleBuildpackR {
	return &buildpackLifecycleBuildpackR{}
}

// buildpackLifecycleBuildpackL is where Load methods for each relationship are stored.
type buildpackLifecycleBuildpackL struct{}

var (
	buildpackLifecycleBuildpackAllColumns            = []string{"id", "guid", "created_at", "updated_at", "admin_buildpack_name", "encrypted_buildpack_url", "encrypted_buildpack_url_salt", "buildpack_lifecycle_data_guid", "encryption_key_label", "version", "buildpack_name", "encryption_iterations"}
	buildpackLifecycleBuildpackColumnsWithoutDefault = []string{"guid", "updated_at", "admin_buildpack_name", "encrypted_buildpack_url", "encrypted_buildpack_url_salt", "buildpack_lifecycle_data_guid", "encryption_key_label", "version", "buildpack_name"}
	buildpackLifecycleBuildpackColumnsWithDefault    = []string{"id", "created_at", "encryption_iterations"}
	buildpackLifecycleBuildpackPrimaryKeyColumns     = []string{"id"}
)

type (
	// BuildpackLifecycleBuildpackSlice is an alias for a slice of pointers to BuildpackLifecycleBuildpack.
	// This should almost always be used instead of []BuildpackLifecycleBuildpack.
	BuildpackLifecycleBuildpackSlice []*BuildpackLifecycleBuildpack

	buildpackLifecycleBuildpackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildpackLifecycleBuildpackType                 = reflect.TypeOf(&BuildpackLifecycleBuildpack{})
	buildpackLifecycleBuildpackMapping              = queries.MakeStructMapping(buildpackLifecycleBuildpackType)
	buildpackLifecycleBuildpackPrimaryKeyMapping, _ = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, buildpackLifecycleBuildpackPrimaryKeyColumns)
	buildpackLifecycleBuildpackInsertCacheMut       sync.RWMutex
	buildpackLifecycleBuildpackInsertCache          = make(map[string]insertCache)
	buildpackLifecycleBuildpackUpdateCacheMut       sync.RWMutex
	buildpackLifecycleBuildpackUpdateCache          = make(map[string]updateCache)
	buildpackLifecycleBuildpackUpsertCacheMut       sync.RWMutex
	buildpackLifecycleBuildpackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type BuildpackLifecycleBuildpackFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLifecycleBuildpack, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLifecycleBuildpackSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single buildpackLifecycleBuildpack record from the query.
func (q buildpackLifecycleBuildpackQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLifecycleBuildpack, error) {
	o := &BuildpackLifecycleBuildpack{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buildpack_lifecycle_buildpacks")
	}

	return o, nil
}

// All returns all BuildpackLifecycleBuildpack records from the query.
func (q buildpackLifecycleBuildpackQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLifecycleBuildpackSlice, error) {
	var o []*BuildpackLifecycleBuildpack

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BuildpackLifecycleBuildpack slice")
	}

	return o, nil
}

// Count returns the count of all BuildpackLifecycleBuildpack records in the query.
func (q buildpackLifecycleBuildpackQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buildpack_lifecycle_buildpacks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q buildpackLifecycleBuildpackQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buildpack_lifecycle_buildpacks exists")
	}

	return count > 0, nil
}

// BuildpackLifecycleDatum pointed to by the foreign key.
func (q buildpackLifecycleBuildpackQuery) BuildpackLifecycleDatum(o *BuildpackLifecycleBuildpack, mods ...qm.QueryMod) buildpackLifecycleDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.BuildpackLifecycleDataGUID),
	}

	queryMods = append(queryMods, mods...)

	query := BuildpackLifecycleData(queryMods...)
	queries.SetFrom(query.Query, "`buildpack_lifecycle_data`")

	return query
}

// LoadBuildpackLifecycleDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (buildpackLifecycleBuildpackL) LoadBuildpackLifecycleDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildpackLifecycleBuildpack interface{}, mods queries.Applicator) error {
	var slice []*BuildpackLifecycleBuildpack
	var object *BuildpackLifecycleBuildpack

	if singular {
		object = maybeBuildpackLifecycleBuildpack.(*BuildpackLifecycleBuildpack)
	} else {
		slice = *maybeBuildpackLifecycleBuildpack.(*[]*BuildpackLifecycleBuildpack)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildpackLifecycleBuildpackR{}
		}
		if !queries.IsNil(object.BuildpackLifecycleDataGUID) {
			args = append(args, object.BuildpackLifecycleDataGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildpackLifecycleBuildpackR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BuildpackLifecycleDataGUID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BuildpackLifecycleDataGUID) {
				args = append(args, obj.BuildpackLifecycleDataGUID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`buildpack_lifecycle_data`),
		qm.WhereIn(`buildpack_lifecycle_data.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load BuildpackLifecycleDatum")
	}

	var resultSlice []*BuildpackLifecycleDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice BuildpackLifecycleDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for buildpack_lifecycle_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for buildpack_lifecycle_data")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.BuildpackLifecycleDatum = foreign
		if foreign.R == nil {
			foreign.R = &buildpackLifecycleDatumR{}
		}
		foreign.R.BuildpackLifecycleBuildpacks = append(foreign.R.BuildpackLifecycleBuildpacks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BuildpackLifecycleDataGUID, foreign.GUID) {
				local.R.BuildpackLifecycleDatum = foreign
				if foreign.R == nil {
					foreign.R = &buildpackLifecycleDatumR{}
				}
				foreign.R.BuildpackLifecycleBuildpacks = append(foreign.R.BuildpackLifecycleBuildpacks, local)
				break
			}
		}
	}

	return nil
}

// SetBuildpackLifecycleDatum of the buildpackLifecycleBuildpack to the related item.
// Sets o.R.BuildpackLifecycleDatum to related.
// Adds o to related.R.BuildpackLifecycleBuildpacks.
func (q buildpackLifecycleBuildpackQuery) SetBuildpackLifecycleDatum(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, insert bool, related *BuildpackLifecycleDatum) error {
	var err error
	if insert {
		if err = BuildpackLifecycleData().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `buildpack_lifecycle_buildpacks` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"buildpack_lifecycle_data_guid"}),
		strmangle.WhereClause("`", "`", 0, buildpackLifecycleBuildpackPrimaryKeyColumns),
	)
	values := []interface{}{related.GUID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.BuildpackLifecycleDataGUID, related.GUID)
	if o.R == nil {
		o.R = &buildpackLifecycleBuildpackR{
			BuildpackLifecycleDatum: related,
		}
	} else {
		o.R.BuildpackLifecycleDatum = related
	}

	if related.R == nil {
		related.R = &buildpackLifecycleDatumR{
			BuildpackLifecycleBuildpacks: BuildpackLifecycleBuildpackSlice{o},
		}
	} else {
		related.R.BuildpackLifecycleBuildpacks = append(related.R.BuildpackLifecycleBuildpacks, o)
	}

	return nil
}

// RemoveBuildpackLifecycleDatum relationship.
// Sets o.R.BuildpackLifecycleDatum to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q buildpackLifecycleBuildpackQuery) RemoveBuildpackLifecycleDatum(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, related *BuildpackLifecycleDatum) error {
	var err error

	queries.SetScanner(&o.BuildpackLifecycleDataGUID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("buildpack_lifecycle_data_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.BuildpackLifecycleDatum = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.BuildpackLifecycleBuildpacks {
		if queries.Equal(o.BuildpackLifecycleDataGUID, ri.BuildpackLifecycleDataGUID) {
			continue
		}

		ln := len(related.R.BuildpackLifecycleBuildpacks)
		if ln > 1 && i < ln-1 {
			related.R.BuildpackLifecycleBuildpacks[i] = related.R.BuildpackLifecycleBuildpacks[ln-1]
		}
		related.R.BuildpackLifecycleBuildpacks = related.R.BuildpackLifecycleBuildpacks[:ln-1]
		break
	}
	return nil
}

// BuildpackLifecycleBuildpacks retrieves all the records using an executor.
func BuildpackLifecycleBuildpacks(mods ...qm.QueryMod) buildpackLifecycleBuildpackQuery {
	mods = append(mods, qm.From("`buildpack_lifecycle_buildpacks`"))
	return buildpackLifecycleBuildpackQuery{NewQuery(mods...)}
}

type BuildpackLifecycleBuildpackFinder interface {
	FindBuildpackLifecycleBuildpack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLifecycleBuildpack, error)
}

// FindBuildpackLifecycleBuildpack retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuildpackLifecycleBuildpack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLifecycleBuildpack, error) {
	buildpackLifecycleBuildpackObj := &BuildpackLifecycleBuildpack{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `buildpack_lifecycle_buildpacks` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildpackLifecycleBuildpackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buildpack_lifecycle_buildpacks")
	}

	return buildpackLifecycleBuildpackObj, nil
}

type BuildpackLifecycleBuildpackInserter interface {
	Insert(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q buildpackLifecycleBuildpackQuery) Insert(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_lifecycle_buildpacks provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(buildpackLifecycleBuildpackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildpackLifecycleBuildpackInsertCacheMut.RLock()
	cache, cached := buildpackLifecycleBuildpackInsertCache[key]
	buildpackLifecycleBuildpackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildpackLifecycleBuildpackAllColumns,
			buildpackLifecycleBuildpackColumnsWithDefault,
			buildpackLifecycleBuildpackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `buildpack_lifecycle_buildpacks` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `buildpack_lifecycle_buildpacks` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `buildpack_lifecycle_buildpacks` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, buildpackLifecycleBuildpackPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into buildpack_lifecycle_buildpacks")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == buildpackLifecycleBuildpackMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for buildpack_lifecycle_buildpacks")
	}

CacheNoHooks:
	if !cached {
		buildpackLifecycleBuildpackInsertCacheMut.Lock()
		buildpackLifecycleBuildpackInsertCache[key] = cache
		buildpackLifecycleBuildpackInsertCacheMut.Unlock()
	}

	return nil
}

type BuildpackLifecycleBuildpackUpdater interface {
	Update(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the BuildpackLifecycleBuildpack.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q buildpackLifecycleBuildpackQuery) Update(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	buildpackLifecycleBuildpackUpdateCacheMut.RLock()
	cache, cached := buildpackLifecycleBuildpackUpdateCache[key]
	buildpackLifecycleBuildpackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildpackLifecycleBuildpackAllColumns,
			buildpackLifecycleBuildpackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buildpack_lifecycle_buildpacks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `buildpack_lifecycle_buildpacks` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, buildpackLifecycleBuildpackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleBuildpackType, buildpackLifecycleBuildpackMapping, append(wl, buildpackLifecycleBuildpackPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update buildpack_lifecycle_buildpacks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buildpack_lifecycle_buildpacks")
	}

	if !cached {
		buildpackLifecycleBuildpackUpdateCacheMut.Lock()
		buildpackLifecycleBuildpackUpdateCache[key] = cache
		buildpackLifecycleBuildpackUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q buildpackLifecycleBuildpackQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buildpack_lifecycle_buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buildpack_lifecycle_buildpacks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q buildpackLifecycleBuildpackQuery) UpdateAllSlice(o BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleBuildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `buildpack_lifecycle_buildpacks` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildpackLifecycleBuildpackPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in buildpackLifecycleBuildpack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all buildpackLifecycleBuildpack")
	}
	return rowsAff, nil
}

type BuildpackLifecycleBuildpackDeleter interface {
	Delete(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single BuildpackLifecycleBuildpack record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q buildpackLifecycleBuildpackQuery) Delete(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BuildpackLifecycleBuildpack provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildpackLifecycleBuildpackPrimaryKeyMapping)
	sql := "DELETE FROM `buildpack_lifecycle_buildpacks` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buildpack_lifecycle_buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buildpack_lifecycle_buildpacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q buildpackLifecycleBuildpackQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildpackLifecycleBuildpackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpack_lifecycle_buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_lifecycle_buildpacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q buildpackLifecycleBuildpackQuery) DeleteAllSlice(o BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleBuildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `buildpack_lifecycle_buildpacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildpackLifecycleBuildpackPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpackLifecycleBuildpack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_lifecycle_buildpacks")
	}

	return rowsAff, nil
}

type BuildpackLifecycleBuildpackReloader interface {
	Reload(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q buildpackLifecycleBuildpackQuery) Reload(o *BuildpackLifecycleBuildpack, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuildpackLifecycleBuildpack(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q buildpackLifecycleBuildpackQuery) ReloadAll(o *BuildpackLifecycleBuildpackSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildpackLifecycleBuildpackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleBuildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `buildpack_lifecycle_buildpacks`.* FROM `buildpack_lifecycle_buildpacks` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, buildpackLifecycleBuildpackPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildpackLifecycleBuildpackSlice")
	}

	*o = slice

	return nil
}

// BuildpackLifecycleBuildpackExists checks if the BuildpackLifecycleBuildpack row exists.
func BuildpackLifecycleBuildpackExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `buildpack_lifecycle_buildpacks` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buildpack_lifecycle_buildpacks exists")
	}

	return exists, nil
}