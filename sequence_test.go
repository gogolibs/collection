package collection_test

import (
	"github.com/gogolibs/collection"
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	sliceAdapter := collection.FromItems(1, 2, 3)
	expected := []int{1, 2, 3}
	actual := collection.ToSlice[int](sliceAdapter)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actual, expected)
	}
}

func TestForIndexedEach(t *testing.T) {
	sliceAdapter := collection.FromItems(1, 2, 3)
	expectedIndexes := []int{0, 1, 2}
	expectedSlice := []int{1, 2, 3}
	actualIndexes := []int{}
	actualSlice := []int{}
	collection.ForIndexedEach[int](sliceAdapter, func(index int, item int) bool {
		actualIndexes = append(actualIndexes, index)
		actualSlice = append(actualSlice, item)
		return true
	})
	if !reflect.DeepEqual(expectedSlice, actualSlice) {
		t.Fatalf("actual slice %#v is not equal to expected %#v", actualSlice, expectedSlice)
	}
	if !reflect.DeepEqual(expectedIndexes, actualIndexes) {
		t.Fatalf("actual slice indexes %#v is not equal to expected %#v", actualIndexes, expectedIndexes)
	}
}
