package controller

import (
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommentCreateController(c *gin.Context) {
	// 评论创建
	var Comment *models.ArticleComment
	if err := c.ShouldBindJSON(&Comment); err != nil {
		zap.L().Error("comment create with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}

	rComment, err := logic.CommentCreate(Comment)
	if err != nil {
		zap.L().Error("logic.CommentCreate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rComment)
}
