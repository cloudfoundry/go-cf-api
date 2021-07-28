// +build psql,db
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

func testServiceUsageEvents(t *testing.T) {
	t.Parallel()

	query := ServiceUsageEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceUsageEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
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

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceUsageEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceUsageEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceUsageEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceUsageEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceUsageEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceUsageEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceUsageEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceUsageEventExists to return true, but got false.")
	}
}

func testServiceUsageEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceUsageEventFound, err := FindServiceUsageEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceUsageEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceUsageEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceUsageEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceUsageEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceUsageEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceUsageEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceUsageEventOne := &ServiceUsageEvent{}
	serviceUsageEventTwo := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, serviceUsageEventOne, serviceUsageEventDBTypes, false, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceUsageEventTwo, serviceUsageEventDBTypes, false, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceUsageEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceUsageEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceUsageEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceUsageEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceUsageEventOne := &ServiceUsageEvent{}
	serviceUsageEventTwo := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, serviceUsageEventOne, serviceUsageEventDBTypes, false, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceUsageEventTwo, serviceUsageEventDBTypes, false, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceUsageEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceUsageEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceUsageEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func serviceUsageEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceUsageEvent) error {
	*o = ServiceUsageEvent{}
	return nil
}

func testServiceUsageEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceUsageEvent{}
	o := &ServiceUsageEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent object: %s", err)
	}

	AddServiceUsageEventHook(boil.BeforeInsertHook, serviceUsageEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventBeforeInsertHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.AfterInsertHook, serviceUsageEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventAfterInsertHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.AfterSelectHook, serviceUsageEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventAfterSelectHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.BeforeUpdateHook, serviceUsageEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventBeforeUpdateHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.AfterUpdateHook, serviceUsageEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventAfterUpdateHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.BeforeDeleteHook, serviceUsageEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventBeforeDeleteHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.AfterDeleteHook, serviceUsageEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventAfterDeleteHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.BeforeUpsertHook, serviceUsageEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventBeforeUpsertHooks = []ServiceUsageEventHook{}

	AddServiceUsageEventHook(boil.AfterUpsertHook, serviceUsageEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceUsageEventAfterUpsertHooks = []ServiceUsageEventHook{}
}

func testServiceUsageEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceUsageEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceUsageEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceUsageEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
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

func testServiceUsageEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceUsageEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceUsageEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceUsageEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceUsageEventDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `State`: `text`, `OrgGUID`: `text`, `SpaceGUID`: `text`, `SpaceName`: `text`, `ServiceInstanceGUID`: `text`, `ServiceInstanceName`: `text`, `ServiceInstanceType`: `text`, `ServicePlanGUID`: `text`, `ServicePlanName`: `text`, `ServiceGUID`: `text`, `ServiceLabel`: `text`, `ServiceBrokerName`: `text`, `ServiceBrokerGUID`: `character varying`}
	_                        = bytes.MinRead
)

func testServiceUsageEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceUsageEventAllColumns) == len(serviceUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceUsageEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceUsageEventAllColumns) == len(serviceUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceUsageEvent{}
	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceUsageEventDBTypes, true, serviceUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceUsageEventAllColumns, serviceUsageEventPrimaryKeyColumns) {
		fields = serviceUsageEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceUsageEventAllColumns,
			serviceUsageEventPrimaryKeyColumns,
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

	slice := ServiceUsageEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceUsageEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceUsageEventAllColumns) == len(serviceUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceUsageEvent{}
	if err = randomize.Struct(seed, &o, serviceUsageEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceUsageEvent: %s", err)
	}

	count, err := ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceUsageEventDBTypes, false, serviceUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceUsageEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceUsageEvent: %s", err)
	}

	count, err = ServiceUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
