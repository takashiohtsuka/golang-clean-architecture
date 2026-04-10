package valueobject

type Region struct {
	id   uint
	name string
}

func NewRegion(id uint, name string) Region {
	return Region{id: id, name: name}
}

func EmptyRegion() Region {
	return Region{}
}

func (r Region) GetID() uint {
	return r.id
}

func (r Region) GetName() string {
	return r.name
}

func (r Region) IsEmpty() bool {
	return r.id == 0
}
