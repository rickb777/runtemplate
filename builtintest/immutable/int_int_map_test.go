// Generated from ../map_test.tpl with Type=int
// options: Mutable:<no value> M:.m

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

func TestIntMapFilter(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 1}, TX1IntIntTuple{1, 2}, TX1IntIntTuple{2, 3})

	b := a.Filter(func(k, v int) bool {
		return v > 2
	})

	exp := NewTX1IntIntMap(TX1IntIntTuple{2, 3})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp, b)
	}
}

func TestIntMapPartition(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 4}, TX1IntIntTuple{2, 11}, TX1IntIntTuple{4, 0})

	b, c := a.Partition(func(k, v int) bool {
		return v > 5
	})

	exp1 := NewTX1IntIntMap(TX1IntIntTuple{2, 11})
	if !b.Equals(exp1) {
		t.Errorf("Expected '%+v' but got '%+v'", exp1.m, b.m)
	}

	exp2 := NewTX1IntIntMap(TX1IntIntTuple{8, 4}, TX1IntIntTuple{4, 0})
	if !c.Equals(exp2) {
		t.Errorf("Expected '%+v' but got '%+v'", exp2.m, c.m)
	}
}

func TestIntMapTransform(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 6}, TX1IntIntTuple{9, 10}, TX1IntIntTuple{10, 5})

	b := a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})

	exp := NewTX1IntIntMap(TX1IntIntTuple{9, 36}, TX1IntIntTuple{10, 100}, TX1IntIntTuple{11, 25})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp.m, b.m)
	}
}

func TestIntMapFlatMap(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{2, 6}, TX1IntIntTuple{1, 10}, TX1IntIntTuple{10, 5})

	b := a.FlatMap(func(k int, v int) []TX1IntIntTuple {
	    if k > 3 {
	        return nil
	    }
		return []TX1IntIntTuple{
		    {k-1, v+1},
		    {k+1, v+2},
		}
	})

	exp := NewTX1IntIntMap(TX1IntIntTuple{1, 7}, TX1IntIntTuple{3, 8},
	    TX1IntIntTuple{0, 11}, TX1IntIntTuple{2, 12})
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp.m, b.m)
	}
}



func TestIntMapMkString(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 4}, TX1IntIntTuple{4, 0})

	c := a.MkString("|")

	if c != "8:4|4:0" && c != "4:0|8:4" {
		t.Errorf("Expected '8:4|4:0' but got %q", c)
	}
}

func TestIntMapMkString3(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 4}, TX1IntIntTuple{4, 0})

	c := a.MkString3("<", ",", ">")

	if c != "<8:4,4:0>" && c != "<4:0,8:4>" {
		t.Errorf("Expected '<8:4,4:0>' but got %q", c)
	}
}
