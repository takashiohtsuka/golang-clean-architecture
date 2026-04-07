package repository

import (
	"context"
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type womanRepository struct {
	db *gorm.DB
}

func NewWomanRepository(db *gorm.DB) outputport.WomanRepository {
	return &womanRepository{db: db}
}

type womanRow struct {
	ID         uint       `gorm:"column:id"`
	CompanyID  uint       `gorm:"column:company_id"`
	Name       string     `gorm:"column:name"`
	Age        *int       `gorm:"column:age"`
	Birthplace *string    `gorm:"column:birthplace"`
	BloodType  *string    `gorm:"column:blood_type"`
	Hobby      *string    `gorm:"column:hobby"`
	IsActive   bool       `gorm:"column:is_active"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (r *womanRow) toEntity() *entity.Woman {
	return &entity.Woman{
		ID:         r.ID,
		CompanyID:  r.CompanyID,
		Name:       r.Name,
		Age:        r.Age,
		Birthplace: r.Birthplace,
		BloodType:  r.BloodType,
		Hobby:      r.Hobby,
		IsActive:   r.IsActive,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
		DeletedAt:  r.DeletedAt,
	}
}

type assignmentRow struct {
	ID        uint       `gorm:"column:id"`
	StoreID   uint       `gorm:"column:store_id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (r *womanRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.WomanEntity], error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, company_id, name, age, birthplace, blood_type, hobby, is_active, created_at, updated_at, deleted_at
	        FROM women WHERE deleted_at IS NULL` + where

	var rows []womanRow
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}

	womanIDs := make([]uint, len(rows))
	for i, row := range rows {
		womanIDs[i] = row.ID
	}

	assignmentsByWomanID, err := r.findAssignmentsByWomanIDs(womanIDs)
	if err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}

	items := make([]entity.WomanEntity, len(rows))
	for i, row := range rows {
		e := row.toEntity()
		e.StoreAssignments = collection.NewCollection(assignmentsByWomanID[row.ID])
		items[i] = e
	}
	return collection.NewCollection(items), nil
}

func (r *womanRepository) FindOne(conditions []query.Condition) (entity.WomanEntity, error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, company_id, name, age, birthplace, blood_type, hobby, is_active, created_at, updated_at, deleted_at
	        FROM women WHERE deleted_at IS NULL` + where + ` LIMIT 1`

	var row womanRow
	if err := r.db.Raw(sql, args...).Scan(&row).Error; err != nil {
		return &entity.NilWoman{}, err
	}
	if row.ID == 0 {
		return &entity.NilWoman{}, nil
	}

	e := row.toEntity()
	assignmentsByWomanID, err := r.findAssignmentsByWomanIDs([]uint{row.ID})
	if err != nil {
		return &entity.NilWoman{}, err
	}
	e.StoreAssignments = collection.NewCollection(assignmentsByWomanID[row.ID])
	return e, nil
}

func (r *womanRepository) Create(ctx context.Context, w *entity.Woman) error {
	db := r.db.WithContext(ctx)
	sql := `INSERT INTO women (company_id, name, age, birthplace, blood_type, hobby, is_active, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	if err := db.Exec(sql, w.CompanyID, w.Name, w.Age, w.Birthplace, w.BloodType, w.Hobby, w.IsActive).Error; err != nil {
		return err
	}

	if w.StoreAssignments.TotalCount() == 0 {
		return nil
	}

	var womanID uint
	if err := db.Raw("SELECT LAST_INSERT_ID()").Scan(&womanID).Error; err != nil {
		return err
	}

	for _, a := range w.StoreAssignments.All() {
		asql := `INSERT INTO woman_store_assignments (woman_id, store_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
		if err := db.Exec(asql, womanID, a.StoreID).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *womanRepository) Update(w *entity.Woman) error {
	sql := `UPDATE women SET company_id = ?, name = ?, age = ?, birthplace = ?, blood_type = ?, hobby = ?, is_active = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, w.CompanyID, w.Name, w.Age, w.Birthplace, w.BloodType, w.Hobby, w.IsActive, w.ID).Error
}

func (r *womanRepository) AssignToStore(womanID uint, storeID uint) error {
	sql := `INSERT INTO woman_store_assignments (woman_id, store_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
	return r.db.Exec(sql, womanID, storeID).Error
}

func (r *womanRepository) RemoveStoreAssignment(womanID uint, storeID uint) error {
	sql := `DELETE FROM woman_store_assignments WHERE woman_id = ? AND store_id = ?`
	return r.db.Exec(sql, womanID, storeID).Error
}

// findAssignmentsByWomanIDs は複数のwomanIDに対するassignmentをまとめて取得しIDでグループ化する。
func (r *womanRepository) findAssignmentsByWomanIDs(womanIDs []uint) (map[uint][]entity.WomanStoreAssignment, error) {
	result := make(map[uint][]entity.WomanStoreAssignment)
	if len(womanIDs) == 0 {
		return result, nil
	}

	type wsaRow struct {
		WomanID   uint       `gorm:"column:woman_id"`
		ID        uint       `gorm:"column:id"`
		StoreID   uint       `gorm:"column:store_id"`
		CreatedAt *time.Time `gorm:"column:created_at"`
		UpdatedAt *time.Time `gorm:"column:updated_at"`
	}

	var rows []wsaRow
	sql := `SELECT id, woman_id, store_id, created_at, updated_at FROM woman_store_assignments WHERE woman_id IN ?`
	if err := r.db.Raw(sql, womanIDs).Scan(&rows).Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		result[row.WomanID] = append(result[row.WomanID], entity.WomanStoreAssignment{
			ID:        row.ID,
			StoreID:   row.StoreID,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		})
	}
	return result, nil
}
