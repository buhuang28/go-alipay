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

// alipay.trade.query(统一收单交易查询)
func TradeQuery(data dto.AlipayTradeQueryReq) (dto.AlipayTradeQueryResponse, error) {
	ceateReq := dto.NewAlipayPublicReq(cst.TRADE_QUERY_METHOD, caches.AlipayConfigCache.AlipayNotifyUrl, data.ToString())
	err := ceateReq.RSASign()
	if err != nil {
		log.Error(err)
		return dto.AlipayTradeQueryResponse{}, err
	}

	resp := utils.HttpPostForm(caches.AlipayConfigCache.AlipayUrl, ceateReq.ToUrlValues())
	if resp.Err != nil {
		log.Error(resp.Err)
		return dto.AlipayTradeQueryResponse{}, resp.Err
	}
	if len(resp.Data) == 0 {
		log.Error(etype.NullDataError)
		return dto.AlipayTradeQueryResponse{}, etype.NullDataError
	}

	queryResp := new(dto.AlipayTradeQueryResp)

	err = json.Unmarshal(resp.Data, queryResp)
	if err != nil {
		log.Error(err)
		return dto.AlipayTradeQueryResponse{}, err
	}
	if queryResp.AlipayTradeQueryResponse.Code != cst.ALIPAY_SUCCESS_CODE {
		log.Errorf("创建预支付订单失败:%s", queryResp.AlipayTradeQueryResponse.Msg)
		return dto.AlipayTradeQueryResponse{}, etype.AlipayTradePrecreateError
	}
	return queryResp.AlipayTradeQueryResponse, nil
}
