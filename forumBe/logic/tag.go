package logic

import (
	"errors"
	"forum/global"
	"forum/models"
)

func TagCreate(t *models.Tag) (rTag *models.Tag, err error) {
	err = global.Db.Create(&t).Error
	if err != nil {
		return nil, errors.New("创建标签失败")
	}
	return t, nil
}

func TagGetList(p *models.Pagination) (rTags []*models.Tag, err error) {
	var tags []*models.Tag
	offset := p.PageSize * (p.Page - 1)
	result := global.Db.Limit(p.PageSize).Offset(offset).Find(&tags)
	err = result.Error
	if err != nil {
		return nil, errors.New("获取标签列表失败")
	}
	rTags = tags
	return
}
