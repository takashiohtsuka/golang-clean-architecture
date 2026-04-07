package entity

import "time"

type Company struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Rank      *string    `json:"rank"`
	IsActive  bool       `json:"is_active"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (c *Company) IsNil() bool      { return c.ID == 0 }
func (c *Company) GetID() uint       { return c.ID }
func (c *Company) GetName() string   { return c.Name }
func (c *Company) GetRank() *string  { return c.Rank }
func (c *Company) GetIsActive() bool { return c.IsActive }
