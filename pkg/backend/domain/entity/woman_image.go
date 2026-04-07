package entity

import "time"

type WomanImage struct {
	ID        uint       `json:"id"`
	WomanID   uint       `json:"woman_id"`
	Path      string     `json:"path"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
