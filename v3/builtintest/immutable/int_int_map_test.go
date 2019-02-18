// Generated from ../map_test.tpl with Type=int
// options: Mutable:<no value> Immutable:true M:.slice()

package immutable

import (
    "bytes"
    "encoding/gob"
	"sort"
	"testing"
	. "github.com/onsi/gomega"
)

func TestImIntIntMapGet(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewTX1IntIntMap1(1, 2)
	v, found := a.Get(1)

	g.Expect(found).To(BeTrue())
	g.Expect(v).To(Equal(2))

	_, found = a.Get(7)

	g.Expect(found).To(BeFalse())
}

func TestImIntIntMapToSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewTX1IntIntMap1(1, 2)
	s := a.ToSlice()

	g.Expect(a.Size()).To(Equal(1))
	g.Expect(len(s)).To(Equal(1))

    // check correct nil handling
    a = nil
    a.ToSlice()
}

func TestImIntIntMapSize(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewTX1IntIntMap()
	a2 := NewTX1IntIntMap1(1, 2)

	g.Expect(a1.Size()).To(Equal(0))
	g.Expect(a2.Size()).To(Equal(1))

    // check correct nil handling
    a1 = nil
    a1.Size()
}

func TestIntIntKeys(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1IntIntMap(TX1IntIntZip(8, 4, 2).Values(4, 0, 5)...)

    k := a.Keys()
    sort.Ints(k)
	g.Expect(k).To(Equal([]int{2, 4, 8}))

    // check correct nil handling
    a = nil
    a.Keys()
}

func TestIntIntValues(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1IntIntMap(TX1IntIntZip(8, 4, 2).Values(4, 0, 5)...)

    v := a.Values()
    sort.Ints(v)
	g.Expect(v).To(Equal([]int{0, 4, 5}))

    // check correct nil handling
    a = nil
    a.Values()
}

func TestIntIntPut(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1IntIntMap(TX1IntIntZip(8, 4, 2).Values(4, 0, 5)...)

    b := a.Put(7, 9)
    g.Expect(b.Size()).To(Equal(4))
    g.Expect(b.ContainsKey(7)).To(BeTrue())

    // check correct nil handling
    a = nil
    c := a.Put(7, 9)
    g.Expect(c.Size()).To(Equal(1))
}

func TestImIntIntMapContainsAllKeys(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 6}, TX1IntIntTuple{1, 10}, TX1IntIntTuple{2, 11})

	if !a.ContainsAllKeys(8, 1, 2) {
		t.Errorf("Got %+v", a)
	}

	if a.ContainsAllKeys(8, 6, 11, 1, 2) {
		t.Errorf("Got %+v", a)
	}

    // check correct nil handling
    a = nil
    a.ContainsAllKeys()
}

func TestImIntIntMapEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewTX1IntIntMap()
	b1 := NewTX1IntIntMap()
	a2 := NewTX1IntIntMap(TX1IntIntTuples{}.Append2(10, 4, 8, 19)...)
	a3 := NewTX1IntIntMap(TX1IntIntTuples{}.Append3(10, 4, 3, 1, 8, 19)...)
	b3 := NewTX1IntIntMap(TX1IntIntTuples{}.Append3(8, 19, 10, 4, 3, 1)...)

	g.Expect(a1.Equals(b1)).To(BeTrue())
	g.Expect(b1.Equals(a1)).To(BeTrue())
	g.Expect(a2.Equals(b1)).To(BeFalse())
	g.Expect(a2.Equals(b3)).To(BeFalse())
	g.Expect(a3.Equals(a2)).To(BeFalse())
	g.Expect(a3.Equals(b3)).To(BeTrue())
	g.Expect(b3.Equals(a3)).To(BeTrue())

    // check correct nil handling
    a1 = nil
    a1.Equals(b1)
    b1.Equals(a1)
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

