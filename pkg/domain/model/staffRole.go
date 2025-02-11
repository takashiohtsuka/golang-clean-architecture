package model

/*
GORMでマッピングされたModelの構造体
*/
type StaffRole struct {
	ID      uint `gorm:"primary_key" json:"id"`
	StaffId uint `json:"staff_id"`
	RoleId  uint `json:"role_id"`
}

func (StaffRole) TableName() string { return "staff_roles" }
