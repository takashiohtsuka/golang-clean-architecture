package interactor

import (
	"context"
	"strconv"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/input"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type UserUsecase struct {
	userRepository outputport.UserRepository
	uow            outputport.UnitOfWork
}

// コンストラクタ
func NewUserUsecase(userRepository outputport.UserRepository, uow outputport.UnitOfWork) *UserUsecase {
	return &UserUsecase{userRepository, uow}
}

func (uu *UserUsecase) List(input input.ListUserInput) (collection.Collection[entity.UserEntity], error) {
	conditions := []query.Condition{}

	if input.Name != "" {
		conditions = append(conditions, query.Where("name", input.Name))
	}
	if input.Age != nil {
		conditions = append(conditions, query.Where("age", *input.Age))
	}

	return uu.userRepository.FindAll(conditions)
}

func (uu *UserUsecase) Create(ctx context.Context, input input.CreateUserInput) (*entity.User, error) {
	user := &entity.User{
		Name: input.Name,
		Age:  strconv.Itoa(input.Age),
	}

	var createdUser *entity.User
	err := uu.uow.Do(ctx, func() error {
		var e error
		createdUser, e = uu.userRepository.Create(user)

		// do mailing
		// do logging
		// do another process
		return e
	})
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
