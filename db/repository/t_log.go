// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// TLog is an object representing the database table.
type TLog struct {
	TXHash     string    `boil:"tx_hash" json:"tx_hash" toml:"tx_hash" yaml:"tx_hash"`
	Index      int64     `boil:"index" json:"index" toml:"index" yaml:"index"`
	Data       string    `boil:"data" json:"data" toml:"data" yaml:"data"`
	CreateTime time.Time `boil:"create_time" json:"create_time" toml:"create_time" yaml:"create_time"`
	UpdateTime time.Time `boil:"update_time" json:"update_time" toml:"update_time" yaml:"update_time"`

	R *tLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TLogColumns = struct {
	TXHash     string
	Index      string
	Data       string
	CreateTime string
	UpdateTime string
}{
	TXHash:     "tx_hash",
	Index:      "index",
	Data:       "data",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

var TLogTableColumns = struct {
	TXHash     string
	Index      string
	Data       string
	CreateTime string
	UpdateTime string
}{
	TXHash:     "t_log.tx_hash",
	Index:      "t_log.index",
	Data:       "t_log.data",
	CreateTime: "t_log.create_time",
	UpdateTime: "t_log.update_time",
}

// Generated where

var TLogWhere = struct {
	TXHash     whereHelperstring
	Index      whereHelperint64
	Data       whereHelperstring
	CreateTime whereHelpertime_Time
	UpdateTime whereHelpertime_Time
}{
	TXHash:     whereHelperstring{field: "\"t_log\".\"tx_hash\""},
	Index:      whereHelperint64{field: "\"t_log\".\"index\""},
	Data:       whereHelperstring{field: "\"t_log\".\"data\""},
	CreateTime: whereHelpertime_Time{field: "\"t_log\".\"create_time\""},
	UpdateTime: whereHelpertime_Time{field: "\"t_log\".\"update_time\""},
}

// TLogRels is where relationship names are stored.
var TLogRels = struct {
	TXHashTTransaction string
}{
	TXHashTTransaction: "TXHashTTransaction",
}

// tLogR is where relationships are stored.
type tLogR struct {
	TXHashTTransaction *TTransaction `boil:"TXHashTTransaction" json:"TXHashTTransaction" toml:"TXHashTTransaction" yaml:"TXHashTTransaction"`
}

// NewStruct creates a new relationship struct
func (*tLogR) NewStruct() *tLogR {
	return &tLogR{}
}

func (r *tLogR) GetTXHashTTransaction() *TTransaction {
	if r == nil {
		return nil
	}
	return r.TXHashTTransaction
}

// tLogL is where Load methods for each relationship are stored.
type tLogL struct{}

var (
	tLogAllColumns            = []string{"tx_hash", "index", "data", "create_time", "update_time"}
	tLogColumnsWithoutDefault = []string{"tx_hash", "index", "data"}
	tLogColumnsWithDefault    = []string{"create_time", "update_time"}
	tLogPrimaryKeyColumns     = []string{"tx_hash", "index"}
	tLogGeneratedColumns      = []string{}
)

type (
	// TLogSlice is an alias for a slice of pointers to TLog.
	// This should almost always be used instead of []TLog.
	TLogSlice []*TLog
	// TLogHook is the signature for custom TLog hook methods
	TLogHook func(context.Context, boil.ContextExecutor, *TLog) error

	tLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tLogType                 = reflect.TypeOf(&TLog{})
	tLogMapping              = queries.MakeStructMapping(tLogType)
	tLogPrimaryKeyMapping, _ = queries.BindMapping(tLogType, tLogMapping, tLogPrimaryKeyColumns)
	tLogInsertCacheMut       sync.RWMutex
	tLogInsertCache          = make(map[string]insertCache)
	tLogUpdateCacheMut       sync.RWMutex
	tLogUpdateCache          = make(map[string]updateCache)
	tLogUpsertCacheMut       sync.RWMutex
	tLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tLogAfterSelectHooks []TLogHook

var tLogBeforeInsertHooks []TLogHook
var tLogAfterInsertHooks []TLogHook

var tLogBeforeUpdateHooks []TLogHook
var tLogAfterUpdateHooks []TLogHook

var tLogBeforeDeleteHooks []TLogHook
var tLogAfterDeleteHooks []TLogHook

var tLogBeforeUpsertHooks []TLogHook
var tLogAfterUpsertHooks []TLogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTLogHook registers your hook function for all future operations.
func AddTLogHook(hookPoint boil.HookPoint, tLogHook TLogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tLogAfterSelectHooks = append(tLogAfterSelectHooks, tLogHook)
	case boil.BeforeInsertHook:
		tLogBeforeInsertHooks = append(tLogBeforeInsertHooks, tLogHook)
	case boil.AfterInsertHook:
		tLogAfterInsertHooks = append(tLogAfterInsertHooks, tLogHook)
	case boil.BeforeUpdateHook:
		tLogBeforeUpdateHooks = append(tLogBeforeUpdateHooks, tLogHook)
	case boil.AfterUpdateHook:
		tLogAfterUpdateHooks = append(tLogAfterUpdateHooks, tLogHook)
	case boil.BeforeDeleteHook:
		tLogBeforeDeleteHooks = append(tLogBeforeDeleteHooks, tLogHook)
	case boil.AfterDeleteHook:
		tLogAfterDeleteHooks = append(tLogAfterDeleteHooks, tLogHook)
	case boil.BeforeUpsertHook:
		tLogBeforeUpsertHooks = append(tLogBeforeUpsertHooks, tLogHook)
	case boil.AfterUpsertHook:
		tLogAfterUpsertHooks = append(tLogAfterUpsertHooks, tLogHook)
	}
}

// OneG returns a single tLog record from the query using the global executor.
func (q tLogQuery) OneG(ctx context.Context) (*TLog, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single tLog record from the query.
func (q tLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TLog, error) {
	o := &TLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for t_log")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all TLog records from the query using the global executor.
func (q tLogQuery) AllG(ctx context.Context) (TLogSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all TLog records from the query.
func (q tLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (TLogSlice, error) {
	var o []*TLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to TLog slice")
	}

	if len(tLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all TLog records in the query using the global executor
func (q tLogQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all TLog records in the query.
func (q tLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count t_log rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q tLogQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q tLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if t_log exists")
	}

	return count > 0, nil
}

// TXHashTTransaction pointed to by the foreign key.
func (o *TLog) TXHashTTransaction(mods ...qm.QueryMod) tTransactionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"tx_hash\" = ?", o.TXHash),
	}

	queryMods = append(queryMods, mods...)

	return TTransactions(queryMods...)
}

// LoadTXHashTTransaction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tLogL) LoadTXHashTTransaction(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTLog interface{}, mods queries.Applicator) error {
	var slice []*TLog
	var object *TLog

	if singular {
		var ok bool
		object, ok = maybeTLog.(*TLog)
		if !ok {
			object = new(TLog)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeTLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeTLog))
			}
		}
	} else {
		s, ok := maybeTLog.(*[]*TLog)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeTLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeTLog))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tLogR{}
		}
		args = append(args, object.TXHash)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tLogR{}
			}

			for _, a := range args {
				if a == obj.TXHash {
					continue Outer
				}
			}

			args = append(args, obj.TXHash)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`t_transaction`),
		qm.WhereIn(`t_transaction.tx_hash in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TTransaction")
	}

	var resultSlice []*TTransaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TTransaction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for t_transaction")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for t_transaction")
	}

	if len(tTransactionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TXHashTTransaction = foreign
		if foreign.R == nil {
			foreign.R = &tTransactionR{}
		}
		foreign.R.TXHashTLogs = append(foreign.R.TXHashTLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TXHash == foreign.TXHash {
				local.R.TXHashTTransaction = foreign
				if foreign.R == nil {
					foreign.R = &tTransactionR{}
				}
				foreign.R.TXHashTLogs = append(foreign.R.TXHashTLogs, local)
				break
			}
		}
	}

	return nil
}

// SetTXHashTTransactionG of the tLog to the related item.
// Sets o.R.TXHashTTransaction to related.
// Adds o to related.R.TXHashTLogs.
// Uses the global database handle.
func (o *TLog) SetTXHashTTransactionG(ctx context.Context, insert bool, related *TTransaction) error {
	return o.SetTXHashTTransaction(ctx, boil.GetContextDB(), insert, related)
}

// SetTXHashTTransaction of the tLog to the related item.
// Sets o.R.TXHashTTransaction to related.
// Adds o to related.R.TXHashTLogs.
func (o *TLog) SetTXHashTTransaction(ctx context.Context, exec boil.ContextExecutor, insert bool, related *TTransaction) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"t_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"tx_hash"}),
		strmangle.WhereClause("\"", "\"", 2, tLogPrimaryKeyColumns),
	)
	values := []interface{}{related.TXHash, o.TXHash, o.Index}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TXHash = related.TXHash
	if o.R == nil {
		o.R = &tLogR{
			TXHashTTransaction: related,
		}
	} else {
		o.R.TXHashTTransaction = related
	}

	if related.R == nil {
		related.R = &tTransactionR{
			TXHashTLogs: TLogSlice{o},
		}
	} else {
		related.R.TXHashTLogs = append(related.R.TXHashTLogs, o)
	}

	return nil
}

// TLogs retrieves all the records using an executor.
func TLogs(mods ...qm.QueryMod) tLogQuery {
	mods = append(mods, qm.From("\"t_log\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"t_log\".*"})
	}

	return tLogQuery{q}
}

// FindTLogG retrieves a single record by ID.
func FindTLogG(ctx context.Context, tXHash string, index int64, selectCols ...string) (*TLog, error) {
	return FindTLog(ctx, boil.GetContextDB(), tXHash, index, selectCols...)
}

// FindTLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTLog(ctx context.Context, exec boil.ContextExecutor, tXHash string, index int64, selectCols ...string) (*TLog, error) {
	tLogObj := &TLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"t_log\" where \"tx_hash\"=$1 AND \"index\"=$2", sel,
	)

	q := queries.Raw(query, tXHash, index)

	err := q.Bind(ctx, exec, tLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: unable to select from t_log")
	}

	if err = tLogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tLogObj, err
	}

	return tLogObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *TLog) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no t_log provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tLogInsertCacheMut.RLock()
	cache, cached := tLogInsertCache[key]
	tLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tLogAllColumns,
			tLogColumnsWithDefault,
			tLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tLogType, tLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tLogType, tLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"t_log\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"t_log\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "repository: unable to insert into t_log")
	}

	if !cached {
		tLogInsertCacheMut.Lock()
		tLogInsertCache[key] = cache
		tLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single TLog record using the global executor.
// See Update for more documentation.
func (o *TLog) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the TLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tLogUpdateCacheMut.RLock()
	cache, cached := tLogUpdateCache[key]
	tLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tLogAllColumns,
			tLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repository: unable to update t_log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"t_log\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, tLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tLogType, tLogMapping, append(wl, tLogPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update t_log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by update for t_log")
	}

	if !cached {
		tLogUpdateCacheMut.Lock()
		tLogUpdateCache[key] = cache
		tLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q tLogQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q tLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all for t_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected for t_log")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TLogSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("repository: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"t_log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, tLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all in tLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected all in update all tLog")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *TLog) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no t_log provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tLogColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tLogUpsertCacheMut.RLock()
	cache, cached := tLogUpsertCache[key]
	tLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tLogAllColumns,
			tLogColumnsWithDefault,
			tLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tLogAllColumns,
			tLogPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("repository: unable to upsert t_log, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(tLogPrimaryKeyColumns))
			copy(conflict, tLogPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"t_log\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(tLogType, tLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tLogType, tLogMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "repository: unable to upsert t_log")
	}

	if !cached {
		tLogUpsertCacheMut.Lock()
		tLogUpsertCache[key] = cache
		tLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single TLog record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *TLog) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single TLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repository: no TLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tLogPrimaryKeyMapping)
	sql := "DELETE FROM \"t_log\" WHERE \"tx_hash\"=$1 AND \"index\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete from t_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by delete for t_log")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q tLogQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q tLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repository: no tLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from t_log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for t_log")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TLogSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"t_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from tLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for t_log")
	}

	if len(tLogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *TLog) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("repository: no TLog provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTLog(ctx, exec, o.TXHash, o.Index)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TLogSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("repository: empty TLogSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"t_log\".* FROM \"t_log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, tLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repository: unable to reload all in TLogSlice")
	}

	*o = slice

	return nil
}

// TLogExistsG checks if the TLog row exists.
func TLogExistsG(ctx context.Context, tXHash string, index int64) (bool, error) {
	return TLogExists(ctx, boil.GetContextDB(), tXHash, index)
}

// TLogExists checks if the TLog row exists.
func TLogExists(ctx context.Context, exec boil.ContextExecutor, tXHash string, index int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"t_log\" where \"tx_hash\"=$1 AND \"index\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tXHash, index)
	}
	row := exec.QueryRowContext(ctx, sql, tXHash, index)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repository: unable to check if t_log exists")
	}

	return exists, nil
}

// Exists checks if the TLog row exists.
func (o *TLog) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TLogExists(ctx, exec, o.TXHash, o.Index)
}
