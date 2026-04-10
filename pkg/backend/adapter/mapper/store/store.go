package store

import (
	"time"

	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

func ToEntity(m *model.Store) *entity.Store {
	return &entity.Store{
		ID:           m.ID,
		CompanyID:    m.CompanyID,
		DistrictID:   m.DistrictID,
		BusinessType: bvo.NewBusinessType(m.BusinessType.Code),
		ContractPlan: bvo.NewContractPlan(m.ContractPlan.Code),
		Name:         m.Name,
		IsActive:     m.IsActive,
		OpenStatus:   entity.OpenStatus(m.OpenStatus),
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
		DeletedAt:    toTimePtr(m.DeletedAt),
	}
}

// ToOrmModel はentityをORM modelに変換する。
// BusinessTypeID / ContractPlanID はVOからID解決できないため、
// 呼び出し元（リポジトリ）が別途セットすること。
func ToOrmModel(e *entity.Store) *model.Store {
	return &model.Store{
		ID:         e.ID,
		CompanyID:  e.CompanyID,
		DistrictID: e.DistrictID,
		Name:       e.Name,
		IsActive:   e.IsActive,
		OpenStatus: string(e.OpenStatus),
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		DeletedAt:  toDeletedAt(e.DeletedAt),
	}
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
