package entity

import (
	"time"

	"golang-clean-architecture/pkg/domain/collection"
)

type Woman struct {
	ID               uint                                      `json:"id"`
	CompanyID        uint                                      `json:"company_id"`
	Name             string                                    `json:"name"`
	Age              *int                                      `json:"age"`
	Birthplace       *string                                   `json:"birthplace"`
	BloodType        *string                                   `json:"blood_type"`
	Hobby            *string                                   `json:"hobby"`
	IsActive         bool                                      `json:"is_active"`
	StoreAssignments collection.Collection[WomanStoreAssignment] `json:"store_assignments"`
	Images           collection.Collection[WomanImage]          `json:"images"`
	Blogs            collection.Collection[BlogEntity]          `json:"blogs"`
	CreatedAt        *time.Time                                 `json:"created_at"`
	UpdatedAt        *time.Time                                `json:"updated_at"`
	DeletedAt        *time.Time                                `json:"deleted_at"`
}

func (w *Woman) IsNil() bool        { return w.ID == 0 }
func (w *Woman) GetID() uint         { return w.ID }
func (w *Woman) GetCompanyID() uint  { return w.CompanyID }
func (w *Woman) GetName() string     { return w.Name }
func (w *Woman) GetAge() *int        { return w.Age }
func (w *Woman) GetBirthplace() *string { return w.Birthplace }
func (w *Woman) GetBloodType() *string  { return w.BloodType }
func (w *Woman) GetHobby() *string      { return w.Hobby }
func (w *Woman) GetIsActive() bool      { return w.IsActive }
func (w *Woman) GetStoreAssignments() collection.Collection[WomanStoreAssignment] {
	return w.StoreAssignments
}
func (w *Woman) GetImages() collection.Collection[WomanImage]                    { return w.Images }
func (w *Woman) GetBlogs() collection.Collection[BlogEntity]                     { return w.Blogs }
