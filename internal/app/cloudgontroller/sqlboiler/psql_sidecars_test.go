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

func testSidecars(t *testing.T) {
	t.Parallel()

	query := Sidecars()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSidecarsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
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

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSidecarsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Sidecars().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSidecarsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SidecarSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSidecarsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SidecarExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Sidecar exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SidecarExists to return true, but got false.")
	}
}

func testSidecarsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	sidecarFound, err := FindSidecar(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if sidecarFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSidecarsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Sidecars().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSidecarsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Sidecars().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSidecarsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	sidecarOne := &Sidecar{}
	sidecarTwo := &Sidecar{}
	if err = randomize.Struct(seed, sidecarOne, sidecarDBTypes, false, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}
	if err = randomize.Struct(seed, sidecarTwo, sidecarDBTypes, false, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sidecarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sidecarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sidecars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSidecarsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	sidecarOne := &Sidecar{}
	sidecarTwo := &Sidecar{}
	if err = randomize.Struct(seed, sidecarOne, sidecarDBTypes, false, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}
	if err = randomize.Struct(seed, sidecarTwo, sidecarDBTypes, false, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = sidecarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = sidecarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func sidecarBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func sidecarAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Sidecar) error {
	*o = Sidecar{}
	return nil
}

func testSidecarsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Sidecar{}
	o := &Sidecar{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, sidecarDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Sidecar object: %s", err)
	}

	AddSidecarHook(boil.BeforeInsertHook, sidecarBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	sidecarBeforeInsertHooks = []SidecarHook{}

	AddSidecarHook(boil.AfterInsertHook, sidecarAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	sidecarAfterInsertHooks = []SidecarHook{}

	AddSidecarHook(boil.AfterSelectHook, sidecarAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	sidecarAfterSelectHooks = []SidecarHook{}

	AddSidecarHook(boil.BeforeUpdateHook, sidecarBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	sidecarBeforeUpdateHooks = []SidecarHook{}

	AddSidecarHook(boil.AfterUpdateHook, sidecarAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	sidecarAfterUpdateHooks = []SidecarHook{}

	AddSidecarHook(boil.BeforeDeleteHook, sidecarBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	sidecarBeforeDeleteHooks = []SidecarHook{}

	AddSidecarHook(boil.AfterDeleteHook, sidecarAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	sidecarAfterDeleteHooks = []SidecarHook{}

	AddSidecarHook(boil.BeforeUpsertHook, sidecarBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	sidecarBeforeUpsertHooks = []SidecarHook{}

	AddSidecarHook(boil.AfterUpsertHook, sidecarAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	sidecarAfterUpsertHooks = []SidecarHook{}
}

func testSidecarsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSidecarsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(sidecarColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSidecarToManySidecarProcessTypes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sidecar
	var b, c SidecarProcessType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, sidecarProcessTypeDBTypes, false, sidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, sidecarProcessTypeDBTypes, false, sidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.SidecarGUID = a.GUID
	c.SidecarGUID = a.GUID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SidecarProcessTypes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.SidecarGUID == b.SidecarGUID {
			bFound = true
		}
		if v.SidecarGUID == c.SidecarGUID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SidecarSlice{&a}
	if err = a.L.LoadSidecarProcessTypes(ctx, tx, false, (*[]*Sidecar)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SidecarProcessTypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SidecarProcessTypes = nil
	if err = a.L.LoadSidecarProcessTypes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SidecarProcessTypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSidecarToManyAddOpSidecarProcessTypes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sidecar
	var b, c, d, e SidecarProcessType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sidecarDBTypes, false, strmangle.SetComplement(sidecarPrimaryKeyColumns, sidecarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*SidecarProcessType{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, sidecarProcessTypeDBTypes, false, strmangle.SetComplement(sidecarProcessTypePrimaryKeyColumns, sidecarProcessTypeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*SidecarProcessType{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSidecarProcessTypes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.GUID != first.SidecarGUID {
			t.Error("foreign key was wrong value", a.GUID, first.SidecarGUID)
		}
		if a.GUID != second.SidecarGUID {
			t.Error("foreign key was wrong value", a.GUID, second.SidecarGUID)
		}

		if first.R.Sidecar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Sidecar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SidecarProcessTypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SidecarProcessTypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SidecarProcessTypes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testSidecarToOneAppUsingApp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Sidecar
	var foreign App

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, sidecarDBTypes, false, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
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

	slice := SidecarSlice{&local}
	if err = local.L.LoadApp(ctx, tx, false, (*[]*Sidecar)(&slice), nil); err != nil {
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

func testSidecarToOneSetOpAppUsingApp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Sidecar
	var b, c App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, sidecarDBTypes, false, strmangle.SetComplement(sidecarPrimaryKeyColumns, sidecarColumnsWithoutDefault)...); err != nil {
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

		if x.R.Sidecars[0] != &a {
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

func testSidecarsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
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

func testSidecarsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SidecarSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSidecarsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Sidecars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	sidecarDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `Name`: `character varying`, `Command`: `character varying`, `AppGUID`: `character varying`, `Memory`: `integer`, `Origin`: `character varying`}
	_              = bytes.MinRead
)

func testSidecarsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(sidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(sidecarAllColumns) == len(sidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSidecarsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(sidecarAllColumns) == len(sidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Sidecar{}
	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, sidecarDBTypes, true, sidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(sidecarAllColumns, sidecarPrimaryKeyColumns) {
		fields = sidecarAllColumns
	} else {
		fields = strmangle.SetComplement(
			sidecarAllColumns,
			sidecarPrimaryKeyColumns,
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

	slice := SidecarSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSidecarsUpsert(t *testing.T) {
	t.Parallel()

	if len(sidecarAllColumns) == len(sidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Sidecar{}
	if err = randomize.Struct(seed, &o, sidecarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sidecar: %s", err)
	}

	count, err := Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, sidecarDBTypes, false, sidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Sidecar struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Sidecar: %s", err)
	}

	count, err = Sidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
