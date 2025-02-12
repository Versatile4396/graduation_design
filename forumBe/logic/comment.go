package logic

import (
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

func CommentCreate(Comment *models.ArticleComment) (rComment *models.ArticleComment, err error) {

	// 生成cid
	cid, err := snowflake.GetID()
	Comment.CommentId = cid
	if err != nil {
		return nil, err
	}
	res := global.Db.Create(&Comment)
	err = res.Error
	if err != nil {
		return nil, err
	}
	return Comment, nil
}
