package interactor

import (
	"golang-clean-architecture/pkg/domain/collection"
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

func (u *storeUsecase) GetList(i input.GetStoreListInput) (collection.Collection[entity.StoreEntity], error) {
	return u.storeRepository.FindAll([]query.Condition{})
}

func (u *storeUsecase) GetDetail(i input.GetStoreDetailInput) (entity.StoreEntity, error) {
	return u.storeRepository.FindOne([]query.Condition{
		query.Where("s.id", i.StoreID),
	})
}
