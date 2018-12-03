// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
{{- if .GobEncode}}
    "bytes"
    "encoding/gob"
{{- end}}
{{- if eq .Key "string"}}
    "encoding/json"
	"strings"
{{- end}}
	"reflect"
	"sort"
	"testing"
	. "github.com/onsi/gomega"
)

func TestIm{{.UKey}}{{.UType}}MapToSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewTX1{{.UKey}}{{.UType}}Map1(1, 2)
	s := a.ToSlice()

	g.Expect(a.Size()).To(Equal(1))
	g.Expect(len(s)).To(Equal(1))

    // check correct nil handling
    a = nil
    a.ToSlice()
}

func TestIm{{.UKey}}{{.UType}}MapSize(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewTX1{{.UKey}}{{.UType}}Map()
	a2 := NewTX1{{.UKey}}{{.UType}}Map1(1, 2)

	g.Expect(a1.Size()).To(Equal(0))
	g.Expect(a2.Size()).To(Equal(1))

    // check correct nil handling
    a1 = nil
    a1.Size()
}

func Test{{.UKey}}{{.UType}}Keys(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 4, 2).Values(4, 0, 5)...)

    k := a.Keys()
    sort.Ints(k)
	if !reflect.DeepEqual(k, []int{2, 4, 8}) {
		t.Errorf("Expected [2,4,8] but got %v", k)
	}

    // check correct nil handling
    a = nil
    a.Keys()
}

func Test{{.UKey}}{{.UType}}Values(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 4, 2).Values(4, 0, 5)...)

    v := a.Values()
    sort.Ints(v)
	if !reflect.DeepEqual(v, []int{0, 4, 5}) {
		t.Errorf("Expected [0,4,5] but got %v", v)
	}

    // check correct nil handling
    a = nil
    a.Values()
}

func TestIm{{.UKey}}{{.UType}}MapContainsAllKeys(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 6}, TX1{{.UKey}}{{.UType}}Tuple{1, 10}, TX1{{.UKey}}{{.UType}}Tuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("Got %+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
    a = nil
    a.ContainsAllKeys()
}

func TestIm{{.UKey}}{{.UType}}MapEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewTX1{{.UKey}}{{.UType}}Map()
	b1 := NewTX1{{.UKey}}{{.UType}}Map()
	a2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuples{}.Append3(8, 19, 10, 4, 3, 1)...)

	g.Expect(a1.Equals(b1)).To(BeTrue())
	g.Expect(b1.Equals(a1)).To(BeTrue())
	g.Expect(a2.Equals(b1)).To(BeFalse())
	g.Expect(a2.Equals(b3)).To(BeFalse())
	g.Expect(a3.Equals(a2)).To(BeFalse())
	g.Expect(a3.Equals(b3)).To(BeTrue())
	g.Expect(b3.Equals(a3)).To(BeTrue())

    // check correct nil handling
    a1 = nil
    a1.Equals(b1)
    b1.Equals(a1)
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

func Test{{.UType}}MapForall(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	found := a.Forall(func(k, v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Forall(func(k, v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Forall(func(k, v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}

    // check correct nil handling
    a = nil
	a.Forall(func(k, v int) bool {
		return v > 0
	})
}

func Test{{.UType}}MapExists(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	found := a.Exists(func(k, v int) bool {
		return v > 2
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Exists(func(k, v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Exists(func(k, v int) bool {
		return v > 0
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	a.Exists(func(k, v int) bool {
		return v > 2
	})
}

func Test{{.UType}}MapForeach(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)
	s := 0

	a.Foreach(func(k, v int) {
		s += v
	})

	if s != 6 {
		t.Errorf("Got %d", s)
	}

    // check correct nil handling
    a = nil
	a.Foreach(func(k, v int) {
		s += v
	})

	if s != 6 {
		t.Errorf("Got %d", s)
	}

    // check correct nil handling
    a = nil
	a.Foreach(func(k, v int) {
		s += v
	})
}

func Test{{.UType}}MapFind(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	b, found := a.Find(func(k, v int) bool {
		return v > 2
	})

	exp := TX1{{.UKey}}{{.UType}}Tuple{2, 3}
	g.Expect(found).To(BeTrue())
	g.Expect(b).To(Equal(exp))

    // check correct nil handling
    a = nil
	a.Find(func(k, v int) bool {
		return v > 2
	})
}

func Test{{.UType}}MapFilter(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	b := a.Filter(func(k, v int) bool {
		return v > 2
	})

	exp := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{2, 3})
	g.Expect(b.Equals(exp)).To(BeTrue())

    // check correct nil handling
    a = nil
	a.Filter(func(k, v int) bool {
		return v > 2
	})
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

	exp2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 4).Values(4, 0)...)
	if !c.Equals(exp2) {
		t.Errorf("Expected '%+v' but got '%+v'", exp2{{.M}}, c{{.M}})
	}

    // check correct nil handling
    a = nil
	a.Partition(func(k, v int) bool {
		return v > 2
	})
}

func Test{{.UType}}MapTransform(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 9, 10).Values(6, 10, 5)...)

	b := a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})

	exp := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(9, 10, 11).Values(36, 100, 25)...)
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp{{.M}}, b{{.M}})
	}

    // check correct nil handling
    a = nil
	a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})
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

    // check correct nil handling
    a = nil
	a.FlatMap(func(k {{.Key}}, v {{.Type}}) []TX1{{.UPrefix}}{{.UKey}}{{.UType}}Tuple {
        return nil
	})
}
{{- if .Mutable}}

