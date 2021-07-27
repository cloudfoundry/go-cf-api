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

func testDeploymentProcesses(t *testing.T) {
	t.Parallel()

	query := DeploymentProcesses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDeploymentProcessesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
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

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentProcessesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DeploymentProcesses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentProcessesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeploymentProcessSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeploymentProcessesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DeploymentProcessExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DeploymentProcess exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DeploymentProcessExists to return true, but got false.")
	}
}

func testDeploymentProcessesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	deploymentProcessFound, err := FindDeploymentProcess(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if deploymentProcessFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDeploymentProcessesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DeploymentProcesses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDeploymentProcessesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DeploymentProcesses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDeploymentProcessesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	deploymentProcessOne := &DeploymentProcess{}
	deploymentProcessTwo := &DeploymentProcess{}
	if err = randomize.Struct(seed, deploymentProcessOne, deploymentProcessDBTypes, false, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}
	if err = randomize.Struct(seed, deploymentProcessTwo, deploymentProcessDBTypes, false, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deploymentProcessOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deploymentProcessTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DeploymentProcesses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDeploymentProcessesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	deploymentProcessOne := &DeploymentProcess{}
	deploymentProcessTwo := &DeploymentProcess{}
	if err = randomize.Struct(seed, deploymentProcessOne, deploymentProcessDBTypes, false, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}
	if err = randomize.Struct(seed, deploymentProcessTwo, deploymentProcessDBTypes, false, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deploymentProcessOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deploymentProcessTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func deploymentProcessBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func deploymentProcessAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DeploymentProcess) error {
	*o = DeploymentProcess{}
	return nil
}

func testDeploymentProcessesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DeploymentProcess{}
	o := &DeploymentProcess{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess object: %s", err)
	}

	AddDeploymentProcessHook(boil.BeforeInsertHook, deploymentProcessBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	deploymentProcessBeforeInsertHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.AfterInsertHook, deploymentProcessAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	deploymentProcessAfterInsertHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.AfterSelectHook, deploymentProcessAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	deploymentProcessAfterSelectHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.BeforeUpdateHook, deploymentProcessBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	deploymentProcessBeforeUpdateHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.AfterUpdateHook, deploymentProcessAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	deploymentProcessAfterUpdateHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.BeforeDeleteHook, deploymentProcessBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	deploymentProcessBeforeDeleteHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.AfterDeleteHook, deploymentProcessAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	deploymentProcessAfterDeleteHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.BeforeUpsertHook, deploymentProcessBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	deploymentProcessBeforeUpsertHooks = []DeploymentProcessHook{}

	AddDeploymentProcessHook(boil.AfterUpsertHook, deploymentProcessAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	deploymentProcessAfterUpsertHooks = []DeploymentProcessHook{}
}

func testDeploymentProcessesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeploymentProcessesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(deploymentProcessColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeploymentProcessToOneDeploymentUsingDeployment(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DeploymentProcess
	var foreign Deployment

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, deploymentDBTypes, false, deploymentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deployment struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.DeploymentGUID, foreign.GUID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Deployment().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.GUID, foreign.GUID) {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := DeploymentProcessSlice{&local}
	if err = local.L.LoadDeployment(ctx, tx, false, (*[]*DeploymentProcess)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Deployment == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Deployment = nil
	if err = local.L.LoadDeployment(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Deployment == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDeploymentProcessToOneSetOpDeploymentUsingDeployment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DeploymentProcess
	var b, c Deployment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deploymentProcessDBTypes, false, strmangle.SetComplement(deploymentProcessPrimaryKeyColumns, deploymentProcessColumnsWithoutDefault)...); err != nil {
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
		err = a.SetDeployment(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Deployment != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DeploymentProcesses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.DeploymentGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.DeploymentGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DeploymentGUID))
		reflect.Indirect(reflect.ValueOf(&a.DeploymentGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.DeploymentGUID, x.GUID) {
			t.Error("foreign key was wrong value", a.DeploymentGUID, x.GUID)
		}
	}
}

func testDeploymentProcessToOneRemoveOpDeploymentUsingDeployment(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DeploymentProcess
	var b Deployment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deploymentProcessDBTypes, false, strmangle.SetComplement(deploymentProcessPrimaryKeyColumns, deploymentProcessColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, deploymentDBTypes, false, strmangle.SetComplement(deploymentPrimaryKeyColumns, deploymentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetDeployment(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveDeployment(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Deployment().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Deployment != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.DeploymentGUID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.DeploymentProcesses) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDeploymentProcessesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
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

func testDeploymentProcessesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeploymentProcessSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDeploymentProcessesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DeploymentProcesses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	deploymentProcessDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ProcessGUID`: `character varying`, `ProcessType`: `character varying`, `DeploymentGUID`: `character varying`}
	_                        = bytes.MinRead
)

func testDeploymentProcessesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(deploymentProcessPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(deploymentProcessAllColumns) == len(deploymentProcessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDeploymentProcessesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(deploymentProcessAllColumns) == len(deploymentProcessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DeploymentProcess{}
	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deploymentProcessDBTypes, true, deploymentProcessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(deploymentProcessAllColumns, deploymentProcessPrimaryKeyColumns) {
		fields = deploymentProcessAllColumns
	} else {
		fields = strmangle.SetComplement(
			deploymentProcessAllColumns,
			deploymentProcessPrimaryKeyColumns,
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

	slice := DeploymentProcessSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDeploymentProcessesUpsert(t *testing.T) {
	t.Parallel()

	if len(deploymentProcessAllColumns) == len(deploymentProcessPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DeploymentProcess{}
	if err = randomize.Struct(seed, &o, deploymentProcessDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DeploymentProcess: %s", err)
	}

	count, err := DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, deploymentProcessDBTypes, false, deploymentProcessPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DeploymentProcess struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DeploymentProcess: %s", err)
	}

	count, err = DeploymentProcesses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
