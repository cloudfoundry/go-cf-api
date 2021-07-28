// +build psql_integration
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

func testBuildAnnotations(t *testing.T) {
	t.Parallel()

	query := BuildAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
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

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BuildAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BuildAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildAnnotationExists to return true, but got false.")
	}
}

func testBuildAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildAnnotationFound, err := FindBuildAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BuildAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BuildAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildAnnotationOne := &BuildAnnotation{}
	buildAnnotationTwo := &BuildAnnotation{}
	if err = randomize.Struct(seed, buildAnnotationOne, buildAnnotationDBTypes, false, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, buildAnnotationTwo, buildAnnotationDBTypes, false, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildAnnotationOne := &BuildAnnotation{}
	buildAnnotationTwo := &BuildAnnotation{}
	if err = randomize.Struct(seed, buildAnnotationOne, buildAnnotationDBTypes, false, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, buildAnnotationTwo, buildAnnotationDBTypes, false, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func buildAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildAnnotation) error {
	*o = BuildAnnotation{}
	return nil
}

func testBuildAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BuildAnnotation{}
	o := &BuildAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation object: %s", err)
	}

	AddBuildAnnotationHook(boil.BeforeInsertHook, buildAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildAnnotationBeforeInsertHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.AfterInsertHook, buildAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildAnnotationAfterInsertHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.AfterSelectHook, buildAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildAnnotationAfterSelectHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.BeforeUpdateHook, buildAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildAnnotationBeforeUpdateHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.AfterUpdateHook, buildAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildAnnotationAfterUpdateHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.BeforeDeleteHook, buildAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildAnnotationBeforeDeleteHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.AfterDeleteHook, buildAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildAnnotationAfterDeleteHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.BeforeUpsertHook, buildAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildAnnotationBeforeUpsertHooks = []BuildAnnotationHook{}

	AddBuildAnnotationHook(boil.AfterUpsertHook, buildAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildAnnotationAfterUpsertHooks = []BuildAnnotationHook{}
}

func testBuildAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildAnnotationToOneBuildUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BuildAnnotation
	var foreign Build

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildDBTypes, false, buildColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Build struct: %s", err)
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

	slice := BuildAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*BuildAnnotation)(&slice), nil); err != nil {
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

func testBuildAnnotationToOneSetOpBuildUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildAnnotation
	var b, c Build

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildAnnotationDBTypes, false, strmangle.SetComplement(buildAnnotationPrimaryKeyColumns, buildAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Build{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceBuildAnnotations[0] != &a {
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

func testBuildAnnotationToOneRemoveOpBuildUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildAnnotation
	var b Build

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildAnnotationDBTypes, false, strmangle.SetComplement(buildAnnotationPrimaryKeyColumns, buildAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildDBTypes, false, strmangle.SetComplement(buildPrimaryKeyColumns, buildColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceBuildAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBuildAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
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

func testBuildAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                      = bytes.MinRead
)

func testBuildAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildAnnotationAllColumns) == len(buildAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildAnnotationAllColumns) == len(buildAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildAnnotation{}
	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildAnnotationDBTypes, true, buildAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildAnnotationAllColumns, buildAnnotationPrimaryKeyColumns) {
		fields = buildAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildAnnotationAllColumns,
			buildAnnotationPrimaryKeyColumns,
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

	slice := BuildAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(buildAnnotationAllColumns) == len(buildAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BuildAnnotation{}
	if err = randomize.Struct(seed, &o, buildAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildAnnotation: %s", err)
	}

	count, err := BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildAnnotationDBTypes, false, buildAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildAnnotation: %s", err)
	}

	count, err = BuildAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
