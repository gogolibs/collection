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
