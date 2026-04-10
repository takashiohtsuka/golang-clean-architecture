package repository

import (
	"context"

	"golang-clean-architecture/pkg/frontend/domain/entity"
	storeMapper "golang-clean-architecture/pkg/frontend/adapter/mapper/store"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) outputport.StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")

	sql := `
		SELECT
			s.id            AS store_id,
			s.district_id   AS district_id,
			d.name          AS district_name,
			d.prefecture_id AS prefecture_id,
			p.name          AS prefecture_name,
			p.region_id     AS region_id,
			r.name          AS region_name,
			bt.code         AS business_type_code,
			s.name          AS store_name,
			w.id            AS woman_id,
			w.name          AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			b.id            AS blog_id,
			b.title         AS blog_title
		FROM stores s
		JOIN districts d    ON s.district_id = d.id
		JOIN prefectures p  ON d.prefecture_id = p.id
		JOIN regions r      ON p.region_id = r.id
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
		ORDER BY w.id, b.id`

	allArgs := append([]any{storeWomenLimit, storeBlogsPerWomanLimit}, args...)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, allArgs...).Scan(&rows).Error; err != nil {
		return &entity.NilStore{}, err
	}
	result := storeMapper.MapToAggregate(rows)
	all := result.All()
	if len(all) == 0 {
		return &entity.NilStore{}, nil
	}
	return all[0], nil
}
