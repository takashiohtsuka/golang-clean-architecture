package repository

import (
	"context"

	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) outputport.StoreRepository {
	return &storeRepository{db: db}
}

const storeSelectSQL = `
	SELECT
		s.id,
		s.company_id,
		bt.code  AS business_type_code,
		cp.code  AS contract_plan_code,
		s.name,
		s.is_active,
		s.open_status,
		s.created_at,
		s.updated_at,
		s.deleted_at
	FROM stores s
	JOIN business_types bt ON s.business_type_id = bt.id
	JOIN contract_plans cp ON s.contract_plan_id = cp.id
	WHERE s.deleted_at IS NULL`

func toStoreEntity(row map[string]any) *entity.Store {
	return &entity.Store{
		ID:           helper.ToUint(row["id"]),
		CompanyID:    helper.ToUint(row["company_id"]),
		BusinessType: bvo.NewBusinessType(func() string { s := helper.ToStringPtr(row["business_type_code"]); if s != nil { return *s }; return "" }()),
		ContractPlan: bvo.NewContractPlan(func() string { s := helper.ToStringPtr(row["contract_plan_code"]); if s != nil { return *s }; return "" }()),
		Name:         func() string { s := helper.ToStringPtr(row["name"]); if s != nil { return *s }; return "" }(),
		IsActive:     helper.ToBool(row["is_active"]),
		OpenStatus:   entity.OpenStatus(func() string { s := helper.ToStringPtr(row["open_status"]); if s != nil { return *s }; return "" }()),
		CreatedAt:    helper.ToTimePtr(row["created_at"]),
		UpdatedAt:    helper.ToTimePtr(row["updated_at"]),
		DeletedAt:    helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *storeRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.StoreEntity], error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")

	var rows []map[string]any
	if err := r.db.Raw(storeSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.StoreEntity](nil), err
	}

	items := make([]entity.StoreEntity, len(rows))
	for i, row := range rows {
		items[i] = toStoreEntity(row)
	}
	return collection.NewCollection(items), nil
}

func (r *storeRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(storeSelectSQL+where+` LIMIT 1`, args...).Scan(&rows).Error; err != nil {
		return &entity.NilStore{}, err
	}
	if len(rows) == 0 {
		return &entity.NilStore{}, nil
	}
	return toStoreEntity(rows[0]), nil
}

func (r *storeRepository) Create(s *entity.Store) error {
	btID, cpID, err := r.resolveIDs(s)
	if err != nil {
		return err
	}
	sql := `INSERT INTO stores (company_id, business_type_id, contract_plan_id, name, is_active, open_status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW())`
	return r.db.Exec(sql, s.CompanyID, btID, cpID, s.Name, s.IsActive, string(s.OpenStatus)).Error
}

func (r *storeRepository) Update(s *entity.Store) error {
	btID, cpID, err := r.resolveIDs(s)
	if err != nil {
		return err
	}
	sql := `UPDATE stores SET company_id = ?, business_type_id = ?, contract_plan_id = ?, name = ?, is_active = ?, open_status = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL`
	return r.db.Exec(sql, s.CompanyID, btID, cpID, s.Name, s.IsActive, string(s.OpenStatus), s.ID).Error
}

func (r *storeRepository) AddWoman(womanID uint, storeID uint) error {
	sql := `INSERT INTO woman_store_assignments (woman_id, store_id, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`
	return r.db.Exec(sql, womanID, storeID).Error
}

func (r *storeRepository) RemoveWoman(womanID uint, storeID uint) error {
	sql := `DELETE FROM woman_store_assignments WHERE woman_id = ? AND store_id = ?`
	return r.db.Exec(sql, womanID, storeID).Error
}

// resolveIDs はStore entityのVOからbusiness_type_idとcontract_plan_idを取得する。
func (r *storeRepository) resolveIDs(s *entity.Store) (btID uint, cpID uint, err error) {
	if err = r.db.Raw(`SELECT id FROM business_types WHERE code = ? LIMIT 1`, s.BusinessType.GetCode()).Scan(&btID).Error; err != nil {
		return
	}
	err = r.db.Raw(`SELECT id FROM contract_plans WHERE code = ? LIMIT 1`, s.ContractPlan.GetCode()).Scan(&cpID).Error
	return
}
