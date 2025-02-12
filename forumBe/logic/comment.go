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

func CommentDelete(cid int64) (err error) {
	res := global.Db.Delete(&models.ArticleComment{}, cid)
	err = res.Error
	if err != nil {
		return err
	}
	return nil
}

func CommentGetList(filter *models.CommentFilter) (rComments []*models.ArticleComment, err error) {
	if filter.Pagination == nil {
		filter.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := global.Db.Model(&rComments)
	// 处理过滤
	if filter.ArticleId != 0 {
		query = query.Where("article_id =?", filter.ArticleId)
	}
	if filter.UserId != 0 {
		query = query.Where("user_id =?", filter.UserId)
	}
	err = query.Error
	offset := filter.Pagination.PageSize * (filter.Pagination.Page - 1)
	query = query.Offset(offset).Limit(filter.Pagination.PageSize)
	query.Find(&rComments)
	if err != nil {
		return nil, err
	}
	return rComments, nil

}
