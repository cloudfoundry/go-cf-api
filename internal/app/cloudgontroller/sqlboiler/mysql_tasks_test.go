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

func testTasks(t *testing.T) {
	t.Parallel()

	query := Tasks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
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

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Tasks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TaskExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Task exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskExists to return true, but got false.")
	}
}

func testTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	taskFound, err := FindTask(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if taskFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Tasks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Tasks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func taskBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Task) error {
	*o = Task{}
	return nil
}

func testTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Task{}
	o := &Task{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Task object: %s", err)
	}

	AddTaskHook(boil.BeforeInsertHook, taskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterInsertHook, taskAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskAfterInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterSelectHook, taskAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskAfterSelectHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpdateHook, taskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpdateHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpdateHook, taskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskAfterUpdateHooks = []TaskHook{}

	AddTaskHook(boil.BeforeDeleteHook, taskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskBeforeDeleteHooks = []TaskHook{}

	AddTaskHook(boil.AfterDeleteHook, taskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskAfterDeleteHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpsertHook, taskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpsertHook, taskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskAfterUpsertHooks = []TaskHook{}
}

func testTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(taskColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskToManyResourceTaskAnnotations(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c TaskAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, taskAnnotationDBTypes, false, taskAnnotationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ResourceGUID, a.GUID)
	queries.Assign(&c.ResourceGUID, a.GUID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ResourceTaskAnnotations().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ResourceGUID, b.ResourceGUID) {
			bFound = true
		}
		if queries.Equal(v.ResourceGUID, c.ResourceGUID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TaskSlice{&a}
	if err = a.L.LoadResourceTaskAnnotations(ctx, tx, false, (*[]*Task)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceTaskAnnotations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ResourceTaskAnnotations = nil
	if err = a.L.LoadResourceTaskAnnotations(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceTaskAnnotations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTaskToManyResourceTaskLabels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c TaskLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, taskLabelDBTypes, false, taskLabelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, taskLabelDBTypes, false, taskLabelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ResourceGUID, a.GUID)
	queries.Assign(&c.ResourceGUID, a.GUID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ResourceTaskLabels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ResourceGUID, b.ResourceGUID) {
			bFound = true
		}
		if queries.Equal(v.ResourceGUID, c.ResourceGUID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TaskSlice{&a}
	if err = a.L.LoadResourceTaskLabels(ctx, tx, false, (*[]*Task)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceTaskLabels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ResourceTaskLabels = nil
	if err = a.L.LoadResourceTaskLabels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceTaskLabels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testTaskToManyAddOpResourceTaskAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskAnnotationDBTypes, false, strmangle.SetComplement(taskAnnotationPrimaryKeyColumns, taskAnnotationColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*TaskAnnotation{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddResourceTaskAnnotations(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.GUID, first.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, first.ResourceGUID)
		}
		if !queries.Equal(a.GUID, second.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, second.ResourceGUID)
		}

		if first.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ResourceTaskAnnotations[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ResourceTaskAnnotations[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ResourceTaskAnnotations().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTaskToManySetOpResourceTaskAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskAnnotationDBTypes, false, strmangle.SetComplement(taskAnnotationPrimaryKeyColumns, taskAnnotationColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetResourceTaskAnnotations(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceTaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetResourceTaskAnnotations(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceTaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.GUID, d.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, d.ResourceGUID)
	}
	if !queries.Equal(a.GUID, e.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, e.ResourceGUID)
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ResourceTaskAnnotations[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ResourceTaskAnnotations[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTaskToManyRemoveOpResourceTaskAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskAnnotationDBTypes, false, strmangle.SetComplement(taskAnnotationPrimaryKeyColumns, taskAnnotationColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddResourceTaskAnnotations(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceTaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveResourceTaskAnnotations(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceTaskAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ResourceTaskAnnotations) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ResourceTaskAnnotations[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ResourceTaskAnnotations[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTaskToManyAddOpResourceTaskLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskLabelDBTypes, false, strmangle.SetComplement(taskLabelPrimaryKeyColumns, taskLabelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*TaskLabel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddResourceTaskLabels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.GUID, first.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, first.ResourceGUID)
		}
		if !queries.Equal(a.GUID, second.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, second.ResourceGUID)
		}

		if first.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ResourceTaskLabels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ResourceTaskLabels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ResourceTaskLabels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTaskToManySetOpResourceTaskLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskLabelDBTypes, false, strmangle.SetComplement(taskLabelPrimaryKeyColumns, taskLabelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetResourceTaskLabels(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceTaskLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetResourceTaskLabels(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceTaskLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.GUID, d.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, d.ResourceGUID)
	}
	if !queries.Equal(a.GUID, e.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, e.ResourceGUID)
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ResourceTaskLabels[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ResourceTaskLabels[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTaskToManyRemoveOpResourceTaskLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c, d, e TaskLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TaskLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, taskLabelDBTypes, false, strmangle.SetComplement(taskLabelPrimaryKeyColumns, taskLabelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddResourceTaskLabels(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceTaskLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveResourceTaskLabels(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceTaskLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ResourceTaskLabels) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ResourceTaskLabels[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ResourceTaskLabels[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTaskToOneAppUsingApp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Task
	var foreign App

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, appDBTypes, false, appColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize App struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AppGUID = foreign.GUID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.App().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.GUID != foreign.GUID {
		t.Errorf("want: %v, got %v", foreign.GUID, check.GUID)
	}

	slice := TaskSlice{&local}
	if err = local.L.LoadApp(ctx, tx, false, (*[]*Task)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.App = nil
	if err = local.L.LoadApp(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.App == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTaskToOneSetOpAppUsingApp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Task
	var b, c App

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
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
		err = a.SetApp(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.App != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Tasks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AppGUID != x.GUID {
			t.Error("foreign key was wrong value", a.AppGUID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AppGUID))
		reflect.Indirect(reflect.ValueOf(&a.AppGUID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AppGUID != x.GUID {
			t.Error("foreign key was wrong value", a.AppGUID, x.GUID)
		}
	}
}

func testTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
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

func testTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Tasks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `Name`: `varchar`, `Command`: `text`, `State`: `varchar`, `MemoryInMB`: `int`, `EncryptedEnvironmentVariables`: `text`, `Salt`: `varchar`, `FailureReason`: `varchar`, `AppGUID`: `varchar`, `DropletGUID`: `varchar`, `SequenceID`: `int`, `DiskInMB`: `int`, `EncryptionKeyLabel`: `varchar`, `EncryptionIterations`: `int`}
	_           = bytes.MinRead
)

func testTasksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskDBTypes, true, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Task{}
	if err = randomize.Struct(seed, o, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskDBTypes, true, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskAllColumns, taskPrimaryKeyColumns) {
		fields = taskAllColumns
	} else {
		fields = strmangle.SetComplement(
			taskAllColumns,
			taskPrimaryKeyColumns,
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

	slice := TaskSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(taskAllColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLTaskUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Task{}
	if err = randomize.Struct(seed, &o, taskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err := Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, taskDBTypes, false, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err = Tasks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
