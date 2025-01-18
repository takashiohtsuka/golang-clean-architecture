package entity

import "time"

/* staff entityの構造体 */
type User struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Age       string     `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
