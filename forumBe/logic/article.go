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

func ArticleGet(aid uint64) (rArticle *models.Article, userInfo *models.User, rABrief *models.ArticleBrief, err error) {
	err = global.Db.Model(&rArticle).Where("article_id =?", aid).First(&rArticle).Error
	if err != nil {
		err = errors.New("获取文章失败")
		return
	}
	err = global.Db.Model(&userInfo).Where("user_id =?", rArticle.UserId).First(&userInfo).Error
	res := global.Db.Model(&models.Article{}).Where("article_id =?", aid).Update("view_count", global.Db.Raw("view_count + 1"))
	err = res.Error
	if err != nil {
		err = errors.New("获取文章作者信息失败")
		return
	}
	ArticleBrief, err := ArticleGetBrief(aid, rArticle.UserId)
	return rArticle, userInfo, ArticleBrief, nil
}

func ArticleDelete(aid int64) (rArticle *models.Article, err error) {
	err = global.Db.Where("article_id = ?", aid).Delete(&rArticle).Error
	if err != nil {
		return nil, errors.New("删除文章失败")
	}
	return nil, nil

}

func ArticleGetList(filter *models.ArticleFilter) (rArticles []*models.Article, rArticleBriefs []*models.ArticleBrief, err error) {
	if filter.Pagination == nil {
		filter.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	query := global.Db.Model(&rArticles)
	// 处理过滤
	if err := HandleFilterInfo(query, filter); err != nil {
		return nil, nil, err
	}
	query.Find(&rArticles)
	if err != nil {
		return nil, nil, errors.New("获取文章列表失败")
	}
	// 获取文章列表 点赞 评论数量
	for _, article := range rArticles {
		ArticleBrief, err := ArticleGetBrief(article.ArticleId, article.UserId)
		if err != nil {
			return nil, nil, err
		}
		rArticleBriefs = append(rArticleBriefs, ArticleBrief)
	}
	return
}

func HandleFilterInfo(query *gorm.DB, filter *models.ArticleFilter) (err error) {
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
	orderDir := "desc"
	if filter.OrderByTime {
		orderDir = "asc"
	}
	offset := filter.Pagination.PageSize * (filter.Pagination.Page - 1)
	query = query.Order(fmt.Sprintf("created_at %s", orderDir)).Offset(offset).Limit(filter.Pagination.PageSize)
	return query.Error
}

func ArticleGetBrief(aid uint64, uid uint64) (rArticleBrief *models.ArticleBrief, err error) {
	var commentCount int64
	var likeCount int64
	var collectionCount int64
	err = global.Db.Model(&models.ArticleComment{}).Where("article_id=?", aid).Count(&commentCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("查询数量时出错:", err)
		return
	}
	err = global.Db.Model(&models.ArticleLike{}).Where("article_id=?", aid).Count(&likeCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("查询数量时出错:", err)
		return
	}
	err = global.Db.Model(&models.ArticleCollection{}).Where("article_id=?", aid).Count(&collectionCount).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("查询数量时出错:", err)
		return
	}
	var u *models.User
	err = global.Db.Model(&u).Where("user_id =?", uid).First(&u).Error
	if err != nil {
		return nil, err
	}
	rArticleBrief = &models.ArticleBrief{
		ArticleId:    aid,
		CommentCount: commentCount,
		LikeCount:    likeCount,
		CollectCount: collectionCount,
		UserName:     u.UserName,
	}
	return rArticleBrief, nil
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
		// 查询是否已经点赞
		var like *models.ArticleLike
		res = global.Db.Model(&models.ArticleLike{}).Where("article_id =? and user_id = ?", l.ArticleId, l.UserId).First(&like)
		err = res.Error
		if err == gorm.ErrRecordNotFound {
			// 没有点赞过，创建点赞记录
			res = global.Db.Create(l)
			err = res.Error
			if err != nil {
				return errors.New("点赞失败")
			}
		} else if err != nil {
			// 其他查询错误
			return errors.New("查询点赞记录失败")
		}
		return nil
	}
	err = res.Error
	if err != nil {
		return errors.New("点赞业务处理失败")
	}
	return nil
}

func ArticleIsLiked(l *models.ArticleLike) (isLiked bool, err error) {
	var res *gorm.DB
	res = global.Db.Model(&models.ArticleLike{}).Where("article_id =? and user_id =?", l.ArticleId, l.UserId).First(&models.ArticleLike{})
	err = res.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("查询点赞状态失败")
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}

func ArticleView(aid string) (err error) {
	res := global.Db.Model(&models.Article{}).Where("article_id =?", aid).Update("view_count", global.Db.Raw("view_count + 1"))
	err = res.Error
	if err != nil {
		return errors.New("浏览失败")
	}
	return nil
}

