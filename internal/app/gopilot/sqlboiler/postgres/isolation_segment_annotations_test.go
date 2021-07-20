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

func testIsolationSegmentAnnotations(t *testing.T) {
	t.Parallel()

	query := IsolationSegmentAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIsolationSegmentAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
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

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := IsolationSegmentAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IsolationSegmentAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIsolationSegmentAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IsolationSegmentAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if IsolationSegmentAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IsolationSegmentAnnotationExists to return true, but got false.")
	}
}

func testIsolationSegmentAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	isolationSegmentAnnotationFound, err := FindIsolationSegmentAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if isolationSegmentAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIsolationSegmentAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = IsolationSegmentAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIsolationSegmentAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := IsolationSegmentAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIsolationSegmentAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	isolationSegmentAnnotationOne := &IsolationSegmentAnnotation{}
	isolationSegmentAnnotationTwo := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, isolationSegmentAnnotationOne, isolationSegmentAnnotationDBTypes, false, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, isolationSegmentAnnotationTwo, isolationSegmentAnnotationDBTypes, false, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = isolationSegmentAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = isolationSegmentAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IsolationSegmentAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIsolationSegmentAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	isolationSegmentAnnotationOne := &IsolationSegmentAnnotation{}
	isolationSegmentAnnotationTwo := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, isolationSegmentAnnotationOne, isolationSegmentAnnotationDBTypes, false, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, isolationSegmentAnnotationTwo, isolationSegmentAnnotationDBTypes, false, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = isolationSegmentAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = isolationSegmentAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func isolationSegmentAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func isolationSegmentAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IsolationSegmentAnnotation) error {
	*o = IsolationSegmentAnnotation{}
	return nil
}

func testIsolationSegmentAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &IsolationSegmentAnnotation{}
	o := &IsolationSegmentAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation object: %s", err)
	}

	AddIsolationSegmentAnnotationHook(boil.BeforeInsertHook, isolationSegmentAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationBeforeInsertHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.AfterInsertHook, isolationSegmentAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationAfterInsertHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.AfterSelectHook, isolationSegmentAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationAfterSelectHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.BeforeUpdateHook, isolationSegmentAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationBeforeUpdateHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.AfterUpdateHook, isolationSegmentAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationAfterUpdateHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.BeforeDeleteHook, isolationSegmentAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationBeforeDeleteHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.AfterDeleteHook, isolationSegmentAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationAfterDeleteHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.BeforeUpsertHook, isolationSegmentAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationBeforeUpsertHooks = []IsolationSegmentAnnotationHook{}

	AddIsolationSegmentAnnotationHook(boil.AfterUpsertHook, isolationSegmentAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	isolationSegmentAnnotationAfterUpsertHooks = []IsolationSegmentAnnotationHook{}
}

func testIsolationSegmentAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIsolationSegmentAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(isolationSegmentAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIsolationSegmentAnnotationToOneIsolationSegmentUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local IsolationSegmentAnnotation
	var foreign IsolationSegment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, isolationSegmentDBTypes, false, isolationSegmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegment struct: %s", err)
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

	slice := IsolationSegmentAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*IsolationSegmentAnnotation)(&slice), nil); err != nil {
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

func testIsolationSegmentAnnotationToOneSetOpIsolationSegmentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IsolationSegmentAnnotation
	var b, c IsolationSegment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, isolationSegmentAnnotationDBTypes, false, strmangle.SetComplement(isolationSegmentAnnotationPrimaryKeyColumns, isolationSegmentAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*IsolationSegment{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceIsolationSegmentAnnotations[0] != &a {
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

func testIsolationSegmentAnnotationToOneRemoveOpIsolationSegmentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IsolationSegmentAnnotation
	var b IsolationSegment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, isolationSegmentAnnotationDBTypes, false, strmangle.SetComplement(isolationSegmentAnnotationPrimaryKeyColumns, isolationSegmentAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, isolationSegmentDBTypes, false, strmangle.SetComplement(isolationSegmentPrimaryKeyColumns, isolationSegmentColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceIsolationSegmentAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testIsolationSegmentAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
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

func testIsolationSegmentAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IsolationSegmentAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIsolationSegmentAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IsolationSegmentAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	isolationSegmentAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                                 = bytes.MinRead
)

func testIsolationSegmentAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(isolationSegmentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(isolationSegmentAnnotationAllColumns) == len(isolationSegmentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIsolationSegmentAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(isolationSegmentAnnotationAllColumns) == len(isolationSegmentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, isolationSegmentAnnotationDBTypes, true, isolationSegmentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(isolationSegmentAnnotationAllColumns, isolationSegmentAnnotationPrimaryKeyColumns) {
		fields = isolationSegmentAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			isolationSegmentAnnotationAllColumns,
			isolationSegmentAnnotationPrimaryKeyColumns,
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

	slice := IsolationSegmentAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIsolationSegmentAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(isolationSegmentAnnotationAllColumns) == len(isolationSegmentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := IsolationSegmentAnnotation{}
	if err = randomize.Struct(seed, &o, isolationSegmentAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IsolationSegmentAnnotation: %s", err)
	}

	count, err := IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, isolationSegmentAnnotationDBTypes, false, isolationSegmentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IsolationSegmentAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IsolationSegmentAnnotation: %s", err)
	}

	count, err = IsolationSegmentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
