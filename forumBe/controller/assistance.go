package controller

import (
	"fmt"
	"forum/global"
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

func AssistanceChangtStatusController(c *gin.Context) {
	type StatusForm struct {
		AssistanceId uint64 `json:"assistance_id"`
		Status       int    `json:"status"`
	}
	var statusForm *StatusForm
	if err := c.ShouldBindJSON(&statusForm); err != nil {
		zap.L().Error("assistance change status with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	fmt.Println(statusForm.Status, statusForm.AssistanceId)
	err := global.Db.Model(&models.Assistance{}).Where("assistance_id =?", statusForm.AssistanceId).Update("status", statusForm.Status).Error
	if err != nil {
		zap.L().Error("assistance change status with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, nil)
}

func AssistanceCommentCreateController(c *gin.Context) {
	var assistanceComment *models.AssistanceComment
	if err := c.ShouldBindJSON(&assistanceComment); err != nil {
		zap.L().Error("assistance comment create with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rAssistanceComment, err := logic.CreateAssistanceComment(assistanceComment)
	if err != nil {
		zap.L().Error("assistance comment create with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rAssistanceComment)
}

func AssistanceCommentDeleteController(c *gin.Context) {
	var id = c.Param("cid")
	postId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		zap.L().Error("assistance comment delete with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	err = logic.DeleteAssistanceComment(postId)
	if err != nil {
		zap.L().Error("assistance comment delete with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func AssistanceCommentGetListController(c *gin.Context) {
	var filter models.AssistanceCommentFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		zap.L().Error("assistance comment get list with invalid Param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rAssistanceComments, err := logic.GetAssistanceCommentList(filter)
	if err != nil {
		zap.L().Error("assistance comment get list with invalid Param", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rAssistanceComments)
}
