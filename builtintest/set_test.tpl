// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
{{- if .GobEncode}}
    "bytes"
    "encoding/gob"
{{- end}}
    "encoding/json"
	"fmt"
	"testing"
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

func TestConvert{{.UType}}Set(t *testing.T) {
	a, ok := ConvertX1{{.UType}}Set(1, 5.1, uint8(2), 7, 3)

	if !ok {
		t.Errorf("Not ok")
	}

	if !a.Equals(NewX1{{.UType}}Set(1, 5, 2, 7, 3)) {
		t.Errorf("Expected 1,5,2,7,3 but got %v", a)
	}

    b, ok := ConvertX1{{.UType}}Set(a.ToInterfaceSlice()...)

	if !ok {
		t.Errorf("Not ok")
	}

	if !a.Equals(b) {
		t.Errorf("Expected %v but got %v", a, b)
	}
}
{{- if .Mutable}}

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
{{- end}}

func Test{{.UType}}SetContainsAll(t *testing.T) {
	a := NewX1{{.UType}}Set(8, 6, 7, 5, 3, 0, 9)

	if !a.ContainsAll(8, 6, 7, 5, 3, 0, 9) {
		t.Error("should contain phone number")
	}

	if a.ContainsAll(8, 6, 11, 5, 3, 0, 9) {
		t.Error("should not have all of these numbers")
	}
}

func Test{{.UType}}SetIsSubset(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 5, 7)
	b := NewX1{{.UType}}Set(3, 5, 7)
	c := NewX1{{.UType}}Set(3, 5, 7, 72)

	if !b.IsSubset(a) {
		t.Errorf("Expected '%s' to be a subset of '%s'", b, a)
	}

	if c.IsSubset(a) {
		t.Errorf("Expected '%s' not to be a subset of '%s'", c, a)
	}

    // check correct nil handling
    a = nil
	if b.IsSubset(a) {
		t.Errorf("Expected '%s' not to be a subset of '%s'", b, a)
	}
	if !a.IsSubset(b) {
		t.Errorf("Expected '%s' to be a subset of '%s'", a, b)
	}
}

func Test{{.UType}}SetIsSuperSet(t *testing.T) {
	a := NewX1{{.UType}}Set(9, 5, 2, 1, 11)
	b := NewX1{{.UType}}Set(5, 2, 11)
	c := NewX1{{.UType}}Set(5, 2, 11, 42)

	if !a.IsSuperset(b) {
		t.Errorf("Expected '%s' to be a superset of '%s'", a, b)
	}

	if a.IsSuperset(c) {
		t.Errorf("Expected '%s' not to be a superset of '%s'", a, b)
	}

    // check correct nil handling
    a = nil
	if !b.IsSuperset(a) {
		t.Errorf("Expected '%s' not to be a superset of '%s'", b, a)
	}
	if a.IsSuperset(b) {
		t.Errorf("Expected '%s' to be a superset of '%s'", a, b)
	}
}

func Test{{.UType}}SetUnion(t *testing.T) {
	a := NewX1{{.UType}}Set()

	b := NewX1{{.UType}}Set(1, 2, 3, 4, 5)

	c := a.Union(b)

	if !c.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5)) {
		t.Errorf("Expected 5 but got %v", c)
	}

	d := NewX1{{.UType}}Set(10, 14, 0)

	e := c.Union(d)
	if !e.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5, 10, 14, 0)) {
		t.Errorf("Expected 8 but got %v", e)
	}

	f := NewX1{{.UType}}Set(14, 3)

	g := f.Union(e)
	if !g.Equals(NewX1{{.UType}}Set(1, 2, 3, 4, 5, 10, 14, 0)) {
		t.Errorf("Expected 8 but got %v", g)
	}

    // check correct nil handling
    a = nil
	c = a.Union(b)
	d = b.Union(a)

	if !c.Equals(b) {
		t.Errorf("Expected 5 but got %s", c)
	}
	if !d.Equals(b) {
		t.Errorf("Expected 5 but got %s", d)
	}
}

