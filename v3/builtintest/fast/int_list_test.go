// Generated from ../list_test.tpl with Type=int
// options: Append:true Find:<no value> Mutable:true M:.slice()

package fast

import (
    "bytes"
    "encoding/gob"
    "encoding/json"
	"testing"
	. "github.com/onsi/gomega"
)

func intRangeOf(from, to int) []int {
	n := 1 + to - from
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		a[i] = i + from
	}
	return a
}

func TestNewIntList_withEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3)
	b := NewX1IntList(1, 3, 2)

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(a.Len()).To(Equal(3))
	g.Expect(a.Get(1)).To(Equal(2))
	g.Expect(a.IsSet()).To(BeFalse())
	g.Expect(a.IsSequence()).To(BeTrue())
	g.Expect(a.Equals(a)).To(BeTrue())
	g.Expect(a.Equals(b)).To(BeFalse())
	g.Expect(a.Equals(nil)).To(BeFalse())
	g.Expect(a.Equals(NewX1IntList(1, 2))).To(BeFalse())

	a = nil
	g.Expect(a.Equals(b)).To(BeFalse())
}

func TestConvertIntList(t *testing.T) {
	g := NewGomegaWithT(t)

	a, ok := ConvertX1IntList(1, 5.1, uint8(2), 7, 3)

	g.Expect(ok).To(BeTrue())

	g.Expect(a.Equals(NewX1IntList(1, 5, 2, 7, 3))).To(BeTrue(), "%v", a)

    b, ok := ConvertX1IntList(a.ToInterfaceSlice()...)

	g.Expect(ok).To(BeTrue())
	g.Expect(a).To(Equal(b))
}

func TestIntListAppend(t *testing.T) {
	g := NewGomegaWithT(t)

	var a *X1IntList

	b := a.Append(1, 2, 3).Append(4, 5).Append(6, 7)

	// a is not mutated when it is nil
	g.Expect(a.Size()).To(Equal(0))

	g.Expect(b.Size()).To(Equal(7))
	g.Expect(b.Get(3)).To(Equal(4))
	g.Expect(b.Last()).To(Equal(7))
}

func TestIntListDoInsertAt(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        i int
        act *X1IntList
        more []int
        exp *X1IntList
    }{
        {
            0,
            nil,
            []int{10, 11},
            NewX1IntList(10, 11),
        },
        {
            0,
            nil,
            nil,
            nil,
        },
        {
            0,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            []int{10, 11},
            NewX1IntList(10, 11, 1, 2, 3, 4, 5, 6),
        },
        {
            2,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            []int{10, 11, 12},
            NewX1IntList(1, 2, 10, 11, 12, 3, 4, 5, 6),
        },
        {
            6,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            []int{10, 11},
            NewX1IntList(1, 2, 3, 4, 5, 6, 10, 11),
        },
        {
            3,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            []int{},
            NewX1IntList(1, 2, 3, 4, 5, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoInsertAt(c.i, c.more...)
        g.Expect(r.Equals(c.exp)).To(BeTrue(), "%d %v", i, r)
    }
}

func TestIntListDoDeleteAt(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        i, n int
        act, exp *X1IntList
    }{
        {
            0, 2,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(3, 4, 5, 6),
        },
        {
            2, 2,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 5, 6),
        },
        {
            4, 2,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3, 4),
        },
        {
            3, 0,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3, 4, 5, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteAt(c.i, c.n)

        g.Expect(c.act.Equals(c.exp)).To(BeTrue(), "%d %v", i, c.act)
        g.Expect(c.act.Equals(r)).To(BeTrue(), "%d %v", i, r)
    }
}

func TestIntListDoDeleteFirst(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        n int
        act, exp *X1IntList
    }{
        {
            0,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3, 4, 5, 6),
        },
        {
            1,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(2, 3, 4, 5, 6),
        },
        {
            3,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(4, 5, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteFirst(c.n)

        g.Expect(c.act.Equals(c.exp)).To(BeTrue(), "%d %v", i, c.act)
        g.Expect(c.act.Equals(r)).To(BeTrue(), "%d %v", i, r)
    }
}

func TestIntListDoDeleteLast(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        n int
        act, exp *X1IntList
    }{
        {
            0,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3, 4, 5, 6),
        },
        {
            1,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3, 4, 5),
        },
        {
            3,
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(1, 2, 3),
        },
    }

    for i, c := range cases {
        r := c.act.DoDeleteLast(c.n)

        g.Expect(c.act.Equals(c.exp)).To(BeTrue(), "%d %v", i, c.act)
        g.Expect(c.act.Equals(r)).To(BeTrue(), "%d %v", i, r)
    }
}

