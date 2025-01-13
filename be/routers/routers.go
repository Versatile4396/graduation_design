package routers

import (
	"be/controller"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.New()
	router.Use()
	user := router.Group("/api/user")
	{
		user.POST("/register", controller.UserRegisterController)
		user.POST("/login", controller.UserLoginController)
		user.POST("/update", controller.UserUpdateController)
		user.GET("/delete", controller.UserDeleteController)
		user.GET("/info", controller.UserInfoController)
	}
	article := router.Group("/api/article")
	{
		article.POST("/create", func(c *gin.Context) {

		})
	}

	return router
}
