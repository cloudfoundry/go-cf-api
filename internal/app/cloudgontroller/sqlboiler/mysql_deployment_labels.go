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

// DeploymentLabel is an object representing the database table.
type DeploymentLabel struct {
	ID           int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID         string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt    time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ResourceGUID null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	KeyPrefix    null.String `boil:"key_prefix" json:"key_prefix,omitempty" toml:"key_prefix" yaml:"key_prefix,omitempty"`
	KeyName      null.String `boil:"key_name" json:"key_name,omitempty" toml:"key_name" yaml:"key_name,omitempty"`
	Value        null.String `boil:"value" json:"value,omitempty" toml:"value" yaml:"value,omitempty"`

	R *deploymentLabelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deploymentLabelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeploymentLabelColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "id",
	GUID:         "guid",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	ResourceGUID: "resource_guid",
	KeyPrefix:    "key_prefix",
	KeyName:      "key_name",
	Value:        "value",
}

var DeploymentLabelTableColumns = struct {
	ID           string
	GUID         string
	CreatedAt    string
	UpdatedAt    string
	ResourceGUID string
	KeyPrefix    string
	KeyName      string
	Value        string
}{
	ID:           "deployment_labels.id",
	GUID:         "deployment_labels.guid",
	CreatedAt:    "deployment_labels.created_at",
	UpdatedAt:    "deployment_labels.updated_at",
	ResourceGUID: "deployment_labels.resource_guid",
	KeyPrefix:    "deployment_labels.key_prefix",
	KeyName:      "deployment_labels.key_name",
	Value:        "deployment_labels.value",
}

// Generated where

var DeploymentLabelWhere = struct {
	ID           whereHelperint
	GUID         whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpernull_Time
	ResourceGUID whereHelpernull_String
	KeyPrefix    whereHelpernull_String
	KeyName      whereHelpernull_String
	Value        whereHelpernull_String
}{
	ID:           whereHelperint{field: "`deployment_labels`.`id`"},
	GUID:         whereHelperstring{field: "`deployment_labels`.`guid`"},
	CreatedAt:    whereHelpertime_Time{field: "`deployment_labels`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`deployment_labels`.`updated_at`"},
	ResourceGUID: whereHelpernull_String{field: "`deployment_labels`.`resource_guid`"},
	KeyPrefix:    whereHelpernull_String{field: "`deployment_labels`.`key_prefix`"},
	KeyName:      whereHelpernull_String{field: "`deployment_labels`.`key_name`"},
	Value:        whereHelpernull_String{field: "`deployment_labels`.`value`"},
}

// DeploymentLabelRels is where relationship names are stored.
var DeploymentLabelRels = struct {
	Resource string
}{
	Resource: "Resource",
}

// deploymentLabelR is where relationships are stored.
type deploymentLabelR struct {
	Resource *Deployment `boil:"Resource" json:"Resource" toml:"Resource" yaml:"Resource"`
}

// NewStruct creates a new relationship struct
func (*deploymentLabelR) NewStruct() *deploymentLabelR {
	return &deploymentLabelR{}
}

// deploymentLabelL is where Load methods for each relationship are stored.
type deploymentLabelL struct{}

var (
	deploymentLabelAllColumns            = []string{"id", "guid", "created_at", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	deploymentLabelColumnsWithoutDefault = []string{"guid", "updated_at", "resource_guid", "key_prefix", "key_name", "value"}
	deploymentLabelColumnsWithDefault    = []string{"id", "created_at"}
	deploymentLabelPrimaryKeyColumns     = []string{"id"}
)

type (
	// DeploymentLabelSlice is an alias for a slice of pointers to DeploymentLabel.
	// This should almost always be used instead of []DeploymentLabel.
	DeploymentLabelSlice []*DeploymentLabel

	deploymentLabelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deploymentLabelType                 = reflect.TypeOf(&DeploymentLabel{})
	deploymentLabelMapping              = queries.MakeStructMapping(deploymentLabelType)
	deploymentLabelPrimaryKeyMapping, _ = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, deploymentLabelPrimaryKeyColumns)
	deploymentLabelInsertCacheMut       sync.RWMutex
	deploymentLabelInsertCache          = make(map[string]insertCache)
	deploymentLabelUpdateCacheMut       sync.RWMutex
	deploymentLabelUpdateCache          = make(map[string]updateCache)
	deploymentLabelUpsertCacheMut       sync.RWMutex
	deploymentLabelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type DeploymentLabelFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*DeploymentLabel, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (DeploymentLabelSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single deploymentLabel record from the query.
func (q deploymentLabelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeploymentLabel, error) {
	o := &DeploymentLabel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for deployment_labels")
	}

	return o, nil
}

// All returns all DeploymentLabel records from the query.
func (q deploymentLabelQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeploymentLabelSlice, error) {
	var o []*DeploymentLabel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeploymentLabel slice")
	}

	return o, nil
}

