package entity

import "time"

/* user entityの構造体 */
type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       string     `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (u *User) IsNil() bool    { return u.ID == 0 }
func (u *User) GetID() uint     { return u.ID }
func (u *User) GetName() string { return u.Name }
func (u *User) GetAge() string  { return u.Age }
