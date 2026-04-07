package registry

import (
	"golang-clean-architecture/pkg/backend/adapter/controller"
	backendrepository "golang-clean-architecture/pkg/backend/adapter/repository"
	"golang-clean-architecture/pkg/backend/usecase/interactor"
	"golang-clean-architecture/pkg/adapter/repository"
)

func (r *registry) NewManagementStaffController() controller.ManagementStaff {
	u := interactor.NewManagementStaffUsecase(
		backendrepository.NewManagementStaffRepository(r.db),
		backendrepository.NewCompanyRepository(r.db),
		backendrepository.NewStoreRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)
	return controller.NewManagementStaffController(u)
}
