//go:generate mockgen -source=$GOFILE -destination=mocks/buildpack_lifecycle_data.go -copyright_file=../../../../buildtags.txt
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

// BuildpackLifecycleDatum is an object representing the database table.
type BuildpackLifecycleDatum struct {
	ID                        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID                      string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt                 time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt                 null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	AppGUID                   null.String `boil:"app_guid" json:"app_guid,omitempty" toml:"app_guid" yaml:"app_guid,omitempty"`
	DropletGUID               null.String `boil:"droplet_guid" json:"droplet_guid,omitempty" toml:"droplet_guid" yaml:"droplet_guid,omitempty"`
	Stack                     null.String `boil:"stack" json:"stack,omitempty" toml:"stack" yaml:"stack,omitempty"`
	EncryptedBuildpackURL     null.String `boil:"encrypted_buildpack_url" json:"encrypted_buildpack_url,omitempty" toml:"encrypted_buildpack_url" yaml:"encrypted_buildpack_url,omitempty"`
	EncryptedBuildpackURLSalt null.String `boil:"encrypted_buildpack_url_salt" json:"encrypted_buildpack_url_salt,omitempty" toml:"encrypted_buildpack_url_salt" yaml:"encrypted_buildpack_url_salt,omitempty"`
	AdminBuildpackName        null.String `boil:"admin_buildpack_name" json:"admin_buildpack_name,omitempty" toml:"admin_buildpack_name" yaml:"admin_buildpack_name,omitempty"`
	BuildGUID                 null.String `boil:"build_guid" json:"build_guid,omitempty" toml:"build_guid" yaml:"build_guid,omitempty"`
	EncryptionKeyLabel        null.String `boil:"encryption_key_label" json:"encryption_key_label,omitempty" toml:"encryption_key_label" yaml:"encryption_key_label,omitempty"`
	EncryptionIterations      int         `boil:"encryption_iterations" json:"encryption_iterations" toml:"encryption_iterations" yaml:"encryption_iterations"`

	R *buildpackLifecycleDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildpackLifecycleDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildpackLifecycleDatumColumns = struct {
	ID                        string
	GUID                      string
	CreatedAt                 string
	UpdatedAt                 string
	AppGUID                   string
	DropletGUID               string
	Stack                     string
	EncryptedBuildpackURL     string
	EncryptedBuildpackURLSalt string
	AdminBuildpackName        string
	BuildGUID                 string
	EncryptionKeyLabel        string
	EncryptionIterations      string
}{
	ID:                        "id",
	GUID:                      "guid",
	CreatedAt:                 "created_at",
	UpdatedAt:                 "updated_at",
	AppGUID:                   "app_guid",
	DropletGUID:               "droplet_guid",
	Stack:                     "stack",
	EncryptedBuildpackURL:     "encrypted_buildpack_url",
	EncryptedBuildpackURLSalt: "encrypted_buildpack_url_salt",
	AdminBuildpackName:        "admin_buildpack_name",
	BuildGUID:                 "build_guid",
	EncryptionKeyLabel:        "encryption_key_label",
	EncryptionIterations:      "encryption_iterations",
}

var BuildpackLifecycleDatumTableColumns = struct {
	ID                        string
	GUID                      string
	CreatedAt                 string
	UpdatedAt                 string
	AppGUID                   string
	DropletGUID               string
	Stack                     string
	EncryptedBuildpackURL     string
	EncryptedBuildpackURLSalt string
	AdminBuildpackName        string
	BuildGUID                 string
	EncryptionKeyLabel        string
	EncryptionIterations      string
}{
	ID:                        "buildpack_lifecycle_data.id",
	GUID:                      "buildpack_lifecycle_data.guid",
	CreatedAt:                 "buildpack_lifecycle_data.created_at",
	UpdatedAt:                 "buildpack_lifecycle_data.updated_at",
	AppGUID:                   "buildpack_lifecycle_data.app_guid",
	DropletGUID:               "buildpack_lifecycle_data.droplet_guid",
	Stack:                     "buildpack_lifecycle_data.stack",
	EncryptedBuildpackURL:     "buildpack_lifecycle_data.encrypted_buildpack_url",
	EncryptedBuildpackURLSalt: "buildpack_lifecycle_data.encrypted_buildpack_url_salt",
	AdminBuildpackName:        "buildpack_lifecycle_data.admin_buildpack_name",
	BuildGUID:                 "buildpack_lifecycle_data.build_guid",
	EncryptionKeyLabel:        "buildpack_lifecycle_data.encryption_key_label",
	EncryptionIterations:      "buildpack_lifecycle_data.encryption_iterations",
}

// Generated where

var BuildpackLifecycleDatumWhere = struct {
	ID                        whereHelperint
	GUID                      whereHelperstring
	CreatedAt                 whereHelpertime_Time
	UpdatedAt                 whereHelpernull_Time
	AppGUID                   whereHelpernull_String
	DropletGUID               whereHelpernull_String
	Stack                     whereHelpernull_String
	EncryptedBuildpackURL     whereHelpernull_String
	EncryptedBuildpackURLSalt whereHelpernull_String
	AdminBuildpackName        whereHelpernull_String
	BuildGUID                 whereHelpernull_String
	EncryptionKeyLabel        whereHelpernull_String
	EncryptionIterations      whereHelperint
}{
	ID:                        whereHelperint{field: "\"buildpack_lifecycle_data\".\"id\""},
	GUID:                      whereHelperstring{field: "\"buildpack_lifecycle_data\".\"guid\""},
	CreatedAt:                 whereHelpertime_Time{field: "\"buildpack_lifecycle_data\".\"created_at\""},
	UpdatedAt:                 whereHelpernull_Time{field: "\"buildpack_lifecycle_data\".\"updated_at\""},
	AppGUID:                   whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"app_guid\""},
	DropletGUID:               whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"droplet_guid\""},
	Stack:                     whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"stack\""},
	EncryptedBuildpackURL:     whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"encrypted_buildpack_url\""},
	EncryptedBuildpackURLSalt: whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"encrypted_buildpack_url_salt\""},
	AdminBuildpackName:        whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"admin_buildpack_name\""},
	BuildGUID:                 whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"build_guid\""},
	EncryptionKeyLabel:        whereHelpernull_String{field: "\"buildpack_lifecycle_data\".\"encryption_key_label\""},
	EncryptionIterations:      whereHelperint{field: "\"buildpack_lifecycle_data\".\"encryption_iterations\""},
}

