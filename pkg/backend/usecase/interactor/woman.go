package interactor

import (
	"context"
	"errors"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/input"
	"golang-clean-architecture/pkg/backend/usecase/inputport"
	backendoutputport "golang-clean-architecture/pkg/backend/usecase/outputport"
	appconfig "golang-clean-architecture/pkg/config"
	"golang-clean-architecture/pkg/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type WomanUsecase struct {
	womanRepository   backendoutputport.WomanRepository
	companyRepository backendoutputport.CompanyRepository
	storeRepository   backendoutputport.StoreRepository
	uow               outputport.UnitOfWork
	storage           backendoutputport.StorageRepository
}

func NewWomanUsecase(
	womanRepository backendoutputport.WomanRepository,
	companyRepository backendoutputport.CompanyRepository,
	storeRepository backendoutputport.StoreRepository,
	uow outputport.UnitOfWork,
	storage backendoutputport.StorageRepository,
) inputport.WomanUsecase {
	return &WomanUsecase{womanRepository, companyRepository, storeRepository, uow, storage}
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
	}
	if hasBTypeStore && len(i.StoreIDs) > 1 {
		return errors.New("業種Bの店舗に所属する場合、他の店舗と同時に所属することはできません")
	}

	woman := &entity.Woman{
		CompanyID:  i.CompanyID,
		Name:       i.Name,
		Age:        i.Age,
		Birthplace: i.Birthplace,
		BloodType:  i.BloodType,
		Hobby:      i.Hobby,
		IsActive:   i.IsActive,
	}

	return u.uow.Do(ctx, func() error {
		womanID, err := u.womanRepository.Create(ctx, woman)
		if err != nil {
			return err
		}
		for _, storeID := range i.StoreIDs {
			if err := u.storeRepository.AddWoman(ctx, womanID, storeID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (u *WomanUsecase) Update(ctx context.Context, i input.UpdateWomanInput) error {
	woman, err := u.womanRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.ID),
	})
	if err != nil {
		return err
	}
	if woman.IsNil() {
		return errors.New("woman not found")
	}

	return u.uow.Do(ctx, func() error {
		updated := &entity.Woman{
			ID:         i.ID,
			CompanyID:  woman.GetCompanyID(),
			Name:       i.Name,
			Age:        i.Age,
			Birthplace: i.Birthplace,
			BloodType:  i.BloodType,
			Hobby:      i.Hobby,
			IsActive:   i.IsActive,
		}
		if err := u.womanRepository.Update(ctx, updated); err != nil {
			return err
		}

		if i.ImageFile == nil {
			return nil
		}

		path, err := u.storage.Upload(
			ctx,
			appconfig.C.Storage.Buckets.WomanImage,
			i.ImageKey,
			i.ImageFile,
			i.ContentType,
		)
		if err != nil {
			return err
		}

		return u.womanRepository.SaveImage(ctx, i.ID, path)
	})
}
