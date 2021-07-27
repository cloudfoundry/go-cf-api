// +build integration mysql
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

func testServiceBrokerUpdateRequestAnnotations(t *testing.T) {
	t.Parallel()

	query := ServiceBrokerUpdateRequestAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceBrokerUpdateRequestAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
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

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceBrokerUpdateRequestAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerUpdateRequestAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceBrokerUpdateRequestAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceBrokerUpdateRequestAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceBrokerUpdateRequestAnnotationExists to return true, but got false.")
	}
}

func testServiceBrokerUpdateRequestAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceBrokerUpdateRequestAnnotationFound, err := FindServiceBrokerUpdateRequestAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceBrokerUpdateRequestAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceBrokerUpdateRequestAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceBrokerUpdateRequestAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerUpdateRequestAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceBrokerUpdateRequestAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceBrokerUpdateRequestAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceBrokerUpdateRequestAnnotationOne := &ServiceBrokerUpdateRequestAnnotation{}
	serviceBrokerUpdateRequestAnnotationTwo := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestAnnotationOne, serviceBrokerUpdateRequestAnnotationDBTypes, false, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestAnnotationTwo, serviceBrokerUpdateRequestAnnotationDBTypes, false, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerUpdateRequestAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerUpdateRequestAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerUpdateRequestAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceBrokerUpdateRequestAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceBrokerUpdateRequestAnnotationOne := &ServiceBrokerUpdateRequestAnnotation{}
	serviceBrokerUpdateRequestAnnotationTwo := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestAnnotationOne, serviceBrokerUpdateRequestAnnotationDBTypes, false, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestAnnotationTwo, serviceBrokerUpdateRequestAnnotationDBTypes, false, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerUpdateRequestAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerUpdateRequestAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBrokerUpdateRequestAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func serviceBrokerUpdateRequestAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestAnnotation) error {
	*o = ServiceBrokerUpdateRequestAnnotation{}
	return nil
}

func testServiceBrokerUpdateRequestAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceBrokerUpdateRequestAnnotation{}
	o := &ServiceBrokerUpdateRequestAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation object: %s", err)
	}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.BeforeInsertHook, serviceBrokerUpdateRequestAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationBeforeInsertHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.AfterInsertHook, serviceBrokerUpdateRequestAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationAfterInsertHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.AfterSelectHook, serviceBrokerUpdateRequestAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationAfterSelectHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.BeforeUpdateHook, serviceBrokerUpdateRequestAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationBeforeUpdateHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.AfterUpdateHook, serviceBrokerUpdateRequestAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationAfterUpdateHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.BeforeDeleteHook, serviceBrokerUpdateRequestAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationBeforeDeleteHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.AfterDeleteHook, serviceBrokerUpdateRequestAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationAfterDeleteHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.BeforeUpsertHook, serviceBrokerUpdateRequestAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationBeforeUpsertHooks = []ServiceBrokerUpdateRequestAnnotationHook{}

	AddServiceBrokerUpdateRequestAnnotationHook(boil.AfterUpsertHook, serviceBrokerUpdateRequestAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestAnnotationAfterUpsertHooks = []ServiceBrokerUpdateRequestAnnotationHook{}
}

func testServiceBrokerUpdateRequestAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerUpdateRequestAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceBrokerUpdateRequestAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerUpdateRequestAnnotationToOneServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceBrokerUpdateRequestAnnotation
	var foreign ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceBrokerUpdateRequestDBTypes, false, serviceBrokerUpdateRequestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequest struct: %s", err)
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

	slice := ServiceBrokerUpdateRequestAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceBrokerUpdateRequestAnnotation)(&slice), nil); err != nil {
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

func testServiceBrokerUpdateRequestAnnotationToOneSetOpServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerUpdateRequestAnnotation
	var b, c ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerUpdateRequestAnnotationDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns, serviceBrokerUpdateRequestAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceBrokerUpdateRequest{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceBrokerUpdateRequestAnnotations[0] != &a {
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

func testServiceBrokerUpdateRequestAnnotationToOneRemoveOpServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerUpdateRequestAnnotation
	var b ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerUpdateRequestAnnotationDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns, serviceBrokerUpdateRequestAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceBrokerUpdateRequestAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceBrokerUpdateRequestAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
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

func testServiceBrokerUpdateRequestAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerUpdateRequestAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerUpdateRequestAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerUpdateRequestAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceBrokerUpdateRequestAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                                           = bytes.MinRead
)

func testServiceBrokerUpdateRequestAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceBrokerUpdateRequestAnnotationAllColumns) == len(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceBrokerUpdateRequestAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerUpdateRequestAnnotationAllColumns) == len(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestAnnotationDBTypes, true, serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceBrokerUpdateRequestAnnotationAllColumns, serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns) {
		fields = serviceBrokerUpdateRequestAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceBrokerUpdateRequestAnnotationAllColumns,
			serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns,
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

	slice := ServiceBrokerUpdateRequestAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceBrokerUpdateRequestAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerUpdateRequestAnnotationAllColumns) == len(serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceBrokerUpdateRequestAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceBrokerUpdateRequestAnnotation{}
	if err = randomize.Struct(seed, &o, serviceBrokerUpdateRequestAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerUpdateRequestAnnotation: %s", err)
	}

	count, err := ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceBrokerUpdateRequestAnnotationDBTypes, false, serviceBrokerUpdateRequestAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerUpdateRequestAnnotation: %s", err)
	}

	count, err = ServiceBrokerUpdateRequestAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
