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

// Pilot is an object representing the database table.
type Pilot struct {
	ID   int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *pilotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pilotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PilotColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

var PilotTableColumns = struct {
	ID   string
	Name string
}{
	ID:   "pilots.id",
	Name: "pilots.name",
}

// Generated where

var PilotWhere = struct {
	ID   whereHelperint64
	Name whereHelperstring
}{
	ID:   whereHelperint64{field: "\"pilots\".\"id\""},
	Name: whereHelperstring{field: "\"pilots\".\"name\""},
}

// PilotRels is where relationship names are stored.
var PilotRels = struct {
	Jets      string
	Languages string
}{
	Jets:      "Jets",
	Languages: "Languages",
}

// pilotR is where relationships are stored.
type pilotR struct {
	Jets      JetSlice      `boil:"Jets" json:"Jets" toml:"Jets" yaml:"Jets"`
	Languages LanguageSlice `boil:"Languages" json:"Languages" toml:"Languages" yaml:"Languages"`
}

// NewStruct creates a new relationship struct
func (*pilotR) NewStruct() *pilotR {
	return &pilotR{}
}

// pilotL is where Load methods for each relationship are stored.
type pilotL struct{}

var (
	pilotAllColumns            = []string{"id", "name"}
	pilotColumnsWithoutDefault = []string{"name"}
	pilotColumnsWithDefault    = []string{"id"}
	pilotPrimaryKeyColumns     = []string{"id"}
)

type (
	// PilotSlice is an alias for a slice of pointers to Pilot.
	// This should almost always be used instead of []Pilot.
	PilotSlice []*Pilot
	// PilotHook is the signature for custom Pilot hook methods
	PilotHook func(context.Context, boil.ContextExecutor, *Pilot) error

	pilotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pilotType                 = reflect.TypeOf(&Pilot{})
	pilotMapping              = queries.MakeStructMapping(pilotType)
	pilotPrimaryKeyMapping, _ = queries.BindMapping(pilotType, pilotMapping, pilotPrimaryKeyColumns)
	pilotInsertCacheMut       sync.RWMutex
	pilotInsertCache          = make(map[string]insertCache)
	pilotUpdateCacheMut       sync.RWMutex
	pilotUpdateCache          = make(map[string]updateCache)
	pilotUpsertCacheMut       sync.RWMutex
	pilotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pilotBeforeInsertHooks []PilotHook
var pilotBeforeUpdateHooks []PilotHook
var pilotBeforeDeleteHooks []PilotHook
var pilotBeforeUpsertHooks []PilotHook

var pilotAfterInsertHooks []PilotHook
var pilotAfterSelectHooks []PilotHook
var pilotAfterUpdateHooks []PilotHook
var pilotAfterDeleteHooks []PilotHook
var pilotAfterUpsertHooks []PilotHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pilot) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pilot) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pilot) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pilot) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pilot) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pilot) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pilot) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pilot) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pilot) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pilotAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPilotHook registers your hook function for all future operations.
func AddPilotHook(hookPoint boil.HookPoint, pilotHook PilotHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pilotBeforeInsertHooks = append(pilotBeforeInsertHooks, pilotHook)
	case boil.BeforeUpdateHook:
		pilotBeforeUpdateHooks = append(pilotBeforeUpdateHooks, pilotHook)
	case boil.BeforeDeleteHook:
		pilotBeforeDeleteHooks = append(pilotBeforeDeleteHooks, pilotHook)
	case boil.BeforeUpsertHook:
		pilotBeforeUpsertHooks = append(pilotBeforeUpsertHooks, pilotHook)
	case boil.AfterInsertHook:
		pilotAfterInsertHooks = append(pilotAfterInsertHooks, pilotHook)
	case boil.AfterSelectHook:
		pilotAfterSelectHooks = append(pilotAfterSelectHooks, pilotHook)
	case boil.AfterUpdateHook:
		pilotAfterUpdateHooks = append(pilotAfterUpdateHooks, pilotHook)
	case boil.AfterDeleteHook:
		pilotAfterDeleteHooks = append(pilotAfterDeleteHooks, pilotHook)
	case boil.AfterUpsertHook:
		pilotAfterUpsertHooks = append(pilotAfterUpsertHooks, pilotHook)
	}
}

