package repository

import (
	"context"
	"strings"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type womanRepository struct {
	db *gorm.DB
}

func NewWomanRepository(db *gorm.DB) outputport.WomanRepository {
	return &womanRepository{db: db}
}

const womanJoinSQL = `
	SELECT
		w.id, w.company_id, w.name, w.age, w.birthplace, w.blood_type, w.hobby, w.is_active,
		w.created_at, w.updated_at, w.deleted_at,
		wi.id AS wi_id, wi.woman_id AS wi_woman_id, wi.path AS wi_path,
		wi.created_at AS wi_created_at, wi.updated_at AS wi_updated_at
	FROM women w
	LEFT JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
	LEFT JOIN woman_images wi ON wi.woman_id = w.id
	WHERE w.deleted_at IS NULL`

func buildWomanWhereClause(conditions []query.Condition) (string, []any) {
	tablePrefix := func(column string) string {
		switch column {
		case "store_id":
			return "wsa." + column
		default:
			return "w." + column
		}
	}

	var clauses []string
	var args []any

	for _, c := range conditions {
		col := tablePrefix(c.Column)
		switch c.Kind {
		case query.KindWhere:
			clauses = append(clauses, col+" = ?")
			args = append(args, c.Value)
		case query.KindWhereIn:
			clauses = append(clauses, col+" IN ?")
			args = append(args, c.Value)
		case query.KindWhereBetween:
			clauses = append(clauses, col+" BETWEEN ? AND ?")
			args = append(args, c.From, c.To)
		case query.KindWhereNotIn:
			clauses = append(clauses, col+" NOT IN ?")
			args = append(args, c.Value)
		}
	}

	if len(clauses) == 0 {
		return "", nil
	}
	return " AND " + strings.Join(clauses, " AND "), args
}

func groupWomen(rows []map[string]any) []*entity.Woman {
	womanMap := make(map[uint]*entity.Woman)
	womanOrder := make([]uint, 0)
	imageSeen := make(map[uint]map[uint]bool)
	imagesByID := make(map[uint][]entity.WomanImage)

	for _, row := range rows {
		womanID := helper.ToUint(row["id"])

		if _, ok := womanMap[womanID]; !ok {
			womanMap[womanID] = &entity.Woman{
				ID:        womanID,
				CompanyID: helper.ToUint(row["company_id"]),
				Name: func() string {
					s := helper.ToStringPtr(row["name"])
					if s != nil {
						return *s
					}
					return ""
				}(),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
				IsActive:   helper.ToBool(row["is_active"]),
				CreatedAt:  helper.ToTimePtr(row["created_at"]),
				UpdatedAt:  helper.ToTimePtr(row["updated_at"]),
				DeletedAt:  helper.ToTimePtr(row["deleted_at"]),
			}
			imageSeen[womanID] = make(map[uint]bool)
			womanOrder = append(womanOrder, womanID)
		}

		imageID := helper.ToUint(row["wi_id"])
		if imageID != 0 && !imageSeen[womanID][imageID] {
			imageSeen[womanID][imageID] = true
			imagesByID[womanID] = append(imagesByID[womanID], entity.WomanImage{
				ID:      imageID,
				WomanID: helper.ToUint(row["wi_woman_id"]),
				Path: func() string {
					s := helper.ToStringPtr(row["wi_path"])
					if s != nil {
						return *s
					}
					return ""
				}(),
				CreatedAt: helper.ToTimePtr(row["wi_created_at"]),
				UpdatedAt: helper.ToTimePtr(row["wi_updated_at"]),
			})
		}
	}

	result := make([]*entity.Woman, len(womanOrder))
	for i, id := range womanOrder {
		w := womanMap[id]
		w.Images = collection.NewCollection(imagesByID[id])
		result[i] = w
	}
	return result
}

func (r *womanRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.WomanEntity], error) {
	where, args := buildWomanWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(womanJoinSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}

	women := groupWomen(rows)
	items := make([]entity.WomanEntity, len(women))
	for i, w := range women {
		items[i] = w
	}
	return collection.NewCollection(items), nil
}

func (r *womanRepository) FindOne(conditions []query.Condition) (entity.WomanEntity, error) {
	where, args := buildWomanWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(womanJoinSQL+where, args...).Scan(&rows).Error; err != nil {
		return &entity.NilWoman{}, err
	}

	women := groupWomen(rows)
	if len(women) == 0 {
		return &entity.NilWoman{}, nil
	}
	return women[0], nil
}

func (r *womanRepository) Create(ctx context.Context, w *entity.Woman) (uint, error) {
	db := r.db.WithContext(ctx)
	sql := `INSERT INTO women (company_id, name, age, birthplace, blood_type, hobby, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	if err := db.Exec(sql, w.CompanyID, w.Name, w.Age, w.Birthplace, w.BloodType, w.Hobby, w.IsActive).Error; err != nil {
		return 0, err
	}
	var womanID uint
	if err := db.Raw("SELECT LAST_INSERT_ID()").Scan(&womanID).Error; err != nil {
		return 0, err
	}
	return womanID, nil
}

func (r *womanRepository) Update(ctx context.Context, w *entity.Woman) error {
	sql := `UPDATE women SET company_id = ?, name = ?, age = ?, birthplace = ?, blood_type = ?, hobby = ?, is_active = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.WithContext(ctx).Exec(sql, w.CompanyID, w.Name, w.Age, w.Birthplace, w.BloodType, w.Hobby, w.IsActive, w.ID).Error
}

func (r *womanRepository) SaveImage(ctx context.Context, womanID uint, path string) error {
	sql := `INSERT INTO woman_images (woman_id, path, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
	return r.db.WithContext(ctx).Exec(sql, womanID, path).Error
}
