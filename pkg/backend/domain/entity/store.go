package entity

import (
	"time"

	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
)

type OpenStatus string

const (
	OpenStatusPublic     OpenStatus = "public"
	OpenStatusRestricted OpenStatus = "restricted"
	OpenStatusPrivate    OpenStatus = "private"
)

type Store struct {
	ID           uint             `json:"id"`
	CompanyID    uint             `json:"company_id"`
	DistrictID   uint             `json:"district_id"`
	BusinessType bvo.BusinessType `json:"business_type"`
	ContractPlan bvo.ContractPlan `json:"contract_plan"`
	Name         string           `json:"name"`
	IsActive     bool             `json:"is_active"`
	OpenStatus   OpenStatus       `json:"open_status"`
	CreatedAt    *time.Time       `json:"created_at"`
	UpdatedAt    *time.Time       `json:"updated_at"`
	DeletedAt    *time.Time       `json:"deleted_at"`
}

func (s *Store) IsNil() bool                      { return s.ID == 0 }
func (s *Store) GetID() uint                       { return s.ID }
func (s *Store) GetCompanyID() uint                { return s.CompanyID }
func (s *Store) GetDistrictID() uint               { return s.DistrictID }
func (s *Store) GetBusinessType() bvo.BusinessType { return s.BusinessType }
func (s *Store) GetContractPlan() bvo.ContractPlan { return s.ContractPlan }
func (s *Store) GetName() string                   { return s.Name }
func (s *Store) GetIsActive() bool                 { return s.IsActive }
func (s *Store) GetOpenStatus() OpenStatus         { return s.OpenStatus }
