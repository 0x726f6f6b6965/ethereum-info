package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/0x726f6f6b6965/ethereum-info/block-service/internal/config"
	"github.com/0x726f6f6b6965/ethereum-info/block-service/internal/service"
	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	"github.com/0x726f6f6b6965/ethereum-info/library/logger"
	blocks "github.com/0x726f6f6b6965/ethereum-info/protos/blocks/v1"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

	db, dbCleanup, err := client.NewPostgres(&cfg.DB)
	if err != nil {
		zaplog.Error("failed to connect db", zap.Error(err))
		return
	}
	defer dbCleanup()

	rpcClient, err := client.InitRPCRelayByHttp(&cfg.RPC)
	if err != nil {
		zaplog.Error("failed to connect rpc", zap.Error(err))
		return
	}

	redisClient := client.InitRedisClient(&cfg.Redis)
	server := service.NewBlockService(db, rpcClient, redisClient, zaplog)

	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		zaplog.Error("failed to listen", zap.Error(err))
		return
	}

	blocks.RegisterBlockServiceServer(grpcServer, server)
	zaplog.Info("server listening", zap.String("addr", lis.Addr().String()))
	if err := grpcServer.Serve(lis); err != nil {
		zaplog.Error("failed to serve", zap.Error(err))
		return
	}
}
