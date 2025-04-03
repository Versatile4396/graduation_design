package controller

import (
	"forum/logic"
	"forum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AssistanceCreateController(c *gin.Context) {
	var assistance *models.Assistance
	if err := c.ShouldBindJSON(&assistance); err != nil {
		zap.L().Error("assistance create with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rAssistance, err := logic.CreateAssistance(assistance)
	if err != nil {
		zap.L().Error("assistance create with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rAssistance)
}

func AssistanceDeleteController(c *gin.Context) {
	aid := c.Param("aid")
	postId, err := strconv.ParseUint(aid, 10, 64)
	if err != nil {
		zap.L().Error("assistance delete with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	err = logic.DeleteAssistance(postId)
	if err != nil {
		zap.L().Error("assistance delete with invalid param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func AssistanceGetListController(c *gin.Context) {
	var filter models.AssistanceFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		zap.L().Error("assistance get list with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rAssistances, err := logic.GetAssistanceList(filter)
	if err != nil {
		zap.L().Error("assistance get list with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rAssistances)
}

func AssistanceUpdateController(c *gin.Context) {
	var assistance *models.Assistance
	if err := c.ShouldBindJSON(&assistance); err != nil {
		zap.L().Error("assistance update with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	err := logic.UpdateAssistance(assistance)
	if err != nil {
		zap.L().Error("assistance update with invalid Param", zap.Error(err))
	}
	ResponseSuccess(c, nil)
}

func AssistanceCommentCreateController(c *gin.Context) {

}

func AssistanceCommentDeleteController(c *gin.Context) {

}

func AssistanceCommentGetListController(c *gin.Context) {

}
