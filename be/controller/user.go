package controller

import (
	"be/Log"
	"be/global"
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
	Log.Fmt(user)
	// c.JSON(200, gin.H{
	// 	"message": user.Email,
	// })
	if err := global.Db.Create(&user).Error; err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "注册失败"+err.Error())
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
