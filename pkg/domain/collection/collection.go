package collection

import "encoding/json"

type Collection[T any] struct {
	items []T
}

// MarshalJSON はjson.Marshalerインターフェースの実装。
// itemsはunexportedフィールドのため、そのままではJSONに出力されない。
// このメソッドを実装することで、json.Marshal()実行時に自動で呼ばれ、
// itemsの中身をJSON配列として出力する。
// 構造体のフィールドとして埋め込んだ場合も、各フィールドごとに自動で呼ばれる。
func (c Collection[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.items)
}

func NewCollection[T any](items []T) Collection[T] {
	return Collection[T]{items: items}
}

func (c Collection[T]) Filter(fn func(T) bool) Collection[T] {
	var result []T
	for _, item := range c.items {
		if fn(item) {
			result = append(result, item)
		}
	}
	return Collection[T]{items: result}
}

func (c Collection[T]) First() T {
	if len(c.items) == 0 {
		var zero T
		return zero
	}
	return c.items[0]
}

func (c Collection[T]) All() []T {
	return c.items
}

func (c Collection[T]) TotalCount() int {
	return len(c.items)
}

func (c Collection[T]) RemoveByFunc(fn func(T) bool) Collection[T] {
	var result []T
	for _, item := range c.items {
		if !fn(item) {
			result = append(result, item)
		}
	}
	return Collection[T]{items: result}
}

func (c Collection[T]) RemoveByIndex(index int) Collection[T] {
	if index < 0 || index >= len(c.items) {
		return c
	}
	result := make([]T, 0, len(c.items)-1)
	result = append(result, c.items[:index]...)
	result = append(result, c.items[index+1:]...)
	return Collection[T]{items: result}
}
