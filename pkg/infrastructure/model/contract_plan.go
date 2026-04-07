package model

/*
GORMでマッピングされたModelの構造体
*/
type ContractPlan struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `json:"code"`
}

func (ContractPlan) TableName() string { return "contract_plans" }
