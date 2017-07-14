// Generated from ../collection_test.tpl with Type=int
// options: Append:<no value>

package immutable

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

func TestNewXIntCollection(t *testing.T) {
	testEmptyIntCollection(t, NewX1IntSet(), "Set")
	testEmptyIntCollection(t, NewX1IntList(), "List")
}

func testEmptyIntCollection(t *testing.T, a Sizer, kind string) {
	if a.Size() != 0 {
		t.Errorf("%s: Expected 0 but got %d", kind, a.Size())
	}
	if !a.IsEmpty() {
		t.Errorf("%s: Expected empty", kind)
	}
	if a.NonEmpty() {
		t.Errorf("%s: Expected empty", kind)
	}
}

func TestIntToSlice(t *testing.T) {
	testIntToSlice(t, NewX1IntSet(1, 2, 3), "Set")
	testIntToSlice(t, NewX1IntList(1, 2, 3), "List")
}

func testIntToSlice(t *testing.T, a X1IntCollection, kind string) {
	s := a.ToSlice()

	if a.Size() != 3 {
		t.Errorf("%s: Expected 3 but got %d", kind, a.Size())
	}

	if len(s) != 3 {
		t.Errorf("%s: Expected 3 but got %d", kind, len(s))
	}
}

func TestIntExists(t *testing.T) {
	testIntExists1(t, NewX1IntSet(1, 2, 3), "Set")
	testIntExists1(t, NewX1IntList(1, 2, 3), "List")
}

func testIntExists1(t *testing.T, a X1IntCollection, kind string) {
	has2 := a.Exists(func(v int) bool {
		return v > 2
	})

	if !has2 {
		t.Errorf("%s: Expected exists for %+v", kind, a)
	}

	has5 := a.Exists(func(v int) bool {
		return v > 5
	})

	if has5 {
		t.Errorf("%s: Expected not exists for %+v", kind, a)
	}
}

func TestIntForall(t *testing.T) {
	testIntForall(t, NewX1IntSet(1, 2, 3), "Set")
	testIntForall(t, NewX1IntList(1, 2, 3), "List")
}

func testIntForall(t *testing.T, a X1IntCollection, kind string) {
	has1 := a.Forall(func(v int) bool {
		return v >= 1
	})

	if !has1 {
		t.Errorf("%s: Expected forall for %+v", kind, a)
	}

	has2 := a.Forall(func(v int) bool {
		return v >= 2
	})

	if has2 {
		t.Errorf("%s: Expected not forall for %+v", kind, a)
	}
}

func TestIntFind(t *testing.T) {
	testIntFind(t, NewX1IntSet(1, 2, 3), "Set")
	testIntFind(t, NewX1IntList(1, 2, 3, 4), "List")
}

func testIntFind(t *testing.T, a X1IntCollection, kind string) {
	b, e := a.Find(func(v int) bool {
		return v > 2
	})

	if !e {
		t.Errorf("%s: Expected '3' but got '%+v'", kind, b)
	}

	if b != 3 {
		t.Errorf("%s: Expected '3' but got '%+v'", kind, b)
	}

	c, e := a.Find(func(v int) bool {
		return v > 20
	})

	if e {
		t.Errorf("%s: Expected false but got '%+v'", kind, c)
	}
}


func TestIntCountBy(t *testing.T) {
	testIntCountBy(t, NewX1IntSet(1, 2, 3), "Set")
	testIntCountBy(t, NewX1IntList(1, 2, 3), "List")
}

func testIntCountBy(t *testing.T, a X1IntCollection, kind string) {
	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	if n != 2 {
		t.Errorf("%s: Expected 2 but got %d", kind, n)
	}
}

func TestIntForeach(t *testing.T) {
	testIntForeach(t, NewX1IntSet(1, 2, 3), "Set")
	testIntForeach(t, NewX1IntList(1, 2, 3), "List")
}

func testIntForeach(t *testing.T, a X1IntCollection, kind string) {
	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	if sum1 != 6 {
		t.Errorf("%s: Expected 6 but got %d", kind, sum1)
	}
}

func TestIntContains(t *testing.T) {
	testIntContains(t, NewX1IntSet(71, 1, 7, 13), "Set")
	testIntContains(t, NewX1IntList(71, 1, 7, 13), "List")
}

func testIntContains(t *testing.T, a X1IntCollection, kind string) {
	if !a.Contains(71) {
		t.Errorf("%s: should contain 71", kind)
	}

	if a.Contains(72) {
		t.Errorf("%s: should not contain 72", kind)
	}

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Errorf("%s: should contain 13, 7, 1", kind)
	}

	if !a.ContainsAll(1, 7, 13) {
		t.Errorf("%s: should contain all 1, 7, 13", kind)
	}

	if a.ContainsAll(1, 3, 5, 7, 9, 11, 13) {
		t.Errorf("%s: should not contain all 1, 3, 5, 7, 9, 11, 13", kind)
	}
}

func TestIntMinMaxSum(t *testing.T) {
	testIntMinMaxSum(t, NewX1IntSet(10, 71, 3, 7, 13), "Set")
	testIntMinMaxSum(t, NewX1IntList(10, 71, 3, 7, 13), "List")
}

func testIntMinMaxSum(t *testing.T, a X1IntCollection, kind string) {
	if a.Min() != 3 {
		t.Errorf("%s: Expected 3 but got %d", kind, a.Min())
	}
	if a.Max() != 71 {
		t.Errorf("%s: Expected 71 but got %d", kind, a.Max())
	}
	if a.Sum() != 104 {
		t.Errorf("%s: Expected 104 but got %d", kind, a.Sum())
	}
}

func TestIntStringer(t *testing.T) {
	testIntStringer(t, NewX1IntSet(10, 71, 3, 7, 13), false, "Set")
	testIntStringer(t, NewX1IntList(10, 71, 3, 7, 13), true, "List")
}

func testIntStringer(t *testing.T, a X1IntCollection, ordered bool, kind string) {
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
		t.Errorf("%s: Expected 5 but got %d", kind, len(sl))
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

