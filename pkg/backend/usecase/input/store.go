package input

type CreateStoreInput struct {
	CompanyID        uint
	BusinessTypeCode string
	ContractPlanCode string
	Name             string
	IsActive         bool
	OpenStatus       string
}
