package outputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreDistrictRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.StoreEntity], error)
}
