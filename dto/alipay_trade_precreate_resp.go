package dto

type AlipayTradePrecreateResp struct {
	AlipayTradePrecreateResponse AlipayTradePrecreateResponse `json:"alipay_trade_precreate_response"`
	Sign                         string                       `json:"sign"`
}

type AlipayTradePrecreateResponse struct {
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
	ShareCode  string `json:"share_code,omitempty"`
}
