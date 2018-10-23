// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Append:{{.Append}} Find:{{.Find}} Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
{{- if .GobEncode}}
    "bytes"
    "encoding/gob"
{{- end}}
	"testing"
)

func {{.LType}}RangeOf(from, to int) []int {
	n := 1 + to - from
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i + from
	}
	return a
}

func TestNew{{.UType}}List(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3)

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if a.Get(1) != 2 {
		t.Errorf("Expected 2 but got %d", a.Get(1))
	}

	if a.IsSet() {
		t.Error("Expected not a set")
	}

	if !a.IsSequence() {
		t.Error("Expected a sequence")
	}
}

func TestConvert{{.UType}}List(t *testing.T) {
	a, ok := ConvertX1{{.UType}}List(1, 5.1, uint8(2), 7, 3)

	if !ok {
		t.Errorf("Not ok")
	}

	if !a.Equals(NewX1{{.UType}}List(1, 5, 2, 7, 3)) {
		t.Errorf("Expected 1,5,2,7,3 but got %v", a)
	}

    b, ok := ConvertX1{{.UType}}List(a.ToInterfaceSlice()...)

	if !ok {
		t.Errorf("Not ok")
	}

	if !a.Equals(b) {
		t.Errorf("Expected %v but got %v", a, b)
	}
}

{{if .Append}}
func Test{{.UType}}ListAppend(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3)

	b := a.Append(4, 5).Append(6, 7)

	if b.Size() != 7 {
		t.Errorf("Expected 5 but got %d", a.Size())
	}

	if b.Get(3) != 4 {
		t.Errorf("Expected 4 but got %d", b.Get(3))
	}

	if b.Last() != 7 {
		t.Errorf("Expected 7 but got %d", b.Last())
	}
}

func Test{{.UType}}ListDoInsertAt(t *testing.T) {
    cases := []struct{
        i int
        more []{{.PType}}
        exp *X1{{.UType}}List
    }{
        {
            0,
            []{{.PType}}{10, 11},
            NewX1{{.UType}}List(10, 11, 1, 2, 3, 4, 5, 6),
        },
        {
            2,
            []{{.PType}}{10, 11, 12},
            NewX1{{.UType}}List(1, 2, 10, 11, 12, 3, 4, 5, 6),
        },
        {
            6,
            []{{.PType}}{10, 11},
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6, 10, 11),
        },
        {
            3,
            []{{.PType}}{},
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
        },
    }

    for i, c := range cases {
        act := NewX1{{.UType}}List(1, 2, 3, 4, 5, 6)
        r := act.DoInsertAt(c.i, c.more...)

        if !act.Equals(c.exp) {
            t.Errorf("%d: Expected %+v but got %+v", i, c.exp{{.M}}, act{{.M}})
        }
        if !act.Equals(r) {
            t.Errorf("%d: Expected %+v but got %+v", i, r{{.M}}, act{{.M}})
        }
    }
}

func Test{{.UType}}ListDoDeleteAt(t *testing.T) {
    cases := []struct{
        i, n int
        act, exp *X1{{.UType}}List
    }{
        {
            0, 2,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(3, 4, 5, 6),
        },
        {
            2, 2,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 5, 6),
        },
        {
            4, 2,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3, 4),
        },
        {
            3, 0,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteAt(c.i, c.n)

        if !c.act.Equals(c.exp) {
            t.Errorf("%d: Expected %+v but got %+v", i, c.exp{{.M}}, c.act{{.M}})
        }
        if !c.act.Equals(r) {
            t.Errorf("%d: Expected %+v but got %+v", i, r{{.M}}, c.act{{.M}})
        }
    }
}

