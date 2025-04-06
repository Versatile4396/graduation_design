package controller

import (
	"errors"
	"fmt"
	mysql "forum/dao"
	"forum/logger"
	"forum/logic"
	"forum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	// 业务处理-注册用户
	rUser, err := logic.SignUp(user)
	if err != nil {
		zap.L().Error("logic.signUp failed", zap.Error(err))
		if err.Error() == mysql.ErrorUserExit {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccessWithMsg(c, rUser, "注册成功")
}

func UserLoginController(c *gin.Context) {
	var u *models.LoginForm
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("login faild", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		return
	}
	logger.Fmt(u)
	// 业务处理 用户登录
	user, err := logic.Login(u)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidParams)
	}
	// 3、返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserId), //js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name":     user.UserName,
		"avatar":        user.Avatar,
		"gender":        user.Gender,
		"nickname":      user.Nickname,
		"overview":      user.Overview,
		"email":         user.Email,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

func UserGetCountController(c *gin.Context) {
	// 获取用户信息
	// 业务处理-获取用户信息
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 64)
	countInfo, err := logic.GetCountController(uid)
	if err != nil {
		zap.L().Error("logic.GetUserInfo failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, countInfo)
}

func UserGetListController(c *gin.Context) {
	filter := &models.UserFilter{}
	if c.Request.ContentLength != 0 {
		if err := c.ShouldBindJSON(&filter); err != nil {
			zap.L().Error("GetUserList with invalid param", zap.Error(err))
			ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
			return
		}
	}
	// 业务处理-获取用户信息
	userList, err := logic.GetUserList(filter)
	if err != nil {
		zap.L().Error("logic.GetUserList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, userList)
}

func UserUpdateController(c *gin.Context) {
	var user *models.UserUpdateForm
	if err := c.ShouldBindJSON(&user); err != nil {
		zap.L().Error("update with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	// 业务处理-更新用户信息
	err := logic.UpdateUser(user)

	if err != nil {
		zap.L().Error("logic.UpdateUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func UserDeleteController(c *gin.Context) {
	// 业务处理-删除用户信息
	type deleteForm struct {
		UserId  uint64 `json:"user_id"`
		Deleted int    `json:"deleted"`
	}
	var fo *deleteForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("delete with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}

	err := logic.DeleteUser(fo.UserId, fo.Deleted)
	if err != nil {
		zap.L().Error("logic.DeleteUser failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
