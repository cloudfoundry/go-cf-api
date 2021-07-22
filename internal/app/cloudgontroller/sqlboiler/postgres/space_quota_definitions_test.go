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

func testSpaceQuotaDefinitions(t *testing.T) {
	t.Parallel()

	query := SpaceQuotaDefinitions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSpaceQuotaDefinitionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
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

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceQuotaDefinitionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SpaceQuotaDefinitions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceQuotaDefinitionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpaceQuotaDefinitionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSpaceQuotaDefinitionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SpaceQuotaDefinitionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SpaceQuotaDefinition exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SpaceQuotaDefinitionExists to return true, but got false.")
	}
}

func testSpaceQuotaDefinitionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	spaceQuotaDefinitionFound, err := FindSpaceQuotaDefinition(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if spaceQuotaDefinitionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSpaceQuotaDefinitionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SpaceQuotaDefinitions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSpaceQuotaDefinitionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SpaceQuotaDefinitions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSpaceQuotaDefinitionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	spaceQuotaDefinitionOne := &SpaceQuotaDefinition{}
	spaceQuotaDefinitionTwo := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, spaceQuotaDefinitionOne, spaceQuotaDefinitionDBTypes, false, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}
	if err = randomize.Struct(seed, spaceQuotaDefinitionTwo, spaceQuotaDefinitionDBTypes, false, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spaceQuotaDefinitionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spaceQuotaDefinitionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpaceQuotaDefinitions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSpaceQuotaDefinitionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	spaceQuotaDefinitionOne := &SpaceQuotaDefinition{}
	spaceQuotaDefinitionTwo := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, spaceQuotaDefinitionOne, spaceQuotaDefinitionDBTypes, false, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}
	if err = randomize.Struct(seed, spaceQuotaDefinitionTwo, spaceQuotaDefinitionDBTypes, false, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = spaceQuotaDefinitionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = spaceQuotaDefinitionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func spaceQuotaDefinitionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func spaceQuotaDefinitionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SpaceQuotaDefinition) error {
	*o = SpaceQuotaDefinition{}
	return nil
}

func testSpaceQuotaDefinitionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SpaceQuotaDefinition{}
	o := &SpaceQuotaDefinition{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition object: %s", err)
	}

	AddSpaceQuotaDefinitionHook(boil.BeforeInsertHook, spaceQuotaDefinitionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionBeforeInsertHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.AfterInsertHook, spaceQuotaDefinitionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionAfterInsertHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.AfterSelectHook, spaceQuotaDefinitionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionAfterSelectHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.BeforeUpdateHook, spaceQuotaDefinitionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionBeforeUpdateHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.AfterUpdateHook, spaceQuotaDefinitionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionAfterUpdateHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.BeforeDeleteHook, spaceQuotaDefinitionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionBeforeDeleteHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.AfterDeleteHook, spaceQuotaDefinitionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionAfterDeleteHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.BeforeUpsertHook, spaceQuotaDefinitionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionBeforeUpsertHooks = []SpaceQuotaDefinitionHook{}

	AddSpaceQuotaDefinitionHook(boil.AfterUpsertHook, spaceQuotaDefinitionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	spaceQuotaDefinitionAfterUpsertHooks = []SpaceQuotaDefinitionHook{}
}

func testSpaceQuotaDefinitionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpaceQuotaDefinitionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(spaceQuotaDefinitionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSpaceQuotaDefinitionToManySpaces(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceQuotaDefinition
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.SpaceQuotaDefinitionID, a.ID)
	queries.Assign(&c.SpaceQuotaDefinitionID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Spaces().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.SpaceQuotaDefinitionID, b.SpaceQuotaDefinitionID) {
			bFound = true
		}
		if queries.Equal(v.SpaceQuotaDefinitionID, c.SpaceQuotaDefinitionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := SpaceQuotaDefinitionSlice{&a}
	if err = a.L.LoadSpaces(ctx, tx, false, (*[]*SpaceQuotaDefinition)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Spaces); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Spaces = nil
	if err = a.L.LoadSpaces(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Spaces); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testSpaceQuotaDefinitionToManyAddOpSpaces(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceQuotaDefinition
	var b, c, d, e Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceQuotaDefinitionDBTypes, false, strmangle.SetComplement(spaceQuotaDefinitionPrimaryKeyColumns, spaceQuotaDefinitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Space{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Space{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSpaces(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.SpaceQuotaDefinitionID) {
			t.Error("foreign key was wrong value", a.ID, first.SpaceQuotaDefinitionID)
		}
		if !queries.Equal(a.ID, second.SpaceQuotaDefinitionID) {
			t.Error("foreign key was wrong value", a.ID, second.SpaceQuotaDefinitionID)
		}

		if first.R.SpaceQuotaDefinition != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.SpaceQuotaDefinition != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Spaces[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Spaces[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Spaces().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testSpaceQuotaDefinitionToManySetOpSpaces(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceQuotaDefinition
	var b, c, d, e Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceQuotaDefinitionDBTypes, false, strmangle.SetComplement(spaceQuotaDefinitionPrimaryKeyColumns, spaceQuotaDefinitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Space{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSpaces(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Spaces().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSpaces(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Spaces().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SpaceQuotaDefinitionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SpaceQuotaDefinitionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.SpaceQuotaDefinitionID) {
		t.Error("foreign key was wrong value", a.ID, d.SpaceQuotaDefinitionID)
	}
	if !queries.Equal(a.ID, e.SpaceQuotaDefinitionID) {
		t.Error("foreign key was wrong value", a.ID, e.SpaceQuotaDefinitionID)
	}

	if b.R.SpaceQuotaDefinition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SpaceQuotaDefinition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SpaceQuotaDefinition != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.SpaceQuotaDefinition != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Spaces[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Spaces[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testSpaceQuotaDefinitionToManyRemoveOpSpaces(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SpaceQuotaDefinition
	var b, c, d, e Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, spaceQuotaDefinitionDBTypes, false, strmangle.SetComplement(spaceQuotaDefinitionPrimaryKeyColumns, spaceQuotaDefinitionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Space{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddSpaces(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Spaces().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSpaces(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Spaces().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.SpaceQuotaDefinitionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.SpaceQuotaDefinitionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.SpaceQuotaDefinition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.SpaceQuotaDefinition != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.SpaceQuotaDefinition != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.SpaceQuotaDefinition != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Spaces) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Spaces[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Spaces[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testSpaceQuotaDefinitionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
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

func testSpaceQuotaDefinitionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SpaceQuotaDefinitionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSpaceQuotaDefinitionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SpaceQuotaDefinitions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	spaceQuotaDefinitionDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `Name`: `text`, `NonBasicServicesAllowed`: `boolean`, `TotalServices`: `integer`, `MemoryLimit`: `integer`, `TotalRoutes`: `integer`, `InstanceMemoryLimit`: `integer`, `OrganizationID`: `integer`, `AppInstanceLimit`: `integer`, `AppTaskLimit`: `integer`, `TotalServiceKeys`: `integer`, `TotalReservedRoutePorts`: `integer`}
	_                           = bytes.MinRead
)

func testSpaceQuotaDefinitionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(spaceQuotaDefinitionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(spaceQuotaDefinitionAllColumns) == len(spaceQuotaDefinitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSpaceQuotaDefinitionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(spaceQuotaDefinitionAllColumns) == len(spaceQuotaDefinitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, spaceQuotaDefinitionDBTypes, true, spaceQuotaDefinitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(spaceQuotaDefinitionAllColumns, spaceQuotaDefinitionPrimaryKeyColumns) {
		fields = spaceQuotaDefinitionAllColumns
	} else {
		fields = strmangle.SetComplement(
			spaceQuotaDefinitionAllColumns,
			spaceQuotaDefinitionPrimaryKeyColumns,
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

	slice := SpaceQuotaDefinitionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSpaceQuotaDefinitionsUpsert(t *testing.T) {
	t.Parallel()

	if len(spaceQuotaDefinitionAllColumns) == len(spaceQuotaDefinitionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SpaceQuotaDefinition{}
	if err = randomize.Struct(seed, &o, spaceQuotaDefinitionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpaceQuotaDefinition: %s", err)
	}

	count, err := SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, spaceQuotaDefinitionDBTypes, false, spaceQuotaDefinitionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SpaceQuotaDefinition struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SpaceQuotaDefinition: %s", err)
	}

	count, err = SpaceQuotaDefinitions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
