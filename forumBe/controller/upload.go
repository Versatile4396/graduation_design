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