func Test{{.UType}}ListDoDeleteFirst(t *testing.T) {
    cases := []struct{
        n int
        act, exp *X1{{.UType}}List
    }{
        {
            0,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
        },
        {
            1,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(2, 3, 4, 5, 6),
        },
        {
            3,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(4, 5, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteFirst(c.n)

        if !c.act.Equals(c.exp) {
            t.Errorf("%d: Expected %+v but got %+v", i, c.exp{{.M}}, c.act{{.M}})
        }
        if !c.act.Equals(r) {
            t.Errorf("%d: Expected %+v but got %+v", i, r{{.M}}, c.act{{.M}})
        }
    }
}

func Test{{.UType}}ListDoDeleteLast(t *testing.T) {
    cases := []struct{
        n int
        act, exp *X1{{.UType}}List
    }{
        {
            0,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
        },
        {
            1,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3, 4, 5),
        },
        {
            3,
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(1, 2, 3),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteLast(c.n)

        if !c.act.Equals(c.exp) {
            t.Errorf("%d: Expected %+v but got %+v", i, c.exp{{.M}}, c.act{{.M}})
        }
        if !c.act.Equals(r) {
            t.Errorf("%d: Expected %+v but got %+v", i, r{{.M}}, c.act{{.M}})
        }
    }
}

{{end}}
{{if and .Mutable .Numeric}}
func Test{{.UType}}ListDoKeepWhere(t *testing.T) {
    cases := []struct{
        act, exp *X1{{.UType}}List
    }{
        {
            NewX1{{.UType}}List(1, 2, 3, 4, 5, 6),
            NewX1{{.UType}}List(2, 4, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoKeepWhere(func (v {{.PType}}) bool {
            return v % 2 == 0
        })

        if !c.act.Equals(c.exp) {
            t.Errorf("%d: Expected %+v but got %+v", i, c.exp{{.M}}, c.act{{.M}})
        }
        if !c.act.Equals(r) {
            t.Errorf("%d: Expected %+v but got %+v", i, r{{.M}}, c.act{{.M}})
        }
    }
}

{{end}}
func Test{{.UType}}ListClone(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2)

	b := a.Clone()

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a{{.M}}, b{{.M}})
	}

	c := a.Clone().Tail()

	if a.Equals(c) {
		t.Errorf("Expected '%+v' not to equal '%+v'", a{{.M}}, c{{.M}})
	}
}

func Test{{.UType}}ListSend(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	b := BuildX1{{.UType}}ListFromChan(a.Send())

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' to equal '%+v'", a{{.M}}, b{{.M}})
	}
}

func Test{{.UType}}ListHeadTailLastInit(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	if a.Head() != 1 {
		t.Errorf("Expected 1 but got %d", a.Head())
	}

	if a.Last() != 4 {
		t.Errorf("Expected 4 but got %d", a.Last())
	}

	tail := a.Tail()
	if !tail.Equals(NewX1{{.UType}}List(2, 3, 4)) {
		t.Errorf("Expected '2, 3, 4' but got '%+v'", tail{{.M}})
	}

	init := a.Init()
	if !init.Equals(NewX1{{.UType}}List(1, 2, 3)) {
		t.Errorf("Expected '1, 2, 3' but got '%+v'", init{{.M}})
	}
}

{{if .Find}}
func Test{{.UType}}ListFind(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	b := a.Find(func(v int) bool {
		return v > 2
	})

	if b != 3 {
		t.Errorf("Expected '3, 4' but got '%+v'", b{{.M}})
	}
}

{{end}}
func Test{{.UType}}ListFilter(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b{{.M}})
	}
}

func Test{{.UType}}ListPartition(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	if !b.Equals(NewX1{{.UType}}List(3, 4)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b{{.M}})
	}

	if !c.Equals(NewX1{{.UType}}List(1, 2)) {
		t.Errorf("Expected '1, 2' but got '%+v'", c{{.M}})
	}
}

func Test{{.UType}}ListTransform(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4)

	b := a.Map(func(v int) int {
		return v * v
	})

    exp := NewX1{{.UType}}List(1, 4, 9, 16)
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp, b{{.M}})
	}
}

func Test{{.UType}}ListFlatMap(t *testing.T) {
	a := NewX1{{.UType}}List(1, 2, 3, 4, 5)

	b := a.FlatMap(func(v {{.Type}}) []{{.Type}} {
	    if v > 3 {
	        return nil
	    }
		return []{{.Type}}{v * 2, v * 3}
	})

    exp := NewX1{{.UType}}List(2, 3, 4, 6, 6, 9)
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp, b{{.M}})
	}
}

{{if .Mutable}}
func Test{{.UType}}ListSort(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	b := a.Sorted()

	if !a.Equals(NewX1{{.UType}}List(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", a{{.M}})
	}

	if !b.Equals(NewX1{{.UType}}List(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b{{.M}})
	}
}

func Test{{.UType}}ListStableSort(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	b := a.StableSorted()

	if !a.Equals(NewX1{{.UType}}List(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", a{{.M}})
	}

	if !b.Equals(NewX1{{.UType}}List(-2, 4, 7, 9, 13)) {
		t.Errorf("Expected '3, 4' but got '%+v'", b{{.M}})
	}
}

func Test{{.UType}}ListReverseOdd(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	b := a.Reverse()

	if b.Equals(a) {
		t.Errorf("Expected reverse of '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}

	c := b.Reverse()

	if !c.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, c{{.M}})
	}
}

func Test{{.UType}}ListReverseEven(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9, 17)

	b := a.Reverse()

	if b.Equals(a) {
		t.Errorf("Expected reverse of '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}

	c := b.Reverse()

	if !c.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, c{{.M}})
	}
}

func Test{{.UType}}ListShuffle(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 1000)...)

	b := a.Shuffle()

	if b.Equals(a) {
		t.Errorf("Expected shuffle of '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}

	// prove that the same set of numbers is present
	b.Sorted()

	if !b.Equals(a) {
		t.Errorf("Expected '%+v' but got '%+v'", a{{.M}}, b{{.M}})
	}
}

{{end}}
func Test{{.UType}}ListTake(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.Take(30)

	if b.Size() != 30 || b.Head() != 1 || b.Last() != 30 {
		t.Errorf("Expected list from 1 to 30 but got '%+v'", b{{.M}})
	}

	c := a.TakeLast(30)

	if c.Size() != 30 || c.Head() != 71 || c.Last() != 100 {
		t.Errorf("Expected list from 71 to 100 but got '%+v'", c{{.M}})
	}

	d := a.Take(101)

	if d.Size() != 100 || d.Head() != 1 || d.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", d{{.M}})
	}

	e := a.TakeLast(101)

	if e.Size() != 100 || e.Head() != 1 || e.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", e{{.M}})
	}
}

