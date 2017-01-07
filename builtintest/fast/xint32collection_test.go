package fast

import (
	"testing"
)

func makeList(ints ...int32) XInt32Collection {
	return NewXInt32List(ints...)
}

func makeSet(ints ...int32) XInt32Collection {
	return NewXInt32Set(ints...)
}

func TestNewCollection(t *testing.T) {
	testNewCollection(t, NewXInt32Set())
	testNewCollection(t, NewXInt32List())
}

func testNewCollection(t *testing.T, a XInt32Collection) {
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

