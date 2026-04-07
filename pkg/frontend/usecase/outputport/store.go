package outputport

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/usecase/query"
)

type StoreRepository interface {
	// FindAll は店舗一覧を所属女性・ブログタイトルとともに返す。
	FindAll(conditions []query.Condition) (collection.Collection[entity.StoreEntity], error)
	// FindOne は条件に一致する店舗を1件返す。
	FindOne(conditions []query.Condition) (entity.StoreEntity, error)
}
