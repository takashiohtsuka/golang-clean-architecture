package stores

import "golang-clean-architecture/pkg/backend/usecase/input"

type Post struct {
	CompanyID        uint   `json:"company_id"          validate:"required"`
	DistrictID       uint   `json:"district_id"         validate:"required"`
	BusinessTypeCode string `json:"business_type_code"  validate:"required"`
	ContractPlanCode string `json:"contract_plan_code"  validate:"required"`
	Name             string `json:"name"                validate:"required,max=100"`
	IsActive         bool   `json:"is_active"`
	OpenStatus       string `json:"open_status"         validate:"required,oneof=open closed"`
}

func (req *Post) ToInput() input.CreateStoreInput {
	return input.CreateStoreInput{
		CompanyID:        req.CompanyID,
		DistrictID:       req.DistrictID,
		BusinessTypeCode: req.BusinessTypeCode,
		ContractPlanCode: req.ContractPlanCode,
		Name:             req.Name,
		IsActive:         req.IsActive,
		OpenStatus:       req.OpenStatus,
	}
}
