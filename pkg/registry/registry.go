package registry

import (
	"golang-clean-architecture/pkg/adapter/controller"

	"github.com/jinzhu/gorm"
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
		User:  r.NewUserController(),
		Staff: r.NewStaffController(),
	}
}
