package caches

import "github.com/buhuang28/go-alipay/config"

var (
	AlipayConfigCache config.AlipayConfig
)

func LoadAlipayConfig(c config.AlipayConfig) {
	AlipayConfigCache = c
}
