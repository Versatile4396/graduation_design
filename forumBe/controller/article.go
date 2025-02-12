package controller

import (
	"forum/logger"
	"forum/logic"
	"forum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ArticleCreateController(c *gin.Context) {
	var article *models.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		zap.L().Error("article create with invalid Param", zap.Error(err))
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
	aid := c.Param("aid")
	postId, err := strconv.ParseInt(aid, 10, 64)
	if err != nil {
		zap.L().Error("article delete with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(aid)
	// 业务处理-文章删除
	rAticle, err := logic.ArticleDelete(postId)
	if err != nil {
		zap.L().Error("logic.ArticleDelete failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rAticle)
}

func ArticleGetController(c *gin.Context) {
	article_id := c.Param("aid")
	postId, err := strconv.ParseInt(article_id, 10, 64)
	if err != nil {
		zap.L().Error("article get with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(article_id)
	// 业务处理-文章获取
	rArticle, rUserInfo, err := logic.ArticleGet(postId)
	if err != nil {
		zap.L().Error("logic.ArticleGet failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	type ArticleUserInfo struct {
		ArticleInfo *models.Article  "json:\"article_info\""
		AuthorInfo  *models.UserInfo "json:\"author_info\""
	}
	userInfo := models.UserInfo{
		UserId:   rArticle.UserId,
		UserName: rUserInfo.UserName,
		Avatar:   rUserInfo.Avatar,
		Email:    rUserInfo.Email,
		Gender:   rUserInfo.Gender,
	}
	resData := ArticleUserInfo{
		rArticle,
		&userInfo,
	}
	ResponseSuccessWithMsg(c, resData, "")
}

func ArticleGetListController(c *gin.Context) {
	// 获取文章列表 通过uid 通过category_id 通过topic_id 通过tag_id 等等
	var params *models.ArticleFilter
	if err := c.ShouldBindJSON(&params); err != nil {
		zap.L().Error("article get list with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(params)
	// 业务处理-文章列表获取
	rArticles, err := logic.ArticleGetList(params)
	if err != nil {
		zap.L().Error("logic.ArticleGetList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rArticles)
}

func ArticleCategoryGetListController(c *gin.Context) {
	// 获取文章分类列表
	// 业务处理-文章分类列表获取
	var pagination *models.Pagination
	if err := c.ShouldBindJSON(&pagination); err != nil {
		zap.L().Error("article category get list with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rArticleCategories, err := logic.ArticleCategoryGetList(pagination)
	if err != nil {
		zap.L().Error("logic.ArticleCategoryGetList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rArticleCategories)
}

func ArticleLikeController(c *gin.Context) {
	// 文章点赞
	var like *models.ArticleLike
	if err := c.ShouldBindJSON(&like); err != nil {
		zap.L().Error("article like with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	logger.Fmt(like)
	// 业务处理-文章点赞
	err := logic.ArticleLike(like)
	if err != nil {
		zap.L().Error("logic.ArticleLike failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func ArticleViewController(c *gin.Context) {
	// 文章浏览
	aid := c.Param("aid")
	// 业务处理-文章浏览
	err := logic.ArticleView(aid)
	if err != nil {
		zap.L().Error("logic.ArticleView failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
