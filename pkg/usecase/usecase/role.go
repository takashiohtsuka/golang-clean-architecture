package usecase

import (
	"errors"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/repository"
)

type roleUsecase struct {
	roleRepository repository.RoleRepository
	dBRepository   repository.DBRepository
}

type Role interface {
	Create(u *entity.Role) (*entity.Role, error)
}

// コンストラクタ
func NewRoleUsecase(r repository.RoleRepository, d repository.DBRepository) Role {
	return &roleUsecase{r, d}
}

func (uu *roleUsecase) Create(u *entity.Role) (*entity.Role, error) {
	data, err := uu.dBRepository.Transaction(func(i interface{}) (interface{}, error) {
		s, err := uu.roleRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return s, err
	})
	role, ok := data.(*entity.Role)

	if !ok {
		return nil, errors.New("cast error")
	}

	if err != nil {
		return nil, err
	}

	return role, nil
}
