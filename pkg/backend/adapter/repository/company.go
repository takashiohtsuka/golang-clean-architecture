package repository

import (
	"context"

	companyMapper "golang-clean-architecture/pkg/backend/adapter/mapper/company"
	"golang-clean-architecture/pkg/backend/domain/entity"
	"golang-clean-architecture/pkg/backend/usecase/outputport"
	"golang-clean-architecture/pkg/domain/collection"
	"golang-clean-architecture/pkg/helper"
	"golang-clean-architecture/pkg/usecase/query"

	"gorm.io/gorm"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) outputport.CompanyRepository {
	return &companyRepository{db: db}
}

const companySelectSQL = "SELECT id, name, `rank`, is_active, created_at, updated_at, deleted_at FROM companies WHERE deleted_at IS NULL"

func toCompanyEntity(row map[string]any) *entity.Company {
	return &entity.Company{
		ID:        helper.ToUint(row["id"]),
		Name:      helper.ToString(row["name"]),
		Rank:      helper.ToStringPtr(row["rank"]),
		IsActive:  helper.ToBool(row["is_active"]),
		CreatedAt: helper.ToTimePtr(row["created_at"]),
		UpdatedAt: helper.ToTimePtr(row["updated_at"]),
		DeletedAt: helper.ToTimePtr(row["deleted_at"]),
	}
}

func (r *companyRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.CompanyEntity], error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(companySelectSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.CompanyEntity](nil), err
	}

	items := make([]entity.CompanyEntity, len(rows))
	for i, row := range rows {
		items[i] = toCompanyEntity(row)
	}
	return collection.NewCollection(items), nil
}

func (r *companyRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.CompanyEntity, error) {
	where, args := buildWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(companySelectSQL+where+" LIMIT 1", args...).Scan(&rows).Error; err != nil {
		return &entity.NilCompany{}, err
	}
	if len(rows) == 0 {
		return &entity.NilCompany{}, nil
	}
	return toCompanyEntity(rows[0]), nil
}

func (r *companyRepository) Create(ctx context.Context, c *entity.Company) error {
	return r.db.WithContext(ctx).Create(companyMapper.ToOrmModel(c)).Error
}

func (r *companyRepository) Update(ctx context.Context, c *entity.Company) error {
	return r.db.WithContext(ctx).Save(companyMapper.ToOrmModel(c)).Error
}
