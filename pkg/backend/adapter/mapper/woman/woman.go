package woman

import (
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

// ToEntity はORM modelをentityに変換する。
// StoreAssignments は含まれないため、リポジトリが別途ロード後にセットすること。
func ToEntity(m *model.Woman) (*entity.Woman, error) {
	return &entity.Woman{
		ID:         m.ID,
		CompanyID:  m.CompanyID,
		Name:       m.Name,
		Age:        m.Age,
		Birthplace: m.Birthplace,
		BloodType:  m.BloodType,
		Hobby:      m.Hobby,
		IsActive:   m.IsActive,
		Images:     collection.NewCollection([]entity.WomanImage{}),
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  toTimePtr(m.DeletedAt),
	}, nil
}

func ToOrmModel(e *entity.Woman) (*model.Woman, error) {
	return &model.Woman{
		ID:         e.ID,
		CompanyID:  e.CompanyID,
		Name:       e.Name,
		Age:        e.Age,
		Birthplace: e.Birthplace,
		BloodType:  e.BloodType,
		Hobby:      e.Hobby,
		IsActive:   e.IsActive,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		DeletedAt:  toDeletedAt(e.DeletedAt),
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
