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

func testProcessLabels(t *testing.T) {
	t.Parallel()

	query := ProcessLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProcessLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
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

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProcessLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProcessLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ProcessLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProcessLabelExists to return true, but got false.")
	}
}

func testProcessLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	processLabelFound, err := FindProcessLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if processLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProcessLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProcessLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProcessLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProcessLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProcessLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	processLabelOne := &ProcessLabel{}
	processLabelTwo := &ProcessLabel{}
	if err = randomize.Struct(seed, processLabelOne, processLabelDBTypes, false, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, processLabelTwo, processLabelDBTypes, false, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = processLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProcessLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	processLabelOne := &ProcessLabel{}
	processLabelTwo := &ProcessLabel{}
	if err = randomize.Struct(seed, processLabelOne, processLabelDBTypes, false, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, processLabelTwo, processLabelDBTypes, false, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = processLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func processLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func processLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessLabel) error {
	*o = ProcessLabel{}
	return nil
}

func testProcessLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ProcessLabel{}
	o := &ProcessLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, processLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ProcessLabel object: %s", err)
	}

	AddProcessLabelHook(boil.BeforeInsertHook, processLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	processLabelBeforeInsertHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.AfterInsertHook, processLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	processLabelAfterInsertHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.AfterSelectHook, processLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	processLabelAfterSelectHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.BeforeUpdateHook, processLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	processLabelBeforeUpdateHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.AfterUpdateHook, processLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	processLabelAfterUpdateHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.BeforeDeleteHook, processLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	processLabelBeforeDeleteHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.AfterDeleteHook, processLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	processLabelAfterDeleteHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.BeforeUpsertHook, processLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	processLabelBeforeUpsertHooks = []ProcessLabelHook{}

	AddProcessLabelHook(boil.AfterUpsertHook, processLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	processLabelAfterUpsertHooks = []ProcessLabelHook{}
}

func testProcessLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(processLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessLabelToOneProcessUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ProcessLabel
	var foreign Process

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, processDBTypes, false, processColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Process struct: %s", err)
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

	slice := ProcessLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ProcessLabel)(&slice), nil); err != nil {
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

func testProcessLabelToOneSetOpProcessUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProcessLabel
	var b, c Process

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, processLabelDBTypes, false, strmangle.SetComplement(processLabelPrimaryKeyColumns, processLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Process{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceProcessLabels[0] != &a {
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

func testProcessLabelToOneRemoveOpProcessUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProcessLabel
	var b Process

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, processLabelDBTypes, false, strmangle.SetComplement(processLabelPrimaryKeyColumns, processLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceProcessLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProcessLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
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

func testProcessLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProcessLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	processLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                   = bytes.MinRead
)

func testProcessLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(processLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(processLabelAllColumns) == len(processLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProcessLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(processLabelAllColumns) == len(processLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessLabel{}
	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processLabelDBTypes, true, processLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(processLabelAllColumns, processLabelPrimaryKeyColumns) {
		fields = processLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			processLabelAllColumns,
			processLabelPrimaryKeyColumns,
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

	slice := ProcessLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProcessLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(processLabelAllColumns) == len(processLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ProcessLabel{}
	if err = randomize.Struct(seed, &o, processLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProcessLabel: %s", err)
	}

	count, err := ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, processLabelDBTypes, false, processLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProcessLabel: %s", err)
	}

	count, err = ProcessLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
