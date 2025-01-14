package main

import (
	"be/config"
	_ "be/docs"
	"be/routers"
	"log/slog"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @BasePath  /api
func main() {
	config.InitConfig()
	router := routers.InitRouters()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	slog.Info("Server started at http://127.0.0.1:5555")
	router.Run(":5555")
}
