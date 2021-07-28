// +build psql,db
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

func testBuildpackLifecycleBuildpacks(t *testing.T) {
	t.Parallel()

	query := BuildpackLifecycleBuildpacks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildpackLifecycleBuildpacksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
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

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleBuildpacksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BuildpackLifecycleBuildpacks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleBuildpacksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackLifecycleBuildpackSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackLifecycleBuildpacksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildpackLifecycleBuildpackExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BuildpackLifecycleBuildpack exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildpackLifecycleBuildpackExists to return true, but got false.")
	}
}

func testBuildpackLifecycleBuildpacksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildpackLifecycleBuildpackFound, err := FindBuildpackLifecycleBuildpack(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildpackLifecycleBuildpackFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildpackLifecycleBuildpacksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BuildpackLifecycleBuildpacks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildpackLifecycleBuildpacksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BuildpackLifecycleBuildpacks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildpackLifecycleBuildpacksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildpackLifecycleBuildpackOne := &BuildpackLifecycleBuildpack{}
	buildpackLifecycleBuildpackTwo := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, buildpackLifecycleBuildpackOne, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackLifecycleBuildpackTwo, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackLifecycleBuildpackOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackLifecycleBuildpackTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackLifecycleBuildpacks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildpackLifecycleBuildpacksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildpackLifecycleBuildpackOne := &BuildpackLifecycleBuildpack{}
	buildpackLifecycleBuildpackTwo := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, buildpackLifecycleBuildpackOne, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackLifecycleBuildpackTwo, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackLifecycleBuildpackOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackLifecycleBuildpackTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildpackLifecycleBuildpackBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func buildpackLifecycleBuildpackAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackLifecycleBuildpack) error {
	*o = BuildpackLifecycleBuildpack{}
	return nil
}

func testBuildpackLifecycleBuildpacksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BuildpackLifecycleBuildpack{}
	o := &BuildpackLifecycleBuildpack{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack object: %s", err)
	}

	AddBuildpackLifecycleBuildpackHook(boil.BeforeInsertHook, buildpackLifecycleBuildpackBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackBeforeInsertHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.AfterInsertHook, buildpackLifecycleBuildpackAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackAfterInsertHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.AfterSelectHook, buildpackLifecycleBuildpackAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackAfterSelectHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.BeforeUpdateHook, buildpackLifecycleBuildpackBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackBeforeUpdateHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.AfterUpdateHook, buildpackLifecycleBuildpackAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackAfterUpdateHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.BeforeDeleteHook, buildpackLifecycleBuildpackBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackBeforeDeleteHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.AfterDeleteHook, buildpackLifecycleBuildpackAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackAfterDeleteHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.BeforeUpsertHook, buildpackLifecycleBuildpackBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackBeforeUpsertHooks = []BuildpackLifecycleBuildpackHook{}

	AddBuildpackLifecycleBuildpackHook(boil.AfterUpsertHook, buildpackLifecycleBuildpackAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackLifecycleBuildpackAfterUpsertHooks = []BuildpackLifecycleBuildpackHook{}
}

func testBuildpackLifecycleBuildpacksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackLifecycleBuildpacksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildpackLifecycleBuildpackColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackLifecycleBuildpackToOneBuildpackLifecycleDatumUsingBuildpackLifecycleDatum(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BuildpackLifecycleBuildpack
	var foreign BuildpackLifecycleDatum

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildpackLifecycleDatumDBTypes, false, buildpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleDatum struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BuildpackLifecycleDataGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.BuildpackLifecycleDatum().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := BuildpackLifecycleBuildpackSlice{&local}
	if err = local.L.LoadBuildpackLifecycleDatum(ctx, tx, false, (*[]*BuildpackLifecycleBuildpack)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BuildpackLifecycleDatum == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BuildpackLifecycleDatum = nil
	if err = local.L.LoadBuildpackLifecycleDatum(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.BuildpackLifecycleDatum == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBuildpackLifecycleBuildpackToOneSetOpBuildpackLifecycleDatumUsingBuildpackLifecycleDatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleBuildpack
	var b, c BuildpackLifecycleDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleBuildpackDBTypes, false, strmangle.SetComplement(buildpackLifecycleBuildpackPrimaryKeyColumns, buildpackLifecycleBuildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*BuildpackLifecycleDatum{&b, &c} {
		err = a.SetBuildpackLifecycleDatum(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BuildpackLifecycleDatum != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BuildpackLifecycleBuildpacks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BuildpackLifecycleDataGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.BuildpackLifecycleDataGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BuildpackLifecycleDataGUID))
		reflect.Indirect(reflect.ValueOf(&a.BuildpackLifecycleDataGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BuildpackLifecycleDataGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.BuildpackLifecycleDataGUID, x.GUID)
		}
	}
}

func testBuildpackLifecycleBuildpackToOneRemoveOpBuildpackLifecycleDatumUsingBuildpackLifecycleDatum(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackLifecycleBuildpack
	var b BuildpackLifecycleDatum

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackLifecycleBuildpackDBTypes, false, strmangle.SetComplement(buildpackLifecycleBuildpackPrimaryKeyColumns, buildpackLifecycleBuildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildpackLifecycleDatumDBTypes, false, strmangle.SetComplement(buildpackLifecycleDatumPrimaryKeyColumns, buildpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBuildpackLifecycleDatum(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBuildpackLifecycleDatum(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.BuildpackLifecycleDatum().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.BuildpackLifecycleDatum != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BuildpackLifecycleDataGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BuildpackLifecycleBuildpacks) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBuildpackLifecycleBuildpacksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
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

func testBuildpackLifecycleBuildpacksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackLifecycleBuildpackSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildpackLifecycleBuildpacksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackLifecycleBuildpacks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildpackLifecycleBuildpackDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `AdminBuildpackName`: `text`, `EncryptedBuildpackURL`: `character varying`, `EncryptedBuildpackURLSalt`: `text`, `BuildpackLifecycleDataGUID`: `text`, `EncryptionKeyLabel`: `character varying`, `Version`: `character varying`, `BuildpackName`: `character varying`, `EncryptionIterations`: `integer`}
	_                                  = bytes.MinRead
)

func testBuildpackLifecycleBuildpacksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildpackLifecycleBuildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildpackLifecycleBuildpackAllColumns) == len(buildpackLifecycleBuildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildpackLifecycleBuildpacksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildpackLifecycleBuildpackAllColumns) == len(buildpackLifecycleBuildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackLifecycleBuildpackDBTypes, true, buildpackLifecycleBuildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildpackLifecycleBuildpackAllColumns, buildpackLifecycleBuildpackPrimaryKeyColumns) {
		fields = buildpackLifecycleBuildpackAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildpackLifecycleBuildpackAllColumns,
			buildpackLifecycleBuildpackPrimaryKeyColumns,
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

	slice := BuildpackLifecycleBuildpackSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildpackLifecycleBuildpacksUpsert(t *testing.T) {
	t.Parallel()

	if len(buildpackLifecycleBuildpackAllColumns) == len(buildpackLifecycleBuildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BuildpackLifecycleBuildpack{}
	if err = randomize.Struct(seed, &o, buildpackLifecycleBuildpackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackLifecycleBuildpack: %s", err)
	}

	count, err := BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildpackLifecycleBuildpackDBTypes, false, buildpackLifecycleBuildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackLifecycleBuildpack struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackLifecycleBuildpack: %s", err)
	}

	count, err = BuildpackLifecycleBuildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
