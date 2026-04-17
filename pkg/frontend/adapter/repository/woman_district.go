package repository

import (
	"context"

	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/frontend/domain/entity"
	womanMapper "golang-clean-architecture/pkg/frontend/adapter/mapper/woman"
	"golang-clean-architecture/pkg/frontend/usecase/outputport"

	"gorm.io/gorm"
)

type womanDistrictRepository struct {
	db *gorm.DB
}

func NewWomanDistrictRepository(db *gorm.DB) outputport.WomanDistrictRepository {
	return &womanDistrictRepository{db: db}
}

func (r *womanDistrictRepository) FindAllByDistrict(ctx context.Context, districtID uint) (collection.Collection[entity.WomanEntity], error) {
	sql := `
		SELECT
			w.id            AS woman_id,
			w.name          AS woman_name,
			w.age,
			w.birthplace,
			w.blood_type,
			w.hobby,
			wsa.id          AS assignment_id,
			wsa.store_id    AS assignment_store_id,
			wi.id           AS image_id,
			wi.path         AS image_path,
			b.id            AS blog_id,
			b.title         AS blog_title
		FROM women w
		JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
		JOIN stores s   ON s.id = wsa.store_id AND s.deleted_at IS NULL AND s.is_active = TRUE
		JOIN districts d    ON s.district_id = d.id
		JOIN prefectures p  ON d.prefecture_id = p.id
		JOIN regions r      ON p.region_id = r.id
		LEFT JOIN woman_images wi ON wi.woman_id = w.id
		LEFT JOIN (
			SELECT id, woman_id, title,
			       ROW_NUMBER() OVER (PARTITION BY woman_id ORDER BY created_at DESC) AS rn
			FROM blogs
			WHERE deleted_at IS NULL AND is_published = TRUE
		) b ON b.woman_id = w.id AND b.rn <= ?
		WHERE w.deleted_at IS NULL AND w.is_active = TRUE
		AND d.id = ?
		ORDER BY w.id, wi.id, b.id`

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, womanBlogsLimit, districtID).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}
	return womanMapper.MapToAggregate(rows), nil
}
