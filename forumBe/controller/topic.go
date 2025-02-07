package controller

import (
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TopicCreateController(c *gin.Context) {
	var topic *models.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		zap.L().Error("topic create with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rTopic, err := logic.TopicCreate(topic)
	if err != nil {
		zap.L().Error("logic.TopicCreate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rTopic)
}

func TopicGetListController(c *gin.Context) {
	var Pagination *models.Pagination
	if err := c.ShouldBindJSON(&Pagination); err != nil {
		zap.L().Error("topic get list with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rTopics, err := logic.TopicGetList(Pagination)
	if err != nil {
		zap.L().Error("logic.TopicGetList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rTopics)
}

func TopicUpdateController(c *gin.Context) {

}

func TopicGetController(c *gin.Context) {

}

func TopicDeleteController(c *gin.Context) {

}
