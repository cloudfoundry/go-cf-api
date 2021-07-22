// +build integration
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

func testRouteBindingAnnotations(t *testing.T) {
	t.Parallel()

	query := RouteBindingAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRouteBindingAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
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

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RouteBindingAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RouteBindingAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RouteBindingAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RouteBindingAnnotationExists to return true, but got false.")
	}
}

func testRouteBindingAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routeBindingAnnotationFound, err := FindRouteBindingAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if routeBindingAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRouteBindingAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RouteBindingAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRouteBindingAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RouteBindingAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRouteBindingAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routeBindingAnnotationOne := &RouteBindingAnnotation{}
	routeBindingAnnotationTwo := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, routeBindingAnnotationOne, routeBindingAnnotationDBTypes, false, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingAnnotationTwo, routeBindingAnnotationDBTypes, false, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRouteBindingAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routeBindingAnnotationOne := &RouteBindingAnnotation{}
	routeBindingAnnotationTwo := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, routeBindingAnnotationOne, routeBindingAnnotationDBTypes, false, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingAnnotationTwo, routeBindingAnnotationDBTypes, false, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routeBindingAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func routeBindingAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingAnnotation) error {
	*o = RouteBindingAnnotation{}
	return nil
}

func testRouteBindingAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RouteBindingAnnotation{}
	o := &RouteBindingAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation object: %s", err)
	}

	AddRouteBindingAnnotationHook(boil.BeforeInsertHook, routeBindingAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationBeforeInsertHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.AfterInsertHook, routeBindingAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationAfterInsertHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.AfterSelectHook, routeBindingAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationAfterSelectHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.BeforeUpdateHook, routeBindingAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationBeforeUpdateHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.AfterUpdateHook, routeBindingAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationAfterUpdateHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.BeforeDeleteHook, routeBindingAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationBeforeDeleteHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.AfterDeleteHook, routeBindingAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationAfterDeleteHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.BeforeUpsertHook, routeBindingAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationBeforeUpsertHooks = []RouteBindingAnnotationHook{}

	AddRouteBindingAnnotationHook(boil.AfterUpsertHook, routeBindingAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingAnnotationAfterUpsertHooks = []RouteBindingAnnotationHook{}
}

func testRouteBindingAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routeBindingAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingAnnotationToOneRouteBindingUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RouteBindingAnnotation
	var foreign RouteBinding

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, routeBindingDBTypes, false, routeBindingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBinding struct: %s", err)
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

	slice := RouteBindingAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*RouteBindingAnnotation)(&slice), nil); err != nil {
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

func testRouteBindingAnnotationToOneSetOpRouteBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingAnnotation
	var b, c RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingAnnotationDBTypes, false, strmangle.SetComplement(routeBindingAnnotationPrimaryKeyColumns, routeBindingAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, routeBindingDBTypes, false, strmangle.SetComplement(routeBindingPrimaryKeyColumns, routeBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, routeBindingDBTypes, false, strmangle.SetComplement(routeBindingPrimaryKeyColumns, routeBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RouteBinding{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceRouteBindingAnnotations[0] != &a {
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

func testRouteBindingAnnotationToOneRemoveOpRouteBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingAnnotation
	var b RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingAnnotationDBTypes, false, strmangle.SetComplement(routeBindingAnnotationPrimaryKeyColumns, routeBindingAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, routeBindingDBTypes, false, strmangle.SetComplement(routeBindingPrimaryKeyColumns, routeBindingColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceRouteBindingAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRouteBindingAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
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

func testRouteBindingAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRouteBindingAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routeBindingAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                             = bytes.MinRead
)

func testRouteBindingAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routeBindingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routeBindingAnnotationAllColumns) == len(routeBindingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRouteBindingAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routeBindingAnnotationAllColumns) == len(routeBindingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingAnnotation{}
	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingAnnotationDBTypes, true, routeBindingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routeBindingAnnotationAllColumns, routeBindingAnnotationPrimaryKeyColumns) {
		fields = routeBindingAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			routeBindingAnnotationAllColumns,
			routeBindingAnnotationPrimaryKeyColumns,
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

	slice := RouteBindingAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRouteBindingAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(routeBindingAnnotationAllColumns) == len(routeBindingAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RouteBindingAnnotation{}
	if err = randomize.Struct(seed, &o, routeBindingAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingAnnotation: %s", err)
	}

	count, err := RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routeBindingAnnotationDBTypes, false, routeBindingAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingAnnotation: %s", err)
	}

	count, err = RouteBindingAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
