package managementstaffs

import "golang-clean-architecture/pkg/backend/usecase/input"

type Post struct {
	CompanyID uint   `json:"company_id" validate:"required"`
	StoreID   uint   `json:"store_id"   validate:"required"`
	Name      string `json:"name"       validate:"required,max=100"`
	Email     string `json:"email"      validate:"required,email"`
}

func (req *Post) ToInput() input.CreateManagementStaffInput {
	return input.CreateManagementStaffInput{
		CompanyID: req.CompanyID,
		StoreID:   req.StoreID,
		Name:      req.Name,
		Email:     req.Email,
	}
}
