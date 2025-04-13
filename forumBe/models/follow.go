package models

import "time"

type Follow struct {
	ID         uint   `gorm:"primaryKey"`
	FollowerID uint64 `json:"follower_id" gorm:"not null"`
	FollowedID uint64 `json:"followed_id" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type FollowFilter struct {
	FollowerID uint64      `json:"follower_id"` // 你关注的
	FollowedID uint64      `json:"followed_id"` // 关注你的
	Pagination *Pagination `json:"pagination"`
}
