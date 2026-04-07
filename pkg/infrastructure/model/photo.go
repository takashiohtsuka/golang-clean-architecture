package model

import "time"

/*
GORMでマッピングされたModelの構造体
*/
type Photo struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	BlogID    uint       `json:"blog_id"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Photo) TableName() string { return "photos" }
