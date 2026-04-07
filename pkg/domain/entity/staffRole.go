package entity

/* staff_roles entityの構造体 */
type StaffRole struct {
	ID      uint `json:"id"`
	StaffId uint `json:"staff_id"`
	RoleId  uint `json:"role_id"`
}

func (sr *StaffRole) IsNil() bool     { return sr.StaffId == 0 }
func (sr *StaffRole) GetID() uint      { return sr.ID }
func (sr *StaffRole) GetStaffId() uint { return sr.StaffId }
func (sr *StaffRole) GetRoleId() uint  { return sr.RoleId }
