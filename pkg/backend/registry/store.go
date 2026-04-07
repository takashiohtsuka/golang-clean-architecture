package registry

import (
	"golang-clean-architecture/pkg/backend/adapter/controller"
	backendrepository "golang-clean-architecture/pkg/backend/adapter/repository"
	"golang-clean-architecture/pkg/backend/usecase/interactor"
	"golang-clean-architecture/pkg/adapter/repository"
)

func (r *registry) NewStoreController() controller.Store {
	u := interactor.NewStoreUsecase(
		backendrepository.NewStoreRepository(r.db),
		backendrepository.NewCompanyRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)
	return controller.NewStoreController(u)
}
