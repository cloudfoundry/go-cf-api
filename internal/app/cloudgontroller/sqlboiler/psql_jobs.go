//go:generate mockgen -source=$GOFILE -destination=mocks/jobs.go
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

// Job is an object representing the database table.
type Job struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID           string      `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt      time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt      null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	State          null.String `boil:"state" json:"state,omitempty" toml:"state" yaml:"state,omitempty"`
	Operation      null.String `boil:"operation" json:"operation,omitempty" toml:"operation" yaml:"operation,omitempty"`
	ResourceGUID   null.String `boil:"resource_guid" json:"resource_guid,omitempty" toml:"resource_guid" yaml:"resource_guid,omitempty"`
	ResourceType   null.String `boil:"resource_type" json:"resource_type,omitempty" toml:"resource_type" yaml:"resource_type,omitempty"`
	DelayedJobGUID null.String `boil:"delayed_job_guid" json:"delayed_job_guid,omitempty" toml:"delayed_job_guid" yaml:"delayed_job_guid,omitempty"`
	CFAPIError     null.String `boil:"cf_api_error" json:"cf_api_error,omitempty" toml:"cf_api_error" yaml:"cf_api_error,omitempty"`

	R *jobR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jobL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JobColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	State          string
	Operation      string
	ResourceGUID   string
	ResourceType   string
	DelayedJobGUID string
	CFAPIError     string
}{
	ID:             "id",
	GUID:           "guid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	State:          "state",
	Operation:      "operation",
	ResourceGUID:   "resource_guid",
	ResourceType:   "resource_type",
	DelayedJobGUID: "delayed_job_guid",
	CFAPIError:     "cf_api_error",
}

var JobTableColumns = struct {
	ID             string
	GUID           string
	CreatedAt      string
	UpdatedAt      string
	State          string
	Operation      string
	ResourceGUID   string
	ResourceType   string
	DelayedJobGUID string
	CFAPIError     string
}{
	ID:             "jobs.id",
	GUID:           "jobs.guid",
	CreatedAt:      "jobs.created_at",
	UpdatedAt:      "jobs.updated_at",
	State:          "jobs.state",
	Operation:      "jobs.operation",
	ResourceGUID:   "jobs.resource_guid",
	ResourceType:   "jobs.resource_type",
	DelayedJobGUID: "jobs.delayed_job_guid",
	CFAPIError:     "jobs.cf_api_error",
}

// Generated where

var JobWhere = struct {
	ID             whereHelperint
	GUID           whereHelperstring
	CreatedAt      whereHelpertime_Time
	UpdatedAt      whereHelpernull_Time
	State          whereHelpernull_String
	Operation      whereHelpernull_String
	ResourceGUID   whereHelpernull_String
	ResourceType   whereHelpernull_String
	DelayedJobGUID whereHelpernull_String
	CFAPIError     whereHelpernull_String
}{
	ID:             whereHelperint{field: "\"jobs\".\"id\""},
	GUID:           whereHelperstring{field: "\"jobs\".\"guid\""},
	CreatedAt:      whereHelpertime_Time{field: "\"jobs\".\"created_at\""},
	UpdatedAt:      whereHelpernull_Time{field: "\"jobs\".\"updated_at\""},
	State:          whereHelpernull_String{field: "\"jobs\".\"state\""},
	Operation:      whereHelpernull_String{field: "\"jobs\".\"operation\""},
	ResourceGUID:   whereHelpernull_String{field: "\"jobs\".\"resource_guid\""},
	ResourceType:   whereHelpernull_String{field: "\"jobs\".\"resource_type\""},
	DelayedJobGUID: whereHelpernull_String{field: "\"jobs\".\"delayed_job_guid\""},
	CFAPIError:     whereHelpernull_String{field: "\"jobs\".\"cf_api_error\""},
}

// JobRels is where relationship names are stored.
var JobRels = struct {
	FKJobJobWarnings string
}{
	FKJobJobWarnings: "FKJobJobWarnings",
}

// jobR is where relationships are stored.
type jobR struct {
	FKJobJobWarnings JobWarningSlice `boil:"FKJobJobWarnings" json:"FKJobJobWarnings" toml:"FKJobJobWarnings" yaml:"FKJobJobWarnings"`
}

// NewStruct creates a new relationship struct
func (*jobR) NewStruct() *jobR {
	return &jobR{}
}

// jobL is where Load methods for each relationship are stored.
type jobL struct{}

var (
	jobAllColumns            = []string{"id", "guid", "created_at", "updated_at", "state", "operation", "resource_guid", "resource_type", "delayed_job_guid", "cf_api_error"}
	jobColumnsWithoutDefault = []string{"guid", "updated_at", "state", "operation", "resource_guid", "resource_type", "delayed_job_guid", "cf_api_error"}
	jobColumnsWithDefault    = []string{"id", "created_at"}
	jobPrimaryKeyColumns     = []string{"id"}
)

type (
	// JobSlice is an alias for a slice of pointers to Job.
	// This should almost always be used instead of []Job.
	JobSlice []*Job

	jobQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jobType                 = reflect.TypeOf(&Job{})
	jobMapping              = queries.MakeStructMapping(jobType)
	jobPrimaryKeyMapping, _ = queries.BindMapping(jobType, jobMapping, jobPrimaryKeyColumns)
	jobInsertCacheMut       sync.RWMutex
	jobInsertCache          = make(map[string]insertCache)
	jobUpdateCacheMut       sync.RWMutex
	jobUpdateCache          = make(map[string]updateCache)
	jobUpsertCacheMut       sync.RWMutex
	jobUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type JobFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*Job, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (JobSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single job record from the query.
func (q jobQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Job, error) {
	o := &Job{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for jobs")
	}

	return o, nil
}

// All returns all Job records from the query.
func (q jobQuery) All(ctx context.Context, exec boil.ContextExecutor) (JobSlice, error) {
	var o []*Job

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Job slice")
	}

	return o, nil
}

// Count returns the count of all Job records in the query.
func (q jobQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count jobs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q jobQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if jobs exists")
	}

	return count > 0, nil
}

// FKJobJobWarnings retrieves all the job_warning's JobWarnings with an executor via fk_jobs_id column.
func (q jobQuery) FKJobJobWarnings(o *Job, mods ...qm.QueryMod) jobWarningQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"job_warnings\".\"fk_jobs_id\"=?", o.ID),
	)

	query := JobWarnings(queryMods...)
	queries.SetFrom(query.Query, "\"job_warnings\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"job_warnings\".*"})
	}

	return query
}

// LoadFKJobJobWarnings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (jobL) LoadFKJobJobWarnings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJob interface{}, mods queries.Applicator) error {
	var slice []*Job
	var object *Job

	if singular {
		object = maybeJob.(*Job)
	} else {
		slice = *maybeJob.(*[]*Job)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jobR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jobR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`job_warnings`),
		qm.WhereIn(`job_warnings.fk_jobs_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load job_warnings")
	}

	var resultSlice []*JobWarning
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice job_warnings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on job_warnings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for job_warnings")
	}

	if singular {
		object.R.FKJobJobWarnings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &jobWarningR{}
			}
			foreign.R.FKJob = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.FKJobsID) {
				local.R.FKJobJobWarnings = append(local.R.FKJobJobWarnings, foreign)
				if foreign.R == nil {
					foreign.R = &jobWarningR{}
				}
				foreign.R.FKJob = local
				break
			}
		}
	}

	return nil
}

