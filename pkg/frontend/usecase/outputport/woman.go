package outputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error)
	FindOne(conditions []query.Condition) (entity.WomanEntity, error)
}
