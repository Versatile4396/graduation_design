package logic

import (
	"forum/global"
	"forum/models"
)

func VideoCreate(video *models.Video) (rVideo *models.Video, err error) {
	// 生成vid
	err = global.Db.Model(&models.Video{}).Create(&video).Error
	if err != nil {
		return nil, err
	}
	rVideo = video
	return
}
