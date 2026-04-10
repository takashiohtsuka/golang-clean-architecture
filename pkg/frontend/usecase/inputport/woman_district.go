package inputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
)

type WomanDistrictUsecase interface {
	GetList(ctx context.Context, i input.GetWomanDistrictListInput) (collection.Collection[entity.WomanEntity], error)
}
