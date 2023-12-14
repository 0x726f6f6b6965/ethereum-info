//go:generate go run ./../tools/generate.go
package graph

import (
	pbBlock "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	pbTrans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Blocks pbBlock.BlockServiceClient
	Trans  pbTrans.TransactionServiceClient
}
