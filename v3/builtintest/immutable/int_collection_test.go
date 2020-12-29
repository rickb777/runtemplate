// Generated from ../collection_test.tpl with Type=int
// options: Append:<no value>

package immutable

import (
	. "github.com/onsi/gomega"
	"sort"
	"testing"
)

type Sizer interface {
	IsEmpty() bool
	NonEmpty() bool
	Size() int
}

func TestNewXIntCollection(t *testing.T) {
	testEmptyIntCollection(t, NewX1IntSet(), "Set")
	testEmptyIntCollection(t, NewX1IntList(), "List")
}

func testEmptyIntCollection(t *testing.T, a Sizer, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Size()).To(Equal(0))
	g.Expect(a.IsEmpty()).To(BeTrue())
	g.Expect(a.NonEmpty()).To(BeFalse())
}

func TestIntToSlice(t *testing.T) {
	testIntToSlice(t, NewX1IntSet(1, 2, 3), "Set")
	testIntToSlice(t, NewX1IntList(1, 2, 3), "List")
}

func testIntToSlice(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	s := a.ToSlice()

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(len(s)).To(Equal(3))
}

func TestIntExists(t *testing.T) {
	testIntExists1(t, NewX1IntSet(1, 2, 3), "Set")
	testIntExists1(t, NewX1IntList(1, 2, 3), "List")
}

func testIntExists1(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	has2 := a.Exists(func(v int) bool {
		return v > 2
	})

	g.Expect(has2).To(BeTrue())

	has5 := a.Exists(func(v int) bool {
		return v > 5
	})

	g.Expect(has5).To(BeFalse())
}

func TestIntForall(t *testing.T) {
	testIntForall(t, NewX1IntSet(1, 2, 3), "Set")
	testIntForall(t, NewX1IntList(1, 2, 3), "List")
}

func testIntForall(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	has1 := a.Forall(func(v int) bool {
		return v >= 1
	})

	g.Expect(has1).To(BeTrue())

	has2 := a.Forall(func(v int) bool {
		return v >= 2
	})

	g.Expect(has2).To(BeFalse())
}

func TestIntFind(t *testing.T) {
	testIntFind(t, NewX1IntSet(1, 2, 3), "Set")
	testIntFind(t, NewX1IntList(1, 2, 3, 4), "List")
}

func testIntFind(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	b, e := a.Find(func(v int) bool {
		return v > 2
	})

	g.Expect(e).To(BeTrue())
	g.Expect(b).To(Equal(3))

	_, e = a.Find(func(v int) bool {
		return v > 20
	})

	g.Expect(e).To(BeFalse())
}

func TestIntCountBy(t *testing.T) {
	testIntCountBy(t, NewX1IntSet(1, 2, 3), "Set")
	testIntCountBy(t, NewX1IntList(1, 2, 3), "List")
}

func testIntCountBy(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	g.Expect(n).To(Equal(2))
}

func TestIntForeach(t *testing.T) {
	testIntForeach(t, NewX1IntSet(1, 2, 3), "Set")
	testIntForeach(t, NewX1IntList(1, 2, 3), "List")
}

func testIntForeach(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	g.Expect(sum1).To(Equal(6))
}

func TestIntContains(t *testing.T) {
	testIntContains(t, NewX1IntSet(71, 1, 7, 13), "Set")
	testIntContains(t, NewX1IntList(71, 1, 7, 13), "List")
}

func testIntContains(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Contains(71)).To(BeTrue())
	g.Expect(a.Contains(72)).To(BeFalse())
	g.Expect((a.Contains(13) && a.Contains(7) && a.Contains(1))).To(BeTrue())
	g.Expect(a.ContainsAll(1, 7, 13)).To(BeTrue())
	g.Expect(a.ContainsAll(1, 3, 5, 7, 9, 11, 13)).To(BeFalse())
}

func TestIntFold(t *testing.T) {
	testIntFold(t, NewX1IntSet(10, 71, 3, 7, 13), "Set")
	testIntFold(t, NewX1IntList(10, 71, 3, 7, 13), "List")
}

func testIntFold(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Fold(10, func(a, b int) int { return a + b })).To(Equal(114))
	g.Expect(a.Fold(0, func(a, b int) int { return a + b })).To(Equal(104))
}

func TestIntMinMaxSum(t *testing.T) {
	testIntMinMaxSum(t, NewX1IntSet(10, 71, 3, 7, 13), "Set")
	testIntMinMaxSum(t, NewX1IntList(10, 71, 3, 7, 13), "List")
}

func testIntMinMaxSum(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Min()).To(Equal(3))
	g.Expect(a.Max()).To(Equal(71))
	g.Expect(a.Sum()).To(Equal(104))
	c := a.MinBy(func(v1, v2 int) bool {
		return v1 < v2
	})
	g.Expect(c).To(Equal(3))
	c = a.MaxBy(func(v1, v2 int) bool {
		return v1 < v2
	})
	g.Expect(c).To(Equal(71))
}

func TestIntSend(t *testing.T) {
	//testIntSend(t, NewX1IntSet(10, 71, 3, 7, 13), "Set")
	testIntSend(t, NewX1IntList(10, 71, 3, 7, 13), "List")
}

func testIntSend(t *testing.T, a X1IntCollection, kind string) {
	g := NewGomegaWithT(t)

	s := make([]int, 0)
	for v := range a.Send() {
		s = append(s, v)
	}
	g.Expect(s).To(Equal([]int{10, 71, 3, 7, 13}))
}

func TestIntStringer(t *testing.T) {
	testIntStringer(t, NewX1IntSet(10, 71, 3, 7, 13), false, "Set")
	testIntStringer(t, NewX1IntList(10, 71, 3, 7, 13), true, "List")
}

func testIntStringer(t *testing.T, a X1IntCollection, ordered bool, kind string) {
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
		t.Errorf("%s: Expected 5 but got %d", kind, len(sl))
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
}
