package router

import (
	"forum/controller"
	"forum/logger"
	"forum/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	// 设置中间件 recovery 中间件会recovery项目中kennel会出现的panic
	router.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.RateLimitMiddleware(2*time.Second, 40))

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.UserRegisterController)
		user.POST("/login", controller.UserLoginController)
	}

	article := router.Group("/api/article")
	{
		article.POST("/create", controller.ArticleCreateController)
		article.POST("/update", controller.ArticleUpdateController)
		article.GET("/getById/:aid", controller.ArticleGetController)
		article.DELETE("/delete/:aid", controller.ArticleDeleteController)
	}
	return router
}
