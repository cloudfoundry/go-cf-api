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

type JobWarningUpserter interface {
	Upsert(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error
}

var mySQLJobWarningUniqueColumns = []string{
	"id",
	"guid",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (q jobWarningQuery) Upsert(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job_warnings provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(jobWarningColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLJobWarningUniqueColumns, o)

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

	jobWarningUpsertCacheMut.RLock()
	cache, cached := jobWarningUpsertCache[key]
	jobWarningUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			jobWarningAllColumns,
			jobWarningColumnsWithDefault,
			jobWarningColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			jobWarningAllColumns,
			jobWarningPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert job_warnings, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`job_warnings`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `job_warnings` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(jobWarningType, jobWarningMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(jobWarningType, jobWarningMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for job_warnings")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == jobWarningMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(jobWarningType, jobWarningMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for job_warnings")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for job_warnings")
	}

CacheNoHooks:
	if !cached {
		jobWarningUpsertCacheMut.Lock()
		jobWarningUpsertCache[key] = cache
		jobWarningUpsertCacheMut.Unlock()
	}

	return nil
}

// JobWarning is an object representing the database table.
type JobWarning struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	GUID      string    `boil:"guid" json:"guid" toml:"guid" yaml:"guid"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	Detail    string    `boil:"detail" json:"detail" toml:"detail" yaml:"detail"`
	JobID     int       `boil:"job_id" json:"job_id" toml:"job_id" yaml:"job_id"`
	FKJobsID  null.Int  `boil:"fk_jobs_id" json:"fk_jobs_id,omitempty" toml:"fk_jobs_id" yaml:"fk_jobs_id,omitempty"`

	R *jobWarningR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L jobWarningL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JobWarningColumns = struct {
	ID        string
	GUID      string
	CreatedAt string
	UpdatedAt string
	Detail    string
	JobID     string
	FKJobsID  string
}{
	ID:        "id",
	GUID:      "guid",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	Detail:    "detail",
	JobID:     "job_id",
	FKJobsID:  "fk_jobs_id",
}

var JobWarningTableColumns = struct {
	ID        string
	GUID      string
	CreatedAt string
	UpdatedAt string
	Detail    string
	JobID     string
	FKJobsID  string
}{
	ID:        "job_warnings.id",
	GUID:      "job_warnings.guid",
	CreatedAt: "job_warnings.created_at",
	UpdatedAt: "job_warnings.updated_at",
	Detail:    "job_warnings.detail",
	JobID:     "job_warnings.job_id",
	FKJobsID:  "job_warnings.fk_jobs_id",
}

// Generated where

var JobWarningWhere = struct {
	ID        whereHelperint
	GUID      whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpernull_Time
	Detail    whereHelperstring
	JobID     whereHelperint
	FKJobsID  whereHelpernull_Int
}{
	ID:        whereHelperint{field: "`job_warnings`.`id`"},
	GUID:      whereHelperstring{field: "`job_warnings`.`guid`"},
	CreatedAt: whereHelpertime_Time{field: "`job_warnings`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`job_warnings`.`updated_at`"},
	Detail:    whereHelperstring{field: "`job_warnings`.`detail`"},
	JobID:     whereHelperint{field: "`job_warnings`.`job_id`"},
	FKJobsID:  whereHelpernull_Int{field: "`job_warnings`.`fk_jobs_id`"},
}

// JobWarningRels is where relationship names are stored.
var JobWarningRels = struct {
	FKJob string
}{
	FKJob: "FKJob",
}

// jobWarningR is where relationships are stored.
type jobWarningR struct {
	FKJob *Job `boil:"FKJob" json:"FKJob" toml:"FKJob" yaml:"FKJob"`
}

// NewStruct creates a new relationship struct
func (*jobWarningR) NewStruct() *jobWarningR {
	return &jobWarningR{}
}

// jobWarningL is where Load methods for each relationship are stored.
type jobWarningL struct{}

var (
	jobWarningAllColumns            = []string{"id", "guid", "created_at", "updated_at", "detail", "job_id", "fk_jobs_id"}
	jobWarningColumnsWithoutDefault = []string{"guid", "updated_at", "detail", "job_id", "fk_jobs_id"}
	jobWarningColumnsWithDefault    = []string{"id", "created_at"}
	jobWarningPrimaryKeyColumns     = []string{"id"}
)

type (
	// JobWarningSlice is an alias for a slice of pointers to JobWarning.
	// This should almost always be used instead of []JobWarning.
	JobWarningSlice []*JobWarning

	jobWarningQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	jobWarningType                 = reflect.TypeOf(&JobWarning{})
	jobWarningMapping              = queries.MakeStructMapping(jobWarningType)
	jobWarningPrimaryKeyMapping, _ = queries.BindMapping(jobWarningType, jobWarningMapping, jobWarningPrimaryKeyColumns)
	jobWarningInsertCacheMut       sync.RWMutex
	jobWarningInsertCache          = make(map[string]insertCache)
	jobWarningUpdateCacheMut       sync.RWMutex
	jobWarningUpdateCache          = make(map[string]updateCache)
	jobWarningUpsertCacheMut       sync.RWMutex
	jobWarningUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

type JobWarningFinisher interface {
	One(ctx context.Context, exec boil.ContextExecutor) (*JobWarning, error)
	Count(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	All(ctx context.Context, exec boil.ContextExecutor) (JobWarningSlice, error)
	Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error)
}

// One returns a single jobWarning record from the query.
func (q jobWarningQuery) One(ctx context.Context, exec boil.ContextExecutor) (*JobWarning, error) {
	o := &JobWarning{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for job_warnings")
	}

	return o, nil
}

// All returns all JobWarning records from the query.
func (q jobWarningQuery) All(ctx context.Context, exec boil.ContextExecutor) (JobWarningSlice, error) {
	var o []*JobWarning

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to JobWarning slice")
	}

	return o, nil
}

// Count returns the count of all JobWarning records in the query.
func (q jobWarningQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count job_warnings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q jobWarningQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if job_warnings exists")
	}

	return count > 0, nil
}

// FKJob pointed to by the foreign key.
func (q jobWarningQuery) FKJob(o *JobWarning, mods ...qm.QueryMod) jobQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.FKJobsID),
	}

	queryMods = append(queryMods, mods...)

	query := Jobs(queryMods...)
	queries.SetFrom(query.Query, "`jobs`")

	return query
}

// LoadFKJob allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (jobWarningL) LoadFKJob(ctx context.Context, e boil.ContextExecutor, singular bool, maybeJobWarning interface{}, mods queries.Applicator) error {
	var slice []*JobWarning
	var object *JobWarning

	if singular {
		object = maybeJobWarning.(*JobWarning)
	} else {
		slice = *maybeJobWarning.(*[]*JobWarning)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &jobWarningR{}
		}
		if !queries.IsNil(object.FKJobsID) {
			args = append(args, object.FKJobsID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &jobWarningR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.FKJobsID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.FKJobsID) {
				args = append(args, obj.FKJobsID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`jobs`),
		qm.WhereIn(`jobs.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Job")
	}

	var resultSlice []*Job
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Job")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for jobs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for jobs")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.FKJob = foreign
		if foreign.R == nil {
			foreign.R = &jobR{}
		}
		foreign.R.FKJobJobWarnings = append(foreign.R.FKJobJobWarnings, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.FKJobsID, foreign.ID) {
				local.R.FKJob = foreign
				if foreign.R == nil {
					foreign.R = &jobR{}
				}
				foreign.R.FKJobJobWarnings = append(foreign.R.FKJobJobWarnings, local)
				break
			}
		}
	}

	return nil
}

