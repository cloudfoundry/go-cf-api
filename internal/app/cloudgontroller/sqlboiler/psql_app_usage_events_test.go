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

func testAppUsageEvents(t *testing.T) {
	t.Parallel()

	query := AppUsageEvents()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAppUsageEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
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

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUsageEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AppUsageEvents().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUsageEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppUsageEventSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAppUsageEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AppUsageEventExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AppUsageEvent exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AppUsageEventExists to return true, but got false.")
	}
}

func testAppUsageEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	appUsageEventFound, err := FindAppUsageEvent(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if appUsageEventFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAppUsageEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AppUsageEvents().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAppUsageEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AppUsageEvents().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAppUsageEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	appUsageEventOne := &AppUsageEvent{}
	appUsageEventTwo := &AppUsageEvent{}
	if err = randomize.Struct(seed, appUsageEventOne, appUsageEventDBTypes, false, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, appUsageEventTwo, appUsageEventDBTypes, false, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appUsageEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appUsageEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppUsageEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAppUsageEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	appUsageEventOne := &AppUsageEvent{}
	appUsageEventTwo := &AppUsageEvent{}
	if err = randomize.Struct(seed, appUsageEventOne, appUsageEventDBTypes, false, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}
	if err = randomize.Struct(seed, appUsageEventTwo, appUsageEventDBTypes, false, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = appUsageEventOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = appUsageEventTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func appUsageEventBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func appUsageEventAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AppUsageEvent) error {
	*o = AppUsageEvent{}
	return nil
}

func testAppUsageEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AppUsageEvent{}
	o := &AppUsageEvent{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent object: %s", err)
	}

	AddAppUsageEventHook(boil.BeforeInsertHook, appUsageEventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	appUsageEventBeforeInsertHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.AfterInsertHook, appUsageEventAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	appUsageEventAfterInsertHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.AfterSelectHook, appUsageEventAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	appUsageEventAfterSelectHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.BeforeUpdateHook, appUsageEventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	appUsageEventBeforeUpdateHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.AfterUpdateHook, appUsageEventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	appUsageEventAfterUpdateHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.BeforeDeleteHook, appUsageEventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	appUsageEventBeforeDeleteHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.AfterDeleteHook, appUsageEventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	appUsageEventAfterDeleteHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.BeforeUpsertHook, appUsageEventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	appUsageEventBeforeUpsertHooks = []AppUsageEventHook{}

	AddAppUsageEventHook(boil.AfterUpsertHook, appUsageEventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	appUsageEventAfterUpsertHooks = []AppUsageEventHook{}
}

func testAppUsageEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppUsageEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(appUsageEventColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAppUsageEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
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

func testAppUsageEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AppUsageEventSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAppUsageEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AppUsageEvents().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	appUsageEventDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `InstanceCount`: `integer`, `MemoryInMBPerInstance`: `integer`, `State`: `text`, `AppGUID`: `text`, `AppName`: `text`, `SpaceGUID`: `text`, `SpaceName`: `text`, `OrgGUID`: `text`, `BuildpackGUID`: `text`, `BuildpackName`: `text`, `PackageState`: `text`, `ParentAppName`: `text`, `ParentAppGUID`: `text`, `ProcessType`: `text`, `TaskGUID`: `text`, `TaskName`: `text`, `PackageGUID`: `text`, `PreviousState`: `text`, `PreviousPackageState`: `text`, `PreviousMemoryInMBPerInstance`: `integer`, `PreviousInstanceCount`: `integer`}
	_                    = bytes.MinRead
)

func testAppUsageEventsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(appUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(appUsageEventAllColumns) == len(appUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAppUsageEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(appUsageEventAllColumns) == len(appUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AppUsageEvent{}
	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, appUsageEventDBTypes, true, appUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(appUsageEventAllColumns, appUsageEventPrimaryKeyColumns) {
		fields = appUsageEventAllColumns
	} else {
		fields = strmangle.SetComplement(
			appUsageEventAllColumns,
			appUsageEventPrimaryKeyColumns,
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

	slice := AppUsageEventSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAppUsageEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(appUsageEventAllColumns) == len(appUsageEventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AppUsageEvent{}
	if err = randomize.Struct(seed, &o, appUsageEventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppUsageEvent: %s", err)
	}

	count, err := AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, appUsageEventDBTypes, false, appUsageEventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AppUsageEvent struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AppUsageEvent: %s", err)
	}

	count, err = AppUsageEvents().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
