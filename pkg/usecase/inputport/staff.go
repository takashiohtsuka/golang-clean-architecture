package inputport

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/domain/entity"
	"golang-clean-architecture/pkg/usecase/input"
)

type StaffUsecase interface {
	List(i input.ListStaffInput) (collection.Collection[entity.StaffEntity], error)
	Create(ctx context.Context, i input.CreateStaffInput) (*entity.Staff, error)
	Update(ctx context.Context, i input.UpdateStaffInput) (bool, error)
}
