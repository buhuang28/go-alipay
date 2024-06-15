package dto

import "encoding/json"

type AlipayTradeQueryReq struct {
	/*
		订单支付时传入的商户订单号,和支付宝交易号不能同时为空。
		trade_no,out_trade_no如果同时存在优先取trade_no
	*/
	OutTradeNo string `json:"out_trade_no,omitempty"`

	/*
		支付宝交易号，和商户订单号不能同时为空
	*/
	TradeNo string `json:"trade_no,omitempty"`

	QueryOptions []string `json:"query_options,omitempty"`
}

func (a *AlipayTradeQueryReq) ToString() string {
	marshal, _ := json.Marshal(a)
	return string(marshal)
}