func ArticleCollection(c *models.ArticleCollection) (err error) {
	var res *gorm.DB
	if !c.IsCollection {
		res = global.Db.Model(&models.ArticleCollection{}).Where("article_id =? and user_id =?", c.ArticleId, c.UserId).Delete(&models.ArticleCollection{})
		err = res.Error
		if err != nil {
			return errors.New("取消收藏失败")
		}
		return nil
	} else {
		// 查询是否已经收藏
		var collection *models.ArticleCollection
		res = global.Db.Model(&models.ArticleCollection{}).Where("article_id =? and user_id =?", c.ArticleId, c.UserId).First(&collection)
		err = res.Error
		if err == gorm.ErrRecordNotFound {
			// 没有收藏过，创建收藏记录
			res = global.Db.Create(c)
			err = res.Error
			if err != nil {
				return errors.New("收藏失败")
			}
		} else if err != nil {
			// 其他查询错误
			return errors.New("查询收藏记录失败")
		}
		return nil
	}
}

func ArticleIsCollection(c *models.ArticleCollection) (isCollection bool, err error) {
	var res *gorm.DB
	res = global.Db.Model(&models.ArticleCollection{}).Where("article_id =? and user_id =?", c.ArticleId, c.UserId).First(&models.ArticleCollection{})
	err = res.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, errors.New("查询收藏状态失败")
	}
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return true, nil
}

func ArticleSearchLike(s *models.ArticleSearch) (rArticles []*models.Article, rABrief []*models.ArticleBrief, err error) {
	if s.Pagination == nil {
		s.Pagination = &models.Pagination{
			Page:     1,
			PageSize: 10,
		}
	}
	offset := s.Pagination.PageSize * (s.Pagination.Page - 1)
	err = global.Db.Model(&rArticles).Where("title LIKE ? OR content LIKE ?", "%"+s.Keyword+"%", "%"+s.Keyword+"%").Offset(offset).Limit(s.Pagination.PageSize).Find(&rArticles).Error
	if s.Title != "" {
		err = global.Db.Model(&rArticles).Where("title like ?", "%"+s.Title+"%").Offset(offset).Limit(s.Pagination.PageSize).Find(&rArticles).Error
	}
	if s.Content != "" {
		err = global.Db.Model(&rArticles).Where("content like?", "%"+s.Content+"%").Offset(offset).Limit(s.Pagination.PageSize).Find(&rArticles).Error
	}
	if err != nil {
		return nil, nil, errors.New("搜索文章失败")
	}
	for _, article := range rArticles {
		ArticleBrief, err := ArticleGetBrief(article.ArticleId, article.UserId)
		if err != nil {
			return nil, nil, err
		}
		rABrief = append(rABrief, ArticleBrief)
	}
	return
}

func GetLikeList(uid uint64) (rArticles []*models.Article, rABrief []*models.ArticleBrief, err error) {
	var rLikeList []*models.ArticleLike
	err = global.Db.Model(&rLikeList).Where("user_id =?", uid).Find(&rLikeList).Error
	if err != nil {
		return nil, nil, errors.New("获取点赞列表失败")
	}
	for _, like := range rLikeList {
		var article *models.Article
		err = global.Db.Model(&article).Where("article_id =?", like.ArticleId).First(&article).Error
		if err != nil {
			return nil, nil, errors.New("获取文章失败")
		}
		rArticles = append(rArticles, article)
		ArticleBrief, err := ArticleGetBrief(article.ArticleId, article.UserId)
		if err != nil {
			return nil, nil, err
		}
		rABrief = append(rABrief, ArticleBrief)
	}
	return
}

func GetCollectionList(uid uint64) (rArticles []*models.Article, rABrief []*models.ArticleBrief, err error) {
	var rCollectionList []*models.ArticleCollection
	err = global.Db.Model(&rCollectionList).Where("user_id =?", uid).Find(&rCollectionList).Error
	if err != nil {
		return nil, nil, errors.New("获取收藏列表失败")
	}
	for _, collection := range rCollectionList {
		var article *models.Article
		err = global.Db.Model(&article).Where("article_id =?", collection.ArticleId).First(&article).Error
		if err != nil {
			return nil, nil, errors.New("获取文章失败")
		}
		rArticles = append(rArticles, article)
		ArticleBrief, err := ArticleGetBrief(article.ArticleId, article.UserId)
		if err != nil {
			return nil, nil, err
		}
		rABrief = append(rABrief, ArticleBrief)
	}
	return
}
