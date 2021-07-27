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

func testKpackLifecycleData(t *testing.T) {
	t.Parallel()

	query := KpackLifecycleData()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testKpackLifecycleDataDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
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

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKpackLifecycleDataQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := KpackLifecycleData().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKpackLifecycleDataSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KpackLifecycleDatumSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testKpackLifecycleDataExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := KpackLifecycleDatumExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if KpackLifecycleDatum exists: %s", err)
	}
	if !e {
		t.Errorf("Expected KpackLifecycleDatumExists to return true, but got false.")
	}
}

func testKpackLifecycleDataFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	kpackLifecycleDatumFound, err := FindKpackLifecycleDatum(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if kpackLifecycleDatumFound == nil {
		t.Error("want a record, got nil")
	}
}

func testKpackLifecycleDataBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = KpackLifecycleData().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testKpackLifecycleDataOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := KpackLifecycleData().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testKpackLifecycleDataAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	kpackLifecycleDatumOne := &KpackLifecycleDatum{}
	kpackLifecycleDatumTwo := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, kpackLifecycleDatumOne, kpackLifecycleDatumDBTypes, false, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, kpackLifecycleDatumTwo, kpackLifecycleDatumDBTypes, false, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = kpackLifecycleDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = kpackLifecycleDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KpackLifecycleData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testKpackLifecycleDataCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	kpackLifecycleDatumOne := &KpackLifecycleDatum{}
	kpackLifecycleDatumTwo := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, kpackLifecycleDatumOne, kpackLifecycleDatumDBTypes, false, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}
	if err = randomize.Struct(seed, kpackLifecycleDatumTwo, kpackLifecycleDatumDBTypes, false, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = kpackLifecycleDatumOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = kpackLifecycleDatumTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func kpackLifecycleDatumBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func kpackLifecycleDatumAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *KpackLifecycleDatum) error {
	*o = KpackLifecycleDatum{}
	return nil
}

func testKpackLifecycleDataHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &KpackLifecycleDatum{}
	o := &KpackLifecycleDatum{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum object: %s", err)
	}

	AddKpackLifecycleDatumHook(boil.BeforeInsertHook, kpackLifecycleDatumBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumBeforeInsertHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.AfterInsertHook, kpackLifecycleDatumAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumAfterInsertHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.AfterSelectHook, kpackLifecycleDatumAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumAfterSelectHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.BeforeUpdateHook, kpackLifecycleDatumBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumBeforeUpdateHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.AfterUpdateHook, kpackLifecycleDatumAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumAfterUpdateHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.BeforeDeleteHook, kpackLifecycleDatumBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumBeforeDeleteHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.AfterDeleteHook, kpackLifecycleDatumAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumAfterDeleteHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.BeforeUpsertHook, kpackLifecycleDatumBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumBeforeUpsertHooks = []KpackLifecycleDatumHook{}

	AddKpackLifecycleDatumHook(boil.AfterUpsertHook, kpackLifecycleDatumAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	kpackLifecycleDatumAfterUpsertHooks = []KpackLifecycleDatumHook{}
}

func testKpackLifecycleDataInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKpackLifecycleDataInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(kpackLifecycleDatumColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testKpackLifecycleDatumToOneAppUsingApp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local KpackLifecycleDatum
	var foreign App

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.AppGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.App().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := KpackLifecycleDatumSlice{&local}
	if err = local.L.LoadApp(ctx, tx, false, (*[]*KpackLifecycleDatum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.App = nil
	if err = local.L.LoadApp(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testKpackLifecycleDatumToOneBuildUsingBuild(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local KpackLifecycleDatum
	var foreign Build

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildDBTypes, false, buildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Build struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.BuildGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Build().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := KpackLifecycleDatumSlice{&local}
	if err = local.L.LoadBuild(ctx, tx, false, (*[]*KpackLifecycleDatum)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Build == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Build = nil
	if err = local.L.LoadBuild(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Build == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testKpackLifecycleDatumToOneSetOpAppUsingApp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a KpackLifecycleDatum
	var b, c App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, kpackLifecycleDatumDBTypes, false, strmangle.SetComplement(kpackLifecycleDatumPrimaryKeyColumns, kpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*App{&b, &c} {
		err = a.SetApp(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.App != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.KpackLifecycleData[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.AppGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.AppGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AppGUID))
		reflect.Indirect(reflect.ValueOf(&a.AppGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.AppGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.AppGUID, x.GUID)
		}
	}
}

func testKpackLifecycleDatumToOneRemoveOpAppUsingApp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a KpackLifecycleDatum
	var b App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, kpackLifecycleDatumDBTypes, false, strmangle.SetComplement(kpackLifecycleDatumPrimaryKeyColumns, kpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetApp(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveApp(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.App().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.App != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.AppGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.KpackLifecycleData) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testKpackLifecycleDatumToOneSetOpBuildUsingBuild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a KpackLifecycleDatum
	var b, c Build

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, kpackLifecycleDatumDBTypes, false, strmangle.SetComplement(kpackLifecycleDatumPrimaryKeyColumns, kpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Build{&b, &c} {
		err = a.SetBuild(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Build != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.KpackLifecycleData[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.BuildGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.BuildGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BuildGUID))
		reflect.Indirect(reflect.ValueOf(&a.BuildGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.BuildGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.BuildGUID, x.GUID)
		}
	}
}

func testKpackLifecycleDatumToOneRemoveOpBuildUsingBuild(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a KpackLifecycleDatum
	var b Build

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, kpackLifecycleDatumDBTypes, false, strmangle.SetComplement(kpackLifecycleDatumPrimaryKeyColumns, kpackLifecycleDatumColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBuild(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBuild(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Build().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Build != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.BuildGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.KpackLifecycleData) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testKpackLifecycleDataReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
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

func testKpackLifecycleDataReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := KpackLifecycleDatumSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testKpackLifecycleDataSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := KpackLifecycleData().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	kpackLifecycleDatumDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `BuildGUID`: `varchar`, `DropletGUID`: `varchar`, `AppGUID`: `varchar`, `Buildpacks`: `varchar`}
	_                          = bytes.MinRead
)

func testKpackLifecycleDataUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(kpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(kpackLifecycleDatumAllColumns) == len(kpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testKpackLifecycleDataSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(kpackLifecycleDatumAllColumns) == len(kpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &KpackLifecycleDatum{}
	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, kpackLifecycleDatumDBTypes, true, kpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(kpackLifecycleDatumAllColumns, kpackLifecycleDatumPrimaryKeyColumns) {
		fields = kpackLifecycleDatumAllColumns
	} else {
		fields = strmangle.SetComplement(
			kpackLifecycleDatumAllColumns,
			kpackLifecycleDatumPrimaryKeyColumns,
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

	slice := KpackLifecycleDatumSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testKpackLifecycleDataUpsert(t *testing.T) {
	t.Parallel()

	if len(kpackLifecycleDatumAllColumns) == len(kpackLifecycleDatumPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLKpackLifecycleDatumUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := KpackLifecycleDatum{}
	if err = randomize.Struct(seed, &o, kpackLifecycleDatumDBTypes, false); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KpackLifecycleDatum: %s", err)
	}

	count, err := KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, kpackLifecycleDatumDBTypes, false, kpackLifecycleDatumPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize KpackLifecycleDatum struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert KpackLifecycleDatum: %s", err)
	}

	count, err = KpackLifecycleData().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
