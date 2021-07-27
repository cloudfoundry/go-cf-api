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

func testFeatureFlags(t *testing.T) {
	t.Parallel()

	query := FeatureFlags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFeatureFlagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
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

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureFlagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FeatureFlags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureFlagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FeatureFlagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFeatureFlagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FeatureFlagExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if FeatureFlag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FeatureFlagExists to return true, but got false.")
	}
}

func testFeatureFlagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	featureFlagFound, err := FindFeatureFlag(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if featureFlagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFeatureFlagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FeatureFlags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFeatureFlagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FeatureFlags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFeatureFlagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	featureFlagOne := &FeatureFlag{}
	featureFlagTwo := &FeatureFlag{}
	if err = randomize.Struct(seed, featureFlagOne, featureFlagDBTypes, false, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}
	if err = randomize.Struct(seed, featureFlagTwo, featureFlagDBTypes, false, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = featureFlagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = featureFlagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FeatureFlags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFeatureFlagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	featureFlagOne := &FeatureFlag{}
	featureFlagTwo := &FeatureFlag{}
	if err = randomize.Struct(seed, featureFlagOne, featureFlagDBTypes, false, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}
	if err = randomize.Struct(seed, featureFlagTwo, featureFlagDBTypes, false, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = featureFlagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = featureFlagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func featureFlagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func featureFlagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FeatureFlag) error {
	*o = FeatureFlag{}
	return nil
}

func testFeatureFlagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FeatureFlag{}
	o := &FeatureFlag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, featureFlagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureFlag object: %s", err)
	}

	AddFeatureFlagHook(boil.BeforeInsertHook, featureFlagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	featureFlagBeforeInsertHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.AfterInsertHook, featureFlagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	featureFlagAfterInsertHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.AfterSelectHook, featureFlagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	featureFlagAfterSelectHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.BeforeUpdateHook, featureFlagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	featureFlagBeforeUpdateHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.AfterUpdateHook, featureFlagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	featureFlagAfterUpdateHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.BeforeDeleteHook, featureFlagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	featureFlagBeforeDeleteHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.AfterDeleteHook, featureFlagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	featureFlagAfterDeleteHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.BeforeUpsertHook, featureFlagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	featureFlagBeforeUpsertHooks = []FeatureFlagHook{}

	AddFeatureFlagHook(boil.AfterUpsertHook, featureFlagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	featureFlagAfterUpsertHooks = []FeatureFlagHook{}
}

func testFeatureFlagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureFlagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(featureFlagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFeatureFlagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
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

func testFeatureFlagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FeatureFlagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFeatureFlagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FeatureFlags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	featureFlagDBTypes = map[string]string{`ID`: `int`, `GUID`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `Name`: `varchar`, `Enabled`: `tinyint`, `ErrorMessage`: `text`}
	_                  = bytes.MinRead
)

func testFeatureFlagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(featureFlagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(featureFlagAllColumns) == len(featureFlagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFeatureFlagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(featureFlagAllColumns) == len(featureFlagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FeatureFlag{}
	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, featureFlagDBTypes, true, featureFlagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(featureFlagAllColumns, featureFlagPrimaryKeyColumns) {
		fields = featureFlagAllColumns
	} else {
		fields = strmangle.SetComplement(
			featureFlagAllColumns,
			featureFlagPrimaryKeyColumns,
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

	slice := FeatureFlagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFeatureFlagsUpsert(t *testing.T) {
	t.Parallel()

	if len(featureFlagAllColumns) == len(featureFlagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFeatureFlagUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FeatureFlag{}
	if err = randomize.Struct(seed, &o, featureFlagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FeatureFlag: %s", err)
	}

	count, err := FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, featureFlagDBTypes, false, featureFlagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FeatureFlag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FeatureFlag: %s", err)
	}

	count, err = FeatureFlags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
