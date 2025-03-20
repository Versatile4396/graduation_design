package router

import (
	"forum/controller"
	"forum/logger"
	"forum/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	// 设置中间件 recovery 中间件会recovery项目中kennel会出现的panic
	router.Use(logger.GinLogger(),
		logger.GinRecovery(true),
		middleware.RateLimitMiddleware(2*time.Second, 4000))
	router.Static("/images", "./images")

	user := router.Group("/api/user")
	{
		user.POST("/register", controller.UserRegisterController)
		user.POST("/login", controller.UserLoginController)
		user.POST("/userinfo", controller.UserLoginController)
		user.POST("/count/info/:uid", controller.UserGetCountController)
	}
	article := router.Group("/api/article")
	{
		article.POST("/create", controller.ArticleCreateController)
		article.POST("/update", controller.ArticleUpdateController)
		article.GET("/getById/:aid", controller.ArticleGetController)
		article.DELETE("/delete/:aid", controller.ArticleDeleteController)
		article.POST("/list", controller.ArticleGetListController)
		article.POST("/category/list", controller.ArticleCategoryGetListController)
		article.POST("/like", controller.ArticleLikeController)
		article.POST("/isLiked", controller.ArticleIsLikedController)
		article.POST("/view/:aid", controller.ArticleViewController)
		article.POST("/collect", controller.ArticleCollectionController)
		article.POST("/isCollected", controller.ArticleIsCollectionController)
		article.POST("/searchLike", controller.ArticleSearchLikeController)
		article.POST("/getLikeList/:uid", controller.GetLikeListController)
		article.POST("/getCollectionList/:uid", controller.GetCollectionListController)
	}
	topic := router.Group("/api/topic")
	{
		topic.POST("/create", controller.TopicCreateController)
		topic.POST("/update", controller.TopicUpdateController)
		topic.GET("/getById/:tid", controller.TopicGetController)
		topic.DELETE("/delete/:tid", controller.TopicDeleteController)
		topic.POST("/list", controller.TopicGetListController)
	}
	tag := router.Group("/api/tag")
	{
		tag.POST("/create", controller.TagCreateController)
		tag.POST("/update", controller.TagUpdateController)
		tag.GET("/getById/:tid", controller.TagGetController)
		tag.DELETE("/delete/:tid", controller.TagDeleteController)
		tag.POST("/list", controller.TagGetListController)
	}
	upload := router.Group("/api/upload")
	{
		upload.POST("/image", controller.UploadImageController)
		upload.POST("/video", controller.UploadVideoController)
	}
	// 视频
	video := router.Group("/api/video")
	{
		video.POST("/create", controller.VideoCreateController)
		video.POST("/update", controller.VideoUpdateController)
		video.GET("/getById/:vid", controller.VideoGetController)
		video.DELETE("/delete/:vid", controller.VideoDeleteController)
		video.POST("/list", controller.VideoGetListController)
	}
	comment := router.Group("/api/comment")
	{
		comment.POST("/create", controller.CommentCreateController)
		comment.POST("/list", controller.CommentGetListController)
		comment.DELETE("/delete/:cid", controller.CommentDeleteController)
	}
	chat := router.Group("/api/chat")
	{
		chat.GET("/ws", controller.WSHandler)
		chat.GET("/history", controller.ChatHistoryController)
	}
	// 处理 所有的ws请求
	ws := router.Group("api/ws")
	{
		ws.GET("/:uid", controller.WSHandler)
	}
	return router
}
