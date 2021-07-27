// +build integration postgres
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

func testServiceDashboardClients(t *testing.T) {
	t.Parallel()

	query := ServiceDashboardClients()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceDashboardClientsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
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

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceDashboardClientsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceDashboardClients().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceDashboardClientsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceDashboardClientSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceDashboardClientsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceDashboardClientExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceDashboardClient exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceDashboardClientExists to return true, but got false.")
	}
}

func testServiceDashboardClientsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceDashboardClientFound, err := FindServiceDashboardClient(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceDashboardClientFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceDashboardClientsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceDashboardClients().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceDashboardClientsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceDashboardClients().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceDashboardClientsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceDashboardClientOne := &ServiceDashboardClient{}
	serviceDashboardClientTwo := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, serviceDashboardClientOne, serviceDashboardClientDBTypes, false, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceDashboardClientTwo, serviceDashboardClientDBTypes, false, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceDashboardClientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceDashboardClientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceDashboardClients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceDashboardClientsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceDashboardClientOne := &ServiceDashboardClient{}
	serviceDashboardClientTwo := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, serviceDashboardClientOne, serviceDashboardClientDBTypes, false, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceDashboardClientTwo, serviceDashboardClientDBTypes, false, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceDashboardClientOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceDashboardClientTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceDashboardClientBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func serviceDashboardClientAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceDashboardClient) error {
	*o = ServiceDashboardClient{}
	return nil
}

func testServiceDashboardClientsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceDashboardClient{}
	o := &ServiceDashboardClient{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient object: %s", err)
	}

	AddServiceDashboardClientHook(boil.BeforeInsertHook, serviceDashboardClientBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientBeforeInsertHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.AfterInsertHook, serviceDashboardClientAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientAfterInsertHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.AfterSelectHook, serviceDashboardClientAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientAfterSelectHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.BeforeUpdateHook, serviceDashboardClientBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientBeforeUpdateHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.AfterUpdateHook, serviceDashboardClientAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientAfterUpdateHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.BeforeDeleteHook, serviceDashboardClientBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientBeforeDeleteHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.AfterDeleteHook, serviceDashboardClientAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientAfterDeleteHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.BeforeUpsertHook, serviceDashboardClientBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientBeforeUpsertHooks = []ServiceDashboardClientHook{}

	AddServiceDashboardClientHook(boil.AfterUpsertHook, serviceDashboardClientAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceDashboardClientAfterUpsertHooks = []ServiceDashboardClientHook{}
}

func testServiceDashboardClientsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceDashboardClientsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceDashboardClientColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceDashboardClientsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
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

func testServiceDashboardClientsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceDashboardClientSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceDashboardClientsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceDashboardClients().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceDashboardClientDBTypes = map[string]string{`ID`: `integer`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `UaaID`: `text`, `ServiceBrokerID`: `integer`}
	_                             = bytes.MinRead
)

func testServiceDashboardClientsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceDashboardClientPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceDashboardClientAllColumns) == len(serviceDashboardClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceDashboardClientsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceDashboardClientAllColumns) == len(serviceDashboardClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceDashboardClient{}
	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceDashboardClientDBTypes, true, serviceDashboardClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceDashboardClientAllColumns, serviceDashboardClientPrimaryKeyColumns) {
		fields = serviceDashboardClientAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceDashboardClientAllColumns,
			serviceDashboardClientPrimaryKeyColumns,
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

	slice := ServiceDashboardClientSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceDashboardClientsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceDashboardClientAllColumns) == len(serviceDashboardClientPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceDashboardClient{}
	if err = randomize.Struct(seed, &o, serviceDashboardClientDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceDashboardClient: %s", err)
	}

	count, err := ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceDashboardClientDBTypes, false, serviceDashboardClientPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceDashboardClient struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceDashboardClient: %s", err)
	}

	count, err = ServiceDashboardClients().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}