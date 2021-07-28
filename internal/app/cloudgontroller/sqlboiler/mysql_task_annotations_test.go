// +build mysql
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

func testTaskAnnotations(t *testing.T) {
	t.Parallel()

	query := TaskAnnotations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTaskAnnotationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
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

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskAnnotationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TaskAnnotations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskAnnotationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskAnnotationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskAnnotationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TaskAnnotationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if TaskAnnotation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskAnnotationExists to return true, but got false.")
	}
}

func testTaskAnnotationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	taskAnnotationFound, err := FindTaskAnnotation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if taskAnnotationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTaskAnnotationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TaskAnnotations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTaskAnnotationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TaskAnnotations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTaskAnnotationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskAnnotationOne := &TaskAnnotation{}
	taskAnnotationTwo := &TaskAnnotation{}
	if err = randomize.Struct(seed, taskAnnotationOne, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, taskAnnotationTwo, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTaskAnnotationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskAnnotationOne := &TaskAnnotation{}
	taskAnnotationTwo := &TaskAnnotation{}
	if err = randomize.Struct(seed, taskAnnotationOne, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}
	if err = randomize.Struct(seed, taskAnnotationTwo, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskAnnotationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskAnnotationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func taskAnnotationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func taskAnnotationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskAnnotation) error {
	*o = TaskAnnotation{}
	return nil
}

func testTaskAnnotationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TaskAnnotation{}
	o := &TaskAnnotation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation object: %s", err)
	}

	AddTaskAnnotationHook(boil.BeforeInsertHook, taskAnnotationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskAnnotationBeforeInsertHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.AfterInsertHook, taskAnnotationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskAnnotationAfterInsertHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.AfterSelectHook, taskAnnotationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskAnnotationAfterSelectHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.BeforeUpdateHook, taskAnnotationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskAnnotationBeforeUpdateHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.AfterUpdateHook, taskAnnotationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskAnnotationAfterUpdateHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.BeforeDeleteHook, taskAnnotationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskAnnotationBeforeDeleteHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.AfterDeleteHook, taskAnnotationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskAnnotationAfterDeleteHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.BeforeUpsertHook, taskAnnotationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskAnnotationBeforeUpsertHooks = []TaskAnnotationHook{}

	AddTaskAnnotationHook(boil.AfterUpsertHook, taskAnnotationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskAnnotationAfterUpsertHooks = []TaskAnnotationHook{}
}

func testTaskAnnotationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskAnnotationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(taskAnnotationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskAnnotationToOneTaskUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TaskAnnotation
	var foreign Task

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
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

	slice := TaskAnnotationSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*TaskAnnotation)(&slice), nil); err != nil {
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

func testTaskAnnotationToOneSetOpTaskUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TaskAnnotation
	var b, c Task

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskAnnotationDBTypes, false, strmangle.SetComplement(taskAnnotationPrimaryKeyColumns, taskAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Task{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceTaskAnnotations[0] != &a {
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

func testTaskAnnotationToOneRemoveOpTaskUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TaskAnnotation
	var b Task

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskAnnotationDBTypes, false, strmangle.SetComplement(taskAnnotationPrimaryKeyColumns, taskAnnotationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceTaskAnnotations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTaskAnnotationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
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

func testTaskAnnotationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskAnnotationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTaskAnnotationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskAnnotations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskAnnotationDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `Key`: `varchar`, `Value`: `varchar`}
	_                     = bytes.MinRead
)

func testTaskAnnotationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(taskAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(taskAnnotationAllColumns) == len(taskAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTaskAnnotationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskAnnotationAllColumns) == len(taskAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskAnnotation{}
	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskAnnotationDBTypes, true, taskAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskAnnotationAllColumns, taskAnnotationPrimaryKeyColumns) {
		fields = taskAnnotationAllColumns
	} else {
		fields = strmangle.SetComplement(
			taskAnnotationAllColumns,
			taskAnnotationPrimaryKeyColumns,
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

	slice := TaskAnnotationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTaskAnnotationsUpsert(t *testing.T) {
	t.Parallel()

	if len(taskAnnotationAllColumns) == len(taskAnnotationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTaskAnnotationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TaskAnnotation{}
	if err = randomize.Struct(seed, &o, taskAnnotationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskAnnotation: %s", err)
	}

	count, err := TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, taskAnnotationDBTypes, false, taskAnnotationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskAnnotation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskAnnotation: %s", err)
	}

	count, err = TaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
