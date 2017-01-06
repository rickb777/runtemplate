package collectiontest1

import (
	"testing"
)

func makeList(ints ...int32) Int32Collection {
	return NewInt32List(ints...)
}

func makeSet(ints ...int32) Int32Collection {
	return NewInt32Set(ints...)
}

func TestNewCollection(t *testing.T) {
	testNewCollection(t, NewInt32Set())
	testNewCollection(t, NewInt32List())
}

func testNewCollection(t *testing.T, a Int32Collection) {
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

