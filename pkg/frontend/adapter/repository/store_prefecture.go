package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	storeMapper "golang-clean-architecture/pkg/frontend/adapter/mapper/store"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"

	"gorm.io/gorm"
)

type storePrefectureRepository struct {
	db *gorm.DB
}

func NewStorePrefectureRepository(db *gorm.DB) outputport.StorePrefectureRepository {
	return &storePrefectureRepository{db: db}
}

func (r *storePrefectureRepository) FindAllByPrefecture(ctx context.Context, prefectureID uint) (collection.Collection[entity.StoreEntity], error) {
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
		WHERE s.deleted_at IS NULL AND s.is_active = TRUE
		AND d.prefecture_id = ?
		ORDER BY s.id, w.id, b.id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, storeWomenLimit, storeBlogsPerWomanLimit, prefectureID).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.StoreEntity](nil), err
	}
	return storeMapper.MapToAggregate(rows), nil
}
