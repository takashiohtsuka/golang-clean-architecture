package role

import (
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
)

func ToEntity(ormRole *model.Role) (*entity.Role, error) {
	return &entity.Role{
		ID:        ormRole.ID,
		Name:      ormRole.Name,
		CreatedAt: ormRole.CreatedAt,
		UpdatedAt: ormRole.UpdatedAt,
		DeletedAt: ormRole.DeletedAt,
	}, nil
}

func ToOrmModel(entityRole *entity.Role) (*model.Role, error) {
	return &model.Role{
		ID:        entityRole.ID,
		Name:      entityRole.Name,
		CreatedAt: entityRole.CreatedAt,
		UpdatedAt: entityRole.UpdatedAt,
		DeletedAt: entityRole.DeletedAt,
	}, nil
}
