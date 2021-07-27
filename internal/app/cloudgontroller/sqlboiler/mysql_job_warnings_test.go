// +build integration mysql
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

func testJobWarnings(t *testing.T) {
	t.Parallel()

	query := JobWarnings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJobWarningsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
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

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobWarningsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := JobWarnings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobWarningsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JobWarningSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobWarningsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JobWarningExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if JobWarning exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JobWarningExists to return true, but got false.")
	}
}

func testJobWarningsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	jobWarningFound, err := FindJobWarning(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if jobWarningFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJobWarningsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = JobWarnings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJobWarningsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := JobWarnings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJobWarningsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jobWarningOne := &JobWarning{}
	jobWarningTwo := &JobWarning{}
	if err = randomize.Struct(seed, jobWarningOne, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}
	if err = randomize.Struct(seed, jobWarningTwo, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jobWarningOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jobWarningTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := JobWarnings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJobWarningsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jobWarningOne := &JobWarning{}
	jobWarningTwo := &JobWarning{}
	if err = randomize.Struct(seed, jobWarningOne, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}
	if err = randomize.Struct(seed, jobWarningTwo, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jobWarningOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jobWarningTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func jobWarningBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func jobWarningAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *JobWarning) error {
	*o = JobWarning{}
	return nil
}

func testJobWarningsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &JobWarning{}
	o := &JobWarning{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jobWarningDBTypes, false); err != nil {
		t.Errorf("Unable to randomize JobWarning object: %s", err)
	}

	AddJobWarningHook(boil.BeforeInsertHook, jobWarningBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jobWarningBeforeInsertHooks = []JobWarningHook{}

	AddJobWarningHook(boil.AfterInsertHook, jobWarningAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jobWarningAfterInsertHooks = []JobWarningHook{}

	AddJobWarningHook(boil.AfterSelectHook, jobWarningAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jobWarningAfterSelectHooks = []JobWarningHook{}

	AddJobWarningHook(boil.BeforeUpdateHook, jobWarningBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jobWarningBeforeUpdateHooks = []JobWarningHook{}

	AddJobWarningHook(boil.AfterUpdateHook, jobWarningAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jobWarningAfterUpdateHooks = []JobWarningHook{}

	AddJobWarningHook(boil.BeforeDeleteHook, jobWarningBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jobWarningBeforeDeleteHooks = []JobWarningHook{}

	AddJobWarningHook(boil.AfterDeleteHook, jobWarningAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jobWarningAfterDeleteHooks = []JobWarningHook{}

	AddJobWarningHook(boil.BeforeUpsertHook, jobWarningBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jobWarningBeforeUpsertHooks = []JobWarningHook{}

	AddJobWarningHook(boil.AfterUpsertHook, jobWarningAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jobWarningAfterUpsertHooks = []JobWarningHook{}
}

func testJobWarningsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJobWarningsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(jobWarningColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJobWarningToOneJobUsingFKJob(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local JobWarning
	var foreign Job

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, jobDBTypes, false, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.FKJobsID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.FKJob().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := JobWarningSlice{&local}
	if err = local.L.LoadFKJob(ctx, tx, false, (*[]*JobWarning)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FKJob == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.FKJob = nil
	if err = local.L.LoadFKJob(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.FKJob == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testJobWarningToOneSetOpJobUsingFKJob(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a JobWarning
	var b, c Job

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobWarningDBTypes, false, strmangle.SetComplement(jobWarningPrimaryKeyColumns, jobWarningColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Job{&b, &c} {
		err = a.SetFKJob(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.FKJob != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.FKJobJobWarnings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.FKJobsID, x.ID) {
			t.Error("foreign key was wrong value", a.FKJobsID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.FKJobsID))
		reflect.Indirect(reflect.ValueOf(&a.FKJobsID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.FKJobsID, x.ID) {
			t.Error("foreign key was wrong value", a.FKJobsID, x.ID)
		}
	}
}

func testJobWarningToOneRemoveOpJobUsingFKJob(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a JobWarning
	var b Job

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobWarningDBTypes, false, strmangle.SetComplement(jobWarningPrimaryKeyColumns, jobWarningColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetFKJob(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveFKJob(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.FKJob().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.FKJob != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.FKJobsID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.FKJobJobWarnings) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testJobWarningsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
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

func testJobWarningsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JobWarningSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJobWarningsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := JobWarnings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jobWarningDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `Detail`: `varchar`, `JobID`: `int`, `FKJobsID`: `int`}
	_                 = bytes.MinRead
)

func testJobWarningsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(jobWarningPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(jobWarningAllColumns) == len(jobWarningPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJobWarningsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jobWarningAllColumns) == len(jobWarningPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &JobWarning{}
	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jobWarningDBTypes, true, jobWarningPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jobWarningAllColumns, jobWarningPrimaryKeyColumns) {
		fields = jobWarningAllColumns
	} else {
		fields = strmangle.SetComplement(
			jobWarningAllColumns,
			jobWarningPrimaryKeyColumns,
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

	slice := JobWarningSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJobWarningsUpsert(t *testing.T) {
	t.Parallel()

	if len(jobWarningAllColumns) == len(jobWarningPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLJobWarningUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := JobWarning{}
	if err = randomize.Struct(seed, &o, jobWarningDBTypes, false); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert JobWarning: %s", err)
	}

	count, err := JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, jobWarningDBTypes, false, jobWarningPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JobWarning struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert JobWarning: %s", err)
	}

	count, err = JobWarnings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
