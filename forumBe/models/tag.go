package models

type Tag struct {
	TagId   int    `json:"tag_id"`
	TagName string `json:"tag_name" validate:"required"`
	TagDesc string `json:"tag_desc" validate:"required"`
}
