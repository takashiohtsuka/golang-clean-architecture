package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
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
	return mapToWomanList(rows), nil
}

func mapToWomanList(rows []map[string]any) collection.Collection[entity.WomanEntity] {
	womanOrder := make([]uint, 0)
	womanMap := make(map[uint]*entity.Woman)
	seenAssignments := make(map[uint]map[uint]bool)
	seenImages := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		womanID := helper.ToUint(row["woman_id"])

		if _, exists := womanMap[womanID]; !exists {
			womanOrder = append(womanOrder, womanID)
			womanMap[womanID] = &entity.Woman{
				ID:         womanID,
				Name:       func() string { s := helper.ToStringPtr(row["woman_name"]); if s != nil { return *s }; return "" }(),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
			}
			seenAssignments[womanID] = make(map[uint]bool)
			seenImages[womanID] = make(map[uint]bool)
			seenBlogs[womanID] = make(map[uint]bool)
		}

		assignmentID := helper.ToUint(row["assignment_id"])
		if assignmentID != 0 && !seenAssignments[womanID][assignmentID] {
			seenAssignments[womanID][assignmentID] = true
			current := womanMap[womanID].StoreAssignments.All()
			current = append(current, entity.WomanStoreAssignment{
				ID:      assignmentID,
				StoreID: helper.ToUint(row["assignment_store_id"]),
			})
			womanMap[womanID].StoreAssignments = collection.NewCollection(current)
		}

		imageID := helper.ToUint(row["image_id"])
		if imageID != 0 && !seenImages[womanID][imageID] {
			seenImages[womanID][imageID] = true
			current := womanMap[womanID].Images.All()
			current = append(current, entity.WomanImage{
				ID:   imageID,
				Path: func() string { s := helper.ToStringPtr(row["image_path"]); if s != nil { return *s }; return "" }(),
			})
			womanMap[womanID].Images = collection.NewCollection(current)
		}

		blogID := helper.ToUint(row["blog_id"])
		if blogID != 0 && !seenBlogs[womanID][blogID] {
			seenBlogs[womanID][blogID] = true
			current := womanMap[womanID].Blogs.All()
			current = append(current, &entity.Blog{
				ID:          blogID,
				WomanID:     womanID,
				Title:       func() string { s := helper.ToStringPtr(row["blog_title"]); if s != nil { return *s }; return "" }(),
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			womanMap[womanID].Blogs = collection.NewCollection(current)
		}
	}

	items := make([]entity.WomanEntity, 0, len(womanOrder))
	for _, wid := range womanOrder {
		items = append(items, womanMap[wid])
	}
	return collection.NewCollection(items)
}

func (r *womanRepository) FindOne(conditions []query.Condition) (entity.WomanEntity, error) {
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
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
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
		Name:       func() string { s := helper.ToStringPtr(base["woman_name"]); if s != nil { return *s }; return "" }(),
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
				Path: func() string { s := helper.ToStringPtr(row["image_path"]); if s != nil { return *s }; return "" }(),
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
				Title:       func() string { s := helper.ToStringPtr(row["blog_title"]); if s != nil { return *s }; return "" }(),
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
						URL: func() string { s := helper.ToStringPtr(row["photo_url"]); if s != nil { return *s }; return "" }(),
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
