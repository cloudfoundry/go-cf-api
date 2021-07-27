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

func testServiceOfferingLabels(t *testing.T) {
	t.Parallel()

	query := ServiceOfferingLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceOfferingLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
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

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceOfferingLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceOfferingLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceOfferingLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceOfferingLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceOfferingLabelExists to return true, but got false.")
	}
}

func testServiceOfferingLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceOfferingLabelFound, err := FindServiceOfferingLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceOfferingLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceOfferingLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceOfferingLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceOfferingLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceOfferingLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceOfferingLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceOfferingLabelOne := &ServiceOfferingLabel{}
	serviceOfferingLabelTwo := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, serviceOfferingLabelOne, serviceOfferingLabelDBTypes, false, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceOfferingLabelTwo, serviceOfferingLabelDBTypes, false, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOfferingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceOfferingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceOfferingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceOfferingLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceOfferingLabelOne := &ServiceOfferingLabel{}
	serviceOfferingLabelTwo := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, serviceOfferingLabelOne, serviceOfferingLabelDBTypes, false, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceOfferingLabelTwo, serviceOfferingLabelDBTypes, false, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOfferingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceOfferingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceOfferingLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func serviceOfferingLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingLabel) error {
	*o = ServiceOfferingLabel{}
	return nil
}

func testServiceOfferingLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceOfferingLabel{}
	o := &ServiceOfferingLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel object: %s", err)
	}

	AddServiceOfferingLabelHook(boil.BeforeInsertHook, serviceOfferingLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelBeforeInsertHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.AfterInsertHook, serviceOfferingLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelAfterInsertHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.AfterSelectHook, serviceOfferingLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelAfterSelectHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.BeforeUpdateHook, serviceOfferingLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelBeforeUpdateHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.AfterUpdateHook, serviceOfferingLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelAfterUpdateHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.BeforeDeleteHook, serviceOfferingLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelBeforeDeleteHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.AfterDeleteHook, serviceOfferingLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelAfterDeleteHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.BeforeUpsertHook, serviceOfferingLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelBeforeUpsertHooks = []ServiceOfferingLabelHook{}

	AddServiceOfferingLabelHook(boil.AfterUpsertHook, serviceOfferingLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingLabelAfterUpsertHooks = []ServiceOfferingLabelHook{}
}

func testServiceOfferingLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceOfferingLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceOfferingLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceOfferingLabelToOneServiceUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceOfferingLabel
	var foreign Service

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceDBTypes, false, serviceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Service struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ResourceGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Resource().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := ServiceOfferingLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceOfferingLabel)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Resource == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Resource = nil
	if err = local.L.LoadResource(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Resource == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServiceOfferingLabelToOneSetOpServiceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceOfferingLabel
	var b, c Service

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceOfferingLabelDBTypes, false, strmangle.SetComplement(serviceOfferingLabelPrimaryKeyColumns, serviceOfferingLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceDBTypes, false, strmangle.SetComplement(servicePrimaryKeyColumns, serviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceDBTypes, false, strmangle.SetComplement(servicePrimaryKeyColumns, serviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Service{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceOfferingLabels[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ResourceGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.ResourceGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ResourceGUID))
		reflect.Indirect(reflect.ValueOf(&a.ResourceGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ResourceGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.ResourceGUID, x.GUID)
		}
	}
}

func testServiceOfferingLabelToOneRemoveOpServiceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceOfferingLabel
	var b Service

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceOfferingLabelDBTypes, false, strmangle.SetComplement(serviceOfferingLabelPrimaryKeyColumns, serviceOfferingLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceDBTypes, false, strmangle.SetComplement(servicePrimaryKeyColumns, serviceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetResource(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveResource(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Resource().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Resource != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ResourceGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ResourceServiceOfferingLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceOfferingLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
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

func testServiceOfferingLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceOfferingLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceOfferingLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceOfferingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceOfferingLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                           = bytes.MinRead
)

func testServiceOfferingLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceOfferingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceOfferingLabelAllColumns) == len(serviceOfferingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceOfferingLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceOfferingLabelAllColumns) == len(serviceOfferingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingLabel{}
	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceOfferingLabelDBTypes, true, serviceOfferingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceOfferingLabelAllColumns, serviceOfferingLabelPrimaryKeyColumns) {
		fields = serviceOfferingLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceOfferingLabelAllColumns,
			serviceOfferingLabelPrimaryKeyColumns,
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

	slice := ServiceOfferingLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceOfferingLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceOfferingLabelAllColumns) == len(serviceOfferingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceOfferingLabel{}
	if err = randomize.Struct(seed, &o, serviceOfferingLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceOfferingLabel: %s", err)
	}

	count, err := ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceOfferingLabelDBTypes, false, serviceOfferingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceOfferingLabel: %s", err)
	}

	count, err = ServiceOfferingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
