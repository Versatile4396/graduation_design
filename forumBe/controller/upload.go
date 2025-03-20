package controller

import (
	"forum/config"
	"forum/models"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func UploadImageController(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 限制图片后缀
	ext := path.Ext(file.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		ResponseError(c, CodeInvalidExt)
		return
	}
	// 限制图片大小
	if file.Size > 1024*1024*5 {
		ResponseError(c, CodeInvalidSize)
		return
	}
	const seprator = string(os.PathSeparator)
	filePath := path.Join(config.AppConf.AppBaseConfig.ImageDir + file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	image := &models.Image{
		ImageUrl:  "http://127.0.0.1" + config.AppConf.AppBaseConfig.Port + seprator + filePath,
		ImageSize: file.Size,
		ImageName: file.Filename,
	}
	image.ImageName = file.Filename
	ResponseSuccess(c, image)
}

func UploadVideoController(c *gin.Context) {
	video_file, err := c.FormFile("video_file")
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 限制视频后缀
	ext := path.Ext(video_file.Filename)
	if ext != ".mp4" {
		ResponseError(c, CodeInvalidExt)
		return
	}
	// 限制视频大小
	if video_file.Size > 1024*1024*50 {
		ResponseError(c, CodeInvalidSize)
		return
	}
	const seprator = string(os.PathSeparator)
	videoPath := path.Join(config.AppConf.AppBaseConfig.VideoDir + video_file.Filename)
	err = c.SaveUploadedFile(video_file, videoPath)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	video := &models.TVideo{
		VideoUrl:  "http://127.0.0.1" + config.AppConf.AppBaseConfig.Port + seprator + videoPath,
		VideoSize: video_file.Size,
		VideoName: video_file.Filename,
	}
	ResponseSuccess(c, video)
}
