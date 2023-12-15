package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/0x726f6f6b6965/ethereum-info/block-service/internal/consts"
	dbTrans "github.com/0x726f6f6b6965/ethereum-info/db/repository"
	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	libGRPC "github.com/0x726f6f6b6965/ethereum-info/library/grpc"
	"github.com/0x726f6f6b6965/ethereum-info/library/helper"
	blocks "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	"github.com/redis/go-redis/v9"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type server struct {
	blocks.UnsafeBlockServiceServer
	rpcClient   *client.RPCRelay
	redisClient *redis.Client
	logger      *zap.Logger
	db          *sql.DB
}

// GetBlock implements v1.BlockServiceServer.
func (s *server) GetBlock(ctx context.Context, req *blocks.GetBlockRequest) (*blocks.GetBlockResponse, error) {
	if req.Id <= 0 {
		return nil, libGRPC.RequiredFieldErr("no id", "id")
	}
	resp := &blocks.GetBlockResponse{}
	// get from db
	mod := []qm.QueryMod{qm.Load(dbTrans.TBlockRels.BlockNumTTransactions), dbTrans.TBlockWhere.BlockNum.EQ(int64(req.Id))}
	info, err := dbTrans.TBlocks(mod...).OneG(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Info("block id not found in db", zap.Uint64("id", req.Id))
		} else {
			s.logger.Error("error in get block", zap.Error(err))
		}
	}

	if info == nil {
		txs := []string{}
		// get from redis
		data, _ := s.redisClient.Get(ctx, fmt.Sprintf(consts.BlockDataKey, req.Id)).Result()
		if !helper.IsEmpty(data) {
			json.Unmarshal([]byte(data), resp)
			iter := s.redisClient.Scan(ctx, 0, fmt.Sprintf(consts.BlockDataKey+":", req.Id), 0).Iterator()
			for iter.Next(ctx) {
				temp := strings.Split(iter.Val(), ":")
				if len(temp) > 3 {
					txs = append(txs, temp[3])
				}
			}
		} else {
			// get from rpc
			block, err := s.rpcClient.GetBlockByNum(ctx, big.NewInt(int64(req.Id)))
			if err != nil {
				if errors.Is(err, errors.New("not found")) {
					return nil, libGRPC.NotFoundErr("block id not found", "id",
						fmt.Sprintf("%d", req.Id))
				}
				s.logger.Error("error in get block", zap.Error(err))
				return nil, libGRPC.InternalErr("error in get block")
			}
			resp = helper.ParseBlockToPb(block)
			for _, tx := range block.Transactions() {
				txs = append(txs, tx.Hash().Hex())
			}
		}
		resp.Transactions = append(resp.Transactions, txs...)
	} else {
		resp = helper.ParseTBlockToPb(info)
		resp.Stable = true
	}

	return resp, nil
}

// GetLatestBlockList implements v1.BlockServiceServer.
func (s *server) GetLatestBlockList(ctx context.Context, req *blocks.GetLatestBlockListRequest) (*blocks.GetLatestBlockListResponse, error) {
	num, err := s.rpcClient.GetCurrentBlock(ctx)
	if err != nil {
		s.logger.Error("error in get block", zap.Error(err))
		return nil, libGRPC.InternalErr("error in get block")
	}

	start := num - req.Limit + 1
	// get from db
	mod := []qm.QueryMod{dbTrans.TBlockWhere.BlockNum.GTE(int64(start))}
	infos, err := dbTrans.TBlocks(mod...).AllG(ctx)
	if err != nil {
		s.logger.Error("error in get block", zap.Error(err))
	}

	resp := &blocks.GetLatestBlockListResponse{}
	for _, info := range infos {
		temp := helper.ParseTBlockToPb(info)
		temp.Stable = true
		resp.Blocks = append(resp.Blocks, temp)
	}

	if len(infos) < int(req.Limit) {
		if err != nil {
			s.logger.Error("error in get block", zap.Error(err))
			return nil, libGRPC.InternalErr("error in get block")
		}
		for i := len(infos); i < int(req.Limit); i++ {
			blockInfo := &blocks.GetBlockResponse{}
			// get from redis
			query := int(num) - i
			data, err := s.redisClient.Get(context.Background(), fmt.Sprintf(consts.BlockDataKey, query)).Result()
			if err != nil {
				// get from rpc
				block, err := s.rpcClient.GetBlockByNum(ctx, big.NewInt(int64(query)))
				if err != nil {
					s.logger.Error("error in get block", zap.Error(err))
					return nil, libGRPC.InternalErr("error in get block")
				}
				blockInfo = helper.ParseBlockToPb(block)
			} else {
				json.Unmarshal([]byte(data), blockInfo)
			}
			resp.Blocks = append(resp.Blocks, blockInfo)
		}
	}

	return resp, nil
}

func NewBlockService(db *sql.DB, rpcClient *client.RPCRelay, redisClient *redis.Client, logger *zap.Logger) blocks.BlockServiceServer {
	ser := &server{
		db:          db,
		rpcClient:   rpcClient,
		redisClient: redisClient,
		logger:      logger,
	}
	boil.SetDB(ser.db)

	return ser
}
