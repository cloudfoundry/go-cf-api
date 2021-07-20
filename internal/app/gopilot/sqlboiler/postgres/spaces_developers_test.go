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

func testSpacesDevelopers(t *testing.T) {
	t.Parallel()

	query := SpacesDevelopers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpacesDevelopersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
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

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesDevelopersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpacesDevelopers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesDevelopersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesDeveloperSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesDevelopersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpacesDeveloperExists(ctx, tx, o.SpacesDevelopersPK)
	if err != nil {
		t.Errorf("Unable to check if SpacesDeveloper exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpacesDeveloperExists to return true, but got false.")
	}
}

func testSpacesDevelopersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spacesDeveloperFound, err := FindSpacesDeveloper(ctx, tx, o.SpacesDevelopersPK)
	if err != nil {
		t.Error(err)
	}

	if spacesDeveloperFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpacesDevelopersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpacesDevelopers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpacesDevelopersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpacesDevelopers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpacesDevelopersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spacesDeveloperOne := &SpacesDeveloper{}
	spacesDeveloperTwo := &SpacesDeveloper{}
	if err = randomize.Struct(seed, spacesDeveloperOne, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesDeveloperTwo, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesDeveloperOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesDeveloperTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesDevelopers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpacesDevelopersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spacesDeveloperOne := &SpacesDeveloper{}
	spacesDeveloperTwo := &SpacesDeveloper{}
	if err = randomize.Struct(seed, spacesDeveloperOne, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesDeveloperTwo, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesDeveloperOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesDeveloperTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spacesDeveloperBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func spacesDeveloperAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesDeveloper) error {
	*o = SpacesDeveloper{}
	return nil
}

func testSpacesDevelopersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpacesDeveloper{}
	o := &SpacesDeveloper{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper object: %s", err)
	}

	AddSpacesDeveloperHook(boil.BeforeInsertHook, spacesDeveloperBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperBeforeInsertHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.AfterInsertHook, spacesDeveloperAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperAfterInsertHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.AfterSelectHook, spacesDeveloperAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperAfterSelectHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.BeforeUpdateHook, spacesDeveloperBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperBeforeUpdateHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.AfterUpdateHook, spacesDeveloperAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperAfterUpdateHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.BeforeDeleteHook, spacesDeveloperBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperBeforeDeleteHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.AfterDeleteHook, spacesDeveloperAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperAfterDeleteHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.BeforeUpsertHook, spacesDeveloperBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperBeforeUpsertHooks = []SpacesDeveloperHook{}

	AddSpacesDeveloperHook(boil.AfterUpsertHook, spacesDeveloperAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spacesDeveloperAfterUpsertHooks = []SpacesDeveloperHook{}
}

func testSpacesDevelopersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesDevelopersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spacesDeveloperColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesDeveloperToOneSpaceUsingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesDeveloper
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
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

	slice := SpacesDeveloperSlice{&local}
	if err = local.L.LoadSpace(ctx, tx, false, (*[]*SpacesDeveloper)(&slice), nil); err != nil {
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

func testSpacesDeveloperToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesDeveloper
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesDeveloperDBTypes, false, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
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

	slice := SpacesDeveloperSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*SpacesDeveloper)(&slice), nil); err != nil {
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

func testSpacesDeveloperToOneSetOpSpaceUsingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesDeveloper
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesDeveloperDBTypes, false, strmangle.SetComplement(spacesDeveloperPrimaryKeyColumns, spacesDeveloperColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesDevelopers[0] != &a {
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
func testSpacesDeveloperToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesDeveloper
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesDeveloperDBTypes, false, strmangle.SetComplement(spacesDeveloperPrimaryKeyColumns, spacesDeveloperColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesDevelopers[0] != &a {
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

func testSpacesDevelopersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
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

func testSpacesDevelopersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesDeveloperSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpacesDevelopersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesDevelopers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spacesDeveloperDBTypes = map[string]string{`SpaceID`: `integer`, `UserID`: `integer`, `SpacesDevelopersPK`: `integer`, `RoleGUID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                      = bytes.MinRead
)

func testSpacesDevelopersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spacesDeveloperPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spacesDeveloperAllColumns) == len(spacesDeveloperPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpacesDevelopersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spacesDeveloperAllColumns) == len(spacesDeveloperPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesDeveloper{}
	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesDeveloperDBTypes, true, spacesDeveloperPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spacesDeveloperAllColumns, spacesDeveloperPrimaryKeyColumns) {
		fields = spacesDeveloperAllColumns
	} else {
		fields = strmangle.SetComplement(
			spacesDeveloperAllColumns,
			spacesDeveloperPrimaryKeyColumns,
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

	slice := SpacesDeveloperSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpacesDevelopersUpsert(t *testing.T) {
	t.Parallel()

	if len(spacesDeveloperAllColumns) == len(spacesDeveloperPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpacesDeveloper{}
	if err = randomize.Struct(seed, &o, spacesDeveloperDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesDeveloper: %s", err)
	}

	count, err := SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spacesDeveloperDBTypes, false, spacesDeveloperPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesDeveloper struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesDeveloper: %s", err)
	}

	count, err = SpacesDevelopers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
