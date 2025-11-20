package main

import (
	"fmt"
	"forum/config"
	"forum/controller"
	mysql "forum/dao"
	_ "forum/docs"
	"forum/logger"
	"forum/pkg/cache"
	"forum/pkg/gpt"
	"forum/pkg/mongodb"
	"forum/pkg/snowflake"
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
	// mongo初始化
	mongodb.Init()
	// redis 初始化
	cache.Init()
	//  websocket 管道监听
	go controller.Manager.Start()
	// 大模型初始化
	gpt.Init()
	// 雪花算法 生成分布式ID
	if err := snowflake.Init(1); err != nil {
		logger.Error("init snowflake failed err:\n")
		return
	}
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator Trans failed,err:%v\n", err)
		return
	}
	// 路由初始化
	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(config.AppConf.AppBaseConfig.Port)
}
// 家电舒适