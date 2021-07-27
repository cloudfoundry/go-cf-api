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

func testRevisionAnnotations(t *testing.T) {
	t.Parallel()

	query := RevisionAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRevisionAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
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

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RevisionAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRevisionAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RevisionAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RevisionAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RevisionAnnotationExists to return true, but got false.")
	}
}

func testRevisionAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	revisionAnnotationFound, err := FindRevisionAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if revisionAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRevisionAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RevisionAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRevisionAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RevisionAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRevisionAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	revisionAnnotationOne := &RevisionAnnotation{}
	revisionAnnotationTwo := &RevisionAnnotation{}
	if err = randomize.Struct(seed, revisionAnnotationOne, revisionAnnotationDBTypes, false, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionAnnotationTwo, revisionAnnotationDBTypes, false, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRevisionAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	revisionAnnotationOne := &RevisionAnnotation{}
	revisionAnnotationTwo := &RevisionAnnotation{}
	if err = randomize.Struct(seed, revisionAnnotationOne, revisionAnnotationDBTypes, false, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, revisionAnnotationTwo, revisionAnnotationDBTypes, false, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = revisionAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = revisionAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func revisionAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func revisionAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RevisionAnnotation) error {
	*o = RevisionAnnotation{}
	return nil
}

func testRevisionAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RevisionAnnotation{}
	o := &RevisionAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation object: %s", err)
	}

	AddRevisionAnnotationHook(boil.BeforeInsertHook, revisionAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationBeforeInsertHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.AfterInsertHook, revisionAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationAfterInsertHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.AfterSelectHook, revisionAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationAfterSelectHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.BeforeUpdateHook, revisionAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationBeforeUpdateHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.AfterUpdateHook, revisionAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationAfterUpdateHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.BeforeDeleteHook, revisionAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationBeforeDeleteHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.AfterDeleteHook, revisionAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationAfterDeleteHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.BeforeUpsertHook, revisionAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationBeforeUpsertHooks = []RevisionAnnotationHook{}

	AddRevisionAnnotationHook(boil.AfterUpsertHook, revisionAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	revisionAnnotationAfterUpsertHooks = []RevisionAnnotationHook{}
}

func testRevisionAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(revisionAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRevisionAnnotationToOneRevisionUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RevisionAnnotation
	var foreign Revision

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, revisionDBTypes, false, revisionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Revision struct: %s", err)
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

	slice := RevisionAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*RevisionAnnotation)(&slice), nil); err != nil {
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

func testRevisionAnnotationToOneSetOpRevisionUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionAnnotation
	var b, c Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionAnnotationDBTypes, false, strmangle.SetComplement(revisionAnnotationPrimaryKeyColumns, revisionAnnotationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceRevisionAnnotations[0] != &a {
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

func testRevisionAnnotationToOneRemoveOpRevisionUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RevisionAnnotation
	var b Revision

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, revisionAnnotationDBTypes, false, strmangle.SetComplement(revisionAnnotationPrimaryKeyColumns, revisionAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, revisionDBTypes, false, strmangle.SetComplement(revisionPrimaryKeyColumns, revisionColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceRevisionAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testRevisionAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
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

func testRevisionAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RevisionAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRevisionAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RevisionAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	revisionAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                         = bytes.MinRead
)

func testRevisionAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(revisionAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(revisionAnnotationAllColumns) == len(revisionAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRevisionAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(revisionAnnotationAllColumns) == len(revisionAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RevisionAnnotation{}
	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, revisionAnnotationDBTypes, true, revisionAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(revisionAnnotationAllColumns, revisionAnnotationPrimaryKeyColumns) {
		fields = revisionAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			revisionAnnotationAllColumns,
			revisionAnnotationPrimaryKeyColumns,
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

	slice := RevisionAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRevisionAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(revisionAnnotationAllColumns) == len(revisionAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RevisionAnnotation{}
	if err = randomize.Struct(seed, &o, revisionAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionAnnotation: %s", err)
	}

	count, err := RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, revisionAnnotationDBTypes, false, revisionAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RevisionAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RevisionAnnotation: %s", err)
	}

	count, err = RevisionAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
