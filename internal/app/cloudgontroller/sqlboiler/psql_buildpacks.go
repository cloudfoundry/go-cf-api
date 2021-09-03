//go:generate mockgen -source=$GOFILE -destination=mocks/buildpacks.go
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

// Buildpack is an object representing the database table.
type Buildpack struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID           string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt      time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Name           string      `boil:"name" json:"name" toml:"name" yaml:"name"`
	Key            null.String `boil:"key" json:"key,omitempty" toml:"key" yaml:"key,omitempty"`
	Position       int         `boil:"position" json:"position" toml:"position" yaml:"position"`
	Enabled        null.Bool   `boil:"enabled" json:"enabled,omitempty" toml:"enabled" yaml:"enabled,omitempty"`
	Locked         null.Bool   `boil:"locked" json:"locked,omitempty" toml:"locked" yaml:"locked,omitempty"`
	Filename       null.String `boil:"filename" json:"filename,omitempty" toml:"filename" yaml:"filename,omitempty"`
	Sha256Checksum null.String `boil:"sha256_checksum" json:"sha256_checksum,omitempty" toml:"sha256_checksum" yaml:"sha256_checksum,omitempty"`
	Stack          null.String `boil:"stack" json:"stack,omitempty" toml:"stack" yaml:"stack,omitempty"`

	R *buildpackR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L buildpackL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BuildpackColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	Name           string
	Key            string
	Position       string
	Enabled        string
	Locked         string
	Filename       string
	Sha256Checksum string
	Stack          string
}{
	ID:             "id",
	GUID:           "guid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Name:           "name",
	Key:            "key",
	Position:       "position",
	Enabled:        "enabled",
	Locked:         "locked",
	Filename:       "filename",
	Sha256Checksum: "sha256_checksum",
	Stack:          "stack",
}

var BuildpackTableColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	Name           string
	Key            string
	Position       string
	Enabled        string
	Locked         string
	Filename       string
	Sha256Checksum string
	Stack          string
}{
	ID:             "buildpacks.id",
	GUID:           "buildpacks.guid",
	CreatedAt:      "buildpacks.created_at",
	UpdatedAt:      "buildpacks.updated_at",
	Name:           "buildpacks.name",
	Key:            "buildpacks.key",
	Position:       "buildpacks.position",
	Enabled:        "buildpacks.enabled",
	Locked:         "buildpacks.locked",
	Filename:       "buildpacks.filename",
	Sha256Checksum: "buildpacks.sha256_checksum",
	Stack:          "buildpacks.stack",
}

// Generated where

var BuildpackWhere = struct {
	ID             whereHelperint
	GUID           whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpernull_Time
	Name           whereHelperstring
	Key            whereHelpernull_String
	Position       whereHelperint
	Enabled        whereHelpernull_Bool
	Locked         whereHelpernull_Bool
	Filename       whereHelpernull_String
	Sha256Checksum whereHelpernull_String
	Stack          whereHelpernull_String
}{
	ID:             whereHelperint{field: "\"buildpacks\".\"id\""},
	GUID:           whereHelperstring{field: "\"buildpacks\".\"guid\""},
	CreatedAt:      whereHelpertime_Time{field: "\"buildpacks\".\"created_at\""},
	UpdatedAt:      whereHelpernull_Time{field: "\"buildpacks\".\"updated_at\""},
	Name:           whereHelperstring{field: "\"buildpacks\".\"name\""},
	Key:            whereHelpernull_String{field: "\"buildpacks\".\"key\""},
	Position:       whereHelperint{field: "\"buildpacks\".\"position\""},
	Enabled:        whereHelpernull_Bool{field: "\"buildpacks\".\"enabled\""},
	Locked:         whereHelpernull_Bool{field: "\"buildpacks\".\"locked\""},
	Filename:       whereHelpernull_String{field: "\"buildpacks\".\"filename\""},
	Sha256Checksum: whereHelpernull_String{field: "\"buildpacks\".\"sha256_checksum\""},
	Stack:          whereHelpernull_String{field: "\"buildpacks\".\"stack\""},
}

