package config

import (
	libCfg "github.com/0x726f6f6b6965/ethereum-info/library/config"
)

type Config struct {
	Name  string              `yaml:"name" help:"the application name"`
	Grpc  libCfg.Grpc         `yaml:"grpc" help:"the application grpc option"`
	DB    libCfg.DBConfig     `yaml:"db" help:"the application db option"`
	Redis libCfg.RedisCfg     `yaml:"redis" help:"the application redis option"`
	RPC   libCfg.RpcClientCfg `yaml:"rpc" help:"the rpc option"`
	Log   libCfg.Log          `yaml:"log" help:"the application log"`
}
