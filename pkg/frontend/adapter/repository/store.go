package repository

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

const (
	storeWomenLimit         = 4
	storeBlogsPerWomanLimit = 3
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) outputport.StoreRepository {
	return &storeRepository{db: db}
}

type storeAggregateRow struct {
	// store
	StoreID          uint   `gorm:"column:store_id"`
	BusinessTypeCode string `gorm:"column:business_type_code"`
	StoreName        string `gorm:"column:store_name"`
	// woman (nullable)
	WomanID    *uint   `gorm:"column:woman_id"`
	WomanName  *string `gorm:"column:woman_name"`
	Age        *int    `gorm:"column:age"`
	Birthplace *string `gorm:"column:birthplace"`
	BloodType  *string `gorm:"column:blood_type"`
	Hobby      *string `gorm:"column:hobby"`
	// blog (nullable)
	BlogID    *uint   `gorm:"column:blog_id"`
	BlogTitle *string `gorm:"column:blog_title"`
}

func (r *storeRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.StoreEntity], error) {
	rows, err := r.query(conditions)
	if err != nil {
		return collection.NewCollection[entity.StoreEntity](nil), err
	}
	return r.mapToAggregate(rows), nil
}

func (r *storeRepository) FindOne(conditions []query.Condition) (entity.StoreEntity, error) {
	rows, err := r.query(conditions)
	if err != nil {
		return &entity.NilStore{}, err
	}
	result := r.mapToAggregate(rows)
	all := result.All()
	if len(all) == 0 {
		return &entity.NilStore{}, nil
	}
	return all[0], nil
}

func (r *storeRepository) query(conditions []query.Condition) ([]storeAggregateRow, error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")

	sql := `
		SELECT
			s.id         AS store_id,
			bt.code      AS business_type_code,
			s.name       AS store_name,
			w.id         AS woman_id,
			w.name       AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			b.id         AS blog_id,
			b.title      AS blog_title
		FROM stores s
		JOIN business_types bt ON s.business_type_id = bt.id
		JOIN contract_plans cp ON s.contract_plan_id = cp.id
		LEFT JOIN (
			SELECT store_id, woman_id,
			       ROW_NUMBER() OVER (PARTITION BY store_id ORDER BY id) AS rn
			FROM woman_store_assignments
		) ranked_wsa ON ranked_wsa.store_id = s.id AND ranked_wsa.rn <= ?
		LEFT JOIN women w ON w.id = ranked_wsa.woman_id AND w.deleted_at IS NULL AND w.is_active = TRUE
		LEFT JOIN (
			SELECT id, woman_id, title,
			       ROW_NUMBER() OVER (PARTITION BY woman_id ORDER BY created_at DESC) AS rn
			FROM blogs
			WHERE deleted_at IS NULL AND is_published = TRUE
		) b ON b.woman_id = w.id AND b.rn <= ?
		WHERE s.deleted_at IS NULL AND s.is_active = TRUE` + where + `
		ORDER BY s.id, w.id, b.id`

	allArgs := append([]any{storeWomenLimit, storeBlogsPerWomanLimit}, args...)

	var rows []storeAggregateRow
	if err := r.db.Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *storeRepository) mapToAggregate(rows []storeAggregateRow) collection.Collection[entity.StoreEntity] {
	storeOrder := make([]uint, 0)
	storeMap := make(map[uint]*entity.Store)
	womanOrderByStore := make(map[uint][]uint)
	womanMap := make(map[uint]*entity.Woman)
	seenWomenByStore := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		if _, exists := storeMap[row.StoreID]; !exists {
			storeOrder = append(storeOrder, row.StoreID)
			storeMap[row.StoreID] = &entity.Store{
				ID:           row.StoreID,
				BusinessType: fvo.NewBusinessType(row.BusinessTypeCode),
				Name:         row.StoreName,
			}
			womanOrderByStore[row.StoreID] = make([]uint, 0)
		}

		if row.WomanID == nil {
			continue
		}
		womanID := *row.WomanID

		if seenWomenByStore[row.StoreID] == nil {
			seenWomenByStore[row.StoreID] = make(map[uint]bool)
		}
		if !seenWomenByStore[row.StoreID][womanID] {
			seenWomenByStore[row.StoreID][womanID] = true
			womanOrderByStore[row.StoreID] = append(womanOrderByStore[row.StoreID], womanID)
		}

		if _, exists := womanMap[womanID]; !exists {
			womanMap[womanID] = &entity.Woman{
				ID:         womanID,
				Name:       *row.WomanName,
				Age:        row.Age,
				Birthplace: row.Birthplace,
				BloodType:  row.BloodType,
				Hobby:      row.Hobby,
			}
			seenBlogs[womanID] = make(map[uint]bool)
		}

		if row.BlogID != nil && !seenBlogs[womanID][*row.BlogID] {
			seenBlogs[womanID][*row.BlogID] = true
			current := womanMap[womanID].Blogs.All()
			current = append(current, &entity.Blog{
				ID:          *row.BlogID,
				WomanID:     womanID,
				Title:       *row.BlogTitle,
				IsPublished: true,
				Photos:      collection.NewCollection[entity.Photo](nil),
			})
			womanMap[womanID].Blogs = collection.NewCollection(current)
		}
	}

	for storeID, womanIDs := range womanOrderByStore {
		women := make([]entity.WomanEntity, 0, len(womanIDs))
		for _, wid := range womanIDs {
			women = append(women, womanMap[wid])
		}
		storeMap[storeID].Women = collection.NewCollection(women)
	}

	items := make([]entity.StoreEntity, 0, len(storeOrder))
	for _, storeID := range storeOrder {
		items = append(items, storeMap[storeID])
	}
	return collection.NewCollection(items)
}
