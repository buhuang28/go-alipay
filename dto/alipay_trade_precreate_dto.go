package dto

// 参考 https://opendocs.alipay.com/apis/02890k?pathHash=0f11e06a
type (
	AlipayTradePrecreateDTO struct {
		OutTradeNo           string         `json:"out_trade_no,omitempty"`
		TotalAmount          string         `json:"total_amount,omitempty"`
		Subject              string         `json:"subject,omitempty"`
		Body                 string         `json:"body,omitempty"`
		StoreID              string         `json:"store_id,omitempty"`
		TimeExpire           string         `json:"time_expire,omitempty"`
		ExtendParams         ExtendParams   `json:"extend_params,omitempty"`
		QueryOptions         []string       `json:"query_options,omitempty"`
		SettleInfo           SettleInfo     `json:"settle_info,omitempty"`
		OperatorID           string         `json:"operator_id,omitempty"`
		ProductCode          string         `json:"product_code,omitempty"`
		GoodsDetail          []GoodsDetail  `json:"goods_detail,omitempty"`
		MerchantOrderNo      string         `json:"merchant_order_no,omitempty"`
		EnablePayChannels    string         `json:"enable_pay_channels,omitempty"`
		DiscountableAmount   string         `json:"discountable_amount,omitempty"`
		BusinessParams       BusinessParams `json:"business_params,omitempty"`
		TimeoutExpress       string         `json:"timeout_express,omitempty"`
		DisablePayChannels   string         `json:"disable_pay_channels,omitempty"`
		UndiscountableAmount string         `json:"undiscountable_amount,omitempty"`
		BkagentReqInfo       BkagentReqInfo `json:"bkagent_req_info,omitempty"`
		SellerID             string         `json:"seller_id,omitempty"`
		TerminalID           string         `json:"terminal_id,omitempty"`
		CodeType             string         `json:"code_type,omitempty"`
	}

	ExtendParams struct {
		SysServiceProviderID  string `json:"sys_service_provider_id,omitempty"`
		TcInstallmentOrderID  string `json:"tc_installment_order_id,omitempty"`
		SpecifiedSellerName   string `json:"specified_seller_name,omitempty"`
		RoyaltyFreeze         string `json:"royalty_freeze,omitempty"`
		CardType              string `json:"card_type,omitempty"`
		TradeComponentOrderID string `json:"trade_component_order_id,omitempty"`
	}

	SettleDetailInfos struct {
		Amount           string `json:"amount,omitempty"`
		TransIn          string `json:"trans_in,omitempty"`
		SettleEntityType string `json:"settle_entity_type,omitempty"`
		SummaryDimension string `json:"summary_dimension,omitempty"`
		ActualAmount     string `json:"actual_amount,omitempty"`
		SettleEntityID   string `json:"settle_entity_id,omitempty"`
		TransInType      string `json:"trans_in_type,omitempty"`
	}

	SettleInfo struct {
		SettlePeriodTime  string              `json:"settle_period_time,omitempty"`
		SettleDetailInfos []SettleDetailInfos `json:"settle_detail_infos,omitempty"`
	}

	GoodsDetail struct {
		OutSkuID       string `json:"out_sku_id,omitempty"`
		GoodsName      string `json:"goods_name,omitempty"`
		Quantity       int    `json:"quantity,omitempty"`
		Price          string `json:"price,omitempty"`
		OutItemID      string `json:"out_item_id,omitempty"`
		GoodsID        string `json:"goods_id,omitempty"`
		GoodsCategory  string `json:"goods_category,omitempty"`
		CategoriesTree string `json:"categories_tree,omitempty"`
		ShowURL        string `json:"show_url,omitempty"`
	}

	BusinessParams struct {
		GoodTaxes              string `json:"good_taxes,omitempty"`
		EnterprisePayInfo      string `json:"enterprise_pay_info,omitempty"`
		ActualOrderTime        string `json:"actual_order_time,omitempty"`
		CampusCard             string `json:"campus_card,omitempty"`
		CardType               string `json:"card_type,omitempty"`
		EnterprisePayAmount    string `json:"enterprise_pay_amount,omitempty"`
		TinyAppMerchantBizType string `json:"tiny_app_merchant_biz_type,omitempty"`
		McCreateTradeIP        string `json:"mc_create_trade_ip,omitempty"`
	}

	BkagentReqInfo struct {
		MerchCode  string `json:"merch_code,omitempty"`
		AcqCode    string `json:"acq_code,omitempty"`
		DeviceType string `json:"device_type,omitempty"`
		Location   string `json:"location,omitempty"`
		SerialNum  string `json:"serial_num,omitempty"`
	}
)

func NewBaseAlipayTradePrecreateDTO(outTradeNo, totalAmount, subject, body string) AlipayTradePrecreateDTO {
	return AlipayTradePrecreateDTO{
		OutTradeNo:  outTradeNo,
		TotalAmount: totalAmount,
		Subject:     subject,
		Body:        body,
	}
}
