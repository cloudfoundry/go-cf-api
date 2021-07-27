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

func testIsolationSegmentLabels(t *testing.T) {
	t.Parallel()

	query := IsolationSegmentLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIsolationSegmentLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
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

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := IsolationSegmentLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IsolationSegmentLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IsolationSegmentLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if IsolationSegmentLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IsolationSegmentLabelExists to return true, but got false.")
	}
}

func testIsolationSegmentLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	isolationSegmentLabelFound, err := FindIsolationSegmentLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if isolationSegmentLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIsolationSegmentLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = IsolationSegmentLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIsolationSegmentLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := IsolationSegmentLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIsolationSegmentLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	isolationSegmentLabelOne := &IsolationSegmentLabel{}
	isolationSegmentLabelTwo := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, isolationSegmentLabelOne, isolationSegmentLabelDBTypes, false, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, isolationSegmentLabelTwo, isolationSegmentLabelDBTypes, false, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = isolationSegmentLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = isolationSegmentLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IsolationSegmentLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIsolationSegmentLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	isolationSegmentLabelOne := &IsolationSegmentLabel{}
	isolationSegmentLabelTwo := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, isolationSegmentLabelOne, isolationSegmentLabelDBTypes, false, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, isolationSegmentLabelTwo, isolationSegmentLabelDBTypes, false, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = isolationSegmentLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = isolationSegmentLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func isolationSegmentLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func isolationSegmentLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentLabel) error {
	*o = IsolationSegmentLabel{}
	return nil
}

func testIsolationSegmentLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &IsolationSegmentLabel{}
	o := &IsolationSegmentLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel object: %s", err)
	}

	AddIsolationSegmentLabelHook(boil.BeforeInsertHook, isolationSegmentLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelBeforeInsertHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.AfterInsertHook, isolationSegmentLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelAfterInsertHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.AfterSelectHook, isolationSegmentLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelAfterSelectHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.BeforeUpdateHook, isolationSegmentLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelBeforeUpdateHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.AfterUpdateHook, isolationSegmentLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelAfterUpdateHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.BeforeDeleteHook, isolationSegmentLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelBeforeDeleteHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.AfterDeleteHook, isolationSegmentLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelAfterDeleteHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.BeforeUpsertHook, isolationSegmentLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelBeforeUpsertHooks = []IsolationSegmentLabelHook{}

	AddIsolationSegmentLabelHook(boil.AfterUpsertHook, isolationSegmentLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentLabelAfterUpsertHooks = []IsolationSegmentLabelHook{}
}

func testIsolationSegmentLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIsolationSegmentLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(isolationSegmentLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIsolationSegmentLabelToOneIsolationSegmentUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local IsolationSegmentLabel
	var foreign IsolationSegment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, isolationSegmentDBTypes, false, isolationSegmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegment struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ResourceGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Resource().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := IsolationSegmentLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*IsolationSegmentLabel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Resource == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Resource = nil
	if err = local.L.LoadResource(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Resource == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIsolationSegmentLabelToOneSetOpIsolationSegmentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IsolationSegmentLabel
	var b, c IsolationSegment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, isolationSegmentLabelDBTypes, false, strmangle.SetComplement(isolationSegmentLabelPrimaryKeyColumns, isolationSegmentLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*IsolationSegment{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceIsolationSegmentLabels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ResourceGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.ResourceGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ResourceGUID))
		reflect.Indirect(reflect.ValueOf(&a.ResourceGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ResourceGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.ResourceGUID, x.GUID)
		}
	}
}

func testIsolationSegmentLabelToOneRemoveOpIsolationSegmentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IsolationSegmentLabel
	var b IsolationSegment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, isolationSegmentLabelDBTypes, false, strmangle.SetComplement(isolationSegmentLabelPrimaryKeyColumns, isolationSegmentLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetResource(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveResource(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Resource().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Resource != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ResourceGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ResourceIsolationSegmentLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIsolationSegmentLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
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

func testIsolationSegmentLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IsolationSegmentLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIsolationSegmentLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IsolationSegmentLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	isolationSegmentLabelDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `KeyName`: `varchar`, `Value`: `varchar`}
	_                            = bytes.MinRead
)

func testIsolationSegmentLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(isolationSegmentLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(isolationSegmentLabelAllColumns) == len(isolationSegmentLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIsolationSegmentLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(isolationSegmentLabelAllColumns) == len(isolationSegmentLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentLabel{}
	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, isolationSegmentLabelDBTypes, true, isolationSegmentLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(isolationSegmentLabelAllColumns, isolationSegmentLabelPrimaryKeyColumns) {
		fields = isolationSegmentLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			isolationSegmentLabelAllColumns,
			isolationSegmentLabelPrimaryKeyColumns,
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

	slice := IsolationSegmentLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIsolationSegmentLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(isolationSegmentLabelAllColumns) == len(isolationSegmentLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLIsolationSegmentLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := IsolationSegmentLabel{}
	if err = randomize.Struct(seed, &o, isolationSegmentLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IsolationSegmentLabel: %s", err)
	}

	count, err := IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, isolationSegmentLabelDBTypes, false, isolationSegmentLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IsolationSegmentLabel: %s", err)
	}

	count, err = IsolationSegmentLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
