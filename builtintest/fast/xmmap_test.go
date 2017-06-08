package fast

import (
	"testing"
)

func TestTxRemove(t *testing.T) {
	a := NewTXIntIntMap1(3, 1)

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

func TestTxContainsKey(t *testing.T) {
	a := NewTXIntIntMap1(13, 1)

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

func TestTxClear(t *testing.T) {
	a := NewTXIntIntMap1(2, 5)

	a.Clear()

	if a.Size() != 0 {
		t.Errorf("%+v", a)
	}
}

func TestTxClone(t *testing.T) {
	a1 := NewTXIntIntMap(TXIntIntTuples{}.Append2(1, 9, 2, 8)...)
	a2 := NewTXIntIntMap(TXIntIntTuples{}.Append3(1, 9, 2, 8, 3, 3)...)

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

//func TestTxSend(t *testing.T) {
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
