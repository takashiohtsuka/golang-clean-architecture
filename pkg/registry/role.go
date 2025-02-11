package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewRoleController() controller.Role {
	role := usecase.NewRoleUsecase(
		repository.NewRoleRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewRoleController(role)
}