// BuildpackRels is where relationship names are stored.
var BuildpackRels = struct {
	ResourceBuildpackAnnotations string
	ResourceBuildpackLabels      string
}{
	ResourceBuildpackAnnotations: "ResourceBuildpackAnnotations",
	ResourceBuildpackLabels:      "ResourceBuildpackLabels",
}

// buildpackR is where relationships are stored.
type buildpackR struct {
	ResourceBuildpackAnnotations BuildpackAnnotationSlice `boil:"ResourceBuildpackAnnotations" json:"ResourceBuildpackAnnotations" toml:"ResourceBuildpackAnnotations" yaml:"ResourceBuildpackAnnotations"`
	ResourceBuildpackLabels      BuildpackLabelSlice      `boil:"ResourceBuildpackLabels" json:"ResourceBuildpackLabels" toml:"ResourceBuildpackLabels" yaml:"ResourceBuildpackLabels"`
}

// NewStruct creates a new relationship struct
func (*buildpackR) NewStruct() *buildpackR {
	return &buildpackR{}
}

// buildpackL is where Load methods for each relationship are stored.
type buildpackL struct{}

var (
	buildpackAllColumns            = []string{"id", "guid", "created_at", "updated_at", "name", "key", "position", "enabled", "locked", "filename", "sha256_checksum", "stack"}
	buildpackColumnsWithoutDefault = []string{"guid", "updated_at", "name", "key", "position", "filename", "sha256_checksum", "stack"}
	buildpackColumnsWithDefault    = []string{"id", "created_at", "enabled", "locked"}
	buildpackPrimaryKeyColumns     = []string{"id"}
)

type (
	// BuildpackSlice is an alias for a slice of pointers to Buildpack.
	// This should almost always be used instead of []Buildpack.
	BuildpackSlice []*Buildpack

	buildpackQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	buildpackType                 = reflect.TypeOf(&Buildpack{})
	buildpackMapping              = queries.MakeStructMapping(buildpackType)
	buildpackPrimaryKeyMapping, _ = queries.BindMapping(buildpackType, buildpackMapping, buildpackPrimaryKeyColumns)
	buildpackInsertCacheMut       sync.RWMutex
	buildpackInsertCache          = make(map[string]insertCache)
	buildpackUpdateCacheMut       sync.RWMutex
	buildpackUpdateCache          = make(map[string]updateCache)
	buildpackUpsertCacheMut       sync.RWMutex
	buildpackUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type BuildpackFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*Buildpack, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (BuildpackSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single buildpack record from the query.
func (q buildpackQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Buildpack, error) {
	o := &Buildpack{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buildpacks")
	}

	return o, nil
}

// All returns all Buildpack records from the query.
func (q buildpackQuery) All(ctx context.Context, exec boil.ContextExecutor) (BuildpackSlice, error) {
	var o []*Buildpack

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Buildpack slice")
	}

	return o, nil
}

// Count returns the count of all Buildpack records in the query.
func (q buildpackQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buildpacks rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q buildpackQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buildpacks exists")
	}

	return count > 0, nil
}

// ResourceBuildpackAnnotations retrieves all the buildpack_annotation's BuildpackAnnotations with an executor via resource_guid column.
func (q buildpackQuery) ResourceBuildpackAnnotations(o *Buildpack, mods ...qm.QueryMod) buildpackAnnotationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"buildpack_annotations\".\"resource_guid\"=?", o.GUID),
	)

	query := BuildpackAnnotations(queryMods...)
	queries.SetFrom(query.Query, "\"buildpack_annotations\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"buildpack_annotations\".*"})
	}

	return query
}

