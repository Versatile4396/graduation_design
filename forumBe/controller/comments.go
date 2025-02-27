package controller

import (
	"forum/logic"
	"forum/models"
	"strconv"

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

func CommentGetListController(c *gin.Context) {
	// 评论列表获取
	var params *models.CommentFilter
	if err := c.ShouldBindJSON(&params); err != nil {
		zap.L().Error("comment get list with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	// 业务处理-评论列表获取
	rComments, err := logic.CommentGetList(params)
	if err != nil {
		zap.L().Error("logic.CommentGetList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rComments)
}

func CommentDeleteController(c *gin.Context) {
	// 评论删除
	var err error
	comment_id := c.Param("cid")
	postId, _ := strconv.ParseInt(comment_id, 10, 64)
	// 业务处理-评论删除
	err = logic.CommentDelete(postId)
	if err != nil {
		zap.L().Error("logic.CommentDelete failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
