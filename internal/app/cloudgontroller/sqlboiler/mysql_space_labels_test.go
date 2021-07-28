// +build mysql_integration
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

func testSpaceLabels(t *testing.T) {
	t.Parallel()

	query := SpaceLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpaceLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
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

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpaceLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpaceLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpaceLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SpaceLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpaceLabelExists to return true, but got false.")
	}
}

func testSpaceLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spaceLabelFound, err := FindSpaceLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if spaceLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpaceLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpaceLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpaceLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpaceLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpaceLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spaceLabelOne := &SpaceLabel{}
	spaceLabelTwo := &SpaceLabel{}
	if err = randomize.Struct(seed, spaceLabelOne, spaceLabelDBTypes, false, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, spaceLabelTwo, spaceLabelDBTypes, false, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spaceLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spaceLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpaceLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpaceLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spaceLabelOne := &SpaceLabel{}
	spaceLabelTwo := &SpaceLabel{}
	if err = randomize.Struct(seed, spaceLabelOne, spaceLabelDBTypes, false, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, spaceLabelTwo, spaceLabelDBTypes, false, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spaceLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spaceLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spaceLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func spaceLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceLabel) error {
	*o = SpaceLabel{}
	return nil
}

func testSpaceLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpaceLabel{}
	o := &SpaceLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpaceLabel object: %s", err)
	}

	AddSpaceLabelHook(boil.BeforeInsertHook, spaceLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spaceLabelBeforeInsertHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.AfterInsertHook, spaceLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spaceLabelAfterInsertHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.AfterSelectHook, spaceLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spaceLabelAfterSelectHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.BeforeUpdateHook, spaceLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spaceLabelBeforeUpdateHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.AfterUpdateHook, spaceLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spaceLabelAfterUpdateHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.BeforeDeleteHook, spaceLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spaceLabelBeforeDeleteHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.AfterDeleteHook, spaceLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spaceLabelAfterDeleteHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.BeforeUpsertHook, spaceLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spaceLabelBeforeUpsertHooks = []SpaceLabelHook{}

	AddSpaceLabelHook(boil.AfterUpsertHook, spaceLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spaceLabelAfterUpsertHooks = []SpaceLabelHook{}
}

func testSpaceLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpaceLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spaceLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpaceLabelToOneSpaceUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpaceLabel
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Space struct: %s", err)
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

	slice := SpaceLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*SpaceLabel)(&slice), nil); err != nil {
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

func testSpaceLabelToOneSetOpSpaceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceLabel
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceLabelDBTypes, false, strmangle.SetComplement(spaceLabelPrimaryKeyColumns, spaceLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Space{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceSpaceLabels[0] != &a {
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

func testSpaceLabelToOneRemoveOpSpaceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceLabel
	var b Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceLabelDBTypes, false, strmangle.SetComplement(spaceLabelPrimaryKeyColumns, spaceLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceSpaceLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testSpaceLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
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

func testSpaceLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpaceLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpaceLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpaceLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spaceLabelDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `KeyName`: `varchar`, `Value`: `varchar`}
	_                 = bytes.MinRead
)

func testSpaceLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spaceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spaceLabelAllColumns) == len(spaceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpaceLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spaceLabelAllColumns) == len(spaceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpaceLabel{}
	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spaceLabelDBTypes, true, spaceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spaceLabelAllColumns, spaceLabelPrimaryKeyColumns) {
		fields = spaceLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			spaceLabelAllColumns,
			spaceLabelPrimaryKeyColumns,
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

	slice := SpaceLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpaceLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(spaceLabelAllColumns) == len(spaceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSpaceLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpaceLabel{}
	if err = randomize.Struct(seed, &o, spaceLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpaceLabel: %s", err)
	}

	count, err := SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spaceLabelDBTypes, false, spaceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpaceLabel: %s", err)
	}

	count, err = SpaceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
