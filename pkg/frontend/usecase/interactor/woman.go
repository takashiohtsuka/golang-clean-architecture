package interactor

import (
	"context"

	"golang-clean-architecture/pkg/apperror"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
	"golang-clean-architecture/pkg/frontend/usecase/inputport"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"
)

type womanUsecase struct {
	womanRepository outputport.WomanRepository
}

func NewWomanUsecase(womanRepository outputport.WomanRepository) inputport.WomanUsecase {
	return &womanUsecase{womanRepository}
}

func (u *womanUsecase) GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[entity.WomanEntity], error) {
	return u.womanRepository.FindAll(ctx, []query.Condition{})
}

func (u *womanUsecase) GetStoreWomanList(ctx context.Context, i input.GetStoreWomanListInput) (collection.Collection[entity.WomanEntity], error) {
	return u.womanRepository.FindAll(ctx, []query.Condition{
		query.Where("wsa.store_id", i.StoreID),
	})
}

func (u *womanUsecase) GetDetail(ctx context.Context, i input.GetWomanDetailInput) (entity.WomanEntity, error) {
	woman, err := u.womanRepository.FindOne(ctx, []query.Condition{
		query.Where("w.id", i.WomanID),
	})
	if err != nil {
		return nil, err
	}
	if woman.IsNil() {
		return nil, apperror.NewNotFoundException("woman not found")
	}
	return woman, nil
}
