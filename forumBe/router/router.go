package router

import (
	"bytes"
	"forum/controller"
	"forum/logger"
	"forum/middleware"
	"io"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetEmptyBodyToJSON 是一个Gin中间件，用于在POST请求中，如果请求体为空，则将其设置为 {}
func SetEmptyBodyToJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			// 读取请求体
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
				c.Abort()
				return
			}
			// 关闭原请求体
			c.Request.Body.Close()

			// 如果请求体为空，则设置为 {}
			if len(body) == 0 {
				body = []byte("{}")
			}

			// 重新设置请求体
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		// 继续处理请求
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(SetEmptyBodyToJSON())
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
		user.POST("/list", controller.UserGetListController)
		user.POST("/update", controller.UserUpdateController)
		user.POST("/delete", controller.UserDeleteController)
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
	assistance := router.Group("/api/assistance")
	{
		assistance.POST("/create", controller.AssistanceCreateController)
		assistance.POST("/delete/:aid", controller.AssistanceDeleteController)
		assistance.POST("/list", controller.AssistanceGetListController)
		assistance.POST("/update", controller.AssistanceUpdateController)
	}
	assistanceComment := router.Group("/api/assistanceComment")
	{
		assistanceComment.POST("/create", controller.AssistanceCommentCreateController)
		assistanceComment.POST("/delete/:cid", controller.AssistanceCommentDeleteController)
		assistanceComment.POST("/list", controller.AssistanceCommentGetListController)
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
