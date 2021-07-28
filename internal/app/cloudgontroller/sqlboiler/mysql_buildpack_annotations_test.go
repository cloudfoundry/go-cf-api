// +build mysql_integration
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

func testBuildpackAnnotations(t *testing.T) {
	t.Parallel()

	query := BuildpackAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildpackAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
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

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BuildpackAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpackAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildpackAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BuildpackAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildpackAnnotationExists to return true, but got false.")
	}
}

func testBuildpackAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildpackAnnotationFound, err := FindBuildpackAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildpackAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildpackAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BuildpackAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildpackAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BuildpackAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildpackAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildpackAnnotationOne := &BuildpackAnnotation{}
	buildpackAnnotationTwo := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, buildpackAnnotationOne, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackAnnotationTwo, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildpackAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildpackAnnotationOne := &BuildpackAnnotation{}
	buildpackAnnotationTwo := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, buildpackAnnotationOne, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackAnnotationTwo, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildpackAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func buildpackAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BuildpackAnnotation) error {
	*o = BuildpackAnnotation{}
	return nil
}

func testBuildpackAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BuildpackAnnotation{}
	o := &BuildpackAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation object: %s", err)
	}

	AddBuildpackAnnotationHook(boil.BeforeInsertHook, buildpackAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationBeforeInsertHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.AfterInsertHook, buildpackAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationAfterInsertHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.AfterSelectHook, buildpackAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationAfterSelectHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.BeforeUpdateHook, buildpackAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationBeforeUpdateHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.AfterUpdateHook, buildpackAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationAfterUpdateHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.BeforeDeleteHook, buildpackAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationBeforeDeleteHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.AfterDeleteHook, buildpackAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationAfterDeleteHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.BeforeUpsertHook, buildpackAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationBeforeUpsertHooks = []BuildpackAnnotationHook{}

	AddBuildpackAnnotationHook(boil.AfterUpsertHook, buildpackAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackAnnotationAfterUpsertHooks = []BuildpackAnnotationHook{}
}

func testBuildpackAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildpackAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackAnnotationToOneBuildpackUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BuildpackAnnotation
	var foreign Buildpack

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, buildpackDBTypes, false, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
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

	slice := BuildpackAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*BuildpackAnnotation)(&slice), nil); err != nil {
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

func testBuildpackAnnotationToOneSetOpBuildpackUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackAnnotation
	var b, c Buildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackAnnotationDBTypes, false, strmangle.SetComplement(buildpackAnnotationPrimaryKeyColumns, buildpackAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Buildpack{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceBuildpackAnnotations[0] != &a {
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

func testBuildpackAnnotationToOneRemoveOpBuildpackUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BuildpackAnnotation
	var b Buildpack

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackAnnotationDBTypes, false, strmangle.SetComplement(buildpackAnnotationPrimaryKeyColumns, buildpackAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceBuildpackAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testBuildpackAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
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

func testBuildpackAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildpackAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BuildpackAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildpackAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                          = bytes.MinRead
)

func testBuildpackAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildpackAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildpackAnnotationAllColumns) == len(buildpackAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildpackAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildpackAnnotationAllColumns) == len(buildpackAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BuildpackAnnotation{}
	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackAnnotationDBTypes, true, buildpackAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildpackAnnotationAllColumns, buildpackAnnotationPrimaryKeyColumns) {
		fields = buildpackAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildpackAnnotationAllColumns,
			buildpackAnnotationPrimaryKeyColumns,
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

	slice := BuildpackAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildpackAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(buildpackAnnotationAllColumns) == len(buildpackAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBuildpackAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BuildpackAnnotation{}
	if err = randomize.Struct(seed, &o, buildpackAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackAnnotation: %s", err)
	}

	count, err := BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildpackAnnotationDBTypes, false, buildpackAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BuildpackAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BuildpackAnnotation: %s", err)
	}

	count, err = BuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
