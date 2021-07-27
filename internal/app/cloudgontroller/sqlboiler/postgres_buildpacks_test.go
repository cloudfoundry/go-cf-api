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

func testBuildpacks(t *testing.T) {
	t.Parallel()

	query := Buildpacks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBuildpacksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
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

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpacksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Buildpacks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpacksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBuildpacksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BuildpackExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Buildpack exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BuildpackExists to return true, but got false.")
	}
}

func testBuildpacksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	buildpackFound, err := FindBuildpack(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if buildpackFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBuildpacksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Buildpacks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBuildpacksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Buildpacks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBuildpacksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	buildpackOne := &Buildpack{}
	buildpackTwo := &Buildpack{}
	if err = randomize.Struct(seed, buildpackOne, buildpackDBTypes, false, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackTwo, buildpackDBTypes, false, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Buildpacks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBuildpacksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	buildpackOne := &Buildpack{}
	buildpackTwo := &Buildpack{}
	if err = randomize.Struct(seed, buildpackOne, buildpackDBTypes, false, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}
	if err = randomize.Struct(seed, buildpackTwo, buildpackDBTypes, false, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = buildpackOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = buildpackTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func buildpackBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func buildpackAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Buildpack) error {
	*o = Buildpack{}
	return nil
}

func testBuildpacksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Buildpack{}
	o := &Buildpack{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, buildpackDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Buildpack object: %s", err)
	}

	AddBuildpackHook(boil.BeforeInsertHook, buildpackBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	buildpackBeforeInsertHooks = []BuildpackHook{}

	AddBuildpackHook(boil.AfterInsertHook, buildpackAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	buildpackAfterInsertHooks = []BuildpackHook{}

	AddBuildpackHook(boil.AfterSelectHook, buildpackAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	buildpackAfterSelectHooks = []BuildpackHook{}

	AddBuildpackHook(boil.BeforeUpdateHook, buildpackBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackBeforeUpdateHooks = []BuildpackHook{}

	AddBuildpackHook(boil.AfterUpdateHook, buildpackAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	buildpackAfterUpdateHooks = []BuildpackHook{}

	AddBuildpackHook(boil.BeforeDeleteHook, buildpackBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackBeforeDeleteHooks = []BuildpackHook{}

	AddBuildpackHook(boil.AfterDeleteHook, buildpackAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	buildpackAfterDeleteHooks = []BuildpackHook{}

	AddBuildpackHook(boil.BeforeUpsertHook, buildpackBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackBeforeUpsertHooks = []BuildpackHook{}

	AddBuildpackHook(boil.AfterUpsertHook, buildpackAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	buildpackAfterUpsertHooks = []BuildpackHook{}
}

func testBuildpacksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpacksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(buildpackColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBuildpackToManyResourceBuildpackAnnotations(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c BuildpackAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildpackAnnotationDBTypes, false, buildpackAnnotationColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ResourceGUID, a.GUID)
	queries.Assign(&c.ResourceGUID, a.GUID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ResourceBuildpackAnnotations().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ResourceGUID, b.ResourceGUID) {
			bFound = true
		}
		if queries.Equal(v.ResourceGUID, c.ResourceGUID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BuildpackSlice{&a}
	if err = a.L.LoadResourceBuildpackAnnotations(ctx, tx, false, (*[]*Buildpack)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceBuildpackAnnotations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ResourceBuildpackAnnotations = nil
	if err = a.L.LoadResourceBuildpackAnnotations(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceBuildpackAnnotations); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBuildpackToManyResourceBuildpackLabels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c BuildpackLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, buildpackLabelDBTypes, false, buildpackLabelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, buildpackLabelDBTypes, false, buildpackLabelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ResourceGUID, a.GUID)
	queries.Assign(&c.ResourceGUID, a.GUID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ResourceBuildpackLabels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ResourceGUID, b.ResourceGUID) {
			bFound = true
		}
		if queries.Equal(v.ResourceGUID, c.ResourceGUID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BuildpackSlice{&a}
	if err = a.L.LoadResourceBuildpackLabels(ctx, tx, false, (*[]*Buildpack)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceBuildpackLabels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ResourceBuildpackLabels = nil
	if err = a.L.LoadResourceBuildpackLabels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ResourceBuildpackLabels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBuildpackToManyAddOpResourceBuildpackAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackAnnotationDBTypes, false, strmangle.SetComplement(buildpackAnnotationPrimaryKeyColumns, buildpackAnnotationColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BuildpackAnnotation{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddResourceBuildpackAnnotations(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.GUID, first.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, first.ResourceGUID)
		}
		if !queries.Equal(a.GUID, second.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, second.ResourceGUID)
		}

		if first.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ResourceBuildpackAnnotations[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ResourceBuildpackAnnotations[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ResourceBuildpackAnnotations().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBuildpackToManySetOpResourceBuildpackAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackAnnotationDBTypes, false, strmangle.SetComplement(buildpackAnnotationPrimaryKeyColumns, buildpackAnnotationColumnsWithoutDefault)...); err != nil {
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

	err = a.SetResourceBuildpackAnnotations(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceBuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetResourceBuildpackAnnotations(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceBuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.GUID, d.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, d.ResourceGUID)
	}
	if !queries.Equal(a.GUID, e.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, e.ResourceGUID)
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ResourceBuildpackAnnotations[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ResourceBuildpackAnnotations[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBuildpackToManyRemoveOpResourceBuildpackAnnotations(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackAnnotation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackAnnotation{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackAnnotationDBTypes, false, strmangle.SetComplement(buildpackAnnotationPrimaryKeyColumns, buildpackAnnotationColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddResourceBuildpackAnnotations(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceBuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveResourceBuildpackAnnotations(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceBuildpackAnnotations().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ResourceBuildpackAnnotations) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ResourceBuildpackAnnotations[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ResourceBuildpackAnnotations[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBuildpackToManyAddOpResourceBuildpackLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLabelDBTypes, false, strmangle.SetComplement(buildpackLabelPrimaryKeyColumns, buildpackLabelColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*BuildpackLabel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddResourceBuildpackLabels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.GUID, first.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, first.ResourceGUID)
		}
		if !queries.Equal(a.GUID, second.ResourceGUID) {
			t.Error("foreign key was wrong value", a.GUID, second.ResourceGUID)
		}

		if first.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Resource != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ResourceBuildpackLabels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ResourceBuildpackLabels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ResourceBuildpackLabels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testBuildpackToManySetOpResourceBuildpackLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLabelDBTypes, false, strmangle.SetComplement(buildpackLabelPrimaryKeyColumns, buildpackLabelColumnsWithoutDefault)...); err != nil {
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

	err = a.SetResourceBuildpackLabels(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceBuildpackLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetResourceBuildpackLabels(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceBuildpackLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.GUID, d.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, d.ResourceGUID)
	}
	if !queries.Equal(a.GUID, e.ResourceGUID) {
		t.Error("foreign key was wrong value", a.GUID, e.ResourceGUID)
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Resource != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ResourceBuildpackLabels[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ResourceBuildpackLabels[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testBuildpackToManyRemoveOpResourceBuildpackLabels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Buildpack
	var b, c, d, e BuildpackLabel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, buildpackDBTypes, false, strmangle.SetComplement(buildpackPrimaryKeyColumns, buildpackColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*BuildpackLabel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, buildpackLabelDBTypes, false, strmangle.SetComplement(buildpackLabelPrimaryKeyColumns, buildpackLabelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddResourceBuildpackLabels(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ResourceBuildpackLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveResourceBuildpackLabels(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ResourceBuildpackLabels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ResourceGUID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ResourceGUID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Resource != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Resource != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ResourceBuildpackLabels) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ResourceBuildpackLabels[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ResourceBuildpackLabels[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testBuildpacksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
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

func testBuildpacksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BuildpackSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBuildpacksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Buildpacks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	buildpackDBTypes = map[string]string{`ID`: `integer`, `GUID`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `Name`: `text`, `Key`: `text`, `Position`: `integer`, `Enabled`: `boolean`, `Locked`: `boolean`, `Filename`: `text`, `Sha256Checksum`: `text`, `Stack`: `character varying`}
	_                = bytes.MinRead
)

func testBuildpacksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(buildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(buildpackAllColumns) == len(buildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBuildpacksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(buildpackAllColumns) == len(buildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Buildpack{}
	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, buildpackDBTypes, true, buildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(buildpackAllColumns, buildpackPrimaryKeyColumns) {
		fields = buildpackAllColumns
	} else {
		fields = strmangle.SetComplement(
			buildpackAllColumns,
			buildpackPrimaryKeyColumns,
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

	slice := BuildpackSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBuildpacksUpsert(t *testing.T) {
	t.Parallel()

	if len(buildpackAllColumns) == len(buildpackPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Buildpack{}
	if err = randomize.Struct(seed, &o, buildpackDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Buildpack: %s", err)
	}

	count, err := Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, buildpackDBTypes, false, buildpackPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Buildpack struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Buildpack: %s", err)
	}

	count, err = Buildpacks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}