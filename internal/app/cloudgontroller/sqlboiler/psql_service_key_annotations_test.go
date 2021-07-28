// +build psql,db
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

func testServiceKeyAnnotations(t *testing.T) {
	t.Parallel()

	query := ServiceKeyAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceKeyAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
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

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceKeyAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceKeyAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceKeyAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceKeyAnnotationExists to return true, but got false.")
	}
}

func testServiceKeyAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceKeyAnnotationFound, err := FindServiceKeyAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceKeyAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceKeyAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceKeyAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceKeyAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceKeyAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceKeyAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceKeyAnnotationOne := &ServiceKeyAnnotation{}
	serviceKeyAnnotationTwo := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, serviceKeyAnnotationOne, serviceKeyAnnotationDBTypes, false, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyAnnotationTwo, serviceKeyAnnotationDBTypes, false, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceKeyAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceKeyAnnotationOne := &ServiceKeyAnnotation{}
	serviceKeyAnnotationTwo := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, serviceKeyAnnotationOne, serviceKeyAnnotationDBTypes, false, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyAnnotationTwo, serviceKeyAnnotationDBTypes, false, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceKeyAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func serviceKeyAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyAnnotation) error {
	*o = ServiceKeyAnnotation{}
	return nil
}

func testServiceKeyAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceKeyAnnotation{}
	o := &ServiceKeyAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation object: %s", err)
	}

	AddServiceKeyAnnotationHook(boil.BeforeInsertHook, serviceKeyAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationBeforeInsertHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.AfterInsertHook, serviceKeyAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationAfterInsertHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.AfterSelectHook, serviceKeyAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationAfterSelectHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.BeforeUpdateHook, serviceKeyAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationBeforeUpdateHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.AfterUpdateHook, serviceKeyAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationAfterUpdateHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.BeforeDeleteHook, serviceKeyAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationBeforeDeleteHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.AfterDeleteHook, serviceKeyAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationAfterDeleteHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.BeforeUpsertHook, serviceKeyAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationBeforeUpsertHooks = []ServiceKeyAnnotationHook{}

	AddServiceKeyAnnotationHook(boil.AfterUpsertHook, serviceKeyAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyAnnotationAfterUpsertHooks = []ServiceKeyAnnotationHook{}
}

func testServiceKeyAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceKeyAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyAnnotationToOneServiceKeyUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceKeyAnnotation
	var foreign ServiceKey

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceKeyDBTypes, false, serviceKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKey struct: %s", err)
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

	slice := ServiceKeyAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceKeyAnnotation)(&slice), nil); err != nil {
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

func testServiceKeyAnnotationToOneSetOpServiceKeyUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyAnnotation
	var b, c ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyAnnotationDBTypes, false, strmangle.SetComplement(serviceKeyAnnotationPrimaryKeyColumns, serviceKeyAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceKey{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceKeyAnnotations[0] != &a {
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

func testServiceKeyAnnotationToOneRemoveOpServiceKeyUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyAnnotation
	var b ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyAnnotationDBTypes, false, strmangle.SetComplement(serviceKeyAnnotationPrimaryKeyColumns, serviceKeyAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceKeyAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceKeyAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
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

func testServiceKeyAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceKeyAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceKeyAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                           = bytes.MinRead
)

func testServiceKeyAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceKeyAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceKeyAnnotationAllColumns) == len(serviceKeyAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceKeyAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceKeyAnnotationAllColumns) == len(serviceKeyAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyAnnotationDBTypes, true, serviceKeyAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceKeyAnnotationAllColumns, serviceKeyAnnotationPrimaryKeyColumns) {
		fields = serviceKeyAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceKeyAnnotationAllColumns,
			serviceKeyAnnotationPrimaryKeyColumns,
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

	slice := ServiceKeyAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceKeyAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceKeyAnnotationAllColumns) == len(serviceKeyAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceKeyAnnotation{}
	if err = randomize.Struct(seed, &o, serviceKeyAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyAnnotation: %s", err)
	}

	count, err := ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceKeyAnnotationDBTypes, false, serviceKeyAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyAnnotation: %s", err)
	}

	count, err = ServiceKeyAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
