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

func testStackLabels(t *testing.T) {
	t.Parallel()

	query := StackLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStackLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
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

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStackLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StackLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStackLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StackLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStackLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StackLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if StackLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StackLabelExists to return true, but got false.")
	}
}

func testStackLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stackLabelFound, err := FindStackLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if stackLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStackLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StackLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStackLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StackLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStackLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stackLabelOne := &StackLabel{}
	stackLabelTwo := &StackLabel{}
	if err = randomize.Struct(seed, stackLabelOne, stackLabelDBTypes, false, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, stackLabelTwo, stackLabelDBTypes, false, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stackLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stackLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StackLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStackLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stackLabelOne := &StackLabel{}
	stackLabelTwo := &StackLabel{}
	if err = randomize.Struct(seed, stackLabelOne, stackLabelDBTypes, false, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, stackLabelTwo, stackLabelDBTypes, false, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stackLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stackLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func stackLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func stackLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StackLabel) error {
	*o = StackLabel{}
	return nil
}

func testStackLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StackLabel{}
	o := &StackLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stackLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StackLabel object: %s", err)
	}

	AddStackLabelHook(boil.BeforeInsertHook, stackLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stackLabelBeforeInsertHooks = []StackLabelHook{}

	AddStackLabelHook(boil.AfterInsertHook, stackLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stackLabelAfterInsertHooks = []StackLabelHook{}

	AddStackLabelHook(boil.AfterSelectHook, stackLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stackLabelAfterSelectHooks = []StackLabelHook{}

	AddStackLabelHook(boil.BeforeUpdateHook, stackLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stackLabelBeforeUpdateHooks = []StackLabelHook{}

	AddStackLabelHook(boil.AfterUpdateHook, stackLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stackLabelAfterUpdateHooks = []StackLabelHook{}

	AddStackLabelHook(boil.BeforeDeleteHook, stackLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stackLabelBeforeDeleteHooks = []StackLabelHook{}

	AddStackLabelHook(boil.AfterDeleteHook, stackLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stackLabelAfterDeleteHooks = []StackLabelHook{}

	AddStackLabelHook(boil.BeforeUpsertHook, stackLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stackLabelBeforeUpsertHooks = []StackLabelHook{}

	AddStackLabelHook(boil.AfterUpsertHook, stackLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stackLabelAfterUpsertHooks = []StackLabelHook{}
}

func testStackLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStackLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stackLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStackLabelToOneStackUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StackLabel
	var foreign Stack

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stackDBTypes, false, stackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stack struct: %s", err)
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

	slice := StackLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*StackLabel)(&slice), nil); err != nil {
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

func testStackLabelToOneSetOpStackUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StackLabel
	var b, c Stack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stackLabelDBTypes, false, strmangle.SetComplement(stackLabelPrimaryKeyColumns, stackLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stackDBTypes, false, strmangle.SetComplement(stackPrimaryKeyColumns, stackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stackDBTypes, false, strmangle.SetComplement(stackPrimaryKeyColumns, stackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stack{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceStackLabels[0] != &a {
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

func testStackLabelToOneRemoveOpStackUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StackLabel
	var b Stack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stackLabelDBTypes, false, strmangle.SetComplement(stackLabelPrimaryKeyColumns, stackLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stackDBTypes, false, strmangle.SetComplement(stackPrimaryKeyColumns, stackColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceStackLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testStackLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
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

func testStackLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StackLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStackLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StackLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stackLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                 = bytes.MinRead
)

func testStackLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stackLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stackLabelAllColumns) == len(stackLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStackLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stackLabelAllColumns) == len(stackLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StackLabel{}
	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stackLabelDBTypes, true, stackLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stackLabelAllColumns, stackLabelPrimaryKeyColumns) {
		fields = stackLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			stackLabelAllColumns,
			stackLabelPrimaryKeyColumns,
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

	slice := StackLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStackLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(stackLabelAllColumns) == len(stackLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StackLabel{}
	if err = randomize.Struct(seed, &o, stackLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StackLabel: %s", err)
	}

	count, err := StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stackLabelDBTypes, false, stackLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StackLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StackLabel: %s", err)
	}

	count, err = StackLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
