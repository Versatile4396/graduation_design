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

func ArticleCopy(a *models.Article) (rArticle *models.Article, err error) {
	return
}
