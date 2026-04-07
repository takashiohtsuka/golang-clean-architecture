package staffs

import "golang-clean-architecture/pkg/usecase/input"

/** request bodyをマッピングする構造体 */
type Post struct {
	Name     string `json:"name"`
	Age      string `json:"age"`
	IsActive string `json:"is_active"`
}

func (req *Post) ToInput() input.CreateStaffInput {
	return input.CreateStaffInput{
		Name:     req.Name,
		Age:      req.Age,
		IsActive: req.IsActive,
	}
}
