package collection_test

import (
	"github.com/gogolibs/collection"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	isEmpty := collection.IsEmpty(collection.FromItems[int]())
	if !isEmpty {
		t.Fatalf("newly created slice must be empty")
	}
}