// Count returns the count of all DeploymentLabel records in the query.
func (q deploymentLabelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count deployment_labels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deploymentLabelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if deployment_labels exists")
	}

	return count > 0, nil
}

// Resource pointed to by the foreign key.
func (q deploymentLabelQuery) Resource(o *DeploymentLabel, mods ...qm.QueryMod) deploymentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.ResourceGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Deployments(queryMods...)
	queries.SetFrom(query.Query, "`deployments`")

	return query
}

// LoadResource allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (deploymentLabelL) LoadResource(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDeploymentLabel interface{}, mods queries.Applicator) error {
	var slice []*DeploymentLabel
	var object *DeploymentLabel

	if singular {
		object = maybeDeploymentLabel.(*DeploymentLabel)
	} else {
		slice = *maybeDeploymentLabel.(*[]*DeploymentLabel)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &deploymentLabelR{}
		}
		if !queries.IsNil(object.ResourceGUID) {
			args = append(args, object.ResourceGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &deploymentLabelR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ResourceGUID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ResourceGUID) {
				args = append(args, obj.ResourceGUID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`deployments`),
		qm.WhereIn(`deployments.guid in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Deployment")
	}

	var resultSlice []*Deployment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Deployment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for deployments")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for deployments")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Resource = foreign
		if foreign.R == nil {
			foreign.R = &deploymentR{}
		}
		foreign.R.ResourceDeploymentLabels = append(foreign.R.ResourceDeploymentLabels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceGUID, foreign.GUID) {
				local.R.Resource = foreign
				if foreign.R == nil {
					foreign.R = &deploymentR{}
				}
				foreign.R.ResourceDeploymentLabels = append(foreign.R.ResourceDeploymentLabels, local)
				break
			}
		}
	}

	return nil
}

// SetResource of the deploymentLabel to the related item.
// Sets o.R.Resource to related.
// Adds o to related.R.ResourceDeploymentLabels.
func (q deploymentLabelQuery) SetResource(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Deployment) error {
	var err error
	if insert {
		if err = Deployments().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `deployment_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_guid"}),
		strmangle.WhereClause("`", "`", 0, deploymentLabelPrimaryKeyColumns),
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

	queries.Assign(&o.ResourceGUID, related.GUID)
	if o.R == nil {
		o.R = &deploymentLabelR{
			Resource: related,
		}
	} else {
		o.R.Resource = related
	}

	if related.R == nil {
		related.R = &deploymentR{
			ResourceDeploymentLabels: DeploymentLabelSlice{o},
		}
	} else {
		related.R.ResourceDeploymentLabels = append(related.R.ResourceDeploymentLabels, o)
	}

	return nil
}

// RemoveResource relationship.
// Sets o.R.Resource to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q deploymentLabelQuery) RemoveResource(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, related *Deployment) error {
	var err error

	queries.SetScanner(&o.ResourceGUID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("resource_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Resource = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ResourceDeploymentLabels {
		if queries.Equal(o.ResourceGUID, ri.ResourceGUID) {
			continue
		}

		ln := len(related.R.ResourceDeploymentLabels)
		if ln > 1 && i < ln-1 {
			related.R.ResourceDeploymentLabels[i] = related.R.ResourceDeploymentLabels[ln-1]
		}
		related.R.ResourceDeploymentLabels = related.R.ResourceDeploymentLabels[:ln-1]
		break
	}
	return nil
}

// DeploymentLabels retrieves all the records using an executor.
func DeploymentLabels(mods ...qm.QueryMod) deploymentLabelQuery {
	mods = append(mods, qm.From("`deployment_labels`"))
	return deploymentLabelQuery{NewQuery(mods...)}
}

type DeploymentLabelFinder interface {
	FindDeploymentLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DeploymentLabel, error)
}

// FindDeploymentLabel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeploymentLabel(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DeploymentLabel, error) {
	deploymentLabelObj := &DeploymentLabel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `deployment_labels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, deploymentLabelObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from deployment_labels")
	}

	return deploymentLabelObj, nil
}

type DeploymentLabelInserter interface {
	Insert(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q deploymentLabelQuery) Insert(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deployment_labels provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(deploymentLabelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deploymentLabelInsertCacheMut.RLock()
	cache, cached := deploymentLabelInsertCache[key]
	deploymentLabelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deploymentLabelAllColumns,
			deploymentLabelColumnsWithDefault,
			deploymentLabelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `deployment_labels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `deployment_labels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `deployment_labels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, deploymentLabelPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into deployment_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deploymentLabelMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for deployment_labels")
	}

