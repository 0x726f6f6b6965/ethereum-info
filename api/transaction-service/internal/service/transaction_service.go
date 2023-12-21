package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	dbTrans "github.com/0x726f6f6b6965/ethereum-info/db/repository"
	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	libGRPC "github.com/0x726f6f6b6965/ethereum-info/library/grpc"
	"github.com/0x726f6f6b6965/ethereum-info/library/helper"
	trans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
	"github.com/0x726f6f6b6965/ethereum-info/transaction-service/internal/consts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

type server struct {
	trans.UnsafeTransactionServiceServer
	rpcClient   *client.RPCRelay
	redisClient *redis.Client
	logger      *zap.Logger
	db          *sql.DB
}

// GetTransaction implements v1.TransactionServiceServer.
func (s *server) GetTransaction(ctx context.Context, req *trans.GetTransactionRequest) (*trans.GetTransactionResponse, error) {
	if helper.IsEmpty(req.TxHash) {
		return nil, libGRPC.RequiredFieldErr("no tx hash", "tx_hash")
	}
	resp := &trans.GetTransactionResponse{}
	// get stable data from redis
	err := client.GetDatabyKey(ctx, s.redisClient, fmt.Sprintf(consts.TransactionStableDataKey, req.GetTxHash()), resp)
	if err == nil {
		return resp, nil
	}
	// get from db
	mod := []qm.QueryMod{qm.Load(dbTrans.TTransactionRels.TXHashTLogs), dbTrans.TTransactionWhere.TXHash.EQ(req.TxHash)}
	info, err := dbTrans.TTransactions(mod...).OneG(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Info("tx hash not found in db", zap.String("tx_hash", req.TxHash))
		} else {
			s.logger.Error("error in get transaction", zap.Error(err))
		}
	}
	if info == nil {
		// get from redis
		iter := s.redisClient.Scan(ctx, 0,
			fmt.Sprintf(consts.TransactionDataScanKey, req.TxHash), 0).Iterator()
		if iter.Next(ctx) {
			data, _ := s.redisClient.Get(ctx, iter.Val()).Result()
			_ = json.Unmarshal([]byte(data), resp)
			if !helper.IsEmpty(resp.Data) {
				resp.Data = helper.Add0xPrefix(resp.Data)
			}
			logIter := s.redisClient.Scan(ctx, 0, fmt.Sprintf(consts.TransactionLogKey+":*", req.TxHash), 0).Iterator()
			for logIter.Next(ctx) {
				logData := &trans.Log{}
				temp, _ := s.redisClient.Get(ctx, logIter.Val()).Result()
				_ = json.Unmarshal([]byte(temp), logData)
				if !helper.IsEmpty(logData.Data) {
					logData.Data = helper.Add0xPrefix(logData.Data)
				}
				resp.Logs = append(resp.Logs, logData)
			}
		} else {
			// get from rpc
			tx, pending, err := s.rpcClient.GetTransaction(ctx, common.HexToHash(req.TxHash))
			if err != nil {
				if errors.Is(err, errors.New("not found")) {
					return nil, libGRPC.NotFoundErr("tx_hash not found", "tx_hash", req.TxHash)
				}
				s.logger.Error("error in get transaction", zap.Error(err))
				return nil, libGRPC.InternalErr("error in get transaction")
			}
			resp = helper.ParseTransToPb(tx)

			if !pending {
				//get log from rpc
				receipt, _ := s.rpcClient.GetTransactionReceipt(ctx, common.HexToHash(req.TxHash))
				resp.Logs = helper.ParseReceiptToPb(receipt)
			}
		}
	} else {
		resp = helper.ParseTTransToPb(info)
	}

	if resp.Stable {
		bytes, _ := json.Marshal(resp)
		s.redisClient.Set(ctx, fmt.Sprintf(consts.TransactionStableDataKey, req.GetTxHash()), bytes, helper.GeneralDuration(5, 5, 10, time.Minute))
	}
	return resp, nil
}

func NewTransactionService(db *sql.DB, rpcClient *client.RPCRelay, redisClient *redis.Client, logger *zap.Logger) trans.TransactionServiceServer {
	ser := &server{
		db:          db,
		rpcClient:   rpcClient,
		redisClient: redisClient,
		logger:      logger,
	}
	boil.SetDB(ser.db)

	return ser
}
