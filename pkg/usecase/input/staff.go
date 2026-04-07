package input

type ListStaffInput struct {
	Name     string
	Age      string
	IsActive *bool
}

type CreateStaffInput struct {
	Name     string
	Age      string
	IsActive string
}

type UpdateStaffInput struct {
	StaffId  uint
	RoleIds  []uint
	Name     string
	Age      string
	IsActive string
}
