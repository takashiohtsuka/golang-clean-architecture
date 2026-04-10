package outputport

import (
	"context"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type CompanyRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.CompanyEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.CompanyEntity, error)
	Create(ctx context.Context, c *entity.Company) error
	Update(ctx context.Context, c *entity.Company) error
}
