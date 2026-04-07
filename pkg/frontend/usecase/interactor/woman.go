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
	conditions := make([]query.Condition, 0)
	if i.StoreID != nil {
		conditions = append(conditions, query.Where("wsa.store_id", *i.StoreID))
	}
	return u.womanRepository.FindAll(ctx, conditions)
}

func (u *womanUsecase) GetDetail(i input.GetWomanDetailInput) (entity.WomanEntity, error) {
	woman, err := u.womanRepository.FindOne([]query.Condition{
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
