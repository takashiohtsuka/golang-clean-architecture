package repository

import (
	"context"

	storeMapper "golang-clean-architecture/pkg/backend/adapter/mapper/store"
	bvo "golang-clean-architecture/pkg/backend/domain/valueobject"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/infrastructure/model"
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
		s.district_id,
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
		DistrictID:   helper.ToUint(row["district_id"]),
		BusinessType: bvo.NewBusinessType(helper.ToString(row["business_type_code"])),
		ContractPlan: bvo.NewContractPlan(helper.ToString(row["contract_plan_code"])),
		Name:         helper.ToString(row["name"]),
		IsActive:     helper.ToBool(row["is_active"]),
		OpenStatus:   entity.OpenStatus(helper.ToString(row["open_status"])),
		CreatedAt:    helper.ToTimePtr(row["created_at"]),
		UpdatedAt:    helper.ToTimePtr(row["updated_at"]),
		DeletedAt:    helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *storeRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.StoreEntity], error) {
	where, args := buildWhereClauseWithPrefix(conditions, "s")

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(storeSelectSQL+where, args...).Scan(&rows).Error; err != nil {
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

func (r *storeRepository) Create(ctx context.Context, s *entity.Store) error {
	m := storeMapper.ToOrmModel(s)
	var err error
	m.BusinessTypeID, m.ContractPlanID, err = r.resolveIDs(ctx, s)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *storeRepository) Update(ctx context.Context, s *entity.Store) error {
	m := storeMapper.ToOrmModel(s)
	var err error
	m.BusinessTypeID, m.ContractPlanID, err = r.resolveIDs(ctx, s)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *storeRepository) AddWoman(ctx context.Context, womanID uint, storeID uint) error {
	return r.db.WithContext(ctx).Create(&model.WomanStoreAssignment{WomanID: womanID, StoreID: storeID}).Error
}

func (r *storeRepository) RemoveWoman(ctx context.Context, womanID uint, storeID uint) error {
	return r.db.WithContext(ctx).Where("woman_id = ? AND store_id = ?", womanID, storeID).Delete(&model.WomanStoreAssignment{}).Error
}

// resolveIDs はStore entityのVOからbusiness_type_idとcontract_plan_idを取得する。
func (r *storeRepository) resolveIDs(ctx context.Context, s *entity.Store) (btID uint, cpID uint, err error) {
	if err = r.db.WithContext(ctx).Raw(`SELECT id FROM business_types WHERE code = ? LIMIT 1`, s.BusinessType.GetCode()).Scan(&btID).Error; err != nil {
		return
	}
	err = r.db.WithContext(ctx).Raw(`SELECT id FROM contract_plans WHERE code = ? LIMIT 1`, s.ContractPlan.GetCode()).Scan(&cpID).Error
	return
}
