// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Append:{{.Append}}

package {{.Package}}

import (
	"testing"
	"sort"
	. "github.com/onsi/gomega"
)

type Sizer interface {
	IsEmpty() bool
	NonEmpty() bool
	Size() int
}

func TestNewX{{.Type.U}}Collection(t *testing.T) {
	testEmpty{{.Type.U}}Collection(t, NewX1{{.Type.U}}Set(), "Set")
	testEmpty{{.Type.U}}Collection(t, NewX1{{.Type.U}}List(), "List")
{{- if .Mutable}}
	testEmpty{{.Type.U}}Collection(t, NewX1{{.Type.U}}Queue(5, false), "Queue")
{{- end}}
}

func testEmpty{{.Type.U}}Collection(t *testing.T, a Sizer, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Size()).To(Equal(0))
	g.Expect(a.IsEmpty()).To(BeTrue())
	g.Expect(a.NonEmpty()).To(BeFalse())
}

func Test{{.Type.U}}ToSlice(t *testing.T) {
	test{{.Type.U}}ToSlice(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}ToSlice(t, NewX1{{.Type.U}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.Type.U}}ToSlice(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}ToSlice(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}ToSlice(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

	s := a.ToSlice()

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(len(s)).To(Equal(3))
}

func Test{{.Type.U}}Exists(t *testing.T) {
	test{{.Type.U}}Exists1(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}Exists1(t, NewX1{{.Type.U}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.Type.U}}Exists1(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}Exists1(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}Exists1(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
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

func Test{{.Type.U}}Forall(t *testing.T) {
	test{{.Type.U}}Forall(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}Forall(t, NewX1{{.Type.U}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.Type.U}}Forall(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}Forall(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}Forall(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
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

func Test{{.Type.U}}Find(t *testing.T) {
	test{{.Type.U}}Find(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}Find(t, NewX1{{.Type.U}}List(1, 2, 3, 4), "List")
{{- if .Mutable}}
	test{{.Type.U}}Find(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}Find(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}Find(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
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

func Test{{.Type.U}}CountBy(t *testing.T) {
	test{{.Type.U}}CountBy(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}CountBy(t, NewX1{{.Type.U}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.Type.U}}CountBy(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}CountBy(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}CountBy(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	g.Expect(n).To(Equal(2))
}

func Test{{.Type.U}}Foreach(t *testing.T) {
	test{{.Type.U}}Foreach(t, NewX1{{.Type.U}}Set(1, 2, 3), "Set")
	test{{.Type.U}}Foreach(t, NewX1{{.Type.U}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.Type.U}}Foreach(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.Type.U}}Foreach(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.Type.U}}Foreach(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	g.Expect(sum1).To(Equal(6))
}

func Test{{.Type.U}}Contains(t *testing.T) {
	test{{.Type.U}}Contains(t, NewX1{{.Type.U}}Set(71, 1, 7, 13), "Set")
	test{{.Type.U}}Contains(t, NewX1{{.Type.U}}List(71, 1, 7, 13), "List")
{{- if .Mutable}}
	test{{.Type.U}}Contains(t, interestingFullQueue(71, 1, 7, 13), "Queue")
	test{{.Type.U}}Contains(t, interestingPartialQueue(71, 1, 7, 13), "Queue")
{{- end}}
}

func test{{.Type.U}}Contains(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Contains(71)).To(BeTrue())
	g.Expect(a.Contains(72)).To(BeFalse())
	g.Expect((a.Contains(13) && a.Contains(7) && a.Contains(1))).To(BeTrue())
	g.Expect(a.ContainsAll(1, 7, 13)).To(BeTrue())
	g.Expect(a.ContainsAll(1, 3, 5, 7, 9, 11, 13)).To(BeFalse())
}

func Test{{.Type.U}}Fold(t *testing.T) {
	test{{.Type.U}}Fold(t, NewX1{{.Type.U}}Set(10, 71, 3, 7, 13), "Set")
	test{{.Type.U}}Fold(t, NewX1{{.Type.U}}List(10, 71, 3, 7, 13), "List")
{{- if .Mutable}}
	test{{.Type.U}}Fold(t, interestingFullQueue(10, 71, 3, 7, 13), "Queue")
	test{{.Type.U}}Fold(t, interestingPartialQueue(10, 13, 3, 7, 71), "Queue")
	test{{.Type.U}}Fold(t, interestingPartialQueue(10, 71, 3, 7, 13), "Queue")
{{- end}}
}

func test{{.Type.U}}Fold(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Fold(10, func(a, b {{.Type}}) {{.Type}} {return a+b} )).To(Equal(114))
	g.Expect(a.Fold(0, func(a, b {{.Type}}) {{.Type}} {return a+b} )).To(Equal(104))
}

func Test{{.Type.U}}MinMaxSum(t *testing.T) {
	test{{.Type.U}}MinMaxSum(t, NewX1{{.Type.U}}Set(10, 71, 3, 7, 13), "Set")
	test{{.Type.U}}MinMaxSum(t, NewX1{{.Type.U}}List(10, 71, 3, 7, 13), "List")
{{- if .Mutable}}
	test{{.Type.U}}MinMaxSum(t, interestingFullQueue(10, 71, 3, 7, 13), "Queue")
	test{{.Type.U}}MinMaxSum(t, interestingPartialQueue(10, 13, 3, 7, 71), "Queue")
	test{{.Type.U}}MinMaxSum(t, interestingPartialQueue(10, 71, 3, 7, 13), "Queue")
{{- end}}
}

func test{{.Type.U}}MinMaxSum(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
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

func Test{{.Type.U}}Send(t *testing.T) {
	//test{{.Type.U}}Send(t, NewX1{{.Type.U}}Set(10, 71, 3, 7, 13), "Set")
	test{{.Type.U}}Send(t, NewX1{{.Type.U}}List(10, 71, 3, 7, 13), "List")
{{- if .Mutable}}
	test{{.Type.U}}Send(t, interestingFullQueue(10, 71, 3, 7, 13), "Queue")
	test{{.Type.U}}Send(t, interestingPartialQueue(10, 71, 3, 7, 13), "Queue")
{{- end}}
}

func test{{.Type.U}}Send(t *testing.T, a X1{{.Type.U}}Collection, kind string) {
	g := NewGomegaWithT(t)

    s := make([]{{.Type}}, 0)
    for v := range a.Send() {
        s = append(s, v)
    }
	g.Expect(s).To(Equal([]{{.Type}}{10, 71, 3, 7, 13}))
}

func Test{{.Type.U}}Stringer(t *testing.T) {
	test{{.Type.U}}Stringer(t, NewX1{{.Type.U}}Set(10, 71, 3, 7, 13), false, "Set")
	test{{.Type.U}}Stringer(t, NewX1{{.Type.U}}List(10, 71, 3, 7, 13), true, "List")
{{- if .Mutable}}
	test{{.Type.U}}Stringer(t, interestingFullQueue(10, 71, 3, 7, 13), true, "Queue")
	test{{.Type.U}}Stringer(t, interestingPartialQueue(10, 71, 3, 7, 13), true, "Queue")
{{- end}}
}

func test{{.Type.U}}Stringer(t *testing.T, a X1{{.Type.U}}Collection, ordered bool, kind string) {
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
{{- if .Mutable}}

func interestingFullQueue(values ...{{.Type}}) *X1{{.Type.U}}Queue {
    // need to use Pop in order to wrap the read & write indexes of the queue
    q := NewX1{{.Type.U}}Queue(len(values), false).Push(0, 1, 2, values[0])
    q.Pop(3)
    q.Push(values[1:]...)
	return q
}

func interestingPartialQueue(values ...{{.Type}}) *X1{{.Type.U}}Queue {
    // need to use Pop in order to wrap the read & write indexes of the queue
    q := NewX1{{.Type.U}}Queue(len(values)+2, false).Push(0, 1, 2, values[0])
    q.Pop(3)
    q.Push(values[1:]...)
	return q
}
{{- end}}