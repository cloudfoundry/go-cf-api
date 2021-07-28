// +build mysql
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

func testRevisionLabels(t *testing.T) {
	t.Parallel()

	query := RevisionLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevisionLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
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

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevisionLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevisionLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevisionLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevisionLabelExists to return true, but got false.")
	}
}

func testRevisionLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revisionLabelFound, err := FindRevisionLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revisionLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevisionLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevisionLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevisionLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevisionLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevisionLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revisionLabelOne := &RevisionLabel{}
	revisionLabelTwo := &RevisionLabel{}
	if err = randomize.Struct(seed, revisionLabelOne, revisionLabelDBTypes, false, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionLabelTwo, revisionLabelDBTypes, false, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevisionLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revisionLabelOne := &RevisionLabel{}
	revisionLabelTwo := &RevisionLabel{}
	if err = randomize.Struct(seed, revisionLabelOne, revisionLabelDBTypes, false, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionLabelTwo, revisionLabelDBTypes, false, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revisionLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func revisionLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionLabel) error {
	*o = RevisionLabel{}
	return nil
}

func testRevisionLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevisionLabel{}
	o := &RevisionLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionLabel object: %s", err)
	}

	AddRevisionLabelHook(boil.BeforeInsertHook, revisionLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revisionLabelBeforeInsertHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.AfterInsertHook, revisionLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revisionLabelAfterInsertHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.AfterSelectHook, revisionLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revisionLabelAfterSelectHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.BeforeUpdateHook, revisionLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revisionLabelBeforeUpdateHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.AfterUpdateHook, revisionLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revisionLabelAfterUpdateHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.BeforeDeleteHook, revisionLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revisionLabelBeforeDeleteHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.AfterDeleteHook, revisionLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revisionLabelAfterDeleteHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.BeforeUpsertHook, revisionLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revisionLabelBeforeUpsertHooks = []RevisionLabelHook{}

	AddRevisionLabelHook(boil.AfterUpsertHook, revisionLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revisionLabelAfterUpsertHooks = []RevisionLabelHook{}
}

func testRevisionLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revisionLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionLabelToOneRevisionUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevisionLabel
	var foreign Revision

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, revisionDBTypes, false, revisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Revision struct: %s", err)
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

	slice := RevisionLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*RevisionLabel)(&slice), nil); err != nil {
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

func testRevisionLabelToOneSetOpRevisionUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionLabel
	var b, c Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionLabelDBTypes, false, strmangle.SetComplement(revisionLabelPrimaryKeyColumns, revisionLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Revision{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceRevisionLabels[0] != &a {
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

func testRevisionLabelToOneRemoveOpRevisionUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionLabel
	var b Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionLabelDBTypes, false, strmangle.SetComplement(revisionLabelPrimaryKeyColumns, revisionLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceRevisionLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRevisionLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
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

func testRevisionLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevisionLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revisionLabelDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `KeyName`: `varchar`, `Value`: `varchar`}
	_                    = bytes.MinRead
)

func testRevisionLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revisionLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revisionLabelAllColumns) == len(revisionLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevisionLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revisionLabelAllColumns) == len(revisionLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionLabel{}
	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionLabelDBTypes, true, revisionLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revisionLabelAllColumns, revisionLabelPrimaryKeyColumns) {
		fields = revisionLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			revisionLabelAllColumns,
			revisionLabelPrimaryKeyColumns,
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

	slice := RevisionLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevisionLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(revisionLabelAllColumns) == len(revisionLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRevisionLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevisionLabel{}
	if err = randomize.Struct(seed, &o, revisionLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionLabel: %s", err)
	}

	count, err := RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revisionLabelDBTypes, false, revisionLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionLabel: %s", err)
	}

	count, err = RevisionLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
