package repository

import (
	"github.com/jinzhu/gorm"
	mapper "golang-clean-architecture/pkg/adapter/mapper/staffRole"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/repository"
)

type staffRoleRepository struct {
	db *gorm.DB
}

func NewStaffRoleRepository(db *gorm.DB) repository.StaffRoleRepository {
	return &staffRoleRepository{db}
}

func (ur *staffRoleRepository) Create(sr *entity.StaffRole) (*entity.StaffRole, error) {
	ormStaffRole, _ := mapper.ToOrmModel(sr)
	if err := ur.db.Create(ormStaffRole).Error; err != nil {
		return nil, err
	}

	return sr, nil
}

func (ur *staffRoleRepository) Update(sr *entity.StaffRole) (*entity.StaffRole, error) {
	ormStaffRole, _ := mapper.ToOrmModel(sr)
	if err := ur.db.Model(&ormStaffRole).Update(ormStaffRole).Error; err != nil {
		return nil, err
	}

	return sr, nil
}

func (ur *staffRoleRepository) FindOne(conditions map[string]interface{}) (*entity.StaffRole, error) {
	modelStaffRole := &model.StaffRole{}

	if err := ur.db.Where(conditions).First(&modelStaffRole).Error; err != nil {
		return nil, err
	}

	entityStaffRole, _ := mapper.ToEntity(modelStaffRole)

	return entityStaffRole, nil
}
