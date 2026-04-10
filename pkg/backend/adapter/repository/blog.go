package repository

import (
	"context"

	blogMapper "golang-clean-architecture/pkg/backend/adapter/mapper/blog"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/infrastructure/model"
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
		Title:       helper.ToString(row["title"]),
		Body:        helper.ToStringPtr(row["body"]),
		IsPublished: helper.ToBool(row["is_published"]),
		CreatedAt:   helper.ToTimePtr(row["created_at"]),
		UpdatedAt:   helper.ToTimePtr(row["updated_at"]),
		DeletedAt:   helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *blogRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.BlogEntity], error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(blogSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.BlogEntity](nil), err
	}
	if len(rows) == 0 {
		return collection.NewCollection[entity.BlogEntity](nil), nil
	}

	blogs := make([]*entity.Blog, len(rows))
	blogIDs := make([]uint, len(rows))
	for i, row := range rows {
		blogs[i] = toBlogEntity(row)
		blogIDs[i] = blogs[i].ID
	}

	photosByBlogID, err := r.findPhotosByBlogIDs(ctx, blogIDs)
	if err != nil {
		return collection.NewCollection[entity.BlogEntity](nil), err
	}

	items := make([]entity.BlogEntity, len(blogs))
	for i, b := range blogs {
		b.Photos = collection.NewCollection(photosByBlogID[b.ID])
		items[i] = b
	}
	return collection.NewCollection(items), nil
}

func (r *blogRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.BlogEntity, error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(blogSelectSQL+where+` LIMIT 1`, args...).Scan(&rows).Error; err != nil {
		return &entity.NilBlog{}, err
	}
	if len(rows) == 0 {
		return &entity.NilBlog{}, nil
	}

	e := toBlogEntity(rows[0])
	photos, err := r.findPhotos(ctx, e.ID)
	if err != nil {
		return &entity.NilBlog{}, err
	}
	e.Photos = photos
	return e, nil
}

func (r *blogRepository) Create(ctx context.Context, b *entity.Blog) error {
	return r.db.WithContext(ctx).Create(blogMapper.ToOrmModel(b)).Error
}

func (r *blogRepository) Update(ctx context.Context, b *entity.Blog) error {
	return r.db.WithContext(ctx).Save(blogMapper.ToOrmModel(b)).Error
}

func (r *blogRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Blog{}, id).Error
}

func (r *blogRepository) findPhotos(ctx context.Context, blogID uint) (collection.Collection[entity.Photo], error) {
	photos, err := r.findPhotosByBlogIDs(ctx, []uint{blogID})
	if err != nil {
		return collection.NewCollection[entity.Photo](nil), err
	}
	return collection.NewCollection(photos[blogID]), nil
}

func (r *blogRepository) findPhotosByBlogIDs(ctx context.Context, blogIDs []uint) (map[uint][]entity.Photo, error) {
	var rows []map[string]any
	sql := `SELECT id, blog_id, url, created_at, updated_at FROM photos WHERE blog_id IN ?`
	if err := r.db.WithContext(ctx).Raw(sql, blogIDs).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make(map[uint][]entity.Photo)
	for _, row := range rows {
		blogID := helper.ToUint(row["blog_id"])
		result[blogID] = append(result[blogID], entity.Photo{
			ID:        helper.ToUint(row["id"]),
			BlogID:    blogID,
			URL:       helper.ToString(row["url"]),
			CreatedAt: helper.ToTimePtr(row["created_at"]),
			UpdatedAt: helper.ToTimePtr(row["updated_at"]),
		})
	}
	return result, nil
}
