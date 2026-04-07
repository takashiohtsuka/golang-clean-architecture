package interactor

import (
	"context"
	"errors"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/input"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type WomanUsecase struct {
	womanRepository   backendoutputport.WomanRepository
	companyRepository backendoutputport.CompanyRepository
	storeRepository   backendoutputport.StoreRepository
	uow               outputport.UnitOfWork
}

func NewWomanUsecase(
	womanRepository backendoutputport.WomanRepository,
	companyRepository backendoutputport.CompanyRepository,
	storeRepository backendoutputport.StoreRepository,
	uow outputport.UnitOfWork,
) inputport.WomanUsecase {
	return &WomanUsecase{womanRepository, companyRepository, storeRepository, uow}
}

func (u *WomanUsecase) Create(ctx context.Context, i input.CreateWomanInput) error {
	company, err := u.companyRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.CompanyID),
	})
	if err != nil {
		return err
	}
	if company.IsNil() {
		return errors.New("company not found")
	}

	assignments := make([]entity.WomanStoreAssignment, 0, len(i.StoreIDs))
	hasBTypeStore := false
	for _, storeID := range i.StoreIDs {
		store, err := u.storeRepository.FindOne(ctx, []query.Condition{
			query.Where("id", storeID),
			query.Where("company_id", i.CompanyID),
		})
		if err != nil {
			return err
		}
		if store.IsNil() {
			return errors.New("store not found")
		}
		if store.GetBusinessType().GetCode() == "B" {
			hasBTypeStore = true
		}
		assignments = append(assignments, entity.WomanStoreAssignment{StoreID: storeID})
	}
	if hasBTypeStore && len(i.StoreIDs) > 1 {
		return errors.New("業種Bの店舗に所属する場合、他の店舗と同時に所属することはできません")
	}

	woman := &entity.Woman{
		CompanyID:        i.CompanyID,
		Name:             i.Name,
		Age:              i.Age,
		Birthplace:       i.Birthplace,
		BloodType:        i.BloodType,
		Hobby:            i.Hobby,
		IsActive:         i.IsActive,
		StoreAssignments: collection.NewCollection(assignments),
	}

	return u.uow.Do(ctx, func() error {
		return u.womanRepository.Create(ctx, woman)
	})
}
