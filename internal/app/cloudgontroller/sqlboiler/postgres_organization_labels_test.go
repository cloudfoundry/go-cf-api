// +build integration,postgres
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

func testOrganizationLabels(t *testing.T) {
	t.Parallel()

	query := OrganizationLabels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrganizationLabelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
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

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationLabelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrganizationLabels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationLabelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationLabelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationLabelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrganizationLabelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OrganizationLabel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationLabelExists to return true, but got false.")
	}
}

func testOrganizationLabelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	organizationLabelFound, err := FindOrganizationLabel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if organizationLabelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrganizationLabelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrganizationLabels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrganizationLabelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrganizationLabels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationLabelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationLabelOne := &OrganizationLabel{}
	organizationLabelTwo := &OrganizationLabel{}
	if err = randomize.Struct(seed, organizationLabelOne, organizationLabelDBTypes, false, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationLabelTwo, organizationLabelDBTypes, false, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationLabelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationLabelOne := &OrganizationLabel{}
	organizationLabelTwo := &OrganizationLabel{}
	if err = randomize.Struct(seed, organizationLabelOne, organizationLabelDBTypes, false, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationLabelTwo, organizationLabelDBTypes, false, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = organizationLabelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = organizationLabelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func organizationLabelBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func organizationLabelAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *OrganizationLabel) error {
	*o = OrganizationLabel{}
	return nil
}

func testOrganizationLabelsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &OrganizationLabel{}
	o := &OrganizationLabel{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, false); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel object: %s", err)
	}

	AddOrganizationLabelHook(boil.BeforeInsertHook, organizationLabelBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationLabelBeforeInsertHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.AfterInsertHook, organizationLabelAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationLabelAfterInsertHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.AfterSelectHook, organizationLabelAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationLabelAfterSelectHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.BeforeUpdateHook, organizationLabelBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationLabelBeforeUpdateHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.AfterUpdateHook, organizationLabelAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationLabelAfterUpdateHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.BeforeDeleteHook, organizationLabelBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationLabelBeforeDeleteHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.AfterDeleteHook, organizationLabelAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationLabelAfterDeleteHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.BeforeUpsertHook, organizationLabelBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationLabelBeforeUpsertHooks = []OrganizationLabelHook{}

	AddOrganizationLabelHook(boil.AfterUpsertHook, organizationLabelAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationLabelAfterUpsertHooks = []OrganizationLabelHook{}
}

func testOrganizationLabelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationLabelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(organizationLabelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationLabelToOneOrganizationUsingResource(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local OrganizationLabel
	var foreign Organization

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
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

	slice := OrganizationLabelSlice{&local}
	if err = local.L.LoadResource(ctx, tx, false, (*[]*OrganizationLabel)(&slice), nil); err != nil {
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

func testOrganizationLabelToOneSetOpOrganizationUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationLabel
	var b, c Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationLabelDBTypes, false, strmangle.SetComplement(organizationLabelPrimaryKeyColumns, organizationLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Organization{&b, &c} {
		err = a.SetResource(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Resource != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ResourceOrganizationLabels[0] != &a {
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

func testOrganizationLabelToOneRemoveOpOrganizationUsingResource(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a OrganizationLabel
	var b Organization

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, organizationLabelDBTypes, false, strmangle.SetComplement(organizationLabelPrimaryKeyColumns, organizationLabelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, organizationDBTypes, false, strmangle.SetComplement(organizationPrimaryKeyColumns, organizationColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.ResourceOrganizationLabels) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOrganizationLabelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
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

func testOrganizationLabelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrganizationLabelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationLabelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrganizationLabels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationLabelDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `ResourceGUID`: `character varying`, `KeyPrefix`: `character varying`, `KeyName`: `character varying`, `Value`: `character varying`}
	_                        = bytes.MinRead
)

func testOrganizationLabelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(organizationLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(organizationLabelAllColumns) == len(organizationLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrganizationLabelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationLabelAllColumns) == len(organizationLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrganizationLabel{}
	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, organizationLabelDBTypes, true, organizationLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationLabelAllColumns, organizationLabelPrimaryKeyColumns) {
		fields = organizationLabelAllColumns
	} else {
		fields = strmangle.SetComplement(
			organizationLabelAllColumns,
			organizationLabelPrimaryKeyColumns,
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

	slice := OrganizationLabelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrganizationLabelsUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationLabelAllColumns) == len(organizationLabelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrganizationLabel{}
	if err = randomize.Struct(seed, &o, organizationLabelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationLabel: %s", err)
	}

	count, err := OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, organizationLabelDBTypes, false, organizationLabelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrganizationLabel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrganizationLabel: %s", err)
	}

	count, err = OrganizationLabels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
