package repository

import (
	"strings"

	"golang-clean-architecture/pkg/usecase/query"
)

func buildWhereClause(conditions []query.Condition) (string, []any) {
	return buildWhereClauseWithPrefix(conditions, "")
}

func buildWhereClauseWithPrefix(conditions []query.Condition, tablePrefix string) (string, []any) {
	var clauses []string
	var args []any

	col := func(name string) string {
		if tablePrefix == "" {
			return name
		}
		return tablePrefix + "." + name
	}

	for _, c := range conditions {
		switch c.Kind {
		case query.KindWhere:
			clauses = append(clauses, col(c.Column)+" = ?")
			args = append(args, c.Value)
		case query.KindWhereIn:
			clauses = append(clauses, col(c.Column)+" IN ?")
			args = append(args, c.Value)
		case query.KindWhereBetween:
			clauses = append(clauses, col(c.Column)+" BETWEEN ? AND ?")
			args = append(args, c.From, c.To)
		case query.KindWhereNotIn:
			clauses = append(clauses, col(c.Column)+" NOT IN ?")
			args = append(args, c.Value)
		}
	}

	if len(clauses) == 0 {
		return "", nil
	}
	return " AND " + strings.Join(clauses, " AND "), args
}
