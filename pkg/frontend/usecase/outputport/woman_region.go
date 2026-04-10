package outputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
)

type WomanRegionRepository interface {
	FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[entity.WomanEntity], error)
}
