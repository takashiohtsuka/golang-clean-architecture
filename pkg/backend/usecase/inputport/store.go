package inputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/usecase/input"
)

type StoreUsecase interface {
	Create(ctx context.Context, i input.CreateStoreInput) error
	Update(ctx context.Context, i input.UpdateStoreInput) error
}
