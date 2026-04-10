package blog

import (
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/infrastructure/model"

	"gorm.io/gorm"
)

// ToEntity はORM modelをentityに変換する。
// Photos は含まれないため、リポジトリが別途ロード後にセットすること。
func ToEntity(m *model.Blog) *entity.Blog {
	return &entity.Blog{
		ID:          m.ID,
		WomanID:     m.WomanID,
		Title:       m.Title,
		Body:        m.Body,
		IsPublished: m.IsPublished,
		Photos:      collection.NewCollection([]entity.Photo{}),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   toTimePtr(m.DeletedAt),
	}
}

func ToOrmModel(e *entity.Blog) *model.Blog {
	return &model.Blog{
		ID:          e.ID,
		WomanID:     e.WomanID,
		Title:       e.Title,
		Body:        e.Body,
		IsPublished: e.IsPublished,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   toDeletedAt(e.DeletedAt),
	}
}

func toTimePtr(d gorm.DeletedAt) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}

func toDeletedAt(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{Time: *t, Valid: true}
	}
	return gorm.DeletedAt{}
}
