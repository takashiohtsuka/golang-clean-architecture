package repository

import (
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	fvo "golang-clean-architecture/pkg/frontend/domain/valueobject"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/helper"
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

func (r *storeRepository) query(conditions []query.Condition) ([]map[string]any, error) {
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

	var rows []map[string]any
	if err := r.db.Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *storeRepository) mapToAggregate(rows []map[string]any) collection.Collection[entity.StoreEntity] {
	storeOrder := make([]uint, 0)
	storeMap := make(map[uint]*entity.Store)
	womanOrderByStore := make(map[uint][]uint)
	womanMap := make(map[uint]*entity.Woman)
	seenWomenByStore := make(map[uint]map[uint]bool)
	seenBlogs := make(map[uint]map[uint]bool)

	for _, row := range rows {
		storeID := helper.ToUint(row["store_id"])

		if _, exists := storeMap[storeID]; !exists {
			storeOrder = append(storeOrder, storeID)
			storeMap[storeID] = &entity.Store{
				ID:           storeID,
				BusinessType: fvo.NewBusinessType(func() string { s := helper.ToStringPtr(row["business_type_code"]); if s != nil { return *s }; return "" }()),
				Name:         func() string { s := helper.ToStringPtr(row["store_name"]); if s != nil { return *s }; return "" }(),
			}
			womanOrderByStore[storeID] = make([]uint, 0)
		}

		womanID := helper.ToUint(row["woman_id"])
		if womanID == 0 {
			continue
		}

		if seenWomenByStore[storeID] == nil {
			seenWomenByStore[storeID] = make(map[uint]bool)
		}
		if !seenWomenByStore[storeID][womanID] {
			seenWomenByStore[storeID][womanID] = true
			womanOrderByStore[storeID] = append(womanOrderByStore[storeID], womanID)
		}

		if _, exists := womanMap[womanID]; !exists {
			womanMap[womanID] = &entity.Woman{
				ID:         womanID,
				Name:       func() string { s := helper.ToStringPtr(row["woman_name"]); if s != nil { return *s }; return "" }(),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
			}
			seenBlogs[womanID] = make(map[uint]bool)
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
