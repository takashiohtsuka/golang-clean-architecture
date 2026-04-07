package staffs

import "golang-clean-architecture/pkg/usecase/input"

/** request bodyをマッピングする構造体 */
type Put struct {
	StaffId  uint   `json:"staff_id"`
	RoleIds  []uint `json:"role_ids"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	IsActive string `json:"is_active"`
}

func (req *Put) ToInput() input.UpdateStaffInput {
	return input.UpdateStaffInput{
		StaffId:  req.StaffId,
		RoleIds:  req.RoleIds,
		Name:     req.Name,
		Age:      req.Age,
		IsActive: req.IsActive,
	}
}
