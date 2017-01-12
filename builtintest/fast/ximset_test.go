package fast

import (
	"testing"
)

func imSet(v ...int) XAppleSet {
	a := make([]Apple, len(v), len(v))
	for i, x := range v {
		a[i] = Apple{x}
	}
	return NewXAppleSet(a...)
}

func imSetOf(from, to int) XAppleSet {
	n := 1 + to - from
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i + from
	}
	return imSet(a...)
}

func TestNewImmutableSet(t *testing.T) {
	a := imSet(1, 2, 3)

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

func TestNewImmutableSetNoDuplicate(t *testing.T) {
	a := imSet(7, 5, 3, 7)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if !(a.Contains(Apple{7}) && a.Contains(Apple{5}) && a.Contains(Apple{3})) {
		t.Error("set should have a 7, 5, and 3 in it.")
	}
}

func TestImmutableSetRemove(t *testing.T) {
	a := imSet(6, 3, 1)

	b := a.Remove(Apple{3})

	if b.Size() != 2 {
		t.Errorf("Expected 2 but got %d", b.Size())
	}

	if !(b.Contains(Apple{6}) && b.Contains(Apple{1})) {
		t.Error("should have only items 6 and 1 in the set")
	}

	c := b.Remove(Apple{6}).Remove(Apple{1})

	if c.Size() != 0 {
		t.Error("should be an empty set after removing 6 and 1")
	}
}

func TestImmutableSetContainsAll(t *testing.T) {
	a := imSet(8, 6, 7, 5, 3, 0, 9)

	if !a.ContainsAll(Apple{8}, Apple{6}, Apple{7}, Apple{5}, Apple{3}, Apple{0}, Apple{9}) {
		t.Error("should contain phone number")
	}

	if a.ContainsAll(Apple{8}, Apple{6}, Apple{11}, Apple{5}, Apple{3}, Apple{0}, Apple{9}) {
		t.Error("should not have all of these numbers")
	}
}

func TestImmutableSetCardinality(t *testing.T) {
	a := imSet()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a = a.Add(Apple{1})

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}
	if a.Cardinality() != 1 {
		t.Errorf("Expected 1 but got %d", a.Cardinality())
	}

	a = a.Remove(Apple{1})

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

	a = a.Add(Apple{9})

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}
}

func TestImmutableSetIsSubset(t *testing.T) {
	a := imSet(1, 2, 3, 5, 7)

	b := imSet(3, 5, 7)

	if !b.IsSubset(a) {
		t.Errorf("Expected '%+v' to be a subset of '%+v'", b, a)
	}

	b = b.Add(Apple{72})

	if b.IsSubset(a) {
		t.Errorf("Expected '%+v' not to be a subset of '%+v'", b, a)
	}
}

func TestImmutableSetIsSuperSet(t *testing.T) {
	a := imSet(9, 5, 2, 1, 11)

	b := imSet(5, 2, 11)

	if !a.IsSuperset(b) {
		t.Errorf("Expected '%+v' to be a superset of '%+v'", a, b)
	}

	b = b.Add(Apple{42})

	if a.IsSuperset(b) {
		t.Errorf("Expected '%+v' not to be a superset of '%+v'", a, b)
	}
}

func TestImmutableSetUnion(t *testing.T) {
	a := imSet()

	b := imSet(1, 2, 3, 4, 5)

	c := a.Union(b)

	if c.Size() != 5 {
		t.Errorf("Expected 5 but got %d", c.Size())
	}

	d := imSet(10, 14, 0)

	e := c.Union(d)
	if e.Size() != 8 {
		t.Errorf("Expected 8 but got %d", e.Size())
	}

	f := imSet(14, 3)

	g := f.Union(e)
	if g.Size() != 8 {
		t.Errorf("Expected 8 but got %d", g.Size())
	}
}

func TestImmutableSetIntersection(t *testing.T) {
	a := imSet(1, 3, 5, 7)

	b := imSet(2, 4, 6)

	c1 := a.Intersect(b)
	c2 := b.Intersect(a)

	if c1.NonEmpty() || c2.NonEmpty() {
		t.Errorf("Expected 0 but got %d and %d", c1.Size(), c2.Size())
	}

	a = a.Add(Apple{10})
	b = b.Add(Apple{10})

	d1 := a.Intersect(b)
	d2 := b.Intersect(a)

	if !(d1.Size() == 1 && d1.Contains(Apple{10})) {
		t.Errorf("d1 should have a size of 1 and contain 10: %+v", d1)
	}

	if !(d2.Size() == 1 && d2.Contains(Apple{10})) {
		t.Errorf("d2 should have a size of 1 and contain 10: %+v", d2)
	}
}

func TestImmutableSetDifference(t *testing.T) {
	a := imSet(1, 2, 3)

	b := imSet(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)

	if !(c.Size() == 1 && c.Contains(Apple{2})) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestImmutableSetSymmetricDifference(t *testing.T) {
	a := imSet(1, 2, 3, 45)

	b := imSet(1, 3, 4, 5, 6, 99)

	c := a.SymmetricDifference(b)

	if !(c.Size() == 6 && c.Contains(Apple{2}) && c.Contains(Apple{45}) && c.Contains(Apple{4}) && c.Contains(Apple{5}) && c.Contains(Apple{6}) && c.Contains(Apple{99})) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

func TestImmutableSetEqual(t *testing.T) {
	a := imSet()
	b := imSet()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	c := imSet(1, 3, 5, 6, 8)
	d := imSet(1, 3, 5, 6, 9)

	if c.Equals(d) {
		t.Errorf("Expected '%+v' not to equal '%+v'", c, d)
	}

	c = c.Add(Apple{9})
	d = d.Add(Apple{8})

	if !c.Equals(d) {
		t.Errorf("Expected '%+v' to equal '%+v'", c, d)
	}
}

func TestImmutableSetClone(t *testing.T) {
	a := imSet(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	a = a.Add(Apple{3})
	if a.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, b)
	}

	c := a.Clone().Remove(Apple{1})

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a, c)
	}
}

func TestImmutableSetSend(t *testing.T) {
	a := imSet(1, 2, 3, 4)

	b := BuildXAppleSetFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}

func TestImmutableSetFilter(t *testing.T) {
	a := imSet(1, 2, 3, 4)

	b := a.Filter(func(v Apple) bool {
		return v.N > 2
	})

	if !b.Equals(imSet(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}
}

func TestImmutableSetPartition(t *testing.T) {
	a := imSet(1, 2, 3, 4)

	b, c := a.Partition(func(v Apple) bool {
		return v.N > 2
	})

	if !b.Equals(imSet(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	if !c.Equals(imSet(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c)
	}
}

