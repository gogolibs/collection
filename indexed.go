package collection

// Indexed is a collection which items can be accessed by index via Get(int).
type Indexed[T any] interface {
	Get(index int) T
}

type SizedIndexed[T any] interface {
	Sized
	Indexed[T]
}

// IndexedMutable is an Indexed with Set method
type IndexedMutable[T any] interface {
	Indexed[T]
	Set(index int, item T)
}

type SizedIndexedMutable[T any] interface {
	Sized
	IndexedMutable[T]
}

// ToIterator creates an *IteratorByIndex (implementation of iterator.Iterator)
// from an SizedIndexed collection.
// IteratorByIndex iterates over the collection by repeatedly calling its Get method
// which may or may not be the optimal iteration approach
// (e.g. it's optimal for slices but not optimal for linked lists).
// Consider using a custom iterator.Iterator implementation for the latter.
func ToIterator[T any](indexed SizedIndexed[T]) *IteratorByIndex[T] {
	return &IteratorByIndex[T]{
		indexed: indexed,
		index:   0,
	}
}

// SwapIndexes swaps two indexes of the IndexedMutable collection.
func SwapIndexes[T any](indexed IndexedMutable[T], index1, index2 int) {
	tmp := indexed.Get(index1)
	indexed.Set(index1, indexed.Get(index2))
	indexed.Set(index2, tmp)
}
