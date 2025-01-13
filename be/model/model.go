package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    MyCode      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

func ResponseError(ctx *gin.Context, c MyCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(ctx *gin.Context, code MyCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}
