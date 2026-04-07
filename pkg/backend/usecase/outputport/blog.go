package outputport

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"
)

type BlogRepository interface {
	FindAll(conditions []query.Condition) (collection.Collection[entity.BlogEntity], error)
	FindOne(conditions []query.Condition) (entity.BlogEntity, error)
	Create(b *entity.Blog) error
	Update(b *entity.Blog) error
	Delete(id uint) error
}
