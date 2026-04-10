package outputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.StoreEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error)
	Create(ctx context.Context, s *entity.Store) error
	Update(ctx context.Context, s *entity.Store) error
	AddWoman(ctx context.Context, womanID uint, storeID uint) error
	RemoveWoman(ctx context.Context, womanID uint, storeID uint) error
}
