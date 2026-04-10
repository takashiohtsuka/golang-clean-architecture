package outputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type ManagementStaffRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.ManagementStaffEntity, error)
	Create(ctx context.Context, m *entity.ManagementStaff) error
	Update(ctx context.Context, m *entity.ManagementStaff) error
}
