package dto

type AlipayTradeQueryResp struct {
	AlipayTradeQueryResponse AlipayTradeQueryResponse `json:"alipay_trade_query_response"`
	Sign                     string                   `json:"sign"`
}
type FundBillList struct {
	FundChannel string `json:"fund_channel"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
}

type AlipayTradeQueryResponse struct {
	Code            string         `json:"code"`
	Msg             string         `json:"msg"`
	TradeNo         string         `json:"trade_no"`
	OutTradeNo      string         `json:"out_trade_no"`
	BuyerLogonID    string         `json:"buyer_logon_id"`
	TradeStatus     string         `json:"trade_status"`
	TotalAmount     string         `json:"total_amount"`
	BuyerPayAmount  string         `json:"buyer_pay_amount"`
	PointAmount     string         `json:"point_amount"`
	InvoiceAmount   string         `json:"invoice_amount"`
	SendPayDate     string         `json:"send_pay_date"`
	ReceiptAmount   string         `json:"receipt_amount"`
	StoreID         string         `json:"store_id"`
	TerminalID      string         `json:"terminal_id"`
	FundBillList    []FundBillList `json:"fund_bill_list"`
	StoreName       string         `json:"store_name"`
	BuyerUserID     string         `json:"buyer_user_id"`
	BuyerOpenID     string         `json:"buyer_open_id"`
	BuyerUserType   string         `json:"buyer_user_type"`
	MdiscountAmount string         `json:"mdiscount_amount"`
	DiscountAmount  string         `json:"discount_amount"`
	ExtInfos        string         `json:"ext_infos"`
}
