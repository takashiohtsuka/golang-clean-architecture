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

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) outputport.CompanyRepository {
	return &companyRepository{db: db}
}

type companyRow struct {
	ID        uint       `gorm:"column:id"`
	Name      string     `gorm:"column:name"`
	Rank      *string    `gorm:"column:rank"`
	IsActive  bool       `gorm:"column:is_active"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (r *companyRow) toEntity() *entity.Company {
	return &entity.Company{
		ID:        r.ID,
		Name:      r.Name,
		Rank:      r.Rank,
		IsActive:  r.IsActive,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		DeletedAt: r.DeletedAt,
	}
}

func (r *companyRepository) FindAll(conditions []query.Condition) (collection.Collection[entity.CompanyEntity], error) {
	where, args := buildWhereClause(conditions)
	sql := "SELECT id, name, `rank`, is_active, created_at, updated_at, deleted_at FROM companies WHERE deleted_at IS NULL" + where

	var rows []companyRow
	if err := r.db.Raw(sql, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.CompanyEntity](nil), err
	}

	items := make([]entity.CompanyEntity, len(rows))
	for i, row := range rows {
		items[i] = row.toEntity()
	}
	return collection.NewCollection(items), nil
}

func (r *companyRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.CompanyEntity, error) {
	where, args := buildWhereClause(conditions)
	sql := "SELECT id, name, `rank`, is_active, created_at, updated_at, deleted_at FROM companies WHERE deleted_at IS NULL" + where + " LIMIT 1"

	var row companyRow
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&row).Error; err != nil {
		return &entity.NilCompany{}, err
	}
	if row.ID == 0 {
		return &entity.NilCompany{}, nil
	}
	return row.toEntity(), nil
}

func (r *companyRepository) Create(c *entity.Company) error {
	sql := "INSERT INTO companies (name, `rank`, is_active, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())"
	return r.db.Exec(sql, c.Name, c.Rank, c.IsActive).Error
}

func (r *companyRepository) Update(c *entity.Company) error {
	sql := "UPDATE companies SET name = ?, `rank` = ?, is_active = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL"
	return r.db.Exec(sql, c.Name, c.Rank, c.IsActive, c.ID).Error
}
