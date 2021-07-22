// +build integration
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

func testPackageAnnotations(t *testing.T) {
	t.Parallel()

	query := PackageAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPackageAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
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

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackageAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := PackageAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackageAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PackageAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPackageAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PackageAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if PackageAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PackageAnnotationExists to return true, but got false.")
	}
}

func testPackageAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	packageAnnotationFound, err := FindPackageAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if packageAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPackageAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = PackageAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPackageAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := PackageAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPackageAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	packageAnnotationOne := &PackageAnnotation{}
	packageAnnotationTwo := &PackageAnnotation{}
	if err = randomize.Struct(seed, packageAnnotationOne, packageAnnotationDBTypes, false, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, packageAnnotationTwo, packageAnnotationDBTypes, false, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = packageAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = packageAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PackageAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPackageAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	packageAnnotationOne := &PackageAnnotation{}
	packageAnnotationTwo := &PackageAnnotation{}
	if err = randomize.Struct(seed, packageAnnotationOne, packageAnnotationDBTypes, false, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, packageAnnotationTwo, packageAnnotationDBTypes, false, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = packageAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = packageAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func packageAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func packageAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *PackageAnnotation) error {
	*o = PackageAnnotation{}
	return nil
}

func testPackageAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &PackageAnnotation{}
	o := &PackageAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation object: %s", err)
	}

	AddPackageAnnotationHook(boil.BeforeInsertHook, packageAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	packageAnnotationBeforeInsertHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.AfterInsertHook, packageAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	packageAnnotationAfterInsertHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.AfterSelectHook, packageAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	packageAnnotationAfterSelectHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.BeforeUpdateHook, packageAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	packageAnnotationBeforeUpdateHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.AfterUpdateHook, packageAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	packageAnnotationAfterUpdateHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.BeforeDeleteHook, packageAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	packageAnnotationBeforeDeleteHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.AfterDeleteHook, packageAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	packageAnnotationAfterDeleteHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.BeforeUpsertHook, packageAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	packageAnnotationBeforeUpsertHooks = []PackageAnnotationHook{}

	AddPackageAnnotationHook(boil.AfterUpsertHook, packageAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	packageAnnotationAfterUpsertHooks = []PackageAnnotationHook{}
}

func testPackageAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPackageAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(packageAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPackageAnnotationToOnePackageUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local PackageAnnotation
	var foreign Package

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, packageDBTypes, false, packageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Package struct: %s", err)
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

	slice := PackageAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*PackageAnnotation)(&slice), nil); err != nil {
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

func testPackageAnnotationToOneSetOpPackageUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PackageAnnotation
	var b, c Package

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, packageAnnotationDBTypes, false, strmangle.SetComplement(packageAnnotationPrimaryKeyColumns, packageAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, packageDBTypes, false, strmangle.SetComplement(packagePrimaryKeyColumns, packageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, packageDBTypes, false, strmangle.SetComplement(packagePrimaryKeyColumns, packageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Package{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourcePackageAnnotations[0] != &a {
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

func testPackageAnnotationToOneRemoveOpPackageUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a PackageAnnotation
	var b Package

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, packageAnnotationDBTypes, false, strmangle.SetComplement(packageAnnotationPrimaryKeyColumns, packageAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, packageDBTypes, false, strmangle.SetComplement(packagePrimaryKeyColumns, packageColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourcePackageAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testPackageAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
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

func testPackageAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PackageAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPackageAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := PackageAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	packageAnnotationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `Key`: `character varying`, `Value`: `character varying`}
	_                        = bytes.MinRead
)

func testPackageAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(packageAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(packageAnnotationAllColumns) == len(packageAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPackageAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(packageAnnotationAllColumns) == len(packageAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &PackageAnnotation{}
	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, packageAnnotationDBTypes, true, packageAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(packageAnnotationAllColumns, packageAnnotationPrimaryKeyColumns) {
		fields = packageAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			packageAnnotationAllColumns,
			packageAnnotationPrimaryKeyColumns,
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

	slice := PackageAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPackageAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(packageAnnotationAllColumns) == len(packageAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := PackageAnnotation{}
	if err = randomize.Struct(seed, &o, packageAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PackageAnnotation: %s", err)
	}

	count, err := PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, packageAnnotationDBTypes, false, packageAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize PackageAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert PackageAnnotation: %s", err)
	}

	count, err = PackageAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
