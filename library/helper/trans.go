package helper

import (
	"strconv"

	dbTrans "github.com/0x726f6f6b6965/ethereum-info/db/repository"
	pbBlock "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	pbTrans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// block to db
func ParseBlockToDB(block *types.Block) *dbTrans.TBlock {
	blockInfo := &dbTrans.TBlock{
		BlockNum:   block.Number().Int64(),
		BlockHash:  block.Hash().Hex(),
		BlockTime:  int64(block.Time()),
		ParentHash: block.ParentHash().Hex(),
	}
	return blockInfo
}

// block to pb
func ParseBlockToPb(block *types.Block) *pbBlock.GetBlockResponse {
	blockInfo := &pbBlock.GetBlockResponse{
		BlockNum:   block.Number().String(),
		BlockHash:  block.Hash().Hex(),
		BlockTime:  strconv.FormatUint(block.Time(), 10),
		ParentHash: block.ParentHash().Hex(),
	}
	return blockInfo
}

// block db to pb
func ParseTBlockToPb(block *dbTrans.TBlock) *pbBlock.GetBlockResponse {
	blockInfo := &pbBlock.GetBlockResponse{
		BlockNum:   strconv.FormatInt(block.BlockNum, 10),
		BlockHash:  block.BlockHash,
		BlockTime:  strconv.FormatInt(block.BlockTime, 10),
		ParentHash: block.ParentHash,
	}
	if block.R != nil {
		for _, tx := range block.R.GetBlockNumTTransactions() {
			blockInfo.Transactions = append(blockInfo.Transactions, tx.TXHash)
		}
	}
	return blockInfo
}

// trans to db without block number
func ParseTransToDB(tx *types.Transaction) *dbTrans.TTransaction {
	from, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	result := &dbTrans.TTransaction{
		TXHash: tx.Hash().String(),
		From:   from.Hex(),
		Nonce:  int64(tx.Nonce()),
		Data:   common.Bytes2Hex(tx.Data()),
		Value:  tx.Value().String(),
	}
	if tx.To() != nil {
		result.To = tx.To().String()
	}
	return result
}

// trans to pb
func ParseTransToPb(tx *types.Transaction) *pbTrans.GetTransactionResponse {
	from, _ := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	result := &pbTrans.GetTransactionResponse{
		TxHash: tx.Hash().Hex(),
		From:   from.Hex(),
		Nonce:  strconv.FormatUint(tx.Nonce(), 10),
		Data:   Add0xPrefix(common.Bytes2Hex(tx.Data())),
		Value:  tx.Value().String(),
	}
	if tx.To() != nil {
		result.To = tx.To().String()
	}
	return result
}

// trans db to pb
func ParseTTransToPb(tx *dbTrans.TTransaction) *pbTrans.GetTransactionResponse {
	result := &pbTrans.GetTransactionResponse{
		TxHash: tx.TXHash,
		From:   tx.From,
		Nonce:  strconv.FormatInt(tx.Nonce, 10),
		Data:   Add0xPrefix(tx.Data),
		Value:  tx.Value,
	}

	if tx.R != nil {
		for _, log := range tx.R.GetTXHashTLogs() {
			temp := &pbTrans.Log{
				Data:  Add0xPrefix(log.Data),
				Index: uint32(log.Index),
			}
			result.Logs = append(result.Logs, temp)
		}
	}
	return result
}

// log to db
func ParseLogToDB(log *types.Log) *dbTrans.TLog {
	result := &dbTrans.TLog{
		TXHash: log.TxHash.Hex(),
		Index:  int64(log.Index),
		Data:   common.Bytes2Hex(log.Data),
	}
	return result
}

// log to pb
func ParseLogToPb(log *types.Log) *pbTrans.Log {
	result := &pbTrans.Log{
		TxHash: log.TxHash.Hex(),
		Data:   common.Bytes2Hex(log.Data),
		Index:  uint32(log.Index),
	}
	return result
}

// receipt to pb
func ParseReceiptToPb(receipt *types.Receipt) []*pbTrans.Log {
	result := []*pbTrans.Log{}
	for _, log := range receipt.Logs {
		temp := &pbTrans.Log{
			Data:  Add0xPrefix(common.Bytes2Hex(log.Data)),
			Index: uint32(log.Index),
		}
		result = append(result, temp)
	}
	return result
}
