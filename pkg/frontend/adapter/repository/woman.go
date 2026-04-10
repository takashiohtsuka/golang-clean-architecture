package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	womanMapper "golang-clean-architecture/pkg/frontend/adapter/mapper/woman"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

const womanBlogsLimit = 3

type womanRepository struct {
	db *gorm.DB
}

func NewWomanRepository(db *gorm.DB) outputport.WomanRepository {
	return &womanRepository{db: db}
}

func (r *womanRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error) {
	where, args := buildWhereClause(conditions)

	sql := `
		SELECT
			w.id         AS woman_id,
			w.name       AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			wsa.id       AS assignment_id,
			wsa.store_id AS assignment_store_id,
			wi.id        AS image_id,
			wi.path      AS image_path,
			b.id         AS blog_id,
			b.title      AS blog_title
		FROM women w
		LEFT JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		LEFT JOIN woman_images wi ON wi.woman_id = w.id
		LEFT JOIN (
			SELECT id, woman_id, title,
			       ROW_NUMBER() OVER (PARTITION BY woman_id ORDER BY created_at DESC) AS rn
			FROM blogs
			WHERE deleted_at IS NULL AND is_published = TRUE
		) b ON b.woman_id = w.id AND b.rn <= ?
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE` + where + `
		ORDER BY w.id, wsa.id, wi.id, b.id`

	allArgs := append([]any{womanBlogsLimit}, args...)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}
	return womanMapper.MapToAggregate(rows), nil
}

func (r *womanRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.WomanEntity, error) {
	where, args := buildWhereClause(conditions)

	sql := `
		SELECT
			w.id         AS woman_id,
			w.name       AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			wsa.id       AS assignment_id,
			wsa.store_id AS assignment_store_id,
			wi.id        AS image_id,
			wi.path      AS image_path,
			b.id         AS blog_id,
			b.title      AS blog_title,
			b.body       AS blog_body,
			p.id         AS photo_id,
			p.url        AS photo_url
		FROM women w
		LEFT JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		LEFT JOIN woman_images wi ON wi.woman_id = w.id
		LEFT JOIN blogs b ON b.woman_id = w.id AND b.deleted_at IS NULL AND b.is_published = TRUE
		LEFT JOIN photos p ON p.blog_id = b.id
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE` + where + `
		ORDER BY w.id, wsa.id, wi.id, b.id, p.id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return &entity.NilWoman{}, err
	}
	if len(rows) == 0 {
		return &entity.NilWoman{}, nil
	}

	return mapToWomanOne(rows), nil
}

func mapToWomanOne(rows []map[string]any) entity.WomanEntity {
	base := rows[0]
	womanID := helper.ToUint(base["woman_id"])

	w := &entity.Woman{
		ID:         womanID,
		Name:       helper.ToString(base["woman_name"]),
		Age:        helper.ToIntPtr(base["age"]),
		Birthplace: helper.ToStringPtr(base["birthplace"]),
		BloodType:  helper.ToStringPtr(base["blood_type"]),
		Hobby:      helper.ToStringPtr(base["hobby"]),
	}

	seenAssignments := make(map[uint]bool)
	seenImages := make(map[uint]bool)
	seenBlogs := make(map[uint]bool)
	seenPhotos := make(map[uint]map[uint]bool)

	for _, row := range rows {
		assignmentID := helper.ToUint(row["assignment_id"])
		if assignmentID != 0 && !seenAssignments[assignmentID] {
			seenAssignments[assignmentID] = true
			current := w.StoreAssignments.All()
			current = append(current, entity.WomanStoreAssignment{
				ID:      assignmentID,
				StoreID: helper.ToUint(row["assignment_store_id"]),
			})
			w.StoreAssignments = collection.NewCollection(current)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[imageID] {
			seenImages[imageID] = true
			current := w.Images.All()
			current = append(current, entity.WomanImage{
				ID:   imageID,
				Path: helper.ToString(row["image_path"]),
			})
			w.Images = collection.NewCollection(current)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[blogID] {
			seenBlogs[blogID] = true
			seenPhotos[blogID] = make(map[uint]bool)
			current := w.Blogs.All()
			current = append(current, &entity.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       helper.ToString(row["blog_title"]),
				Body:        helper.ToStringPtr(row["blog_body"]),
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			w.Blogs = collection.NewCollection(current)
		}

		photoID := helper.ToUint(row["photo_id"])
		if photoID != 0 && blogID != 0 && !seenPhotos[blogID][photoID] {
			seenPhotos[blogID][photoID] = true
			blogs := w.Blogs.All()
			for i, b := range blogs {
				if b.GetID() == blogID {
					photos := b.GetPhotos().All()
					photos = append(photos, entity.Photo{
						ID:  photoID,
						URL: helper.ToString(row["photo_url"]),
					})
					blogs[i].(*entity.Blog).Photos = collection.NewCollection(photos)
					break
				}
			}
			w.Blogs = collection.NewCollection(blogs)
		}
	}

	return w
}
