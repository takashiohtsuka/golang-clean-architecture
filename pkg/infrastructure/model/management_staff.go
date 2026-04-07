package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type ManagementStaff struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CompanyID uint           `json:"company_id"`
	StoreID   uint           `json:"store_id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	CreatedAt *time.Time     `json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (ManagementStaff) TableName() string { return "management_staffs" }
