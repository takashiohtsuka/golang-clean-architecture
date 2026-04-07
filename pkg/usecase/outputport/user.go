package outputport

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

/*
interfaceだが具象classとなるgoファイルは/adapter/repository/user.go
*/
type UserRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.UserEntity], error)
	Create(u *entity.User) (*entity.User, error)
	WithTx(tx any) UserRepository
}
