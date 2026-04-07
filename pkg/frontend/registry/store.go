package registry

import (
	"golang-clean-architecture/pkg/frontend/adapter/controller"
	frontendrepository "golang-clean-architecture/pkg/frontend/adapter/repository"
	"golang-clean-architecture/pkg/frontend/usecase/interactor"
)

func (r *registry) NewStoreController() controller.Store {
	u := interactor.NewStoreUsecase(
		frontendrepository.NewStoreRepository(r.db),
	)
	return controller.NewStoreController(u)
}
