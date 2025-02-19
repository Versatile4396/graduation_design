package logic

import (
	"errors"
	"fmt"
	"forum/global"
	"forum/models"
	"forum/pkg/snowflake"
	"time"

	"gorm.io/gorm"
)

func ArticleCreate(a *models.Article) (rArticle *models.Article, err error) {
	// 生成aid
	aid, err := snowflake.GetID()
	if err != nil {
		err = errors.New("创建文章ID失败")
		return nil, err
	}
	now := time.Now()
	a.ArticleId = aid
	a.CreateTime = now
	a.UpdateTime = now
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

func ArticleGet(aid int64) (rArticle *models.Article, userInfo *models.User, err error) {
	err = global.Db.Model(&rArticle).Where("article_id =?", aid).First(&rArticle).Error
	if err != nil {
		return nil, nil, errors.New("获取文章失败")
	}
	err = global.Db.Model(&userInfo).Where("user_id =?", rArticle.UserId).First(&userInfo).Error
	res := global.Db.Model(&models.Article{}).Where("article_id =?", aid).Update("view_count", global.Db.Raw("view_count + 1"))
	err = res.Error
	if err != nil {
		return nil, nil, errors.New("获取文章作者信息失败")
	}
	return rArticle, userInfo, nil
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
	orderDir := "desc"
	if filter.OrderByTime {
		orderDir = "asc"
	}
	offset := filter.Pagination.PageSize * (filter.Pagination.Page - 1)
	query = query.Order(fmt.Sprintf("created_at %s", orderDir)).Offset(offset).Limit(filter.Pagination.PageSize)
	query.Find(&rArticles)
	if err != nil {
		return nil, errors.New("获取文章列表失败")
	}
	return rArticles, nil
}

func ArticleCategoryGetList(p *models.Pagination) (rArticleCategories []*models.ArticleCategory, err error) {
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

func ArticleLike(l *models.ArticleLike) (err error) {
	var res *gorm.DB
	if !l.IsLike {
		res = global.Db.Model(&models.ArticleLike{}).Where("article_id =? and user_id =?", l.ArticleId, l.UserId).Delete(&models.ArticleLike{})
		err = res.Error
		if err != nil {
			return errors.New("取消点赞失败")
		}
		return nil
	} else {
		res = global.Db.Create(&l)
	}
	err = res.Error
	if err != nil {
		return errors.New("点赞业务处理失败")
	}
	return nil
}

func ArticleView(aid string) (err error) {
	res := global.Db.Model(&models.Article{}).Where("article_id =?", aid).Update("view_count", global.Db.Raw("view_count + 1"))
	err = res.Error
	if err != nil {
		return errors.New("浏览失败")
	}
	return nil
}
