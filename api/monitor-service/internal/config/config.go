package config

import (
	"time"

	libCfg "github.com/0x726f6f6b6965/ethereum-info/library/config"
)

type MonitorCfg struct {
	GoBlockNum    int           `json:"go_block_num" yaml:"go_block_num"`
	GoTransNum    int           `json:"go_trans_num" yaml:"go_trans_num"`
	GoLogNum      int           `json:"go_log_num" yaml:"go_log_num"`
	GoReceiptsNum int           `json:"go_recepits_num" yaml:"go_recepits_num"`
	StartBlock    int64         `json:"start_block" yaml:"start_block"`
	RollBack      int64         `json:"roll_back" yaml:"roll_back"`
	SaveBlockNum  int64         `json:"save_block_num" yaml:"save_block_num"`
	Interval      time.Duration `json:"interval" yaml:"interval"`
}

type Config struct {
	Monitors struct {
		Stable   MonitorCfg `yaml:"stable"`
		Unstable MonitorCfg `yaml:"unstable"`
	} `yaml:"monitors"`
	DB    libCfg.DBConfig     `yaml:"db" help:"the application db option"`
	Redis libCfg.RedisCfg     `yaml:"redis" help:"the application redis option"`
	RPC   libCfg.RpcClientCfg `yaml:"rpc" help:"the rpc option"`
	Log   libCfg.Log          `yaml:"log" help:"the application log"`
}
