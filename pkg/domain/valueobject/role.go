package valueobject

type Role struct {
	ValueObject[string]
}

func NewRole(name string) Role {
	return Role{NewValueObject(name)}
}

func EmptyRole() Role {
	return Role{}
}

func (r Role) GetName() string {
	return r.Get()
}

func (r Role) IsEmpty() bool {
	return r.Get() == ""
}
