// +build integration
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

func testServiceKeyOperations(t *testing.T) {
	t.Parallel()

	query := ServiceKeyOperations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceKeyOperationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
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

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyOperationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceKeyOperations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyOperationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyOperationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyOperationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceKeyOperationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceKeyOperation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceKeyOperationExists to return true, but got false.")
	}
}

func testServiceKeyOperationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceKeyOperationFound, err := FindServiceKeyOperation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceKeyOperationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceKeyOperationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceKeyOperations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceKeyOperationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceKeyOperations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceKeyOperationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceKeyOperationOne := &ServiceKeyOperation{}
	serviceKeyOperationTwo := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, serviceKeyOperationOne, serviceKeyOperationDBTypes, false, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyOperationTwo, serviceKeyOperationDBTypes, false, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceKeyOperationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceKeyOperationOne := &ServiceKeyOperation{}
	serviceKeyOperationTwo := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, serviceKeyOperationOne, serviceKeyOperationDBTypes, false, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyOperationTwo, serviceKeyOperationDBTypes, false, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyOperationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyOperationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceKeyOperationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func serviceKeyOperationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyOperation) error {
	*o = ServiceKeyOperation{}
	return nil
}

func testServiceKeyOperationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceKeyOperation{}
	o := &ServiceKeyOperation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation object: %s", err)
	}

	AddServiceKeyOperationHook(boil.BeforeInsertHook, serviceKeyOperationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationBeforeInsertHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.AfterInsertHook, serviceKeyOperationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationAfterInsertHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.AfterSelectHook, serviceKeyOperationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationAfterSelectHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.BeforeUpdateHook, serviceKeyOperationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationBeforeUpdateHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.AfterUpdateHook, serviceKeyOperationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationAfterUpdateHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.BeforeDeleteHook, serviceKeyOperationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationBeforeDeleteHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.AfterDeleteHook, serviceKeyOperationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationAfterDeleteHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.BeforeUpsertHook, serviceKeyOperationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationBeforeUpsertHooks = []ServiceKeyOperationHook{}

	AddServiceKeyOperationHook(boil.AfterUpsertHook, serviceKeyOperationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyOperationAfterUpsertHooks = []ServiceKeyOperationHook{}
}

func testServiceKeyOperationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyOperationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceKeyOperationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyOperationToOneServiceKeyUsingServiceKey(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceKeyOperation
	var foreign ServiceKey

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceKeyDBTypes, false, serviceKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKey struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.ServiceKeyID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ServiceKey().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ServiceKeyOperationSlice{&local}
	if err = local.L.LoadServiceKey(ctx, tx, false, (*[]*ServiceKeyOperation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceKey == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ServiceKey = nil
	if err = local.L.LoadServiceKey(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ServiceKey == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testServiceKeyOperationToOneSetOpServiceKeyUsingServiceKey(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyOperation
	var b, c ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyOperationDBTypes, false, strmangle.SetComplement(serviceKeyOperationPrimaryKeyColumns, serviceKeyOperationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*ServiceKey{&b, &c} {
		err = a.SetServiceKey(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ServiceKey != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ServiceKeyOperation != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.ServiceKeyID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceKeyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ServiceKeyID))
		reflect.Indirect(reflect.ValueOf(&a.ServiceKeyID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ServiceKeyID, x.ID) {
			t.Error("foreign key was wrong value", a.ServiceKeyID, x.ID)
		}
	}
}

func testServiceKeyOperationToOneRemoveOpServiceKeyUsingServiceKey(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyOperation
	var b ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyOperationDBTypes, false, strmangle.SetComplement(serviceKeyOperationPrimaryKeyColumns, serviceKeyOperationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetServiceKey(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveServiceKey(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ServiceKey().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ServiceKey != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.ServiceKeyID) {
		t.Error("foreign key value should be nil")
	}

	if b.R.ServiceKeyOperation != nil {
		t.Error("failed to remove a from b's relationships")
	}

}

func testServiceKeyOperationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
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

func testServiceKeyOperationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyOperationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceKeyOperationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyOperations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceKeyOperationDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ServiceKeyID`: `integer`, `State`: `character varying`, `Type`: `character varying`, `Description`: `character varying`, `BrokerProvidedOperation`: `character varying`}
	_                          = bytes.MinRead
)

func testServiceKeyOperationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceKeyOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceKeyOperationAllColumns) == len(serviceKeyOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceKeyOperationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceKeyOperationAllColumns) == len(serviceKeyOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyOperation{}
	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyOperationDBTypes, true, serviceKeyOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceKeyOperationAllColumns, serviceKeyOperationPrimaryKeyColumns) {
		fields = serviceKeyOperationAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceKeyOperationAllColumns,
			serviceKeyOperationPrimaryKeyColumns,
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

	slice := ServiceKeyOperationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceKeyOperationsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceKeyOperationAllColumns) == len(serviceKeyOperationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceKeyOperation{}
	if err = randomize.Struct(seed, &o, serviceKeyOperationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyOperation: %s", err)
	}

	count, err := ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceKeyOperationDBTypes, false, serviceKeyOperationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyOperation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyOperation: %s", err)
	}

	count, err = ServiceKeyOperations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
