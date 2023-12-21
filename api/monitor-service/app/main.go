package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0x726f6f6b6965/ethereum-info/library/client"
	"github.com/0x726f6f6b6965/ethereum-info/library/logger"
	"github.com/0x726f6f6b6965/ethereum-info/monitor-service/internal/config"
	"github.com/0x726f6f6b6965/ethereum-info/monitor-service/internal/service"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
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

	fwServer := service.NewDBMonitor(&cfg.Monitors.Stable, &cfg.RPC, &cfg.Redis, db, zaplog)
	fwCtx, fwCancel := context.WithCancel(context.Background())
	fwServer.MonitorBlocks(fwCtx)

	unServer := service.NewRedisMonitor(&cfg.Monitors.Unstable, &cfg.RPC, &cfg.Redis, db, zaplog)
	unCtx, unCancel := context.WithCancel(context.Background())
	unServer.MonitorBlocks(unCtx)

	shoutdown := make(chan os.Signal, 1)
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGQUIT)
	signal.Notify(shoutdown, os.Interrupt)

	select {
	case <-stop:
		fwServer.StopMonitor()
		unServer.StopMonitor()
		for fwServer.IsMonitoring() || unServer.IsMonitoring() {
			time.Sleep(5 * time.Second)
		}
	case <-shoutdown:
		fwCancel()
		unCancel()
		for fwServer.IsMonitoring() || unServer.IsMonitoring() {
			time.Sleep(5 * time.Second)
		}
	}

}
