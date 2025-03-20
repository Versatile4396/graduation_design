package models

import "time"

type Video struct {
	VideoId    string    `json:"video_id" gorm:"-"`
	UserId     uint64    `json:"user_id"`
	Title      string    `json:"title"`
	Abstract   string    `json:"abstract"`
	UploadTime time.Time `json:"upload_time" gorm:"-"`
	Duration   int64     `json:"duration"`
	Cover      string    `json:"cover"`
	TagIds     string    `json:"tag_ids"`
	Status     string    `json:"status" gorm:"-"`
	Category   string    `json:"category"`
	VideoUrl   string    `json:"video_url"`
	VideoSize  int64     `json:"video_size"`
}
