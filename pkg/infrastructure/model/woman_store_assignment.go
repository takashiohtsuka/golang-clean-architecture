package model

import "time"

/*
GORMでマッピングされたModelの構造体
*/
type WomanStoreAssignment struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	WomanID   uint       `json:"woman_id"`
	StoreID   uint       `json:"store_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (WomanStoreAssignment) TableName() string { return "woman_store_assignments" }
