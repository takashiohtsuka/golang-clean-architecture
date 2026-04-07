package entity

import (
	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/domain/collection"
)

type CompanyEntity interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetRank() *string
	GetIsActive() bool
}

type StoreEntity interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetBusinessType() bvo.BusinessType
	GetContractPlan() bvo.ContractPlan
	GetName() string
	GetIsActive() bool
	GetOpenStatus() OpenStatus
}

type WomanEntity interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetName() string
	GetAge() *int
	GetBirthplace() *string
	GetBloodType() *string
	GetHobby() *string
	GetIsActive() bool
	GetStoreAssignments() collection.Collection[WomanStoreAssignment]
}

type BlogEntity interface {
	IsNil() bool
	GetID() uint
	GetWomanID() uint
	GetTitle() string
	GetBody() *string
	GetIsPublished() bool
	GetPhotos() collection.Collection[Photo]
}

type ManagementStaffEntity interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetStoreID() uint
	GetName() string
	GetEmail() string
}