// One returns a single pilot record from the query.
func (q pilotQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pilot, error) {
	o := &Pilot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pilots")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Pilot records from the query.
func (q pilotQuery) All(ctx context.Context, exec boil.ContextExecutor) (PilotSlice, error) {
	var o []*Pilot

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pilot slice")
	}

	if len(pilotAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Pilot records in the query.
func (q pilotQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pilots rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pilotQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pilots exists")
	}

	return count > 0, nil
}

// Jets retrieves all the jet's Jets with an executor.
func (o *Pilot) Jets(mods ...qm.QueryMod) jetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"jets\".\"pilot_id\"=?", o.ID),
	)

	query := Jets(queryMods...)
	queries.SetFrom(query.Query, "\"jets\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"jets\".*"})
	}

	return query
}

// Languages retrieves all the language's Languages with an executor.
func (o *Pilot) Languages(mods ...qm.QueryMod) languageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.InnerJoin("\"pilot_languages\" on \"languages\".\"id\" = \"pilot_languages\".\"language_id\""),
		qm.Where("\"pilot_languages\".\"pilot_id\"=?", o.ID),
	)

	query := Languages(queryMods...)
	queries.SetFrom(query.Query, "\"languages\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"languages\".*"})
	}

	return query
}

// LoadJets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadJets(ctx context.Context, e boil.ContextExecutor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
			}

			for _, a := range args {
				if a == obj.ID {
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
		qm.From(`jets`),
		qm.WhereIn(`jets.pilot_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load jets")
	}

	var resultSlice []*Jet
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice jets")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on jets")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for jets")
	}

	if len(jetAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Jets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &jetR{}
			}
			foreign.R.Pilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PilotID {
				local.R.Jets = append(local.R.Jets, foreign)
				if foreign.R == nil {
					foreign.R = &jetR{}
				}
				foreign.R.Pilot = local
				break
			}
		}
	}

	return nil
}

// LoadLanguages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadLanguages(ctx context.Context, e boil.ContextExecutor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
			}

			for _, a := range args {
				if a == obj.ID {
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
		qm.Select("\"languages\".id, \"languages\".language, \"a\".\"pilot_id\""),
		qm.From("\"languages\""),
		qm.InnerJoin("\"pilot_languages\" as \"a\" on \"languages\".\"id\" = \"a\".\"language_id\""),
		qm.WhereIn("\"a\".\"pilot_id\" in ?", args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load languages")
	}

	var resultSlice []*Language

	var localJoinCols []int64
	for results.Next() {
		one := new(Language)
		var localJoinCol int64

		err = results.Scan(&one.ID, &one.Language, &localJoinCol)
		if err != nil {
			return errors.Wrap(err, "failed to scan eager loaded results for languages")
		}
		if err = results.Err(); err != nil {
			return errors.Wrap(err, "failed to plebian-bind eager loaded slice languages")
		}

		resultSlice = append(resultSlice, one)
		localJoinCols = append(localJoinCols, localJoinCol)
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on languages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for languages")
	}

	if len(languageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Languages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &languageR{}
			}
			foreign.R.Pilots = append(foreign.R.Pilots, object)
		}
		return nil
	}

	for i, foreign := range resultSlice {
		localJoinCol := localJoinCols[i]
		for _, local := range slice {
			if local.ID == localJoinCol {
				local.R.Languages = append(local.R.Languages, foreign)
				if foreign.R == nil {
					foreign.R = &languageR{}
				}
				foreign.R.Pilots = append(foreign.R.Pilots, local)
				break
			}
		}
	}

	return nil
}

// AddJets adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Jets.
// Sets related.R.Pilot appropriately.
func (o *Pilot) AddJets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Jet) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PilotID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"jets\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pilot_id"}),
				strmangle.WhereClause("\"", "\"", 2, jetPrimaryKeyColumns),
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

			rel.PilotID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pilotR{
			Jets: related,
		}
	} else {
		o.R.Jets = append(o.R.Jets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &jetR{
				Pilot: o,
			}
		} else {
			rel.R.Pilot = o
		}
	}
	return nil
}

// AddLanguages adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Languages.
// Sets related.R.Pilots appropriately.
func (o *Pilot) AddLanguages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Language) error {
	var err error
	for _, rel := range related {
		if insert {
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		}
	}

	for _, rel := range related {
		query := "insert into \"pilot_languages\" (\"pilot_id\", \"language_id\") values ($1, $2)"
		values := []interface{}{o.ID, rel.ID}

		if boil.IsDebug(ctx) {
			writer := boil.DebugWriterFrom(ctx)
			fmt.Fprintln(writer, query)
			fmt.Fprintln(writer, values)
		}
		_, err = exec.ExecContext(ctx, query, values...)
		if err != nil {
			return errors.Wrap(err, "failed to insert into join table")
		}
	}
	if o.R == nil {
		o.R = &pilotR{
			Languages: related,
		}
	} else {
		o.R.Languages = append(o.R.Languages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &languageR{
				Pilots: PilotSlice{o},
			}
		} else {
			rel.R.Pilots = append(rel.R.Pilots, o)
		}
	}
	return nil
}

// SetLanguages removes all previously related items of the
// pilot replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Pilots's Languages accordingly.
// Replaces o.R.Languages with related.
// Sets related.R.Pilots's Languages accordingly.
func (o *Pilot) SetLanguages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Language) error {
	query := "delete from \"pilot_languages\" where \"pilot_id\" = $1"
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

	removeLanguagesFromPilotsSlice(o, related)
	if o.R != nil {
		o.R.Languages = nil
	}
	return o.AddLanguages(ctx, exec, insert, related...)
}

// RemoveLanguages relationships from objects passed in.
// Removes related items from R.Languages (uses pointer comparison, removal does not keep order)
// Sets related.R.Pilots.
func (o *Pilot) RemoveLanguages(ctx context.Context, exec boil.ContextExecutor, related ...*Language) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	query := fmt.Sprintf(
		"delete from \"pilot_languages\" where \"pilot_id\" = $1 and \"language_id\" in (%s)",
		strmangle.Placeholders(dialect.UseIndexPlaceholders, len(related), 2, 1),
	)
	values := []interface{}{o.ID}
	for _, rel := range related {
		values = append(values, rel.ID)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err = exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}
	removeLanguagesFromPilotsSlice(o, related)
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Languages {
			if rel != ri {
				continue
			}

			ln := len(o.R.Languages)
			if ln > 1 && i < ln-1 {
				o.R.Languages[i] = o.R.Languages[ln-1]
			}
			o.R.Languages = o.R.Languages[:ln-1]
			break
		}
	}

	return nil
}

func removeLanguagesFromPilotsSlice(o *Pilot, related []*Language) {
	for _, rel := range related {
		if rel.R == nil {
			continue
		}
		for i, ri := range rel.R.Pilots {
			if o.ID != ri.ID {
				continue
			}

			ln := len(rel.R.Pilots)
			if ln > 1 && i < ln-1 {
				rel.R.Pilots[i] = rel.R.Pilots[ln-1]
			}
			rel.R.Pilots = rel.R.Pilots[:ln-1]
			break
		}
	}
}

// Pilots retrieves all the records using an executor.
func Pilots(mods ...qm.QueryMod) pilotQuery {
	mods = append(mods, qm.From("\"pilots\""))
	return pilotQuery{NewQuery(mods...)}
}

// FindPilot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPilot(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Pilot, error) {
	pilotObj := &Pilot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pilots\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pilotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pilots")
	}

	if err = pilotObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pilotObj, err
	}

	return pilotObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pilot) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pilotInsertCacheMut.RLock()
	cache, cached := pilotInsertCache[key]
	pilotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pilots\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pilots\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into pilots")
	}

	if !cached {
		pilotInsertCacheMut.Lock()
		pilotInsertCache[key] = cache
		pilotInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Pilot.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pilot) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pilotUpdateCacheMut.RLock()
	cache, cached := pilotUpdateCache[key]
	pilotUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pilots, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pilots\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pilotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, append(wl, pilotPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pilots row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pilots")
	}

	if !cached {
		pilotUpdateCacheMut.Lock()
		pilotUpdateCache[key] = cache
		pilotUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pilotQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pilots")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PilotSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pilots\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pilotPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pilot")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pilot) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)

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

	pilotUpsertCacheMut.RLock()
	cache, cached := pilotUpsertCache[key]
	pilotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert pilots, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(pilotPrimaryKeyColumns))
			copy(conflict, pilotPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pilots\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert pilots")
	}

	if !cached {
		pilotUpsertCacheMut.Lock()
		pilotUpsertCache[key] = cache
		pilotUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Pilot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pilot) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pilot provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pilotPrimaryKeyMapping)
	sql := "DELETE FROM \"pilots\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pilots")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pilotQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pilotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PilotSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pilotBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"pilots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	if len(pilotAfterDeleteHooks) != 0 {
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
func (o *Pilot) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPilot(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PilotSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PilotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pilots\".* FROM \"pilots\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PilotSlice")
	}

	*o = slice

	return nil
}

// PilotExists checks if the Pilot row exists.
func PilotExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pilots\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pilots exists")
	}

	return exists, nil
}
