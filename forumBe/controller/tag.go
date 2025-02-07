package controller

import (
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TagCreateController(c *gin.Context) {
	var tag *models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		zap.L().Error("tag create with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rTag, err := logic.TagCreate(tag)
	if err != nil {
		zap.L().Error("logic.TagCreate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rTag)

}
func TagGetListController(c *gin.Context) {
	var Pagination *models.Pagination
	if err := c.ShouldBindJSON(&Pagination); err != nil {
		zap.L().Error("tag get list with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rTags, err := logic.TagGetList(Pagination)
	if err != nil {
		zap.L().Error("logic.TagGetList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rTags)
}

func TagUpdateController(c *gin.Context) {

}

func TagGetController(c *gin.Context) {

}

func TagDeleteController(c *gin.Context) {

}
