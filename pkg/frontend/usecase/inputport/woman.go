package inputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
)

type WomanUsecase interface {
	GetList(ctx context.Context, i input.GetWomanListInput) (collection.Collection[entity.WomanEntity], error)
	GetDetail(i input.GetWomanDetailInput) (entity.WomanEntity, error)
}
