package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Blog struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	WomanID     uint           `json:"woman_id"`
	Title       string         `json:"title"`
	Body        *string        `json:"body"`
	IsPublished bool           `json:"is_published"`
	CreatedAt   *time.Time     `json:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (Blog) TableName() string { return "blogs" }
