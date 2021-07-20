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

func testSpacesManagers(t *testing.T) {
	t.Parallel()

	query := SpacesManagers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpacesManagersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
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

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesManagersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpacesManagers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesManagersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesManagerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesManagersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpacesManagerExists(ctx, tx, o.SpacesManagersPK)
	if err != nil {
		t.Errorf("Unable to check if SpacesManager exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpacesManagerExists to return true, but got false.")
	}
}

func testSpacesManagersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spacesManagerFound, err := FindSpacesManager(ctx, tx, o.SpacesManagersPK)
	if err != nil {
		t.Error(err)
	}

	if spacesManagerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpacesManagersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpacesManagers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpacesManagersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpacesManagers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpacesManagersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spacesManagerOne := &SpacesManager{}
	spacesManagerTwo := &SpacesManager{}
	if err = randomize.Struct(seed, spacesManagerOne, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesManagerTwo, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpacesManagersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spacesManagerOne := &SpacesManager{}
	spacesManagerTwo := &SpacesManager{}
	if err = randomize.Struct(seed, spacesManagerOne, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesManagerTwo, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spacesManagerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func spacesManagerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesManager) error {
	*o = SpacesManager{}
	return nil
}

func testSpacesManagersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpacesManager{}
	o := &SpacesManager{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpacesManager object: %s", err)
	}

	AddSpacesManagerHook(boil.BeforeInsertHook, spacesManagerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spacesManagerBeforeInsertHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.AfterInsertHook, spacesManagerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spacesManagerAfterInsertHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.AfterSelectHook, spacesManagerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spacesManagerAfterSelectHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.BeforeUpdateHook, spacesManagerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spacesManagerBeforeUpdateHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.AfterUpdateHook, spacesManagerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spacesManagerAfterUpdateHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.BeforeDeleteHook, spacesManagerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spacesManagerBeforeDeleteHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.AfterDeleteHook, spacesManagerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spacesManagerAfterDeleteHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.BeforeUpsertHook, spacesManagerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spacesManagerBeforeUpsertHooks = []SpacesManagerHook{}

	AddSpacesManagerHook(boil.AfterUpsertHook, spacesManagerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spacesManagerAfterUpsertHooks = []SpacesManagerHook{}
}

func testSpacesManagersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesManagersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spacesManagerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesManagerToOneSpaceUsingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesManager
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Space struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SpaceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Space().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SpacesManagerSlice{&local}
	if err = local.L.LoadSpace(ctx, tx, false, (*[]*SpacesManager)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Space == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Space = nil
	if err = local.L.LoadSpace(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Space == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSpacesManagerToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesManager
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesManagerDBTypes, false, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SpacesManagerSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*SpacesManager)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSpacesManagerToOneSetOpSpaceUsingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesManager
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesManagerDBTypes, false, strmangle.SetComplement(spacesManagerPrimaryKeyColumns, spacesManagerColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSpace(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Space != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SpacesManagers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SpaceID != x.ID {
			t.Error("foreign key was wrong value", a.SpaceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SpaceID))
		reflect.Indirect(reflect.ValueOf(&a.SpaceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SpaceID != x.ID {
			t.Error("foreign key was wrong value", a.SpaceID, x.ID)
		}
	}
}
func testSpacesManagerToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesManager
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesManagerDBTypes, false, strmangle.SetComplement(spacesManagerPrimaryKeyColumns, spacesManagerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SpacesManagers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testSpacesManagersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
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

func testSpacesManagersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesManagerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpacesManagersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spacesManagerDBTypes = map[string]string{`SpaceID`: `integer`, `UserID`: `integer`, `SpacesManagersPK`: `integer`, `RoleGUID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testSpacesManagersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spacesManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spacesManagerAllColumns) == len(spacesManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpacesManagersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spacesManagerAllColumns) == len(spacesManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesManager{}
	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesManagerDBTypes, true, spacesManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spacesManagerAllColumns, spacesManagerPrimaryKeyColumns) {
		fields = spacesManagerAllColumns
	} else {
		fields = strmangle.SetComplement(
			spacesManagerAllColumns,
			spacesManagerPrimaryKeyColumns,
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

	slice := SpacesManagerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpacesManagersUpsert(t *testing.T) {
	t.Parallel()

	if len(spacesManagerAllColumns) == len(spacesManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpacesManager{}
	if err = randomize.Struct(seed, &o, spacesManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesManager: %s", err)
	}

	count, err := SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spacesManagerDBTypes, false, spacesManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesManager struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesManager: %s", err)
	}

	count, err = SpacesManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