// ResourceBuildpackLabels retrieves all the buildpack_label's BuildpackLabels with an executor via resource_guid column.
func (q buildpackQuery) ResourceBuildpackLabels(o *Buildpack, mods ...qm.QueryMod) buildpackLabelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"buildpack_labels\".\"resource_guid\"=?", o.GUID),
	)

	query := BuildpackLabels(queryMods...)
	queries.SetFrom(query.Query, "\"buildpack_labels\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"buildpack_labels\".*"})
	}

	return query
}

// LoadResourceBuildpackAnnotations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (buildpackL) LoadResourceBuildpackAnnotations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildpack interface{}, mods queries.Applicator) error {
	var slice []*Buildpack
	var object *Buildpack

	if singular {
		object = maybeBuildpack.(*Buildpack)
	} else {
		slice = *maybeBuildpack.(*[]*Buildpack)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildpackR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildpackR{}
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
		qm.From(`buildpack_annotations`),
		qm.WhereIn(`buildpack_annotations.resource_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load buildpack_annotations")
	}

	var resultSlice []*BuildpackAnnotation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice buildpack_annotations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on buildpack_annotations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for buildpack_annotations")
	}

	if singular {
		object.R.ResourceBuildpackAnnotations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &buildpackAnnotationR{}
			}
			foreign.R.Resource = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.GUID, foreign.ResourceGUID) {
				local.R.ResourceBuildpackAnnotations = append(local.R.ResourceBuildpackAnnotations, foreign)
				if foreign.R == nil {
					foreign.R = &buildpackAnnotationR{}
				}
				foreign.R.Resource = local
				break
			}
		}
	}

	return nil
}

// LoadResourceBuildpackLabels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (buildpackL) LoadResourceBuildpackLabels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBuildpack interface{}, mods queries.Applicator) error {
	var slice []*Buildpack
	var object *Buildpack

	if singular {
		object = maybeBuildpack.(*Buildpack)
	} else {
		slice = *maybeBuildpack.(*[]*Buildpack)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &buildpackR{}
		}
		args = append(args, object.GUID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &buildpackR{}
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
		qm.From(`buildpack_labels`),
		qm.WhereIn(`buildpack_labels.resource_guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load buildpack_labels")
	}

	var resultSlice []*BuildpackLabel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice buildpack_labels")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on buildpack_labels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for buildpack_labels")
	}

	if singular {
		object.R.ResourceBuildpackLabels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &buildpackLabelR{}
			}
			foreign.R.Resource = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.GUID, foreign.ResourceGUID) {
				local.R.ResourceBuildpackLabels = append(local.R.ResourceBuildpackLabels, foreign)
				if foreign.R == nil {
					foreign.R = &buildpackLabelR{}
				}
				foreign.R.Resource = local
				break
			}
		}
	}

	return nil
}

// AddResourceBuildpackAnnotations adds the given related objects to the existing relationships
// of the buildpack, optionally inserting them as new records.
// Appends related to o.R.ResourceBuildpackAnnotations.
// Sets related.R.Resource appropriately.
func (q buildpackQuery) AddResourceBuildpackAnnotations(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackAnnotation) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ResourceGUID, o.GUID)
			if err = BuildpackAnnotations().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"buildpack_annotations\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
				strmangle.WhereClause("\"", "\"", 2, buildpackAnnotationPrimaryKeyColumns),
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

			queries.Assign(&rel.ResourceGUID, o.GUID)
		}
	}

	if o.R == nil {
		o.R = &buildpackR{
			ResourceBuildpackAnnotations: related,
		}
	} else {
		o.R.ResourceBuildpackAnnotations = append(o.R.ResourceBuildpackAnnotations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &buildpackAnnotationR{
				Resource: o,
			}
		} else {
			rel.R.Resource = o
		}
	}
	return nil
}

