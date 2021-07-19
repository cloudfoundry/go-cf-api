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

func testJets(t *testing.T) {
	t.Parallel()

	query := Jets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
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

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Jets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Jet exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JetExists to return true, but got false.")
	}
}

func testJetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	jetFound, err := FindJet(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if jetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Jets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Jets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jetOne := &Jet{}
	jetTwo := &Jet{}
	if err = randomize.Struct(seed, jetOne, jetDBTypes, false, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}
	if err = randomize.Struct(seed, jetTwo, jetDBTypes, false, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jetOne := &Jet{}
	jetTwo := &Jet{}
	if err = randomize.Struct(seed, jetOne, jetDBTypes, false, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}
	if err = randomize.Struct(seed, jetTwo, jetDBTypes, false, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func jetBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func jetAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Jet) error {
	*o = Jet{}
	return nil
}

func testJetsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Jet{}
	o := &Jet{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jetDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Jet object: %s", err)
	}

	AddJetHook(boil.BeforeInsertHook, jetBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jetBeforeInsertHooks = []JetHook{}

	AddJetHook(boil.AfterInsertHook, jetAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jetAfterInsertHooks = []JetHook{}

	AddJetHook(boil.AfterSelectHook, jetAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jetAfterSelectHooks = []JetHook{}

	AddJetHook(boil.BeforeUpdateHook, jetBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jetBeforeUpdateHooks = []JetHook{}

	AddJetHook(boil.AfterUpdateHook, jetAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jetAfterUpdateHooks = []JetHook{}

	AddJetHook(boil.BeforeDeleteHook, jetBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jetBeforeDeleteHooks = []JetHook{}

	AddJetHook(boil.AfterDeleteHook, jetAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jetAfterDeleteHooks = []JetHook{}

	AddJetHook(boil.BeforeUpsertHook, jetBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jetBeforeUpsertHooks = []JetHook{}

	AddJetHook(boil.AfterUpsertHook, jetAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jetAfterUpsertHooks = []JetHook{}
}

func testJetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(jetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJetToOnePilotUsingPilot(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Jet
	var foreign Pilot

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jetDBTypes, false, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pilotDBTypes, false, pilotColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pilot struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PilotID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Pilot().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := JetSlice{&local}
	if err = local.L.LoadPilot(ctx, tx, false, (*[]*Jet)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pilot == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Pilot = nil
	if err = local.L.LoadPilot(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Pilot == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testJetToOneSetOpPilotUsingPilot(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Jet
	var b, c Pilot

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jetDBTypes, false, strmangle.SetComplement(jetPrimaryKeyColumns, jetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pilotDBTypes, false, strmangle.SetComplement(pilotPrimaryKeyColumns, pilotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pilotDBTypes, false, strmangle.SetComplement(pilotPrimaryKeyColumns, pilotColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pilot{&b, &c} {
		err = a.SetPilot(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Pilot != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Jets[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PilotID != x.ID {
			t.Error("foreign key was wrong value", a.PilotID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PilotID))
		reflect.Indirect(reflect.ValueOf(&a.PilotID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PilotID != x.ID {
			t.Error("foreign key was wrong value", a.PilotID, x.ID)
		}
	}
}

func testJetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
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

func testJetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jetDBTypes = map[string]string{`ID`: `bigint`, `PilotID`: `bigint`, `Age`: `smallint`, `Name`: `text`, `Color`: `text`}
	_          = bytes.MinRead
)

func testJetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(jetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(jetAllColumns) == len(jetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jetDBTypes, true, jetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jetAllColumns) == len(jetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Jet{}
	if err = randomize.Struct(seed, o, jetDBTypes, true, jetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jetDBTypes, true, jetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jetAllColumns, jetPrimaryKeyColumns) {
		fields = jetAllColumns
	} else {
		fields = strmangle.SetComplement(
			jetAllColumns,
			jetPrimaryKeyColumns,
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

	slice := JetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJetsUpsert(t *testing.T) {
	t.Parallel()

	if len(jetAllColumns) == len(jetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Jet{}
	if err = randomize.Struct(seed, &o, jetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Jet: %s", err)
	}

	count, err := Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, jetDBTypes, false, jetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Jet struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Jet: %s", err)
	}

	count, err = Jets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
