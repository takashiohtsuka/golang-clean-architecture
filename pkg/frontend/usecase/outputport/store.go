package outputport

import (
	"context"

	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreRepository interface {
	FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error)
}
