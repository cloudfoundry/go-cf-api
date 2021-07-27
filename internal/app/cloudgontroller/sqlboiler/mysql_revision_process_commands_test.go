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

func testRevisionProcessCommands(t *testing.T) {
	t.Parallel()

	query := RevisionProcessCommands()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevisionProcessCommandsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
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

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionProcessCommandsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevisionProcessCommands().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionProcessCommandsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionProcessCommandSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionProcessCommandsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevisionProcessCommandExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevisionProcessCommand exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevisionProcessCommandExists to return true, but got false.")
	}
}

func testRevisionProcessCommandsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revisionProcessCommandFound, err := FindRevisionProcessCommand(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revisionProcessCommandFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevisionProcessCommandsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevisionProcessCommands().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevisionProcessCommandsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevisionProcessCommands().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevisionProcessCommandsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revisionProcessCommandOne := &RevisionProcessCommand{}
	revisionProcessCommandTwo := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, revisionProcessCommandOne, revisionProcessCommandDBTypes, false, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionProcessCommandTwo, revisionProcessCommandDBTypes, false, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionProcessCommandOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionProcessCommandTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionProcessCommands().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevisionProcessCommandsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revisionProcessCommandOne := &RevisionProcessCommand{}
	revisionProcessCommandTwo := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, revisionProcessCommandOne, revisionProcessCommandDBTypes, false, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionProcessCommandTwo, revisionProcessCommandDBTypes, false, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionProcessCommandOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionProcessCommandTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revisionProcessCommandBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func revisionProcessCommandAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionProcessCommand) error {
	*o = RevisionProcessCommand{}
	return nil
}

func testRevisionProcessCommandsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevisionProcessCommand{}
	o := &RevisionProcessCommand{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand object: %s", err)
	}

	AddRevisionProcessCommandHook(boil.BeforeInsertHook, revisionProcessCommandBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandBeforeInsertHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.AfterInsertHook, revisionProcessCommandAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandAfterInsertHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.AfterSelectHook, revisionProcessCommandAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandAfterSelectHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.BeforeUpdateHook, revisionProcessCommandBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandBeforeUpdateHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.AfterUpdateHook, revisionProcessCommandAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandAfterUpdateHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.BeforeDeleteHook, revisionProcessCommandBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandBeforeDeleteHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.AfterDeleteHook, revisionProcessCommandAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandAfterDeleteHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.BeforeUpsertHook, revisionProcessCommandBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandBeforeUpsertHooks = []RevisionProcessCommandHook{}

	AddRevisionProcessCommandHook(boil.AfterUpsertHook, revisionProcessCommandAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revisionProcessCommandAfterUpsertHooks = []RevisionProcessCommandHook{}
}

func testRevisionProcessCommandsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionProcessCommandsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revisionProcessCommandColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionProcessCommandToOneRevisionUsingRevision(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevisionProcessCommand
	var foreign Revision

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revisionProcessCommandDBTypes, false, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, revisionDBTypes, false, revisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Revision struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RevisionGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Revision().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := RevisionProcessCommandSlice{&local}
	if err = local.L.LoadRevision(ctx, tx, false, (*[]*RevisionProcessCommand)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Revision = nil
	if err = local.L.LoadRevision(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRevisionProcessCommandToOneSetOpRevisionUsingRevision(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionProcessCommand
	var b, c Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionProcessCommandDBTypes, false, strmangle.SetComplement(revisionProcessCommandPrimaryKeyColumns, revisionProcessCommandColumnsWithoutDefault)...); err != nil {
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
		err = a.SetRevision(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Revision != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RevisionProcessCommands[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RevisionGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RevisionGUID))
		reflect.Indirect(reflect.ValueOf(&a.RevisionGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RevisionGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionGUID, x.GUID)
		}
	}
}

func testRevisionProcessCommandsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
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

func testRevisionProcessCommandsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionProcessCommandSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevisionProcessCommandsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionProcessCommands().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revisionProcessCommandDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `RevisionGUID`: `varchar`, `ProcessType`: `varchar`, `ProcessCommand`: `varchar`}
	_                             = bytes.MinRead
)

func testRevisionProcessCommandsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revisionProcessCommandPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revisionProcessCommandAllColumns) == len(revisionProcessCommandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevisionProcessCommandsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revisionProcessCommandAllColumns) == len(revisionProcessCommandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionProcessCommand{}
	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionProcessCommandDBTypes, true, revisionProcessCommandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revisionProcessCommandAllColumns, revisionProcessCommandPrimaryKeyColumns) {
		fields = revisionProcessCommandAllColumns
	} else {
		fields = strmangle.SetComplement(
			revisionProcessCommandAllColumns,
			revisionProcessCommandPrimaryKeyColumns,
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

	slice := RevisionProcessCommandSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevisionProcessCommandsUpsert(t *testing.T) {
	t.Parallel()

	if len(revisionProcessCommandAllColumns) == len(revisionProcessCommandPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRevisionProcessCommandUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevisionProcessCommand{}
	if err = randomize.Struct(seed, &o, revisionProcessCommandDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionProcessCommand: %s", err)
	}

	count, err := RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revisionProcessCommandDBTypes, false, revisionProcessCommandPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionProcessCommand struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionProcessCommand: %s", err)
	}

	count, err = RevisionProcessCommands().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
