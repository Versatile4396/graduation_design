package controller

import (
	"forum/logger"
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ArticleCreateController(c *gin.Context) {
	var article *models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		zap.L().Error("article create with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(article)
	// 业务处理-文章创建
	rArticle, err := logic.ArticleCreate(article)
	if err != nil {
		zap.L().Error("logic.ArticleCreate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rArticle)
}

func ArticleUpdateController(c *gin.Context) {
	var article *models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		zap.L().Error("article update with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(article)
	// 业务处理-文章更新
	rArticle, err := logic.ArticleUpdate(article)
	if err != nil {
		zap.L().Error("logic.ArticleUpdate failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rArticle)
}
func ArticleDeleteController(c *gin.Context) {
	var article *models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		zap.L().Error("article delete with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(article)
	// 业务处理-文章删除
	err := logic.ArticleDelete(article)
	if err != nil {
		zap.L().Error("logic.ArticleDelete failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
}

func ArticleGetController(c *gin.Context) {
	var article *models.Article
	if err := c.BindJSON(&article); err != nil {
		zap.L().Error("article get with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(article)
	// 业务处理-文章获取
	rArticle, err := logic.ArticleGet(article)
	if err != nil {
		zap.L().Error("logic.ArticleGet failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rArticle)
}
