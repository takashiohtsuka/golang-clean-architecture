package photo

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/infrastructure/model"
)

func ToEntity(m *model.Photo) (*entity.Photo, error) {
	return &entity.Photo{
		ID:        m.ID,
		BlogID:    m.BlogID,
		URL:       m.URL,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}, nil
}

func ToOrmModel(e *entity.Photo) (*model.Photo, error) {
	return &model.Photo{
		ID:        e.ID,
		BlogID:    e.BlogID,
		URL:       e.URL,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}, nil
}