func Test{{.UType}}SetIntersection(t *testing.T) {
	a1 := NewX1{{.UType}}Set(1, 3, 5, 7)
	a2 := NewX1{{.UType}}Set(1, 3, 5, 7, 10)

	b1 := NewX1{{.UType}}Set(0, 2, 4, 6)
	b2 := NewX1{{.UType}}Set(2, 4, 6, 10)

	c1 := a1.Intersect(a2)
	c2 := b1.Intersect(a1)

	if !c1.Equals(NewX1IntSet(1, 3, 5, 7)) {
		t.Errorf("Expected 4 but got %v", c1)
	}

	if c2.NonEmpty() {
		t.Errorf("Expected 4 but got %v", c2)
	}

	d1 := a1.Intersect(a2)
	d2 := b2.Intersect(a2)

	if !d1.Equals(NewX1IntSet(1, 3, 5, 7)) {
		t.Errorf("Expected 4 but got %s", d1)
	}

	if !d2.Equals(NewX1IntSet(10)) {
		t.Errorf("Expected 1 but got %s", d2)
	}

    // check correct nil handling
    a1 = nil
	c1 = a1.Intersect(b1)
	d1 = b1.Intersect(a1)

	if c1.NonEmpty() {
		t.Errorf("Expected empty but got %s", c1)
	}
	if d1.NonEmpty() {
		t.Errorf("Expected empty but got %s", d1)
	}
}

func Test{{.UType}}SetDifference(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3)
	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)
	d := b.Difference(a)

	if !c.Equals(NewX1IntSet(2)) {
		t.Errorf("Expected [2] but got %s", c)
	}

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	if !c.Equals(a) {
		t.Errorf("Expected none but got %s", c)
	}
	if !d.Equals(b) {
		t.Errorf("Expected 5 but got %s", d)
	}
}

func Test{{.UType}}SetSymmetricDifference(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 50)
	b := NewX1{{.UType}}Set(1, 3, 4, 5, 6, 99)

	c := a.SymmetricDifference(b)
	d := b.SymmetricDifference(a)

	if !c.Equals(NewX1IntSet(2, 4, 5, 6, 50, 99)) {
		t.Errorf("Expected none but got %s", c)
	}

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	if !c.Equals(a) {
		t.Errorf("Expected none but got %s", c)
	}
	if !d.Equals(b) {
		t.Errorf("Expected 5 but got %s", d)
	}
}

func Test{{.UType}}SetEquals(t *testing.T) {
	a := NewX1{{.UType}}Set()
	b := NewX1{{.UType}}Set()

	if !a.Equals(b) {
		t.Errorf("Expected '%s' to equal '%s'", a, b)
	}

	c := NewX1{{.UType}}Set(1, 3, 5, 6, 8)
	d := NewX1{{.UType}}Set(1, 3, 5, 6, 9)

	if c.Equals(d) {
		t.Errorf("Expected '%s' not to equal '%s'", c, d)
	}

    // check correct nil handling
    a = nil
    if a.Equals(b) != b.Equals(a) {
        t.Fail()
    }
}

func Test{{.UType}}SetSend(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	b := BuildX1{{.UType}}SetFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%s' to equal '%s'", a, b)
	}
{{- if .Mutable}}

    // check correct nil handling
	a = nil
	b = BuildX1{{.UType}}SetFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a{{.M}}, b{{.M}})
	}
{{- end}}
}

