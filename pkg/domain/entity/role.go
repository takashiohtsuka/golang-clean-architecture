package entity

import "time"

/* role entityの構造体 */
type Role struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (r *Role) IsNil() bool    { return r.ID == 0 }
func (r *Role) GetID() uint     { return r.ID }
func (r *Role) GetName() string { return r.Name }
