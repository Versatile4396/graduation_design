package controller

import (
	"errors"
	"fmt"
	mysql "forum/dao"
	"forum/logger"
	"forum/logic"
	"forum/models"
	"forum/pkg/cache"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
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
	// // 验证码
	code, err := cache.RedisClient.Get(c, user.Email).Result()
	if err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, "验证码错误")
		return
	}
	if code != user.Captcha {
		ResponseErrorWithMsg(c, CodeCaptchaWrong, "验证码错误")
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

func SendEmailCodeController(c *gin.Context) {
	var fo *models.UserValidateEmailForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("sendEmailCode with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	// 业务处理-发送验证码 有问题 不管了
	vCode, err := SendEmailCode(fo.Email)
	// 保存验证码到redis中
	cache.RedisClient.Set(c, fo.Email[0], vCode, 5*time.Minute)
	if err != nil {
		zap.L().Error("SendEmailCode failed", zap.Error(err))
		ResponseSuccessWithMsg(c, vCode, "验证码发送成功")
		return
	}
	ResponseSuccessWithMsg(c, vCode, "验证码发送成功")
}

func SendEmailCode(em []string) (string, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("发件人笔名 <1808368025@qq.com>")
	e.To = em
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	//设置文件发送的内容
	content := fmt.Sprintf(`
		<div>
			<div>
				尊敬的%s，您好！
			</div>
			<div style="padding: 8px 40px 8px 50px;">
				<p>您于 %s 提交的邮箱验证，本次验证码为<u><strong>%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
			</div>
			<div>
				<p>此邮箱为系统邮箱，请勿回复。</p>
			</div>
		</div>
	`, em[0], t, vCode)
	e.Text = []byte(content)
	//设置服务器相关的配置
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "1808368025@qq.com", "lxexguvmgcuqcgdb", "smtp.qq.com"))
	return vCode, err
}

func UserValidateEmailController(c *gin.Context) {
	var fo *models.UserValidateEmailForm
	if err := c.ShouldBindJSON(&fo); err != nil {
		zap.L().Error("validateEmail with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
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
