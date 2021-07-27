// +build integration postgres
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

func testBuildpackLifecycleData(t *testing.T) {
	t.Parallel()

	query := BuildpackLifecycleData()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildpackLifecycleDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
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

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BuildpackLifecycleData().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackLifecycleDatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildpackLifecycleDatumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BuildpackLifecycleDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildpackLifecycleDatumExists to return true, but got false.")
	}
}

func testBuildpackLifecycleDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildpackLifecycleDatumFound, err := FindBuildpackLifecycleDatum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildpackLifecycleDatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildpackLifecycleDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BuildpackLifecycleData().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildpackLifecycleDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BuildpackLifecycleData().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildpackLifecycleDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildpackLifecycleDatumOne := &BuildpackLifecycleDatum{}
	buildpackLifecycleDatumTwo := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, buildpackLifecycleDatumOne, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackLifecycleDatumTwo, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackLifecycleDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackLifecycleDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackLifecycleData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildpackLifecycleDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildpackLifecycleDatumOne := &BuildpackLifecycleDatum{}
	buildpackLifecycleDatumTwo := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, buildpackLifecycleDatumOne, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackLifecycleDatumTwo, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackLifecycleDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackLifecycleDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildpackLifecycleDatumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func buildpackLifecycleDatumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleDatum) error {
	*o = BuildpackLifecycleDatum{}
	return nil
}

func testBuildpackLifecycleDataHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BuildpackLifecycleDatum{}
	o := &BuildpackLifecycleDatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum object: %s", err)
	}

	AddBuildpackLifecycleDatumHook(boil.BeforeInsertHook, buildpackLifecycleDatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumBeforeInsertHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.AfterInsertHook, buildpackLifecycleDatumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumAfterInsertHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.AfterSelectHook, buildpackLifecycleDatumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumAfterSelectHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.BeforeUpdateHook, buildpackLifecycleDatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumBeforeUpdateHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.AfterUpdateHook, buildpackLifecycleDatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumAfterUpdateHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.BeforeDeleteHook, buildpackLifecycleDatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumBeforeDeleteHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.AfterDeleteHook, buildpackLifecycleDatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumAfterDeleteHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.BeforeUpsertHook, buildpackLifecycleDatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumBeforeUpsertHooks = []BuildpackLifecycleDatumHook{}

	AddBuildpackLifecycleDatumHook(boil.AfterUpsertHook, buildpackLifecycleDatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleDatumAfterUpsertHooks = []BuildpackLifecycleDatumHook{}
}

func testBuildpackLifecycleDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackLifecycleDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildpackLifecycleDatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackLifecycleDatumToManyBuildpackLifecycleBuildpacks(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleDatum
	var b, c BuildpackLifecycleBuildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.BuildpackLifecycleDataGUID, a.GUID)
	queries.Assign(&c.BuildpackLifecycleDataGUID, a.GUID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.BuildpackLifecycleBuildpacks().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.BuildpackLifecycleDataGUID, b.BuildpackLifecycleDataGUID) {
			bFound = true
		}
		if queries.Equal(v.BuildpackLifecycleDataGUID, c.BuildpackLifecycleDataGUID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BuildpackLifecycleDatumSlice{&a}
	if err = a.L.LoadBuildpackLifecycleBuildpacks(ctx, tx, false, (*[]*BuildpackLifecycleDatum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildpackLifecycleBuildpacks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.BuildpackLifecycleBuildpacks = nil
	if err = a.L.LoadBuildpackLifecycleBuildpacks(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.BuildpackLifecycleBuildpacks); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBuildpackLifecycleDatumToManyAddOpBuildpackLifecycleBuildpacks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleDatum
	var b, c, d, e BuildpackLifecycleBuildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLifecycleBuildpack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLifecycleBuildpackDBTypes, false, strmangle.SetComplement(buildpackLifecycleBuildpackPrimaryKeyColumns, buildpackLifecycleBuildpackColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BuildpackLifecycleBuildpack{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddBuildpackLifecycleBuildpacks(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.GUID, first.BuildpackLifecycleDataGUID) {
			t.Error("foreign key was wrong value", a.GUID, first.BuildpackLifecycleDataGUID)
		}
		if !queries.Equal(a.GUID, second.BuildpackLifecycleDataGUID) {
			t.Error("foreign key was wrong value", a.GUID, second.BuildpackLifecycleDataGUID)
		}

		if first.R.BuildpackLifecycleDatum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.BuildpackLifecycleDatum != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.BuildpackLifecycleBuildpacks[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.BuildpackLifecycleBuildpacks[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.BuildpackLifecycleBuildpacks().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBuildpackLifecycleDatumToManySetOpBuildpackLifecycleBuildpacks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleDatum
	var b, c, d, e BuildpackLifecycleBuildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLifecycleBuildpack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLifecycleBuildpackDBTypes, false, strmangle.SetComplement(buildpackLifecycleBuildpackPrimaryKeyColumns, buildpackLifecycleBuildpackColumnsWithoutDefault)...); err != nil {
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

	err = a.SetBuildpackLifecycleBuildpacks(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetBuildpackLifecycleBuildpacks(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BuildpackLifecycleDataGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BuildpackLifecycleDataGUID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.GUID, d.BuildpackLifecycleDataGUID) {
		t.Error("foreign key was wrong value", a.GUID, d.BuildpackLifecycleDataGUID)
	}
	if !queries.Equal(a.GUID, e.BuildpackLifecycleDataGUID) {
		t.Error("foreign key was wrong value", a.GUID, e.BuildpackLifecycleDataGUID)
	}

	if b.R.BuildpackLifecycleDatum != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BuildpackLifecycleDatum != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BuildpackLifecycleDatum != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.BuildpackLifecycleDatum != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.BuildpackLifecycleBuildpacks[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.BuildpackLifecycleBuildpacks[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBuildpackLifecycleDatumToManyRemoveOpBuildpackLifecycleBuildpacks(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleDatum
	var b, c, d, e BuildpackLifecycleBuildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLifecycleBuildpack{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLifecycleBuildpackDBTypes, false, strmangle.SetComplement(buildpackLifecycleBuildpackPrimaryKeyColumns, buildpackLifecycleBuildpackColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddBuildpackLifecycleBuildpacks(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveBuildpackLifecycleBuildpacks(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.BuildpackLifecycleDataGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.BuildpackLifecycleDataGUID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.BuildpackLifecycleDatum != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.BuildpackLifecycleDatum != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.BuildpackLifecycleDatum != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.BuildpackLifecycleDatum != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.BuildpackLifecycleBuildpacks) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.BuildpackLifecycleBuildpacks[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.BuildpackLifecycleBuildpacks[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBuildpackLifecycleDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
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

func testBuildpackLifecycleDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackLifecycleDatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildpackLifecycleDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackLifecycleData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildpackLifecycleDatumDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `AppGUID`: `text`, `DropletGUID`: `text`, `Stack`: `text`, `EncryptedBuildpackURL`: `text`, `EncryptedBuildpackURLSalt`: `text`, `AdminBuildpackName`: `text`, `BuildGUID`: `text`, `EncryptionKeyLabel`: `character varying`, `EncryptionIterations`: `integer`}
	_                              = bytes.MinRead
)

func testBuildpackLifecycleDataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildpackLifecycleDatumAllColumns) == len(buildpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildpackLifecycleDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildpackLifecycleDatumAllColumns) == len(buildpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackLifecycleDatumDBTypes, true, buildpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildpackLifecycleDatumAllColumns, buildpackLifecycleDatumPrimaryKeyColumns) {
		fields = buildpackLifecycleDatumAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildpackLifecycleDatumAllColumns,
			buildpackLifecycleDatumPrimaryKeyColumns,
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

	slice := BuildpackLifecycleDatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildpackLifecycleDataUpsert(t *testing.T) {
	t.Parallel()

	if len(buildpackLifecycleDatumAllColumns) == len(buildpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BuildpackLifecycleDatum{}
	if err = randomize.Struct(seed, &o, buildpackLifecycleDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackLifecycleDatum: %s", err)
	}

	count, err := BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackLifecycleDatum: %s", err)
	}

	count, err = BuildpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}