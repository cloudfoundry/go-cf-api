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

func testRouteBindingLabels(t *testing.T) {
	t.Parallel()

	query := RouteBindingLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRouteBindingLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
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

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RouteBindingLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteBindingLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RouteBindingLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RouteBindingLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RouteBindingLabelExists to return true, but got false.")
	}
}

func testRouteBindingLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routeBindingLabelFound, err := FindRouteBindingLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if routeBindingLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRouteBindingLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RouteBindingLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRouteBindingLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RouteBindingLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRouteBindingLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routeBindingLabelOne := &RouteBindingLabel{}
	routeBindingLabelTwo := &RouteBindingLabel{}
	if err = randomize.Struct(seed, routeBindingLabelOne, routeBindingLabelDBTypes, false, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingLabelTwo, routeBindingLabelDBTypes, false, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRouteBindingLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routeBindingLabelOne := &RouteBindingLabel{}
	routeBindingLabelTwo := &RouteBindingLabel{}
	if err = randomize.Struct(seed, routeBindingLabelOne, routeBindingLabelDBTypes, false, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, routeBindingLabelTwo, routeBindingLabelDBTypes, false, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeBindingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeBindingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routeBindingLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func routeBindingLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteBindingLabel) error {
	*o = RouteBindingLabel{}
	return nil
}

func testRouteBindingLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RouteBindingLabel{}
	o := &RouteBindingLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel object: %s", err)
	}

	AddRouteBindingLabelHook(boil.BeforeInsertHook, routeBindingLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelBeforeInsertHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.AfterInsertHook, routeBindingLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelAfterInsertHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.AfterSelectHook, routeBindingLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelAfterSelectHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.BeforeUpdateHook, routeBindingLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelBeforeUpdateHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.AfterUpdateHook, routeBindingLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelAfterUpdateHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.BeforeDeleteHook, routeBindingLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelBeforeDeleteHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.AfterDeleteHook, routeBindingLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelAfterDeleteHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.BeforeUpsertHook, routeBindingLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelBeforeUpsertHooks = []RouteBindingLabelHook{}

	AddRouteBindingLabelHook(boil.AfterUpsertHook, routeBindingLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routeBindingLabelAfterUpsertHooks = []RouteBindingLabelHook{}
}

func testRouteBindingLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routeBindingLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteBindingLabelToOneRouteBindingUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RouteBindingLabel
	var foreign RouteBinding

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
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

	slice := RouteBindingLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*RouteBindingLabel)(&slice), nil); err != nil {
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

func testRouteBindingLabelToOneSetOpRouteBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingLabel
	var b, c RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingLabelDBTypes, false, strmangle.SetComplement(routeBindingLabelPrimaryKeyColumns, routeBindingLabelColumnsWithoutDefault)...); err != nil {
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

		if x.R.ResourceRouteBindingLabels[0] != &a {
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

func testRouteBindingLabelToOneRemoveOpRouteBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteBindingLabel
	var b RouteBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeBindingLabelDBTypes, false, strmangle.SetComplement(routeBindingLabelPrimaryKeyColumns, routeBindingLabelColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceRouteBindingLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRouteBindingLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
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

func testRouteBindingLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteBindingLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRouteBindingLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteBindingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routeBindingLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                        = bytes.MinRead
)

func testRouteBindingLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routeBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routeBindingLabelAllColumns) == len(routeBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRouteBindingLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routeBindingLabelAllColumns) == len(routeBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteBindingLabel{}
	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeBindingLabelDBTypes, true, routeBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routeBindingLabelAllColumns, routeBindingLabelPrimaryKeyColumns) {
		fields = routeBindingLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			routeBindingLabelAllColumns,
			routeBindingLabelPrimaryKeyColumns,
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

	slice := RouteBindingLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRouteBindingLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(routeBindingLabelAllColumns) == len(routeBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RouteBindingLabel{}
	if err = randomize.Struct(seed, &o, routeBindingLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingLabel: %s", err)
	}

	count, err := RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routeBindingLabelDBTypes, false, routeBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteBindingLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteBindingLabel: %s", err)
	}

	count, err = RouteBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
