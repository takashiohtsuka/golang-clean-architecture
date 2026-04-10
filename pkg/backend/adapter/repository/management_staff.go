package repository

import (
	"context"

	managementstaffMapper "golang-clean-architecture/pkg/backend/adapter/mapper/managementstaff"
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
		Name:      helper.ToString(row["name"]),
		Email:     helper.ToString(row["email"]),
		CreatedAt: helper.ToTimePtr(row["created_at"]),
		UpdatedAt: helper.ToTimePtr(row["updated_at"]),
		DeletedAt: helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *managementStaffRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(managementStaffSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.ManagementStaffEntity](nil), err
	}

	items := make([]entity.ManagementStaffEntity, len(rows))
	for i, row := range rows {
		items[i] = toManagementStaffEntity(row)
	}
	return collection.NewCollection(items), nil
}

func (r *managementStaffRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.ManagementStaffEntity, error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(managementStaffSelectSQL+where+` LIMIT 1`, args...).Scan(&rows).Error; err != nil {
		return &entity.NilManagementStaff{}, err
	}
	if len(rows) == 0 {
		return &entity.NilManagementStaff{}, nil
	}
	return toManagementStaffEntity(rows[0]), nil
}

func (r *managementStaffRepository) Create(ctx context.Context, e *entity.ManagementStaff) error {
	return r.db.WithContext(ctx).Create(managementstaffMapper.ToOrmModel(e)).Error
}

func (r *managementStaffRepository) Update(ctx context.Context, e *entity.ManagementStaff) error {
	return r.db.WithContext(ctx).Save(managementstaffMapper.ToOrmModel(e)).Error
}
