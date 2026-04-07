package women

import "golang-clean-architecture/pkg/backend/usecase/input"

type Post struct {
	CompanyID  uint    `json:"company_id"  validate:"required"`
	Name       string  `json:"name"        validate:"required,max=100"`
	Age        *int    `json:"age"         validate:"omitempty,min=18,max=60"`
	Birthplace *string `json:"birthplace"  validate:"omitempty,max=100"`
	BloodType  *string `json:"blood_type"  validate:"omitempty,oneof=A B O AB"`
	Hobby      *string `json:"hobby"       validate:"omitempty,max=200"`
	IsActive   bool    `json:"is_active"`
	StoreIDs   []uint  `json:"store_ids"`
}

func (req *Post) ToInput() input.CreateWomanInput {
	return input.CreateWomanInput{
		CompanyID:  req.CompanyID,
		Name:       req.Name,
		Age:        req.Age,
		Birthplace: req.Birthplace,
		BloodType:  req.BloodType,
		Hobby:      req.Hobby,
		IsActive:   req.IsActive,
		StoreIDs:   req.StoreIDs,
	}
}
