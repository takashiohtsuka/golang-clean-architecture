package model

import (
	"time"

	"gorm.io/gorm"
)

/*
GORMでマッピングされたModelの構造体
*/
type Store struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CompanyID      uint           `json:"company_id"`
	DistrictID     uint           `json:"district_id"`
	BusinessTypeID uint           `json:"business_type_id"`
	BusinessType   BusinessType   `json:"business_type"`
	ContractPlanID uint           `json:"contract_plan_id"`
	ContractPlan   ContractPlan   `json:"contract_plan"`
	Name           string         `json:"name"`
	IsActive       bool           `json:"is_active"`
	OpenStatus     string         `json:"open_status"`
	CreatedAt      *time.Time     `json:"created_at"`
	UpdatedAt      *time.Time     `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}

func (Store) TableName() string { return "stores" }
