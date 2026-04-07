package repository

import (
	"golang-clean-architecture/pkg/usecase/query"
	"gorm.io/gorm"
)

func buildQuery(db *gorm.DB, conditions []query.Condition) *gorm.DB {
	for _, c := range conditions {
		switch c.Kind {
		case query.KindWhere:
			db = db.Where(c.Column+" = ?", c.Value)
		case query.KindWhereIn:
			db = db.Where(c.Column+" IN ?", c.Value)
		case query.KindWhereBetween:
			db = db.Where(c.Column+" BETWEEN ? AND ?", c.From, c.To)
		case query.KindWhereNotIn:
			db = db.Where(c.Column+" NOT IN ?", c.Value)
		}
	}
	return db
}
