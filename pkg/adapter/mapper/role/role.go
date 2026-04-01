package role

import (
	"time"

	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"

	"gorm.io/gorm"
)

func ToEntity(ormRole *model.Role) (*entity.Role, error) {
	return &entity.Role{
		ID:        ormRole.ID,
		Name:      ormRole.Name,
		CreatedAt: ormRole.CreatedAt,
		UpdatedAt: ormRole.UpdatedAt,
		DeletedAt: toTimePtr(ormRole.DeletedAt),
	}, nil
}

func ToOrmModel(entityRole *entity.Role) (*model.Role, error) {
	return &model.Role{
		ID:        entityRole.ID,
		Name:      entityRole.Name,
		CreatedAt: entityRole.CreatedAt,
		UpdatedAt: entityRole.UpdatedAt,
		DeletedAt: toDeletedAt(entityRole.DeletedAt),
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
