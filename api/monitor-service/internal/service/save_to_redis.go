package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	libConfig "github.com/0x726f6f6b6965/ethereum-info/library/config"
	"github.com/0x726f6f6b6965/ethereum-info/monitor-service/internal/config"
	"github.com/0x726f6f6b6965/ethereum-info/monitor-service/internal/consts"
	"github.com/redis/go-redis/v9"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

func NewRedisMonitor(cfg *config.MonitorCfg, rpcCfg *libConfig.RpcClientCfg, rdbCfg *libConfig.RedisCfg, db *sql.DB, logger *zap.Logger) MonitorService {
	rpc, _ := client.InitRPCRelayByHttp(rpcCfg)
	redisClient := client.InitRedisClient(rdbCfg)
	ser := &server{
		rpcRelay:      rpc,
		rpcCfg:        rpcCfg,
		redisClient:   redisClient,
		db:            db,
		logger:        logger,
		goBlockNum:    cfg.GoBlockNum,
		goTransNum:    cfg.GoTransNum,
		goLogNum:      cfg.GoLogNum,
		goReceiptsNum: cfg.GoReceiptsNum,
		saveBlockNum:  cfg.SaveBlockNum,
		startBlock:    cfg.StartBlock,
		rollBack:      cfg.RollBack,
		intreval:      cfg.Interval,
		done:          make(chan bool),
		monitorType:   consts.UNSTABLE,
	}
	boil.SetDB(ser.db)
	return ser
}

func (s *server) saveToRedis(ctx context.Context, ticker *time.Ticker) {
	_ = s.redisClient.Set(ctx, "rollback", s.rollBack, -1)
	start, _ := s.loadLastUnstableBlock(ctx)
	for {
		select {
		case <-ticker.C:
			blockNum, err := s.rpcRelay.GetCurrentBlock(ctx)
			if err != nil || s.rpcRelay == nil {
				s.ConnectEth(ctx)
				continue
			}
			// unstable
			current := int64(blockNum) - s.rollBack
			if current < start && start-(s.rollBack/3) > current {
				current = start - (s.rollBack / 3)
			}

			blocks := s.queryBlocks(ctx, current, int64(blockNum), true)
			txs := s.insertTransactions(ctx, blocks, true)
			receipts := s.queryReceipts(ctx, txs)
			s.insertLogs(ctx, receipts, true)

			err = s.saveLastUnstableBlock(ctx, int64(blockNum))
			if err != nil {
				continue
			}
			start = int64(blockNum)

		case <-ctx.Done():
			s.StopMonitor()
		case <-s.done:
			return
		}
	}
}

// saveLastUnstableBlock - save this number as the last unstable block number
func (s *server) saveLastUnstableBlock(ctx context.Context, block int64) error {
	if s.redisClient == nil {
		s.startBlock = block
		return nil
	}
	err := s.redisClient.Set(ctx, consts.LastBlockKey, block, -1).Err()
	return err
}

// loadLastUnstableBlock - load the last unstable block number
func (s *server) loadLastUnstableBlock(ctx context.Context) (int64, error) {
	if s.redisClient == nil {
		return s.startBlock, nil
	}
	block, err := s.redisClient.Get(ctx, consts.LastBlockKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return s.startBlock, nil
		} else {
			return s.startBlock, err
		}
	}
	if block == "" {
		return s.startBlock, nil
	}
	startBlock, _ := strconv.ParseInt(block, 10, 64)
	return startBlock, nil
}
