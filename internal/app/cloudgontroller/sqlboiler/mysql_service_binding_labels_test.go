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

func testServiceBindingLabels(t *testing.T) {
	t.Parallel()

	query := ServiceBindingLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceBindingLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
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

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceBindingLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBindingLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceBindingLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceBindingLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceBindingLabelExists to return true, but got false.")
	}
}

func testServiceBindingLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceBindingLabelFound, err := FindServiceBindingLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceBindingLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceBindingLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceBindingLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceBindingLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceBindingLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceBindingLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceBindingLabelOne := &ServiceBindingLabel{}
	serviceBindingLabelTwo := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, serviceBindingLabelOne, serviceBindingLabelDBTypes, false, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBindingLabelTwo, serviceBindingLabelDBTypes, false, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBindingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBindingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBindingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceBindingLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceBindingLabelOne := &ServiceBindingLabel{}
	serviceBindingLabelTwo := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, serviceBindingLabelOne, serviceBindingLabelDBTypes, false, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBindingLabelTwo, serviceBindingLabelDBTypes, false, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBindingLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBindingLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBindingLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func serviceBindingLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingLabel) error {
	*o = ServiceBindingLabel{}
	return nil
}

func testServiceBindingLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceBindingLabel{}
	o := &ServiceBindingLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel object: %s", err)
	}

	AddServiceBindingLabelHook(boil.BeforeInsertHook, serviceBindingLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelBeforeInsertHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.AfterInsertHook, serviceBindingLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelAfterInsertHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.AfterSelectHook, serviceBindingLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelAfterSelectHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.BeforeUpdateHook, serviceBindingLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelBeforeUpdateHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.AfterUpdateHook, serviceBindingLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelAfterUpdateHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.BeforeDeleteHook, serviceBindingLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelBeforeDeleteHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.AfterDeleteHook, serviceBindingLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelAfterDeleteHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.BeforeUpsertHook, serviceBindingLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelBeforeUpsertHooks = []ServiceBindingLabelHook{}

	AddServiceBindingLabelHook(boil.AfterUpsertHook, serviceBindingLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingLabelAfterUpsertHooks = []ServiceBindingLabelHook{}
}

func testServiceBindingLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBindingLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceBindingLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBindingLabelToOneServiceBindingUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceBindingLabel
	var foreign ServiceBinding

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceBindingDBTypes, false, serviceBindingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBinding struct: %s", err)
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

	slice := ServiceBindingLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceBindingLabel)(&slice), nil); err != nil {
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

func testServiceBindingLabelToOneSetOpServiceBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBindingLabel
	var b, c ServiceBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBindingLabelDBTypes, false, strmangle.SetComplement(serviceBindingLabelPrimaryKeyColumns, serviceBindingLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBindingDBTypes, false, strmangle.SetComplement(serviceBindingPrimaryKeyColumns, serviceBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceBindingDBTypes, false, strmangle.SetComplement(serviceBindingPrimaryKeyColumns, serviceBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceBinding{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceBindingLabels[0] != &a {
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

func testServiceBindingLabelToOneRemoveOpServiceBindingUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBindingLabel
	var b ServiceBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBindingLabelDBTypes, false, strmangle.SetComplement(serviceBindingLabelPrimaryKeyColumns, serviceBindingLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBindingDBTypes, false, strmangle.SetComplement(serviceBindingPrimaryKeyColumns, serviceBindingColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceBindingLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceBindingLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
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

func testServiceBindingLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBindingLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceBindingLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBindingLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceBindingLabelDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `ResourceGUID`: `varchar`, `KeyPrefix`: `varchar`, `KeyName`: `varchar`, `Value`: `varchar`}
	_                          = bytes.MinRead
)

func testServiceBindingLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceBindingLabelAllColumns) == len(serviceBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceBindingLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceBindingLabelAllColumns) == len(serviceBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingLabel{}
	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBindingLabelDBTypes, true, serviceBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceBindingLabelAllColumns, serviceBindingLabelPrimaryKeyColumns) {
		fields = serviceBindingLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceBindingLabelAllColumns,
			serviceBindingLabelPrimaryKeyColumns,
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

	slice := ServiceBindingLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceBindingLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceBindingLabelAllColumns) == len(serviceBindingLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLServiceBindingLabelUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceBindingLabel{}
	if err = randomize.Struct(seed, &o, serviceBindingLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBindingLabel: %s", err)
	}

	count, err := ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceBindingLabelDBTypes, false, serviceBindingLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBindingLabel: %s", err)
	}

	count, err = ServiceBindingLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}