package logic

import (
	"errors"
	"fmt"
	"forum/global"
	"forum/models"
)

func TopicCreate(t *models.Topic) (rTopic *models.Topic, err error) {
	err = global.Db.Create(&t).Error
	if err != nil {
		return nil, errors.New("创建话题失败")
	}
	return
}

func TopicGetList(p *models.Pagination) (rTopics []*models.Topic, err error) {
	var topics []*models.Topic
	offset := p.PageSize * (p.Page - 1)
	result := global.Db.Model(&models.Topic{}).Find(&topics)
	fmt.Println(offset, p.Page, p.PageSize, "pagesize", result)
	err = result.Error
	if err != nil {
		return nil, errors.New("获取话题列表失败")
	}
	rTopics = topics
	return
}
