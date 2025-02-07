package models

type Image struct {
	ImageUrl  string `json:"image_url"`
	ImageSize int64  `json:"image_size"`
	ImageName string `json:"image_name"`
}