// SetResourceBuildpackAnnotations removes all previously related items of the
// buildpack replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Resource's ResourceBuildpackAnnotations accordingly.
// Replaces o.R.ResourceBuildpackAnnotations with related.
// Sets related.R.Resource's ResourceBuildpackAnnotations accordingly.
func (q buildpackQuery) SetResourceBuildpackAnnotations(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackAnnotation) error {
	query := "update \"buildpack_annotations\" set \"resource_guid\" = null where \"resource_guid\" = $1"
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
		for _, rel := range o.R.ResourceBuildpackAnnotations {
			queries.SetScanner(&rel.ResourceGUID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Resource = nil
		}

		o.R.ResourceBuildpackAnnotations = nil
	}
	return q.AddResourceBuildpackAnnotations(o, ctx, exec, insert, related...)
}

// RemoveResourceBuildpackAnnotations relationships from objects passed in.
// Removes related items from R.ResourceBuildpackAnnotations (uses pointer comparison, removal does not keep order)
// Sets related.R.Resource.
func (q buildpackQuery) RemoveResourceBuildpackAnnotations(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, related ...*BuildpackAnnotation) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ResourceGUID, nil)
		if rel.R != nil {
			rel.R.Resource = nil
		}
		if _, err = BuildpackAnnotations().Update(rel, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ResourceBuildpackAnnotations {
			if rel != ri {
				continue
			}

			ln := len(o.R.ResourceBuildpackAnnotations)
			if ln > 1 && i < ln-1 {
				o.R.ResourceBuildpackAnnotations[i] = o.R.ResourceBuildpackAnnotations[ln-1]
			}
			o.R.ResourceBuildpackAnnotations = o.R.ResourceBuildpackAnnotations[:ln-1]
			break
		}
	}

	return nil
}

// AddResourceBuildpackLabels adds the given related objects to the existing relationships
// of the buildpack, optionally inserting them as new records.
// Appends related to o.R.ResourceBuildpackLabels.
// Sets related.R.Resource appropriately.
func (q buildpackQuery) AddResourceBuildpackLabels(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackLabel) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ResourceGUID, o.GUID)
			if err = BuildpackLabels().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"buildpack_labels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"resource_guid"}),
				strmangle.WhereClause("\"", "\"", 2, buildpackLabelPrimaryKeyColumns),
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

			queries.Assign(&rel.ResourceGUID, o.GUID)
		}
	}

	if o.R == nil {
		o.R = &buildpackR{
			ResourceBuildpackLabels: related,
		}
	} else {
		o.R.ResourceBuildpackLabels = append(o.R.ResourceBuildpackLabels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &buildpackLabelR{
				Resource: o,
			}
		} else {
			rel.R.Resource = o
		}
	}
	return nil
}

// SetResourceBuildpackLabels removes all previously related items of the
// buildpack replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Resource's ResourceBuildpackLabels accordingly.
// Replaces o.R.ResourceBuildpackLabels with related.
// Sets related.R.Resource's ResourceBuildpackLabels accordingly.
func (q buildpackQuery) SetResourceBuildpackLabels(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*BuildpackLabel) error {
	query := "update \"buildpack_labels\" set \"resource_guid\" = null where \"resource_guid\" = $1"
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
		for _, rel := range o.R.ResourceBuildpackLabels {
			queries.SetScanner(&rel.ResourceGUID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Resource = nil
		}

		o.R.ResourceBuildpackLabels = nil
	}
	return q.AddResourceBuildpackLabels(o, ctx, exec, insert, related...)
}

// RemoveResourceBuildpackLabels relationships from objects passed in.
// Removes related items from R.ResourceBuildpackLabels (uses pointer comparison, removal does not keep order)
// Sets related.R.Resource.
func (q buildpackQuery) RemoveResourceBuildpackLabels(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, related ...*BuildpackLabel) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ResourceGUID, nil)
		if rel.R != nil {
			rel.R.Resource = nil
		}
		if _, err = BuildpackLabels().Update(rel, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ResourceBuildpackLabels {
			if rel != ri {
				continue
			}

			ln := len(o.R.ResourceBuildpackLabels)
			if ln > 1 && i < ln-1 {
				o.R.ResourceBuildpackLabels[i] = o.R.ResourceBuildpackLabels[ln-1]
			}
			o.R.ResourceBuildpackLabels = o.R.ResourceBuildpackLabels[:ln-1]
			break
		}
	}

	return nil
}

