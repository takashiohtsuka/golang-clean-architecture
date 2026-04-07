package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Role struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Role) TableName() string { return "roles" }
