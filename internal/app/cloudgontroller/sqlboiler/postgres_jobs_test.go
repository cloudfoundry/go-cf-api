// +build integration,postgres
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

func testJobs(t *testing.T) {
	t.Parallel()

	query := Jobs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJobsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
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

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Jobs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JobSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJobsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JobExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Job exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JobExists to return true, but got false.")
	}
}

func testJobsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	jobFound, err := FindJob(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if jobFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJobsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Jobs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJobsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Jobs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJobsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jobOne := &Job{}
	jobTwo := &Job{}
	if err = randomize.Struct(seed, jobOne, jobDBTypes, false, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}
	if err = randomize.Struct(seed, jobTwo, jobDBTypes, false, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJobsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jobOne := &Job{}
	jobTwo := &Job{}
	if err = randomize.Struct(seed, jobOne, jobDBTypes, false, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}
	if err = randomize.Struct(seed, jobTwo, jobDBTypes, false, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func jobBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func jobAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Job) error {
	*o = Job{}
	return nil
}

func testJobsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Job{}
	o := &Job{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jobDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Job object: %s", err)
	}

	AddJobHook(boil.BeforeInsertHook, jobBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jobBeforeInsertHooks = []JobHook{}

	AddJobHook(boil.AfterInsertHook, jobAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jobAfterInsertHooks = []JobHook{}

	AddJobHook(boil.AfterSelectHook, jobAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jobAfterSelectHooks = []JobHook{}

	AddJobHook(boil.BeforeUpdateHook, jobBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jobBeforeUpdateHooks = []JobHook{}

	AddJobHook(boil.AfterUpdateHook, jobAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jobAfterUpdateHooks = []JobHook{}

	AddJobHook(boil.BeforeDeleteHook, jobBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jobBeforeDeleteHooks = []JobHook{}

	AddJobHook(boil.AfterDeleteHook, jobAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jobAfterDeleteHooks = []JobHook{}

	AddJobHook(boil.BeforeUpsertHook, jobBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jobBeforeUpsertHooks = []JobHook{}

	AddJobHook(boil.AfterUpsertHook, jobAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jobAfterUpsertHooks = []JobHook{}
}

func testJobsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJobsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(jobColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJobToManyFKJobJobWarnings(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Job
	var b, c JobWarning

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, jobWarningDBTypes, false, jobWarningColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.FKJobsID, a.ID)
	queries.Assign(&c.FKJobsID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FKJobJobWarnings().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.FKJobsID, b.FKJobsID) {
			bFound = true
		}
		if queries.Equal(v.FKJobsID, c.FKJobsID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := JobSlice{&a}
	if err = a.L.LoadFKJobJobWarnings(ctx, tx, false, (*[]*Job)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKJobJobWarnings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FKJobJobWarnings = nil
	if err = a.L.LoadFKJobJobWarnings(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FKJobJobWarnings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testJobToManyAddOpFKJobJobWarnings(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Job
	var b, c, d, e JobWarning

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JobWarning{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jobWarningDBTypes, false, strmangle.SetComplement(jobWarningPrimaryKeyColumns, jobWarningColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*JobWarning{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFKJobJobWarnings(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.FKJobsID) {
			t.Error("foreign key was wrong value", a.ID, first.FKJobsID)
		}
		if !queries.Equal(a.ID, second.FKJobsID) {
			t.Error("foreign key was wrong value", a.ID, second.FKJobsID)
		}

		if first.R.FKJob != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FKJob != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FKJobJobWarnings[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FKJobJobWarnings[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FKJobJobWarnings().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testJobToManySetOpFKJobJobWarnings(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Job
	var b, c, d, e JobWarning

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JobWarning{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jobWarningDBTypes, false, strmangle.SetComplement(jobWarningPrimaryKeyColumns, jobWarningColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetFKJobJobWarnings(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.FKJobJobWarnings().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetFKJobJobWarnings(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.FKJobJobWarnings().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.FKJobsID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.FKJobsID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.FKJobsID) {
		t.Error("foreign key was wrong value", a.ID, d.FKJobsID)
	}
	if !queries.Equal(a.ID, e.FKJobsID) {
		t.Error("foreign key was wrong value", a.ID, e.FKJobsID)
	}

	if b.R.FKJob != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.FKJob != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.FKJob != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.FKJob != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.FKJobJobWarnings[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.FKJobJobWarnings[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testJobToManyRemoveOpFKJobJobWarnings(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Job
	var b, c, d, e JobWarning

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jobDBTypes, false, strmangle.SetComplement(jobPrimaryKeyColumns, jobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*JobWarning{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, jobWarningDBTypes, false, strmangle.SetComplement(jobWarningPrimaryKeyColumns, jobWarningColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddFKJobJobWarnings(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.FKJobJobWarnings().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveFKJobJobWarnings(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.FKJobJobWarnings().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.FKJobsID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.FKJobsID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.FKJob != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.FKJob != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.FKJob != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.FKJob != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.FKJobJobWarnings) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.FKJobJobWarnings[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.FKJobJobWarnings[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testJobsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
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

func testJobsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JobSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJobsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jobDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `State`: `text`, `Operation`: `text`, `ResourceGUID`: `text`, `ResourceType`: `text`, `DelayedJobGUID`: `text`, `CFAPIError`: `character varying`}
	_          = bytes.MinRead
)

func testJobsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(jobPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(jobAllColumns) == len(jobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jobDBTypes, true, jobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJobsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jobAllColumns) == len(jobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Job{}
	if err = randomize.Struct(seed, o, jobDBTypes, true, jobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jobDBTypes, true, jobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jobAllColumns, jobPrimaryKeyColumns) {
		fields = jobAllColumns
	} else {
		fields = strmangle.SetComplement(
			jobAllColumns,
			jobPrimaryKeyColumns,
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

	slice := JobSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJobsUpsert(t *testing.T) {
	t.Parallel()

	if len(jobAllColumns) == len(jobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Job{}
	if err = randomize.Struct(seed, &o, jobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Job: %s", err)
	}

	count, err := Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, jobDBTypes, false, jobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Job struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Job: %s", err)
	}

	count, err = Jobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
