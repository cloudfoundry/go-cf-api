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

func testOrganizationsManagers(t *testing.T) {
	t.Parallel()

	query := OrganizationsManagers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationsManagersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
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

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsManagersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationsManagers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsManagersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsManagerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsManagersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationsManagerExists(ctx, tx, o.OrganizationsManagersPK)
	if err != nil {
		t.Errorf("Unable to check if OrganizationsManager exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationsManagerExists to return true, but got false.")
	}
}

func testOrganizationsManagersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationsManagerFound, err := FindOrganizationsManager(ctx, tx, o.OrganizationsManagersPK)
	if err != nil {
		t.Error(err)
	}

	if organizationsManagerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationsManagersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationsManagers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationsManagersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationsManagers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationsManagersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationsManagerOne := &OrganizationsManager{}
	organizationsManagerTwo := &OrganizationsManager{}
	if err = randomize.Struct(seed, organizationsManagerOne, organizationsManagerDBTypes, false, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsManagerTwo, organizationsManagerDBTypes, false, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationsManagersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationsManagerOne := &OrganizationsManager{}
	organizationsManagerTwo := &OrganizationsManager{}
	if err = randomize.Struct(seed, organizationsManagerOne, organizationsManagerDBTypes, false, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsManagerTwo, organizationsManagerDBTypes, false, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsManagerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsManagerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationsManagerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func organizationsManagerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsManager) error {
	*o = OrganizationsManager{}
	return nil
}

func testOrganizationsManagersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationsManager{}
	o := &OrganizationsManager{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager object: %s", err)
	}

	AddOrganizationsManagerHook(boil.BeforeInsertHook, organizationsManagerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationsManagerBeforeInsertHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.AfterInsertHook, organizationsManagerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationsManagerAfterInsertHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.AfterSelectHook, organizationsManagerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationsManagerAfterSelectHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.BeforeUpdateHook, organizationsManagerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsManagerBeforeUpdateHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.AfterUpdateHook, organizationsManagerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsManagerAfterUpdateHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.BeforeDeleteHook, organizationsManagerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsManagerBeforeDeleteHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.AfterDeleteHook, organizationsManagerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsManagerAfterDeleteHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.BeforeUpsertHook, organizationsManagerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsManagerBeforeUpsertHooks = []OrganizationsManagerHook{}

	AddOrganizationsManagerHook(boil.AfterUpsertHook, organizationsManagerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsManagerAfterUpsertHooks = []OrganizationsManagerHook{}
}

func testOrganizationsManagersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsManagersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationsManagerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsManagerToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsManager
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsManagerDBTypes, false, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
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

	slice := OrganizationsManagerSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OrganizationsManager)(&slice), nil); err != nil {
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

func testOrganizationsManagerToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsManager
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsManagerDBTypes, false, strmangle.SetComplement(organizationsManagerPrimaryKeyColumns, organizationsManagerColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationsManagers[0] != &a {
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

func testOrganizationsManagersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
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

func testOrganizationsManagersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsManagerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsManagersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsManagers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationsManagerDBTypes = map[string]string{`OrganizationID`: `integer`, `UserID`: `integer`, `OrganizationsManagersPK`: `integer`, `RoleGUID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                           = bytes.MinRead
)

func testOrganizationsManagersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationsManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationsManagerAllColumns) == len(organizationsManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationsManagersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationsManagerAllColumns) == len(organizationsManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsManager{}
	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsManagerDBTypes, true, organizationsManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationsManagerAllColumns, organizationsManagerPrimaryKeyColumns) {
		fields = organizationsManagerAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationsManagerAllColumns,
			organizationsManagerPrimaryKeyColumns,
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

	slice := OrganizationsManagerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationsManagersUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationsManagerAllColumns) == len(organizationsManagerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationsManager{}
	if err = randomize.Struct(seed, &o, organizationsManagerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsManager: %s", err)
	}

	count, err := OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationsManagerDBTypes, false, organizationsManagerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsManager struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsManager: %s", err)
	}

	count, err = OrganizationsManagers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
