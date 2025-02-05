package models

type Article struct {
	ArticleId  uint64 `json:"article_id"`
	Content    string `json:"content"`
	Title      string `json:"title"`
	UserId     string `json:"user_id"`
	CategoryId int    `json:"category_id"`
	TopicTitle int    `json:"topic_title" `
	Cover      string `json:"cover"`
}