CacheNoHooks:
	if !cached {
		deploymentLabelInsertCacheMut.Lock()
		deploymentLabelInsertCache[key] = cache
		deploymentLabelInsertCacheMut.Unlock()
	}

	return nil
}

type DeploymentLabelUpdater interface {
	Update(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the DeploymentLabel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q deploymentLabelQuery) Update(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	deploymentLabelUpdateCacheMut.RLock()
	cache, cached := deploymentLabelUpdateCache[key]
	deploymentLabelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deploymentLabelAllColumns,
			deploymentLabelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update deployment_labels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `deployment_labels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, deploymentLabelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, append(wl, deploymentLabelPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update deployment_labels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for deployment_labels")
	}

	if !cached {
		deploymentLabelUpdateCacheMut.Lock()
		deploymentLabelUpdateCache[key] = cache
		deploymentLabelUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q deploymentLabelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for deployment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for deployment_labels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q deploymentLabelQuery) UpdateAllSlice(o DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `deployment_labels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentLabelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deploymentLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deploymentLabel")
	}
	return rowsAff, nil
}

type DeploymentLabelUpserter interface {
	Upsert(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLDeploymentLabelUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q deploymentLabelQuery) Upsert(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deployment_labels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(deploymentLabelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDeploymentLabelUniqueColumns, o)

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

	deploymentLabelUpsertCacheMut.RLock()
	cache, cached := deploymentLabelUpsertCache[key]
	deploymentLabelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deploymentLabelAllColumns,
			deploymentLabelColumnsWithDefault,
			deploymentLabelColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			deploymentLabelAllColumns,
			deploymentLabelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert deployment_labels, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`deployment_labels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `deployment_labels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for deployment_labels")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deploymentLabelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(deploymentLabelType, deploymentLabelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for deployment_labels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for deployment_labels")
	}

CacheNoHooks:
	if !cached {
		deploymentLabelUpsertCacheMut.Lock()
		deploymentLabelUpsertCache[key] = cache
		deploymentLabelUpsertCacheMut.Unlock()
	}

	return nil
}

type DeploymentLabelDeleter interface {
	Delete(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single DeploymentLabel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q deploymentLabelQuery) Delete(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeploymentLabel provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deploymentLabelPrimaryKeyMapping)
	sql := "DELETE FROM `deployment_labels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from deployment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for deployment_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deploymentLabelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deploymentLabelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deployment_labels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deployment_labels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q deploymentLabelQuery) DeleteAllSlice(o DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `deployment_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentLabelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deploymentLabel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deployment_labels")
	}

	return rowsAff, nil
}

type DeploymentLabelReloader interface {
	Reload(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q deploymentLabelQuery) Reload(o *DeploymentLabel, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDeploymentLabel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q deploymentLabelQuery) ReloadAll(o *DeploymentLabelSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeploymentLabelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentLabelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `deployment_labels`.* FROM `deployment_labels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentLabelPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeploymentLabelSlice")
	}

	*o = slice

	return nil
}

// DeploymentLabelExists checks if the DeploymentLabel row exists.
func DeploymentLabelExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `deployment_labels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if deployment_labels exists")
	}

	return exists, nil
}
