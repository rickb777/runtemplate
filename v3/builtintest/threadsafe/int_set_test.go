// Generated from ../set_test.tpl with Type=int
// options: Mutable:true M:.slice()

package threadsafe

import (
    "bytes"
    "encoding/gob"
    "encoding/json"
	"fmt"
	"testing"
	. "github.com/onsi/gomega"
)

func TestNewIntSet_withEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3)
	b := NewX1IntSet(2, 3, 1)

	g.Expect(a.Size()).To(Equal(3))

	g.Expect(a.IsSet()).To(BeTrue())
	g.Expect(a.IsSequence()).To(BeFalse())
	g.Expect(a.Equals(b)).To(BeTrue())
	g.Expect(a.Equals(nil)).To(BeFalse())
	g.Expect(a.Equals(NewX1IntSet(2))).To(BeFalse())
	g.Expect(a.Equals(NewX1IntSet(1, 2, 4))).To(BeFalse())

	a = nil
	g.Expect(a.Equals(b)).To(BeFalse())
	g.Expect(a.Equals(nil)).To(BeTrue())
}

func TestNewIntSetNoDuplicate(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(7, 5, 3, 7)

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(a.Contains(7)).To(BeTrue())
	g.Expect(a.Contains(5)).To(BeTrue())
	g.Expect(a.Contains(3)).To(BeTrue())
}

func TestConvertIntSet(t *testing.T) {
	g := NewGomegaWithT(t)

	a, ok := ConvertX1IntSet(1, 5.1, uint8(2), 7, 3)

	g.Expect(ok).To(BeTrue())
	g.Expect(a.Equals(NewX1IntSet(1, 5, 2, 7, 3))).To(BeTrue())

    b, ok := ConvertX1IntSet(a.ToInterfaceSlice()...)

	g.Expect(ok).To(BeTrue())
	g.Expect(a.Equals(b)).To(BeTrue())
}

func TestMutableIntSetCardinality(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet()

	g.Expect(a.Size()).To(Equal(0))

	a.Add(1)

	g.Expect(a.Size()).To(Equal(1))
	g.Expect(a.Cardinality()).To(Equal(1))

	a.Remove(1)

	g.Expect(a.Size()).To(Equal(0))

	a.Add(9)

	g.Expect(a.Size()).To(Equal(1))
}

func TestMutableIntSetRemove(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(6, 3, 1)

	a.Remove(3)

	g.Expect(a.Size()).To(Equal(2))
	g.Expect(a.Contains(6) && a.Contains(1)).To(BeTrue())

	a.Remove(6)
	a.Remove(1)

	g.Expect(a.Size()).To(Equal(0))
}

func TestIntSetContainsAll(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(8, 6, 7, 5, 3, 0, 9)

	g.Expect(a.ContainsAll(8, 6, 7, 5, 3, 0, 9)).To(BeTrue())
	g.Expect(a.ContainsAll(8, 6, 11, 5, 3, 0, 9)).To(BeFalse())
}

func TestIntSetIsSubset(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 5, 7)
	b := NewX1IntSet(3, 5, 7)
	c := NewX1IntSet(3, 5, 7, 72)

	g.Expect(b.IsSubset(a)).To(BeTrue())
	g.Expect(c.IsSubset(a)).To(BeFalse())

    // check correct nil handling
    a = nil
	g.Expect(b.IsSubset(a)).To(BeFalse())
	g.Expect(a.IsSubset(b)).To(BeTrue())
}

func TestIntSetIsSuperSet(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(9, 5, 2, 1, 11)
	b := NewX1IntSet(5, 2, 11)
	c := NewX1IntSet(5, 2, 11, 42)

	g.Expect(a.IsSuperset(b)).To(BeTrue())
	g.Expect(a.IsSuperset(c)).To(BeFalse())

    // check correct nil handling
    a = nil
	g.Expect(b.IsSuperset(a)).To(BeTrue())
	g.Expect(a.IsSuperset(b)).To(BeFalse())
}