// BuildpackLifecycleDatumRels is where relationship names are stored.
var BuildpackLifecycleDatumRels = struct {
	BuildpackLifecycleBuildpacks string
}{
	BuildpackLifecycleBuildpacks: "BuildpackLifecycleBuildpacks",
}

// buildpackLifecycleDatumR is where relationships are stored.
type buildpackLifecycleDatumR struct {
	BuildpackLifecycleBuildpacks BuildpackLifecycleBuildpackSlice `boil:"BuildpackLifecycleBuildpacks" json:"BuildpackLifecycleBuildpacks" toml:"BuildpackLifecycleBuildpacks" yaml:"BuildpackLifecycleBuildpacks"`
}

// NewStruct creates a new relationship struct
func (*buildpackLifecycleDatumR) NewStruct() *buildpackLifecycleDatumR {
	return &buildpackLifecycleDatumR{}
}

// buildpackLifecycleDatumL is where Load methods for each relationship are stored.
type buildpackLifecycleDatumL struct{}

var (
	buildpackLifecycleDatumAllColumns            = []string{"id", "guid", "created_at", "updated_at", "app_guid", "droplet_guid", "stack", "encrypted_buildpack_url", "encrypted_buildpack_url_salt", "admin_buildpack_name", "build_guid", "encryption_key_label", "encryption_iterations"}
	buildpackLifecycleDatumColumnsWithoutDefault = []string{"guid", "updated_at", "app_guid", "droplet_guid", "stack", "encrypted_buildpack_url", "encrypted_buildpack_url_salt", "admin_buildpack_name", "build_guid", "encryption_key_label"}
	buildpackLifecycleDatumColumnsWithDefault    = []string{"id", "created_at", "encryption_iterations"}
	buildpackLifecycleDatumPrimaryKeyColumns     = []string{"id"}
)