func TestIntListDoKeepWhere(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        act, exp *X1IntList
    }{
        {
            nil,
            nil,
        },
        {
            NewX1IntList(1, 2, 3, 4, 5, 6),
            NewX1IntList(2, 4, 6),
        },
    }

    for i, c := range cases {
        r := c.act.DoKeepWhere(func (v int) bool {
            return v % 2 == 0
        })

        g.Expect(c.act.Equals(c.exp)).To(BeTrue(), "%d %v", i, c.act)
        g.Expect(c.act.Equals(r)).To(BeTrue(), "%d %v", i, r)
    }
}

func TestIntListClone(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b := a.Clone()

	g.Expect(a).To(Equal(b))

	c := a.Clone().Tail()

	g.Expect(a.slice()).NotTo(Equal(c.slice()))

	a = nil
	b = a.Clone()

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListToSet(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b := a.ToSet()

	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.Contains(1)).To(BeTrue())
	g.Expect(a.Contains(2)).To(BeTrue())
	g.Expect(a.Contains(3)).To(BeTrue())
	g.Expect(a.Contains(4)).To(BeTrue())

    c := a.ToList()
	g.Expect(c.slice()).To(Equal(a.slice()))

	a = nil
	b = a.ToSet()

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListToSlice(t *testing.T) {
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

func TestIntListSend(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b := BuildX1IntListFromChan(a.Send())

	g.Expect(a.slice()).To(Equal(b.slice()))

    // check correct nil handling
	a = nil
	b = BuildX1IntListFromChan(a.Send())

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListGetHeadTailLastInit(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)

	g.Expect(a.Get(0)).To(Equal(1))
	g.Expect(a.Get(3)).To(Equal(4))
	g.Expect(a.Head()).To(Equal(1))
	g.Expect(a.Last()).To(Equal(4))
	g.Expect(a.HeadOption()).To(Equal(1))
	g.Expect(a.LastOption()).To(Equal(4))

	tail := a.Tail()

	g.Expect(tail.Equals(NewX1IntList(2, 3, 4))).To(BeTrue())

	init := a.Init()

	g.Expect(init.Equals(NewX1IntList(1, 2, 3))).To(BeTrue())

    // check correct nil handling
    a = nil

	g.Expect(a.HeadOption()).To(Equal(0))
	g.Expect(a.LastOption()).To(Equal(0))
}

func TestIntListContains(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)

	found := a.Contains(3)

	g.Expect(found).To(BeTrue())

	found = a.Contains(5)

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
    a.Contains(3)
}

func TestIntListContainsAll(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)

	found := a.ContainsAll(3)

	g.Expect(found).To(BeTrue())

	found = a.ContainsAll(1, 3, 5, 7)

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
    a.ContainsAll(3)
}

func TestIntListFind(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b, found := a.Find(func(v int) bool {
		return v > 2
	})

	g.Expect(found).To(BeTrue())

	g.Expect(b).To(Equal(3))

	b, found = a.Find(func(v int) bool {
		return v > 100
	})

	g.Expect(found).To(BeFalse())

    // check correct nil handling
    a = nil
	_, found = a.Find(func(v int) bool {
		return v > 0
	})

	g.Expect(found).To(BeFalse())
}

