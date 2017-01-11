package threadsafe

import (
	"testing"
)

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

func TestToSlice(t *testing.T) {
	testToSlice(t, NewXInt32Set(1, 2, 3))
	testToSlice(t, NewXInt32List(1, 2, 3))
}

func testToSlice(t *testing.T, a XInt32Collection) {
	s := a.ToSlice()

	if len(s) != 3 {
		t.Errorf("Expected 3 but got %d", len(s))
	}
}

