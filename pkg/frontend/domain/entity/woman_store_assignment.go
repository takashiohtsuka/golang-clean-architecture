package entity

import "time"

type WomanStoreAssignment struct {
	ID        uint       `json:"id"`
	StoreID   uint       `json:"store_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
