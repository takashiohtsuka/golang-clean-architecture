package inputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/input"
)

type UserUsecase interface {
	List(i input.ListUserInput) (collection.Collection[entity.UserEntity], error)
	Create(ctx context.Context, i input.CreateUserInput) (*entity.User, error)
}
