package inputport

import (
	"context"

	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/input"
)

type StoreUsecase interface {
	GetDetail(ctx context.Context, i input.GetStoreDetailInput) (entity.StoreEntity, error)
}
