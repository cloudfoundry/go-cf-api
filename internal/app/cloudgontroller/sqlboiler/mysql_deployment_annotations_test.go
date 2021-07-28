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

func testDeploymentAnnotations(t *testing.T) {
	t.Parallel()

	query := DeploymentAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDeploymentAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
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

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DeploymentAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeploymentAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DeploymentAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DeploymentAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DeploymentAnnotationExists to return true, but got false.")
	}
}

func testDeploymentAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	deploymentAnnotationFound, err := FindDeploymentAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if deploymentAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDeploymentAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DeploymentAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDeploymentAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DeploymentAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDeploymentAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	deploymentAnnotationOne := &DeploymentAnnotation{}
	deploymentAnnotationTwo := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, deploymentAnnotationOne, deploymentAnnotationDBTypes, false, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, deploymentAnnotationTwo, deploymentAnnotationDBTypes, false, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deploymentAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deploymentAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DeploymentAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDeploymentAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	deploymentAnnotationOne := &DeploymentAnnotation{}
	deploymentAnnotationTwo := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, deploymentAnnotationOne, deploymentAnnotationDBTypes, false, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, deploymentAnnotationTwo, deploymentAnnotationDBTypes, false, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deploymentAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deploymentAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func deploymentAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func deploymentAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentAnnotation) error {
	*o = DeploymentAnnotation{}
	return nil
}

func testDeploymentAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DeploymentAnnotation{}
	o := &DeploymentAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation object: %s", err)
	}

	AddDeploymentAnnotationHook(boil.BeforeInsertHook, deploymentAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationBeforeInsertHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.AfterInsertHook, deploymentAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationAfterInsertHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.AfterSelectHook, deploymentAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationAfterSelectHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.BeforeUpdateHook, deploymentAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationBeforeUpdateHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.AfterUpdateHook, deploymentAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationAfterUpdateHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.BeforeDeleteHook, deploymentAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationBeforeDeleteHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.AfterDeleteHook, deploymentAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationAfterDeleteHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.BeforeUpsertHook, deploymentAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationBeforeUpsertHooks = []DeploymentAnnotationHook{}

	AddDeploymentAnnotationHook(boil.AfterUpsertHook, deploymentAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	deploymentAnnotationAfterUpsertHooks = []DeploymentAnnotationHook{}
}

func testDeploymentAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeploymentAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(deploymentAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeploymentAnnotationToOneDeploymentUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DeploymentAnnotation
	var foreign Deployment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, deploymentDBTypes, false, deploymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deployment struct: %s", err)
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

	slice := DeploymentAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*DeploymentAnnotation)(&slice), nil); err != nil {
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

func testDeploymentAnnotationToOneSetOpDeploymentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DeploymentAnnotation
	var b, c Deployment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deploymentAnnotationDBTypes, false, strmangle.SetComplement(deploymentAnnotationPrimaryKeyColumns, deploymentAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, deploymentDBTypes, false, strmangle.SetComplement(deploymentPrimaryKeyColumns, deploymentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, deploymentDBTypes, false, strmangle.SetComplement(deploymentPrimaryKeyColumns, deploymentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Deployment{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceDeploymentAnnotations[0] != &a {
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

func testDeploymentAnnotationToOneRemoveOpDeploymentUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DeploymentAnnotation
	var b Deployment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deploymentAnnotationDBTypes, false, strmangle.SetComplement(deploymentAnnotationPrimaryKeyColumns, deploymentAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, deploymentDBTypes, false, strmangle.SetComplement(deploymentPrimaryKeyColumns, deploymentColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceDeploymentAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDeploymentAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
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

func testDeploymentAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeploymentAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDeploymentAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DeploymentAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	deploymentAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                           = bytes.MinRead
)

func testDeploymentAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(deploymentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(deploymentAnnotationAllColumns) == len(deploymentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDeploymentAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(deploymentAnnotationAllColumns) == len(deploymentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentAnnotation{}
	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deploymentAnnotationDBTypes, true, deploymentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(deploymentAnnotationAllColumns, deploymentAnnotationPrimaryKeyColumns) {
		fields = deploymentAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			deploymentAnnotationAllColumns,
			deploymentAnnotationPrimaryKeyColumns,
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

	slice := DeploymentAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDeploymentAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(deploymentAnnotationAllColumns) == len(deploymentAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDeploymentAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DeploymentAnnotation{}
	if err = randomize.Struct(seed, &o, deploymentAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DeploymentAnnotation: %s", err)
	}

	count, err := DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, deploymentAnnotationDBTypes, false, deploymentAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DeploymentAnnotation: %s", err)
	}

	count, err = DeploymentAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