func TestIntSetUnion(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet()

	b := NewX1IntSet(1, 2, 3, 4, 5)

	c := a.Union(b)

	g.Expect(c.Equals(NewX1IntSet(1, 2, 3, 4, 5))).To(BeTrue())

	d := NewX1IntSet(10, 14, 0)

	e := c.Union(d)
	g.Expect(e.Equals(NewX1IntSet(1, 2, 3, 4, 5, 10, 14, 0))).To(BeTrue())

	a = NewX1IntSet(14, 3)

	f := a.Union(e)
	g.Expect(f.Equals(NewX1IntSet(1, 2, 3, 4, 5, 10, 14, 0))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Union(b)
	d = b.Union(a)

	g.Expect(c.Equals(b)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func TestIntSetIntersection(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewX1IntSet(1, 3, 5, 7)
	a2 := NewX1IntSet(1, 3, 5, 7, 10)

	b1 := NewX1IntSet(0, 2, 4, 6)
	b2 := NewX1IntSet(2, 4, 6, 10)

	c1 := a1.Intersect(a2)
	c2 := b1.Intersect(a1)

	g.Expect(c1.Equals(NewX1IntSet(1, 3, 5, 7))).To(BeTrue())
	g.Expect(c2.NonEmpty()).To(BeFalse())

	d1 := a1.Intersect(a2)
	d2 := b2.Intersect(a2)

	g.Expect(d1.Equals(NewX1IntSet(1, 3, 5, 7))).To(BeTrue(), d1.String())
	g.Expect(d2.Equals(NewX1IntSet(10))).To(BeTrue(), d2.String())

    // check correct nil handling
    a1 = nil
	c1 = a1.Intersect(b1)
	d1 = b1.Intersect(a1)

	g.Expect(c1.NonEmpty()).To(BeFalse())
	g.Expect(d1.NonEmpty()).To(BeFalse())
}

func TestIntSetDifference(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3)
	b := NewX1IntSet(1, 3, 4, 5, 6, 99)

	c := a.Difference(b)
	d := b.Difference(a)

	g.Expect(c.Equals(NewX1IntSet(2))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	g.Expect(c.Equals(a)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func TestIntSetSymmetricDifference(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 50)
	b := NewX1IntSet(1, 3, 4, 5, 6, 99)

	c := a.SymmetricDifference(b)
	d := b.SymmetricDifference(a)

	g.Expect(c.Equals(NewX1IntSet(2, 4, 5, 6, 50, 99))).To(BeTrue())

    // check correct nil handling
    a = nil
	c = a.Difference(b)
	d = b.Difference(a)

	g.Expect(c.Equals(a)).To(BeTrue())
	g.Expect(d.Equals(b)).To(BeTrue())
}

func TestIntSetEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet()
	b := NewX1IntSet()

	g.Expect(a.Equals(b)).To(BeTrue())

	c := NewX1IntSet(1, 3, 5, 6, 8)
	d := NewX1IntSet(1, 3, 5, 6, 9)

	g.Expect(c.Equals(d)).To(BeFalse())

    // check correct nil handling
    a = nil
    g.Expect(a.Equals(b)).To(BeTrue())
    g.Expect(b.Equals(a)).To(BeTrue())
}

func TestIntSetToList(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)
	b := a.ToList()

	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.Contains(1)).To(BeTrue())
	g.Expect(a.Contains(2)).To(BeTrue())
	g.Expect(a.Contains(3)).To(BeTrue())
	g.Expect(a.Contains(4)).To(BeTrue())

    c := a.ToSet()
	g.Expect(c.Size()).To(Equal(4))
	g.Expect(c.Contains(1)).To(BeTrue())
	g.Expect(c.Contains(2)).To(BeTrue())
	g.Expect(c.Contains(3)).To(BeTrue())
	g.Expect(c.Contains(4)).To(BeTrue())

	a = nil
	b = a.ToList()

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntSetToSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b1 := a.ToSlice()
	b2 := a.ToInterfaceSlice()

	g.Expect(b1).To(Equal([]int{1, 2, 3, 4}))
	g.Expect(b2).To(Equal([]interface{}{1, 2, 3, 4}))

	a = nil
	b1 = a.ToSlice()
	b2 = a.ToInterfaceSlice()

	g.Expect(b1).To(HaveLen(0))
	g.Expect(b2).To(HaveLen(0))
}

func TestIntSetSend(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)
	b := BuildX1IntSetFromChan(a.Send())

	g.Expect(a.Equals(b)).To(BeTrue())

    // check correct nil handling
	a = nil
	b = BuildX1IntSetFromChan(a.Send())

	g.Expect(a.Equals(b)).To(BeTrue())
}

