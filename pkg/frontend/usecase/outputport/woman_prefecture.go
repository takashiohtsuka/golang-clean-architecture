package outputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
)

type WomanPrefectureRepository interface {
	FindAllByPrefecture(ctx context.Context, prefectureID uint) (collection.Collection[entity.WomanEntity], error)
}
