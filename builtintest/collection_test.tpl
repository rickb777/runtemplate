// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Append:{{.Append}}

package {{.Package}}

import (
	"testing"
	"sort"
	"encoding/json"
)

type Sizer interface {
	IsEmpty() bool
	NonEmpty() bool
	Size() int
}

func TestNewX{{.UType}}Collection(t *testing.T) {
	testEmpty{{.UType}}Collection(t, NewX1{{.UType}}Set())
	testEmpty{{.UType}}Collection(t, NewX1{{.UType}}List())
}

func testEmpty{{.UType}}Collection(t *testing.T, a Sizer) {
	if a.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a.Size())
	}
	if !a.IsEmpty() {
		t.Error("Expected empty")
	}
	if a.NonEmpty() {
		t.Error("Expected empty")
	}
}

func Test{{.UType}}ToSlice(t *testing.T) {
	test{{.UType}}ToSlice(t, NewX1{{.UType}}Set(1, 2, 3))
	test{{.UType}}ToSlice(t, NewX1{{.UType}}List(1, 2, 3))
}

func test{{.UType}}ToSlice(t *testing.T, a X1{{.UType}}Collection) {
	s := a.ToSlice()

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if len(s) != 3 {
		t.Errorf("Expected 3 but got %d", len(s))
	}
}

func Test{{.UType}}Exists(t *testing.T) {
	test{{.UType}}Exists1(t, NewX1{{.UType}}Set(1, 2, 3))
	test{{.UType}}Exists1(t, NewX1{{.UType}}List(1, 2, 3))
}

func test{{.UType}}Exists1(t *testing.T, a X1{{.UType}}Collection) {
	has2 := a.Exists(func(v int) bool {
		return v > 2
	})

	if !has2 {
		t.Errorf("Expected exists for %+v", a)
	}

	has5 := a.Exists(func(v int) bool {
		return v > 5
	})

	if has5 {
		t.Errorf("Expected not exists for %+v", a)
	}
}

func Test{{.UType}}Forall(t *testing.T) {
	test{{.UType}}Forall(t, NewX1{{.UType}}Set(1, 2, 3))
	test{{.UType}}Forall(t, NewX1{{.UType}}List(1, 2, 3))
}

func test{{.UType}}Forall(t *testing.T, a X1{{.UType}}Collection) {
	has1 := a.Forall(func(v int) bool {
		return v >= 1
	})

	if !has1 {
		t.Errorf("Expected forall for %+v", a)
	}

	has2 := a.Forall(func(v int) bool {
		return v >= 2
	})

	if has2 {
		t.Errorf("Expected not forall for %+v", a)
	}
}

func Test{{.UType}}Find(t *testing.T) {
	test{{.UType}}Find(t, NewX1{{.UType}}List(1, 2, 3, 4))
	test{{.UType}}Find(t, NewX1{{.UType}}Set(1, 2, 3, 4))
}

func test{{.UType}}Find(t *testing.T, a X1{{.UType}}Collection) {
	b, e := a.Find(func(v int) bool {
		return v > 2
	})

	if !e {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	if b != 3 {
		t.Errorf("Expected '3, 4' but got '%+v'", b)
	}

	c, e := a.Find(func(v int) bool {
		return v > 20
	})

	if e {
		t.Errorf("Expected false but got '%+v'", c)
	}
}


func Test{{.UType}}CountBy(t *testing.T) {
	test{{.UType}}CountBy(t, NewX1{{.UType}}Set(1, 2, 3))
	test{{.UType}}CountBy(t, NewX1{{.UType}}List(1, 2, 3))
}

func test{{.UType}}CountBy(t *testing.T, a X1{{.UType}}Collection) {
	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	if n != 2 {
		t.Errorf("Expected 2 but got %d", n)
	}
}

func Test{{.UType}}Foreach(t *testing.T) {
	test{{.UType}}Foreach(t, NewX1{{.UType}}Set(1, 2, 3))
	test{{.UType}}Foreach(t, NewX1{{.UType}}List(1, 2, 3))
}

func test{{.UType}}Foreach(t *testing.T, a X1{{.UType}}Collection) {
	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	if sum1 != 6 {
		t.Errorf("Expected 6 but got %d", sum1)
	}
}

func Test{{.UType}}Contains(t *testing.T) {
	test{{.UType}}Contains(t, NewX1{{.UType}}Set(71, 1, 7, 13))
	test{{.UType}}Contains(t, NewX1{{.UType}}List(71, 1, 7, 13))
}

func test{{.UType}}Contains(t *testing.T, a X1{{.UType}}Collection) {
	if !a.Contains(71) {
		t.Error("should contain 71")
	}

	if a.Contains(72) {
		t.Error("should not contain 72")
	}

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("should contain 13, 7, 1")
	}

	if !a.ContainsAll(1, 7, 13) {
		t.Error("should contain all 1, 7, 13")
	}

	if a.ContainsAll(1, 3, 5, 7, 9, 11, 13) {
		t.Error("should not contain all 1, 3, 5, 7, 9, 11, 13")
	}
}

