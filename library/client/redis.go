package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/0x726f6f6b6965/ethereum-info/library/config"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient(cfg *config.RedisCfg) *redis.Client {
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.User,
		Password: cfg.Password,
		DB:       cfg.DB,
	}
	if cfg.MaxRetries != 0 {
		opt.MaxRetries = cfg.MaxRetries
	}
	return redis.NewClient(opt)
}

func GetDatabyKey(ctx context.Context, c *redis.Client, key string, data interface{}) error {
	value, err := c.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(value), data)
}
