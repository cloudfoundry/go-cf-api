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

func testServiceBindingOperations(t *testing.T) {
	t.Parallel()

	query := ServiceBindingOperations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceBindingOperationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
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

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingOperationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceBindingOperations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingOperationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBindingOperationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceBindingOperationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceBindingOperationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceBindingOperation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceBindingOperationExists to return true, but got false.")
	}
}

func testServiceBindingOperationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceBindingOperationFound, err := FindServiceBindingOperation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceBindingOperationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceBindingOperationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceBindingOperations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceBindingOperationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceBindingOperations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceBindingOperationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceBindingOperationOne := &ServiceBindingOperation{}
	serviceBindingOperationTwo := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, serviceBindingOperationOne, serviceBindingOperationDBTypes, false, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBindingOperationTwo, serviceBindingOperationDBTypes, false, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBindingOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBindingOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBindingOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceBindingOperationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceBindingOperationOne := &ServiceBindingOperation{}
	serviceBindingOperationTwo := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, serviceBindingOperationOne, serviceBindingOperationDBTypes, false, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceBindingOperationTwo, serviceBindingOperationDBTypes, false, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceBindingOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceBindingOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceBindingOperationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func serviceBindingOperationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceBindingOperation) error {
	*o = ServiceBindingOperation{}
	return nil
}

func testServiceBindingOperationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceBindingOperation{}
	o := &ServiceBindingOperation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation object: %s", err)
	}

	AddServiceBindingOperationHook(boil.BeforeInsertHook, serviceBindingOperationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationBeforeInsertHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.AfterInsertHook, serviceBindingOperationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationAfterInsertHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.AfterSelectHook, serviceBindingOperationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationAfterSelectHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.BeforeUpdateHook, serviceBindingOperationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationBeforeUpdateHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.AfterUpdateHook, serviceBindingOperationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationAfterUpdateHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.BeforeDeleteHook, serviceBindingOperationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationBeforeDeleteHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.AfterDeleteHook, serviceBindingOperationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationAfterDeleteHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.BeforeUpsertHook, serviceBindingOperationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationBeforeUpsertHooks = []ServiceBindingOperationHook{}

	AddServiceBindingOperationHook(boil.AfterUpsertHook, serviceBindingOperationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceBindingOperationAfterUpsertHooks = []ServiceBindingOperationHook{}
}

func testServiceBindingOperationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBindingOperationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceBindingOperationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceBindingOperationToOneServiceBindingUsingServiceBinding(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceBindingOperation
	var foreign ServiceBinding

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceBindingDBTypes, false, serviceBindingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBinding struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ServiceBindingID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServiceBinding().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ServiceBindingOperationSlice{&local}
	if err = local.L.LoadServiceBinding(ctx, tx, false, (*[]*ServiceBindingOperation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceBinding == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ServiceBinding = nil
	if err = local.L.LoadServiceBinding(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceBinding == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServiceBindingOperationToOneSetOpServiceBindingUsingServiceBinding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBindingOperation
	var b, c ServiceBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBindingOperationDBTypes, false, strmangle.SetComplement(serviceBindingOperationPrimaryKeyColumns, serviceBindingOperationColumnsWithoutDefault)...); err != nil {
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
		err = a.SetServiceBinding(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ServiceBinding != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServiceBindingOperation != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ServiceBindingID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceBindingID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServiceBindingID))
		reflect.Indirect(reflect.ValueOf(&a.ServiceBindingID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ServiceBindingID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceBindingID, x.ID)
		}
	}
}

func testServiceBindingOperationToOneRemoveOpServiceBindingUsingServiceBinding(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceBindingOperation
	var b ServiceBinding

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceBindingOperationDBTypes, false, strmangle.SetComplement(serviceBindingOperationPrimaryKeyColumns, serviceBindingOperationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceBindingDBTypes, false, strmangle.SetComplement(serviceBindingPrimaryKeyColumns, serviceBindingColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetServiceBinding(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveServiceBinding(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ServiceBinding().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ServiceBinding != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ServiceBindingID) {
		t.Error("foreign key value should be nil")
	}

	if b.R.ServiceBindingOperation != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testServiceBindingOperationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
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

func testServiceBindingOperationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceBindingOperationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceBindingOperationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceBindingOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceBindingOperationDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ServiceBindingID`: `integer`, `State`: `character varying`, `Type`: `character varying`, `Description`: `character varying`, `BrokerProvidedOperation`: `character varying`}
	_                              = bytes.MinRead
)

func testServiceBindingOperationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceBindingOperationAllColumns) == len(serviceBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceBindingOperationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceBindingOperationAllColumns) == len(serviceBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceBindingOperation{}
	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceBindingOperationDBTypes, true, serviceBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceBindingOperationAllColumns, serviceBindingOperationPrimaryKeyColumns) {
		fields = serviceBindingOperationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceBindingOperationAllColumns,
			serviceBindingOperationPrimaryKeyColumns,
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

	slice := ServiceBindingOperationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceBindingOperationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceBindingOperationAllColumns) == len(serviceBindingOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceBindingOperation{}
	if err = randomize.Struct(seed, &o, serviceBindingOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBindingOperation: %s", err)
	}

	count, err := ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceBindingOperationDBTypes, false, serviceBindingOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceBindingOperation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceBindingOperation: %s", err)
	}

	count, err = ServiceBindingOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
