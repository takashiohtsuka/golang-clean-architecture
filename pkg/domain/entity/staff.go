package entity

import (
	"time"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/valueobject"
)

/* staff entityの構造体 */
type Staff struct {
	ID        uint                                    `json:"id"`
	Name      string                                  `json:"name"`
	Age       string                                  `json:"age"`
	IsActive  string                                  `json:"is_active"`
	Roles     collection.Collection[valueobject.Role] `json:"roles"`
	CreatedAt *time.Time                              `json:"created_at"`
	UpdatedAt *time.Time                              `json:"updated_at"`
	DeletedAt *time.Time                              `json:"deleted_at"`
}

func (s *Staff) IsNil() bool                                        { return s.ID == 0 }
func (s *Staff) GetID() uint                                         { return s.ID }
func (s *Staff) GetName() string                                     { return s.Name }
func (s *Staff) GetAge() string                                      { return s.Age }
func (s *Staff) GetIsActive() string                                 { return s.IsActive }
func (s *Staff) GetRoles() collection.Collection[valueobject.Role]   { return s.Roles }
