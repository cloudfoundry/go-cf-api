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

func testRouteBindingOperations(t *testing.T) {
	t.Parallel()

	query := RouteBindingOperations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRouteBindingOperationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
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

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingOperationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RouteBindingOperations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingOperationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingOperationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingOperationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RouteBindingOperationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RouteBindingOperation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RouteBindingOperationExists to return true, but got false.")
	}
}

func testRouteBindingOperationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routeBindingOperationFound, err := FindRouteBindingOperation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if routeBindingOperationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRouteBindingOperationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RouteBindingOperations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRouteBindingOperationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RouteBindingOperations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRouteBindingOperationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routeBindingOperationOne := &RouteBindingOperation{}
	routeBindingOperationTwo := &RouteBindingOperation{}
	if err = randomize.Struct(seed, routeBindingOperationOne, routeBindingOperationDBTypes, false, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingOperationTwo, routeBindingOperationDBTypes, false, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRouteBindingOperationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routeBindingOperationOne := &RouteBindingOperation{}
	routeBindingOperationTwo := &RouteBindingOperation{}
	if err = randomize.Struct(seed, routeBindingOperationOne, routeBindingOperationDBTypes, false, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingOperationTwo, routeBindingOperationDBTypes, false, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routeBindingOperationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func routeBindingOperationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingOperation) error {
	*o = RouteBindingOperation{}
	return nil
}

func testRouteBindingOperationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RouteBindingOperation{}
	o := &RouteBindingOperation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation object: %s", err)
	}

	AddRouteBindingOperationHook(boil.BeforeInsertHook, routeBindingOperationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationBeforeInsertHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.AfterInsertHook, routeBindingOperationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationAfterInsertHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.AfterSelectHook, routeBindingOperationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationAfterSelectHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.BeforeUpdateHook, routeBindingOperationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationBeforeUpdateHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.AfterUpdateHook, routeBindingOperationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationAfterUpdateHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.BeforeDeleteHook, routeBindingOperationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationBeforeDeleteHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.AfterDeleteHook, routeBindingOperationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationAfterDeleteHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.BeforeUpsertHook, routeBindingOperationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationBeforeUpsertHooks = []RouteBindingOperationHook{}

	AddRouteBindingOperationHook(boil.AfterUpsertHook, routeBindingOperationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingOperationAfterUpsertHooks = []RouteBindingOperationHook{}
}

func testRouteBindingOperationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingOperationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routeBindingOperationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingOperationToOneRouteBindingUsingRouteBinding(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RouteBindingOperation
	var foreign RouteBinding

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, routeBindingDBTypes, false, routeBindingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBinding struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.RouteBindingID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RouteBinding().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RouteBindingOperationSlice{&local}
	if err = local.L.LoadRouteBinding(ctx, tx, false, (*[]*RouteBindingOperation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RouteBinding == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RouteBinding = nil
	if err = local.L.LoadRouteBinding(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RouteBinding == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRouteBindingOperationToOneSetOpRouteBindingUsingRouteBinding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingOperation
	var b, c RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingOperationDBTypes, false, strmangle.SetComplement(routeBindingOperationPrimaryKeyColumns, routeBindingOperationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetRouteBinding(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RouteBinding != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RouteBindingOperation != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.RouteBindingID, x.ID) {
			t.Error("foreign key was wrong value", a.RouteBindingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RouteBindingID))
		reflect.Indirect(reflect.ValueOf(&a.RouteBindingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.RouteBindingID, x.ID) {
			t.Error("foreign key was wrong value", a.RouteBindingID, x.ID)
		}
	}
}

func testRouteBindingOperationToOneRemoveOpRouteBindingUsingRouteBinding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingOperation
	var b RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingOperationDBTypes, false, strmangle.SetComplement(routeBindingOperationPrimaryKeyColumns, routeBindingOperationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, routeBindingDBTypes, false, strmangle.SetComplement(routeBindingPrimaryKeyColumns, routeBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetRouteBinding(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveRouteBinding(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.RouteBinding().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.RouteBinding != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.RouteBindingID) {
		t.Error("foreign key value should be nil")
	}

	if b.R.RouteBindingOperation != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testRouteBindingOperationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
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

func testRouteBindingOperationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingOperationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRouteBindingOperationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routeBindingOperationDBTypes = map[string]string{`ID`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `RouteBindingID`: `int`, `State`: `varchar`, `Type`: `varchar`, `Description`: `varchar`, `BrokerProvidedOperation`: `varchar`}
	_                            = bytes.MinRead
)

func testRouteBindingOperationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routeBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routeBindingOperationAllColumns) == len(routeBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRouteBindingOperationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routeBindingOperationAllColumns) == len(routeBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingOperation{}
	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingOperationDBTypes, true, routeBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routeBindingOperationAllColumns, routeBindingOperationPrimaryKeyColumns) {
		fields = routeBindingOperationAllColumns
	} else {
		fields = strmangle.SetComplement(
			routeBindingOperationAllColumns,
			routeBindingOperationPrimaryKeyColumns,
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

	slice := RouteBindingOperationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRouteBindingOperationsUpsert(t *testing.T) {
	t.Parallel()

	if len(routeBindingOperationAllColumns) == len(routeBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRouteBindingOperationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RouteBindingOperation{}
	if err = randomize.Struct(seed, &o, routeBindingOperationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingOperation: %s", err)
	}

	count, err := RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routeBindingOperationDBTypes, false, routeBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingOperation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingOperation: %s", err)
	}

	count, err = RouteBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
