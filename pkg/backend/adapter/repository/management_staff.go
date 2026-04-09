package repository

import (
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type managementStaffRepository struct {
	db *gorm.DB
}

func NewManagementStaffRepository(db *gorm.DB) outputport.ManagementStaffRepository {
	return &managementStaffRepository{db: db}
}

const managementStaffSelectSQL = `
	SELECT id, company_id, store_id, name, email, created_at, updated_at, deleted_at
	FROM management_staffs WHERE deleted_at IS NULL`

func toManagementStaffEntity(row map[string]any) *entity.ManagementStaff {
	return &entity.ManagementStaff{
		ID:        helper.ToUint(row["id"]),
		CompanyID: helper.ToUint(row["company_id"]),
		StoreID:   helper.ToUint(row["store_id"]),
		Name:      func() string { s := helper.ToStringPtr(row["name"]); if s != nil { return *s }; return "" }(),
		Email:     func() string { s := helper.ToStringPtr(row["email"]); if s != nil { return *s }; return "" }(),
		CreatedAt: helper.ToTimePtr(row["created_at"]),
		UpdatedAt: helper.ToTimePtr(row["updated_at"]),
		DeletedAt: helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *managementStaffRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(managementStaffSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.ManagementStaffEntity](nil), err
	}

	items := make([]entity.ManagementStaffEntity, len(rows))
	for i, row := range rows {
		items[i] = toManagementStaffEntity(row)
	}
	return collection.NewCollection(items), nil
}

func (r *managementStaffRepository) FindOne(conditions []query.Condition) (entity.ManagementStaffEntity, error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.Raw(managementStaffSelectSQL+where+` LIMIT 1`, args...).Scan(&rows).Error; err != nil {
		return &entity.NilManagementStaff{}, err
	}
	if len(rows) == 0 {
		return &entity.NilManagementStaff{}, nil
	}
	return toManagementStaffEntity(rows[0]), nil
}

func (r *managementStaffRepository) Create(m *entity.ManagementStaff) error {
	sql := `INSERT INTO management_staffs (company_id, store_id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
	return r.db.Exec(sql, m.CompanyID, m.StoreID, m.Name, m.Email).Error
}

func (r *managementStaffRepository) Update(m *entity.ManagementStaff) error {
	sql := `UPDATE management_staffs SET company_id = ?, store_id = ?, name = ?, email = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, m.CompanyID, m.StoreID, m.Name, m.Email, m.ID).Error
}
