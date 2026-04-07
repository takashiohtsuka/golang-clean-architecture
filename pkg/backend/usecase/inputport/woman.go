package inputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/usecase/input"
)

type WomanUsecase interface {
	Create(ctx context.Context, i input.CreateWomanInput) error
}
