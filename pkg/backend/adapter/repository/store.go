package repository

import (
	"context"
	"time"

	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) outputport.StoreRepository {
	return &storeRepository{db: db}
}

type storeRow struct {
	ID               uint       `gorm:"column:id"`
	CompanyID        uint       `gorm:"column:company_id"`
	BusinessTypeCode string     `gorm:"column:business_type_code"`
	ContractPlanCode string     `gorm:"column:contract_plan_code"`
	Name             string     `gorm:"column:name"`
	IsActive         bool       `gorm:"column:is_active"`
	OpenStatus       string     `gorm:"column:open_status"`
	CreatedAt        *time.Time `gorm:"column:created_at"`
	UpdatedAt        *time.Time `gorm:"column:updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at"`
}

func (r *storeRow) toEntity() *entity.Store {
	return &entity.Store{
		ID:           r.ID,
		CompanyID:    r.CompanyID,
		BusinessType: bvo.NewBusinessType(r.BusinessTypeCode),
		ContractPlan: bvo.NewContractPlan(r.ContractPlanCode),
		Name:         r.Name,
		IsActive:     r.IsActive,
		OpenStatus:   entity.OpenStatus(r.OpenStatus),
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
		DeletedAt:    r.DeletedAt,
	}
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

func (r *storeRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.StoreEntity], error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")
	var rows []storeRow
	if err := r.db.Raw(storeSelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.StoreEntity](nil), err
	}

	items := make([]entity.StoreEntity, len(rows))
	for i, row := range rows {
		items[i] = row.toEntity()
	}
	return collection.NewCollection(items), nil
}

func (r *storeRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.StoreEntity, error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")
	var row storeRow
	if err := r.db.WithContext(ctx).Raw(storeSelectSQL+where+` LIMIT 1`, args...).Scan(&row).Error; err != nil {
		return &entity.NilStore{}, err
	}
	if row.ID == 0 {
		return &entity.NilStore{}, nil
	}
	return row.toEntity(), nil
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

// resolveIDs はStore entityのVOからbusiness_type_idとcontract_plan_idを取得する。
func (r *storeRepository) resolveIDs(s *entity.Store) (btID uint, cpID uint, err error) {
	if err = r.db.Raw(`SELECT id FROM business_types WHERE code = ? LIMIT 1`, s.BusinessType.GetCode()).Scan(&btID).Error; err != nil {
		return
	}
	err = r.db.Raw(`SELECT id FROM contract_plans WHERE code = ? LIMIT 1`, s.ContractPlan.GetCode()).Scan(&cpID).Error
	return
}
