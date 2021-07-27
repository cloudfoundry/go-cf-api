// +build integration,postgres
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

func testProcessAnnotations(t *testing.T) {
	t.Parallel()

	query := ProcessAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProcessAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
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

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProcessAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProcessAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProcessAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ProcessAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProcessAnnotationExists to return true, but got false.")
	}
}

func testProcessAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	processAnnotationFound, err := FindProcessAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if processAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProcessAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProcessAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProcessAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProcessAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProcessAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	processAnnotationOne := &ProcessAnnotation{}
	processAnnotationTwo := &ProcessAnnotation{}
	if err = randomize.Struct(seed, processAnnotationOne, processAnnotationDBTypes, false, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, processAnnotationTwo, processAnnotationDBTypes, false, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = processAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProcessAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	processAnnotationOne := &ProcessAnnotation{}
	processAnnotationTwo := &ProcessAnnotation{}
	if err = randomize.Struct(seed, processAnnotationOne, processAnnotationDBTypes, false, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, processAnnotationTwo, processAnnotationDBTypes, false, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = processAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = processAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func processAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func processAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ProcessAnnotation) error {
	*o = ProcessAnnotation{}
	return nil
}

func testProcessAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ProcessAnnotation{}
	o := &ProcessAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation object: %s", err)
	}

	AddProcessAnnotationHook(boil.BeforeInsertHook, processAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	processAnnotationBeforeInsertHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.AfterInsertHook, processAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	processAnnotationAfterInsertHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.AfterSelectHook, processAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	processAnnotationAfterSelectHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.BeforeUpdateHook, processAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	processAnnotationBeforeUpdateHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.AfterUpdateHook, processAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	processAnnotationAfterUpdateHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.BeforeDeleteHook, processAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	processAnnotationBeforeDeleteHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.AfterDeleteHook, processAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	processAnnotationAfterDeleteHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.BeforeUpsertHook, processAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	processAnnotationBeforeUpsertHooks = []ProcessAnnotationHook{}

	AddProcessAnnotationHook(boil.AfterUpsertHook, processAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	processAnnotationAfterUpsertHooks = []ProcessAnnotationHook{}
}

func testProcessAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(processAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProcessAnnotationToOneProcessUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ProcessAnnotation
	var foreign Process

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, processDBTypes, false, processColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Process struct: %s", err)
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

	slice := ProcessAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ProcessAnnotation)(&slice), nil); err != nil {
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

func testProcessAnnotationToOneSetOpProcessUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProcessAnnotation
	var b, c Process

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, processAnnotationDBTypes, false, strmangle.SetComplement(processAnnotationPrimaryKeyColumns, processAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Process{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceProcessAnnotations[0] != &a {
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

func testProcessAnnotationToOneRemoveOpProcessUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ProcessAnnotation
	var b Process

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, processAnnotationDBTypes, false, strmangle.SetComplement(processAnnotationPrimaryKeyColumns, processAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, processDBTypes, false, strmangle.SetComplement(processPrimaryKeyColumns, processColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceProcessAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProcessAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
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

func testProcessAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProcessAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProcessAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProcessAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	processAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                        = bytes.MinRead
)

func testProcessAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(processAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(processAnnotationAllColumns) == len(processAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProcessAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(processAnnotationAllColumns) == len(processAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProcessAnnotation{}
	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, processAnnotationDBTypes, true, processAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(processAnnotationAllColumns, processAnnotationPrimaryKeyColumns) {
		fields = processAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			processAnnotationAllColumns,
			processAnnotationPrimaryKeyColumns,
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

	slice := ProcessAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProcessAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(processAnnotationAllColumns) == len(processAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ProcessAnnotation{}
	if err = randomize.Struct(seed, &o, processAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProcessAnnotation: %s", err)
	}

	count, err := ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, processAnnotationDBTypes, false, processAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProcessAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProcessAnnotation: %s", err)
	}

	count, err = ProcessAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
