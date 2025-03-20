package controller

import (
	"forum/logic"
	"forum/models"

	"github.com/gin-gonic/gin"
)

func VideoCreateController(c *gin.Context) {
	var video *models.Video
	if err := c.ShouldBindJSON(&video); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, "参数传递错误")
		return
	}
	rVideo, err := logic.VideoCreate(video)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, rVideo)
}

func VideoUpdateController(c *gin.Context) {

}

func VideoGetController(c *gin.Context) {

}

func VideoDeleteController(c *gin.Context) {

}

func VideoGetListController(c *gin.Context) {

}
