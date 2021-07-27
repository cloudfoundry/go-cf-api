// +build integration,mysql
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

func testAppAnnotations(t *testing.T) {
	t.Parallel()

	query := AppAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAppAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
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

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AppAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AppAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AppAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppAnnotationExists to return true, but got false.")
	}
}

func testAppAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	appAnnotationFound, err := FindAppAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if appAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAppAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AppAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAppAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AppAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appAnnotationOne := &AppAnnotation{}
	appAnnotationTwo := &AppAnnotation{}
	if err = randomize.Struct(seed, appAnnotationOne, appAnnotationDBTypes, false, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, appAnnotationTwo, appAnnotationDBTypes, false, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appAnnotationOne := &AppAnnotation{}
	appAnnotationTwo := &AppAnnotation{}
	if err = randomize.Struct(seed, appAnnotationOne, appAnnotationDBTypes, false, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, appAnnotationTwo, appAnnotationDBTypes, false, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func appAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func appAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AppAnnotation) error {
	*o = AppAnnotation{}
	return nil
}

func testAppAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AppAnnotation{}
	o := &AppAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AppAnnotation object: %s", err)
	}

	AddAppAnnotationHook(boil.BeforeInsertHook, appAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	appAnnotationBeforeInsertHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.AfterInsertHook, appAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	appAnnotationAfterInsertHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.AfterSelectHook, appAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	appAnnotationAfterSelectHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.BeforeUpdateHook, appAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	appAnnotationBeforeUpdateHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.AfterUpdateHook, appAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	appAnnotationAfterUpdateHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.BeforeDeleteHook, appAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	appAnnotationBeforeDeleteHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.AfterDeleteHook, appAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	appAnnotationAfterDeleteHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.BeforeUpsertHook, appAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	appAnnotationBeforeUpsertHooks = []AppAnnotationHook{}

	AddAppAnnotationHook(boil.AfterUpsertHook, appAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	appAnnotationAfterUpsertHooks = []AppAnnotationHook{}
}

func testAppAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(appAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppAnnotationToOneAppUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local AppAnnotation
	var foreign App

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
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

	slice := AppAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*AppAnnotation)(&slice), nil); err != nil {
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

func testAppAnnotationToOneSetOpAppUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AppAnnotation
	var b, c App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, appAnnotationDBTypes, false, strmangle.SetComplement(appAnnotationPrimaryKeyColumns, appAnnotationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceAppAnnotations[0] != &a {
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

func testAppAnnotationToOneRemoveOpAppUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AppAnnotation
	var b App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, appAnnotationDBTypes, false, strmangle.SetComplement(appAnnotationPrimaryKeyColumns, appAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, appDBTypes, false, strmangle.SetComplement(appPrimaryKeyColumns, appColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceAppAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testAppAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
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

func testAppAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAppAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                    = bytes.MinRead
)

func testAppAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(appAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(appAnnotationAllColumns) == len(appAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAppAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appAnnotationAllColumns) == len(appAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppAnnotation{}
	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appAnnotationDBTypes, true, appAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appAnnotationAllColumns, appAnnotationPrimaryKeyColumns) {
		fields = appAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			appAnnotationAllColumns,
			appAnnotationPrimaryKeyColumns,
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

	slice := AppAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAppAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(appAnnotationAllColumns) == len(appAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLAppAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AppAnnotation{}
	if err = randomize.Struct(seed, &o, appAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppAnnotation: %s", err)
	}

	count, err := AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, appAnnotationDBTypes, false, appAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppAnnotation: %s", err)
	}

	count, err = AppAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
