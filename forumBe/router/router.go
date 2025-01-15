package router

import (
	"forum/controller"
	"forum/logger"
	"forum/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	// 设置中间件 recovery 中间件会recovery项目中kennel会出现的panic
	router.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.RateLimitMiddleware(2*time.Second, 40))

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.UserRegisterController)
	}
	return router
}
