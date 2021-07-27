// +build integration,mysql
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

func testRequestCounts(t *testing.T) {
	t.Parallel()

	query := RequestCounts()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRequestCountsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
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

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestCountsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RequestCounts().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestCountsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestCountSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestCountsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RequestCountExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RequestCount exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RequestCountExists to return true, but got false.")
	}
}

func testRequestCountsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	requestCountFound, err := FindRequestCount(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if requestCountFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRequestCountsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RequestCounts().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRequestCountsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RequestCounts().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRequestCountsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	requestCountOne := &RequestCount{}
	requestCountTwo := &RequestCount{}
	if err = randomize.Struct(seed, requestCountOne, requestCountDBTypes, false, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}
	if err = randomize.Struct(seed, requestCountTwo, requestCountDBTypes, false, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestCountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestCountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestCounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRequestCountsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	requestCountOne := &RequestCount{}
	requestCountTwo := &RequestCount{}
	if err = randomize.Struct(seed, requestCountOne, requestCountDBTypes, false, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}
	if err = randomize.Struct(seed, requestCountTwo, requestCountDBTypes, false, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestCountOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestCountTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func requestCountBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func requestCountAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestCount) error {
	*o = RequestCount{}
	return nil
}

func testRequestCountsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RequestCount{}
	o := &RequestCount{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, requestCountDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RequestCount object: %s", err)
	}

	AddRequestCountHook(boil.BeforeInsertHook, requestCountBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	requestCountBeforeInsertHooks = []RequestCountHook{}

	AddRequestCountHook(boil.AfterInsertHook, requestCountAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	requestCountAfterInsertHooks = []RequestCountHook{}

	AddRequestCountHook(boil.AfterSelectHook, requestCountAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	requestCountAfterSelectHooks = []RequestCountHook{}

	AddRequestCountHook(boil.BeforeUpdateHook, requestCountBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	requestCountBeforeUpdateHooks = []RequestCountHook{}

	AddRequestCountHook(boil.AfterUpdateHook, requestCountAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	requestCountAfterUpdateHooks = []RequestCountHook{}

	AddRequestCountHook(boil.BeforeDeleteHook, requestCountBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	requestCountBeforeDeleteHooks = []RequestCountHook{}

	AddRequestCountHook(boil.AfterDeleteHook, requestCountAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	requestCountAfterDeleteHooks = []RequestCountHook{}

	AddRequestCountHook(boil.BeforeUpsertHook, requestCountBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	requestCountBeforeUpsertHooks = []RequestCountHook{}

	AddRequestCountHook(boil.AfterUpsertHook, requestCountAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	requestCountAfterUpsertHooks = []RequestCountHook{}
}

func testRequestCountsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestCountsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(requestCountColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestCountsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
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

func testRequestCountsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestCountSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRequestCountsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestCounts().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	requestCountDBTypes = map[string]string{`ID`: `int`, `UserGUID`: `varchar`, `Count`: `int`, `ValidUntil`: `datetime`}
	_                   = bytes.MinRead
)

func testRequestCountsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(requestCountPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(requestCountAllColumns) == len(requestCountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRequestCountsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(requestCountAllColumns) == len(requestCountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestCount{}
	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestCountDBTypes, true, requestCountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(requestCountAllColumns, requestCountPrimaryKeyColumns) {
		fields = requestCountAllColumns
	} else {
		fields = strmangle.SetComplement(
			requestCountAllColumns,
			requestCountPrimaryKeyColumns,
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

	slice := RequestCountSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRequestCountsUpsert(t *testing.T) {
	t.Parallel()

	if len(requestCountAllColumns) == len(requestCountPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRequestCountUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RequestCount{}
	if err = randomize.Struct(seed, &o, requestCountDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestCount: %s", err)
	}

	count, err := RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, requestCountDBTypes, false, requestCountPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestCount struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestCount: %s", err)
	}

	count, err = RequestCounts().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
