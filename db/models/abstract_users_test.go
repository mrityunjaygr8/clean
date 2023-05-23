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

func testAbstractUsers(t *testing.T) {
	t.Parallel()

	query := AbstractUsers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAbstractUsersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
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

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAbstractUsersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AbstractUsers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAbstractUsersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AbstractUserSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAbstractUsersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AbstractUserExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AbstractUser exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AbstractUserExists to return true, but got false.")
	}
}

func testAbstractUsersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	abstractUserFound, err := FindAbstractUser(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if abstractUserFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAbstractUsersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AbstractUsers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAbstractUsersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AbstractUsers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAbstractUsersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	abstractUserOne := &AbstractUser{}
	abstractUserTwo := &AbstractUser{}
	if err = randomize.Struct(seed, abstractUserOne, abstractUserDBTypes, false, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}
	if err = randomize.Struct(seed, abstractUserTwo, abstractUserDBTypes, false, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = abstractUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = abstractUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AbstractUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAbstractUsersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	abstractUserOne := &AbstractUser{}
	abstractUserTwo := &AbstractUser{}
	if err = randomize.Struct(seed, abstractUserOne, abstractUserDBTypes, false, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}
	if err = randomize.Struct(seed, abstractUserTwo, abstractUserDBTypes, false, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = abstractUserOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = abstractUserTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func abstractUserBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func abstractUserAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *AbstractUser) error {
	*o = AbstractUser{}
	return nil
}

func testAbstractUsersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &AbstractUser{}
	o := &AbstractUser{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, abstractUserDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AbstractUser object: %s", err)
	}

	AddAbstractUserHook(boil.BeforeInsertHook, abstractUserBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	abstractUserBeforeInsertHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.AfterInsertHook, abstractUserAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	abstractUserAfterInsertHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.AfterSelectHook, abstractUserAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	abstractUserAfterSelectHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.BeforeUpdateHook, abstractUserBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	abstractUserBeforeUpdateHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.AfterUpdateHook, abstractUserAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	abstractUserAfterUpdateHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.BeforeDeleteHook, abstractUserBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	abstractUserBeforeDeleteHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.AfterDeleteHook, abstractUserAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	abstractUserAfterDeleteHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.BeforeUpsertHook, abstractUserBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	abstractUserBeforeUpsertHooks = []AbstractUserHook{}

	AddAbstractUserHook(boil.AfterUpsertHook, abstractUserAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	abstractUserAfterUpsertHooks = []AbstractUserHook{}
}

func testAbstractUsersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAbstractUsersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(abstractUserColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAbstractUsersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
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

func testAbstractUsersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AbstractUserSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAbstractUsersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AbstractUsers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	abstractUserDBTypes = map[string]string{`ID`: `character varying`, `InternalID`: `integer`, `Email`: `character varying`, `Password`: `character varying`}
	_                   = bytes.MinRead
)

func testAbstractUsersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(abstractUserPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(abstractUserAllColumns) == len(abstractUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAbstractUsersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(abstractUserAllColumns) == len(abstractUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AbstractUser{}
	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, abstractUserDBTypes, true, abstractUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(abstractUserAllColumns, abstractUserPrimaryKeyColumns) {
		fields = abstractUserAllColumns
	} else {
		fields = strmangle.SetComplement(
			abstractUserAllColumns,
			abstractUserPrimaryKeyColumns,
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

	slice := AbstractUserSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAbstractUsersUpsert(t *testing.T) {
	t.Parallel()

	if len(abstractUserAllColumns) == len(abstractUserPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AbstractUser{}
	if err = randomize.Struct(seed, &o, abstractUserDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AbstractUser: %s", err)
	}

	count, err := AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, abstractUserDBTypes, false, abstractUserPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AbstractUser struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AbstractUser: %s", err)
	}

	count, err = AbstractUsers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
