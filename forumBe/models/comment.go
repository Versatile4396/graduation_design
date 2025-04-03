package models

import "time"

type ArticleComment struct {
	CommentId       uint64     `json:"comment_id"`
	UserId          int64      `json:"user_id"`
	ArticleId       int64      `json:"article_id"`
	Content         string     `json:"content" binding:"required"`
	ParentCommentId *uint64    `json:"parent_comment_id"`
	UserInfo        *UserInfo  `json:"user_info" gorm:"-"`
	CreateTime      *time.Time `json:"create_time" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime      *time.Time `json:"update_time" gorm:"default:CURRENT_TIMESTAMP"`
}

type CommentFilter struct {
	ArticleId  uint64      `json:"article_id"`
	UserId     uint64      `json:"user_id"`
	Pagination *Pagination `json:"pagination"`
}

type AssistanceComment struct {
	CommentId       uint64     `json:"comment_id"`
	UserId          int64      `json:"user_id"`
	ArticleId       int64      `json:"article_id"`
	Content         string     `json:"content" binding:"required"`
	ParentCommentId *uint64    `json:"parent_comment_id"`
	UserInfo        *UserInfo  `json:"user_info" gorm:"-"`
	CreateTime      *time.Time `json:"create_time" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime      *time.Time `json:"update_time" gorm:"default:CURRENT_TIMESTAMP"`
}
