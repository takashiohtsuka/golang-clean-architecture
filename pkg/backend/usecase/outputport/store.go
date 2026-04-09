package outputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.StoreEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error)
	Create(s *entity.Store) error
	Update(s *entity.Store) error
	AddWoman(womanID uint, storeID uint) error
	RemoveWoman(womanID uint, storeID uint) error
}
