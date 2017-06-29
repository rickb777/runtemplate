// Generated from ../map_test.tpl with Type=int
// options: Mutable:<no value>

package immutable

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

