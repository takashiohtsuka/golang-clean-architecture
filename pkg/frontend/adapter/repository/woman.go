package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
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

// --- row structs ---

type womanListRow struct {
	WomanID           uint    `gorm:"column:woman_id"`
	WomanName         string  `gorm:"column:woman_name"`
	Age               *int    `gorm:"column:age"`
	Birthplace        *string `gorm:"column:birthplace"`
	BloodType         *string `gorm:"column:blood_type"`
	Hobby             *string `gorm:"column:hobby"`
	AssignmentID      *uint   `gorm:"column:assignment_id"`
	AssignmentStoreID *uint   `gorm:"column:assignment_store_id"`
	ImageID           *uint   `gorm:"column:image_id"`
	ImagePath         *string `gorm:"column:image_path"`
	BlogID            *uint   `gorm:"column:blog_id"`
	BlogTitle         *string `gorm:"column:blog_title"`
}

type womanDetailRow struct {
	WomanID           uint    `gorm:"column:woman_id"`
	WomanName         string  `gorm:"column:woman_name"`
	Age               *int    `gorm:"column:age"`
	Birthplace        *string `gorm:"column:birthplace"`
	BloodType         *string `gorm:"column:blood_type"`
	Hobby             *string `gorm:"column:hobby"`
	AssignmentID      *uint   `gorm:"column:assignment_id"`
	AssignmentStoreID *uint   `gorm:"column:assignment_store_id"`
	ImageID           *uint   `gorm:"column:image_id"`
	ImagePath         *string `gorm:"column:image_path"`
	BlogID            *uint   `gorm:"column:blog_id"`
	BlogTitle         *string `gorm:"column:blog_title"`
	BlogBody          *string `gorm:"column:blog_body"`
	PhotoID           *uint   `gorm:"column:photo_id"`
	PhotoURL          *string `gorm:"column:photo_url"`
}

// --- FindAll ---

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

	var rows []womanListRow
	if err := r.db.WithContext(ctx).Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}
	return mapToWomanList(rows), nil
}

func mapToWomanList(rows []womanListRow) collection.Collection[entity.WomanEntity] {
	womanOrder := make([]uint, 0)
	womanMap := make(map[uint]*entity.Woman)
	seenAssignments := make(map[uint]map[uint]bool)
	seenImages := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		if _, exists := womanMap[row.WomanID]; !exists {
			womanOrder = append(womanOrder, row.WomanID)
			womanMap[row.WomanID] = &entity.Woman{
				ID:         row.WomanID,
				Name:       row.WomanName,
				Age:        row.Age,
				Birthplace: row.Birthplace,
				BloodType:  row.BloodType,
				Hobby:      row.Hobby,
			}
			seenAssignments[row.WomanID] = make(map[uint]bool)
			seenImages[row.WomanID] = make(map[uint]bool)
			seenBlogs[row.WomanID] = make(map[uint]bool)
		}

		if row.AssignmentID != nil && !seenAssignments[row.WomanID][*row.AssignmentID] {
			seenAssignments[row.WomanID][*row.AssignmentID] = true
			current := womanMap[row.WomanID].StoreAssignments.All()
			current = append(current, entity.WomanStoreAssignment{
				ID:      *row.AssignmentID,
				StoreID: *row.AssignmentStoreID,
			})
			womanMap[row.WomanID].StoreAssignments = collection.NewCollection(current)
		}

		if row.ImageID != nil && !seenImages[row.WomanID][*row.ImageID] {
			seenImages[row.WomanID][*row.ImageID] = true
			current := womanMap[row.WomanID].Images.All()
			current = append(current, entity.WomanImage{
				ID:   *row.ImageID,
				Path: *row.ImagePath,
			})
			womanMap[row.WomanID].Images = collection.NewCollection(current)
		}

		if row.BlogID != nil && !seenBlogs[row.WomanID][*row.BlogID] {
			seenBlogs[row.WomanID][*row.BlogID] = true
			current := womanMap[row.WomanID].Blogs.All()
			current = append(current, &entity.Blog{
				ID:          *row.BlogID,
				WomanID:     row.WomanID,
				Title:       *row.BlogTitle,
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			womanMap[row.WomanID].Blogs = collection.NewCollection(current)
		}
	}

	items := make([]entity.WomanEntity, 0, len(womanOrder))
	for _, wid := range womanOrder {
		items = append(items, womanMap[wid])
	}
	return collection.NewCollection(items)
}

// --- FindOne ---

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

	var rows []womanDetailRow
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return &entity.NilWoman{}, err
	}
	if len(rows) == 0 {
		return &entity.NilWoman{}, nil
	}
	return mapToWomanOne(rows), nil
}

func mapToWomanOne(rows []womanDetailRow) entity.WomanEntity {
	base := rows[0]
	w := &entity.Woman{
		ID:         base.WomanID,
		Name:       base.WomanName,
		Age:        base.Age,
		Birthplace: base.Birthplace,
		BloodType:  base.BloodType,
		Hobby:      base.Hobby,
	}

	seenAssignments := make(map[uint]bool)
	seenImages := make(map[uint]bool)
	seenBlogs := make(map[uint]bool)
	seenPhotos := make(map[uint]map[uint]bool)

	for _, row := range rows {
		if row.AssignmentID != nil && !seenAssignments[*row.AssignmentID] {
			seenAssignments[*row.AssignmentID] = true
			current := w.StoreAssignments.All()
			current = append(current, entity.WomanStoreAssignment{
				ID:      *row.AssignmentID,
				StoreID: *row.AssignmentStoreID,
			})
			w.StoreAssignments = collection.NewCollection(current)
		}

		if row.ImageID != nil && !seenImages[*row.ImageID] {
			seenImages[*row.ImageID] = true
			current := w.Images.All()
			current = append(current, entity.WomanImage{
				ID:   *row.ImageID,
				Path: *row.ImagePath,
			})
			w.Images = collection.NewCollection(current)
		}

		if row.BlogID != nil && !seenBlogs[*row.BlogID] {
			seenBlogs[*row.BlogID] = true
			seenPhotos[*row.BlogID] = make(map[uint]bool)
			current := w.Blogs.All()
			current = append(current, &entity.Blog{
				ID:          *row.BlogID,
				WomanID:     base.WomanID,
				Title:       *row.BlogTitle,
				Body:        row.BlogBody,
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			w.Blogs = collection.NewCollection(current)
		}

		if row.PhotoID != nil && row.BlogID != nil && !seenPhotos[*row.BlogID][*row.PhotoID] {
			seenPhotos[*row.BlogID][*row.PhotoID] = true
			blogs := w.Blogs.All()
			for i, b := range blogs {
				if b.GetID() == *row.BlogID {
					photos := b.GetPhotos().All()
					photos = append(photos, entity.Photo{
						ID:  *row.PhotoID,
						URL: *row.PhotoURL,
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
