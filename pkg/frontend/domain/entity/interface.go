package entity

import (
	"golang-clean-architecture/pkg/domain/collection"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
)

type StoreEntity interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetDistrict() fvo.District
	GetPrefecture() fvo.Prefecture
	GetRegion() fvo.Region
	GetBusinessType() fvo.BusinessType
	GetContractPlan() fvo.ContractPlan
	GetName() string
	GetIsActive() bool
	GetOpenStatus() OpenStatus
	GetWomen() collection.Collection[WomanEntity]
}

type WomanEntity interface {
	IsNil() bool
	GetID() uint
	GetCompanyID() uint
	GetDistrict() fvo.District
	GetPrefecture() fvo.Prefecture
	GetRegion() fvo.Region
	GetName() string
	GetAge() *int
	GetBirthplace() *string
	GetBloodType() *string
	GetHobby() *string
	GetIsActive() bool
	GetStoreAssignments() collection.Collection[WomanStoreAssignment]
	GetImages() collection.Collection[WomanImage]
	GetBlogs() collection.Collection[BlogEntity]
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
