package input

type CreateStoreInput struct {
	CompanyID        uint
	DistrictID       uint
	BusinessTypeCode string
	ContractPlanCode string
	Name             string
	IsActive         bool
	OpenStatus       string
}

type UpdateStoreInput struct {
	ID               uint
	DistrictID       uint
	BusinessTypeCode string
	ContractPlanCode string
	Name             string
	IsActive         bool
	OpenStatus       string
}
