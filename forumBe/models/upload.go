package models

type Image struct {
	ImageUrl  string `json:"image_url"`
	ImageSize int64  `json:"image_size"`
	ImageName string `json:"image_name"`
}
type TVideo struct {
	VideoUrl  string `json:"video_url"`
	VideoSize int64  `json:"video_size"`
	VideoName string `json:"video_name"`
}