// Buildpacks retrieves all the records using an executor.
func Buildpacks(mods ...qm.QueryMod) buildpackQuery {
	mods = append(mods, qm.From("\"buildpacks\""))
	return buildpackQuery{NewQuery(mods...)}
}

type BuildpackFinder interface {
	FindBuildpack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Buildpack, error)
}

// FindBuildpack retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBuildpack(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Buildpack, error) {
	buildpackObj := &Buildpack{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"buildpacks\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, buildpackObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buildpacks")
	}

	return buildpackObj, nil
}

type BuildpackInserter interface {
	Insert(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q buildpackQuery) Insert(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpacks provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(buildpackColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	buildpackInsertCacheMut.RLock()
	cache, cached := buildpackInsertCache[key]
	buildpackInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			buildpackAllColumns,
			buildpackColumnsWithDefault,
			buildpackColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(buildpackType, buildpackMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(buildpackType, buildpackMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"buildpacks\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"buildpacks\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into buildpacks")
	}

	if !cached {
		buildpackInsertCacheMut.Lock()
		buildpackInsertCache[key] = cache
		buildpackInsertCacheMut.Unlock()
	}

	return nil
}

type BuildpackUpdater interface {
	Update(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o BuildpackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the Buildpack.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q buildpackQuery) Update(o *Buildpack, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	buildpackUpdateCacheMut.RLock()
	cache, cached := buildpackUpdateCache[key]
	buildpackUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			buildpackAllColumns,
			buildpackPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buildpacks, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"buildpacks\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, buildpackPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(buildpackType, buildpackMapping, append(wl, buildpackPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update buildpacks row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buildpacks")
	}

	if !cached {
		buildpackUpdateCacheMut.Lock()
		buildpackUpdateCache[key] = cache
		buildpackUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q buildpackQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buildpacks")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q buildpackQuery) UpdateAllSlice(o BuildpackSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"buildpacks\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, buildpackPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in buildpack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all buildpack")
	}
	return rowsAff, nil
}

type BuildpackDeleter interface {
	Delete(o *Buildpack, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single Buildpack record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q buildpackQuery) Delete(o *Buildpack, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Buildpack provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), buildpackPrimaryKeyMapping)
	sql := "DELETE FROM \"buildpacks\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buildpacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q buildpackQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no buildpackQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpacks")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpacks")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q buildpackQuery) DeleteAllSlice(o BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"buildpacks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buildpack slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buildpacks")
	}

	return rowsAff, nil
}

type BuildpackReloader interface {
	Reload(o *Buildpack, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q buildpackQuery) Reload(o *Buildpack, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBuildpack(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q buildpackQuery) ReloadAll(o *BuildpackSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BuildpackSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), buildpackPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"buildpacks\".* FROM \"buildpacks\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, buildpackPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BuildpackSlice")
	}

	*o = slice

	return nil
}

// BuildpackExists checks if the Buildpack row exists.
func BuildpackExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"buildpacks\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buildpacks exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Buildpack) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buildpacks provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(buildpackColumnsWithDefault, o)

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

	buildpackUpsertCacheMut.RLock()
	cache, cached := buildpackUpsertCache[key]
	buildpackUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			buildpackAllColumns,
			buildpackColumnsWithDefault,
			buildpackColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			buildpackAllColumns,
			buildpackPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert buildpacks, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(buildpackPrimaryKeyColumns))
			copy(conflict, buildpackPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"buildpacks\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(buildpackType, buildpackMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(buildpackType, buildpackMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert buildpacks")
	}

	if !cached {
		buildpackUpsertCacheMut.Lock()
		buildpackUpsertCache[key] = cache
		buildpackUpsertCacheMut.Unlock()
	}

	return nil
}
