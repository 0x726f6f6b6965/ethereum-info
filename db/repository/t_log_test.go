// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

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

func testTLogs(t *testing.T) {
	t.Parallel()

	query := TLogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTLogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
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

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTLogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TLogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTLogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TLogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTLogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TLogExists(ctx, tx, o.TXHash, o.Index)
	if err != nil {
		t.Errorf("Unable to check if TLog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TLogExists to return true, but got false.")
	}
}

func testTLogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	tLogFound, err := FindTLog(ctx, tx, o.TXHash, o.Index)
	if err != nil {
		t.Error(err)
	}

	if tLogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTLogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TLogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTLogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TLogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTLogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tLogOne := &TLog{}
	tLogTwo := &TLog{}
	if err = randomize.Struct(seed, tLogOne, tLogDBTypes, false, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}
	if err = randomize.Struct(seed, tLogTwo, tLogDBTypes, false, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTLogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tLogOne := &TLog{}
	tLogTwo := &TLog{}
	if err = randomize.Struct(seed, tLogOne, tLogDBTypes, false, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}
	if err = randomize.Struct(seed, tLogTwo, tLogDBTypes, false, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = tLogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = tLogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tLogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func tLogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TLog) error {
	*o = TLog{}
	return nil
}

func testTLogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TLog{}
	o := &TLog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tLogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TLog object: %s", err)
	}

	AddTLogHook(boil.BeforeInsertHook, tLogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tLogBeforeInsertHooks = []TLogHook{}

	AddTLogHook(boil.AfterInsertHook, tLogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tLogAfterInsertHooks = []TLogHook{}

	AddTLogHook(boil.AfterSelectHook, tLogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tLogAfterSelectHooks = []TLogHook{}

	AddTLogHook(boil.BeforeUpdateHook, tLogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tLogBeforeUpdateHooks = []TLogHook{}

	AddTLogHook(boil.AfterUpdateHook, tLogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tLogAfterUpdateHooks = []TLogHook{}

	AddTLogHook(boil.BeforeDeleteHook, tLogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tLogBeforeDeleteHooks = []TLogHook{}

	AddTLogHook(boil.AfterDeleteHook, tLogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tLogAfterDeleteHooks = []TLogHook{}

	AddTLogHook(boil.BeforeUpsertHook, tLogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tLogBeforeUpsertHooks = []TLogHook{}

	AddTLogHook(boil.AfterUpsertHook, tLogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tLogAfterUpsertHooks = []TLogHook{}
}

func testTLogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTLogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(tLogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTLogToOneTTransactionUsingTXHashTTransaction(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TLog
	var foreign TTransaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, tLogDBTypes, false, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tTransactionDBTypes, false, tTransactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TTransaction struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TXHash = foreign.TXHash
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.TXHashTTransaction().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.TXHash != foreign.TXHash {
		t.Errorf("want: %v, got %v", foreign.TXHash, check.TXHash)
	}

	ranAfterSelectHook := false
	AddTTransactionHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *TTransaction) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TLogSlice{&local}
	if err = local.L.LoadTXHashTTransaction(ctx, tx, false, (*[]*TLog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TXHashTTransaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.TXHashTTransaction = nil
	if err = local.L.LoadTXHashTTransaction(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.TXHashTTransaction == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTLogToOneSetOpTTransactionUsingTXHashTTransaction(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TLog
	var b, c TTransaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tLogDBTypes, false, strmangle.SetComplement(tLogPrimaryKeyColumns, tLogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tTransactionDBTypes, false, strmangle.SetComplement(tTransactionPrimaryKeyColumns, tTransactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tTransactionDBTypes, false, strmangle.SetComplement(tTransactionPrimaryKeyColumns, tTransactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*TTransaction{&b, &c} {
		err = a.SetTXHashTTransaction(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.TXHashTTransaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TXHashTLogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TXHash != x.TXHash {
			t.Error("foreign key was wrong value", a.TXHash)
		}

		if exists, err := TLogExists(ctx, tx, a.TXHash, a.Index); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testTLogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
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

func testTLogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TLogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTLogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TLogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tLogDBTypes = map[string]string{`TXHash`: `character varying`, `Index`: `bigint`, `Data`: `character varying`, `CreateTime`: `timestamp without time zone`, `UpdateTime`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testTLogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(tLogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(tLogAllColumns) == len(tLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTLogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tLogAllColumns) == len(tLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TLog{}
	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, tLogDBTypes, true, tLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tLogAllColumns, tLogPrimaryKeyColumns) {
		fields = tLogAllColumns
	} else {
		fields = strmangle.SetComplement(
			tLogAllColumns,
			tLogPrimaryKeyColumns,
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

	slice := TLogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTLogsUpsert(t *testing.T) {
	t.Parallel()

	if len(tLogAllColumns) == len(tLogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TLog{}
	if err = randomize.Struct(seed, &o, tLogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TLog: %s", err)
	}

	count, err := TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, tLogDBTypes, false, tLogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TLog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TLog: %s", err)
	}

	count, err = TLogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
