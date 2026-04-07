package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewRoleController() controller.Role {
	role := interactor.NewRoleUsecase(
		repository.NewRoleRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)

	return controller.NewRoleController(role)
}
