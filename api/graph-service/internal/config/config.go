package config

import (
	libCfg "github.com/0x726f6f6b6965/ethereum-info/library/config"
)

type Config struct {
	Name    string      `yaml:"name" help:"the application name"`
	Rest    libCfg.Rest `yaml:"rest" help:"the application rest option"`
	Log     libCfg.Log  `yaml:"log" help:"the application log"`
	Clients struct {
		Block       libCfg.Client `yaml:"block-svc"`
		Transaction libCfg.Client `yaml:"transaction-svc"`
	} `yaml:"clients"`
}
