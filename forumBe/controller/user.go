package controller

import (
	mysql "forum/dao"
	"forum/logger"
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	userId, err := logic.SignUp(user)
	if err != nil {
		zap.L().Error("logic.signUp failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	res := &models.User{
		UserId:   userId,
		Email:    user.Email,
		Gender:   user.Gender,
		UserName: user.UserName,
	}
	ResponseSuccessWithMsg(c, res, "注册成功")
}

func UserLoginController(c *gin.Context) {
	var user *models.LoginForm
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("login faild", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	logger.Fmt(user)

	// 业务处理 用户登录
}
