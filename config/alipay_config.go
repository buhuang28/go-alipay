package config

type AlipayConfig struct {
	AlipayUrl       string `json:"alipay_url"`
	AppPrivateKey   string `json:"app_private_key"`
	AppPublicKey    string `json:"app_public_key"`
	AlipayPublicKey string `json:"alipay_public_key"`
	AlipayAppId     string `json:"alipay_app_id"`
	AlipayNotifyUrl string `json:"alipay_notify_url"`
}
