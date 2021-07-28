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

func testServiceInstanceLabels(t *testing.T) {
	t.Parallel()

	query := ServiceInstanceLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceInstanceLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
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

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceInstanceLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceInstanceLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceInstanceLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceInstanceLabelExists to return true, but got false.")
	}
}

func testServiceInstanceLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceInstanceLabelFound, err := FindServiceInstanceLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceInstanceLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceInstanceLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceInstanceLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceInstanceLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceInstanceLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceInstanceLabelOne := &ServiceInstanceLabel{}
	serviceInstanceLabelTwo := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, serviceInstanceLabelOne, serviceInstanceLabelDBTypes, false, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceLabelTwo, serviceInstanceLabelDBTypes, false, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceInstanceLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceInstanceLabelOne := &ServiceInstanceLabel{}
	serviceInstanceLabelTwo := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, serviceInstanceLabelOne, serviceInstanceLabelDBTypes, false, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceLabelTwo, serviceInstanceLabelDBTypes, false, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceInstanceLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func serviceInstanceLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceLabel) error {
	*o = ServiceInstanceLabel{}
	return nil
}

func testServiceInstanceLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceInstanceLabel{}
	o := &ServiceInstanceLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel object: %s", err)
	}

	AddServiceInstanceLabelHook(boil.BeforeInsertHook, serviceInstanceLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelBeforeInsertHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.AfterInsertHook, serviceInstanceLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelAfterInsertHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.AfterSelectHook, serviceInstanceLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelAfterSelectHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.BeforeUpdateHook, serviceInstanceLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelBeforeUpdateHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.AfterUpdateHook, serviceInstanceLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelAfterUpdateHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.BeforeDeleteHook, serviceInstanceLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelBeforeDeleteHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.AfterDeleteHook, serviceInstanceLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelAfterDeleteHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.BeforeUpsertHook, serviceInstanceLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelBeforeUpsertHooks = []ServiceInstanceLabelHook{}

	AddServiceInstanceLabelHook(boil.AfterUpsertHook, serviceInstanceLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceLabelAfterUpsertHooks = []ServiceInstanceLabelHook{}
}

func testServiceInstanceLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceInstanceLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceLabelToOneServiceInstanceUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceInstanceLabel
	var foreign ServiceInstance

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceInstanceDBTypes, false, serviceInstanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstance struct: %s", err)
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

	slice := ServiceInstanceLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceInstanceLabel)(&slice), nil); err != nil {
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

func testServiceInstanceLabelToOneSetOpServiceInstanceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceLabel
	var b, c ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceLabelDBTypes, false, strmangle.SetComplement(serviceInstanceLabelPrimaryKeyColumns, serviceInstanceLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceInstance{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceInstanceLabels[0] != &a {
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

func testServiceInstanceLabelToOneRemoveOpServiceInstanceUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceLabel
	var b ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceLabelDBTypes, false, strmangle.SetComplement(serviceInstanceLabelPrimaryKeyColumns, serviceInstanceLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceInstanceLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceInstanceLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
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

func testServiceInstanceLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceInstanceLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                           = bytes.MinRead
)

func testServiceInstanceLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceInstanceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceInstanceLabelAllColumns) == len(serviceInstanceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceInstanceLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceLabelAllColumns) == len(serviceInstanceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceLabel{}
	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceLabelDBTypes, true, serviceInstanceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceInstanceLabelAllColumns, serviceInstanceLabelPrimaryKeyColumns) {
		fields = serviceInstanceLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceInstanceLabelAllColumns,
			serviceInstanceLabelPrimaryKeyColumns,
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

	slice := ServiceInstanceLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceInstanceLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceLabelAllColumns) == len(serviceInstanceLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceInstanceLabel{}
	if err = randomize.Struct(seed, &o, serviceInstanceLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceLabel: %s", err)
	}

	count, err := ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceInstanceLabelDBTypes, false, serviceInstanceLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceLabel: %s", err)
	}

	count, err = ServiceInstanceLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
