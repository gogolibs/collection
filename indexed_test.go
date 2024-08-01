package collection_test

import (
	"github.com/gogolibs/collection"
	"github.com/gogolibs/iterator"
	"reflect"
	"testing"
)

func TestToIterator(t *testing.T) {
	slice := collection.FromItems(1, 2, 3)
	expected := []int{1, 2, 3}
	actual := iterator.ToSlice[int](collection.ToIterator[int](slice), 3)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actual, expected)
	}
}

func TestSwapIndexes(t *testing.T) {
	slice := collection.FromItems(1, 2, 3)
	collection.SwapIndexes[int](slice, 0, 2)
	expected := []int{3, 2, 1}
	actual := slice.Slice
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actual, expected)
	}
}
