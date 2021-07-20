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

func testLockings(t *testing.T) {
	t.Parallel()

	query := Lockings()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testLockingsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
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

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLockingsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Lockings().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLockingsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LockingSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testLockingsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := LockingExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Locking exists: %s", err)
	}
	if !e {
		t.Errorf("Expected LockingExists to return true, but got false.")
	}
}

func testLockingsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	lockingFound, err := FindLocking(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if lockingFound == nil {
		t.Error("want a record, got nil")
	}
}

func testLockingsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Lockings().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testLockingsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Lockings().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testLockingsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	lockingOne := &Locking{}
	lockingTwo := &Locking{}
	if err = randomize.Struct(seed, lockingOne, lockingDBTypes, false, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}
	if err = randomize.Struct(seed, lockingTwo, lockingDBTypes, false, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lockingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lockingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lockings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testLockingsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	lockingOne := &Locking{}
	lockingTwo := &Locking{}
	if err = randomize.Struct(seed, lockingOne, lockingDBTypes, false, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}
	if err = randomize.Struct(seed, lockingTwo, lockingDBTypes, false, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = lockingOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = lockingTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func lockingBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func lockingAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Locking) error {
	*o = Locking{}
	return nil
}

func testLockingsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Locking{}
	o := &Locking{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, lockingDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Locking object: %s", err)
	}

	AddLockingHook(boil.BeforeInsertHook, lockingBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	lockingBeforeInsertHooks = []LockingHook{}

	AddLockingHook(boil.AfterInsertHook, lockingAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	lockingAfterInsertHooks = []LockingHook{}

	AddLockingHook(boil.AfterSelectHook, lockingAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	lockingAfterSelectHooks = []LockingHook{}

	AddLockingHook(boil.BeforeUpdateHook, lockingBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	lockingBeforeUpdateHooks = []LockingHook{}

	AddLockingHook(boil.AfterUpdateHook, lockingAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	lockingAfterUpdateHooks = []LockingHook{}

	AddLockingHook(boil.BeforeDeleteHook, lockingBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	lockingBeforeDeleteHooks = []LockingHook{}

	AddLockingHook(boil.AfterDeleteHook, lockingAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	lockingAfterDeleteHooks = []LockingHook{}

	AddLockingHook(boil.BeforeUpsertHook, lockingBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	lockingBeforeUpsertHooks = []LockingHook{}

	AddLockingHook(boil.AfterUpsertHook, lockingAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	lockingAfterUpsertHooks = []LockingHook{}
}

func testLockingsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLockingsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(lockingColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testLockingsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
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

func testLockingsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := LockingSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testLockingsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Lockings().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	lockingDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`}
	_              = bytes.MinRead
)

func testLockingsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(lockingPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(lockingAllColumns) == len(lockingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testLockingsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(lockingAllColumns) == len(lockingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Locking{}
	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, lockingDBTypes, true, lockingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(lockingAllColumns, lockingPrimaryKeyColumns) {
		fields = lockingAllColumns
	} else {
		fields = strmangle.SetComplement(
			lockingAllColumns,
			lockingPrimaryKeyColumns,
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

	slice := LockingSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testLockingsUpsert(t *testing.T) {
	t.Parallel()

	if len(lockingAllColumns) == len(lockingPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Locking{}
	if err = randomize.Struct(seed, &o, lockingDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Locking: %s", err)
	}

	count, err := Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, lockingDBTypes, false, lockingPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Locking struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Locking: %s", err)
	}

	count, err = Lockings().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
