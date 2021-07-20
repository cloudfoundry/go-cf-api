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

func testRevisionSidecars(t *testing.T) {
	t.Parallel()

	query := RevisionSidecars()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevisionSidecarsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
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

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevisionSidecars().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionSidecarSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionSidecarsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevisionSidecarExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevisionSidecar exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevisionSidecarExists to return true, but got false.")
	}
}

func testRevisionSidecarsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revisionSidecarFound, err := FindRevisionSidecar(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revisionSidecarFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevisionSidecarsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevisionSidecars().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevisionSidecarsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevisionSidecars().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevisionSidecarsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revisionSidecarOne := &RevisionSidecar{}
	revisionSidecarTwo := &RevisionSidecar{}
	if err = randomize.Struct(seed, revisionSidecarOne, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionSidecarTwo, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionSidecarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionSidecarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionSidecars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevisionSidecarsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revisionSidecarOne := &RevisionSidecar{}
	revisionSidecarTwo := &RevisionSidecar{}
	if err = randomize.Struct(seed, revisionSidecarOne, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionSidecarTwo, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionSidecarOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionSidecarTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revisionSidecarBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func revisionSidecarAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionSidecar) error {
	*o = RevisionSidecar{}
	return nil
}

func testRevisionSidecarsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevisionSidecar{}
	o := &RevisionSidecar{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar object: %s", err)
	}

	AddRevisionSidecarHook(boil.BeforeInsertHook, revisionSidecarBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarBeforeInsertHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.AfterInsertHook, revisionSidecarAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarAfterInsertHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.AfterSelectHook, revisionSidecarAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revisionSidecarAfterSelectHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.BeforeUpdateHook, revisionSidecarBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revisionSidecarBeforeUpdateHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.AfterUpdateHook, revisionSidecarAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revisionSidecarAfterUpdateHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.BeforeDeleteHook, revisionSidecarBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revisionSidecarBeforeDeleteHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.AfterDeleteHook, revisionSidecarAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revisionSidecarAfterDeleteHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.BeforeUpsertHook, revisionSidecarBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarBeforeUpsertHooks = []RevisionSidecarHook{}

	AddRevisionSidecarHook(boil.AfterUpsertHook, revisionSidecarAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revisionSidecarAfterUpsertHooks = []RevisionSidecarHook{}
}

func testRevisionSidecarsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionSidecarsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revisionSidecarColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionSidecarToManyRevisionSidecarProcessTypes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionSidecar
	var b, c RevisionSidecarProcessType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, revisionSidecarProcessTypeDBTypes, false, revisionSidecarProcessTypeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RevisionSidecarGUID = a.GUID
	c.RevisionSidecarGUID = a.GUID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.RevisionSidecarProcessTypes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.RevisionSidecarGUID == b.RevisionSidecarGUID {
			bFound = true
		}
		if v.RevisionSidecarGUID == c.RevisionSidecarGUID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RevisionSidecarSlice{&a}
	if err = a.L.LoadRevisionSidecarProcessTypes(ctx, tx, false, (*[]*RevisionSidecar)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RevisionSidecarProcessTypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RevisionSidecarProcessTypes = nil
	if err = a.L.LoadRevisionSidecarProcessTypes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RevisionSidecarProcessTypes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testRevisionSidecarToManyAddOpRevisionSidecarProcessTypes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionSidecar
	var b, c, d, e RevisionSidecarProcessType

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionSidecarDBTypes, false, strmangle.SetComplement(revisionSidecarPrimaryKeyColumns, revisionSidecarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RevisionSidecarProcessType{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, revisionSidecarProcessTypeDBTypes, false, strmangle.SetComplement(revisionSidecarProcessTypePrimaryKeyColumns, revisionSidecarProcessTypeColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RevisionSidecarProcessType{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRevisionSidecarProcessTypes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.GUID != first.RevisionSidecarGUID {
			t.Error("foreign key was wrong value", a.GUID, first.RevisionSidecarGUID)
		}
		if a.GUID != second.RevisionSidecarGUID {
			t.Error("foreign key was wrong value", a.GUID, second.RevisionSidecarGUID)
		}

		if first.R.RevisionSidecar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.RevisionSidecar != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RevisionSidecarProcessTypes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RevisionSidecarProcessTypes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RevisionSidecarProcessTypes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testRevisionSidecarToOneRevisionUsingRevision(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevisionSidecar
	var foreign Revision

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revisionSidecarDBTypes, false, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, revisionDBTypes, false, revisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Revision struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RevisionGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Revision().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := RevisionSidecarSlice{&local}
	if err = local.L.LoadRevision(ctx, tx, false, (*[]*RevisionSidecar)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Revision = nil
	if err = local.L.LoadRevision(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Revision == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRevisionSidecarToOneSetOpRevisionUsingRevision(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionSidecar
	var b, c Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionSidecarDBTypes, false, strmangle.SetComplement(revisionSidecarPrimaryKeyColumns, revisionSidecarColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Revision{&b, &c} {
		err = a.SetRevision(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Revision != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RevisionSidecars[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RevisionGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RevisionGUID))
		reflect.Indirect(reflect.ValueOf(&a.RevisionGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RevisionGUID != x.GUID {
			t.Error("foreign key was wrong value", a.RevisionGUID, x.GUID)
		}
	}
}

func testRevisionSidecarsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
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

func testRevisionSidecarsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionSidecarSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevisionSidecarsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionSidecars().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revisionSidecarDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `Name`: `character varying`, `Command`: `character varying`, `RevisionGUID`: `character varying`, `Memory`: `integer`}
	_                      = bytes.MinRead
)

func testRevisionSidecarsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revisionSidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revisionSidecarAllColumns) == len(revisionSidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevisionSidecarsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revisionSidecarAllColumns) == len(revisionSidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionSidecar{}
	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionSidecarDBTypes, true, revisionSidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revisionSidecarAllColumns, revisionSidecarPrimaryKeyColumns) {
		fields = revisionSidecarAllColumns
	} else {
		fields = strmangle.SetComplement(
			revisionSidecarAllColumns,
			revisionSidecarPrimaryKeyColumns,
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

	slice := RevisionSidecarSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevisionSidecarsUpsert(t *testing.T) {
	t.Parallel()

	if len(revisionSidecarAllColumns) == len(revisionSidecarPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevisionSidecar{}
	if err = randomize.Struct(seed, &o, revisionSidecarDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionSidecar: %s", err)
	}

	count, err := RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revisionSidecarDBTypes, false, revisionSidecarPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionSidecar struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionSidecar: %s", err)
	}

	count, err = RevisionSidecars().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
