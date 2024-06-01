package alipay

import (
	"encoding/json"
	"github.com/buhuang28/go-alipay/caches"
	"github.com/buhuang28/go-alipay/cst"
	"github.com/buhuang28/go-alipay/dto"
	"github.com/buhuang28/go-alipay/etype"
	"github.com/buhuang28/go-alipay/utils"
	log "github.com/sirupsen/logrus"
)

// 预创建订单 -- 返回路径
func TradePrecreate(data dto.AlipayTradePrecreateReq) (string, error) {
	ceateReq := dto.NewAlipayPublicReq(cst.TRADE_PRECREATE_METHOD, caches.AlipayConfigCache.AlipayNotifyUrl, data.ToString())
	err := ceateReq.RSASign()
	if err != nil {
		log.Error(err)
		return "", err
	}

	resp := utils.HttpPostForm(caches.AlipayConfigCache.AlipayUrl, ceateReq.ToUrlValues())
	if resp.Err != nil {
		log.Error(resp.Err)
		return "", resp.Err
	}
	if len(resp.Data) == 0 {
		log.Error(etype.NullDataError)
		return "", etype.NullDataError
	}

	createResp := new(dto.AlipayTradePrecreateResp)

	err = json.Unmarshal(resp.Data, createResp)
	if err != nil {
		log.Error(err)
		return "", err
	}
	if createResp.AlipayTradePrecreateResponse.Code != cst.ALIPAY_SUCCESS_CODE {
		log.Errorf("创建预支付订单失败:%s", createResp.AlipayTradePrecreateResponse.Msg)
		return "", etype.AlipayTradePrecreateError
	}
	return createResp.AlipayTradePrecreateResponse.QrCode, nil
}
