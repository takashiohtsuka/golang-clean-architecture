package query

type Kind int

const (
	KindWhere Kind = iota
	KindWhereIn
	KindWhereBetween
	KindWhereNotIn
)

type Condition struct {
	Kind   Kind
	Column string
	Value  any // Where / WhereIn / WhereNotIn
	From   any // WhereBetween
	To     any // WhereBetween
}

func Where(column string, value any) Condition {
	return Condition{Kind: KindWhere, Column: column, Value: value}
}

func WhereIn(column string, values any) Condition {
	return Condition{Kind: KindWhereIn, Column: column, Value: values}
}

func WhereBetween(column string, from, to any) Condition {
	return Condition{Kind: KindWhereBetween, Column: column, From: from, To: to}
}

func WhereNotIn(column string, values any) Condition {
	return Condition{Kind: KindWhereNotIn, Column: column, Value: values}
}
