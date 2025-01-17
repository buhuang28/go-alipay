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

// alipay.trade.precreate(统一收单线下交易预创建)
func TradePrecreate(data dto.AlipayTradePrecreateReq) (dto.AlipayTradePrecreateResponse, error) {
	ceateReq := dto.NewAlipayPublicReq(cst.TRADE_PRECREATE_METHOD, caches.AlipayConfig.AlipayNotifyUrl, data.ToString())
	err := ceateReq.RSASign()
	if err != nil {
		log.Error(err)
		return dto.AlipayTradePrecreateResponse{}, err
	}

	resp := utils.HttpPostForm(caches.AlipayConfig.AlipayUrl, ceateReq.ToUrlValues())
	if resp.Err != nil {
		log.Error(resp.Err)
		return dto.AlipayTradePrecreateResponse{}, resp.Err
	}
	if len(resp.Data) == 0 {
		log.Error(etype.NullDataError)
		return dto.AlipayTradePrecreateResponse{}, etype.NullDataError
	}

	createResp := new(dto.AlipayTradePrecreateResp)

	err = json.Unmarshal(resp.Data, createResp)
	if err != nil {
		log.Error(err)
		return dto.AlipayTradePrecreateResponse{}, err
	}
	if createResp.AlipayTradePrecreateResponse.Code != cst.ALIPAY_SUCCESS_CODE {
		log.Errorf("创建预支付订单失败:%s", createResp.AlipayTradePrecreateResponse.Msg)
		return dto.AlipayTradePrecreateResponse{}, etype.AlipayTradePrecreateError
	}
	return createResp.AlipayTradePrecreateResponse, nil
}
