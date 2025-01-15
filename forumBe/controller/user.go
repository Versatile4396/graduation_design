package controller

import (
	mysql "forum/dao"
	"forum/logger"
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Ping godoc
// @Summary      用户注册
// @Description  用户注册接口
// @Success      200  {string}  string
// @body        name body     string  true  "name"
// @Produce      json
// @Router       /register [post]
func UserRegisterController(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(user)
	// 业务处理-注册用户
	if err := logic.SignUp(user); err != nil {
		zap.L().Error("logic.signUp failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "注册成功")
}
