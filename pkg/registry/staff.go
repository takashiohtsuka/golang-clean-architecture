package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewStaffController() controller.Staff {
	s := interactor.NewStaffUsecase(
		repository.NewStaffRepository(r.db),
		repository.NewUnitOfWork(r.db),
		repository.NewRoleRepository(r.db),
	)

	return controller.NewStaffController(s)
}
