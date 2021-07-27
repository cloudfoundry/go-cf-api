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

func testServiceInstanceAnnotations(t *testing.T) {
	t.Parallel()

	query := ServiceInstanceAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceInstanceAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
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

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceInstanceAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceInstanceAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceInstanceAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceInstanceAnnotationExists to return true, but got false.")
	}
}

func testServiceInstanceAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceInstanceAnnotationFound, err := FindServiceInstanceAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceInstanceAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceInstanceAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceInstanceAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceInstanceAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceInstanceAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceInstanceAnnotationOne := &ServiceInstanceAnnotation{}
	serviceInstanceAnnotationTwo := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, serviceInstanceAnnotationOne, serviceInstanceAnnotationDBTypes, false, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceAnnotationTwo, serviceInstanceAnnotationDBTypes, false, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceInstanceAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceInstanceAnnotationOne := &ServiceInstanceAnnotation{}
	serviceInstanceAnnotationTwo := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, serviceInstanceAnnotationOne, serviceInstanceAnnotationDBTypes, false, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceAnnotationTwo, serviceInstanceAnnotationDBTypes, false, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceInstanceAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func serviceInstanceAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceAnnotation) error {
	*o = ServiceInstanceAnnotation{}
	return nil
}

func testServiceInstanceAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceInstanceAnnotation{}
	o := &ServiceInstanceAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation object: %s", err)
	}

	AddServiceInstanceAnnotationHook(boil.BeforeInsertHook, serviceInstanceAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationBeforeInsertHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.AfterInsertHook, serviceInstanceAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationAfterInsertHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.AfterSelectHook, serviceInstanceAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationAfterSelectHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.BeforeUpdateHook, serviceInstanceAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationBeforeUpdateHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.AfterUpdateHook, serviceInstanceAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationAfterUpdateHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.BeforeDeleteHook, serviceInstanceAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationBeforeDeleteHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.AfterDeleteHook, serviceInstanceAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationAfterDeleteHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.BeforeUpsertHook, serviceInstanceAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationBeforeUpsertHooks = []ServiceInstanceAnnotationHook{}

	AddServiceInstanceAnnotationHook(boil.AfterUpsertHook, serviceInstanceAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceAnnotationAfterUpsertHooks = []ServiceInstanceAnnotationHook{}
}

func testServiceInstanceAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceInstanceAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceAnnotationToOneServiceInstanceUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceInstanceAnnotation
	var foreign ServiceInstance

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceInstanceDBTypes, false, serviceInstanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstance struct: %s", err)
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

	slice := ServiceInstanceAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceInstanceAnnotation)(&slice), nil); err != nil {
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

func testServiceInstanceAnnotationToOneSetOpServiceInstanceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceAnnotation
	var b, c ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceAnnotationDBTypes, false, strmangle.SetComplement(serviceInstanceAnnotationPrimaryKeyColumns, serviceInstanceAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceInstance{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceInstanceAnnotations[0] != &a {
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

func testServiceInstanceAnnotationToOneRemoveOpServiceInstanceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceAnnotation
	var b ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceAnnotationDBTypes, false, strmangle.SetComplement(serviceInstanceAnnotationPrimaryKeyColumns, serviceInstanceAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceInstanceAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceInstanceAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
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

func testServiceInstanceAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceInstanceAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                                = bytes.MinRead
)

func testServiceInstanceAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceInstanceAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceInstanceAnnotationAllColumns) == len(serviceInstanceAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceInstanceAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceAnnotationAllColumns) == len(serviceInstanceAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceAnnotationDBTypes, true, serviceInstanceAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceInstanceAnnotationAllColumns, serviceInstanceAnnotationPrimaryKeyColumns) {
		fields = serviceInstanceAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceInstanceAnnotationAllColumns,
			serviceInstanceAnnotationPrimaryKeyColumns,
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

	slice := ServiceInstanceAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceInstanceAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceAnnotationAllColumns) == len(serviceInstanceAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceInstanceAnnotation{}
	if err = randomize.Struct(seed, &o, serviceInstanceAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceAnnotation: %s", err)
	}

	count, err := ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceInstanceAnnotationDBTypes, false, serviceInstanceAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceAnnotation: %s", err)
	}

	count, err = ServiceInstanceAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
