// +build psql_integration
// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testClockJobs(t *testing.T) {
	t.Parallel()

	query := ClockJobs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testClockJobsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClockJobsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ClockJobs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClockJobsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClockJobSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testClockJobsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ClockJobExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ClockJob exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ClockJobExists to return true, but got false.")
	}
}

func testClockJobsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	clockJobFound, err := FindClockJob(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if clockJobFound == nil {
		t.Error("want a record, got nil")
	}
}

func testClockJobsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ClockJobs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testClockJobsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ClockJobs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testClockJobsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	clockJobOne := &ClockJob{}
	clockJobTwo := &ClockJob{}
	if err = randomize.Struct(seed, clockJobOne, clockJobDBTypes, false, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}
	if err = randomize.Struct(seed, clockJobTwo, clockJobDBTypes, false, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clockJobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clockJobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClockJobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testClockJobsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	clockJobOne := &ClockJob{}
	clockJobTwo := &ClockJob{}
	if err = randomize.Struct(seed, clockJobOne, clockJobDBTypes, false, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}
	if err = randomize.Struct(seed, clockJobTwo, clockJobDBTypes, false, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = clockJobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = clockJobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func clockJobBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func clockJobAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ClockJob) error {
	*o = ClockJob{}
	return nil
}

func testClockJobsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ClockJob{}
	o := &ClockJob{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, clockJobDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ClockJob object: %s", err)
	}

	AddClockJobHook(boil.BeforeInsertHook, clockJobBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	clockJobBeforeInsertHooks = []ClockJobHook{}

	AddClockJobHook(boil.AfterInsertHook, clockJobAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	clockJobAfterInsertHooks = []ClockJobHook{}

	AddClockJobHook(boil.AfterSelectHook, clockJobAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	clockJobAfterSelectHooks = []ClockJobHook{}

	AddClockJobHook(boil.BeforeUpdateHook, clockJobBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	clockJobBeforeUpdateHooks = []ClockJobHook{}

	AddClockJobHook(boil.AfterUpdateHook, clockJobAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	clockJobAfterUpdateHooks = []ClockJobHook{}

	AddClockJobHook(boil.BeforeDeleteHook, clockJobBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	clockJobBeforeDeleteHooks = []ClockJobHook{}

	AddClockJobHook(boil.AfterDeleteHook, clockJobAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	clockJobAfterDeleteHooks = []ClockJobHook{}

	AddClockJobHook(boil.BeforeUpsertHook, clockJobBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	clockJobBeforeUpsertHooks = []ClockJobHook{}

	AddClockJobHook(boil.AfterUpsertHook, clockJobAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	clockJobAfterUpsertHooks = []ClockJobHook{}
}

func testClockJobsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClockJobsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(clockJobColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testClockJobsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClockJobsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ClockJobSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testClockJobsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ClockJobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	clockJobDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`, `LastStartedAt`: `timestamp without time zone`, `LastCompletedAt`: `timestamp without time zone`}
	_               = bytes.MinRead
)

func testClockJobsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(clockJobPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(clockJobAllColumns) == len(clockJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testClockJobsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(clockJobAllColumns) == len(clockJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ClockJob{}
	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, clockJobDBTypes, true, clockJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(clockJobAllColumns, clockJobPrimaryKeyColumns) {
		fields = clockJobAllColumns
	} else {
		fields = strmangle.SetComplement(
			clockJobAllColumns,
			clockJobPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ClockJobSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testClockJobsUpsert(t *testing.T) {
	t.Parallel()

	if len(clockJobAllColumns) == len(clockJobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ClockJob{}
	if err = randomize.Struct(seed, &o, clockJobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClockJob: %s", err)
	}

	count, err := ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, clockJobDBTypes, false, clockJobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ClockJob struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ClockJob: %s", err)
	}

	count, err = ClockJobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}