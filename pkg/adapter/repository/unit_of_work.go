package repository

import (
	"context"

	"golang-clean-architecture/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type unitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) outputport.UnitOfWork {
	return &unitOfWork{db}
}

func (u *unitOfWork) Do(ctx context.Context, fn func() error) error {
	return u.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn()
	})
}
