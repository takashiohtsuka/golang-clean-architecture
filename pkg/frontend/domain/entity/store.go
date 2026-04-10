package entity

import (
	"time"

	"golang-clean-architecture/pkg/domain/collection"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
)

type OpenStatus string

const (
	OpenStatusOpen   OpenStatus = "open"
	OpenStatusClosed OpenStatus = "closed"
)

type Store struct {
	ID           uint                               `json:"id"`
	CompanyID    uint                               `json:"company_id"`
	District     fvo.District                       `json:"district"`
	Prefecture   fvo.Prefecture                     `json:"prefecture"`
	Region       fvo.Region                         `json:"region"`
	BusinessType fvo.BusinessType                   `json:"business_type"`
	ContractPlan fvo.ContractPlan                   `json:"contract_plan"`
	Name         string                             `json:"name"`
	IsActive     bool                               `json:"is_active"`
	OpenStatus   OpenStatus                         `json:"open_status"`
	Women        collection.Collection[WomanEntity] `json:"women"`
	CreatedAt    *time.Time                         `json:"created_at"`
	UpdatedAt    *time.Time                         `json:"updated_at"`
	DeletedAt    *time.Time                         `json:"deleted_at"`
}

func (s *Store) IsNil() bool                                    { return s.ID == 0 }
func (s *Store) GetID() uint                                     { return s.ID }
func (s *Store) GetCompanyID() uint                              { return s.CompanyID }
func (s *Store) GetDistrict() fvo.District                       { return s.District }
func (s *Store) GetPrefecture() fvo.Prefecture                   { return s.Prefecture }
func (s *Store) GetRegion() fvo.Region                           { return s.Region }
func (s *Store) GetBusinessType() fvo.BusinessType               { return s.BusinessType }
func (s *Store) GetContractPlan() fvo.ContractPlan               { return s.ContractPlan }
func (s *Store) GetName() string                                 { return s.Name }
func (s *Store) GetIsActive() bool                               { return s.IsActive }
func (s *Store) GetOpenStatus() OpenStatus                       { return s.OpenStatus }
func (s *Store) GetWomen() collection.Collection[WomanEntity]    { return s.Women }
