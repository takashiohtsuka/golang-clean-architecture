package repository

import "golang-clean-architecture/pkg/domain/model"

/*
interfaceだが具象classとなるgoファイルは/adapter/repository/user.go
*/
type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
	Create(u *model.User) (*model.User, error)
}