// AddFKJobJobWarnings adds the given related objects to the existing relationships
// of the job, optionally inserting them as new records.
// Appends related to o.R.FKJobJobWarnings.
// Sets related.R.FKJob appropriately.
func (q jobQuery) AddFKJobJobWarnings(o *Job, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*JobWarning) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.FKJobsID, o.ID)
			if err = JobWarnings().Insert(rel, ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"job_warnings\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"fk_jobs_id"}),
				strmangle.WhereClause("\"", "\"", 2, jobWarningPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.FKJobsID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &jobR{
			FKJobJobWarnings: related,
		}
	} else {
		o.R.FKJobJobWarnings = append(o.R.FKJobJobWarnings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jobWarningR{
				FKJob: o,
			}
		} else {
			rel.R.FKJob = o
		}
	}
	return nil
}

// SetFKJobJobWarnings removes all previously related items of the
// job replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.FKJob's FKJobJobWarnings accordingly.
// Replaces o.R.FKJobJobWarnings with related.
// Sets related.R.FKJob's FKJobJobWarnings accordingly.
func (q jobQuery) SetFKJobJobWarnings(o *Job, ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*JobWarning) error {
	query := "update \"job_warnings\" set \"fk_jobs_id\" = null where \"fk_jobs_id\" = $1"
	values := []interface{}{o.ID}
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
		for _, rel := range o.R.FKJobJobWarnings {
			queries.SetScanner(&rel.FKJobsID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.FKJob = nil
		}

		o.R.FKJobJobWarnings = nil
	}
	return q.AddFKJobJobWarnings(o, ctx, exec, insert, related...)
}

