package interactor

import (
	"errors"
	"golang-clean-architecture/pkg/domain/model"
	"golang-clean-architecture/pkg/usecase/outputport"
)

type UserUsecase struct {
	userRepository outputport.UserRepository
	dBRepository   outputport.DBRepository
}

// コンストラクタ
func NewUserUsecase(r outputport.UserRepository, d outputport.DBRepository) *UserUsecase {
	return &UserUsecase{r, d}
}

func (uu *UserUsecase) List(u []*model.User) ([]*model.User, error) {
	u, err := uu.userRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *UserUsecase) Create(u *model.User) (*model.User, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		u, err := uu.userRepository.WithTx(i).Create(u)

		// do mailing
		// do logging
		// do another process
		return u, err
	})
	user, ok := data.(*model.User)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
