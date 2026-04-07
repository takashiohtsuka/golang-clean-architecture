package repository

import (
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) outputport.BlogRepository {
	return &blogRepository{db: db}
}

type blogRow struct {
	ID          uint       `gorm:"column:id"`
	WomanID     uint       `gorm:"column:woman_id"`
	Title       string     `gorm:"column:title"`
	Body        *string    `gorm:"column:body"`
	IsPublished bool       `gorm:"column:is_published"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (r *blogRow) toEntity() *entity.Blog {
	return &entity.Blog{
		ID:          r.ID,
		WomanID:     r.WomanID,
		Title:       r.Title,
		Body:        r.Body,
		IsPublished: r.IsPublished,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt,
	}
}

type photoRow struct {
	ID        uint       `gorm:"column:id"`
	BlogID    uint       `gorm:"column:blog_id"`
	URL       string     `gorm:"column:url"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (r *blogRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.BlogEntity], error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, woman_id, title, body, is_published, created_at, updated_at, deleted_at
	        FROM blogs WHERE deleted_at IS NULL` + where

	var rows []blogRow
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.BlogEntity](nil), err
	}

	items := make([]entity.BlogEntity, len(rows))
	for i, row := range rows {
		e := row.toEntity()
		e.Photos = collection.NewCollection[entity.Photo](nil)
		items[i] = e
	}
	return collection.NewCollection(items), nil
}

func (r *blogRepository) FindOne(conditions []query.Condition) (entity.BlogEntity, error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, woman_id, title, body, is_published, created_at, updated_at, deleted_at
	        FROM blogs WHERE deleted_at IS NULL` + where + ` LIMIT 1`

	var row blogRow
	if err := r.db.Raw(sql, args...).Scan(&row).Error; err != nil {
		return &entity.NilBlog{}, err
	}
	if row.ID == 0 {
		return &entity.NilBlog{}, nil
	}

	e := row.toEntity()
	photos, err := r.findPhotos(row.ID)
	if err != nil {
		return &entity.NilBlog{}, err
	}
	e.Photos = photos
	return e, nil
}

func (r *blogRepository) Create(b *entity.Blog) error {
	sql := `INSERT INTO blogs (woman_id, title, body, is_published, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
	return r.db.Exec(sql, b.WomanID, b.Title, b.Body, b.IsPublished).Error
}

func (r *blogRepository) Update(b *entity.Blog) error {
	sql := `UPDATE blogs SET woman_id = ?, title = ?, body = ?, is_published = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, b.WomanID, b.Title, b.Body, b.IsPublished, b.ID).Error
}

func (r *blogRepository) Delete(id uint) error {
	sql := `UPDATE blogs SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, id).Error
}

func (r *blogRepository) findPhotos(blogID uint) (collection.Collection[entity.Photo], error) {
	var rows []photoRow
	sql := `SELECT id, blog_id, url, created_at, updated_at FROM photos WHERE blog_id = ?`
	if err := r.db.Raw(sql, blogID).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.Photo](nil), err
	}

	items := make([]entity.Photo, len(rows))
	for i, row := range rows {
		items[i] = entity.Photo{
			ID:        row.ID,
			BlogID:    row.BlogID,
			URL:       row.URL,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}
	return collection.NewCollection(items), nil
}
