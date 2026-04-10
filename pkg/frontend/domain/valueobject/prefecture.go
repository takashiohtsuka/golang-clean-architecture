package valueobject

type Prefecture struct {
	id   uint
	name string
}

func NewPrefecture(id uint, name string) Prefecture {
	return Prefecture{id: id, name: name}
}

func EmptyPrefecture() Prefecture {
	return Prefecture{}
}

func (p Prefecture) GetID() uint {
	return p.id
}

func (p Prefecture) GetName() string {
	return p.name
}

func (p Prefecture) IsEmpty() bool {
	return p.id == 0
}
