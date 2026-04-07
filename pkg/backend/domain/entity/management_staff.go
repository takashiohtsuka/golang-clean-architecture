package entity

import "time"

type ManagementStaff struct {
	ID        uint       `json:"id"`
	CompanyID uint       `json:"company_id"`
	StoreID   uint       `json:"store_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (m *ManagementStaff) IsNil() bool        { return m.ID == 0 }
func (m *ManagementStaff) GetID() uint         { return m.ID }
func (m *ManagementStaff) GetCompanyID() uint  { return m.CompanyID }
func (m *ManagementStaff) GetStoreID() uint    { return m.StoreID }
func (m *ManagementStaff) GetName() string     { return m.Name }
func (m *ManagementStaff) GetEmail() string    { return m.Email }
