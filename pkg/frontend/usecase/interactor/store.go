package interactor

import (
	"context"

	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type storeUsecase struct {
	storeRepository outputport.StoreRepository
}

func NewStoreUsecase(storeRepository outputport.StoreRepository) inputport.StoreUsecase {
	return &storeUsecase{storeRepository}
}

func (u *storeUsecase) GetDetail(ctx context.Context, i input.GetStoreDetailInput) (entity.StoreEntity, error) {
	return u.storeRepository.FindOne(ctx, []query.Condition{
		query.Where("id", i.StoreID),
	})
}
