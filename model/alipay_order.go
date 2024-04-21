package model

// 支付宝订单表
type AlipayOrder struct {
	ID uint `gorm:"primaryKey;autoIncrement;not null;"`

	//【描述】商户订单号。
	//由商家自定义，64个字符以内，仅支持字母、数字、下划线且需保证在商户端不重复。
	//【示例值】20150320010101001
	OutTeadeNo string `gorm:"size:64"`

	/*
		total_amount必选price(12)
		【描述】订单总金额，单位为元，精确到小数点后两位，取值范围为 [0.01,100000000]，金额不能为 0。
		如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	*/
	TotalAmount float64

	/*
		【描述】订单标题。
		注意：不可使用特殊字符，如 /，=，& 等。
		【示例值】Iphone6 16G
	*/
	Subject string `gorm:"size:256"`

	/*
		【描述】销售产品码。如果签约的是当面付快捷版，则传 OFFLINE_PAYMENT；其它支付宝当面付产品传 FACE_TO_FACE_PAYMENT；不传则默认使用 FACE_TO_FACE_PAYMENT。
		【示例值】FACE_TO_FACE_PAYMENT
	*/
	ProductCode string `gorm:"size:64"`

	/*
		【描述】卖家支付宝用户 ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户 ID。不允许收款账号与付款方账号相同
		【示例值】2088102146225135
	*/
	SellerId string `gorm:"size:32"`

	/*
		【描述】订单附加信息。
		如果请求时传递了该参数，将在异步通知、对账单中原样返回，同时会在商户和用户的pc账单详情中作为交易描述展示
		【示例值】Iphone6 16G
	*/
	Body string `gorm:"size:128"`

	//商品订单ID，关联商品订单
	ItemOrderId string

	//支付状态
	PayStatus int
}

func (a *AlipayOrder) TableName() string {
	return "alipay_order"
}