// RemoveFKJobJobWarnings relationships from objects passed in.
// Removes related items from R.FKJobJobWarnings (uses pointer comparison, removal does not keep order)
// Sets related.R.FKJob.
func (q jobQuery) RemoveFKJobJobWarnings(o *Job, ctx context.Context, exec boil.ContextExecutor, related ...*JobWarning) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.FKJobsID, nil)
		if rel.R != nil {
			rel.R.FKJob = nil
		}
		if _, err = JobWarnings().Update(rel, ctx, exec, boil.Whitelist("fk_jobs_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.FKJobJobWarnings {
			if rel != ri {
				continue
			}

			ln := len(o.R.FKJobJobWarnings)
			if ln > 1 && i < ln-1 {
				o.R.FKJobJobWarnings[i] = o.R.FKJobJobWarnings[ln-1]
			}
			o.R.FKJobJobWarnings = o.R.FKJobJobWarnings[:ln-1]
			break
		}
	}

	return nil
}

// Jobs retrieves all the records using an executor.
func Jobs(mods ...qm.QueryMod) jobQuery {
	mods = append(mods, qm.From("\"jobs\""))
	return jobQuery{NewQuery(mods...)}
}

type JobFinder interface {
	FindJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Job, error)
}

// FindJob retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJob(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Job, error) {
	jobObj := &Job{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"jobs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, jobObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from jobs")
	}

	return jobObj, nil
}

type JobInserter interface {
	Insert(o *Job, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q jobQuery) Insert(o *Job, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no jobs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(jobColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	jobInsertCacheMut.RLock()
	cache, cached := jobInsertCache[key]
	jobInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			jobAllColumns,
			jobColumnsWithDefault,
			jobColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jobType, jobMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"jobs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"jobs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into jobs")
	}

	if !cached {
		jobInsertCacheMut.Lock()
		jobInsertCache[key] = cache
		jobInsertCacheMut.Unlock()
	}

	return nil
}

type JobUpdater interface {
	Update(o *Job, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o JobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the Job.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q jobQuery) Update(o *Job, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	jobUpdateCacheMut.RLock()
	cache, cached := jobUpdateCache[key]
	jobUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			jobAllColumns,
			jobPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update jobs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"jobs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, jobPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, append(wl, jobPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update jobs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for jobs")
	}

	if !cached {
		jobUpdateCacheMut.Lock()
		jobUpdateCache[key] = cache
		jobUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q jobQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for jobs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q jobQuery) UpdateAllSlice(o JobSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"jobs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, jobPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in job slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all job")
	}
	return rowsAff, nil
}

type JobDeleter interface {
	Delete(o *Job, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o JobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single Job record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q jobQuery) Delete(o *Job, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Job provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jobPrimaryKeyMapping)
	sql := "DELETE FROM \"jobs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q jobQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no jobQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from jobs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for jobs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q jobQuery) DeleteAllSlice(o JobSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"jobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from job slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for jobs")
	}

	return rowsAff, nil
}

type JobReloader interface {
	Reload(o *Job, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *JobSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q jobQuery) Reload(o *Job, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJob(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q jobQuery) ReloadAll(o *JobSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JobSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"jobs\".* FROM \"jobs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, jobPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JobSlice")
	}

	*o = slice

	return nil
}

// JobExists checks if the Job row exists.
func JobExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"jobs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if jobs exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Job) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no jobs provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(jobColumnsWithDefault, o)

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

	jobUpsertCacheMut.RLock()
	cache, cached := jobUpsertCache[key]
	jobUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			jobAllColumns,
			jobColumnsWithDefault,
			jobColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			jobAllColumns,
			jobPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert jobs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(jobPrimaryKeyColumns))
			copy(conflict, jobPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"jobs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(jobType, jobMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jobType, jobMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert jobs")
	}

	if !cached {
		jobUpsertCacheMut.Lock()
		jobUpsertCache[key] = cache
		jobUpsertCacheMut.Unlock()
	}

	return nil
}
