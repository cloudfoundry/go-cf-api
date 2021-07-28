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

func testServiceBrokerUpdateRequestLabels(t *testing.T) {
	t.Parallel()

	query := ServiceBrokerUpdateRequestLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceBrokerUpdateRequestLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
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

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceBrokerUpdateRequestLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerUpdateRequestLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBrokerUpdateRequestLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceBrokerUpdateRequestLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceBrokerUpdateRequestLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceBrokerUpdateRequestLabelExists to return true, but got false.")
	}
}

func testServiceBrokerUpdateRequestLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceBrokerUpdateRequestLabelFound, err := FindServiceBrokerUpdateRequestLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceBrokerUpdateRequestLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceBrokerUpdateRequestLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceBrokerUpdateRequestLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerUpdateRequestLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceBrokerUpdateRequestLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceBrokerUpdateRequestLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceBrokerUpdateRequestLabelOne := &ServiceBrokerUpdateRequestLabel{}
	serviceBrokerUpdateRequestLabelTwo := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestLabelOne, serviceBrokerUpdateRequestLabelDBTypes, false, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestLabelTwo, serviceBrokerUpdateRequestLabelDBTypes, false, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerUpdateRequestLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerUpdateRequestLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerUpdateRequestLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceBrokerUpdateRequestLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceBrokerUpdateRequestLabelOne := &ServiceBrokerUpdateRequestLabel{}
	serviceBrokerUpdateRequestLabelTwo := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestLabelOne, serviceBrokerUpdateRequestLabelDBTypes, false, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBrokerUpdateRequestLabelTwo, serviceBrokerUpdateRequestLabelDBTypes, false, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBrokerUpdateRequestLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBrokerUpdateRequestLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBrokerUpdateRequestLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func serviceBrokerUpdateRequestLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBrokerUpdateRequestLabel) error {
	*o = ServiceBrokerUpdateRequestLabel{}
	return nil
}

func testServiceBrokerUpdateRequestLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceBrokerUpdateRequestLabel{}
	o := &ServiceBrokerUpdateRequestLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel object: %s", err)
	}

	AddServiceBrokerUpdateRequestLabelHook(boil.BeforeInsertHook, serviceBrokerUpdateRequestLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelBeforeInsertHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.AfterInsertHook, serviceBrokerUpdateRequestLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelAfterInsertHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.AfterSelectHook, serviceBrokerUpdateRequestLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelAfterSelectHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.BeforeUpdateHook, serviceBrokerUpdateRequestLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelBeforeUpdateHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.AfterUpdateHook, serviceBrokerUpdateRequestLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelAfterUpdateHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.BeforeDeleteHook, serviceBrokerUpdateRequestLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelBeforeDeleteHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.AfterDeleteHook, serviceBrokerUpdateRequestLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelAfterDeleteHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.BeforeUpsertHook, serviceBrokerUpdateRequestLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelBeforeUpsertHooks = []ServiceBrokerUpdateRequestLabelHook{}

	AddServiceBrokerUpdateRequestLabelHook(boil.AfterUpsertHook, serviceBrokerUpdateRequestLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBrokerUpdateRequestLabelAfterUpsertHooks = []ServiceBrokerUpdateRequestLabelHook{}
}

func testServiceBrokerUpdateRequestLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerUpdateRequestLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceBrokerUpdateRequestLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBrokerUpdateRequestLabelToOneServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceBrokerUpdateRequestLabel
	var foreign ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceBrokerUpdateRequestDBTypes, false, serviceBrokerUpdateRequestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequest struct: %s", err)
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

	slice := ServiceBrokerUpdateRequestLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceBrokerUpdateRequestLabel)(&slice), nil); err != nil {
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

func testServiceBrokerUpdateRequestLabelToOneSetOpServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerUpdateRequestLabel
	var b, c ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerUpdateRequestLabelDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestLabelPrimaryKeyColumns, serviceBrokerUpdateRequestLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceBrokerUpdateRequest{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceBrokerUpdateRequestLabels[0] != &a {
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

func testServiceBrokerUpdateRequestLabelToOneRemoveOpServiceBrokerUpdateRequestUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBrokerUpdateRequestLabel
	var b ServiceBrokerUpdateRequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBrokerUpdateRequestLabelDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestLabelPrimaryKeyColumns, serviceBrokerUpdateRequestLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBrokerUpdateRequestDBTypes, false, strmangle.SetComplement(serviceBrokerUpdateRequestPrimaryKeyColumns, serviceBrokerUpdateRequestColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceBrokerUpdateRequestLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceBrokerUpdateRequestLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
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

func testServiceBrokerUpdateRequestLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBrokerUpdateRequestLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceBrokerUpdateRequestLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBrokerUpdateRequestLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceBrokerUpdateRequestLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                                      = bytes.MinRead
)

func testServiceBrokerUpdateRequestLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceBrokerUpdateRequestLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceBrokerUpdateRequestLabelAllColumns) == len(serviceBrokerUpdateRequestLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceBrokerUpdateRequestLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerUpdateRequestLabelAllColumns) == len(serviceBrokerUpdateRequestLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBrokerUpdateRequestLabelDBTypes, true, serviceBrokerUpdateRequestLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceBrokerUpdateRequestLabelAllColumns, serviceBrokerUpdateRequestLabelPrimaryKeyColumns) {
		fields = serviceBrokerUpdateRequestLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceBrokerUpdateRequestLabelAllColumns,
			serviceBrokerUpdateRequestLabelPrimaryKeyColumns,
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

	slice := ServiceBrokerUpdateRequestLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceBrokerUpdateRequestLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceBrokerUpdateRequestLabelAllColumns) == len(serviceBrokerUpdateRequestLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceBrokerUpdateRequestLabel{}
	if err = randomize.Struct(seed, &o, serviceBrokerUpdateRequestLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerUpdateRequestLabel: %s", err)
	}

	count, err := ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceBrokerUpdateRequestLabelDBTypes, false, serviceBrokerUpdateRequestLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBrokerUpdateRequestLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBrokerUpdateRequestLabel: %s", err)
	}

	count, err = ServiceBrokerUpdateRequestLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}