func Test{{.UType}}SetForall(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	found := a.Forall(func(v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Forall(func(v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Forall(func(v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}
}

func Test{{.UType}}SetExists(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	found := a.Exists(func(v int) bool {
		return v > 2
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Exists(func(v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Exists(func(v int) bool {
		return v > 0
	})

	if found {
		t.Errorf("Expected not to find.")
	}
}

func Test{{.UType}}SetForeach(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)
	s := 0

	a.Foreach(func(v int) {
		s += v
	})

	if s != 10 {
		t.Errorf("Got %d", s)
	}

    // check correct nil handling
    a = nil
	a.Foreach(func(v int) {
		s += v
	})

	if s != 10 {
		t.Errorf("Got %d", s)
	}
}

func Test{{.UType}}SetFilter(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}Set(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%s'", b)
	}

    // check correct nil handling
	a = nil
	a.Filter(func(v int) bool {
		return v > 2
	})
}

func Test{{.UType}}SetPartition(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}Set(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%s'", b)
	}

	if !c.Equals(NewX1{{.UType}}Set(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%s'", c)
	}

    // check correct nil handling
	a = nil
	a.Partition(func(v int) bool {
		return v > 2
	})
}

func Test{{.UType}}SetTransform(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.Map(func(v {{.Type}}) {{.Type}} {
		return v * v
	})

	if !b.Equals(NewX1{{.UType}}Set(1, 4, 9, 16)) {
		t.Errorf("Expected '1, 4, 9, 16' but got '%s'", b)
	}

    // check correct nil handling
	a = nil
	a.Map(func(v {{.Type}}) {{.Type}} {
		return v * v
	})
}

func Test{{.UType}}SetFlatMap(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.FlatMap(func(v {{.Type}}) []{{.Type}} {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})

    exp := NewX1{{.UType}}Set(2, 3, 4, 6, 6, 9)
	if !b.Equals(exp) {
		t.Errorf("Expected '%s' but got '%s'", exp, b)
	}

    // check correct nil handling
	a = nil
	a.FlatMap(func(v {{.Type}}) []{{.Type}} {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})
}

func Test{{.UType}}SetStringMap(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2, 3, 4)

	b := a.StringMap()

	for _, c := range a.ToSlice() {
		s := fmt.Sprintf("%d", c)
		if _, ok := b[s]; !ok {
			t.Errorf("Expected '%s' but got '%v'", s, b)
		}
	}

    // check correct nil handling
	a = nil
	a.StringMap()
}
{{- if .Mutable}}

func TestMutable{{.UType}}SetClear(t *testing.T) {
	a := NewX1{{.UType}}Set(2, 5, 9, 10)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}

    // check correct nil handling
	a = nil
	a.Clear()
}

func TestMutable{{.UType}}SetClone(t *testing.T) {
	a := NewX1{{.UType}}Set(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%s' to equal '%s'", a, b)
	}

	a.Add(3)
	if a.Equals(b) {
		t.Errorf("Expected '%s' not to equal '%s'", a, b)
	}

	c := a.Clone()
	c.Remove(1)

	if a.Equals(c) {
		t.Errorf("Expected '%s' not to equal '%s'", a, c)
	}

    // check correct nil handling
	a = nil
	a.Clone()
}
{{- end}}

func Test{{.UType}}SetMkString(t *testing.T) {
	a := NewX1{{.UType}}Set(13, 4)

	c := a.MkString("|")

	if c != "13|4" && c != "4|13" {
		t.Errorf("Expected '13|4' but got %q", c)
	}

    // check correct nil handling
	a = nil
	a.MkString("|")
}

func Test{{.UType}}SetMkString3(t *testing.T) {
	a := NewX1{{.UType}}Set(13, 4)

	c := a.MkString3("<", ", ", ">")

	if c != "<13, 4>" && c != "<4, 13>" {
		t.Errorf("Expected '13|4' but got %q", c)
	}

    // check correct nil handling
	a = nil
	a.MkString3("<", ",", ">")
}

{{if .GobEncode}}
func Test{{.UType}}SetGobEncode(t *testing.T) {
	a := NewX1{{.UType}}Set(13, 4, 7, -2, 9)
	b := NewX1{{.UType}}Set()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

	if err != nil {
		t.Errorf("Got %v", err)
	}

    err = gob.NewDecoder(buf).Decode(&b)

	if err != nil {
		t.Errorf("Got %v", err)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%s' but got '%s'", a, b)
	}
}

{{end}}
func Test{{.UType}}SetJsonEncode(t *testing.T) {
	a := NewX1{{.UType}}Set(13, 4, 7, -2, 9)
	b := NewX1{{.UType}}Set()

    buf, err := json.Marshal(a)

	if err != nil {
		t.Errorf("Got %v", err)
	}

    err = json.Unmarshal(buf, &b)

	if err != nil {
		t.Errorf("Got %v", err)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%s' but got '%s'", a, b)
	}
}
