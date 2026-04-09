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
	Create(ctx context.Context, w *entity.Woman) (uint, error)
	Update(ctx context.Context, w *entity.Woman) error
	SaveImage(ctx context.Context, womanID uint, path string) error
}
