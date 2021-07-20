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

func testDomainLabels(t *testing.T) {
	t.Parallel()

	query := DomainLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDomainLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
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

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DomainLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DomainLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DomainLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DomainLabelExists to return true, but got false.")
	}
}

func testDomainLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	domainLabelFound, err := FindDomainLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if domainLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDomainLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DomainLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDomainLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DomainLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDomainLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	domainLabelOne := &DomainLabel{}
	domainLabelTwo := &DomainLabel{}
	if err = randomize.Struct(seed, domainLabelOne, domainLabelDBTypes, false, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, domainLabelTwo, domainLabelDBTypes, false, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDomainLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	domainLabelOne := &DomainLabel{}
	domainLabelTwo := &DomainLabel{}
	if err = randomize.Struct(seed, domainLabelOne, domainLabelDBTypes, false, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, domainLabelTwo, domainLabelDBTypes, false, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func domainLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func domainLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DomainLabel) error {
	*o = DomainLabel{}
	return nil
}

func testDomainLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DomainLabel{}
	o := &DomainLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, domainLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DomainLabel object: %s", err)
	}

	AddDomainLabelHook(boil.BeforeInsertHook, domainLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	domainLabelBeforeInsertHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.AfterInsertHook, domainLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	domainLabelAfterInsertHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.AfterSelectHook, domainLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	domainLabelAfterSelectHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.BeforeUpdateHook, domainLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	domainLabelBeforeUpdateHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.AfterUpdateHook, domainLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	domainLabelAfterUpdateHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.BeforeDeleteHook, domainLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	domainLabelBeforeDeleteHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.AfterDeleteHook, domainLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	domainLabelAfterDeleteHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.BeforeUpsertHook, domainLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	domainLabelBeforeUpsertHooks = []DomainLabelHook{}

	AddDomainLabelHook(boil.AfterUpsertHook, domainLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	domainLabelAfterUpsertHooks = []DomainLabelHook{}
}

func testDomainLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(domainLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainLabelToOneDomainUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DomainLabel
	var foreign Domain

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, domainDBTypes, false, domainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Domain struct: %s", err)
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

	slice := DomainLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*DomainLabel)(&slice), nil); err != nil {
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

func testDomainLabelToOneSetOpDomainUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainLabel
	var b, c Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainLabelDBTypes, false, strmangle.SetComplement(domainLabelPrimaryKeyColumns, domainLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Domain{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceDomainLabels[0] != &a {
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

func testDomainLabelToOneRemoveOpDomainUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainLabel
	var b Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainLabelDBTypes, false, strmangle.SetComplement(domainLabelPrimaryKeyColumns, domainLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceDomainLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDomainLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
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

func testDomainLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDomainLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	domainLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                  = bytes.MinRead
)

func testDomainLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(domainLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(domainLabelAllColumns) == len(domainLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDomainLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(domainLabelAllColumns) == len(domainLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainLabel{}
	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainLabelDBTypes, true, domainLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(domainLabelAllColumns, domainLabelPrimaryKeyColumns) {
		fields = domainLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			domainLabelAllColumns,
			domainLabelPrimaryKeyColumns,
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

	slice := DomainLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDomainLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(domainLabelAllColumns) == len(domainLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DomainLabel{}
	if err = randomize.Struct(seed, &o, domainLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainLabel: %s", err)
	}

	count, err := DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, domainLabelDBTypes, false, domainLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainLabel: %s", err)
	}

	count, err = DomainLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
