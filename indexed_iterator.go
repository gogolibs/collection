package collection

// IteratorByIndex is returned by ToIterator function. See ToIterator docs for more information.
type IteratorByIndex[T any] struct {
	indexed SizedIndexed[T]
	index   int
}

func (i *IteratorByIndex[T]) HasNext() bool {
	return i.index < i.indexed.Size()
}

func (i *IteratorByIndex[T]) Next() T {
	item := i.indexed.Get(i.index)
	i.index++
	return item
}
