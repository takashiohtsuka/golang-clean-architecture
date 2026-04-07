package inputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/usecase/input"
)

type CompanyUsecase interface {
	Create(ctx context.Context, i input.CreateCompanyInput) error
}
