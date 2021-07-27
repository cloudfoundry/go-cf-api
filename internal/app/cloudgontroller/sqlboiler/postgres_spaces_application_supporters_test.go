// +build integration,postgres
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

func testSpacesApplicationSupporters(t *testing.T) {
	t.Parallel()

	query := SpacesApplicationSupporters()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpacesApplicationSupportersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
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

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesApplicationSupportersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpacesApplicationSupporters().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesApplicationSupportersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesApplicationSupporterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesApplicationSupportersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpacesApplicationSupporterExists(ctx, tx, o.SpacesApplicationSupportersPK)
	if err != nil {
		t.Errorf("Unable to check if SpacesApplicationSupporter exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpacesApplicationSupporterExists to return true, but got false.")
	}
}

func testSpacesApplicationSupportersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spacesApplicationSupporterFound, err := FindSpacesApplicationSupporter(ctx, tx, o.SpacesApplicationSupportersPK)
	if err != nil {
		t.Error(err)
	}

	if spacesApplicationSupporterFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpacesApplicationSupportersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpacesApplicationSupporters().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpacesApplicationSupportersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpacesApplicationSupporters().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpacesApplicationSupportersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spacesApplicationSupporterOne := &SpacesApplicationSupporter{}
	spacesApplicationSupporterTwo := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, spacesApplicationSupporterOne, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesApplicationSupporterTwo, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesApplicationSupporterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesApplicationSupporterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesApplicationSupporters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpacesApplicationSupportersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spacesApplicationSupporterOne := &SpacesApplicationSupporter{}
	spacesApplicationSupporterTwo := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, spacesApplicationSupporterOne, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesApplicationSupporterTwo, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesApplicationSupporterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesApplicationSupporterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spacesApplicationSupporterBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func spacesApplicationSupporterAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesApplicationSupporter) error {
	*o = SpacesApplicationSupporter{}
	return nil
}

func testSpacesApplicationSupportersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpacesApplicationSupporter{}
	o := &SpacesApplicationSupporter{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter object: %s", err)
	}

	AddSpacesApplicationSupporterHook(boil.BeforeInsertHook, spacesApplicationSupporterBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterBeforeInsertHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.AfterInsertHook, spacesApplicationSupporterAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterAfterInsertHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.AfterSelectHook, spacesApplicationSupporterAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterAfterSelectHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.BeforeUpdateHook, spacesApplicationSupporterBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterBeforeUpdateHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.AfterUpdateHook, spacesApplicationSupporterAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterAfterUpdateHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.BeforeDeleteHook, spacesApplicationSupporterBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterBeforeDeleteHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.AfterDeleteHook, spacesApplicationSupporterAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterAfterDeleteHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.BeforeUpsertHook, spacesApplicationSupporterBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterBeforeUpsertHooks = []SpacesApplicationSupporterHook{}

	AddSpacesApplicationSupporterHook(boil.AfterUpsertHook, spacesApplicationSupporterAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spacesApplicationSupporterAfterUpsertHooks = []SpacesApplicationSupporterHook{}
}

func testSpacesApplicationSupportersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesApplicationSupportersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spacesApplicationSupporterColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesApplicationSupporterToOneSpaceUsingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesApplicationSupporter
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
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

	slice := SpacesApplicationSupporterSlice{&local}
	if err = local.L.LoadSpace(ctx, tx, false, (*[]*SpacesApplicationSupporter)(&slice), nil); err != nil {
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

func testSpacesApplicationSupporterToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesApplicationSupporter
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
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

	slice := SpacesApplicationSupporterSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*SpacesApplicationSupporter)(&slice), nil); err != nil {
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

func testSpacesApplicationSupporterToOneSetOpSpaceUsingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesApplicationSupporter
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesApplicationSupporterDBTypes, false, strmangle.SetComplement(spacesApplicationSupporterPrimaryKeyColumns, spacesApplicationSupporterColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesApplicationSupporters[0] != &a {
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
func testSpacesApplicationSupporterToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesApplicationSupporter
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesApplicationSupporterDBTypes, false, strmangle.SetComplement(spacesApplicationSupporterPrimaryKeyColumns, spacesApplicationSupporterColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesApplicationSupporters[0] != &a {
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

func testSpacesApplicationSupportersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
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

func testSpacesApplicationSupportersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesApplicationSupporterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpacesApplicationSupportersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesApplicationSupporters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spacesApplicationSupporterDBTypes = map[string]string{`SpacesApplicationSupportersPK`: `integer`, `RoleGUID`: `character varying`, `SpaceID`: `integer`, `UserID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                                 = bytes.MinRead
)

func testSpacesApplicationSupportersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spacesApplicationSupporterPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spacesApplicationSupporterAllColumns) == len(spacesApplicationSupporterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpacesApplicationSupportersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spacesApplicationSupporterAllColumns) == len(spacesApplicationSupporterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesApplicationSupporterDBTypes, true, spacesApplicationSupporterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spacesApplicationSupporterAllColumns, spacesApplicationSupporterPrimaryKeyColumns) {
		fields = spacesApplicationSupporterAllColumns
	} else {
		fields = strmangle.SetComplement(
			spacesApplicationSupporterAllColumns,
			spacesApplicationSupporterPrimaryKeyColumns,
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

	slice := SpacesApplicationSupporterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpacesApplicationSupportersUpsert(t *testing.T) {
	t.Parallel()

	if len(spacesApplicationSupporterAllColumns) == len(spacesApplicationSupporterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpacesApplicationSupporter{}
	if err = randomize.Struct(seed, &o, spacesApplicationSupporterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesApplicationSupporter: %s", err)
	}

	count, err := SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spacesApplicationSupporterDBTypes, false, spacesApplicationSupporterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesApplicationSupporter struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesApplicationSupporter: %s", err)
	}

	count, err = SpacesApplicationSupporters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
