package collection

import "github.com/gogolibs/iterator"

// Sequence is a collection items of which can be iterated over using an iterator.Iterator.
type Sequence[T any] interface {
	Iterator() iterator.Iterator[T]
}

// ForEach is an equivalent for iterator.ForEach for a Sequence.
func ForEach[T any](sequence Sequence[T], f func(T) bool) bool {
	return iterator.ForEach(sequence.Iterator(), f)
}

// ForIndexedEach is ForEach with added index.
func ForIndexedEach[T any](sequence Sequence[T], f func(int, T) bool) bool {
	index := 0
	return ForEach[T](sequence, func(item T) bool {
		proceed := f(index, item)
		index++
		return proceed
	})
}

type SizedSequence[T any] interface {
	Sized
	Sequence[T]
}

// ToSlice converts a SizedSequence to a slice.
func ToSlice[T any](sequence SizedSequence[T]) []T {
	slice := make([]T, 0, sequence.Size())
	ForEach[T](sequence, func(item T) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}
