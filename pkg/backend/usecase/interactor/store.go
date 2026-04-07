package interactor

import (
	"context"
	"errors"

	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/input"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreUsecase struct {
	storeRepository   backendoutputport.StoreRepository
	companyRepository backendoutputport.CompanyRepository
	uow               outputport.UnitOfWork
}

func NewStoreUsecase(
	storeRepository backendoutputport.StoreRepository,
	companyRepository backendoutputport.CompanyRepository,
	uow outputport.UnitOfWork,
) inputport.StoreUsecase {
	return &StoreUsecase{storeRepository, companyRepository, uow}
}

func (u *StoreUsecase) Create(ctx context.Context, i input.CreateStoreInput) error {
	company, err := u.companyRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.CompanyID),
	})
	if err != nil {
		return err
	}
	if company.IsNil() {
		return errors.New("company not found")
	}

	store := &entity.Store{
		CompanyID:    i.CompanyID,
		BusinessType: bvo.NewBusinessType(i.BusinessTypeCode),
		ContractPlan: bvo.NewContractPlan(i.ContractPlanCode),
		Name:         i.Name,
		IsActive:     i.IsActive,
		OpenStatus:   entity.OpenStatus(i.OpenStatus),
	}

	return u.uow.Do(ctx, func() error {
		return u.storeRepository.Create(store)
	})
}
