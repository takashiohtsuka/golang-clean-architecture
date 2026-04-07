package outputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.WomanEntity], error)
	FindOne(conditions []query.Condition) (entity.WomanEntity, error)
	Create(ctx context.Context, w *entity.Woman) error
	Update(w *entity.Woman) error
	AssignToStore(womanID uint, storeID uint) error
	RemoveStoreAssignment(womanID uint, storeID uint) error
}
