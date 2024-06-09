package dto

type AlipayNotifyReq struct {
	GmtCreate      string `form:"gmt_create"`
	Charset        string `form:"charset"`
	SellerEmail    string `form:"seller_email"`
	Subject        string `form:"subject"`
	Sign           string `form:"sign"`
	BuyerId        string `form:"buyer_id"`
	Body           string `form:"body"`
	InvoiceAmount  string `form:"invoice_amount"`
	NotifyId       string `form:"notify_id"`
	FundBillList   string `form:"fund_bill_list"`
	NotifyType     string `form:"notify_type"`
	TradeStatus    string `form:"trade_status"`
	ReceiptAmount  string `form:"receipt_amount"`
	BuyerPayAmount string `form:"buyer_pay_amount"`
	AppId          string `form:"app_id"`
	SignType       string `form:"sign_type"`
	SellerId       string `form:"seller_id"`
	GmtPayment     string `form:"gmt_payment"`
	NotifyTime     string `form:"notify_time"`
	Version        string `form:"version"`
	OutTradeNo     string `form:"out_trade_no"`
	TotalAmount    string `form:"total_amount"`
	TradeNo        string `form:"trade_no"`
	AuthAppId      string `form:"auth_app_id"`
	BuyerLogonId   string `form:"buyer_logon_id"`
	PointAmount    string `form:"point_amount"`
}
