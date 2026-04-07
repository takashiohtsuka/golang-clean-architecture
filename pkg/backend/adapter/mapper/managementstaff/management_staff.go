package managementstaff

import (
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

func ToEntity(m *model.ManagementStaff) (*entity.ManagementStaff, error) {
	return &entity.ManagementStaff{
		ID:        m.ID,
		CompanyID: m.CompanyID,
		StoreID:   m.StoreID,
		Name:      m.Name,
		Email:     m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: toTimePtr(m.DeletedAt),
	}, nil
}

func ToOrmModel(e *entity.ManagementStaff) (*model.ManagementStaff, error) {
	return &model.ManagementStaff{
		ID:        e.ID,
		CompanyID: e.CompanyID,
		StoreID:   e.StoreID,
		Name:      e.Name,
		Email:     e.Email,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: toDeletedAt(e.DeletedAt),
	}, nil
}

func toTimePtr(d gorm.DeletedAt) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}

func toDeletedAt(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{Time: *t, Valid: true}
	}
	return gorm.DeletedAt{}
}
