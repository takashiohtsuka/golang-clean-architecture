package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewURLDownloadSequentialController() controller.URLDownloadSequential {
	return controller.NewURLDownloadSequentialController(interactor.NewURLDownloadSequentialUsecase())
}
