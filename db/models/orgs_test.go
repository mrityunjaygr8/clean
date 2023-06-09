// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

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

func testOrgs(t *testing.T) {
	t.Parallel()

	query := Orgs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrgsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
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

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrgsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Orgs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrgsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrgSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrgsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrgExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Org exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrgExists to return true, but got false.")
	}
}

func testOrgsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orgFound, err := FindOrg(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orgFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrgsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Orgs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrgsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Orgs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrgsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orgOne := &Org{}
	orgTwo := &Org{}
	if err = randomize.Struct(seed, orgOne, orgDBTypes, false, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}
	if err = randomize.Struct(seed, orgTwo, orgDBTypes, false, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orgOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orgTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orgs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrgsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orgOne := &Org{}
	orgTwo := &Org{}
	if err = randomize.Struct(seed, orgOne, orgDBTypes, false, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}
	if err = randomize.Struct(seed, orgTwo, orgDBTypes, false, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orgOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orgTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func orgBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func orgAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Org) error {
	*o = Org{}
	return nil
}

func testOrgsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Org{}
	o := &Org{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, orgDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Org object: %s", err)
	}

	AddOrgHook(boil.BeforeInsertHook, orgBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	orgBeforeInsertHooks = []OrgHook{}

	AddOrgHook(boil.AfterInsertHook, orgAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	orgAfterInsertHooks = []OrgHook{}

	AddOrgHook(boil.AfterSelectHook, orgAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	orgAfterSelectHooks = []OrgHook{}

	AddOrgHook(boil.BeforeUpdateHook, orgBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	orgBeforeUpdateHooks = []OrgHook{}

	AddOrgHook(boil.AfterUpdateHook, orgAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	orgAfterUpdateHooks = []OrgHook{}

	AddOrgHook(boil.BeforeDeleteHook, orgBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	orgBeforeDeleteHooks = []OrgHook{}

	AddOrgHook(boil.AfterDeleteHook, orgAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	orgAfterDeleteHooks = []OrgHook{}

	AddOrgHook(boil.BeforeUpsertHook, orgBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	orgBeforeUpsertHooks = []OrgHook{}

	AddOrgHook(boil.AfterUpsertHook, orgAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	orgAfterUpsertHooks = []OrgHook{}
}

func testOrgsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrgsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orgColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrgToManyOrgUsers(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Org
	var b, c OrgUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, orgUserDBTypes, false, orgUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orgUserDBTypes, false, orgUserColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.OrgID = a.InternalID
	c.OrgID = a.InternalID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OrgUsers().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.OrgID == b.OrgID {
			bFound = true
		}
		if v.OrgID == c.OrgID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrgSlice{&a}
	if err = a.L.LoadOrgUsers(ctx, tx, false, (*[]*Org)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrgUsers = nil
	if err = a.L.LoadOrgUsers(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrgUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOrgToManyAddOpOrgUsers(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Org
	var b, c, d, e OrgUser

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orgDBTypes, false, strmangle.SetComplement(orgPrimaryKeyColumns, orgColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OrgUser{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orgUserDBTypes, false, strmangle.SetComplement(orgUserPrimaryKeyColumns, orgUserColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*OrgUser{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrgUsers(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.InternalID != first.OrgID {
			t.Error("foreign key was wrong value", a.InternalID, first.OrgID)
		}
		if a.InternalID != second.OrgID {
			t.Error("foreign key was wrong value", a.InternalID, second.OrgID)
		}

		if first.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Org != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrgUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrgUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrgUsers().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOrgsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
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

func testOrgsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrgSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrgsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orgs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orgDBTypes = map[string]string{`InternalID`: `integer`, `ID`: `character varying`, `Name`: `character varying`}
	_          = bytes.MinRead
)

func testOrgsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orgPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orgAllColumns) == len(orgPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orgDBTypes, true, orgPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrgsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orgAllColumns) == len(orgPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Org{}
	if err = randomize.Struct(seed, o, orgDBTypes, true, orgColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orgDBTypes, true, orgPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orgAllColumns, orgPrimaryKeyColumns) {
		fields = orgAllColumns
	} else {
		fields = strmangle.SetComplement(
			orgAllColumns,
			orgPrimaryKeyColumns,
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

	slice := OrgSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrgsUpsert(t *testing.T) {
	t.Parallel()

	if len(orgAllColumns) == len(orgPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Org{}
	if err = randomize.Struct(seed, &o, orgDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Org: %s", err)
	}

	count, err := Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, orgDBTypes, false, orgPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Org struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Org: %s", err)
	}

	count, err = Orgs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
