// +build integration,mysql
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

func testOrganizationsAuditors(t *testing.T) {
	t.Parallel()

	query := OrganizationsAuditors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationsAuditorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
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

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsAuditorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationsAuditors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsAuditorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsAuditorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsAuditorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationsAuditorExists(ctx, tx, o.OrganizationsAuditorsPK)
	if err != nil {
		t.Errorf("Unable to check if OrganizationsAuditor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationsAuditorExists to return true, but got false.")
	}
}

func testOrganizationsAuditorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationsAuditorFound, err := FindOrganizationsAuditor(ctx, tx, o.OrganizationsAuditorsPK)
	if err != nil {
		t.Error(err)
	}

	if organizationsAuditorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationsAuditorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationsAuditors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationsAuditorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationsAuditors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationsAuditorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationsAuditorOne := &OrganizationsAuditor{}
	organizationsAuditorTwo := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, organizationsAuditorOne, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsAuditorTwo, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsAuditorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsAuditorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsAuditors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationsAuditorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationsAuditorOne := &OrganizationsAuditor{}
	organizationsAuditorTwo := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, organizationsAuditorOne, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsAuditorTwo, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsAuditorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsAuditorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationsAuditorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func organizationsAuditorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsAuditor) error {
	*o = OrganizationsAuditor{}
	return nil
}

func testOrganizationsAuditorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationsAuditor{}
	o := &OrganizationsAuditor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor object: %s", err)
	}

	AddOrganizationsAuditorHook(boil.BeforeInsertHook, organizationsAuditorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorBeforeInsertHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.AfterInsertHook, organizationsAuditorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorAfterInsertHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.AfterSelectHook, organizationsAuditorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorAfterSelectHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.BeforeUpdateHook, organizationsAuditorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorBeforeUpdateHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.AfterUpdateHook, organizationsAuditorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorAfterUpdateHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.BeforeDeleteHook, organizationsAuditorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorBeforeDeleteHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.AfterDeleteHook, organizationsAuditorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorAfterDeleteHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.BeforeUpsertHook, organizationsAuditorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorBeforeUpsertHooks = []OrganizationsAuditorHook{}

	AddOrganizationsAuditorHook(boil.AfterUpsertHook, organizationsAuditorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsAuditorAfterUpsertHooks = []OrganizationsAuditorHook{}
}

func testOrganizationsAuditorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsAuditorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationsAuditorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsAuditorToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsAuditor
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
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

	slice := OrganizationsAuditorSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*OrganizationsAuditor)(&slice), nil); err != nil {
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

func testOrganizationsAuditorToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsAuditor
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsAuditorDBTypes, false, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
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

	slice := OrganizationsAuditorSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*OrganizationsAuditor)(&slice), nil); err != nil {
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

func testOrganizationsAuditorToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsAuditor
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsAuditorDBTypes, false, strmangle.SetComplement(organizationsAuditorPrimaryKeyColumns, organizationsAuditorColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationsAuditors[0] != &a {
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
func testOrganizationsAuditorToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsAuditor
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsAuditorDBTypes, false, strmangle.SetComplement(organizationsAuditorPrimaryKeyColumns, organizationsAuditorColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationsAuditors[0] != &a {
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

func testOrganizationsAuditorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
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

func testOrganizationsAuditorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsAuditorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsAuditorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsAuditors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationsAuditorDBTypes = map[string]string{`OrganizationID`: `int`, `UserID`: `int`, `OrganizationsAuditorsPK`: `int`, `RoleGUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_                           = bytes.MinRead
)

func testOrganizationsAuditorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationsAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationsAuditorAllColumns) == len(organizationsAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationsAuditorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationsAuditorAllColumns) == len(organizationsAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsAuditor{}
	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsAuditorDBTypes, true, organizationsAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationsAuditorAllColumns, organizationsAuditorPrimaryKeyColumns) {
		fields = organizationsAuditorAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationsAuditorAllColumns,
			organizationsAuditorPrimaryKeyColumns,
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

	slice := OrganizationsAuditorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationsAuditorsUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationsAuditorAllColumns) == len(organizationsAuditorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOrganizationsAuditorUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationsAuditor{}
	if err = randomize.Struct(seed, &o, organizationsAuditorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsAuditor: %s", err)
	}

	count, err := OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationsAuditorDBTypes, false, organizationsAuditorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsAuditor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsAuditor: %s", err)
	}

	count, err = OrganizationsAuditors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
