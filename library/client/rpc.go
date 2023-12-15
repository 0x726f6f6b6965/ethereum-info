package client

import (
	"context"
	"math/big"
	"time"

	"github.com/0x726f6f6b6965/ethereum-info/library/config"
	"github.com/0x726f6f6b6965/ethereum-info/library/helper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RPCRelay struct {
	client *ethclient.Client
	cfg    *config.RpcClientCfg
	retry  int
}

func InitRPCRelayByHttp(cfg *config.RpcClientCfg) (relay *RPCRelay, err error) {
	var (
		timeoutCtx, cancel = context.WithTimeout(context.Background(), time.Duration(cfg.ConnectTimeout)*time.Second)
		randKey            int
	)
	if len(cfg.Servers) <= 1 {
		randKey = 0
	} else {
		randKey = helper.RandInt(0, len(cfg.Servers)-1)
	}
	defer cancel()
	client, err := ethclient.DialContext(timeoutCtx, cfg.Servers[randKey])
	if err != nil {
		return
	}

	relay = &RPCRelay{
		client: client,
		retry:  cfg.Retry,
		cfg:    cfg,
	}
	return
}

func (r *RPCRelay) GetBlockByNum(ctx context.Context, number *big.Int) (result *types.Block, err error) {
	for try := 0; try <= r.retry; try++ {
		result, err = r.client.BlockByNumber(ctx, number)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return
}

func (r *RPCRelay) GetTransactionReceipt(ctx context.Context, txHash common.Hash) (receipt *types.Receipt, err error) {
	for try := 0; try <= r.retry; try++ {
		receipt, err = r.client.TransactionReceipt(ctx, txHash)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return
}

func (r *RPCRelay) GetTransaction(ctx context.Context, txHash common.Hash) (tx *types.Transaction, pending bool, err error) {
	for try := 0; try <= r.retry; try++ {
		tx, pending, err = r.client.TransactionByHash(ctx, txHash)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return
}

func (r *RPCRelay) GetCurrentBlock(ctx context.Context) (num uint64, err error) {
	for try := 0; try <= r.retry; try++ {
		num, err = r.client.BlockNumber(ctx)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	return
}
