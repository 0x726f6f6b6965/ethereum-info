package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	dbTrans "github.com/0x726f6f6b6965/ethereum-info/db/repository"
	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	libConfig "github.com/0x726f6f6b6965/ethereum-info/library/config"
	"github.com/0x726f6f6b6965/ethereum-info/library/helper"
	"github.com/0x726f6f6b6965/ethereum-info/monitor-service/internal/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/redis/go-redis/v9"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type MonitorService interface {
	SetGoBlockNum(num int)
	SetGoTransactionNum(num int)
	SetGoLogNum(num int)
	SetGoReceiptsNum(num int)
	SetSaveBlockNum(num int64)
	SetRollBackNum(num int64)
	MonitorBlocks(ctx context.Context)
	ConnectEth(ctx context.Context)
	StopMonitor()
	IsMonitoring() bool
}

type server struct {
	rpcRelay      *client.RPCRelay
	rpcCfg        *libConfig.RpcClientCfg
	redisClient   *redis.Client
	db            *sql.DB
	logger        *zap.Logger
	startBlock    int64
	rollBack      int64
	saveBlockNum  int64
	goBlockNum    int
	goTransNum    int
	goLogNum      int
	goReceiptsNum int
	intreval      time.Duration
	monitoring    uint32
	stop          func()
	done          chan bool
	monitorType   consts.MonitorType
}

func (s *server) SetGoBlockNum(num int) {
	s.goBlockNum = num
}
func (s *server) SetGoTransactionNum(num int) {
	s.goTransNum = num
}
func (s *server) SetGoLogNum(num int) {
	s.goLogNum = num
}
func (s *server) SetGoReceiptsNum(num int) {
	s.goReceiptsNum = num
}
func (s *server) SetSaveBlockNum(num int64) {
	s.saveBlockNum = num
}
func (s *server) SetRollBackNum(num int64) {
	s.rollBack = num
	_ = s.redisClient.Set(context.Background(), "rollback", s.rollBack, -1)
}
func (s *server) MonitorBlocks(ctx context.Context) {
	if atomic.LoadUint32(&s.monitoring) == 1 {
		return
	}

	ticker := time.NewTicker(s.intreval)
	if s.monitorType == consts.FORWARD {
		go s.saveToDB(ctx, ticker)
	} else if s.monitorType == consts.UNSTABLE {
		go s.saveToRedis(ctx, ticker)
	}
	atomic.StoreUint32(&s.monitoring, 1)
	s.stop = ticker.Stop
}

func (s *server) ConnectEth(ctx context.Context) {
	client, err := client.InitRPCRelayByHttp(s.rpcCfg)
	if err != nil {
		return
	}
	s.rpcRelay = client
}

func (s *server) StopMonitor() {
	if atomic.LoadUint32(&s.monitoring) != 1 {
		return
	}

	atomic.StoreUint32(&s.monitoring, 0)
	s.stop()
	s.done <- true
}

func (s *server) IsMonitoring() bool {
	return atomic.LoadUint32(&s.monitoring) == 1
}

func (s *server) queryBlocks(ctx context.Context, start int64, end int64, unstable bool) chan *types.Block {
	blocks := make(chan *types.Block)
	limiter := make(chan struct{}, s.goBlockNum)
	go func() {
		defer close(blocks)
		defer close(limiter)
		var wg sync.WaitGroup
		defer wg.Wait()
		for i := start; i <= end; i++ {
			query := big.NewInt(i)
			// limit goroutine
			limiter <- struct{}{}
			wg.Add(1)
			go func(num *big.Int) {
				defer func() {
					<-limiter
					wg.Done()
				}()
				block, err := s.rpcRelay.GetBlockByNum(ctx, num)
				if err != nil {
					// @TBD: add error message to dead queue
					s.logger.Error("rpc get block error",
						zap.Int64("block_num", num.Int64()), zap.Error(err))
					_ = s.redisClient.Set(ctx, fmt.Sprintf(consts.ErrQueryBlocksKey, num.Int64()),
						err.Error(), helper.GeneralDuration(6, 2, 12, time.Hour))
					return
				}

				if unstable {
					// save block pb format to redis
					pbBlockInfo := helper.ParseBlockToPb(block)
					pbBlockInfo.Stable = false
					bytes, _ := json.Marshal(pbBlockInfo)
					_ = s.redisClient.Set(ctx,
						fmt.Sprintf(consts.BlockDataKey, block.Number().Int64()), bytes,
						helper.GeneralDuration(15, 5, 10, time.Minute))
				} else {
					// insert stable block info to db
					blockInfo := helper.ParseBlockToDB(block)
					if blockErr := blockInfo.UpsertG(ctx, true, []string{dbTrans.TBlockColumns.BlockNum}, boil.Infer(), boil.Infer()); blockErr != nil {
						s.logger.Error("insert block error",
							zap.Int64("block_num", block.Number().Int64()), zap.Error(blockErr))
						// @TBD: add error message to dead queue
						_ = s.redisClient.Set(ctx, fmt.Sprintf(consts.ErrInsertBlocksKey, block.Number().Int64()),
							blockErr.Error(), helper.GeneralDuration(6, 2, 12, time.Hour))
					}
				}
				s.logger.Debug("insert block success", zap.Int64("block_num", block.Number().Int64()))
				blocks <- block
			}(query)
		}
	}()
	return blocks
}

