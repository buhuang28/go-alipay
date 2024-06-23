package dto

type AlipayTradeCloseResp struct {
	AlipayTradeCloseResponse AlipayTradeCloseResponse `json:"alipay_trade_close_response"`
	Sign                     string                   `json:"sign"`
}

type AlipayTradeCloseResponse struct {
	Code       string `json:"code,omitempty"`
	Msg        string `json:"msg,omitempty"`
	OutTradeNo string `json:"out_trade_no,omitempty"`
	QrCode     string `json:"qr_code,omitempty"`
}
