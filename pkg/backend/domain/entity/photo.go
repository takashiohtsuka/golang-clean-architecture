package entity

import "time"

type Photo struct {
	ID        uint       `json:"id"`
	BlogID    uint       `json:"blog_id"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (p *Photo) IsNil() bool  { return p.ID == 0 }
func (p *Photo) GetID() uint   { return p.ID }
