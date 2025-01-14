package main

import (
	"be/Log"
	"be/config"
	_ "be/docs"
	"be/routers"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @BasePath  /api
func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	config.InitDB()

	if err := Log.Init(config.AppConfig.LogConfig, config.AppConfig.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	router := routers.InitRouters()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Println("Server started at http://127.0.0.1:5555")
	router.Run(":5555")
}
