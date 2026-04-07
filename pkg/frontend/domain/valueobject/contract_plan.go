package valueobject

import "golang-clean-architecture/pkg/domain/valueobject"

type ContractPlan struct {
	valueobject.ValueObject[string]
}

func NewContractPlan(code string) ContractPlan {
	return ContractPlan{valueobject.NewValueObject(code)}
}

func EmptyContractPlan() ContractPlan {
	return ContractPlan{}
}

func (cp ContractPlan) GetCode() string {
	return cp.Get()
}

func (cp ContractPlan) IsEmpty() bool {
	return cp.Get() == ""
}
