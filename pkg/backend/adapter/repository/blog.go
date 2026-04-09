package repository

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) outputport.BlogRepository {
	return &blogRepository{db: db}
}

const blogSelectSQL = `
	SELECT id, woman_id, title, body, is_published, created_at, updated_at, deleted_at
	FROM blogs WHERE deleted_at IS NULL`

func toBlogEntity(row map[string]any) *entity.Blog {
	return &entity.Blog{
		ID:          helper.ToUint(row["id"]),
		WomanID:     helper.ToUint(row["woman_id"]),
		Title:       func() string { s := helper.ToStringPtr(row["title"]); if s != nil { return *s }; return "" }(),
		Body:        helper.ToStringPtr(row["body"]),
		IsPublished: helper.ToBool(row["is_published"]),
		CreatedAt:   helper.ToTimePtr(row["created_at"]),
		UpdatedAt:   helper.ToTimePtr(row["updated_at"]),
		DeletedAt:   helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *blogRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.BlogEntity], error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(blogSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.BlogEntity](nil), err
	}

	items := make([]entity.BlogEntity, len(rows))
	for i, row := range rows {
		e := toBlogEntity(row)
		e.Photos = collection.NewCollection[entity.Photo](nil)
		items[i] = e
	}
	return collection.NewCollection(items), nil
}

func (r *blogRepository) FindOne(conditions []query.Condition) (entity.BlogEntity, error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(blogSelectSQL+where+` LIMIT 1`, args...).Scan(&rows).Error; err != nil {
		return &entity.NilBlog{}, err
	}
	if len(rows) == 0 {
		return &entity.NilBlog{}, nil
	}

	e := toBlogEntity(rows[0])
	photos, err := r.findPhotos(e.ID)
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
	var rows []map[string]any
	sql := `SELECT id, blog_id, url, created_at, updated_at FROM photos WHERE blog_id = ?`
	if err := r.db.Raw(sql, blogID).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.Photo](nil), err
	}

	items := make([]entity.Photo, len(rows))
	for i, row := range rows {
		items[i] = entity.Photo{
			ID:        helper.ToUint(row["id"]),
			BlogID:    helper.ToUint(row["blog_id"]),
			URL:       func() string { s := helper.ToStringPtr(row["url"]); if s != nil { return *s }; return "" }(),
			CreatedAt: helper.ToTimePtr(row["created_at"]),
			UpdatedAt: helper.ToTimePtr(row["updated_at"]),
		}
	}
	return collection.NewCollection(items), nil
}
