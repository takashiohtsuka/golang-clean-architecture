package mapper

import (
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
)

func ToEntity(ormStaff *model.Staff) (*entity.Staff, error) {
	return &entity.Staff{
		ID:        ormStaff.ID,
		Name:      ormStaff.Name,
		Age:       ormStaff.Age,
		IsActive:  ormStaff.IsActive,
		CreatedAt: ormStaff.CreatedAt,
		UpdatedAt: ormStaff.UpdatedAt,
		DeletedAt: ormStaff.DeletedAt,
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
		DeletedAt: entityStaff.DeletedAt,
	}, nil
}
