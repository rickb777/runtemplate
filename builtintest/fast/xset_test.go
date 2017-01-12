package fast

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	a := NewXInt32Set(1, 2, 3)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if !a.IsSet() {
		t.Error("Expected a set")
	}

	if a.IsSequence() {
		t.Error("Expected not a sequence")
	}
}

func TestNewSetNoDuplicate(t *testing.T) {
	a := NewXInt32Set(7, 5, 3, 7)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("set should have a 7, 5, and 3 in it.")
	}
}

func TestSetRemove(t *testing.T) {
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
	if a.Cardinality() != 1 {
		t.Errorf("Expected 1 but got %d", a.Cardinality())
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

	b := NewXInt32Set(3, 5, 7)

	if !b.IsSubset(a) {
		t.Errorf("Expected '%+v' to be a subset of '%+v'", b, a)
	}

	b.Add(72)

	if b.IsSubset(a) {
		t.Errorf("Expected '%+v' not to be a subset of '%+v'", b, a)
	}
}

func TestIsSuperSet(t *testing.T) {
	a := NewXInt32Set(9, 5, 2, 1, 11)

	b := NewXInt32Set(5, 2, 11)

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

	b := NewXInt32Set(1, 2, 3, 4, 5)

	c := a.Union(b)

	if c.Size() != 5 {
		t.Errorf("Expected 5 but got %d", c.Size())
	}

	d := NewXInt32Set(10, 14, 0)

	e := c.Union(d)
	if e.Size() != 8 {
		t.Errorf("Expected 8 but got %d", e.Size())
	}

	f := NewXInt32Set(14, 3)

	g := f.Union(e)
	if g.Size() != 8 {
		t.Errorf("Expected 8 but got %d", g.Size())
	}
}

func TestSetIntersection(t *testing.T) {
	a := NewXInt32Set(1, 3, 5, 7)

	b := NewXInt32Set(2, 4, 6)

	c1 := a.Intersect(b)
	c2 := b.Intersect(a)

	if c1.NonEmpty() || c2.NonEmpty() {
		t.Errorf("Expected 0 but got %d and %d", c1.Size(), c2.Size())
	}

	a.Add(10)
	b.Add(10)

	d1 := a.Intersect(b)
	d2 := b.Intersect(a)

	if !(d1.Size() == 1 && d1.Contains(10)) {
		t.Errorf("d1 should have a size of 1 and contain 10: %+v", d1)
	}

	if !(d2.Size() == 1 && d2.Contains(10)) {
		t.Errorf("d2 should have a size of 1 and contain 10: %+v", d2)
	}
}

func TestSetDifference(t *testing.T) {
	a := NewXInt32Set(1, 2, 3)

	b := NewXInt32Set(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)

	if !(c.Size() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestSetSymmetricDifference(t *testing.T) {
	a := NewXInt32Set(1, 2, 3, 45)

	b := NewXInt32Set(1, 3, 4, 5, 6, 99)

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
	a := NewXInt32Set(1, 2)

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

func TestSetSend(t *testing.T) {
	a := NewXInt32Set(1, 2, 3, 4)

	b := BuildXInt32SetFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}
