package mapper

import (
	"time"

	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

func ToEntity(ormStaff *model.Staff) (*entity.Staff, error) {
	return &entity.Staff{
		ID:        ormStaff.ID,
		Name:      ormStaff.Name,
		Age:       ormStaff.Age,
		IsActive:  ormStaff.IsActive,
		CreatedAt: ormStaff.CreatedAt,
		UpdatedAt: ormStaff.UpdatedAt,
		DeletedAt: toTimePtr(ormStaff.DeletedAt),
	}, nil
}

func ToOrmModel(entityStaff *entity.Staff) (*model.Staff, error) {
	return &model.Staff{
		ID:        entityStaff.ID,
		Name:      entityStaff.Name,
		Age:       entityStaff.Age,
		IsActive:  entityStaff.IsActive,
		CreatedAt: entityStaff.CreatedAt,
		UpdatedAt: entityStaff.UpdatedAt,
		DeletedAt: toDeletedAt(entityStaff.DeletedAt),
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
