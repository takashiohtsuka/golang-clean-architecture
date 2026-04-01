package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"
	"golang-clean-architecture/pkg/adapter/repository"
	"golang-clean-architecture/pkg/usecase/interactor"
)

/*
*
UserControllerでDIするUsecaseとRepositryをコンストラクタインジェクションでDI
*/
func (r *registry) NewUserController() controller.User {
	u := interactor.NewUserUsecase(
		repository.NewUserRepository(r.db),
		repository.NewDBRepository(r.db),
	)

	return controller.NewUserController(u)
}
