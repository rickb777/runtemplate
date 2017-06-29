// Generated from ../map_test.tpl with Type=int
// options: Mutable:true

package threadsafe

import (
	"testing"
)

func TestImIntIntMapToSlice(t *testing.T) {
	a := NewTX1IntIntMap1(1, 2)
	s := a.ToSlice()

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	if len(s) != 1 {
		t.Errorf("Expected 1 but got %d", len(s))
	}
}

func TestImIntIntMapContainsAllKeys(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 6}, TX1IntIntTuple{1, 10}, TX1IntIntTuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("%+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("%+v", a)
	}
}

func TestImIntIntMapCardinality(t *testing.T) {
	a1 := NewTX1IntIntMap()
	a2 := NewTX1IntIntMap1(1, 2)

	if a1.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a1.Size())
	}

	if a2.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a2.Size())
	}
}

func TestImIntIntMapEquals(t *testing.T) {
	a1 := NewTX1IntIntMap()
	b1 := NewTX1IntIntMap()
	a2 := NewTX1IntIntMap(TX1IntIntTuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTX1IntIntMap(TX1IntIntTuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTX1IntIntMap(TX1IntIntTuples{}.Append3(8, 19, 10, 4, 3, 1)...)

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

//func TestImIntIntMapSend(t *testing.T) {
//	a := NewTX1IntIntMap(1, 2, 3, 4)
//
//	b := NewTX1IntIntMap()
//	for val := range a.Send() {
//		b.Add(val)
//	}
//
//	if !a.Equals(b) {
//		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
//	}
//}



func TestMuIntIntMapRemove(t *testing.T) {
	a := NewTX1IntIntMap1(3, 1)

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

func TestMuIntIntMapContainsKey(t *testing.T) {
	a := NewTX1IntIntMap1(13, 1)

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

func TestMuIntIntMapClear(t *testing.T) {
	a := NewTX1IntIntMap1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("%+v", a)
	}
}

func TestMuIntIntMapClone(t *testing.T) {
	a1 := NewTX1IntIntMap(TX1IntIntTuples{}.Append2(1, 9, 2, 8)...)
	a2 := NewTX1IntIntMap(TX1IntIntTuples{}.Append3(1, 9, 2, 8, 3, 3)...)

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

