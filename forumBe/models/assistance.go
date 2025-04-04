package models

import "time"

type Assistance struct {
	AssistanceId uint64    `json:"assistance_id"`
	Content      string    `json:"content"`
	Title        string    `json:"title"`
	UserId       uint64    `json:"user_id"`
	CategoryId   int       `json:"category_id"`
	TopicId      int       `json:"topic_id" `
	TagId        int       `json:"tag_id"`
	Cover        string    `json:"cover"`
	Abstract     string    `json:"abstract"`
	CreateTime   time.Time `json:"create_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdateTime   time.Time `json:"update_at" gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
}
type AssistanceFilter struct {
	AssistanceId uint64 `json:"assistance_id"`
	Title        string `json:"title"`
	UserId       uint64 `json:"user_id"`
	CategoryId   uint64 `json:"category_id"`
	TopicId      uint64 `json:"topic_id" `
	TagId        uint64 `json:"tag_id"`
	Status       int    `json:"status"`
	OrderByTime  bool   `json:"order_by_time" gorm:"default:0"`
	Pagination   *Pagination
}
