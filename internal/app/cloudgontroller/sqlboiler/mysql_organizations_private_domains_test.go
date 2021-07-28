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

func testOrganizationsPrivateDomains(t *testing.T) {
	t.Parallel()

	query := OrganizationsPrivateDomains()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationsPrivateDomainsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
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

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsPrivateDomainsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationsPrivateDomains().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsPrivateDomainsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsPrivateDomainSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsPrivateDomainsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationsPrivateDomainExists(ctx, tx, o.OrganizationsPrivateDomainsPK)
	if err != nil {
		t.Errorf("Unable to check if OrganizationsPrivateDomain exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationsPrivateDomainExists to return true, but got false.")
	}
}

func testOrganizationsPrivateDomainsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationsPrivateDomainFound, err := FindOrganizationsPrivateDomain(ctx, tx, o.OrganizationsPrivateDomainsPK)
	if err != nil {
		t.Error(err)
	}

	if organizationsPrivateDomainFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationsPrivateDomainsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationsPrivateDomains().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationsPrivateDomainsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationsPrivateDomains().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationsPrivateDomainsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationsPrivateDomainOne := &OrganizationsPrivateDomain{}
	organizationsPrivateDomainTwo := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, organizationsPrivateDomainOne, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsPrivateDomainTwo, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsPrivateDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsPrivateDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsPrivateDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationsPrivateDomainsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationsPrivateDomainOne := &OrganizationsPrivateDomain{}
	organizationsPrivateDomainTwo := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, organizationsPrivateDomainOne, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationsPrivateDomainTwo, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationsPrivateDomainOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationsPrivateDomainTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationsPrivateDomainBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func organizationsPrivateDomainAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationsPrivateDomain) error {
	*o = OrganizationsPrivateDomain{}
	return nil
}

func testOrganizationsPrivateDomainsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationsPrivateDomain{}
	o := &OrganizationsPrivateDomain{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain object: %s", err)
	}

	AddOrganizationsPrivateDomainHook(boil.BeforeInsertHook, organizationsPrivateDomainBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainBeforeInsertHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.AfterInsertHook, organizationsPrivateDomainAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainAfterInsertHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.AfterSelectHook, organizationsPrivateDomainAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainAfterSelectHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.BeforeUpdateHook, organizationsPrivateDomainBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainBeforeUpdateHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.AfterUpdateHook, organizationsPrivateDomainAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainAfterUpdateHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.BeforeDeleteHook, organizationsPrivateDomainBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainBeforeDeleteHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.AfterDeleteHook, organizationsPrivateDomainAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainAfterDeleteHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.BeforeUpsertHook, organizationsPrivateDomainBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainBeforeUpsertHooks = []OrganizationsPrivateDomainHook{}

	AddOrganizationsPrivateDomainHook(boil.AfterUpsertHook, organizationsPrivateDomainAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationsPrivateDomainAfterUpsertHooks = []OrganizationsPrivateDomainHook{}
}

func testOrganizationsPrivateDomainsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsPrivateDomainsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationsPrivateDomainColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsPrivateDomainToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsPrivateDomain
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
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

	slice := OrganizationsPrivateDomainSlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*OrganizationsPrivateDomain)(&slice), nil); err != nil {
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

func testOrganizationsPrivateDomainToOneDomainUsingPrivateDomain(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationsPrivateDomain
	var foreign Domain

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, domainDBTypes, false, domainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Domain struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PrivateDomainID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.PrivateDomain().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OrganizationsPrivateDomainSlice{&local}
	if err = local.L.LoadPrivateDomain(ctx, tx, false, (*[]*OrganizationsPrivateDomain)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PrivateDomain == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.PrivateDomain = nil
	if err = local.L.LoadPrivateDomain(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.PrivateDomain == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOrganizationsPrivateDomainToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsPrivateDomain
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsPrivateDomainDBTypes, false, strmangle.SetComplement(organizationsPrivateDomainPrimaryKeyColumns, organizationsPrivateDomainColumnsWithoutDefault)...); err != nil {
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

		if x.R.OrganizationsPrivateDomains[0] != &a {
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
func testOrganizationsPrivateDomainToOneSetOpDomainUsingPrivateDomain(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationsPrivateDomain
	var b, c Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationsPrivateDomainDBTypes, false, strmangle.SetComplement(organizationsPrivateDomainPrimaryKeyColumns, organizationsPrivateDomainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Domain{&b, &c} {
		err = a.SetPrivateDomain(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.PrivateDomain != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.PrivateDomainOrganizationsPrivateDomains[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PrivateDomainID != x.ID {
			t.Error("foreign key was wrong value", a.PrivateDomainID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PrivateDomainID))
		reflect.Indirect(reflect.ValueOf(&a.PrivateDomainID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PrivateDomainID != x.ID {
			t.Error("foreign key was wrong value", a.PrivateDomainID, x.ID)
		}
	}
}

func testOrganizationsPrivateDomainsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
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

func testOrganizationsPrivateDomainsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationsPrivateDomainSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsPrivateDomainsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationsPrivateDomains().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationsPrivateDomainDBTypes = map[string]string{`OrganizationID`: `int`, `PrivateDomainID`: `int`, `OrganizationsPrivateDomainsPK`: `int`}
	_                                 = bytes.MinRead
)

func testOrganizationsPrivateDomainsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationsPrivateDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationsPrivateDomainAllColumns) == len(organizationsPrivateDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationsPrivateDomainsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationsPrivateDomainAllColumns) == len(organizationsPrivateDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationsPrivateDomainDBTypes, true, organizationsPrivateDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationsPrivateDomainAllColumns, organizationsPrivateDomainPrimaryKeyColumns) {
		fields = organizationsPrivateDomainAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationsPrivateDomainAllColumns,
			organizationsPrivateDomainPrimaryKeyColumns,
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

	slice := OrganizationsPrivateDomainSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationsPrivateDomainsUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationsPrivateDomainAllColumns) == len(organizationsPrivateDomainPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLOrganizationsPrivateDomainUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationsPrivateDomain{}
	if err = randomize.Struct(seed, &o, organizationsPrivateDomainDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsPrivateDomain: %s", err)
	}

	count, err := OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationsPrivateDomainDBTypes, false, organizationsPrivateDomainPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationsPrivateDomain struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationsPrivateDomain: %s", err)
	}

	count, err = OrganizationsPrivateDomains().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
