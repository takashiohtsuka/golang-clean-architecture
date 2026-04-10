package entity

import (
	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/domain/collection"
)

type NilCompany struct{}

func (n *NilCompany) IsNil() bool      { return true }
func (n *NilCompany) GetID() uint       { return 0 }
func (n *NilCompany) GetName() string   { return "" }
func (n *NilCompany) GetRank() *string  { return nil }
func (n *NilCompany) GetIsActive() bool { return false }

type NilStore struct{}

func (n *NilStore) IsNil() bool                      { return true }
func (n *NilStore) GetID() uint                       { return 0 }
func (n *NilStore) GetCompanyID() uint                { return 0 }
func (n *NilStore) GetDistrictID() uint               { return 0 }
func (n *NilStore) GetBusinessType() bvo.BusinessType { return bvo.EmptyBusinessType() }
func (n *NilStore) GetContractPlan() bvo.ContractPlan { return bvo.EmptyContractPlan() }
func (n *NilStore) GetName() string                   { return "" }
func (n *NilStore) GetIsActive() bool                 { return false }
func (n *NilStore) GetOpenStatus() OpenStatus         { return "" }

type NilWoman struct{}

func (n *NilWoman) IsNil() bool                                                    { return true }
func (n *NilWoman) GetID() uint                                                     { return 0 }
func (n *NilWoman) GetCompanyID() uint                                              { return 0 }
func (n *NilWoman) GetName() string                                                 { return "" }
func (n *NilWoman) GetAge() *int                                                    { return nil }
func (n *NilWoman) GetBirthplace() *string                                          { return nil }
func (n *NilWoman) GetBloodType() *string                                           { return nil }
func (n *NilWoman) GetHobby() *string                                               { return nil }
func (n *NilWoman) GetIsActive() bool                                    { return false }
func (n *NilWoman) GetImages() collection.Collection[WomanImage] {
	return collection.NewCollection[WomanImage](nil)
}

type NilBlog struct{}

func (n *NilBlog) IsNil() bool                                 { return true }
func (n *NilBlog) GetID() uint                                  { return 0 }
func (n *NilBlog) GetWomanID() uint                             { return 0 }
func (n *NilBlog) GetTitle() string                             { return "" }
func (n *NilBlog) GetBody() *string                             { return nil }
func (n *NilBlog) GetIsPublished() bool                         { return false }
func (n *NilBlog) GetPhotos() collection.Collection[Photo]      { return collection.NewCollection[Photo](nil) }

type NilManagementStaff struct{}

func (n *NilManagementStaff) IsNil() bool        { return true }
func (n *NilManagementStaff) GetID() uint         { return 0 }
func (n *NilManagementStaff) GetCompanyID() uint  { return 0 }
func (n *NilManagementStaff) GetStoreID() uint    { return 0 }
func (n *NilManagementStaff) GetName() string     { return "" }
func (n *NilManagementStaff) GetEmail() string    { return "" }
