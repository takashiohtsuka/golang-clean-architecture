package valueobject

type District struct {
	id   uint
	name string
}

func NewDistrict(id uint, name string) District {
	return District{id: id, name: name}
}

func EmptyDistrict() District {
	return District{}
}

func (d District) GetID() uint {
	return d.id
}

func (d District) GetName() string {
	return d.name
}

func (d District) IsEmpty() bool {
	return d.id == 0
}
