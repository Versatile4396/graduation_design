package models

type ArticleCategory struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentId     int    `json:"parent_id"`
}
