package dto

import (
	"bytes"
	"github.com/buhuang28/go-alipay/caches"
	"github.com/buhuang28/go-alipay/cst"
	"github.com/buhuang28/go-alipay/utils"
	log "github.com/sirupsen/logrus"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type AlipayPublicReq struct {
	AppId        string `json:"app_id"`                   //支付宝分配给开发者的应用ID
	Method       string `json:"method"`                   //接口名称
	Format       string `json:"format,omitempty"`         //仅支持JSON
	Charset      string `json:"charset"`                  //请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType     string `json:"sign_type"`                //商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Sign         string `json:"sign"`                     //商户请求参数的签名串
	Timestamp    string `json:"timestamp"`                //发送请求的时间，格式"yyyy-MM-dd HH:mm:ss"
	Version      string `json:"version"`                  //调用的接口版本，固定为：1.0
	NotifyUrl    string `json:"notify_url,omitempty"`     //支付宝服务器主动通知商户服务器里指定的页面http/https路径。
	AppAuthToken string `json:"app_auth_token,omitempty"` //详见应用授权概述
	BizContent   string `json:"biz_content"`              //请求参数的集合，最大长度不限，除公共参数外所有请求参数都必须放在这个参数中传递，具体参照各产品快速接入文档
}

func NewAlipayPublicReq(method, notifyUrl, bizContent string) AlipayPublicReq {
	return AlipayPublicReq{
		AppId:      caches.AlipayConfigCache.AlipayAppId,
		Method:     method,
		Format:     cst.ALIPAY_FORMAT,
		Charset:    cst.ALIPAY_CHARSET,
		SignType:   cst.ALIPAY_SIGN_TYPE,
		Timestamp:  time.Now().Format(time.DateTime),
		Version:    cst.ALIPAY_VERSION,
		NotifyUrl:  notifyUrl,
		BizContent: bizContent,
	}
}

func (a *AlipayPublicReq) RSASign() error {
	/*
		加签原理
		1. 获取所有支付宝开放平台的 post 内容，不包括字节类型参数，如文件、字节流，剔除 sign 字段，剔除值为空的参数；
		2. 按照第一个字符的键值 ASCII 码递增排序（字母升序排序），如果遇到相同字符则按照第二个字符的键值 ASCII 码递增排序，以此类推；
		3. 将排序后的参数与其对应值，组合成 参数=参数值 的格式，并且把这些参数用 & 字符连接起来，此时生成的字符串为待签名字符串。
	*/

	//代码生成参考 gen_sign_test.go
	var buf bytes.Buffer
	if a.AppAuthToken != "" {
		buf.WriteString("&app_auth_token=")
		buf.WriteString(a.AppAuthToken)
	}
	buf.WriteString("&app_id=")
	buf.WriteString(a.AppId)
	buf.WriteString("&biz_content=")
	buf.WriteString(a.BizContent)
	buf.WriteString("&charset=")
	buf.WriteString(a.Charset)
	if a.Format != "" {
		buf.WriteString("format=")
		buf.WriteString(a.Format)
	}
	buf.WriteString("&method=")
	buf.WriteString(a.Method)
	if a.NotifyUrl != "" {
		buf.WriteString("notify_url=")
		buf.WriteString(a.NotifyUrl)
	}
	buf.WriteString("&sign=")
	buf.WriteString(a.Sign)
	buf.WriteString("&sign_type=")
	buf.WriteString(a.SignType)
	buf.WriteString("&timestamp=")
	buf.WriteString(a.Timestamp)
	buf.WriteString("&version=")
	buf.WriteString(a.Version)

	sign, err := utils.RsaSha256(caches.AlipayConfigCache.AppPrivateKey, buf.String())
	if err != nil {
		log.Error(err)
		return err
	}
	a.Sign = sign
	return nil
}

func (a *AlipayPublicReq) ToUrlValues() url.Values {
	var data url.Values
	tof := reflect.TypeOf(a)
	vof := reflect.ValueOf(a)
	for i := 0; i < tof.NumField(); i++ {
		tag := tof.Field(i).Tag.Get("json")
		tag, _, ok := strings.Cut(tag, ",omitempty")
		v := vof.Field(i).String()
		if ok {
			if v == "" {
				data.Add(tag, "")
			}
		} else {
			data.Add(tag, v)
		}
	}
	return data
}
