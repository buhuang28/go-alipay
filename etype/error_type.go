package etype

import "errors"

var (
	NullDataError             = errors.New("空数据")
	FilePathError             = errors.New("错误的文件路径")
	AlipayTradePrecreateError = errors.New("预创建订单失败")
)
