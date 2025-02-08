package models

type Article struct {
	ArticleId  uint64 `json:"article_id"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	UserId     string `json:"user_id"`
	CategoryId int    `json:"category_id"`
	TopicId    int    `json:"topic_id" `
	TagId      int    `json:"tag_id"`
	Cover      string `json:"cover"`
	Abstract   string `json:"abstract"`
}

type ArticleFilter struct {
	ArticleId  uint64 `json:"article_id"`
	Title      string `json:"title"`
	UserId     string `json:"user_id"`
	CategoryId uint64 `json:"category_id"`
	TopicId    uint64 `json:"topic_id" `
	TagId      uint64 `json:"tag_id"`
	Pagination Pagination
}
