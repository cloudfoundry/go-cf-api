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

func testDropletLabels(t *testing.T) {
	t.Parallel()

	query := DropletLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDropletLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
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

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDropletLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DropletLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDropletLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DropletLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDropletLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DropletLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DropletLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DropletLabelExists to return true, but got false.")
	}
}

func testDropletLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	dropletLabelFound, err := FindDropletLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if dropletLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDropletLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DropletLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDropletLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DropletLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDropletLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	dropletLabelOne := &DropletLabel{}
	dropletLabelTwo := &DropletLabel{}
	if err = randomize.Struct(seed, dropletLabelOne, dropletLabelDBTypes, false, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, dropletLabelTwo, dropletLabelDBTypes, false, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dropletLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dropletLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DropletLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDropletLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	dropletLabelOne := &DropletLabel{}
	dropletLabelTwo := &DropletLabel{}
	if err = randomize.Struct(seed, dropletLabelOne, dropletLabelDBTypes, false, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, dropletLabelTwo, dropletLabelDBTypes, false, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = dropletLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = dropletLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func dropletLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func dropletLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DropletLabel) error {
	*o = DropletLabel{}
	return nil
}

func testDropletLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DropletLabel{}
	o := &DropletLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DropletLabel object: %s", err)
	}

	AddDropletLabelHook(boil.BeforeInsertHook, dropletLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	dropletLabelBeforeInsertHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.AfterInsertHook, dropletLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	dropletLabelAfterInsertHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.AfterSelectHook, dropletLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	dropletLabelAfterSelectHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.BeforeUpdateHook, dropletLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	dropletLabelBeforeUpdateHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.AfterUpdateHook, dropletLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	dropletLabelAfterUpdateHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.BeforeDeleteHook, dropletLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	dropletLabelBeforeDeleteHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.AfterDeleteHook, dropletLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	dropletLabelAfterDeleteHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.BeforeUpsertHook, dropletLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	dropletLabelBeforeUpsertHooks = []DropletLabelHook{}

	AddDropletLabelHook(boil.AfterUpsertHook, dropletLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	dropletLabelAfterUpsertHooks = []DropletLabelHook{}
}

func testDropletLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDropletLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(dropletLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDropletLabelToOneDropletUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DropletLabel
	var foreign Droplet

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, dropletDBTypes, false, dropletColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Droplet struct: %s", err)
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

	slice := DropletLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*DropletLabel)(&slice), nil); err != nil {
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

func testDropletLabelToOneSetOpDropletUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DropletLabel
	var b, c Droplet

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dropletLabelDBTypes, false, strmangle.SetComplement(dropletLabelPrimaryKeyColumns, dropletLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dropletDBTypes, false, strmangle.SetComplement(dropletPrimaryKeyColumns, dropletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, dropletDBTypes, false, strmangle.SetComplement(dropletPrimaryKeyColumns, dropletColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Droplet{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceDropletLabels[0] != &a {
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

func testDropletLabelToOneRemoveOpDropletUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DropletLabel
	var b Droplet

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, dropletLabelDBTypes, false, strmangle.SetComplement(dropletLabelPrimaryKeyColumns, dropletLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, dropletDBTypes, false, strmangle.SetComplement(dropletPrimaryKeyColumns, dropletColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceDropletLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDropletLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
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

func testDropletLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DropletLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDropletLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DropletLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	dropletLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                   = bytes.MinRead
)

func testDropletLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(dropletLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(dropletLabelAllColumns) == len(dropletLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDropletLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(dropletLabelAllColumns) == len(dropletLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DropletLabel{}
	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, dropletLabelDBTypes, true, dropletLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(dropletLabelAllColumns, dropletLabelPrimaryKeyColumns) {
		fields = dropletLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			dropletLabelAllColumns,
			dropletLabelPrimaryKeyColumns,
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

	slice := DropletLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDropletLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(dropletLabelAllColumns) == len(dropletLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DropletLabel{}
	if err = randomize.Struct(seed, &o, dropletLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DropletLabel: %s", err)
	}

	count, err := DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, dropletLabelDBTypes, false, dropletLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DropletLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DropletLabel: %s", err)
	}

	count, err = DropletLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
