package models

import "time"

type Article struct {
	ArticleId  uint64    `json:"article_id"`
	Content    string    `json:"content"`
	Title      string    `json:"title"`
	UserId     uint64    `json:"user_id"`
	CategoryId int       `json:"category_id"`
	TopicId    int       `json:"topic_id" `
	TagId      int       `json:"tag_id"`
	Cover      string    `json:"cover"`
	Abstract   string    `json:"abstract"`
	ViewCount  int       `json:"view_count" grom:"default:0"`
	CreateTime time.Time `json:"create_at" gorm:"column:created_at"`
	UpdateTime time.Time `json:"update_at" gorm:"column:updated_at"`
}

type ArticleBrief struct {
	ArticleId    uint64 `json:"article_id"`
	Title        string `json:"title"`
	UserId       uint64 `json:"user_id"`
	UserName     string `json:"username"`
	CategoryId   int64  `json:"category_id"`
	CommentCount int64  `json:"comment_count"`
	LikeCount    int64  `json:"like_count"`
	CollectCount int64  `json:"collect_count"`
}
type ArticleFilter struct {
	ArticleId   uint64 `json:"article_id"`
	Title       string `json:"title"`
	UserId      string `json:"user_id"`
	CategoryId  uint64 `json:"category_id"`
	TopicId     uint64 `json:"topic_id" `
	TagId       uint64 `json:"tag_id"`
	OrderByTime bool   `json:"order_by_time" gorm:"default:0"`
	Pagination  *Pagination
}

type ArticleLike struct {
	ArticleId uint64 `json:"article_id"`
	UserId    uint64 `json:"user_id"`
	IsLike    bool   `json:"is_like" gorm:"-"`
}

type ArticleCollection struct {
	ArticleId    uint64 `json:"article_id"`
	UserId       uint64 `json:"user_id"`
	IsCollection bool   `json:"is_collection" gorm:"-"`
}
