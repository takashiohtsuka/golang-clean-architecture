package outputport

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type ManagementStaffRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error)
	FindOne(conditions []query.Condition) (entity.ManagementStaffEntity, error)
	Create(m *entity.ManagementStaff) error
	Update(m *entity.ManagementStaff) error
}
