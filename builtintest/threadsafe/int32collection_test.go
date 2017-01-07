package threadsafe

import (
	"testing"
)

func makeList(ints ...int32) XInt32Collection {
	return NewXInt32List(ints...)
}

func TestNewXCollection(t *testing.T) {
	testNewXCollection(t, NewXInt32Set())
	testNewXCollection(t, NewXInt32List())
}

func testNewXCollection(t *testing.T, a XInt32Collection) {
	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
	if !a.IsEmpty() {
		t.Error("Expected empty")
	}
	if a.NonEmpty() {
		t.Error("Expected empty")
	}
}

