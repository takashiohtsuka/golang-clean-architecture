package registry

import (
	"golang-clean-architecture/pkg/frontend/adapter/controller"
	frontendrepository "golang-clean-architecture/pkg/frontend/adapter/repository"
	"golang-clean-architecture/pkg/frontend/usecase/interactor"
)

func (r *registry) NewWomanDistrictController() controller.WomanDistrict {
	u := interactor.NewWomanDistrictUsecase(
		frontendrepository.NewWomanDistrictRepository(r.db),
	)
	return controller.NewWomanDistrictController(u)
}
