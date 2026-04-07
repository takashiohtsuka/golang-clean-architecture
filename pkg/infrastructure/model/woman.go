package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Woman struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CompanyID  uint           `json:"company_id"`
	Name       string         `json:"name"`
	Age        *int           `json:"age"`
	Birthplace *string        `json:"birthplace"`
	BloodType  *string        `json:"blood_type"`
	Hobby      *string        `json:"hobby"`
	IsActive   bool           `json:"is_active"`
	CreatedAt  *time.Time     `json:"created_at"`
	UpdatedAt  *time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (Woman) TableName() string { return "women" }
