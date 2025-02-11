package logic

import (
	"errors"
	"fmt"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
)

func ArticleCreate(a *models.Article) (rArticle *models.Article, err error) {
	// 生成aid
	aid, err := snowflake.GetID()
	if err != nil {
		err = errors.New("创建文章ID失败")
		return nil, err
	}
	a.ArticleId = aid
	err = global.Db.Create(&a).Error
	if err != nil {
		return nil, errors.New("创建文章失败")
	}
	return a, nil
}

func ArticleUpdate(a *models.Article) (rArticle *models.Article, err error) {
	fmt.Println(a)
	err = global.Db.Model(&a).Where("article_id = ?", a.ArticleId).Updates(a).Error
	if err != nil {
		return nil, errors.New("更新文章失败")
	}
	return a, nil

}

func ArticleGet(aid int64) (rArticle *models.Article, err error) {
	err = global.Db.Model(&rArticle).Where("article_id =?", aid).First(&rArticle).Error
	if err != nil {
		return nil, errors.New("获取文章失败")
	}
	return rArticle, nil
}

func ArticleDelete(aid int64) (rArticle *models.Article, err error) {
	err = global.Db.Where("article_id = ?", aid).Delete(&rArticle).Error
	if err != nil {
		return nil, errors.New("删除文章失败")
	}
	return nil, nil

}

func ArticleGetList(filter *models.ArticleFilter) (rArticles []*models.Article, err error) {
	if filter.Pagination == nil {
		filter.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := global.Db.Model(&rArticles)
	// 处理过滤
	if filter.CategoryId != 0 {
		query = query.Where("category_id =?", filter.CategoryId)
	}
	if filter.TopicId != 0 {
		query = query.Where("topic_id =?", filter.TopicId)
	}
	if filter.TagId != 0 {
		query = query.Where("tag_id =?", filter.TagId)
	}
	if filter.UserId != "" {
		fmt.Println("userid", filter.UserId)
		query = query.Where("user_id =?", filter.UserId)
	}
	err = query.Error
	offset := filter.Pagination.PageSize * (filter.Pagination.Page - 1)
	query = query.Offset(offset).Limit(filter.Pagination.PageSize)
	query.Find(&rArticles)
	if err != nil {
		return nil, errors.New("获取文章列表失败")
	}
	return rArticles, nil
}

func ArticleCategoryGetList(p *models.Pagination) (rArticleCategories []*models.ArticleCategory, err error) {
	fmt.Println(p, "asd")
	if p.Page == 0 && p.PageSize == 0 {
		p = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	offset := p.PageSize * (p.Page - 1)
	result := global.Db.Limit(p.PageSize).Offset(offset).Find(&rArticleCategories)
	err = result.Error
	if err != nil {
		return nil, errors.New("获取文章分类列表失败")
	}
	return
}