func Test{{.UType}}MinMaxSum(t *testing.T) {
	test{{.UType}}MinMaxSum(t, NewX1{{.UType}}Set(10, 71, 3, 7, 13))
	test{{.UType}}MinMaxSum(t, NewX1{{.UType}}List(10, 71, 3, 7, 13))
}

func test{{.UType}}MinMaxSum(t *testing.T, a X1{{.UType}}Collection) {
	if a.Min() != 3 {
		t.Errorf("Expected 3 but got %d", a.Min())
	}
	if a.Max() != 71 {
		t.Errorf("Expected 71 but got %d", a.Max())
	}
	if a.Sum() != 104 {
		t.Errorf("Expected 104 but got %d", a.Sum())
	}
}

func Test{{.UType}}Stringer(t *testing.T) {
	test{{.UType}}Stringer(t, NewX1{{.UType}}Set(10, 71, 3, 7, 13), false)
	test{{.UType}}Stringer(t, NewX1{{.UType}}List(10, 71, 3, 7, 13), true)
}

func test{{.UType}}Stringer(t *testing.T, a X1{{.UType}}Collection, ordered bool) {
	s1 := a.String()
	if ordered && s1 != "[10, 71, 3, 7, 13]" {
		t.Errorf("Got %s for %+v", s1, a)
	} else if len(s1) != 18 {
		t.Errorf("Got %s for %+v", s1, a)
	}

	s2 := a.MkString("|")
	if ordered && s2 != "10|71|3|7|13" {
		t.Errorf("Got %s for %+v", s2, a)
	} else if len(s2) != 12 {
		t.Errorf("Got %s for %+v", s2, a)
	}

	s3 := a.MkString3("<", "|", ">")
	if ordered && s3 != "<10|71|3|7|13>" {
		t.Errorf("Got %s for %+v", s3, a)
	} else if len(s3) != 14 {
		t.Errorf("Got %s for %+v", s3, a)
	}

	sl := a.StringList()
	if len(sl) != 5 {
		t.Errorf("Expected 5 but got %d", len(sl))
	}

	sort.Sort(sort.StringSlice(sl))
	if sl[0] != "10" {
		t.Errorf("Got %s", sl[0])
	}
	if sl[1] != "13" {
		t.Errorf("Got %s", sl[1])
	}
	if sl[2] != "3" {
		t.Errorf("Got %s", sl[2])
	}
	if sl[3] != "7" {
		t.Errorf("Got %s", sl[3])
	}
	if sl[4] != "71" {
		t.Errorf("Got %s", sl[4])
	}

	var m json.Marshaler = a
	json, err := m.MarshalJSON()
	if err != nil {
		t.Errorf("Got %v", err)
	}
	if ordered && string(json) != `["10", "71", "3", "7", "13"]` {
		t.Errorf("Got %s for %+v", json, a)
	} else if len(json) != 28 {
		t.Errorf("Got %s for %+v", json, a)
	}
}

