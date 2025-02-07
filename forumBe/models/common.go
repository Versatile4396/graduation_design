package models

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pagesize"`
}
