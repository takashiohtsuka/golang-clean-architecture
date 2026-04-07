package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Company struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Rank      *string        `json:"rank"`
	IsActive  bool           `json:"is_active"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Company) TableName() string { return "companies" }
