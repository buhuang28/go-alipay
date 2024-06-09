package controller

import (
	"bytes"
	"github.com/buhuang28/go-alipay/api/service"
	"github.com/buhuang28/go-alipay/cst"
	"github.com/buhuang28/go-alipay/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
)

var (
	Alipay AlipayController
)

type AlipayController struct {
}

func (a *AlipayController) ReciveNotify(c *gin.Context) {
	log.Info("收到支付回调")
	raw, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.String(200, cst.NOTIFU_ERROR)
		return
	}
	var req dto.AlipayNotifyReq
	c.Request.Body = io.NopCloser(bytes.NewBuffer(raw))
	err = c.Bind(&req)
	log.Infof("收到订单:%s 的支付回调", req.OutTradeNo)
	if err != nil {
		log.Error(err)
		c.String(200, cst.NOTIFU_ERROR)
		return
	}
	err = service.AlipayVerifySign(string(raw))
	if err != nil {
		log.Infof("订单:%s 回调校验错误:%v", req.OutTradeNo, err)
		c.String(200, cst.NOTIFU_ERROR)
	} else {
		log.Infof("订单:%s 回调校验成功", req.OutTradeNo)
		c.String(200, cst.NOTIFY_SUCCESS)
	}
	return
}
