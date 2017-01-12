package threadsafe

import (
	"testing"
	"sort"
	"encoding/json"
)

func TestNewXCollection(t *testing.T) {
	testNewXCollection1(t, NewXInt32Set())
	testNewXCollection1(t, NewXInt32List())
	testNewXCollection2(t, NewXAppleList())
	testNewXCollection2(t, NewXAppleSet())
}

func testNewXCollection1(t *testing.T, a XInt32Collection) {
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

func testNewXCollection2(t *testing.T, a XAppleCollection) {
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

func TestToSlice(t *testing.T) {
	testToSlice(t, NewXInt32Set(1, 2, 3))
	testToSlice(t, NewXInt32List(1, 2, 3))
}

func testToSlice(t *testing.T, a XInt32Collection) {
	s := a.ToSlice()

	if len(s) != 3 {
		t.Errorf("Expected 3 but got %d", len(s))
	}
}

func TestExists(t *testing.T) {
	testExists(t, NewXInt32Set(1, 2, 3))
	testExists(t, NewXInt32List(1, 2, 3))
}

func testExists(t *testing.T, a XInt32Collection) {
	has2 := a.Exists(func(v int32) bool {
		return v > 2
	})

	if !has2 {
		t.Errorf("Expected exists for %+v", a)
	}

	has5 := a.Exists(func(v int32) bool {
		return v > 5
	})

	if has5 {
		t.Errorf("Expected not exists for %+v", a)
	}
}

func TestForall(t *testing.T) {
	testForall(t, NewXInt32Set(1, 2, 3))
	testForall(t, NewXInt32List(1, 2, 3))
}

func testForall(t *testing.T, a XInt32Collection) {
	has1 := a.Forall(func(v int32) bool {
		return v >= 1
	})

	if !has1 {
		t.Errorf("Expected forall for %+v", a)
	}

	has2 := a.Forall(func(v int32) bool {
		return v >= 2
	})

	if has2 {
		t.Errorf("Expected not forall for %+v", a)
	}
}

func TestCountBy(t *testing.T) {
	testCountBy(t, NewXInt32Set(1, 2, 3))
	testCountBy(t, NewXInt32List(1, 2, 3))
}

func testCountBy(t *testing.T, a XInt32Collection) {
	n := a.CountBy(func(v int32) bool {
		return v >= 2
	})

	if n != 2 {
		t.Errorf("Expected 2 but got %d", n)
	}
}

func TestForeach(t *testing.T) {
	testForeach(t, NewXInt32Set(1, 2, 3))
	testForeach(t, NewXInt32List(1, 2, 3))
}

func testForeach(t *testing.T, a XInt32Collection) {
	sum1 := int32(0)
	a.Foreach(func(v int32) {
		sum1 += v
	})

	if sum1 != 6 {
		t.Errorf("Expected 6 but got %d", sum1)
	}
}

func TestContains(t *testing.T) {
	testContains(t, NewXInt32Set(71, 1, 7, 13))
	testContains(t, NewXInt32List(71, 1, 7, 13))
}

func testContains(t *testing.T, a XInt32Collection) {
	if !a.Contains(71) {
		t.Error("should contain 71")
	}

	if a.Contains(72) {
		t.Error("should not contain 72")
	}

	if !(a.Contains(13) && a.Contains(7) && a.Contains(1)) {
		t.Error("should contain 13, 7, 1")
	}
}

func TestMinMaxSum(t *testing.T) {
	testMinMaxSum(t, NewXInt32Set(10, 71, 3, 7, 13))
	testMinMaxSum(t, NewXInt32List(10, 71, 3, 7, 13))
}

func testMinMaxSum(t *testing.T, a XInt32Collection) {
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

func TestStringer(t *testing.T) {
	testStringer(t, NewXInt32Set(10, 71, 3, 7, 13), false)
	testStringer(t, NewXInt32List(10, 71, 3, 7, 13), true)
}

func testStringer(t *testing.T, a XInt32Collection, ordered bool) {
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

