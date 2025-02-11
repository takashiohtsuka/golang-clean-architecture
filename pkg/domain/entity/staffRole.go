package entity

/* staff_roles entityの構造体 */
type StaffRole struct {
	ID      uint `json:"id"`
	StaffId uint `json:"staff_id"`
	RoleId  uint `json:"role_id"`
}
