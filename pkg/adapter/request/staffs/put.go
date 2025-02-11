package staffs

/** request bodyをマッピングする構造体 */
type Put struct {
	StaffId uint   `json:"staff_id"`
	RoleId  uint   `json:"role_id"`
	Name    string `json:"name"`
}
