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

func testServiceKeyLabels(t *testing.T) {
	t.Parallel()

	query := ServiceKeyLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testServiceKeyLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
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

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ServiceKeyLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testServiceKeyLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ServiceKeyLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if ServiceKeyLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ServiceKeyLabelExists to return true, but got false.")
	}
}

func testServiceKeyLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	serviceKeyLabelFound, err := FindServiceKeyLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if serviceKeyLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testServiceKeyLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ServiceKeyLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testServiceKeyLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ServiceKeyLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testServiceKeyLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	serviceKeyLabelOne := &ServiceKeyLabel{}
	serviceKeyLabelTwo := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, serviceKeyLabelOne, serviceKeyLabelDBTypes, false, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyLabelTwo, serviceKeyLabelDBTypes, false, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testServiceKeyLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	serviceKeyLabelOne := &ServiceKeyLabel{}
	serviceKeyLabelTwo := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, serviceKeyLabelOne, serviceKeyLabelDBTypes, false, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, serviceKeyLabelTwo, serviceKeyLabelDBTypes, false, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = serviceKeyLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = serviceKeyLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func serviceKeyLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func serviceKeyLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *ServiceKeyLabel) error {
	*o = ServiceKeyLabel{}
	return nil
}

func testServiceKeyLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &ServiceKeyLabel{}
	o := &ServiceKeyLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel object: %s", err)
	}

	AddServiceKeyLabelHook(boil.BeforeInsertHook, serviceKeyLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelBeforeInsertHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.AfterInsertHook, serviceKeyLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelAfterInsertHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.AfterSelectHook, serviceKeyLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelAfterSelectHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.BeforeUpdateHook, serviceKeyLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelBeforeUpdateHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.AfterUpdateHook, serviceKeyLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelAfterUpdateHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.BeforeDeleteHook, serviceKeyLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelBeforeDeleteHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.AfterDeleteHook, serviceKeyLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelAfterDeleteHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.BeforeUpsertHook, serviceKeyLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelBeforeUpsertHooks = []ServiceKeyLabelHook{}

	AddServiceKeyLabelHook(boil.AfterUpsertHook, serviceKeyLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	serviceKeyLabelAfterUpsertHooks = []ServiceKeyLabelHook{}
}

func testServiceKeyLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(serviceKeyLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testServiceKeyLabelToOneServiceKeyUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local ServiceKeyLabel
	var foreign ServiceKey

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, serviceKeyDBTypes, false, serviceKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKey struct: %s", err)
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

	slice := ServiceKeyLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*ServiceKeyLabel)(&slice), nil); err != nil {
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

func testServiceKeyLabelToOneSetOpServiceKeyUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyLabel
	var b, c ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyLabelDBTypes, false, strmangle.SetComplement(serviceKeyLabelPrimaryKeyColumns, serviceKeyLabelColumnsWithoutDefault)...); err != nil {
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
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceServiceKeyLabels[0] != &a {
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

func testServiceKeyLabelToOneRemoveOpServiceKeyUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a ServiceKeyLabel
	var b ServiceKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, serviceKeyLabelDBTypes, false, strmangle.SetComplement(serviceKeyLabelPrimaryKeyColumns, serviceKeyLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, serviceKeyDBTypes, false, strmangle.SetComplement(serviceKeyPrimaryKeyColumns, serviceKeyColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceServiceKeyLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testServiceKeyLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
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

func testServiceKeyLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ServiceKeyLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testServiceKeyLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ServiceKeyLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	serviceKeyLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                      = bytes.MinRead
)

func testServiceKeyLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(serviceKeyLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(serviceKeyLabelAllColumns) == len(serviceKeyLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testServiceKeyLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(serviceKeyLabelAllColumns) == len(serviceKeyLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ServiceKeyLabel{}
	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, serviceKeyLabelDBTypes, true, serviceKeyLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(serviceKeyLabelAllColumns, serviceKeyLabelPrimaryKeyColumns) {
		fields = serviceKeyLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			serviceKeyLabelAllColumns,
			serviceKeyLabelPrimaryKeyColumns,
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

	slice := ServiceKeyLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testServiceKeyLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(serviceKeyLabelAllColumns) == len(serviceKeyLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ServiceKeyLabel{}
	if err = randomize.Struct(seed, &o, serviceKeyLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyLabel: %s", err)
	}

	count, err := ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, serviceKeyLabelDBTypes, false, serviceKeyLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ServiceKeyLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ServiceKeyLabel: %s", err)
	}

	count, err = ServiceKeyLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
