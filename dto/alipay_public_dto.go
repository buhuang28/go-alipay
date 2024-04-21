package dto

type AlipayPublicDTO struct {
	AppId        string `json:"app_id,omitempty"`         //支付宝分配给开发者的应用ID
	Method       string `json:"method,omitempty"`         //接口名称
	Format       string `json:"format,omitempty"`         //仅支持JSON
	Charset      string `json:"charset,omitempty"`        //请求使用的编码格式，如utf-8,gbk,gb2312等
	Sign         string `json:"sign,omitempty"`           //商户请求参数的签名串
	Timestamp    string `json:"timestamp,omitempty"`      //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `json:"version,omitempty"`        //调用的接口版本，固定为：1.0
	NotifyUrl    string `json:"notify_url,omitempty"`     //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	AppAuthToken string `json:"app_auth_token,omitempty"` //详见应用授权概述
	BizContent   string `json:"biz_content,omitempty"`    //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}
