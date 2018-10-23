// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
{{- if .GobEncode}}
    "bytes"
    "encoding/gob"
{{- end}}
	"testing"
)

func TestIm{{.UKey}}{{.UType}}MapToSlice(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map1(1, 2)
	s := a.ToSlice()

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	if len(s) != 1 {
		t.Errorf("Expected 1 but got %d", len(s))
	}
}

func TestIm{{.UKey}}{{.UType}}MapContainsAllKeys(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 6}, TX1{{.UKey}}{{.UType}}Tuple{1, 10}, TX1{{.UKey}}{{.UType}}Tuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("%+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("%+v", a)
	}
}

func TestIm{{.UKey}}{{.UType}}MapCardinality(t *testing.T) {
	a1 := NewTX1{{.UKey}}{{.UType}}Map()
	a2 := NewTX1{{.UKey}}{{.UType}}Map1(1, 2)

	if a1.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a1.Size())
	}

	if a2.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a2.Size())
	}
}

func TestIm{{.UKey}}{{.UType}}MapEquals(t *testing.T) {
	a1 := NewTX1{{.UKey}}{{.UType}}Map()
	b1 := NewTX1{{.UKey}}{{.UType}}Map()
	a2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append3(8, 19, 10, 4, 3, 1)...)

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
//	a := NewTX1{{.UKey}}{{.UType}}Map(1, 2, 3, 4)
//
//	b := NewTX1{{.UKey}}{{.UType}}Map()
//	for val := range a.Send() {
//		b.Add(val)
//	}
//
//	if !a.Equals(b) {
//		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
//	}
//}

func Test{{.UType}}MapFilter(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 1}, TX1{{.UKey}}{{.UType}}Tuple{1, 2}, TX1{{.UKey}}{{.UType}}Tuple{2, 3})

	b := a.Filter(func(k, v int) bool {
		return v > 2
	})

	exp := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{2, 3})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp, b)
	}
}

func Test{{.UType}}MapPartition(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 2, 4).Values(4, 11, 0)...)

	b, c := a.Partition(func(k, v int) bool {
		return v > 5
	})

	exp1 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{2, 11})
	if !b.Equals(exp1) {
		t.Errorf("Expected '%+v' but got '%+v'", exp1{{.M}}, b{{.M}})
	}

	exp2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 4}, TX1{{.UKey}}{{.UType}}Tuple{4, 0})
	if !c.Equals(exp2) {
		t.Errorf("Expected '%+v' but got '%+v'", exp2{{.M}}, c{{.M}})
	}
}

func Test{{.UType}}MapTransform(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 9, 10).Values(6, 10, 5)...)

	b := a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})

	exp := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{9, 36}, TX1{{.UKey}}{{.UType}}Tuple{10, 100}, TX1{{.UKey}}{{.UType}}Tuple{11, 25})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp{{.M}}, b{{.M}})
	}
}

func Test{{.UType}}MapFlatMap(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(2, 1, 18).Values(6, 10, 5)...)

	b := a.FlatMap(func(k {{.Key}}, v {{.Type}}) []TX1{{.UPrefix}}{{.UKey}}{{.UType}}Tuple {
	    if k > 3 {
	        return nil
	    }
		return []TX1{{.UPrefix}}{{.UKey}}{{.UType}}Tuple{
		    {k-1, v+1},
		    {k+1, v+2},
		}
	})

	exp := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{1, 7}, TX1{{.UKey}}{{.UType}}Tuple{3, 8},
	    TX1{{.UKey}}{{.UType}}Tuple{0, 11}, TX1{{.UKey}}{{.UType}}Tuple{2, 12})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp{{.M}}, b{{.M}})
	}
}

{{if .Mutable}}

func TestMu{{.UKey}}{{.UType}}MapRemove(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map1(3, 1)

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
	a := NewTX1{{.UKey}}{{.UType}}Map1(13, 1)

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
	a := NewTX1{{.UKey}}{{.UType}}Map1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("%+v", a)
	}
}

func TestMu{{.UKey}}{{.UType}}MapClone(t *testing.T) {
	a1 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append2(1, 9, 2, 8)...)
	a2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append3(1, 9, 2, 8, 3, 3)...)

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

func Test{{.UType}}MapMkString(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 4}, TX1{{.UKey}}{{.UType}}Tuple{4, 0})

	c := a.MkString("|")

	if c != "8:4|4:0" && c != "4:0|8:4" {
		t.Errorf("Expected '8:4|4:0' but got %q", c)
	}
}

func Test{{.UType}}MapMkString3(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 4}, TX1{{.UKey}}{{.UType}}Tuple{4, 0})

	c := a.MkString3("<", ",", ">")

	if c != "<8:4,4:0>" && c != "<4:0,8:4>" {
		t.Errorf("Expected '<8:4,4:0>' but got %q", c)
	}
}

{{if .GobEncode}}
func Test{{.UType}}MapGobEncode(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 9, -2, 8, 3, 3).Values(-5, 10, 13, 17, 19, 23)...)
	b := NewTX1{{.UKey}}{{.UType}}Map()

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
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}
}

{{end}}

