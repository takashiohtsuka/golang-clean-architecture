package outputport

type DBRepository interface {
	Transaction(func(interface{}) (interface{}, error)) (interface{}, error)
}