type (
	// BuildpackLifecycleDatumSlice is an alias for a slice of pointers to BuildpackLifecycleDatum.
	// This should almost always be used instead of []BuildpackLifecycleDatum.
	BuildpackLifecycleDatumSlice []*BuildpackLifecycleDatum

	buildpackLifecycleDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildpackLifecycleDatumType                 = reflect.TypeOf(&BuildpackLifecycleDatum{})
	buildpackLifecycleDatumMapping              = queries.MakeStructMapping(buildpackLifecycleDatumType)
	buildpackLifecycleDatumPrimaryKeyMapping, _ = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, buildpackLifecycleDatumPrimaryKeyColumns)
	buildpackLifecycleDatumInsertCacheMut       sync.RWMutex
	buildpackLifecycleDatumInsertCache          = make(map[string]insertCache)
	buildpackLifecycleDatumUpdateCacheMut       sync.RWMutex
	buildpackLifecycleDatumUpdateCache          = make(map[string]updateCache)
	buildpackLifecycleDatumUpsertCacheMut       sync.RWMutex
	buildpackLifecycleDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type BuildpackLifecycleDatumFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLifecycleDatum, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLifecycleDatumSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single buildpackLifecycleDatum record from the query.
func (q buildpackLifecycleDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BuildpackLifecycleDatum, error) {
	o := &BuildpackLifecycleDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buildpack_lifecycle_data")
	}

	return o, nil
}

// All returns all BuildpackLifecycleDatum records from the query.
func (q buildpackLifecycleDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildpackLifecycleDatumSlice, error) {
	var o []*BuildpackLifecycleDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BuildpackLifecycleDatum slice")
	}

	return o, nil
}

// Count returns the count of all BuildpackLifecycleDatum records in the query.
func (q buildpackLifecycleDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buildpack_lifecycle_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q buildpackLifecycleDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buildpack_lifecycle_data exists")
	}

	return count > 0, nil
}

// BuildpackLifecycleBuildpacks retrieves all the buildpack_lifecycle_buildpack's BuildpackLifecycleBuildpacks with an executor.
func (q buildpackLifecycleDatumQuery) BuildpackLifecycleBuildpacks(o *BuildpackLifecycleDatum, mods ...qm.QueryMod) buildpackLifecycleBuildpackQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"buildpack_lifecycle_buildpacks\".\"buildpack_lifecycle_data_guid\"=?", o.GUID),
	)

	query := BuildpackLifecycleBuildpacks(queryMods...)
	queries.SetFrom(query.Query, "\"buildpack_lifecycle_buildpacks\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"buildpack_lifecycle_buildpacks\".*"})
	}

	return query
}

