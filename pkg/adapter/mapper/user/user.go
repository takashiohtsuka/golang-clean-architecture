package mapper

import (
	"strconv"
	"time"

	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

func ToEntity(ormUser *model.User) (*entity.User, error) {
	return &entity.User{
		ID:        ormUser.ID,
		Name:      ormUser.Name,
		Age:       strconv.Itoa(ormUser.Age),
		CreatedAt: ormUser.CreatedAt,
		UpdatedAt: ormUser.UpdatedAt,
		DeletedAt: toTimePtr(ormUser.DeletedAt),
	}, nil
}

func ToOrmModel(entityUser *entity.User) (*model.User, error) {
	age, _ := strconv.Atoi(entityUser.Age)
	return &model.User{
		ID:        entityUser.ID,
		Name:      entityUser.Name,
		Age:       age,
		CreatedAt: entityUser.CreatedAt,
		UpdatedAt: entityUser.UpdatedAt,
		DeletedAt: toDeletedAt(entityUser.DeletedAt),
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
