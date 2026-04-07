package repository

import (
	"time"

	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type managementStaffRepository struct {
	db *gorm.DB
}

func NewManagementStaffRepository(db *gorm.DB) outputport.ManagementStaffRepository {
	return &managementStaffRepository{db: db}
}

type managementStaffRow struct {
	ID        uint       `gorm:"column:id"`
	CompanyID uint       `gorm:"column:company_id"`
	StoreID   uint       `gorm:"column:store_id"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (r *managementStaffRow) toEntity() *entity.ManagementStaff {
	return &entity.ManagementStaff{
		ID:        r.ID,
		CompanyID: r.CompanyID,
		StoreID:   r.StoreID,
		Name:      r.Name,
		Email:     r.Email,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		DeletedAt: r.DeletedAt,
	}
}

func (r *managementStaffRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, company_id, store_id, name, email, created_at, updated_at, deleted_at
	        FROM management_staffs WHERE deleted_at IS NULL` + where

	var rows []managementStaffRow
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.ManagementStaffEntity](nil), err
	}

	items := make([]entity.ManagementStaffEntity, len(rows))
	for i, row := range rows {
		items[i] = row.toEntity()
	}
	return collection.NewCollection(items), nil
}

func (r *managementStaffRepository) FindOne(conditions []query.Condition) (entity.ManagementStaffEntity, error) {
	where, args := buildWhereClause(conditions)
	sql := `SELECT id, company_id, store_id, name, email, created_at, updated_at, deleted_at
	        FROM management_staffs WHERE deleted_at IS NULL` + where + ` LIMIT 1`

	var row managementStaffRow
	if err := r.db.Raw(sql, args...).Scan(&row).Error; err != nil {
		return &entity.NilManagementStaff{}, err
	}
	if row.ID == 0 {
		return &entity.NilManagementStaff{}, nil
	}
	return row.toEntity(), nil
}

func (r *managementStaffRepository) Create(m *entity.ManagementStaff) error {
	sql := `INSERT INTO management_staffs (company_id, store_id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
	return r.db.Exec(sql, m.CompanyID, m.StoreID, m.Name, m.Email).Error
}

func (r *managementStaffRepository) Update(m *entity.ManagementStaff) error {
	sql := `UPDATE management_staffs SET company_id = ?, store_id = ?, name = ?, email = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, m.CompanyID, m.StoreID, m.Name, m.Email, m.ID).Error
}
