package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"encoding/json"

	"github.com/0x726f6f6b6965/ethereum-info/graph-service/internal/model"
	pbBlock "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	pbTrans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

// Block is the resolver for the block field.
func (r *queryResolver) Block(ctx context.Context, num uint64) (*model.Block, error) {
	resp, err := r.Blocks.GetBlock(ctx, &pbBlock.GetBlockRequest{Id: num})
	if err != nil {
		return nil, err
	}
	b, err := protojson.Marshal(resp)
	if err != nil {
		return nil, err
	}
	result := &model.Block{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LatestBlocks is the resolver for the latestBlocks field.
func (r *queryResolver) LatestBlocks(ctx context.Context, num uint64) (*model.Blocks, error) {
	resp, err := r.Blocks.GetLatestBlockList(ctx, &pbBlock.GetLatestBlockListRequest{Limit: num})
	if err != nil {
		return nil, err
	}
	b, err := protojson.Marshal(resp)
	if err != nil {
		return nil, err
	}
	result := &model.Blocks{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Transaction is the resolver for the transaction field.
func (r *queryResolver) Transaction(ctx context.Context, txHash string) (*model.Transaction, error) {
	resp, err := r.Trans.GetTransaction(ctx, &pbTrans.GetTransactionRequest{TxHash: txHash})
	if err != nil {
		return nil, err
	}
	b, err := protojson.Marshal(resp)
	if err != nil {
		return nil, err
	}
	result := &model.Transaction{}
	err = json.Unmarshal(b, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
