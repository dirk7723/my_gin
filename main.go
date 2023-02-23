package main

import (
	"log"

	"github.com/gin-gonic/gin"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
	"shutuiche.com/luka/go_test/global"
	"shutuiche.com/luka/go_test/router"
)

// var (
// 	DB *gorm.DB
// )

// init
func init() {
	//setting
	err := global.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	//logger
	err = global.SetupLogger()
	if err != nil {
		log.Fatalf("init.SetupLogger err: %v", err)
	}

	//access logger
	err = global.SetupAccessLogger()
	if err != nil {
		log.Fatalf("init.SetupAccessLogger err: %v", err)
	}

	var _ *gorm.DB
	_, err = global.SetupDB()
	if err != nil {
		log.Fatalf("init.SetupDB err: %v", err)
	}

	global.Logger.Infof("------应用init结束")
	//global.Logger.
}

func main() {
	global.Logger.Infof("------应用main函数开始")
	//设置运行模式
	gin.SetMode(global.ServerSetting.RunMode)
	//引入路由
	r := router.Router()
	//run
	r.Run(":" + global.ServerSetting.HttpPort)
}
