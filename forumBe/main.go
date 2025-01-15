package main

import (
	"fmt"
	"forum/config"
	"forum/dao/mysql"
	_ "forum/docs"
	"forum/logger"
	"forum/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 配置读取
	if err := config.InitConfig(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 日志初始化
	if err := logger.Init(config.AppConf.LogConfig, "dev"); err != nil {
		fmt.Printf("logger init failed, err:%v\n", err)
		return
	}
	// mysql初始化
	if err := mysql.InitDB(); err != nil {
		fmt.Printf("Db init failed, err:%v\n", err)
		return
	}

	// 路由初始化

	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.AppConf.AppBaseConfig.Port)
}
