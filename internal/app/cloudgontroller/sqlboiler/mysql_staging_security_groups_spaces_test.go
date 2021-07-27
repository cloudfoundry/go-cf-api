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

func testStagingSecurityGroupsSpaces(t *testing.T) {
	t.Parallel()

	query := StagingSecurityGroupsSpaces()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStagingSecurityGroupsSpacesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
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

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStagingSecurityGroupsSpacesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := StagingSecurityGroupsSpaces().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStagingSecurityGroupsSpacesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StagingSecurityGroupsSpaceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStagingSecurityGroupsSpacesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StagingSecurityGroupsSpaceExists(ctx, tx, o.StagingSecurityGroupsSpacesPK)
	if err != nil {
		t.Errorf("Unable to check if StagingSecurityGroupsSpace exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StagingSecurityGroupsSpaceExists to return true, but got false.")
	}
}

func testStagingSecurityGroupsSpacesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stagingSecurityGroupsSpaceFound, err := FindStagingSecurityGroupsSpace(ctx, tx, o.StagingSecurityGroupsSpacesPK)
	if err != nil {
		t.Error(err)
	}

	if stagingSecurityGroupsSpaceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStagingSecurityGroupsSpacesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = StagingSecurityGroupsSpaces().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStagingSecurityGroupsSpacesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := StagingSecurityGroupsSpaces().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStagingSecurityGroupsSpacesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stagingSecurityGroupsSpaceOne := &StagingSecurityGroupsSpace{}
	stagingSecurityGroupsSpaceTwo := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, stagingSecurityGroupsSpaceOne, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}
	if err = randomize.Struct(seed, stagingSecurityGroupsSpaceTwo, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stagingSecurityGroupsSpaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stagingSecurityGroupsSpaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StagingSecurityGroupsSpaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStagingSecurityGroupsSpacesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stagingSecurityGroupsSpaceOne := &StagingSecurityGroupsSpace{}
	stagingSecurityGroupsSpaceTwo := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, stagingSecurityGroupsSpaceOne, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}
	if err = randomize.Struct(seed, stagingSecurityGroupsSpaceTwo, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stagingSecurityGroupsSpaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stagingSecurityGroupsSpaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func stagingSecurityGroupsSpaceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func stagingSecurityGroupsSpaceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *StagingSecurityGroupsSpace) error {
	*o = StagingSecurityGroupsSpace{}
	return nil
}

func testStagingSecurityGroupsSpacesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &StagingSecurityGroupsSpace{}
	o := &StagingSecurityGroupsSpace{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace object: %s", err)
	}

	AddStagingSecurityGroupsSpaceHook(boil.BeforeInsertHook, stagingSecurityGroupsSpaceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceBeforeInsertHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.AfterInsertHook, stagingSecurityGroupsSpaceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceAfterInsertHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.AfterSelectHook, stagingSecurityGroupsSpaceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceAfterSelectHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.BeforeUpdateHook, stagingSecurityGroupsSpaceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceBeforeUpdateHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.AfterUpdateHook, stagingSecurityGroupsSpaceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceAfterUpdateHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.BeforeDeleteHook, stagingSecurityGroupsSpaceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceBeforeDeleteHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.AfterDeleteHook, stagingSecurityGroupsSpaceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceAfterDeleteHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.BeforeUpsertHook, stagingSecurityGroupsSpaceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceBeforeUpsertHooks = []StagingSecurityGroupsSpaceHook{}

	AddStagingSecurityGroupsSpaceHook(boil.AfterUpsertHook, stagingSecurityGroupsSpaceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stagingSecurityGroupsSpaceAfterUpsertHooks = []StagingSecurityGroupsSpaceHook{}
}

func testStagingSecurityGroupsSpacesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStagingSecurityGroupsSpacesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stagingSecurityGroupsSpaceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStagingSecurityGroupsSpaceToOneSecurityGroupUsingStagingSecurityGroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StagingSecurityGroupsSpace
	var foreign SecurityGroup

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, securityGroupDBTypes, false, securityGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroup struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StagingSecurityGroupID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.StagingSecurityGroup().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StagingSecurityGroupsSpaceSlice{&local}
	if err = local.L.LoadStagingSecurityGroup(ctx, tx, false, (*[]*StagingSecurityGroupsSpace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.StagingSecurityGroup == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StagingSecurityGroup = nil
	if err = local.L.LoadStagingSecurityGroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.StagingSecurityGroup == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStagingSecurityGroupsSpaceToOneSpaceUsingStagingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local StagingSecurityGroupsSpace
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Space struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StagingSpaceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.StagingSpace().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StagingSecurityGroupsSpaceSlice{&local}
	if err = local.L.LoadStagingSpace(ctx, tx, false, (*[]*StagingSecurityGroupsSpace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.StagingSpace == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.StagingSpace = nil
	if err = local.L.LoadStagingSpace(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.StagingSpace == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStagingSecurityGroupsSpaceToOneSetOpSecurityGroupUsingStagingSecurityGroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StagingSecurityGroupsSpace
	var b, c SecurityGroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stagingSecurityGroupsSpaceDBTypes, false, strmangle.SetComplement(stagingSecurityGroupsSpacePrimaryKeyColumns, stagingSecurityGroupsSpaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, securityGroupDBTypes, false, strmangle.SetComplement(securityGroupPrimaryKeyColumns, securityGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, securityGroupDBTypes, false, strmangle.SetComplement(securityGroupPrimaryKeyColumns, securityGroupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SecurityGroup{&b, &c} {
		err = a.SetStagingSecurityGroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StagingSecurityGroup != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StagingSecurityGroupStagingSecurityGroupsSpaces[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StagingSecurityGroupID != x.ID {
			t.Error("foreign key was wrong value", a.StagingSecurityGroupID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StagingSecurityGroupID))
		reflect.Indirect(reflect.ValueOf(&a.StagingSecurityGroupID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StagingSecurityGroupID != x.ID {
			t.Error("foreign key was wrong value", a.StagingSecurityGroupID, x.ID)
		}
	}
}
func testStagingSecurityGroupsSpaceToOneSetOpSpaceUsingStagingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a StagingSecurityGroupsSpace
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stagingSecurityGroupsSpaceDBTypes, false, strmangle.SetComplement(stagingSecurityGroupsSpacePrimaryKeyColumns, stagingSecurityGroupsSpaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, spaceDBTypes, false, strmangle.SetComplement(spacePrimaryKeyColumns, spaceColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Space{&b, &c} {
		err = a.SetStagingSpace(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.StagingSpace != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.StagingSpaceStagingSecurityGroupsSpaces[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StagingSpaceID != x.ID {
			t.Error("foreign key was wrong value", a.StagingSpaceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StagingSpaceID))
		reflect.Indirect(reflect.ValueOf(&a.StagingSpaceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StagingSpaceID != x.ID {
			t.Error("foreign key was wrong value", a.StagingSpaceID, x.ID)
		}
	}
}

func testStagingSecurityGroupsSpacesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
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

func testStagingSecurityGroupsSpacesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StagingSecurityGroupsSpaceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStagingSecurityGroupsSpacesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := StagingSecurityGroupsSpaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stagingSecurityGroupsSpaceDBTypes = map[string]string{`StagingSecurityGroupID`: `int`, `StagingSpaceID`: `int`, `StagingSecurityGroupsSpacesPK`: `int`}
	_                                 = bytes.MinRead
)

func testStagingSecurityGroupsSpacesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stagingSecurityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stagingSecurityGroupsSpaceAllColumns) == len(stagingSecurityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStagingSecurityGroupsSpacesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stagingSecurityGroupsSpaceAllColumns) == len(stagingSecurityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stagingSecurityGroupsSpaceDBTypes, true, stagingSecurityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stagingSecurityGroupsSpaceAllColumns, stagingSecurityGroupsSpacePrimaryKeyColumns) {
		fields = stagingSecurityGroupsSpaceAllColumns
	} else {
		fields = strmangle.SetComplement(
			stagingSecurityGroupsSpaceAllColumns,
			stagingSecurityGroupsSpacePrimaryKeyColumns,
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

	slice := StagingSecurityGroupsSpaceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStagingSecurityGroupsSpacesUpsert(t *testing.T) {
	t.Parallel()

	if len(stagingSecurityGroupsSpaceAllColumns) == len(stagingSecurityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLStagingSecurityGroupsSpaceUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := StagingSecurityGroupsSpace{}
	if err = randomize.Struct(seed, &o, stagingSecurityGroupsSpaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StagingSecurityGroupsSpace: %s", err)
	}

	count, err := StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stagingSecurityGroupsSpaceDBTypes, false, stagingSecurityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize StagingSecurityGroupsSpace struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert StagingSecurityGroupsSpace: %s", err)
	}

	count, err = StagingSecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
