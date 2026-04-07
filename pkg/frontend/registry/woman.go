package registry

import (
	"golang-clean-architecture/pkg/frontend/adapter/controller"
	frontendrepository "golang-clean-architecture/pkg/frontend/adapter/repository"
	"golang-clean-architecture/pkg/frontend/usecase/interactor"
)

func (r *registry) NewWomanController() controller.Woman {
	u := interactor.NewWomanUsecase(
		frontendrepository.NewWomanRepository(r.db),
	)
	return controller.NewWomanController(u)
}
