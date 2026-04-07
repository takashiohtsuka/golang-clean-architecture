package repository

import (
	"errors"

	"gorm.io/gorm"
	mapper "golang-clean-architecture/pkg/adapter/mapper/role"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) outputport.RoleRepository {
	return &roleRepository{db}
}

func (ur *roleRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.RoleEntity], error) {
	var ormRoles []*model.Role
	if err := buildQuery(ur.db, conditions).Find(&ormRoles).Error; err != nil {
		return collection.NewCollection[entity.RoleEntity](nil), err
	}

	if len(ormRoles) == 0 {
		return collection.NewCollection[entity.RoleEntity](nil), nil
	}

	items := make([]entity.RoleEntity, len(ormRoles))
	for i, ormRole := range ormRoles {
		items[i], _ = mapper.ToEntity(ormRole)
	}

	return collection.NewCollection(items), nil
}

func (ur *roleRepository) Create(r *entity.Role) (*entity.Role, error) {
	ormRole, _ := mapper.ToOrmModel(r)
	if err := ur.db.Create(ormRole).Error; err != nil {
		return nil, err
	}

	return r, nil
}

func (ur *roleRepository) FindOne(conditions []query.Condition) (entity.RoleEntity, error) {
	modelRole := &model.Role{}

	if err := buildQuery(ur.db, conditions).First(&modelRole).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &entity.NilRole{}, nil
		}
		return &entity.NilRole{}, err
	}

	entityRole, _ := mapper.ToEntity(modelRole)

	return entityRole, nil
}
