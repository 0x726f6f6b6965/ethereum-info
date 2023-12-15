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
	"github.com/0x726f6f6b6965/ethereum-info/library/logger"
	pbBlock "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	pbTrans "github.com/0x726f6f6b6965/ethereum-info/protos/transaction/v1"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func main() {
	godotenv.Load()
	path := os.Getenv("CONFIG")
	var cfg config.Config
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("read yaml error", err)
		return
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal("unmarshal yaml error", err)
		return
	}
	zaplog, cleanup, err := logger.NewLogger(&cfg.Log)
	if err != nil {
		log.Fatal("create log error", err)
		return
	}
	defer cleanup()

	connBlock, err := client.NewGrpcConn(
		context.Background(),
		fmt.Sprintf("%s:%d", cfg.Clients.Block.Host, cfg.Clients.Block.GrpcPort),
		cfg.Clients.Block.ConnectionTimeout)
	if err != nil {
		zaplog.Error("connect block service error", zap.Error(err))
		return
	}
	connTrans, err := client.NewGrpcConn(
		context.Background(),
		fmt.Sprintf("%s:%d", cfg.Clients.Transaction.Host, cfg.Clients.Transaction.GrpcPort),
		cfg.Clients.Transaction.ConnectionTimeout)
	if err != nil {
		zaplog.Error("connect transaction service error", zap.Error(err))
		return
	}
	resolver := &graph.Resolver{
		Blocks: pbBlock.NewBlockServiceClient(connBlock.GetConn()),
		Trans:  pbTrans.NewTransactionServiceClient(connTrans.GetConn()),
		Log:    zaplog,
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	zaplog.Info("create api success", zap.String("endpoint", fmt.Sprintf("http://localhost:%d/", cfg.Rest.Port)))
	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.Rest.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Rest.Port), nil))
}
