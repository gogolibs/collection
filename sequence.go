package collection

import "github.com/gogolibs/iterator"

// Sequence is a collection items of which can be iterated over using an iterator.Iterator.
type Sequence[T any] interface {
	Iterator() iterator.Iterator[T]
}

type SizedSequence[T any] interface {
	Sized
	Sequence[T]
}

// ToSlice converts a SizedSequence to a slice.
func ToSlice[T any](sequence SizedSequence[T]) []T {
	slice := make([]T, 0, sequence.Size())
	iterator.ForEach(sequence.Iterator(), func(item T) bool {
		slice = append(slice, item)
		return true
	})
	return slice
}
