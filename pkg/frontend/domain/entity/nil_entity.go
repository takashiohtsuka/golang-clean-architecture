package entity

import (
	"golang-clean-architecture/pkg/domain/collection"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
)

type NilStore struct{}

func (n *NilStore) IsNil() bool                                   { return true }
func (n *NilStore) GetID() uint                                    { return 0 }
func (n *NilStore) GetCompanyID() uint                             { return 0 }
func (n *NilStore) GetDistrict() fvo.District                      { return fvo.EmptyDistrict() }
func (n *NilStore) GetPrefecture() fvo.Prefecture                  { return fvo.EmptyPrefecture() }
func (n *NilStore) GetRegion() fvo.Region                          { return fvo.EmptyRegion() }
func (n *NilStore) GetBusinessType() fvo.BusinessType              { return fvo.EmptyBusinessType() }
func (n *NilStore) GetContractPlan() fvo.ContractPlan              { return fvo.EmptyContractPlan() }
func (n *NilStore) GetName() string                                { return "" }
func (n *NilStore) GetIsActive() bool                              { return false }
func (n *NilStore) GetOpenStatus() OpenStatus                      { return "" }
func (n *NilStore) GetWomen() collection.Collection[WomanEntity]   {
	return collection.NewCollection[WomanEntity](nil)
}

type NilWoman struct{}

func (n *NilWoman) IsNil() bool                  { return true }
func (n *NilWoman) GetID() uint                   { return 0 }
func (n *NilWoman) GetCompanyID() uint            { return 0 }
func (n *NilWoman) GetDistrict() fvo.District     { return fvo.EmptyDistrict() }
func (n *NilWoman) GetPrefecture() fvo.Prefecture { return fvo.EmptyPrefecture() }
func (n *NilWoman) GetRegion() fvo.Region         { return fvo.EmptyRegion() }
func (n *NilWoman) GetName() string               { return "" }
func (n *NilWoman) GetAge() *int                  { return nil }
func (n *NilWoman) GetBirthplace() *string        { return nil }
func (n *NilWoman) GetBloodType() *string         { return nil }
func (n *NilWoman) GetHobby() *string             { return nil }
func (n *NilWoman) GetIsActive() bool             { return false }
func (n *NilWoman) GetStoreAssignments() collection.Collection[WomanStoreAssignment] {
	return collection.NewCollection[WomanStoreAssignment](nil)
}
func (n *NilWoman) GetImages() collection.Collection[WomanImage] {
	return collection.NewCollection[WomanImage](nil)
}
func (n *NilWoman) GetBlogs() collection.Collection[BlogEntity] {
	return collection.NewCollection[BlogEntity](nil)
}

type NilBlog struct{}

func (n *NilBlog) IsNil() bool                              { return true }
func (n *NilBlog) GetID() uint                               { return 0 }
func (n *NilBlog) GetWomanID() uint                          { return 0 }
func (n *NilBlog) GetTitle() string                          { return "" }
func (n *NilBlog) GetBody() *string                          { return nil }
func (n *NilBlog) GetIsPublished() bool                      { return false }
func (n *NilBlog) GetPhotos() collection.Collection[Photo]   {
	return collection.NewCollection[Photo](nil)
}
