package fast

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
	testEmptyIntCollection(t, NewXIntSet())
	testEmptyIntCollection(t, NewXIntList())
}

func testEmptyIntCollection(t *testing.T, a Sizer) {
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

func TestIntToSlice(t *testing.T) {
	testIntToSlice(t, NewXIntSet(1, 2, 3))
	testIntToSlice(t, NewXIntList(1, 2, 3))
}

func testIntToSlice(t *testing.T, a XIntCollection) {
	s := a.ToSlice()

	if a.Size() != 3 {
		t.Errorf("Expected 3 but got %d", a.Size())
	}

	if len(s) != 3 {
		t.Errorf("Expected 3 but got %d", len(s))
	}
}

func TestIntExists(t *testing.T) {
	testIntExists1(t, NewXIntSet(1, 2, 3))
	testIntExists1(t, NewXIntList(1, 2, 3))
}

func testIntExists1(t *testing.T, a XIntCollection) {
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

func TestIntForall(t *testing.T) {
	testIntForall(t, NewXIntSet(1, 2, 3))
	testIntForall(t, NewXIntList(1, 2, 3))
}

func testIntForall(t *testing.T, a XIntCollection) {
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

func TestIntCountBy(t *testing.T) {
	testIntCountBy(t, NewXIntSet(1, 2, 3))
	testIntCountBy(t, NewXIntList(1, 2, 3))
}

func testIntCountBy(t *testing.T, a XIntCollection) {
	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	if n != 2 {
		t.Errorf("Expected 2 but got %d", n)
	}
}

func TestIntForeach(t *testing.T) {
	testIntForeach(t, NewXIntSet(1, 2, 3))
	testIntForeach(t, NewXIntList(1, 2, 3))
}

func testIntForeach(t *testing.T, a XIntCollection) {
	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	if sum1 != 6 {
		t.Errorf("Expected 6 but got %d", sum1)
	}
}

func TestIntContains(t *testing.T) {
	testIntContains(t, NewXIntSet(71, 1, 7, 13))
	testIntContains(t, NewXIntList(71, 1, 7, 13))
}

func testIntContains(t *testing.T, a XIntCollection) {
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

func TestIntMinMaxSum(t *testing.T) {
	testIntMinMaxSum(t, NewXIntSet(10, 71, 3, 7, 13))
	testIntMinMaxSum(t, NewXIntList(10, 71, 3, 7, 13))
}

func testIntMinMaxSum(t *testing.T, a XIntCollection) {
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

func TestIntStringer(t *testing.T) {
	testIntStringer(t, NewXIntSet(10, 71, 3, 7, 13), false)
	testIntStringer(t, NewXIntList(10, 71, 3, 7, 13), true)
}

func testIntStringer(t *testing.T, a XIntCollection, ordered bool) {
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

