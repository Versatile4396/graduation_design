package controller

import (
	"forum/global"
	"forum/logger"
	"forum/models"

	"github.com/gin-gonic/gin"
)

func UserRegisterController(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
	}
	logger.Fmt(user)
	// c.JSON(200, gin.H{
	// 	"message": user.Email,
	// })
	if err := global.Db.Create(&user).Error; err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "注册失败"+err.Error())
	}
}
