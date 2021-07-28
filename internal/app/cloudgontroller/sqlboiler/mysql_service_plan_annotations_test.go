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

func testServicePlanAnnotations(t *testing.T) {
	t.Parallel()

	query := ServicePlanAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServicePlanAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
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

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServicePlanAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServicePlanAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServicePlanAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServicePlanAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServicePlanAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServicePlanAnnotationExists to return true, but got false.")
	}
}

func testServicePlanAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	servicePlanAnnotationFound, err := FindServicePlanAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if servicePlanAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServicePlanAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServicePlanAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServicePlanAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServicePlanAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServicePlanAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	servicePlanAnnotationOne := &ServicePlanAnnotation{}
	servicePlanAnnotationTwo := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, servicePlanAnnotationOne, servicePlanAnnotationDBTypes, false, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, servicePlanAnnotationTwo, servicePlanAnnotationDBTypes, false, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = servicePlanAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = servicePlanAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServicePlanAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServicePlanAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	servicePlanAnnotationOne := &ServicePlanAnnotation{}
	servicePlanAnnotationTwo := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, servicePlanAnnotationOne, servicePlanAnnotationDBTypes, false, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, servicePlanAnnotationTwo, servicePlanAnnotationDBTypes, false, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = servicePlanAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = servicePlanAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func servicePlanAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func servicePlanAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServicePlanAnnotation) error {
	*o = ServicePlanAnnotation{}
	return nil
}

func testServicePlanAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServicePlanAnnotation{}
	o := &ServicePlanAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation object: %s", err)
	}

	AddServicePlanAnnotationHook(boil.BeforeInsertHook, servicePlanAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationBeforeInsertHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.AfterInsertHook, servicePlanAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationAfterInsertHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.AfterSelectHook, servicePlanAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationAfterSelectHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.BeforeUpdateHook, servicePlanAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationBeforeUpdateHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.AfterUpdateHook, servicePlanAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationAfterUpdateHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.BeforeDeleteHook, servicePlanAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationBeforeDeleteHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.AfterDeleteHook, servicePlanAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationAfterDeleteHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.BeforeUpsertHook, servicePlanAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationBeforeUpsertHooks = []ServicePlanAnnotationHook{}

	AddServicePlanAnnotationHook(boil.AfterUpsertHook, servicePlanAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	servicePlanAnnotationAfterUpsertHooks = []ServicePlanAnnotationHook{}
}

func testServicePlanAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicePlanAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(servicePlanAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServicePlanAnnotationToOneServicePlanUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServicePlanAnnotation
	var foreign ServicePlan

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, servicePlanDBTypes, false, servicePlanColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlan struct: %s", err)
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

	slice := ServicePlanAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServicePlanAnnotation)(&slice), nil); err != nil {
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

func testServicePlanAnnotationToOneSetOpServicePlanUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServicePlanAnnotation
	var b, c ServicePlan

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, servicePlanAnnotationDBTypes, false, strmangle.SetComplement(servicePlanAnnotationPrimaryKeyColumns, servicePlanAnnotationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServicePlanAnnotations[0] != &a {
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

func testServicePlanAnnotationToOneRemoveOpServicePlanUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServicePlanAnnotation
	var b ServicePlan

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, servicePlanAnnotationDBTypes, false, strmangle.SetComplement(servicePlanAnnotationPrimaryKeyColumns, servicePlanAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, servicePlanDBTypes, false, strmangle.SetComplement(servicePlanPrimaryKeyColumns, servicePlanColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServicePlanAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServicePlanAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
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

func testServicePlanAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServicePlanAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServicePlanAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServicePlanAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	servicePlanAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                            = bytes.MinRead
)

func testServicePlanAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(servicePlanAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(servicePlanAnnotationAllColumns) == len(servicePlanAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServicePlanAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(servicePlanAnnotationAllColumns) == len(servicePlanAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServicePlanAnnotation{}
	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, servicePlanAnnotationDBTypes, true, servicePlanAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(servicePlanAnnotationAllColumns, servicePlanAnnotationPrimaryKeyColumns) {
		fields = servicePlanAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			servicePlanAnnotationAllColumns,
			servicePlanAnnotationPrimaryKeyColumns,
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

	slice := ServicePlanAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServicePlanAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(servicePlanAnnotationAllColumns) == len(servicePlanAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServicePlanAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServicePlanAnnotation{}
	if err = randomize.Struct(seed, &o, servicePlanAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServicePlanAnnotation: %s", err)
	}

	count, err := ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, servicePlanAnnotationDBTypes, false, servicePlanAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServicePlanAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServicePlanAnnotation: %s", err)
	}

	count, err = ServicePlanAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
