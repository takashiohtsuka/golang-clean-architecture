package registry

import (
	"golang-clean-architecture/pkg/backend/adapter/controller"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"

	"gorm.io/gorm"
)

type registry struct {
	db      *gorm.DB
	storage backendoutputport.StorageRepository
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB, storage backendoutputport.StorageRepository) Registry {
	return &registry{db, storage}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Company:         r.NewCompanyController(),
		Store:           r.NewStoreController(),
		ManagementStaff: r.NewManagementStaffController(),
		Woman:           r.NewWomanController(),
	}
}
