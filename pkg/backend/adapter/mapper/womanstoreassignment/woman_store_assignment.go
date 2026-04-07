package womanstoreassignment

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"
)

func ToEntity(m *model.WomanStoreAssignment) (*entity.WomanStoreAssignment, error) {
	return &entity.WomanStoreAssignment{
		ID:        m.ID,
		StoreID:   m.StoreID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func ToOrmModel(e *entity.WomanStoreAssignment, womanID uint) (*model.WomanStoreAssignment, error) {
	return &model.WomanStoreAssignment{
		ID:        e.ID,
		WomanID:   womanID,
		StoreID:   e.StoreID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}, nil
}
