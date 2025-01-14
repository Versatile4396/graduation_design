package controller

import (
	models "be/model"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary      用户注册
// @Description  用户注册接口
// @Success      200  {string}  string
// @body        name body     string  true  "name"
// @Produce      json
// @Router       /register [post]
func UserRegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
	}

}

func UserLoginController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}

func UserUpdateController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
func UserDeleteController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
func UserInfoController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
