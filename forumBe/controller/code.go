package controller

type MyCode int64

const (
	CodeSuccess            MyCode = 1000
	CodeInvalidParams      MyCode = 1001
	CodeUserExist          MyCode = 1002
	CodeUserNotExist       MyCode = 1003
	CodeInvalidPassword    MyCode = 1004
	CodeServerBusy         MyCode = 1005
	CodeInvalidToken       MyCode = 1006
	CodeInvalidAuthFormat  MyCode = 1007
	CodeNotLogin           MyCode = 1008
	ErrVoteRepeated        MyCode = 1009
	ErrorVoteTimeExpire    MyCode = 1010
	CodeInvalidExt         MyCode = 1011
	CodeInvalidSize        MyCode = 1012
	CodeLimiteTimes        MyCode = 1013
	CodeConnectionSuccess  MyCode = 1014
	CodeCaptchaFail        MyCode = 1015
	CodeCaptchaWrong       MyCode = 1016
	CodeFollowAlreadyExist MyCode = 1017
)

var msgFlags = map[MyCode]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidExt:         "文件后缀名不合法",
	CodeInvalidSize:        "文件大小超出限制",
	CodeInvalidToken:       "无效的Token",
	CodeInvalidAuthFormat:  "认证格式有误",
	CodeNotLogin:           "未登录",
	ErrVoteRepeated:        "请勿重复投票",
	ErrorVoteTimeExpire:    "投票时间已过",
	CodeCaptchaFail:        "验证码发送失败",
	CodeCaptchaWrong:       "验证码错误",
	CodeFollowAlreadyExist: "已经关注该用户",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
