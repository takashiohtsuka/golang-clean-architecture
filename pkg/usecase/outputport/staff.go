package outputport

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type StaffRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.StaffEntity], error)
	FindOne(conditions []query.Condition) (entity.StaffEntity, error)
	Create(u *entity.Staff) (*entity.Staff, error)
	Update(s *entity.Staff) (*entity.Staff, error)
	ReplaceRoles(staffId uint, roleIds []uint) error
}
