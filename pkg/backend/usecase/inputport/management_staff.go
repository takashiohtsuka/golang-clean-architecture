package inputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/usecase/input"
)

type ManagementStaffUsecase interface {
	Create(ctx context.Context, i input.CreateManagementStaffInput) error
}
