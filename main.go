package main

import (
	"github.com/buhuang28/go-alipay/config"
	"github.com/buhuang28/go-alipay/db"
	"github.com/buhuang28/go-alipay/logs"
)

func init() {
	logs.LogInit()
	c := config.LoadConfig("./config.json")
	db.DbInit(c.DBConfig)

}

func main() {

}
