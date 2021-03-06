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

type DeploymentProcessUpserter interface {
	Upsert(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLDeploymentProcessUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q deploymentProcessQuery) Upsert(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deployment_processes provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(deploymentProcessColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDeploymentProcessUniqueColumns, o)

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

	deploymentProcessUpsertCacheMut.RLock()
	cache, cached := deploymentProcessUpsertCache[key]
	deploymentProcessUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deploymentProcessAllColumns,
			deploymentProcessColumnsWithDefault,
			deploymentProcessColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			deploymentProcessAllColumns,
			deploymentProcessPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert deployment_processes, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`deployment_processes`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `deployment_processes` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for deployment_processes")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deploymentProcessMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for deployment_processes")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for deployment_processes")
	}

CacheNoHooks:
	if !cached {
		deploymentProcessUpsertCacheMut.Lock()
		deploymentProcessUpsertCache[key] = cache
		deploymentProcessUpsertCacheMut.Unlock()
	}

	return nil
}

// DeploymentProcess is an object representing the database table.
type DeploymentProcess struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID           string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt      time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	ProcessGUID    null.String `boil:"process_guid" json:"process_guid,omitempty" toml:"process_guid" yaml:"process_guid,omitempty"`
	ProcessType    null.String `boil:"process_type" json:"process_type,omitempty" toml:"process_type" yaml:"process_type,omitempty"`
	DeploymentGUID null.String `boil:"deployment_guid" json:"deployment_guid,omitempty" toml:"deployment_guid" yaml:"deployment_guid,omitempty"`

	R *deploymentProcessR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deploymentProcessL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeploymentProcessColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	ProcessGUID    string
	ProcessType    string
	DeploymentGUID string
}{
	ID:             "id",
	GUID:           "guid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	ProcessGUID:    "process_guid",
	ProcessType:    "process_type",
	DeploymentGUID: "deployment_guid",
}

var DeploymentProcessTableColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	ProcessGUID    string
	ProcessType    string
	DeploymentGUID string
}{
	ID:             "deployment_processes.id",
	GUID:           "deployment_processes.guid",
	CreatedAt:      "deployment_processes.created_at",
	UpdatedAt:      "deployment_processes.updated_at",
	ProcessGUID:    "deployment_processes.process_guid",
	ProcessType:    "deployment_processes.process_type",
	DeploymentGUID: "deployment_processes.deployment_guid",
}

// Generated where

var DeploymentProcessWhere = struct {
	ID             whereHelperint
	GUID           whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpernull_Time
	ProcessGUID    whereHelpernull_String
	ProcessType    whereHelpernull_String
	DeploymentGUID whereHelpernull_String
}{
	ID:             whereHelperint{field: "`deployment_processes`.`id`"},
	GUID:           whereHelperstring{field: "`deployment_processes`.`guid`"},
	CreatedAt:      whereHelpertime_Time{field: "`deployment_processes`.`created_at`"},
	UpdatedAt:      whereHelpernull_Time{field: "`deployment_processes`.`updated_at`"},
	ProcessGUID:    whereHelpernull_String{field: "`deployment_processes`.`process_guid`"},
	ProcessType:    whereHelpernull_String{field: "`deployment_processes`.`process_type`"},
	DeploymentGUID: whereHelpernull_String{field: "`deployment_processes`.`deployment_guid`"},
}

// DeploymentProcessRels is where relationship names are stored.
var DeploymentProcessRels = struct {
	Deployment string
}{
	Deployment: "Deployment",
}

// deploymentProcessR is where relationships are stored.
type deploymentProcessR struct {
	Deployment *Deployment `boil:"Deployment" json:"Deployment" toml:"Deployment" yaml:"Deployment"`
}

// NewStruct creates a new relationship struct
func (*deploymentProcessR) NewStruct() *deploymentProcessR {
	return &deploymentProcessR{}
}

// deploymentProcessL is where Load methods for each relationship are stored.
type deploymentProcessL struct{}

var (
	deploymentProcessAllColumns            = []string{"id", "guid", "created_at", "updated_at", "process_guid", "process_type", "deployment_guid"}
	deploymentProcessColumnsWithoutDefault = []string{"guid", "updated_at", "process_guid", "process_type", "deployment_guid"}
	deploymentProcessColumnsWithDefault    = []string{"id", "created_at"}
	deploymentProcessPrimaryKeyColumns     = []string{"id"}
)

