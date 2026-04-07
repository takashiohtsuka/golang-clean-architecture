package valueobject

import "encoding/json"

type ValueObject[T any] struct {
	value T
}

func NewValueObject[T any](v T) ValueObject[T] {
	return ValueObject[T]{value: v}
}

func (vo ValueObject[T]) Get() T {
	return vo.value
}

// MarshalJSON はjson.Marshalerインターフェースの実装。
// valueはunexportedフィールドのため、そのままではJSONに出力されない。
// ValueObjectを継承する全ての型でMarshalJSONの個別実装が不要になる。
func (vo ValueObject[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(vo.value)
}
