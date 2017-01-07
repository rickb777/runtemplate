package maps

import (
	"testing"
)

func TestExToSlice(t *testing.T) {
	a := NewEXIntIntMap1(1, 2)
	s := a.ToSlice()

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	if len(s) != 1 {
		t.Errorf("Expected 1 but got %d", len(s))
	}
}

func TestExRemove(t *testing.T) {
	a := NewEXIntIntMap1(3, 1)

	a.Put(1, 5)
	a.Put(2, 5)
	a.Remove(3)

	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	if !(a.ContainsKey(1) && a.ContainsKey(2)) {
		t.Errorf("%+v", a.m)
	}

	a.Remove(2)
	a.Remove(1)

	if a.Size() != 0 {
		t.Errorf("%+v", a.m)
	}
}

func TestExContainsKey(t *testing.T) {
	a := NewEXIntIntMap1(13, 1)

	a.Put(71, 13)

	if !a.ContainsKey(71) {
		t.Error("should contain 71")
	}

	a.Remove(71)

	if a.ContainsKey(71) {
		t.Error("should not contain 71")
	}

	a.Put(9, 5)

	if !(a.ContainsKey(9) && a.ContainsKey(13)) {
		t.Errorf("%+v", a.m)
	}
}

func TestExContainsAllKeys(t *testing.T) {
	a := NewEXIntIntMap1(8, 6)

	a.Put(1, 10)
	a.Put(2, 11)

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("%+v", a.m)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("%+v", a.m)
	}
}

func TestExClear(t *testing.T) {
	a := NewEXIntIntMap1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("%+v", a.m)
	}
}

func TestExCardinality(t *testing.T) {
	a := NewEXIntIntMap()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a.Put(1, 2)

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	a.Remove(1)

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a.Put(9, 5)

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
}

func TestExEquals(t *testing.T) {
	a := NewEXIntIntMap()
	b := NewEXIntIntMap()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}

	a.Put(10, 4)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, b.m)
	}

	b.Put(10, 4)

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}

	b.Put(8, 8)
	b.Put(3, 1)
	b.Put(47, 49)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, b.m)
	}

	a.Put(8, 8)
	a.Put(3, 1)
	a.Put(47, 49)

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}

	a.Put(47, 1)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, b.m)
	}
}

func TestExClone(t *testing.T) {
	a := NewEXIntIntMap()
	a.Put(1, 9)
	a.Put(2, 8)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a.m, b.m)
	}

	a.Put(3, 3)
	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, b.m)
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a.m, c.m)
	}
}

//func TestExSend(t *testing.T) {
//	a := NewEXIntIntMap(1, 2, 3, 4)
//
//	b := NewEXIntIntMap()
//	for val := range a.Send() {
//		b.Add(val)
//	}
//
//	if !a.Equals(b) {
//		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
//	}
//}
