package interactor

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type RoleUsecase struct {
	roleRepository outputport.RoleRepository
	uow            outputport.UnitOfWork
}

// コンストラクタ
func NewRoleUsecase(r outputport.RoleRepository, uow outputport.UnitOfWork) *RoleUsecase {
	return &RoleUsecase{r, uow}
}

func (uu *RoleUsecase) List() (collection.Collection[entity.RoleEntity], error) {
	return uu.roleRepository.FindAll([]query.Condition{})
}

func (uu *RoleUsecase) Create(ctx context.Context, u *entity.Role) (*entity.Role, error) {
	var role *entity.Role
	err := uu.uow.Do(ctx, func() error {
		var e error
		role, e = uu.roleRepository.Create(u)

		// do mailing
		// do logging
		// do another process
		return e
	})
	if err != nil {
		return nil, err
	}

	return role, nil
}