func TestMu{{.UKey}}{{.UType}}MapPop(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	v, y := a.Pop(8)

	if !y {
		t.Errorf("Expected popped value")
	}
	if v != 2 {
		t.Errorf("Expected popped 2 but got %d", v)
	}
	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	v, y = a.Pop(8)

	if y {
		t.Errorf("Expected no popped value")
	}
	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	if !(a.ContainsKey(1) && a.ContainsKey(2)) {
		t.Errorf("Got %+v", a)
	}

	a.Pop(1)
	a.Pop(2)

	if a.NonEmpty() {
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
	a = nil
	_, y = a.Pop(1)
	if y {
		t.Errorf("Expected no popped value")
	}
}

func TestMu{{.UKey}}{{.UType}}MapRemove(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 8, 2).Values(1, 2, 3)...)

	a.Remove(8)

	if a.Size() != 2 {
		t.Errorf("Expected 2 but got %d", a.Size())
	}

	if !(a.ContainsKey(1) && a.ContainsKey(2)) {
		t.Errorf("Got %+v", a)
	}

	a.Remove(2)
	a.Remove(1)

	if a.NonEmpty() {
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
	a = nil
	a.Remove(1)
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
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
    a = nil
	if a.ContainsKey(71) {
		t.Error("should not contain 71")
	}
}

func TestMu{{.UKey}}{{.UType}}MapClear(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
    a = nil
    a.Clear()
}

func TestMu{{.UKey}}{{.UType}}MapClone(t *testing.T) {
	a1 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 2).Values(9, 8)...)
	a2 := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 2, 3).Values(9, 8, 7)...)

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

    // check correct nil handling
    a1 = nil
    a1.Clone()
}
{{- end}}

func Test{{.UType}}MapMkString(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(8, 4).Values(4, 0)...)

	c := a.MkString("|")

	if c != "8:4|4:0" && c != "4:0|8:4" {
		t.Errorf("Expected '8:4|4:0' but got %q", c)
	}

    // check correct nil handling
    a = nil
	a.MkString("|")
}

func Test{{.UType}}MapMkString3(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Tuple{8, 4}, TX1{{.UKey}}{{.UType}}Tuple{4, 0})

	c := a.MkString3("<", ",", ">")

	if c != "<8:4,4:0>" && c != "<4:0,8:4>" {
		t.Errorf("Expected '<8:4,4:0>' but got %q", c)
	}

    // check correct nil handling
    a = nil
	a.MkString3("<", ",", ">")
}
{{- if .GobEncode}}

func Test{{.UType}}MapGobEncode(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 9, -2, 8, 3, 3).Values(-5, 10, 13, 17, 19, 23)...)
	b := NewTX1{{.UKey}}{{.UType}}Map()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

	if err != nil {
		t.Errorf("%v", err)
	}

    err = gob.NewDecoder(buf).Decode(&b)

	if err != nil {
		t.Errorf("%v", err)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}
}
{{- end}}
{{- if eq .Key "string"}}

func Test{{.UType}}MapJsonEncode(t *testing.T) {
	a := NewTX1{{.UKey}}{{.UType}}Map(TX1{{.UKey}}{{.UType}}Zip(1, 9, -2, 8, 3, 3).Values(-5, 10, 13, 17, 19, 23)...)
	b := NewTX1{{.UKey}}{{.UType}}Map()

    buf, err := json.Marshal(a)

	if err != nil {
		t.Errorf("%v", err)
	}

    got := strings.TrimSpace(string(buf))
    exp := `{"-2":13,"1":-5,"3":23,"8":17,"9":10}`
	if got != exp {
		t.Errorf("Expected %s but got '%+v'", exp, got)
	}

    err = json.Unmarshal(buf, &b)

	if err != nil {
		t.Errorf("%v", err)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}
}
{{- end}}
