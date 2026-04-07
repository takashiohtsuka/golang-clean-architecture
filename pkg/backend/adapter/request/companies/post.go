package companies

import "golang-clean-architecture/pkg/backend/usecase/input"

type Post struct {
	Name     string  `json:"name"     validate:"required,max=100"`
	Rank     *string `json:"rank"     validate:"omitempty,max=50"`
	IsActive bool    `json:"is_active"`
}

func (req *Post) ToInput() input.CreateCompanyInput {
	return input.CreateCompanyInput{
		Name:     req.Name,
		Rank:     req.Rank,
		IsActive: req.IsActive,
	}
}
