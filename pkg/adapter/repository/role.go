package repository

import (
	"gorm.io/gorm"
	mapper "golang-clean-architecture/pkg/adapter/mapper/role"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/outputport"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) outputport.RoleRepository {
	return &roleRepository{db}
}

func (ur *roleRepository) Create(r *entity.Role) (*entity.Role, error) {
	ormRole, _ := mapper.ToOrmModel(r)
	if err := ur.db.Create(ormRole).Error; err != nil {
		return nil, err
	}

	return r, nil
}

func (ur *roleRepository) FindOne(conditions map[string]interface{}) (*entity.Role, error) {
	modelRole := &model.Role{}
	
	if err := ur.db.Model(&model.Role{}).Where(conditions).First(&modelRole).Error; err != nil {
		return nil, err
	}

	entityRole, _ := mapper.ToEntity(modelRole)

	return entityRole, nil
}
