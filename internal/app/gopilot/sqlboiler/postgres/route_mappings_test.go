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

func testRouteMappings(t *testing.T) {
	t.Parallel()

	query := RouteMappings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRouteMappingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
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

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteMappingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RouteMappings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteMappingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteMappingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRouteMappingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RouteMappingExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RouteMapping exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RouteMappingExists to return true, but got false.")
	}
}

func testRouteMappingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	routeMappingFound, err := FindRouteMapping(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if routeMappingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRouteMappingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RouteMappings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRouteMappingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RouteMappings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRouteMappingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	routeMappingOne := &RouteMapping{}
	routeMappingTwo := &RouteMapping{}
	if err = randomize.Struct(seed, routeMappingOne, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}
	if err = randomize.Struct(seed, routeMappingTwo, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeMappingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeMappingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteMappings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRouteMappingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	routeMappingOne := &RouteMapping{}
	routeMappingTwo := &RouteMapping{}
	if err = randomize.Struct(seed, routeMappingOne, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}
	if err = randomize.Struct(seed, routeMappingTwo, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = routeMappingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = routeMappingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func routeMappingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func routeMappingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RouteMapping) error {
	*o = RouteMapping{}
	return nil
}

func testRouteMappingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RouteMapping{}
	o := &RouteMapping{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, routeMappingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RouteMapping object: %s", err)
	}

	AddRouteMappingHook(boil.BeforeInsertHook, routeMappingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	routeMappingBeforeInsertHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.AfterInsertHook, routeMappingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	routeMappingAfterInsertHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.AfterSelectHook, routeMappingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	routeMappingAfterSelectHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.BeforeUpdateHook, routeMappingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	routeMappingBeforeUpdateHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.AfterUpdateHook, routeMappingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	routeMappingAfterUpdateHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.BeforeDeleteHook, routeMappingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	routeMappingBeforeDeleteHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.AfterDeleteHook, routeMappingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	routeMappingAfterDeleteHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.BeforeUpsertHook, routeMappingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	routeMappingBeforeUpsertHooks = []RouteMappingHook{}

	AddRouteMappingHook(boil.AfterUpsertHook, routeMappingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	routeMappingAfterUpsertHooks = []RouteMappingHook{}
}

func testRouteMappingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteMappingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(routeMappingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRouteMappingToOneAppUsingApp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RouteMapping
	var foreign App

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AppGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.App().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := RouteMappingSlice{&local}
	if err = local.L.LoadApp(ctx, tx, false, (*[]*RouteMapping)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.App = nil
	if err = local.L.LoadApp(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRouteMappingToOneRouteUsingRoute(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RouteMapping
	var foreign Route

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, routeMappingDBTypes, false, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, routeDBTypes, false, routeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Route struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RouteGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Route().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := RouteMappingSlice{&local}
	if err = local.L.LoadRoute(ctx, tx, false, (*[]*RouteMapping)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Route == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Route = nil
	if err = local.L.LoadRoute(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Route == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRouteMappingToOneSetOpAppUsingApp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteMapping
	var b, c App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeMappingDBTypes, false, strmangle.SetComplement(routeMappingPrimaryKeyColumns, routeMappingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*App{&b, &c} {
		err = a.SetApp(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.App != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RouteMappings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AppGUID != x.GUID {
			t.Error("foreign key was wrong value", a.AppGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AppGUID))
		reflect.Indirect(reflect.ValueOf(&a.AppGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AppGUID != x.GUID {
			t.Error("foreign key was wrong value", a.AppGUID, x.GUID)
		}
	}
}
func testRouteMappingToOneSetOpRouteUsingRoute(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RouteMapping
	var b, c Route

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, routeMappingDBTypes, false, strmangle.SetComplement(routeMappingPrimaryKeyColumns, routeMappingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, routeDBTypes, false, strmangle.SetComplement(routePrimaryKeyColumns, routeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, routeDBTypes, false, strmangle.SetComplement(routePrimaryKeyColumns, routeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Route{&b, &c} {
		err = a.SetRoute(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Route != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RouteMappings[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RouteGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RouteGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RouteGUID))
		reflect.Indirect(reflect.ValueOf(&a.RouteGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RouteGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RouteGUID, x.GUID)
		}
	}
}

func testRouteMappingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
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

func testRouteMappingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RouteMappingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRouteMappingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RouteMappings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	routeMappingDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `AppPort`: `integer`, `GUID`: `text`, `AppGUID`: `text`, `RouteGUID`: `text`, `ProcessType`: `text`, `Weight`: `integer`}
	_                   = bytes.MinRead
)

func testRouteMappingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(routeMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(routeMappingAllColumns) == len(routeMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRouteMappingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(routeMappingAllColumns) == len(routeMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RouteMapping{}
	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, routeMappingDBTypes, true, routeMappingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(routeMappingAllColumns, routeMappingPrimaryKeyColumns) {
		fields = routeMappingAllColumns
	} else {
		fields = strmangle.SetComplement(
			routeMappingAllColumns,
			routeMappingPrimaryKeyColumns,
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

	slice := RouteMappingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRouteMappingsUpsert(t *testing.T) {
	t.Parallel()

	if len(routeMappingAllColumns) == len(routeMappingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RouteMapping{}
	if err = randomize.Struct(seed, &o, routeMappingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteMapping: %s", err)
	}

	count, err := RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, routeMappingDBTypes, false, routeMappingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RouteMapping struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RouteMapping: %s", err)
	}

	count, err = RouteMappings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
