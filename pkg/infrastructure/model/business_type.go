package model

/*
GORMでマッピングされたModelの構造体
*/
type BusinessType struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `json:"code"`
}

func (BusinessType) TableName() string { return "business_types" }
