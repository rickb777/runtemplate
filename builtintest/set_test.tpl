// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Mutable:{{.Mutable}}

package {{.Package}}

import (
	"testing"
	"fmt"
)

func TestNew{{.UType}}Set(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3)

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

func TestNew{{.UType}}SetNoDuplicate(t *testing.T) {
	a := NewX1{{.UType}}Set(7, 5, 3, 7)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if !(a.Contains(7) && a.Contains(5) && a.Contains(3)) {
		t.Error("set should have a 7, 5, and 3 in it.")
	}
}

{{if .Mutable}}
func TestMutable{{.UType}}SetRemove(t *testing.T) {
	a := NewX1{{.UType}}Set(6, 3, 1)

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

{{end}}
func Test{{.UType}}SetContainsAll(t *testing.T) {
	a := NewX1{{.UType}}Set(8, 6, 7, 5, 3, 0, 9)

	if !a.ContainsAll(8, 6, 7, 5, 3, 0, 9) {
		t.Error("should contain phone number")
	}

	if a.ContainsAll(8, 6, 11, 5, 3, 0, 9) {
		t.Error("should not have all of these numbers")
	}
}

{{if .Mutable}}
func TestMutable{{.UType}}SetCardinality(t *testing.T) {
	a := NewX1{{.UType}}Set()

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
}

{{end}}
func Test{{.UType}}SetIsSubset(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 5, 7)
	b := NewX1{{.UType}}Set(3, 5, 7)
	c := NewX1{{.UType}}Set(3, 5, 7, 72)

	if !b.IsSubset(a) {
		t.Errorf("Expected '%+v' to be a subset of '%+v'", b, a)
	}

	if c.IsSubset(a) {
		t.Errorf("Expected '%+v' not to be a subset of '%+v'", c, a)
	}
}

func Test{{.UType}}SetIsSuperSet(t *testing.T) {
	a := NewX1{{.UType}}Set(9, 5, 2, 1, 11)
	b := NewX1{{.UType}}Set(5, 2, 11)
	c := NewX1{{.UType}}Set(5, 2, 11, 42)

	if !a.IsSuperset(b) {
		t.Errorf("Expected '%+v' to be a superset of '%+v'", a, b)
	}

	if a.IsSuperset(c) {
		t.Errorf("Expected '%+v' not to be a superset of '%+v'", a, b)
	}
}

func Test{{.UType}}SetUnion(t *testing.T) {
	a := NewX1{{.UType}}Set()

	b := NewX1{{.UType}}Set(1, 2, 3, 4, 5)

	c := a.Union(b)

	if c.Size() != 5 {
		t.Errorf("Expected 5 but got %d", c.Size())
	}

	d := NewX1{{.UType}}Set(10, 14, 0)

	e := c.Union(d)
	if e.Size() != 8 {
		t.Errorf("Expected 8 but got %d", e.Size())
	}

	f := NewX1{{.UType}}Set(14, 3)

	g := f.Union(e)
	if g.Size() != 8 {
		t.Errorf("Expected 8 but got %d", g.Size())
	}
}

func Test{{.UType}}SetIntersection(t *testing.T) {
	a1 := NewX1{{.UType}}Set(1, 3, 5, 7)
	a2 := NewX1{{.UType}}Set(1, 3, 5, 7, 10)

	b1 := NewX1{{.UType}}Set(2, 4, 6)
	b2 := NewX1{{.UType}}Set(2, 4, 6, 10)

	c1 := a1.Intersect(b1)
	c2 := b1.Intersect(a1)

	if c1.NonEmpty() || c2.NonEmpty() {
		t.Errorf("Expected 0 but got %d and %d", c1.Size(), c2.Size())
	}

	d1 := a2.Intersect(b2)
	d2 := b2.Intersect(a2)

	if !(d1.Size() == 1 && d1.Contains(10)) {
		t.Errorf("d1 should have a size of 1 and contain 10: %+v", d1)
	}

	if !(d2.Size() == 1 && d2.Contains(10)) {
		t.Errorf("d2 should have a size of 1 and contain 10: %+v", d2)
	}
}

func TestMutable{{.UType}}SetDifference(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3)

	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)

	if !(c.Size() == 1 && c.Contains(2)) {
		t.Error("the difference of set a to b is the set of 1 item: 2")
	}
}

func TestMutable{{.UType}}SetSymmetricDifference(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 45)

	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.SymmetricDifference(b)

	if !(c.Size() == 6 && c.Contains(2) && c.Contains(45) && c.Contains(4) && c.Contains(5) && c.Contains(6) && c.Contains(99)) {
		t.Error("the symmetric difference of set a to b is the set of 6 items: 2, 45, 4, 5, 6, 99")
	}
}

func TestMutable{{.UType}}SetEqual(t *testing.T) {
	a := NewX1{{.UType}}Set()
	b := NewX1{{.UType}}Set()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}

	c := NewX1{{.UType}}Set(1, 3, 5, 6, 8)
	d := NewX1{{.UType}}Set(1, 3, 5, 6, 9)

	if c.Equals(d) {
		t.Errorf("Expected '%+v' not to equal '%+v'", c, d)
	}
}

func TestMutable{{.UType}}SetSend(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := BuildX1{{.UType}}SetFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
	}
}

func TestMutable{{.UType}}SetFilter(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}Set(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}
}

func TestMutable{{.UType}}SetPartition(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}Set(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	if !c.Equals(NewX1{{.UType}}Set(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c)
	}
}

func TestMutable{{.UType}}SetStringMap(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.StringMap()

	for _, c := range a.ToSlice() {
		s := fmt.Sprintf("%d", c)
		if _, ok := b[s]; !ok {
			t.Errorf("Expected '%s' but got '%+v'", s, b)
		}
	}
}

{{if .Mutable}}
func TestMutable{{.UType}}SetClear(t *testing.T) {
	a := NewX1{{.UType}}Set(2, 5, 9, 10)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
}

func TestMutable{{.UType}}SetClone(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2)

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
{{end}}