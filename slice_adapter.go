package collection

import "github.com/gogolibs/iterator"

func FromItems[T any](items ...T) *SliceAdapter[T] {
	return NewSliceAdapter(items)
}

func NewSliceAdapter[T any](slice []T) *SliceAdapter[T] {
	return &SliceAdapter[T]{
		Slice: slice,
	}
}

type SliceAdapter[T any] struct {
	Slice []T
}

func (a *SliceAdapter[T]) Size() int {
	return len(a.Slice)
}

func (a *SliceAdapter[T]) Get(index int) T {
	return a.Slice[index]
}

func (a *SliceAdapter[T]) Set(index int, item T) {
	a.Slice[index] = item
}

func (a *SliceAdapter[T]) Iterator() iterator.Iterator[T] {
	return iterator.FromSlice(a.Slice)
}
