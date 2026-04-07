package interactor

import (
	"context"
	"errors"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/input"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type ManagementStaffUsecase struct {
	managementStaffRepository backendoutputport.ManagementStaffRepository
	companyRepository         backendoutputport.CompanyRepository
	storeRepository           backendoutputport.StoreRepository
	uow                       outputport.UnitOfWork
}

func NewManagementStaffUsecase(
	managementStaffRepository backendoutputport.ManagementStaffRepository,
	companyRepository backendoutputport.CompanyRepository,
	storeRepository backendoutputport.StoreRepository,
	uow outputport.UnitOfWork,
) inputport.ManagementStaffUsecase {
	return &ManagementStaffUsecase{managementStaffRepository, companyRepository, storeRepository, uow}
}

func (u *ManagementStaffUsecase) Create(ctx context.Context, i input.CreateManagementStaffInput) error {
	company, err := u.companyRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.CompanyID),
	})
	if err != nil {
		return err
	}
	if company.IsNil() {
		return errors.New("company not found")
	}

	store, err := u.storeRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.StoreID),
		query.Where("company_id", i.CompanyID),
	})
	if err != nil {
		return err
	}
	if store.IsNil() {
		return errors.New("store not found")
	}

	staff := &entity.ManagementStaff{
		CompanyID: i.CompanyID,
		StoreID:   i.StoreID,
		Name:      i.Name,
		Email:     i.Email,
	}

	return u.uow.Do(ctx, func() error {
		return u.managementStaffRepository.Create(staff)
	})
}