// LoadBuildpackLifecycleBuildpacks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (buildpackLifecycleDatumL) LoadBuildpackLifecycleBuildpacks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildpackLifecycleDatum interface{}, mods queries.Applicator) error {
	var slice []*BuildpackLifecycleDatum
	var object *BuildpackLifecycleDatum

	if singular {
		object = maybeBuildpackLifecycleDatum.(*BuildpackLifecycleDatum)
	} else {
		slice = *maybeBuildpackLifecycleDatum.(*[]*BuildpackLifecycleDatum)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildpackLifecycleDatumR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildpackLifecycleDatumR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.GUID) {
					continue Outer
				}
			}

			args = append(args, obj.GUID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`buildpack_lifecycle_buildpacks`),
		qm.WhereIn(`buildpack_lifecycle_buildpacks.buildpack_lifecycle_data_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load buildpack_lifecycle_buildpacks")
	}

	var resultSlice []*BuildpackLifecycleBuildpack
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice buildpack_lifecycle_buildpacks")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on buildpack_lifecycle_buildpacks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for buildpack_lifecycle_buildpacks")
	}

	if singular {
		object.R.BuildpackLifecycleBuildpacks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &buildpackLifecycleBuildpackR{}
			}
			foreign.R.BuildpackLifecycleDatum = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.GUID, foreign.BuildpackLifecycleDataGUID) {
				local.R.BuildpackLifecycleBuildpacks = append(local.R.BuildpackLifecycleBuildpacks, foreign)
				if foreign.R == nil {
					foreign.R = &buildpackLifecycleBuildpackR{}
				}
				foreign.R.BuildpackLifecycleDatum = local
				break
			}
		}
	}

	return nil
}

// AddBuildpackLifecycleBuildpacks adds the given related objects to the existing relationships
// of the buildpack_lifecycle_datum, optionally inserting them as new records.
// Appends related to o.R.BuildpackLifecycleBuildpacks.
// Sets related.R.BuildpackLifecycleDatum appropriately.
func (q buildpackLifecycleDatumQuery) AddBuildpackLifecycleBuildpacks(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackLifecycleBuildpack) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.BuildpackLifecycleDataGUID, o.GUID)
			if err = BuildpackLifecycleBuildpacks().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"buildpack_lifecycle_buildpacks\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"buildpack_lifecycle_data_guid"}),
				strmangle.WhereClause("\"", "\"", 2, buildpackLifecycleBuildpackPrimaryKeyColumns),
			)
			values := []interface{}{o.GUID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.BuildpackLifecycleDataGUID, o.GUID)
		}
	}

	if o.R == nil {
		o.R = &buildpackLifecycleDatumR{
			BuildpackLifecycleBuildpacks: related,
		}
	} else {
		o.R.BuildpackLifecycleBuildpacks = append(o.R.BuildpackLifecycleBuildpacks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &buildpackLifecycleBuildpackR{
				BuildpackLifecycleDatum: o,
			}
		} else {
			rel.R.BuildpackLifecycleDatum = o
		}
	}
	return nil
}

// SetBuildpackLifecycleBuildpacks removes all previously related items of the
// buildpack_lifecycle_datum replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.BuildpackLifecycleDatum's BuildpackLifecycleBuildpacks accordingly.
// Replaces o.R.BuildpackLifecycleBuildpacks with related.
// Sets related.R.BuildpackLifecycleDatum's BuildpackLifecycleBuildpacks accordingly.
func (q buildpackLifecycleDatumQuery) SetBuildpackLifecycleBuildpacks(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackLifecycleBuildpack) error {
	query := "update \"buildpack_lifecycle_buildpacks\" set \"buildpack_lifecycle_data_guid\" = null where \"buildpack_lifecycle_data_guid\" = $1"
	values := []interface{}{o.GUID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.BuildpackLifecycleBuildpacks {
			queries.SetScanner(&rel.BuildpackLifecycleDataGUID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.BuildpackLifecycleDatum = nil
		}

		o.R.BuildpackLifecycleBuildpacks = nil
	}
	return q.AddBuildpackLifecycleBuildpacks(o, ctx, exec, insert, related...)
}

// RemoveBuildpackLifecycleBuildpacks relationships from objects passed in.
// Removes related items from R.BuildpackLifecycleBuildpacks (uses pointer comparison, removal does not keep order)
// Sets related.R.BuildpackLifecycleDatum.
func (q buildpackLifecycleDatumQuery) RemoveBuildpackLifecycleBuildpacks(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, related ...*BuildpackLifecycleBuildpack) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.BuildpackLifecycleDataGUID, nil)
		if rel.R != nil {
			rel.R.BuildpackLifecycleDatum = nil
		}
		if _, err = BuildpackLifecycleBuildpacks().Update(rel, ctx, exec, boil.Whitelist("buildpack_lifecycle_data_guid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.BuildpackLifecycleBuildpacks {
			if rel != ri {
				continue
			}

			ln := len(o.R.BuildpackLifecycleBuildpacks)
			if ln > 1 && i < ln-1 {
				o.R.BuildpackLifecycleBuildpacks[i] = o.R.BuildpackLifecycleBuildpacks[ln-1]
			}
			o.R.BuildpackLifecycleBuildpacks = o.R.BuildpackLifecycleBuildpacks[:ln-1]
			break
		}
	}

	return nil
}

// BuildpackLifecycleData retrieves all the records using an executor.
func BuildpackLifecycleData(mods ...qm.QueryMod) buildpackLifecycleDatumQuery {
	mods = append(mods, qm.From("\"buildpack_lifecycle_data\""))
	return buildpackLifecycleDatumQuery{NewQuery(mods...)}
}

type BuildpackLifecycleDatumFinder interface {
	FindBuildpackLifecycleDatum(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLifecycleDatum, error)
}

// FindBuildpackLifecycleDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuildpackLifecycleDatum(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*BuildpackLifecycleDatum, error) {
	buildpackLifecycleDatumObj := &BuildpackLifecycleDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"buildpack_lifecycle_data\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildpackLifecycleDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buildpack_lifecycle_data")
	}

	return buildpackLifecycleDatumObj, nil
}

type BuildpackLifecycleDatumInserter interface {
	Insert(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q buildpackLifecycleDatumQuery) Insert(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_lifecycle_data provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(buildpackLifecycleDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildpackLifecycleDatumInsertCacheMut.RLock()
	cache, cached := buildpackLifecycleDatumInsertCache[key]
	buildpackLifecycleDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildpackLifecycleDatumAllColumns,
			buildpackLifecycleDatumColumnsWithDefault,
			buildpackLifecycleDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"buildpack_lifecycle_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"buildpack_lifecycle_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into buildpack_lifecycle_data")
	}

	if !cached {
		buildpackLifecycleDatumInsertCacheMut.Lock()
		buildpackLifecycleDatumInsertCache[key] = cache
		buildpackLifecycleDatumInsertCacheMut.Unlock()
	}

	return nil
}

type BuildpackLifecycleDatumUpdater interface {
	Update(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the BuildpackLifecycleDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q buildpackLifecycleDatumQuery) Update(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	buildpackLifecycleDatumUpdateCacheMut.RLock()
	cache, cached := buildpackLifecycleDatumUpdateCache[key]
	buildpackLifecycleDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildpackLifecycleDatumAllColumns,
			buildpackLifecycleDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buildpack_lifecycle_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"buildpack_lifecycle_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, buildpackLifecycleDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, append(wl, buildpackLifecycleDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update buildpack_lifecycle_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buildpack_lifecycle_data")
	}

	if !cached {
		buildpackLifecycleDatumUpdateCacheMut.Lock()
		buildpackLifecycleDatumUpdateCache[key] = cache
		buildpackLifecycleDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q buildpackLifecycleDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buildpack_lifecycle_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buildpack_lifecycle_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q buildpackLifecycleDatumQuery) UpdateAllSlice(o BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"buildpack_lifecycle_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, buildpackLifecycleDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in buildpackLifecycleDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all buildpackLifecycleDatum")
	}
	return rowsAff, nil
}

type BuildpackLifecycleDatumUpserter interface {
	Upsert(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q buildpackLifecycleDatumQuery) Upsert(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpack_lifecycle_data provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(buildpackLifecycleDatumColumnsWithDefault, o)

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

	buildpackLifecycleDatumUpsertCacheMut.RLock()
	cache, cached := buildpackLifecycleDatumUpsertCache[key]
	buildpackLifecycleDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			buildpackLifecycleDatumAllColumns,
			buildpackLifecycleDatumColumnsWithDefault,
			buildpackLifecycleDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			buildpackLifecycleDatumAllColumns,
			buildpackLifecycleDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert buildpack_lifecycle_data, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(buildpackLifecycleDatumPrimaryKeyColumns))
			copy(conflict, buildpackLifecycleDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"buildpack_lifecycle_data\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildpackLifecycleDatumType, buildpackLifecycleDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert buildpack_lifecycle_data")
	}

	if !cached {
		buildpackLifecycleDatumUpsertCacheMut.Lock()
		buildpackLifecycleDatumUpsertCache[key] = cache
		buildpackLifecycleDatumUpsertCacheMut.Unlock()
	}

	return nil
}

type BuildpackLifecycleDatumDeleter interface {
	Delete(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single BuildpackLifecycleDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q buildpackLifecycleDatumQuery) Delete(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BuildpackLifecycleDatum provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildpackLifecycleDatumPrimaryKeyMapping)
	sql := "DELETE FROM \"buildpack_lifecycle_data\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buildpack_lifecycle_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buildpack_lifecycle_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q buildpackLifecycleDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildpackLifecycleDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpack_lifecycle_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_lifecycle_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q buildpackLifecycleDatumQuery) DeleteAllSlice(o BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"buildpack_lifecycle_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackLifecycleDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpackLifecycleDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpack_lifecycle_data")
	}

	return rowsAff, nil
}

type BuildpackLifecycleDatumReloader interface {
	Reload(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q buildpackLifecycleDatumQuery) Reload(o *BuildpackLifecycleDatum, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuildpackLifecycleDatum(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q buildpackLifecycleDatumQuery) ReloadAll(o *BuildpackLifecycleDatumSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildpackLifecycleDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackLifecycleDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"buildpack_lifecycle_data\".* FROM \"buildpack_lifecycle_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackLifecycleDatumPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildpackLifecycleDatumSlice")
	}

	*o = slice

	return nil
}

// BuildpackLifecycleDatumExists checks if the BuildpackLifecycleDatum row exists.
func BuildpackLifecycleDatumExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"buildpack_lifecycle_data\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buildpack_lifecycle_data exists")
	}

	return exists, nil
}
