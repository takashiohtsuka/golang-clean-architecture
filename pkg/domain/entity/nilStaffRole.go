package entity

type NilStaffRole struct{}

func (n *NilStaffRole) IsNil() bool     { return true }
func (n *NilStaffRole) GetID() uint      { return 0 }
func (n *NilStaffRole) GetStaffId() uint { return 0 }
func (n *NilStaffRole) GetRoleId() uint  { return 0 }
