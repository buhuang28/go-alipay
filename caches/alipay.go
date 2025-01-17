package caches

import "github.com/buhuang28/go-alipay/config"

var (
	AlipayConfig config.AlipayConfig
)

func LoadAlipayConfig(c config.AlipayConfig) {
	AlipayConfig = c
}
