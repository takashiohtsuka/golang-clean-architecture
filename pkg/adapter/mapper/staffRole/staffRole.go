package staffRole

import (
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/domain/model"
)

func ToEntity(ormStaffRole *model.StaffRole) (*entity.StaffRole, error) {
	return &entity.StaffRole{
		ID:      ormStaffRole.ID,
		StaffId: ormStaffRole.StaffId,
		RoleId:  ormStaffRole.RoleId,
	}, nil
}

func ToOrmModel(entityStaffRole *entity.StaffRole) (*model.StaffRole, error) {
	return &model.StaffRole{
		ID:      entityStaffRole.ID,
		StaffId: entityStaffRole.StaffId,
		RoleId:  entityStaffRole.RoleId,
	}, nil
}
