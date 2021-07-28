// +build mysql_integration
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

func testServiceOfferingAnnotations(t *testing.T) {
	t.Parallel()

	query := ServiceOfferingAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceOfferingAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
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

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceOfferingAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceOfferingAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceOfferingAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceOfferingAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceOfferingAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceOfferingAnnotationExists to return true, but got false.")
	}
}

func testServiceOfferingAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceOfferingAnnotationFound, err := FindServiceOfferingAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceOfferingAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceOfferingAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceOfferingAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceOfferingAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceOfferingAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceOfferingAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceOfferingAnnotationOne := &ServiceOfferingAnnotation{}
	serviceOfferingAnnotationTwo := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, serviceOfferingAnnotationOne, serviceOfferingAnnotationDBTypes, false, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceOfferingAnnotationTwo, serviceOfferingAnnotationDBTypes, false, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOfferingAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceOfferingAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceOfferingAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceOfferingAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceOfferingAnnotationOne := &ServiceOfferingAnnotation{}
	serviceOfferingAnnotationTwo := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, serviceOfferingAnnotationOne, serviceOfferingAnnotationDBTypes, false, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceOfferingAnnotationTwo, serviceOfferingAnnotationDBTypes, false, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceOfferingAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceOfferingAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceOfferingAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func serviceOfferingAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceOfferingAnnotation) error {
	*o = ServiceOfferingAnnotation{}
	return nil
}

func testServiceOfferingAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceOfferingAnnotation{}
	o := &ServiceOfferingAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation object: %s", err)
	}

	AddServiceOfferingAnnotationHook(boil.BeforeInsertHook, serviceOfferingAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationBeforeInsertHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.AfterInsertHook, serviceOfferingAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationAfterInsertHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.AfterSelectHook, serviceOfferingAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationAfterSelectHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.BeforeUpdateHook, serviceOfferingAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationBeforeUpdateHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.AfterUpdateHook, serviceOfferingAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationAfterUpdateHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.BeforeDeleteHook, serviceOfferingAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationBeforeDeleteHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.AfterDeleteHook, serviceOfferingAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationAfterDeleteHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.BeforeUpsertHook, serviceOfferingAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationBeforeUpsertHooks = []ServiceOfferingAnnotationHook{}

	AddServiceOfferingAnnotationHook(boil.AfterUpsertHook, serviceOfferingAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceOfferingAnnotationAfterUpsertHooks = []ServiceOfferingAnnotationHook{}
}

func testServiceOfferingAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceOfferingAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceOfferingAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceOfferingAnnotationToOneServiceUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceOfferingAnnotation
	var foreign Service

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
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

	slice := ServiceOfferingAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceOfferingAnnotation)(&slice), nil); err != nil {
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

func testServiceOfferingAnnotationToOneSetOpServiceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceOfferingAnnotation
	var b, c Service

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceOfferingAnnotationDBTypes, false, strmangle.SetComplement(serviceOfferingAnnotationPrimaryKeyColumns, serviceOfferingAnnotationColumnsWithoutDefault)...); err != nil {
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

		if x.R.ResourceServiceOfferingAnnotations[0] != &a {
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

func testServiceOfferingAnnotationToOneRemoveOpServiceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceOfferingAnnotation
	var b Service

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceOfferingAnnotationDBTypes, false, strmangle.SetComplement(serviceOfferingAnnotationPrimaryKeyColumns, serviceOfferingAnnotationColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceOfferingAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceOfferingAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
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

func testServiceOfferingAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceOfferingAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceOfferingAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceOfferingAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceOfferingAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                                = bytes.MinRead
)

func testServiceOfferingAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceOfferingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceOfferingAnnotationAllColumns) == len(serviceOfferingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceOfferingAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceOfferingAnnotationAllColumns) == len(serviceOfferingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceOfferingAnnotationDBTypes, true, serviceOfferingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceOfferingAnnotationAllColumns, serviceOfferingAnnotationPrimaryKeyColumns) {
		fields = serviceOfferingAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceOfferingAnnotationAllColumns,
			serviceOfferingAnnotationPrimaryKeyColumns,
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

	slice := ServiceOfferingAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceOfferingAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceOfferingAnnotationAllColumns) == len(serviceOfferingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceOfferingAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceOfferingAnnotation{}
	if err = randomize.Struct(seed, &o, serviceOfferingAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceOfferingAnnotation: %s", err)
	}

	count, err := ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceOfferingAnnotationDBTypes, false, serviceOfferingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceOfferingAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceOfferingAnnotation: %s", err)
	}

	count, err = ServiceOfferingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
