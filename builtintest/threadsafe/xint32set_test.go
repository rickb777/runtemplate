package threadsafe

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	a := NewXInt32Set(1, 2, 3)
	s := a.ToSlice()

	if len(s) != 3 {
		t.Errorf("Expected 3 but got %d", len(s))
	}
}

func TestAddSet(t *testing.T) {
	a := NewXInt32Set(1, 2, 3)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}
}

func TestAddSetNoDuplicate(t *testing.T) {
	a := NewXInt32Set(7, 5, 3, 7)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("set should have a 7, 5, and 3 in it.")
	}
}

func TestRemoveSet(t *testing.T) {
	a := NewXInt32Set(6, 3, 1)

	a.Remove(3)

	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	if !(a.Contains(6) && a.Contains(1)) {
		t.Error("should have only items 6 and 1 in the set")
	}

	a.Remove(6)
	a.Remove(1)

	if a.Size() != 0 {
		t.Error("should be an empty set after removing 6 and 1")
	}
}

func TestContainsSet(t *testing.T) {
	a := NewXInt32Set()

	a.Add(71)

	if !a.Contains(71) {
		t.Error("should contain 71")
	}

	a.Remove(71)

	if a.Contains(71) {
		t.Error("should not contain 71")
	}

	a.Add(13)
	a.Add(7)
	a.Add(1)

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("should contain 13, 7, 1")
	}
}

func TestContainsAllSet(t *testing.T) {
	a := NewXInt32Set(8, 6, 7, 5, 3, 0, 9)

	if !a.ContainsAll(8, 6, 7, 5, 3, 0, 9) {
		t.Error("should contain phone number")
	}

	if a.ContainsAll(8, 6, 11, 5, 3, 0, 9) {
		t.Error("should not have all of these numbers")
	}
}

func TestClearSet(t *testing.T) {
	a := NewXInt32Set(2, 5, 9, 10)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
}

func TestCardinality(t *testing.T) {
	a := NewXInt32Set()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a.Add(1)

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	a.Remove(1)

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a.Add(9)

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
}

func TestIsSubset(t *testing.T) {
	a := NewXInt32Set(1, 2, 3, 5, 7)

	b := NewXInt32Set()
	b.Add(3)
	b.Add(5)
	b.Add(7)

	if !b.IsSubset(a) {
		t.Errorf("Expected '%+v' to be a subset of '%+v'", b, a)
	}

	b.Add(72)

	if b.IsSubset(a) {
		t.Errorf("Expected '%+v' not to be a subset of '%+v'", b, a)
	}
}

func TestIsSuperSet(t *testing.T) {
	a := NewXInt32Set()
	a.Add(9)
	a.Add(5)
	a.Add(2)
	a.Add(1)
	a.Add(11)

	b := NewXInt32Set()
	b.Add(5)
	b.Add(2)
	b.Add(11)

	if !a.IsSuperset(b) {
		t.Errorf("Expected '%+v' to be a superset of '%+v'", a, b)
	}

	b.Add(42)

	if a.IsSuperset(b) {
		t.Errorf("Expected '%+v' not to be a superset of '%+v'", a, b)
	}
}

func TestSetUnion(t *testing.T) {
	a := NewXInt32Set()

	b := NewXInt32Set()
	b.Add(1)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	c := a.Union(b)

	if c.Size() != 5 {
		t.Errorf("Expected 5 but got %d", c.Size())
	}

	d := NewXInt32Set()
	d.Add(10)
	d.Add(14)
	d.Add(0)

	e := c.Union(d)
	if e.Size() != 8 {
		t.Errorf("Expected 8 but got %d", e.Size())
	}

	f := NewXInt32Set()
	f.Add(14)
	f.Add(3)

	g := f.Union(e)
	if g.Size() != 8 {
		t.Errorf("Expected 8 but got %d", g.Size())
	}
}

func TestSetIntersection(t *testing.T) {
	a := NewXInt32Set()
	a.Add(1)
	a.Add(3)
	a.Add(5)

	b := NewXInt32Set()
	a.Add(2)
	a.Add(4)
	a.Add(6)

	c := a.Intersect(b)

	if c.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a.Add(10)
	b.Add(10)

	d := a.Intersect(b)

	if !(d.Size() == 1 && d.Contains(10)) {
		t.Error("set d should have a size of 1 and contain the item 10")
	}
}

func TestSetDifference(t *testing.T) {
	a := NewXInt32Set()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewXInt32Set()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.Difference(b)

	if !(c.Size() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	a := NewXInt32Set()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	a.Add(45)

	b := NewXInt32Set()
	b.Add(1)
	b.Add(3)
	b.Add(4)
	b.Add(5)
	b.Add(6)
	b.Add(99)

	c := a.SymmetricDifference(b)

	if !(c.Size() == 6 && c.Contains(2) && c.Contains(45) && c.Contains(4) && c.Contains(5) && c.Contains(6) && c.Contains(99)) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

func TestSetEqual(t *testing.T) {
	a := NewXInt32Set()
	b := NewXInt32Set()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	a.Add(10)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, b)
	}

	b.Add(10)

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	b.Add(8)
	b.Add(3)
	b.Add(47)

	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, b)
	}

	a.Add(8)
	a.Add(3)
	a.Add(47)

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}

func TestSetClone(t *testing.T) {
	a := NewXInt32Set()
	a.Add(1)
	a.Add(2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	a.Add(3)
	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, b)
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, c)
	}
}

func TestSend(t *testing.T) {
	a := NewXInt32Set(1, 2, 3, 4)

	b := NewXInt32Set()
	for val := range a.Send() {
		b.Add(val)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}
