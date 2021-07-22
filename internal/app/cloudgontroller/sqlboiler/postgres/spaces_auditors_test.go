// +build integration
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

func testSpacesAuditors(t *testing.T) {
	t.Parallel()

	query := SpacesAuditors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpacesAuditorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
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

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesAuditorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpacesAuditors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesAuditorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesAuditorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpacesAuditorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpacesAuditorExists(ctx, tx, o.SpacesAuditorsPK)
	if err != nil {
		t.Errorf("Unable to check if SpacesAuditor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpacesAuditorExists to return true, but got false.")
	}
}

func testSpacesAuditorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spacesAuditorFound, err := FindSpacesAuditor(ctx, tx, o.SpacesAuditorsPK)
	if err != nil {
		t.Error(err)
	}

	if spacesAuditorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpacesAuditorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpacesAuditors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpacesAuditorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpacesAuditors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpacesAuditorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spacesAuditorOne := &SpacesAuditor{}
	spacesAuditorTwo := &SpacesAuditor{}
	if err = randomize.Struct(seed, spacesAuditorOne, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesAuditorTwo, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesAuditorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesAuditorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesAuditors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpacesAuditorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spacesAuditorOne := &SpacesAuditor{}
	spacesAuditorTwo := &SpacesAuditor{}
	if err = randomize.Struct(seed, spacesAuditorOne, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}
	if err = randomize.Struct(seed, spacesAuditorTwo, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spacesAuditorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spacesAuditorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spacesAuditorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func spacesAuditorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpacesAuditor) error {
	*o = SpacesAuditor{}
	return nil
}

func testSpacesAuditorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpacesAuditor{}
	o := &SpacesAuditor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor object: %s", err)
	}

	AddSpacesAuditorHook(boil.BeforeInsertHook, spacesAuditorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spacesAuditorBeforeInsertHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.AfterInsertHook, spacesAuditorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spacesAuditorAfterInsertHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.AfterSelectHook, spacesAuditorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spacesAuditorAfterSelectHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.BeforeUpdateHook, spacesAuditorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spacesAuditorBeforeUpdateHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.AfterUpdateHook, spacesAuditorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spacesAuditorAfterUpdateHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.BeforeDeleteHook, spacesAuditorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spacesAuditorBeforeDeleteHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.AfterDeleteHook, spacesAuditorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spacesAuditorAfterDeleteHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.BeforeUpsertHook, spacesAuditorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spacesAuditorBeforeUpsertHooks = []SpacesAuditorHook{}

	AddSpacesAuditorHook(boil.AfterUpsertHook, spacesAuditorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spacesAuditorAfterUpsertHooks = []SpacesAuditorHook{}
}

func testSpacesAuditorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesAuditorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spacesAuditorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpacesAuditorToOneSpaceUsingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesAuditor
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
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

	slice := SpacesAuditorSlice{&local}
	if err = local.L.LoadSpace(ctx, tx, false, (*[]*SpacesAuditor)(&slice), nil); err != nil {
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

func testSpacesAuditorToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SpacesAuditor
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, spacesAuditorDBTypes, false, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
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

	slice := SpacesAuditorSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*SpacesAuditor)(&slice), nil); err != nil {
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

func testSpacesAuditorToOneSetOpSpaceUsingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesAuditor
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesAuditorDBTypes, false, strmangle.SetComplement(spacesAuditorPrimaryKeyColumns, spacesAuditorColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesAuditors[0] != &a {
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
func testSpacesAuditorToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpacesAuditor
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spacesAuditorDBTypes, false, strmangle.SetComplement(spacesAuditorPrimaryKeyColumns, spacesAuditorColumnsWithoutDefault)...); err != nil {
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

		if x.R.SpacesAuditors[0] != &a {
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

func testSpacesAuditorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
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

func testSpacesAuditorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpacesAuditorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpacesAuditorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpacesAuditors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spacesAuditorDBTypes = map[string]string{`SpaceID`: `integer`, `UserID`: `integer`, `SpacesAuditorsPK`: `integer`, `RoleGUID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                    = bytes.MinRead
)

func testSpacesAuditorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spacesAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spacesAuditorAllColumns) == len(spacesAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpacesAuditorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spacesAuditorAllColumns) == len(spacesAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpacesAuditor{}
	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spacesAuditorDBTypes, true, spacesAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spacesAuditorAllColumns, spacesAuditorPrimaryKeyColumns) {
		fields = spacesAuditorAllColumns
	} else {
		fields = strmangle.SetComplement(
			spacesAuditorAllColumns,
			spacesAuditorPrimaryKeyColumns,
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

	slice := SpacesAuditorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpacesAuditorsUpsert(t *testing.T) {
	t.Parallel()

	if len(spacesAuditorAllColumns) == len(spacesAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpacesAuditor{}
	if err = randomize.Struct(seed, &o, spacesAuditorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesAuditor: %s", err)
	}

	count, err := SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spacesAuditorDBTypes, false, spacesAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpacesAuditor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpacesAuditor: %s", err)
	}

	count, err = SpacesAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
