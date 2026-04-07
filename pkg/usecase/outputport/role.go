package outputport

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type RoleRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.RoleEntity], error)
	FindOne(conditions []query.Condition) (entity.RoleEntity, error)
	Create(u *entity.Role) (*entity.Role, error)
}