func TestIntMapForall(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(1, 8, 2).Values(1, 2, 3)...)

	found := a.Forall(func(k, v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Forall(func(k, v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Forall(func(k, v int) bool {
		return v > 0
	})

	if !found {
		t.Errorf("Expected to find.")
	}

    // check correct nil handling
    a = nil
	a.Forall(func(k, v int) bool {
		return v > 0
	})
}

func TestIntMapExists(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(1, 8, 2).Values(1, 2, 3)...)

	found := a.Exists(func(k, v int) bool {
		return v > 2
	})

	if !found {
		t.Errorf("Expected to find.")
	}

	found = a.Exists(func(k, v int) bool {
		return v > 100
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	found = a.Exists(func(k, v int) bool {
		return v > 0
	})

	if found {
		t.Errorf("Expected not to find.")
	}

    // check correct nil handling
    a = nil
	a.Exists(func(k, v int) bool {
		return v > 2
	})
}

func TestIntMapForeach(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(1, 8, 2).Values(1, 2, 3)...)
	s := 0

	a.Foreach(func(k, v int) {
		s += v
	})

	if s != 6 {
		t.Errorf("Got %d", s)
	}

    // check correct nil handling
    a = nil
	a.Foreach(func(k, v int) {
		s += v
	})

	if s != 6 {
		t.Errorf("Got %d", s)
	}

    // check correct nil handling
    a = nil
	a.Foreach(func(k, v int) {
		s += v
	})
}

func TestIntMapFind(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1IntIntMap(TX1IntIntZip(1, 8, 2).Values(1, 2, 3)...)

	b, found := a.Find(func(k, v int) bool {
		return v > 2
	})

	exp := TX1IntIntTuple{2, 3}
	g.Expect(found).To(BeTrue())
	g.Expect(b).To(Equal(exp))

	_, found = a.Find(func(k, v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	a.Find(func(k, v int) bool {
		return v > 2
	})
}

func TestIntMapFilter(t *testing.T) {
    g := NewGomegaWithT(t)

	a := NewTX1IntIntMap(TX1IntIntZip(1, 8, 2).Values(1, 2, 3)...)

	b := a.Filter(func(k, v int) bool {
		return v > 2
	})

	exp := NewTX1IntIntMap(TX1IntIntTuple{2, 3})
	g.Expect(b.Equals(exp)).To(BeTrue())

    // check correct nil handling
    a = nil
	a.Filter(func(k, v int) bool {
		return v > 2
	})
}

func TestIntMapPartition(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(8, 2, 4).Values(4, 11, 0)...)

	b, c := a.Partition(func(k, v int) bool {
		return v > 5
	})

	exp1 := NewTX1IntIntMap(TX1IntIntTuple{2, 11})
	if !b.Equals(exp1) {
		t.Errorf("Expected '%+v' but got '%+v'", exp1.slice(), b.slice())
	}

	exp2 := NewTX1IntIntMap(TX1IntIntZip(8, 4).Values(4, 0)...)
	if !c.Equals(exp2) {
		t.Errorf("Expected '%+v' but got '%+v'", exp2.slice(), c.slice())
	}

    // check correct nil handling
    a = nil
	a.Partition(func(k, v int) bool {
		return v > 2
	})
}

func TestIntMapTransform(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(8, 9, 10).Values(6, 10, 5)...)

	b := a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})

	exp := NewTX1IntIntMap(TX1IntIntZip(9, 10, 11).Values(36, 100, 25)...)
	if !b.Equals(exp) {
		t.Errorf("Expected '%+v' but got '%+v'", exp.slice(), b.slice())
	}

    // check correct nil handling
    a = nil
	a.Map(func(k, v int) (int, int) {
		return k + 1, v * v
	})
}

func TestIntMapFlatMap(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(2, 1, 18).Values(6, 10, 5)...)

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
		t.Errorf("Expected '%+v' but got '%+v'", exp.slice(), b.slice())
	}

    // check correct nil handling
    a = nil
	a.FlatMap(func(k int, v int) []TX1IntIntTuple {
        return nil
	})
}

func TestIntMapMkString(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(8, 4).Values(4, 0)...)

	c := a.MkString("|")

	if c != "8:4|4:0" && c != "4:0|8:4" {
		t.Errorf("Expected '8:4|4:0' but got %q", c)
	}

    // check correct nil handling
    a = nil
	a.MkString("|")
}

func TestIntMapMkString3(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntTuple{8, 4}, TX1IntIntTuple{4, 0})

	c := a.MkString3("<", ",", ">")

	if c != "<8:4,4:0>" && c != "<4:0,8:4>" {
		t.Errorf("Expected '<8:4,4:0>' but got %q", c)
	}

    // check correct nil handling
    a = nil
	a.MkString3("<", ",", ">")
}

func TestIntMapGobEncode(t *testing.T) {
	a := NewTX1IntIntMap(TX1IntIntZip(1, 9, -2, 8, 3, 3).Values(-5, 10, 13, 17, 19, 23)...)
	b := NewTX1IntIntMap()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

	if err != nil {
		t.Errorf("%v", err)
	}

    err = gob.NewDecoder(buf).Decode(&b)

	if err != nil {
		t.Errorf("%v", err)
	}

	if !a.Equals(b) {
		t.Errorf("Expected '%+v' but got '%+v'", a.slice(), b.slice())
	}
}
