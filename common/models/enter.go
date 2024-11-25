package models

import "time"

type Models struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
	Key   string `form:"key"`
}