type (
	// DeploymentProcessSlice is an alias for a slice of pointers to DeploymentProcess.
	// This should almost always be used instead of []DeploymentProcess.
	DeploymentProcessSlice []*DeploymentProcess

	deploymentProcessQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deploymentProcessType                 = reflect.TypeOf(&DeploymentProcess{})
	deploymentProcessMapping              = queries.MakeStructMapping(deploymentProcessType)
	deploymentProcessPrimaryKeyMapping, _ = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, deploymentProcessPrimaryKeyColumns)
	deploymentProcessInsertCacheMut       sync.RWMutex
	deploymentProcessInsertCache          = make(map[string]insertCache)
	deploymentProcessUpdateCacheMut       sync.RWMutex
	deploymentProcessUpdateCache          = make(map[string]updateCache)
	deploymentProcessUpsertCacheMut       sync.RWMutex
	deploymentProcessUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type DeploymentProcessFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*DeploymentProcess, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (DeploymentProcessSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single deploymentProcess record from the query.
func (q deploymentProcessQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeploymentProcess, error) {
	o := &DeploymentProcess{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for deployment_processes")
	}

	return o, nil
}

// All returns all DeploymentProcess records from the query.
func (q deploymentProcessQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeploymentProcessSlice, error) {
	var o []*DeploymentProcess

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeploymentProcess slice")
	}

	return o, nil
}

// Count returns the count of all DeploymentProcess records in the query.
func (q deploymentProcessQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count deployment_processes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deploymentProcessQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if deployment_processes exists")
	}

	return count > 0, nil
}

// Deployment pointed to by the foreign key.
func (q deploymentProcessQuery) Deployment(o *DeploymentProcess, mods ...qm.QueryMod) deploymentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`guid` = ?", o.DeploymentGUID),
	}

	queryMods = append(queryMods, mods...)

	query := Deployments(queryMods...)
	queries.SetFrom(query.Query, "`deployments`")

	return query
}

// LoadDeployment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (deploymentProcessL) LoadDeployment(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDeploymentProcess interface{}, mods queries.Applicator) error {
	var slice []*DeploymentProcess
	var object *DeploymentProcess

	if singular {
		object = maybeDeploymentProcess.(*DeploymentProcess)
	} else {
		slice = *maybeDeploymentProcess.(*[]*DeploymentProcess)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &deploymentProcessR{}
		}
		if !queries.IsNil(object.DeploymentGUID) {
			args = append(args, object.DeploymentGUID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &deploymentProcessR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DeploymentGUID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DeploymentGUID) {
				args = append(args, obj.DeploymentGUID)
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
		object.R.Deployment = foreign
		if foreign.R == nil {
			foreign.R = &deploymentR{}
		}
		foreign.R.DeploymentProcesses = append(foreign.R.DeploymentProcesses, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DeploymentGUID, foreign.GUID) {
				local.R.Deployment = foreign
				if foreign.R == nil {
					foreign.R = &deploymentR{}
				}
				foreign.R.DeploymentProcesses = append(foreign.R.DeploymentProcesses, local)
				break
			}
		}
	}

	return nil
}

// SetDeployment of the deploymentProcess to the related item.
// Sets o.R.Deployment to related.
// Adds o to related.R.DeploymentProcesses.
func (q deploymentProcessQuery) SetDeployment(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Deployment) error {
	var err error
	if insert {
		if err = Deployments().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `deployment_processes` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"deployment_guid"}),
		strmangle.WhereClause("`", "`", 0, deploymentProcessPrimaryKeyColumns),
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

	queries.Assign(&o.DeploymentGUID, related.GUID)
	if o.R == nil {
		o.R = &deploymentProcessR{
			Deployment: related,
		}
	} else {
		o.R.Deployment = related
	}

	if related.R == nil {
		related.R = &deploymentR{
			DeploymentProcesses: DeploymentProcessSlice{o},
		}
	} else {
		related.R.DeploymentProcesses = append(related.R.DeploymentProcesses, o)
	}

	return nil
}

