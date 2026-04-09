package entity

import "time"

type WomanStoreAssignment struct {
	ID        uint       `json:"id"`
	WomanID   uint       `json:"woman_id"`
	StoreID   uint       `json:"store_id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *WomanStoreAssignment) IsNil() bool  { return s.ID == 0 }
func (s *WomanStoreAssignment) GetID() uint   { return s.ID }
