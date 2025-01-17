package config

import (
	"encoding/json"
	"github.com/buhuang28/go-alipay/etype"
	"github.com/buhuang28/go-alipay/utils"
)

type Config struct {
	AlipayConfig AlipayConfig `json:"alipay_config,omitempty"`
}

func LoadConfig(cfgPath string) *Config {
	if cfgPath == "" {
		panic(etype.FilePathError)
	}
	fileBytes := utils.ReadFile(cfgPath)
	if len(fileBytes) == 0 {
		panic(etype.NullDataError)
	}
	cfg := new(Config)
	err := json.Unmarshal(fileBytes, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func LoadConfigBytes(fileBytes []byte) *Config {
	cfg := new(Config)
	err := json.Unmarshal(fileBytes, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
