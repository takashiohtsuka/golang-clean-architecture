package valueobject

import "golang-clean-architecture/pkg/domain/valueobject"

type BusinessType struct {
	valueobject.ValueObject[string]
}

func NewBusinessType(code string) BusinessType {
	return BusinessType{valueobject.NewValueObject(code)}
}

func EmptyBusinessType() BusinessType {
	return BusinessType{}
}

func (bt BusinessType) GetCode() string {
	return bt.Get()
}

func (bt BusinessType) IsEmpty() bool {
	return bt.Get() == ""
}
