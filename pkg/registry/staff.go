package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/usecase"
)

func (r *registry) NewStaffController() controller.Staff {
	s := usecase.NewStaffUsecase(
		repository.NewStaffRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewStaffController(s)
}
