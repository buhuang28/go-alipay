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

// alipay.trade.close(统一收单交易关闭接口)
func TradeClose(data dto.AlipayTradeCloseReq) (dto.AlipayTradeCloseResponse, error) {
	ceateReq := dto.NewAlipayPublicReq(cst.TRADE_PRECREATE_METHOD, caches.AlipayConfig.AlipayNotifyUrl, data.ToString())
	err := ceateReq.RSASign()
	if err != nil {
		log.Error(err)
		return dto.AlipayTradeCloseResponse{}, err
	}

	resp := utils.HttpPostForm(caches.AlipayConfig.AlipayUrl, ceateReq.ToUrlValues())
	if resp.Err != nil {
		log.Error(resp.Err)
		return dto.AlipayTradeCloseResponse{}, resp.Err
	}
	if len(resp.Data) == 0 {
		log.Error(etype.NullDataError)
		return dto.AlipayTradeCloseResponse{}, etype.NullDataError
	}

	createResp := new(dto.AlipayTradeCloseResp)

	err = json.Unmarshal(resp.Data, createResp)
	if err != nil {
		log.Error(err)
		return dto.AlipayTradeCloseResponse{}, err
	}
	if createResp.AlipayTradeCloseResponse.Code != cst.ALIPAY_SUCCESS_CODE {
		log.Errorf("关闭订单失败:%s", createResp.AlipayTradeCloseResponse.Msg)
		return dto.AlipayTradeCloseResponse{}, etype.AlipayTradePrecreateError
	}
	return createResp.AlipayTradeCloseResponse, nil
}
