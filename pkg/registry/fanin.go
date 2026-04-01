package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewFanInController() controller.FanIn {
	return controller.NewFanInController(interactor.NewFanInUsecase())
}
