package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/usecase/interactor"
)

func (r *registry) NewURLDownloadConcurrentController() controller.URLDownloadConcurrent {
	return controller.NewURLDownloadConcurrentController(interactor.NewURLDownloadConcurrentUsecase())
}
