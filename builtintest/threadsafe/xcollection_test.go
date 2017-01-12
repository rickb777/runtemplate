package threadsafe

import (
	"testing"
)

func TestNewXCollection(t *testing.T) {
	testNewXCollection1(t, NewXInt32Set())
	testNewXCollection1(t, NewXInt32List())
	testNewXCollection2(t, NewXAppleList())
	testNewXCollection2(t, NewXAppleSet())
}

func testNewXCollection1(t *testing.T, a XInt32Collection) {
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

func testNewXCollection2(t *testing.T, a XAppleCollection) {
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

func TestExists(t *testing.T) {
	testExists(t, NewXInt32Set(1, 2, 3))
	testExists(t, NewXInt32List(1, 2, 3))
}

func testExists(t *testing.T, a XInt32Collection) {
	has2 := a.Exists(func(v int32) bool {
		return v > 2
	})

	if !has2 {
		t.Errorf("Expected exists for %+v", a)
	}

	has5 := a.Exists(func(v int32) bool {
		return v > 5
	})

	if has5 {
		t.Errorf("Expected not exists for %+v", a)
	}
}

func TestForall(t *testing.T) {
	testForall(t, NewXInt32Set(1, 2, 3))
	testForall(t, NewXInt32List(1, 2, 3))
}

func testForall(t *testing.T, a XInt32Collection) {
	has1 := a.Forall(func(v int32) bool {
		return v >= 1
	})

	if !has1 {
		t.Errorf("Expected forall for %+v", a)
	}

	has2 := a.Forall(func(v int32) bool {
		return v >= 2
	})

	if has2 {
		t.Errorf("Expected not forall for %+v", a)
	}
}

func TestCountBy(t *testing.T) {
	testCountBy(t, NewXInt32Set(1, 2, 3))
	testCountBy(t, NewXInt32List(1, 2, 3))
}

func testCountBy(t *testing.T, a XInt32Collection) {
	n := a.CountBy(func(v int32) bool {
		return v >= 2
	})

	if n != 2 {
		t.Errorf("Expected 2 but got %d", n)
	}
}

func TestForeach(t *testing.T) {
	testForeach(t, NewXInt32Set(1, 2, 3))
	testForeach(t, NewXInt32List(1, 2, 3))
}

func testForeach(t *testing.T, a XInt32Collection) {
	sum1 := int32(0)
	a.Foreach(func(v int32) {
		sum1 += v
	})

	if sum1 != 6 {
		t.Errorf("Expected 6 but got %d", sum1)
	}
}