func TestIntListForall(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
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

func TestIntListExists(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
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

func TestIntListForeach(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
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

func TestIntListFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b := a.Filter(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1IntList(3, 4))).To(BeTrue(), "%v", b)

	b = a.Filter(func(v int) bool {
		return v > 100
	})

	g.Expect(b.IsEmpty()).To(BeTrue())

    a = nil
	b = a.Filter(func(v int) bool {
		return v > 0
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListPartition(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1IntList(3, 4))).To(BeTrue(), "%v", b)
	g.Expect(c.Equals(NewX1IntList(1, 2))).To(BeTrue(), "%v", c)

	a = nil
	b, c = a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
	g.Expect(c.IsEmpty()).To(BeTrue())
}

func TestIntListTransform(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4)
	b := a.Map(func(v int) int {
		return v * v
	})

    exp := NewX1IntList(1, 4, 9, 16)
	g.Expect(b.Equals(exp)).To(BeTrue(), "%v %v", b, exp)

	a = nil
	b = a.Map(func(v int) int {
		return v * v
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListFlatMap(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 2, 3, 4, 5)
	b := a.FlatMap(func(v int) []int {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})

    exp := NewX1IntList(2, 3, 4, 6, 6, 9)

	g.Expect(b.Equals(exp)).To(BeTrue(), "%v %v", b, exp)

    // check correct nil handling
	a = nil
	b = a.FlatMap(func(v int) []int {
	    if v > 3 {
	        return nil
	    }
		return []int{v * 2, v * 3}
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func TestIntListSorted(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)
	b := a.Sorted()


	g.Expect(a.Equals(NewX1IntList(-2, 4, 7, 9, 13))).To(BeTrue(), "%v", a)
	g.Expect(b.Equals(NewX1IntList(-2, 4, 7, 9, 13))).To(BeTrue(), "%v", b)

    // check correct nil handling
	a = nil
	a.Sorted()
}

func TestIntListStableSorted(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)
	b := a.StableSorted()


	g.Expect(a.Equals(NewX1IntList(-2, 4, 7, 9, 13))).To(BeTrue(), "%v", a)
	g.Expect(b.Equals(NewX1IntList(-2, 4, 7, 9, 13))).To(BeTrue(), "%v", b)

    // check correct nil handling
	a = nil
	a.StableSorted()
}

func TestIntListReverseOdd(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)

	b := a.Reverse()

	g.Expect(b.Equals(a)).To(BeFalse(), "%v %v", a, b)

	c := b.Reverse()

	g.Expect(c.Equals(a)).To(BeTrue(), "%v %v", a, c)

    // check correct nil handling
	a = nil
	a.Reverse()
}

func TestIntListReverseEven(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9, 17)

	b := a.Reverse()

	g.Expect(b.Equals(a)).To(BeFalse(), "%v %v", a, b)

	c := b.Reverse()

	g.Expect(c.Equals(a)).To(BeTrue(), "%v %v", a, c)
}

func TestIntListDoReverse(t *testing.T) {
	g := NewGomegaWithT(t)

	a1 := NewX1IntList(13, 4, 7, -2, 9, 17)
	a2 := a1.Clone()

	b := a2.DoReverse()

	g.Expect(b.Equals(a1)).To(BeFalse(), "%v %v", a1, b)

	c := b.DoReverse()

	g.Expect(c.Equals(a1)).To(BeTrue(), "%v %v", a1, c)

    // check correct nil handling
	a1 = nil
	a1.DoReverse()
}

func TestIntListShuffle(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.Shuffle()

	g.Expect(b.Equals(a)).To(BeFalse(), "%v %v", a, b)

	// prove that the same set of numbers is present
	c := b.Sorted()

	g.Expect(c.Equals(a)).To(BeTrue(), "%v %v", a, c)

    // check correct nil handling
	a = nil
	a.Shuffle()
}

func TestIntListDoShuffle(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.Clone()
	c := b.DoShuffle()

	g.Expect(b.Equals(a)).To(BeFalse(), "%v %v", a, b)
	g.Expect(b.Equals(c)).To(BeTrue(), "%v %v", c, b)

	// prove that the same set of numbers is present
	d := b.Sorted()

	g.Expect(d.Equals(a)).To(BeTrue(), "%v %v", a, d)

    // check correct nil handling
	a = nil
	a.DoShuffle()
}

func TestIntListTake(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.Take(30)

	g.Expect(b.Size()).To(Equal(30))
	g.Expect(b.Head()).To(Equal(1))
	g.Expect(b.Last()).To(Equal(30))

	c := a.TakeLast(30)

	g.Expect(c.Size()).To(Equal(30))
	g.Expect(c.Head()).To(Equal(71))
	g.Expect(c.Last()).To(Equal(100))

	d := a.Take(101)

	g.Expect(d.Size()).To(Equal(100))
	g.Expect(d.Head()).To(Equal(1))
	g.Expect(d.Last()).To(Equal(100))

	e := a.TakeLast(101)

	g.Expect(e.Size()).To(Equal(100))
	g.Expect(e.Head()).To(Equal(1))
	g.Expect(e.Last()).To(Equal(100))

    // check correct nil handling
	a = nil
	a.Take(0)
	a.TakeLast(0)
}

func TestIntListDrop(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.Drop(70)

	g.Expect(b.Size()).To(Equal(30))
	g.Expect(b.Head()).To(Equal(71))
	g.Expect(b.Last()).To(Equal(100))

	c := a.DropLast(75)

	g.Expect(c.Size()).To(Equal(25))
	g.Expect(c.Head()).To(Equal(1))
	g.Expect(c.Last()).To(Equal(25))

	d := a.Drop(101)

	g.Expect(d.Size()).To(Equal(0))

	e := a.DropLast(101)

	g.Expect(e.Size()).To(Equal(0))

    // check correct nil handling
	a = nil
	a.Drop(0)
	a.DropLast(0)
}

func TestIntListTakeWhile(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.TakeWhile(func(v int) bool {
		return v <= 20
	})

	g.Expect(b.Size()).To(Equal(20))
	g.Expect(b.Head()).To(Equal(1))
	g.Expect(b.Last()).To(Equal(20))

	c := a.TakeWhile(func(v int) bool {
		return true
	})

	g.Expect(c.Size()).To(Equal(100))
	g.Expect(c.Head()).To(Equal(1))
	g.Expect(c.Last()).To(Equal(100))

    // check correct nil handling
	a = nil
	a.TakeWhile(func(v int) bool {
		return v <= 20
	})
}

func TestIntListDropWhile(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.DropWhile(func(v int) bool {
		return v <= 80
	})

	g.Expect(b.Size()).To(Equal(20))
	g.Expect(b.Head()).To(Equal(81))
	g.Expect(b.Last()).To(Equal(100))

	c := a.DropWhile(func(v int) bool {
		return true
	})

	g.Expect(c.Size()).To(Equal(0))

    // check correct nil handling
	a = nil
	a.DropWhile(func(v int) bool {
		return v <= 20
	})
}

func TestIntListDistinctBy(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(1, 1, 1, 2, 1, 2, 3, 4, 5, 3, 3, 5)

	c := a.DistinctBy(func(v1, v2 int) bool {
		return v1 == v2
	})

	g.Expect(c.Equals(NewX1IntList(1, 2, 3, 4, 5))).To(BeTrue(), "%v", c)

    // check correct nil handling
    a = nil
	a.DistinctBy(func(v1, v2 int) bool {
		return v1 == v2
	})
}

func TestIntListIndexWhere(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.IndexWhere(func(v int) bool {
		return v >= 47
	})

	g.Expect(b).To(Equal(46))

	c := a.IndexWhere(func(v int) bool {
		return false
	})

	g.Expect(c).To(Equal(-1))

	d := a.IndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 10)

	g.Expect(d).To(Equal(11))
}

func TestIntListLastIndexWhere(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(intRangeOf(1, 100)...)

	b := a.LastIndexWhere(func(v int) bool {
		return v <= 47
	})

	g.Expect(b).To(Equal(46))

	c := a.LastIndexWhere(func(v int) bool {
		return false
	})

	g.Expect(c).To(Equal(-1))

	d := a.LastIndexWhere2(func(v int) bool {
		return v % 3 == 0
	}, 61)

	g.Expect(d).To(Equal(59))
}

func TestIntListMkString(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)

	c := a.MkString("|")

	g.Expect(c).To(Equal("13|4|7|-2|9"))

    // check correct nil handling
    a = nil
	a.MkString("|")
}

