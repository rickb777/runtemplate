package immutable

import (
	"testing"
)

func TestImMapToSlice(t *testing.T) {
	a := NewTXIntIntMap1(1, 2)
	s := a.ToSlice()

	if a.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a.Size())
	}

	if len(s) != 1 {
		t.Errorf("Expected 1 but got %d", len(s))
	}
}

func TestImMapContainsAllKeys(t *testing.T) {
	a := NewTXIntIntMap(TXIntIntTuple{8, 6}, TXIntIntTuple{1, 10}, TXIntIntTuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("%+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("%+v", a)
	}
}

func TestImMapCardinality(t *testing.T) {
	a1 := NewTXIntIntMap()
	a2 := NewTXIntIntMap1(1, 2)

	if a1.Size() != 0 {
		t.Errorf("Expected 0 but got %d", a1.Size())
	}

	if a2.Size() != 1 {
		t.Errorf("Expected 1 but got %d", a2.Size())
	}
}

func TestImMapEquals(t *testing.T) {
	a1 := NewTXIntIntMap()
	b1 := NewTXIntIntMap()
	a2 := NewTXIntIntMap(TXIntIntTuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTXIntIntMap(TXIntIntTuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTXIntIntMap(TXIntIntTuples{}.Append3(8, 19, 10, 4, 3, 1)...)

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

//func TestImMapSend(t *testing.T) {
//	a := NewTXIntIntMap(1, 2, 3, 4)
//
//	b := NewTXIntIntMap()
//	for val := range a.Send() {
//		b.Add(val)
//	}
//
//	if !a.Equals(b) {
//		t.Errorf("Expected '%+v' to equal '%+v'", a, b)
//	}
//}
