// +build mysql_integration
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

func testSecurityGroupsSpaces(t *testing.T) {
	t.Parallel()

	query := SecurityGroupsSpaces()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSecurityGroupsSpacesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
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

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSecurityGroupsSpacesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SecurityGroupsSpaces().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSecurityGroupsSpacesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SecurityGroupsSpaceSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSecurityGroupsSpacesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SecurityGroupsSpaceExists(ctx, tx, o.SecurityGroupsSpacesPK)
	if err != nil {
		t.Errorf("Unable to check if SecurityGroupsSpace exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SecurityGroupsSpaceExists to return true, but got false.")
	}
}

func testSecurityGroupsSpacesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	securityGroupsSpaceFound, err := FindSecurityGroupsSpace(ctx, tx, o.SecurityGroupsSpacesPK)
	if err != nil {
		t.Error(err)
	}

	if securityGroupsSpaceFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSecurityGroupsSpacesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SecurityGroupsSpaces().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSecurityGroupsSpacesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SecurityGroupsSpaces().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSecurityGroupsSpacesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	securityGroupsSpaceOne := &SecurityGroupsSpace{}
	securityGroupsSpaceTwo := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, securityGroupsSpaceOne, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}
	if err = randomize.Struct(seed, securityGroupsSpaceTwo, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = securityGroupsSpaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = securityGroupsSpaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SecurityGroupsSpaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSecurityGroupsSpacesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	securityGroupsSpaceOne := &SecurityGroupsSpace{}
	securityGroupsSpaceTwo := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, securityGroupsSpaceOne, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}
	if err = randomize.Struct(seed, securityGroupsSpaceTwo, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = securityGroupsSpaceOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = securityGroupsSpaceTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func securityGroupsSpaceBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func securityGroupsSpaceAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *SecurityGroupsSpace) error {
	*o = SecurityGroupsSpace{}
	return nil
}

func testSecurityGroupsSpacesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &SecurityGroupsSpace{}
	o := &SecurityGroupsSpace{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace object: %s", err)
	}

	AddSecurityGroupsSpaceHook(boil.BeforeInsertHook, securityGroupsSpaceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceBeforeInsertHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.AfterInsertHook, securityGroupsSpaceAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceAfterInsertHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.AfterSelectHook, securityGroupsSpaceAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceAfterSelectHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.BeforeUpdateHook, securityGroupsSpaceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceBeforeUpdateHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.AfterUpdateHook, securityGroupsSpaceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceAfterUpdateHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.BeforeDeleteHook, securityGroupsSpaceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceBeforeDeleteHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.AfterDeleteHook, securityGroupsSpaceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceAfterDeleteHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.BeforeUpsertHook, securityGroupsSpaceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceBeforeUpsertHooks = []SecurityGroupsSpaceHook{}

	AddSecurityGroupsSpaceHook(boil.AfterUpsertHook, securityGroupsSpaceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	securityGroupsSpaceAfterUpsertHooks = []SecurityGroupsSpaceHook{}
}

func testSecurityGroupsSpacesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSecurityGroupsSpacesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(securityGroupsSpaceColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSecurityGroupsSpaceToOneSecurityGroupUsingSecurityGroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SecurityGroupsSpace
	var foreign SecurityGroup

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, securityGroupDBTypes, false, securityGroupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroup struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SecurityGroupID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SecurityGroup().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SecurityGroupsSpaceSlice{&local}
	if err = local.L.LoadSecurityGroup(ctx, tx, false, (*[]*SecurityGroupsSpace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SecurityGroup == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SecurityGroup = nil
	if err = local.L.LoadSecurityGroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SecurityGroup == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSecurityGroupsSpaceToOneSpaceUsingSpace(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local SecurityGroupsSpace
	var foreign Space

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, securityGroupsSpaceDBTypes, false, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, spaceDBTypes, false, spaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Space struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.SpaceID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Space().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SecurityGroupsSpaceSlice{&local}
	if err = local.L.LoadSpace(ctx, tx, false, (*[]*SecurityGroupsSpace)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Space == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Space = nil
	if err = local.L.LoadSpace(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Space == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSecurityGroupsSpaceToOneSetOpSecurityGroupUsingSecurityGroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SecurityGroupsSpace
	var b, c SecurityGroup

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, securityGroupsSpaceDBTypes, false, strmangle.SetComplement(securityGroupsSpacePrimaryKeyColumns, securityGroupsSpaceColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSecurityGroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SecurityGroup != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SecurityGroupsSpaces[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SecurityGroupID != x.ID {
			t.Error("foreign key was wrong value", a.SecurityGroupID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SecurityGroupID))
		reflect.Indirect(reflect.ValueOf(&a.SecurityGroupID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SecurityGroupID != x.ID {
			t.Error("foreign key was wrong value", a.SecurityGroupID, x.ID)
		}
	}
}
func testSecurityGroupsSpaceToOneSetOpSpaceUsingSpace(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a SecurityGroupsSpace
	var b, c Space

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, securityGroupsSpaceDBTypes, false, strmangle.SetComplement(securityGroupsSpacePrimaryKeyColumns, securityGroupsSpaceColumnsWithoutDefault)...); err != nil {
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
		err = a.SetSpace(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Space != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SecurityGroupsSpaces[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SpaceID != x.ID {
			t.Error("foreign key was wrong value", a.SpaceID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SpaceID))
		reflect.Indirect(reflect.ValueOf(&a.SpaceID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SpaceID != x.ID {
			t.Error("foreign key was wrong value", a.SpaceID, x.ID)
		}
	}
}

func testSecurityGroupsSpacesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
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

func testSecurityGroupsSpacesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SecurityGroupsSpaceSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSecurityGroupsSpacesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SecurityGroupsSpaces().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	securityGroupsSpaceDBTypes = map[string]string{`SecurityGroupID`: `int`, `SpaceID`: `int`, `SecurityGroupsSpacesPK`: `int`}
	_                          = bytes.MinRead
)

func testSecurityGroupsSpacesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(securityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(securityGroupsSpaceAllColumns) == len(securityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSecurityGroupsSpacesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(securityGroupsSpaceAllColumns) == len(securityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SecurityGroupsSpace{}
	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpaceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, securityGroupsSpaceDBTypes, true, securityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(securityGroupsSpaceAllColumns, securityGroupsSpacePrimaryKeyColumns) {
		fields = securityGroupsSpaceAllColumns
	} else {
		fields = strmangle.SetComplement(
			securityGroupsSpaceAllColumns,
			securityGroupsSpacePrimaryKeyColumns,
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

	slice := SecurityGroupsSpaceSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSecurityGroupsSpacesUpsert(t *testing.T) {
	t.Parallel()

	if len(securityGroupsSpaceAllColumns) == len(securityGroupsSpacePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLSecurityGroupsSpaceUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SecurityGroupsSpace{}
	if err = randomize.Struct(seed, &o, securityGroupsSpaceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SecurityGroupsSpace: %s", err)
	}

	count, err := SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, securityGroupsSpaceDBTypes, false, securityGroupsSpacePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SecurityGroupsSpace struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SecurityGroupsSpace: %s", err)
	}

	count, err = SecurityGroupsSpaces().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
