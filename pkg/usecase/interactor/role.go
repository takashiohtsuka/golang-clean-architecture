package interactor

import (
	"errors"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/outputport"
)

type RoleUsecase struct {
	roleRepository outputport.RoleRepository
	dBRepository   outputport.DBRepository
}

// コンストラクタ
func NewRoleUsecase(r outputport.RoleRepository, d outputport.DBRepository) *RoleUsecase {
	return &RoleUsecase{r, d}
}

func (uu *RoleUsecase) Create(u *entity.Role) (*entity.Role, error) {
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
