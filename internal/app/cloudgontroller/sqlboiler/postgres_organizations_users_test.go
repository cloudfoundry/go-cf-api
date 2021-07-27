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

func testOrganizationsUsers(t *testing.T) {
	t.Parallel()

	query := OrganizationsUsers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationsUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
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

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationsUsers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsUserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationsUserExists(ctx, tx, o.OrganizationsUsersPK)
	if err != nil {
		t.Errorf("Unable to check if OrganizationsUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationsUserExists to return true, but got false.")
	}
}

func testOrganizationsUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationsUserFound, err := FindOrganizationsUser(ctx, tx, o.OrganizationsUsersPK)
	if err != nil {
		t.Error(err)
	}

	if organizationsUserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationsUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationsUsers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationsUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationsUsers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationsUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationsUserOne := &OrganizationsUser{}
	organizationsUserTwo := &OrganizationsUser{}
	if err = randomize.Struct(seed, organizationsUserOne, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsUserTwo, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationsUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationsUserOne := &OrganizationsUser{}
	organizationsUserTwo := &OrganizationsUser{}
	if err = randomize.Struct(seed, organizationsUserOne, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsUserTwo, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationsUserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func organizationsUserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsUser) error {
	*o = OrganizationsUser{}
	return nil
}

func testOrganizationsUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationsUser{}
	o := &OrganizationsUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser object: %s", err)
	}

	AddOrganizationsUserHook(boil.BeforeInsertHook, organizationsUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationsUserBeforeInsertHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.AfterInsertHook, organizationsUserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationsUserAfterInsertHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.AfterSelectHook, organizationsUserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationsUserAfterSelectHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.BeforeUpdateHook, organizationsUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsUserBeforeUpdateHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.AfterUpdateHook, organizationsUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsUserAfterUpdateHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.BeforeDeleteHook, organizationsUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsUserBeforeDeleteHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.AfterDeleteHook, organizationsUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsUserAfterDeleteHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.BeforeUpsertHook, organizationsUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsUserBeforeUpsertHooks = []OrganizationsUserHook{}

	AddOrganizationsUserHook(boil.AfterUpsertHook, organizationsUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsUserAfterUpsertHooks = []OrganizationsUserHook{}
}

func testOrganizationsUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationsUserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsUserToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsUser
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OrganizationID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Organization().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OrganizationsUserSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*OrganizationsUser)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Organization = nil
	if err = local.L.LoadOrganization(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Organization == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganizationsUserToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsUser
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsUserDBTypes, false, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
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

	slice := OrganizationsUserSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OrganizationsUser)(&slice), nil); err != nil {
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

func testOrganizationsUserToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsUser
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsUserDBTypes, false, strmangle.SetComplement(organizationsUserPrimaryKeyColumns, organizationsUserColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organization{&b, &c} {
		err = a.SetOrganization(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Organization != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OrganizationsUsers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OrganizationID))
		reflect.Indirect(reflect.ValueOf(&a.OrganizationID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OrganizationID != x.ID {
			t.Error("foreign key was wrong value", a.OrganizationID, x.ID)
		}
	}
}
func testOrganizationsUserToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsUser
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsUserDBTypes, false, strmangle.SetComplement(organizationsUserPrimaryKeyColumns, organizationsUserColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationsUsers[0] != &a {
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

func testOrganizationsUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
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

func testOrganizationsUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsUserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationsUserDBTypes = map[string]string{`OrganizationID`: `integer`, `UserID`: `integer`, `OrganizationsUsersPK`: `integer`, `RoleGUID`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_                        = bytes.MinRead
)

func testOrganizationsUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationsUserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationsUserAllColumns) == len(organizationsUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationsUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationsUserAllColumns) == len(organizationsUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsUser{}
	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsUserDBTypes, true, organizationsUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationsUserAllColumns, organizationsUserPrimaryKeyColumns) {
		fields = organizationsUserAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationsUserAllColumns,
			organizationsUserPrimaryKeyColumns,
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

	slice := OrganizationsUserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationsUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationsUserAllColumns) == len(organizationsUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationsUser{}
	if err = randomize.Struct(seed, &o, organizationsUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsUser: %s", err)
	}

	count, err := OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationsUserDBTypes, false, organizationsUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsUser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsUser: %s", err)
	}

	count, err = OrganizationsUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
