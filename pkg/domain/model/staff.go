package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Staff struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Age       string         `json:"age"`
	IsActive  string         `json:"is_active"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Staff) TableName() string { return "staffs" }
