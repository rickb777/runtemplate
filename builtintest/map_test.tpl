package {{.Package}}

import (
	"testing"
)

func TestIm{{.UKey}}{{.UType}}MapToSlice(t *testing.T) {
	a := NewTX{{.UKey}}{{.UType}}Map1(1, 2)
	s := a.ToSlice()

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	if len(s) != 1 {
		t.Errorf("Expected 1 but got %d", len(s))
	}
}

func TestIm{{.UKey}}{{.UType}}MapContainsAllKeys(t *testing.T) {
	a := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuple{8, 6}, TX{{.UKey}}{{.UType}}Tuple{1, 10}, TX{{.UKey}}{{.UType}}Tuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("%+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("%+v", a)
	}
}

func TestIm{{.UKey}}{{.UType}}MapCardinality(t *testing.T) {
	a1 := NewTX{{.UKey}}{{.UType}}Map()
	a2 := NewTX{{.UKey}}{{.UType}}Map1(1, 2)

	if a1.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a1.Size())
	}

	if a2.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a2.Size())
	}
}

func TestIm{{.UKey}}{{.UType}}MapEquals(t *testing.T) {
	a1 := NewTX{{.UKey}}{{.UType}}Map()
	b1 := NewTX{{.UKey}}{{.UType}}Map()
	a2 := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuples{}.Append3(8, 19, 10, 4, 3, 1)...)

	if !a1.Equals(b1) {
		t.Errorf("Expected '%+v' to equal '%+v'", a1, b1)
	}

	if !b1.Equals(a1) {
		t.Errorf("Expected '%+v' to equal '%+v'", a1, b1)
	}

	if a2.Equals(b1) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a2, b1)
	}

	if a2.Equals(b3) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a2, b3)
	}

	if a3.Equals(a2) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a2, b3)
	}

	if !a3.Equals(b3) {
		t.Errorf("Expected '%+v' to equal '%+v'", a3, b3)
	}

	if !b3.Equals(a3) {
		t.Errorf("Expected '%+v' to equal '%+v'", a3, b3)
	}
}

//func TestIm{{.UKey}}{{.UType}}MapSend(t *testing.T) {
//	a := NewTX{{.UKey}}{{.UType}}Map(1, 2, 3, 4)
//
//	b := NewTX{{.UKey}}{{.UType}}Map()
//	for val := range a.Send() {
//		b.Add(val)
//	}
//
//	if !a.Equals(b) {
//		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
//	}
//}

{{if .Mutable}}

func TestMu{{.UKey}}{{.UType}}MapRemove(t *testing.T) {
	a := NewTX{{.UKey}}{{.UType}}Map1(3, 1)

	a.Put(1, 5)
	a.Put(2, 5)
	a.Remove(3)

	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	if !(a.ContainsKey(1) && a.ContainsKey(2)) {
		t.Errorf("%+v", a)
	}

	a.Remove(2)
	a.Remove(1)

	if a.Size() != 0 {
		t.Errorf("%+v", a)
	}
}

func TestMu{{.UKey}}{{.UType}}MapContainsKey(t *testing.T) {
	a := NewTX{{.UKey}}{{.UType}}Map1(13, 1)

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
		t.Errorf("%+v", a)
	}
}

func TestMu{{.UKey}}{{.UType}}MapClear(t *testing.T) {
	a := NewTX{{.UKey}}{{.UType}}Map1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("%+v", a)
	}
}

func TestMu{{.UKey}}{{.UType}}MapClone(t *testing.T) {
	a1 := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuples{}.Append2(1, 9, 2, 8)...)
	a2 := NewTX{{.UKey}}{{.UType}}Map(TX{{.UKey}}{{.UType}}Tuples{}.Append3(1, 9, 2, 8, 3, 3)...)

	b := a1.Clone()

	if !a1.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a1, b)
	}

	if a2.Equals(b) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a2, b)
	}

	c := a2.Clone()
	c = c.Filter(func(k, v int) bool {
		return k != 1
	})

	if a2.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a2, c)
	}
}

{{end}}