// RemoveDeployment relationship.
// Sets o.R.Deployment to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q deploymentProcessQuery) RemoveDeployment(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, related *Deployment) error {
	var err error

	queries.SetScanner(&o.DeploymentGUID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("deployment_guid")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Deployment = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.DeploymentProcesses {
		if queries.Equal(o.DeploymentGUID, ri.DeploymentGUID) {
			continue
		}

		ln := len(related.R.DeploymentProcesses)
		if ln > 1 && i < ln-1 {
			related.R.DeploymentProcesses[i] = related.R.DeploymentProcesses[ln-1]
		}
		related.R.DeploymentProcesses = related.R.DeploymentProcesses[:ln-1]
		break
	}
	return nil
}

// DeploymentProcesses retrieves all the records using an executor.
func DeploymentProcesses(mods ...qm.QueryMod) deploymentProcessQuery {
	mods = append(mods, qm.From("`deployment_processes`"))
	return deploymentProcessQuery{NewQuery(mods...)}
}

type DeploymentProcessFinder interface {
	FindDeploymentProcess(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DeploymentProcess, error)
}

// FindDeploymentProcess retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeploymentProcess(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DeploymentProcess, error) {
	deploymentProcessObj := &DeploymentProcess{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `deployment_processes` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, deploymentProcessObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from deployment_processes")
	}

	return deploymentProcessObj, nil
}

type DeploymentProcessInserter interface {
	Insert(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q deploymentProcessQuery) Insert(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deployment_processes provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(deploymentProcessColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deploymentProcessInsertCacheMut.RLock()
	cache, cached := deploymentProcessInsertCache[key]
	deploymentProcessInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deploymentProcessAllColumns,
			deploymentProcessColumnsWithDefault,
			deploymentProcessColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `deployment_processes` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `deployment_processes` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `deployment_processes` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, deploymentProcessPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into deployment_processes")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deploymentProcessMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for deployment_processes")
	}

CacheNoHooks:
	if !cached {
		deploymentProcessInsertCacheMut.Lock()
		deploymentProcessInsertCache[key] = cache
		deploymentProcessInsertCacheMut.Unlock()
	}

	return nil
}

type DeploymentProcessUpdater interface {
	Update(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the DeploymentProcess.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q deploymentProcessQuery) Update(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	deploymentProcessUpdateCacheMut.RLock()
	cache, cached := deploymentProcessUpdateCache[key]
	deploymentProcessUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deploymentProcessAllColumns,
			deploymentProcessPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update deployment_processes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `deployment_processes` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, deploymentProcessPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deploymentProcessType, deploymentProcessMapping, append(wl, deploymentProcessPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update deployment_processes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for deployment_processes")
	}

	if !cached {
		deploymentProcessUpdateCacheMut.Lock()
		deploymentProcessUpdateCache[key] = cache
		deploymentProcessUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q deploymentProcessQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for deployment_processes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for deployment_processes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q deploymentProcessQuery) UpdateAllSlice(o DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentProcessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `deployment_processes` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentProcessPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deploymentProcess slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deploymentProcess")
	}
	return rowsAff, nil
}

type DeploymentProcessDeleter interface {
	Delete(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single DeploymentProcess record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q deploymentProcessQuery) Delete(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeploymentProcess provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deploymentProcessPrimaryKeyMapping)
	sql := "DELETE FROM `deployment_processes` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from deployment_processes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for deployment_processes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deploymentProcessQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deploymentProcessQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deployment_processes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deployment_processes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q deploymentProcessQuery) DeleteAllSlice(o DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentProcessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `deployment_processes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentProcessPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deploymentProcess slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deployment_processes")
	}

	return rowsAff, nil
}

type DeploymentProcessReloader interface {
	Reload(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q deploymentProcessQuery) Reload(o *DeploymentProcess, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDeploymentProcess(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q deploymentProcessQuery) ReloadAll(o *DeploymentProcessSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeploymentProcessSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deploymentProcessPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `deployment_processes`.* FROM `deployment_processes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deploymentProcessPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeploymentProcessSlice")
	}

	*o = slice

	return nil
}

// DeploymentProcessExists checks if the DeploymentProcess row exists.
func DeploymentProcessExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `deployment_processes` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if deployment_processes exists")
	}

	return exists, nil
}
