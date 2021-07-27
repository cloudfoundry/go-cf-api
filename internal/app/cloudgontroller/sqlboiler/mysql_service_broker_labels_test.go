// +build integration mysql
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

func testServiceBrokerLabels(t *testing.T) {
	t.Parallel()

	query := ServiceBrokerLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceBrokerLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
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

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceBrokerLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceBrokerLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceBrokerLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceBrokerLabelExists to return true, but got false.")
	}
}

func testServiceBrokerLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceBrokerLabelFound, err := FindServiceBrokerLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceBrokerLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceBrokerLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceBrokerLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceBrokerLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceBrokerLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceBrokerLabelOne := &ServiceBrokerLabel{}
	serviceBrokerLabelTwo := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, serviceBrokerLabelOne, serviceBrokerLabelDBTypes, false, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerLabelTwo, serviceBrokerLabelDBTypes, false, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceBrokerLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceBrokerLabelOne := &ServiceBrokerLabel{}
	serviceBrokerLabelTwo := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, serviceBrokerLabelOne, serviceBrokerLabelDBTypes, false, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerLabelTwo, serviceBrokerLabelDBTypes, false, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBrokerLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func serviceBrokerLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerLabel) error {
	*o = ServiceBrokerLabel{}
	return nil
}

func testServiceBrokerLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceBrokerLabel{}
	o := &ServiceBrokerLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel object: %s", err)
	}

	AddServiceBrokerLabelHook(boil.BeforeInsertHook, serviceBrokerLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelBeforeInsertHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.AfterInsertHook, serviceBrokerLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelAfterInsertHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.AfterSelectHook, serviceBrokerLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelAfterSelectHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.BeforeUpdateHook, serviceBrokerLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelBeforeUpdateHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.AfterUpdateHook, serviceBrokerLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelAfterUpdateHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.BeforeDeleteHook, serviceBrokerLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelBeforeDeleteHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.AfterDeleteHook, serviceBrokerLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelAfterDeleteHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.BeforeUpsertHook, serviceBrokerLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelBeforeUpsertHooks = []ServiceBrokerLabelHook{}

	AddServiceBrokerLabelHook(boil.AfterUpsertHook, serviceBrokerLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerLabelAfterUpsertHooks = []ServiceBrokerLabelHook{}
}

func testServiceBrokerLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceBrokerLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerLabelToOneServiceBrokerUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceBrokerLabel
	var foreign ServiceBroker

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceBrokerDBTypes, false, serviceBrokerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBroker struct: %s", err)
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

	slice := ServiceBrokerLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceBrokerLabel)(&slice), nil); err != nil {
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

func testServiceBrokerLabelToOneSetOpServiceBrokerUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerLabel
	var b, c ServiceBroker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerLabelDBTypes, false, strmangle.SetComplement(serviceBrokerLabelPrimaryKeyColumns, serviceBrokerLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerDBTypes, false, strmangle.SetComplement(serviceBrokerPrimaryKeyColumns, serviceBrokerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceBrokerDBTypes, false, strmangle.SetComplement(serviceBrokerPrimaryKeyColumns, serviceBrokerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceBroker{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceBrokerLabels[0] != &a {
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

func testServiceBrokerLabelToOneRemoveOpServiceBrokerUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerLabel
	var b ServiceBroker

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerLabelDBTypes, false, strmangle.SetComplement(serviceBrokerLabelPrimaryKeyColumns, serviceBrokerLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerDBTypes, false, strmangle.SetComplement(serviceBrokerPrimaryKeyColumns, serviceBrokerColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceBrokerLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceBrokerLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
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

func testServiceBrokerLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceBrokerLabelDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `KeyName`: `varchar`, `Value`: `varchar`}
	_                         = bytes.MinRead
)

func testServiceBrokerLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceBrokerLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceBrokerLabelAllColumns) == len(serviceBrokerLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceBrokerLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerLabelAllColumns) == len(serviceBrokerLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerLabelDBTypes, true, serviceBrokerLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceBrokerLabelAllColumns, serviceBrokerLabelPrimaryKeyColumns) {
		fields = serviceBrokerLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceBrokerLabelAllColumns,
			serviceBrokerLabelPrimaryKeyColumns,
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

	slice := ServiceBrokerLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceBrokerLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerLabelAllColumns) == len(serviceBrokerLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceBrokerLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceBrokerLabel{}
	if err = randomize.Struct(seed, &o, serviceBrokerLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerLabel: %s", err)
	}

	count, err := ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceBrokerLabelDBTypes, false, serviceBrokerLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerLabel: %s", err)
	}

	count, err = ServiceBrokerLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