func (s *server) insertTransactions(ctx context.Context, blocks chan *types.Block, unstable bool) chan common.Hash {
	txHashs := make(chan common.Hash)
	limiter := make(chan struct{}, s.goTransNum)
	go func() {
		defer close(txHashs)
		defer close(limiter)
		var wg sync.WaitGroup
		defer wg.Wait()
		for block := range blocks {
			for _, tx := range block.Transactions() {
				limiter <- struct{}{}
				wg.Add(1)
				go func(trans *types.Transaction, blockNum int64) {
					defer func() {
						<-limiter
						wg.Done()
					}()

					if unstable {
						// save tx pb format to redis
						pbTxInfo := helper.ParseTransToPb(trans)
						bytes, _ := json.Marshal(pbTxInfo)
						_ = s.redisClient.Set(ctx,
							fmt.Sprintf(consts.TransactionDataKey, blockNum, trans.Hash().String()), bytes,
							helper.GeneralDuration(25, 5, 10, time.Minute))
					} else {
						// insert stable tx info to db
						data := helper.ParseTransToDB(trans)
						data.BlockNum = blockNum
						if err := data.UpsertG(ctx, true, []string{dbTrans.TTransactionColumns.TXHash}, boil.Infer(), boil.Infer()); err != nil {
							// @TBD: add error message to dead queue
							s.logger.Error("insert transaction error",
								zap.String("tx_hash", trans.Hash().String()), zap.Error(err))
							_ = s.redisClient.Set(ctx,
								fmt.Sprintf(consts.ErrInsertTransactionsKey, trans.Hash().Hex()),
								err.Error(), helper.GeneralDuration(6, 2, 12, time.Hour))
							return
						}
					}
					s.logger.Debug("insert transaction success", zap.String("tx_hash", trans.Hash().String()))
					txHashs <- trans.Hash()
				}(tx, int64(block.NumberU64()))
			}
		}
	}()
	return txHashs
}

func (s *server) queryReceipts(ctx context.Context, txHash chan common.Hash) chan *types.Receipt {
	receipts := make(chan *types.Receipt)
	limiter := make(chan struct{}, s.goReceiptsNum)
	go func() {
		defer close(receipts)
		defer close(limiter)
		var wg sync.WaitGroup
		defer wg.Wait()
		for tx := range txHash {
			limiter <- struct{}{}
			wg.Add(1)
			go func(query common.Hash) {
				defer func() {
					<-limiter
					wg.Done()
				}()
				receipt, err := s.rpcRelay.GetTransactionReceipt(ctx, query)
				if err != nil {
					// @TBD: add error message to dead queue
					s.logger.Error("rpc get receipt error",
						zap.String("tx_hash", query.Hex()), zap.Error(err))
					_ = s.redisClient.Set(ctx,
						fmt.Sprintf(consts.ErrQueryReceiptsKey, query.Hex()),
						err.Error(), helper.GeneralDuration(6, 2, 12, time.Hour))
					return
				}
				receipts <- receipt
				s.logger.Debug("rpc get receipt success",
					zap.String("tx_hash", query.Hex()))
			}(tx)
		}
	}()
	return receipts
}

func (s *server) insertLogs(ctx context.Context, receipts chan *types.Receipt, unstable bool) {
	limiter := make(chan struct{}, s.goLogNum)
	defer close(limiter)
	var wg sync.WaitGroup
	defer wg.Wait()
	for receipt := range receipts {
		for _, log := range receipt.Logs {
			limiter <- struct{}{}
			wg.Add(1)
			go func(logInfo *types.Log) {
				defer func() {
					<-limiter
					wg.Done()
				}()

				if unstable {
					// save log pb format to redis
					pbLogInfo := helper.ParseLogToPb(logInfo)
					bytes, _ := json.Marshal(pbLogInfo)
					_ = s.redisClient.Set(ctx,
						fmt.Sprintf(consts.LogDataKey, logInfo.TxHash.Hex(), int64(logInfo.Index)),
						bytes, helper.GeneralDuration(35, 5, 10, time.Minute))
				} else {
					// insert stable log info to db
					data := helper.ParseLogToDB(logInfo)
					if err := data.UpsertG(ctx, true, []string{dbTrans.TLogColumns.Index, dbTrans.TLogColumns.TXHash}, boil.Infer(), boil.Infer()); err != nil {
						// @TBD: add error message to dead queue
						s.logger.Error("insert log error",
							zap.String("tx_hash", logInfo.TxHash.Hex()),
							zap.Uint("index", logInfo.Index), zap.Error(err))
						_ = s.redisClient.Set(ctx,
							fmt.Sprintf(consts.ErrInsertLogsKey, logInfo.TxHash.Hex(), logInfo.Index),
							err.Error(), helper.GeneralDuration(6, 2, 12, time.Hour))
						return
					}
				}
				s.logger.Debug("insert log success",
					zap.String("tx_hash", logInfo.TxHash.Hex()),
					zap.Uint("index", logInfo.Index))
			}(log)
		}
	}
}
