package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

/*
*
コンストラクタインジェクションでDIしている
*/
func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		User:                  r.NewUserController(),
		Staff:                 r.NewStaffController(),
		Role:                  r.NewRoleController(),
		FanIn:                 r.NewFanInController(),
		URLDownloadSequential: r.NewURLDownloadSequentialController(),
		URLDownloadConcurrent: r.NewURLDownloadConcurrentController(),
	}
}
