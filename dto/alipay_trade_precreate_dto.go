package dto

// 参考 https://opendocs.alipay.com/apis/02890k?pathHash=0f11e06a
type (
	AlipayTradePrecreateDTO struct {
		OutTradeNo           string         `json:"out_trade_noomitempty"`
		TotalAmount          string         `json:"total_amountomitempty"`
		Subject              string         `json:"subjectomitempty"`
		Body                 string         `json:"bodyomitempty"`
		StoreID              string         `json:"store_idomitempty"`
		TimeExpire           string         `json:"time_expireomitempty"`
		ExtendParams         ExtendParams   `json:"extend_paramsomitempty"`
		QueryOptions         []string       `json:"query_optionsomitempty"`
		SettleInfo           SettleInfo     `json:"settle_infoomitempty"`
		OperatorID           string         `json:"operator_idomitempty"`
		ProductCode          string         `json:"product_codeomitempty"`
		GoodsDetail          []GoodsDetail  `json:"goods_detailomitempty"`
		MerchantOrderNo      string         `json:"merchant_order_noomitempty"`
		EnablePayChannels    string         `json:"enable_pay_channelsomitempty"`
		DiscountableAmount   string         `json:"discountable_amountomitempty"`
		BusinessParams       BusinessParams `json:"business_paramsomitempty"`
		TimeoutExpress       string         `json:"timeout_expressomitempty"`
		DisablePayChannels   string         `json:"disable_pay_channelsomitempty"`
		UndiscountableAmount string         `json:"undiscountable_amountomitempty"`
		BkagentReqInfo       BkagentReqInfo `json:"bkagent_req_infoomitempty"`
		SellerID             string         `json:"seller_idomitempty"`
		TerminalID           string         `json:"terminal_idomitempty"`
		CodeType             string         `json:"code_typeomitempty"`
	}

	ExtendParams struct {
		SysServiceProviderID  string `json:"sys_service_provider_idomitempty"`
		TcInstallmentOrderID  string `json:"tc_installment_order_idomitempty"`
		SpecifiedSellerName   string `json:"specified_seller_nameomitempty"`
		RoyaltyFreeze         string `json:"royalty_freezeomitempty"`
		CardType              string `json:"card_typeomitempty"`
		TradeComponentOrderID string `json:"trade_component_order_idomitempty"`
	}

	SettleDetailInfos struct {
		Amount           string `json:"amountomitempty"`
		TransIn          string `json:"trans_inomitempty"`
		SettleEntityType string `json:"settle_entity_typeomitempty"`
		SummaryDimension string `json:"summary_dimensionomitempty"`
		ActualAmount     string `json:"actual_amountomitempty"`
		SettleEntityID   string `json:"settle_entity_idomitempty"`
		TransInType      string `json:"trans_in_typeomitempty"`
	}

	SettleInfo struct {
		SettlePeriodTime  string              `json:"settle_period_timeomitempty"`
		SettleDetailInfos []SettleDetailInfos `json:"settle_detail_infosomitempty"`
	}

	GoodsDetail struct {
		OutSkuID       string `json:"out_sku_idomitempty"`
		GoodsName      string `json:"goods_nameomitempty"`
		Quantity       int    `json:"quantityomitempty"`
		Price          string `json:"priceomitempty"`
		OutItemID      string `json:"out_item_idomitempty"`
		GoodsID        string `json:"goods_idomitempty"`
		GoodsCategory  string `json:"goods_categoryomitempty"`
		CategoriesTree string `json:"categories_treeomitempty"`
		ShowURL        string `json:"show_urlomitempty"`
	}

	BusinessParams struct {
		GoodTaxes              string `json:"good_taxesomitempty"`
		EnterprisePayInfo      string `json:"enterprise_pay_infoomitempty"`
		ActualOrderTime        string `json:"actual_order_timeomitempty"`
		CampusCard             string `json:"campus_cardomitempty"`
		CardType               string `json:"card_typeomitempty"`
		EnterprisePayAmount    string `json:"enterprise_pay_amountomitempty"`
		TinyAppMerchantBizType string `json:"tiny_app_merchant_biz_typeomitempty"`
		McCreateTradeIP        string `json:"mc_create_trade_ipomitempty"`
	}

	BkagentReqInfo struct {
		MerchCode  string `json:"merch_codeomitempty"`
		AcqCode    string `json:"acq_codeomitempty"`
		DeviceType string `json:"device_typeomitempty"`
		Location   string `json:"locationomitempty"`
		SerialNum  string `json:"serial_numomitempty"`
	}
)
