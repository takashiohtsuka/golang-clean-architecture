package model

import "time"

type WomanImage struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	WomanID   uint       `json:"woman_id"`
	Path      string     `json:"path"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (WomanImage) TableName() string { return "woman_images" }
