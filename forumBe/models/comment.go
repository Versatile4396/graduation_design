package models

type ArticleComment struct {
	CommentId       uint64  `json:"comment_id"`
	UserId          int64   `json:"user_id"`
	ArticleId       int64   `json:"article_id"`
	Content         string  `json:"content"`
	ParentCommentId *uint64 `json:"parent_comment_id"`
	CreateTime      int64   `json:"create_time" gorm:"-"`
	UpdateTime      int64   `json:"update_time" gorm:"-"`
}

type CommentFilter struct {
	ArticleId  uint64 `json:"article_id"`
	UserId     uint64 `json:"user_id"`
	Pagination *Pagination
}