func TestIntListMkString3(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)

	c := a.MkString3("<", ", ", ">")

	g.Expect(c).To(Equal("<13, 4, 7, -2, 9>"))

    // check correct nil handling
    a = nil
	a.MkString3("<", ", ", ">")
}

func TestIntListGobEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)
	b := NewX1IntList()

    buf := &bytes.Buffer{}
    err := gob.NewEncoder(buf).Encode(a)

    g.Expect(err).NotTo(HaveOccurred())

    err = gob.NewDecoder(buf).Decode(&b)

    g.Expect(err).NotTo(HaveOccurred())

	g.Expect(a).To(Equal(b))
}

func TestIntListJsonEncode(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1IntList(13, 4, 7, -2, 9)
	b := NewX1IntList()

    buf, err := json.Marshal(a)

    g.Expect(err).NotTo(HaveOccurred())

    err = json.Unmarshal(buf, &b)

    g.Expect(err).NotTo(HaveOccurred())

	g.Expect(a).To(Equal(b))
}

func BenchmarkIntListTakeDrop(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    stuff := a.Take(100)
	    a.Drop(50).Append(stuff.ToSlice()...)
    }
}

func BenchmarkIntListTakeDropLast(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    stuff := a.TakeLast(100)
	    a.DropLast(50).Append(stuff.ToSlice()...)
    }
}

func BenchmarkIntListDistinctBy(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.DistinctBy(func(v1, v2 int) bool {
		    return v1 == v2
	    })
    }
}

func BenchmarkIntListCountBy(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.CountBy(func(v int) bool {
		    return v > 100
	    })
    }
}

func BenchmarkIntListFilter(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Filter(func(v int) bool {
		    return v > 100
	    })
    }
}

func BenchmarkIntListDoKeepWhere(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.DoKeepWhere(func(v int) bool {
		    return v > 100
	    })
    }
}

func BenchmarkIntListMap(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Map(func(v int) int {
		    return v
	    })
    }
}

func BenchmarkIntListFlatMap(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.FlatMap(func(v int) []int {
		    return []int{v}
	    })
    }
}

func BenchmarkIntListSum(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1IntList(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Sum()
    }
}

func fibonacciX1IntList(n int) *X1IntList {
    a := NewX1IntList()
    i0 := 1
    i1 := 1
    for j := 0; j < n; j++ {
        i2 := i0 + i1
        a.Append(i2)
        i1, i0 = i2, i1
    }
    return a
}