// +build mysql,db
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

func testServicePlanVisibilities(t *testing.T) {
	t.Parallel()

	query := ServicePlanVisibilities()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServicePlanVisibilitiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
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

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanVisibilitiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServicePlanVisibilities().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanVisibilitiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServicePlanVisibilitySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanVisibilitiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServicePlanVisibilityExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServicePlanVisibility exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServicePlanVisibilityExists to return true, but got false.")
	}
}

func testServicePlanVisibilitiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	servicePlanVisibilityFound, err := FindServicePlanVisibility(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if servicePlanVisibilityFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServicePlanVisibilitiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServicePlanVisibilities().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServicePlanVisibilitiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServicePlanVisibilities().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServicePlanVisibilitiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	servicePlanVisibilityOne := &ServicePlanVisibility{}
	servicePlanVisibilityTwo := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, servicePlanVisibilityOne, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}
	if err = randomize.Struct(seed, servicePlanVisibilityTwo, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = servicePlanVisibilityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = servicePlanVisibilityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServicePlanVisibilities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServicePlanVisibilitiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	servicePlanVisibilityOne := &ServicePlanVisibility{}
	servicePlanVisibilityTwo := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, servicePlanVisibilityOne, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}
	if err = randomize.Struct(seed, servicePlanVisibilityTwo, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = servicePlanVisibilityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = servicePlanVisibilityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func servicePlanVisibilityBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func servicePlanVisibilityAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanVisibility) error {
	*o = ServicePlanVisibility{}
	return nil
}

func testServicePlanVisibilitiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServicePlanVisibility{}
	o := &ServicePlanVisibility{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility object: %s", err)
	}

	AddServicePlanVisibilityHook(boil.BeforeInsertHook, servicePlanVisibilityBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityBeforeInsertHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.AfterInsertHook, servicePlanVisibilityAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityAfterInsertHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.AfterSelectHook, servicePlanVisibilityAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityAfterSelectHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.BeforeUpdateHook, servicePlanVisibilityBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityBeforeUpdateHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.AfterUpdateHook, servicePlanVisibilityAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityAfterUpdateHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.BeforeDeleteHook, servicePlanVisibilityBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityBeforeDeleteHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.AfterDeleteHook, servicePlanVisibilityAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityAfterDeleteHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.BeforeUpsertHook, servicePlanVisibilityBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityBeforeUpsertHooks = []ServicePlanVisibilityHook{}

	AddServicePlanVisibilityHook(boil.AfterUpsertHook, servicePlanVisibilityAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	servicePlanVisibilityAfterUpsertHooks = []ServicePlanVisibilityHook{}
}

func testServicePlanVisibilitiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicePlanVisibilitiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(servicePlanVisibilityColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicePlanVisibilityToOneOrganizationUsingOrganization(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServicePlanVisibility
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
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

	slice := ServicePlanVisibilitySlice{&local}
	if err = local.L.LoadOrganization(ctx, tx, false, (*[]*ServicePlanVisibility)(&slice), nil); err != nil {
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

func testServicePlanVisibilityToOneServicePlanUsingServicePlan(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServicePlanVisibility
	var foreign ServicePlan

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, servicePlanVisibilityDBTypes, false, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, servicePlanDBTypes, false, servicePlanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlan struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ServicePlanID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServicePlan().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ServicePlanVisibilitySlice{&local}
	if err = local.L.LoadServicePlan(ctx, tx, false, (*[]*ServicePlanVisibility)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServicePlan == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ServicePlan = nil
	if err = local.L.LoadServicePlan(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServicePlan == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServicePlanVisibilityToOneSetOpOrganizationUsingOrganization(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServicePlanVisibility
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, servicePlanVisibilityDBTypes, false, strmangle.SetComplement(servicePlanVisibilityPrimaryKeyColumns, servicePlanVisibilityColumnsWithoutDefault)...); err != nil {
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

		if x.R.ServicePlanVisibilities[0] != &a {
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
func testServicePlanVisibilityToOneSetOpServicePlanUsingServicePlan(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServicePlanVisibility
	var b, c ServicePlan

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, servicePlanVisibilityDBTypes, false, strmangle.SetComplement(servicePlanVisibilityPrimaryKeyColumns, servicePlanVisibilityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, servicePlanDBTypes, false, strmangle.SetComplement(servicePlanPrimaryKeyColumns, servicePlanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, servicePlanDBTypes, false, strmangle.SetComplement(servicePlanPrimaryKeyColumns, servicePlanColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServicePlan{&b, &c} {
		err = a.SetServicePlan(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ServicePlan != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServicePlanVisibilities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ServicePlanID != x.ID {
			t.Error("foreign key was wrong value", a.ServicePlanID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServicePlanID))
		reflect.Indirect(reflect.ValueOf(&a.ServicePlanID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ServicePlanID != x.ID {
			t.Error("foreign key was wrong value", a.ServicePlanID, x.ID)
		}
	}
}

func testServicePlanVisibilitiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
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

func testServicePlanVisibilitiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServicePlanVisibilitySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServicePlanVisibilitiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServicePlanVisibilities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	servicePlanVisibilityDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ServicePlanID`: `int`, `OrganizationID`: `int`}
	_                            = bytes.MinRead
)

func testServicePlanVisibilitiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(servicePlanVisibilityPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(servicePlanVisibilityAllColumns) == len(servicePlanVisibilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServicePlanVisibilitiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(servicePlanVisibilityAllColumns) == len(servicePlanVisibilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanVisibility{}
	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, servicePlanVisibilityDBTypes, true, servicePlanVisibilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(servicePlanVisibilityAllColumns, servicePlanVisibilityPrimaryKeyColumns) {
		fields = servicePlanVisibilityAllColumns
	} else {
		fields = strmangle.SetComplement(
			servicePlanVisibilityAllColumns,
			servicePlanVisibilityPrimaryKeyColumns,
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

	slice := ServicePlanVisibilitySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServicePlanVisibilitiesUpsert(t *testing.T) {
	t.Parallel()

	if len(servicePlanVisibilityAllColumns) == len(servicePlanVisibilityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServicePlanVisibilityUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServicePlanVisibility{}
	if err = randomize.Struct(seed, &o, servicePlanVisibilityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServicePlanVisibility: %s", err)
	}

	count, err := ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, servicePlanVisibilityDBTypes, false, servicePlanVisibilityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanVisibility struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServicePlanVisibility: %s", err)
	}

	count, err = ServicePlanVisibilities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
