package outputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
)

type WomanDistrictRepository interface {
	FindAllByDistrict(ctx context.Context, districtID uint) (collection.Collection[entity.WomanEntity], error)
}