func Test{{.UType}}ListDrop(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.Drop(70)

	if b.Size() != 30 || b.Head() != 71 || b.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b{{.M}})
	}

	c := a.DropLast(75)

	if c.Size() != 25 || c.Head() != 1 || c.Last() != 25 {
		t.Errorf("Expected list from 1 to 25 but got '%+v'", c{{.M}})
	}

	d := a.Drop(101)

	if d.Size() != 0 {
		t.Errorf("Expected empty list but got '%+v'", d{{.M}})
	}

	e := a.DropLast(101)

	if e.Size() != 0 {
		t.Errorf("Expected empty list but got '%+v'", e{{.M}})
	}
}

func Test{{.UType}}ListTakeWhile(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.TakeWhile(func(v int) bool {
		return v <= 20
	})

	if b.Size() != 20 || b.Head() != 1 || b.Last() != 20 {
		t.Errorf("Expected list from 1 to 20 but got '%+v'", b{{.M}})
	}

	c := a.TakeWhile(func(v int) bool {
		return true
	})

	if c.Size() != 100 || c.Head() != 1 || c.Last() != 100 {
		t.Errorf("Expected list from 1 to 100 but got '%+v'", b{{.M}})
	}
}

func Test{{.UType}}ListDropWhile(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.DropWhile(func(v int) bool {
		return v <= 80
	})

	if b.Size() != 20 || b.Head() != 81 || b.Last() != 100 {
		t.Errorf("Expected list from 81 to 100 but got '%+v'", b{{.M}})
	}

	c := a.DropWhile(func(v int) bool {
		return true
	})

	if c.Size() != 0 {
		t.Errorf("Expected empty list but got '%+v'", b{{.M}})
	}
}

func Test{{.UType}}ListDistinctBy(t *testing.T) {
	a := NewX1{{.UType}}List(1, 1, 1, 2, 1, 2, 3, 4, 5, 3, 3, 5)

	c := a.DistinctBy(func(v1, v2 int) bool {
		return v1 == v2
	})

	if !c.Equals(NewX1{{.UType}}List(1, 2, 3, 4, 5)) {
		t.Errorf("Expected 1 to 5 but got '%+v'", c{{.M}})
	}
}

func Test{{.UType}}ListIndexWhere(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.IndexWhere(func(v int) bool {
		return v >= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.IndexWhere(func(v int) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.IndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 10)

	if d != 11 {
		t.Errorf("Expected 11 but got %d", d)
	}
}

func Test{{.UType}}ListLastIndexWhere(t *testing.T) {
	a := NewX1{{.UType}}List({{.LType}}RangeOf(1, 100)...)

	b := a.LastIndexWhere(func(v int) bool {
		return v <= 47
	})

	if b != 46 {
		t.Errorf("Expected 46 but got %d", b)
	}

	c := a.LastIndexWhere(func(v int) bool {
		return false
	})

	if c != -1 {
		t.Errorf("Expected -1 but got %d", c)
	}

	d := a.LastIndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 61)

	if d != 59 {
		t.Errorf("Expected 59 but got %d", d)
	}
}

func Test{{.UType}}ListMinBy(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	c := a.MinBy(func(v1, v2 int) bool {
		return v1 > v2
	})

	if c != 13 {
		t.Errorf("Expected 13 but got '%+v'", c)
	}
}

func Test{{.UType}}ListMaxBy(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	c := a.MaxBy(func(v1, v2 int) bool {
		return v1 > v2
	})

	if c != -2 {
		t.Errorf("Expected -2 but got '%+v'", c)
	}
}

func Test{{.UType}}ListMkString(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	c := a.MkString("|")

	if c != "13|4|7|-2|9" {
		t.Errorf("Expected '13|4|7|-2|9' but got %q", c)
	}
}

func Test{{.UType}}ListMkString3(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)

	c := a.MkString3("<", ", ", ">")

	if c != "<13, 4, 7, -2, 9>" {
		t.Errorf("Expected '<13, 4, 7, -2, 9>' but got %q", c)
	}
}

{{if .GobEncode}}
func Test{{.UType}}ListGobEncode(t *testing.T) {
	a := NewX1{{.UType}}List(13, 4, 7, -2, 9)
	b := NewX1{{.UType}}List()

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