func TestIntSetForall(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)
	found := a.Forall(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeTrue())

	found = a.Forall(func(v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	found = a.Forall(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeTrue())
}

func TestIntSetExists(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)
	found := a.Exists(func(v int) bool {
		return v > 2
	})

	g.Expect(found).To(BeTrue())

	found = a.Exists(func(v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	found = a.Exists(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeFalse())
}

func TestIntSetForeach(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)
	s := 0

	a.Foreach(func(v int) {
		s += v
	})

	g.Expect(s).To(Equal(10))

    // check correct nil handling
    a = nil
	a.Foreach(func(v int) {
		s += v
	})

	g.Expect(s).To(Equal(10))
}

func TestIntSetFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1IntSet(3, 4))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Filter(func(v int) bool {
		return v > 2
	})
}

func TestIntSetPartition(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1IntSet(3, 4))).To(BeTrue())
	g.Expect(c.Equals(NewX1IntSet(1, 2))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Partition(func(v int) bool {
		return v > 2
	})
}

func TestIntSetTransform(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)

	b := a.Map(func(v int) int {
		return v * v
	})

	g.Expect(b.Equals(NewX1IntSet(1, 4, 9, 16))).To(BeTrue())

    // check correct nil handling
	a = nil
	a.Map(func(v int) int {
		return v * v
	})
}

func TestIntSetFlatMap(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)

	b := a.FlatMap(func(v int) []int {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})

    exp := NewX1IntSet(2, 3, 4, 6, 6, 9)
	g.Expect(b.Equals(exp)).To(BeTrue())

    // check correct nil handling
	a = nil
	a.FlatMap(func(v int) []int {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})
}

func TestIntSetStringMap(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2, 3, 4)

	b := a.StringMap()

	for _, c := range a.ToSlice() {
		s := fmt.Sprintf("%d", c)
		_, ok := b[s]
		g.Expect(ok).To(BeTrue())
	}

    // check correct nil handling
	a = nil
	a.StringMap()
}

func TestMutableIntSetClear(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(2, 5, 9, 10)

	a.Clear()

	g.Expect(a.Size()).To(Equal(0))

    // check correct nil handling
	a = nil
	a.Clear()
}

func TestMutableIntSetClone(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(1, 2)

	b := a.Clone()

	g.Expect(a.Equals(b)).To(BeTrue())

	a.Add(3)
	g.Expect(a.Equals(b)).To(BeFalse())

	c := a.Clone()
	c.Remove(1)

	g.Expect(a.Equals(c)).To(BeFalse())

    // check correct nil handling
	a = nil
	a.Clone()
}

func TestIntSetMkString(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(13, 4)

	c := a.MkString("|")

	g.Expect(c).To(Or(Equal("13|4"), Equal("4|13")))

    // check correct nil handling
	a = nil
	a.MkString("|")
}

func TestIntSetMkString3(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(13, 4)

	c := a.MkString3("<", ", ", ">")

	g.Expect(c).To(Or(Equal("<13, 4>"), Equal("<4, 13>")))

    // check correct nil handling
	a = nil
	a.MkString3("<", ",", ">")
}


func TestIntSetGobEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(13, 4, 7, -2, 9)
	b := NewX1IntSet()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

	g.Expect(err).NotTo(HaveOccurred())

    err = gob.NewDecoder(buf).Decode(&b)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(a.Equals(b)).To(BeTrue())
}


func TestIntSetJsonEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntSet(13, 4, 7, -2, 9)
	b := NewX1IntSet()

    buf, err := json.Marshal(a)

	g.Expect(err).NotTo(HaveOccurred())

    err = json.Unmarshal(buf, &b)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(a.Equals(b)).To(BeTrue())
}
