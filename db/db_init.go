package db

import (
	"fmt"
	"github.com/buhuang28/go-alipay/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DbLink *gorm.DB
)

func DbInit(c config.DBConfig) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Port, c.Ip, &c.Port, c.DBName)
	DbLink, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(err)
		panic(err)
	}
}
