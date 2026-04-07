package inputport

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
)

type StoreUsecase interface {
	GetList(i input.GetStoreListInput) (collection.Collection[entity.StoreEntity], error)
	GetDetail(i input.GetStoreDetailInput) (entity.StoreEntity, error)
}
