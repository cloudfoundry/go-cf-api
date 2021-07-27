// +build integration,mysql
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

func testRevisionSidecarProcessTypes(t *testing.T) {
	t.Parallel()

	query := RevisionSidecarProcessTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevisionSidecarProcessTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
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

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarProcessTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevisionSidecarProcessTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarProcessTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionSidecarProcessTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarProcessTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevisionSidecarProcessTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevisionSidecarProcessType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevisionSidecarProcessTypeExists to return true, but got false.")
	}
}

func testRevisionSidecarProcessTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revisionSidecarProcessTypeFound, err := FindRevisionSidecarProcessType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revisionSidecarProcessTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevisionSidecarProcessTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevisionSidecarProcessTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevisionSidecarProcessTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevisionSidecarProcessTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevisionSidecarProcessTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revisionSidecarProcessTypeOne := &RevisionSidecarProcessType{}
	revisionSidecarProcessTypeTwo := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, revisionSidecarProcessTypeOne, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionSidecarProcessTypeTwo, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionSidecarProcessTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionSidecarProcessTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionSidecarProcessTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevisionSidecarProcessTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revisionSidecarProcessTypeOne := &RevisionSidecarProcessType{}
	revisionSidecarProcessTypeTwo := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, revisionSidecarProcessTypeOne, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionSidecarProcessTypeTwo, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionSidecarProcessTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionSidecarProcessTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revisionSidecarProcessTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func revisionSidecarProcessTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecarProcessType) error {
	*o = RevisionSidecarProcessType{}
	return nil
}

func testRevisionSidecarProcessTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevisionSidecarProcessType{}
	o := &RevisionSidecarProcessType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType object: %s", err)
	}

	AddRevisionSidecarProcessTypeHook(boil.BeforeInsertHook, revisionSidecarProcessTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeBeforeInsertHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.AfterInsertHook, revisionSidecarProcessTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeAfterInsertHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.AfterSelectHook, revisionSidecarProcessTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeAfterSelectHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.BeforeUpdateHook, revisionSidecarProcessTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeBeforeUpdateHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.AfterUpdateHook, revisionSidecarProcessTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeAfterUpdateHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.BeforeDeleteHook, revisionSidecarProcessTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeBeforeDeleteHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.AfterDeleteHook, revisionSidecarProcessTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeAfterDeleteHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.BeforeUpsertHook, revisionSidecarProcessTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeBeforeUpsertHooks = []RevisionSidecarProcessTypeHook{}

	AddRevisionSidecarProcessTypeHook(boil.AfterUpsertHook, revisionSidecarProcessTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarProcessTypeAfterUpsertHooks = []RevisionSidecarProcessTypeHook{}
}

func testRevisionSidecarProcessTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionSidecarProcessTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revisionSidecarProcessTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionSidecarProcessTypeToOneRevisionSidecarUsingRevisionSidecar(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevisionSidecarProcessType
	var foreign RevisionSidecar

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RevisionSidecarGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RevisionSidecar().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := RevisionSidecarProcessTypeSlice{&local}
	if err = local.L.LoadRevisionSidecar(ctx, tx, false, (*[]*RevisionSidecarProcessType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RevisionSidecar == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RevisionSidecar = nil
	if err = local.L.LoadRevisionSidecar(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RevisionSidecar == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRevisionSidecarProcessTypeToOneSetOpRevisionSidecarUsingRevisionSidecar(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionSidecarProcessType
	var b, c RevisionSidecar

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionSidecarProcessTypeDBTypes, false, strmangle.SetComplement(revisionSidecarProcessTypePrimaryKeyColumns, revisionSidecarProcessTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, revisionSidecarDBTypes, false, strmangle.SetComplement(revisionSidecarPrimaryKeyColumns, revisionSidecarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, revisionSidecarDBTypes, false, strmangle.SetComplement(revisionSidecarPrimaryKeyColumns, revisionSidecarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RevisionSidecar{&b, &c} {
		err = a.SetRevisionSidecar(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RevisionSidecar != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RevisionSidecarProcessTypes[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RevisionSidecarGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionSidecarGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RevisionSidecarGUID))
		reflect.Indirect(reflect.ValueOf(&a.RevisionSidecarGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RevisionSidecarGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionSidecarGUID, x.GUID)
		}
	}
}

func testRevisionSidecarProcessTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
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

func testRevisionSidecarProcessTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionSidecarProcessTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevisionSidecarProcessTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionSidecarProcessTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revisionSidecarProcessTypeDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `Type`: `varchar`, `RevisionSidecarGUID`: `varchar`}
	_                                 = bytes.MinRead
)

func testRevisionSidecarProcessTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revisionSidecarProcessTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revisionSidecarProcessTypeAllColumns) == len(revisionSidecarProcessTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevisionSidecarProcessTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revisionSidecarProcessTypeAllColumns) == len(revisionSidecarProcessTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionSidecarProcessTypeDBTypes, true, revisionSidecarProcessTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revisionSidecarProcessTypeAllColumns, revisionSidecarProcessTypePrimaryKeyColumns) {
		fields = revisionSidecarProcessTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			revisionSidecarProcessTypeAllColumns,
			revisionSidecarProcessTypePrimaryKeyColumns,
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

	slice := RevisionSidecarProcessTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevisionSidecarProcessTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(revisionSidecarProcessTypeAllColumns) == len(revisionSidecarProcessTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRevisionSidecarProcessTypeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevisionSidecarProcessType{}
	if err = randomize.Struct(seed, &o, revisionSidecarProcessTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionSidecarProcessType: %s", err)
	}

	count, err := RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecarProcessType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionSidecarProcessType: %s", err)
	}

	count, err = RevisionSidecarProcessTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
