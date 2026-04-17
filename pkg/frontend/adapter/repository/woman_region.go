package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	womanMapper "golang-clean-architecture/pkg/frontend/adapter/mapper/woman"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"

	"gorm.io/gorm"
)

const womanRegionPickupLimit = 6

type womanRegionRepository struct {
	db *gorm.DB
}

func NewWomanRegionRepository(db *gorm.DB) outputport.WomanRegionRepository {
	return &womanRegionRepository{db: db}
}

func (r *womanRegionRepository) FindPickupByRegion(ctx context.Context, regionID uint) (collection.Collection[entity.WomanEntity], error) {
	sql := `
		SELECT
			w.id            AS woman_id,
			w.name          AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			wi.id           AS image_id,
			wi.path         AS image_path,
			b.id            AS blog_id,
			b.title         AS blog_title
		FROM women w
		JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		JOIN stores s   ON s.id = wsa.store_id AND s.deleted_at IS NULL AND s.is_active = TRUE
		JOIN regions r      ON s.region_id = r.id
		LEFT JOIN woman_images wi ON wi.woman_id = w.id
		LEFT JOIN (
			SELECT id, woman_id, title,
			       ROW_NUMBER() OVER (PARTITION BY woman_id ORDER BY created_at DESC) AS rn
			FROM blogs
			WHERE deleted_at IS NULL AND is_published = TRUE
		) b ON b.woman_id = w.id AND b.rn <= ?
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE
		AND w.id IN (
			SELECT w2.id
			FROM women w2
			JOIN woman_store_assignments wsa2 ON wsa2.woman_id = w2.id
			JOIN stores s2  ON s2.id = wsa2.store_id AND s2.deleted_at IS NULL AND s2.is_active = TRUE
			JOIN districts d2   ON s2.district_id = d2.id
			JOIN prefectures p2 ON d2.prefecture_id = p2.id
			WHERE w2.deleted_at IS NULL AND w2.is_active = TRUE
			AND p2.region_id = ?
			ORDER BY w2.id
			LIMIT ?
		)
		ORDER BY w.id, wi.id, b.id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, womanBlogsLimit, regionID, womanRegionPickupLimit).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}
	return womanMapper.MapToAggregate(rows), nil
}