// SetFKJob of the jobWarning to the related item.
// Sets o.R.FKJob to related.
// Adds o to related.R.FKJobJobWarnings.
func (q jobWarningQuery) SetFKJob(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, insert bool, related *Job) error {
	var err error
	if insert {
		if err = Jobs().Insert(related, ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `job_warnings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"fk_jobs_id"}),
		strmangle.WhereClause("`", "`", 0, jobWarningPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.FKJobsID, related.ID)
	if o.R == nil {
		o.R = &jobWarningR{
			FKJob: related,
		}
	} else {
		o.R.FKJob = related
	}

	if related.R == nil {
		related.R = &jobR{
			FKJobJobWarnings: JobWarningSlice{o},
		}
	} else {
		related.R.FKJobJobWarnings = append(related.R.FKJobJobWarnings, o)
	}

	return nil
}

// RemoveFKJob relationship.
// Sets o.R.FKJob to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (q jobWarningQuery) RemoveFKJob(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, related *Job) error {
	var err error

	queries.SetScanner(&o.FKJobsID, nil)
	if _, err = q.Update(o, ctx, exec, boil.Whitelist("fk_jobs_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.FKJob = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.FKJobJobWarnings {
		if queries.Equal(o.FKJobsID, ri.FKJobsID) {
			continue
		}

		ln := len(related.R.FKJobJobWarnings)
		if ln > 1 && i < ln-1 {
			related.R.FKJobJobWarnings[i] = related.R.FKJobJobWarnings[ln-1]
		}
		related.R.FKJobJobWarnings = related.R.FKJobJobWarnings[:ln-1]
		break
	}
	return nil
}

// JobWarnings retrieves all the records using an executor.
func JobWarnings(mods ...qm.QueryMod) jobWarningQuery {
	mods = append(mods, qm.From("`job_warnings`"))
	return jobWarningQuery{NewQuery(mods...)}
}

type JobWarningFinder interface {
	FindJobWarning(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*JobWarning, error)
}

// FindJobWarning retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJobWarning(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*JobWarning, error) {
	jobWarningObj := &JobWarning{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `job_warnings` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, jobWarningObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from job_warnings")
	}

	return jobWarningObj, nil
}

type JobWarningInserter interface {
	Insert(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (q jobWarningQuery) Insert(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no job_warnings provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(jobWarningColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	jobWarningInsertCacheMut.RLock()
	cache, cached := jobWarningInsertCache[key]
	jobWarningInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			jobWarningAllColumns,
			jobWarningColumnsWithDefault,
			jobWarningColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(jobWarningType, jobWarningMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(jobWarningType, jobWarningMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `job_warnings` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `job_warnings` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `job_warnings` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, jobWarningPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into job_warnings")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == jobWarningMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for job_warnings")
	}

CacheNoHooks:
	if !cached {
		jobWarningInsertCacheMut.Lock()
		jobWarningInsertCache[key] = cache
		jobWarningInsertCacheMut.Unlock()
	}

	return nil
}

type JobWarningUpdater interface {
	Update(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error)
	UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
	UpdateAllSlice(o JobWarningSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error)
}

// Update uses an executor to update the JobWarning.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (q jobWarningQuery) Update(o *JobWarning, ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	jobWarningUpdateCacheMut.RLock()
	cache, cached := jobWarningUpdateCache[key]
	jobWarningUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			jobWarningAllColumns,
			jobWarningPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update job_warnings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `job_warnings` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, jobWarningPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(jobWarningType, jobWarningMapping, append(wl, jobWarningPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update job_warnings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for job_warnings")
	}

	if !cached {
		jobWarningUpdateCacheMut.Lock()
		jobWarningUpdateCache[key] = cache
		jobWarningUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q jobWarningQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for job_warnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for job_warnings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (q jobWarningQuery) UpdateAllSlice(o JobWarningSlice, ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobWarningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `job_warnings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobWarningPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in jobWarning slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all jobWarning")
	}
	return rowsAff, nil
}

type JobWarningDeleter interface {
	Delete(o *JobWarning, ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error)
	DeleteAllSlice(o JobWarningSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error)
}

// Delete deletes a single JobWarning record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (q jobWarningQuery) Delete(o *JobWarning, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no JobWarning provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), jobWarningPrimaryKeyMapping)
	sql := "DELETE FROM `job_warnings` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from job_warnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for job_warnings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q jobWarningQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no jobWarningQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from job_warnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job_warnings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (q jobWarningQuery) DeleteAllSlice(o JobWarningSlice, ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobWarningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `job_warnings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobWarningPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from jobWarning slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for job_warnings")
	}

	return rowsAff, nil
}

type JobWarningReloader interface {
	Reload(o *JobWarning, ctx context.Context, exec boil.ContextExecutor) error
	ReloadAll(o *JobWarningSlice, ctx context.Context, exec boil.ContextExecutor) error
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (q jobWarningQuery) Reload(o *JobWarning, ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindJobWarning(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (q jobWarningQuery) ReloadAll(o *JobWarningSlice, ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JobWarningSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), jobWarningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `job_warnings`.* FROM `job_warnings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, jobWarningPrimaryKeyColumns, len(*o))

	query := queries.Raw(sql, args...)

	err := query.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JobWarningSlice")
	}

	*o = slice

	return nil
}

// JobWarningExists checks if the JobWarning row exists.
func JobWarningExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `job_warnings` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if job_warnings exists")
	}

	return exists, nil
}
