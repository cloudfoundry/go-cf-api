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

func testServiceInstanceOperations(t *testing.T) {
	t.Parallel()

	query := ServiceInstanceOperations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceInstanceOperationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
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

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceOperationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceInstanceOperations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceOperationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceOperationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceInstanceOperationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceInstanceOperationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceInstanceOperation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceInstanceOperationExists to return true, but got false.")
	}
}

func testServiceInstanceOperationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceInstanceOperationFound, err := FindServiceInstanceOperation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceInstanceOperationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceInstanceOperationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceInstanceOperations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceOperationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceInstanceOperations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceInstanceOperationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceInstanceOperationOne := &ServiceInstanceOperation{}
	serviceInstanceOperationTwo := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, serviceInstanceOperationOne, serviceInstanceOperationDBTypes, false, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceOperationTwo, serviceInstanceOperationDBTypes, false, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceInstanceOperationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceInstanceOperationOne := &ServiceInstanceOperation{}
	serviceInstanceOperationTwo := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, serviceInstanceOperationOne, serviceInstanceOperationDBTypes, false, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceInstanceOperationTwo, serviceInstanceOperationDBTypes, false, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceInstanceOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceInstanceOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceInstanceOperationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func serviceInstanceOperationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceInstanceOperation) error {
	*o = ServiceInstanceOperation{}
	return nil
}

func testServiceInstanceOperationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceInstanceOperation{}
	o := &ServiceInstanceOperation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation object: %s", err)
	}

	AddServiceInstanceOperationHook(boil.BeforeInsertHook, serviceInstanceOperationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationBeforeInsertHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.AfterInsertHook, serviceInstanceOperationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationAfterInsertHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.AfterSelectHook, serviceInstanceOperationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationAfterSelectHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.BeforeUpdateHook, serviceInstanceOperationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationBeforeUpdateHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.AfterUpdateHook, serviceInstanceOperationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationAfterUpdateHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.BeforeDeleteHook, serviceInstanceOperationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationBeforeDeleteHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.AfterDeleteHook, serviceInstanceOperationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationAfterDeleteHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.BeforeUpsertHook, serviceInstanceOperationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationBeforeUpsertHooks = []ServiceInstanceOperationHook{}

	AddServiceInstanceOperationHook(boil.AfterUpsertHook, serviceInstanceOperationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceInstanceOperationAfterUpsertHooks = []ServiceInstanceOperationHook{}
}

func testServiceInstanceOperationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceOperationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceInstanceOperationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceInstanceOperationToOneServiceInstanceUsingServiceInstance(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceInstanceOperation
	var foreign ServiceInstance

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceInstanceDBTypes, false, serviceInstanceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstance struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ServiceInstanceID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServiceInstance().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ServiceInstanceOperationSlice{&local}
	if err = local.L.LoadServiceInstance(ctx, tx, false, (*[]*ServiceInstanceOperation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceInstance == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ServiceInstance = nil
	if err = local.L.LoadServiceInstance(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceInstance == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServiceInstanceOperationToOneSetOpServiceInstanceUsingServiceInstance(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceOperation
	var b, c ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceOperationDBTypes, false, strmangle.SetComplement(serviceInstanceOperationPrimaryKeyColumns, serviceInstanceOperationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetServiceInstance(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ServiceInstance != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServiceInstanceOperations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ServiceInstanceID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceInstanceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServiceInstanceID))
		reflect.Indirect(reflect.ValueOf(&a.ServiceInstanceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ServiceInstanceID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceInstanceID, x.ID)
		}
	}
}

func testServiceInstanceOperationToOneRemoveOpServiceInstanceUsingServiceInstance(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceInstanceOperation
	var b ServiceInstance

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceInstanceOperationDBTypes, false, strmangle.SetComplement(serviceInstanceOperationPrimaryKeyColumns, serviceInstanceOperationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceInstanceDBTypes, false, strmangle.SetComplement(serviceInstancePrimaryKeyColumns, serviceInstanceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetServiceInstance(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveServiceInstance(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ServiceInstance().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ServiceInstance != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ServiceInstanceID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.ServiceInstanceOperations) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceInstanceOperationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
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

func testServiceInstanceOperationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceInstanceOperationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceInstanceOperationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceInstanceOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceInstanceOperationDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ServiceInstanceID`: `integer`, `Type`: `text`, `State`: `text`, `Description`: `text`, `ProposedChanges`: `text`, `BrokerProvidedOperation`: `text`}
	_                               = bytes.MinRead
)

func testServiceInstanceOperationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceInstanceOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceInstanceOperationAllColumns) == len(serviceInstanceOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceInstanceOperationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceOperationAllColumns) == len(serviceInstanceOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceInstanceOperation{}
	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceInstanceOperationDBTypes, true, serviceInstanceOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceInstanceOperationAllColumns, serviceInstanceOperationPrimaryKeyColumns) {
		fields = serviceInstanceOperationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceInstanceOperationAllColumns,
			serviceInstanceOperationPrimaryKeyColumns,
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

	slice := ServiceInstanceOperationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceInstanceOperationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceInstanceOperationAllColumns) == len(serviceInstanceOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceInstanceOperation{}
	if err = randomize.Struct(seed, &o, serviceInstanceOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceOperation: %s", err)
	}

	count, err := ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceInstanceOperationDBTypes, false, serviceInstanceOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceInstanceOperation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceInstanceOperation: %s", err)
	}

	count, err = ServiceInstanceOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
