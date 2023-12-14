package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/0x726f6f6b6965/ethereum-info/graph-service/graph"
	"github.com/0x726f6f6b6965/ethereum-info/graph-service/internal/client"
	"github.com/0x726f6f6b6965/ethereum-info/graph-service/internal/config"
	pbBlock "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	pbTrans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	path := os.Getenv("CONFIG")
	var cfg config.Config
	err := hclsimple.DecodeFile(path, nil, &cfg)
	if err != nil {
		return
	}

	connBlock, err := client.NewGrpcConn(
		context.Background(),
		fmt.Sprintf("%s:%d", cfg.Clients.Block.Host, cfg.Clients.Block.GrpcPort),
		cfg.Clients.Block.ConnectionTimeout)
	if err != nil {
		return
	}
	connTrans, err := client.NewGrpcConn(
		context.Background(),
		fmt.Sprintf("%s:%d", cfg.Clients.Transaction.Host, cfg.Clients.Transaction.GrpcPort),
		cfg.Clients.Transaction.ConnectionTimeout)
	if err != nil {
		return
	}
	resolver := &graph.Resolver{
		Blocks: pbBlock.NewBlockServiceClient(connBlock.GetConn()),
		Trans:  pbTrans.NewTransactionServiceClient(connTrans.GetConn()),
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.Rest.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Rest.Port), nil))
}
