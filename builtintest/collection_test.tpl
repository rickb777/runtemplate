// Generated from {{.TemplateFile}} with Type={{.PType}}
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

func TestNewX{{.UType}}Collection(t *testing.T) {
	testEmpty{{.UType}}Collection(t, NewX1{{.UType}}Set(), "Set")
	testEmpty{{.UType}}Collection(t, NewX1{{.UType}}List(), "List")
{{- if .Mutable}}
	testEmpty{{.UType}}Collection(t, NewX1{{.UType}}Queue(5, false), "Queue")
{{- end}}
}

func testEmpty{{.UType}}Collection(t *testing.T, a Sizer, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Size()).To(Equal(0))
	g.Expect(a.IsEmpty()).To(BeTrue())
	g.Expect(a.NonEmpty()).To(BeFalse())
}

func Test{{.UType}}ToSlice(t *testing.T) {
	test{{.UType}}ToSlice(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}ToSlice(t, NewX1{{.UType}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.UType}}ToSlice(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}ToSlice(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}ToSlice(t *testing.T, a X1{{.UType}}Collection, kind string) {
	g := NewGomegaWithT(t)

	s := a.ToSlice()

	g.Expect(a.Size()).To(Equal(3))
	g.Expect(len(s)).To(Equal(3))
}

func Test{{.UType}}Exists(t *testing.T) {
	test{{.UType}}Exists1(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}Exists1(t, NewX1{{.UType}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.UType}}Exists1(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}Exists1(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}Exists1(t *testing.T, a X1{{.UType}}Collection, kind string) {
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

func Test{{.UType}}Forall(t *testing.T) {
	test{{.UType}}Forall(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}Forall(t, NewX1{{.UType}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.UType}}Forall(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}Forall(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}Forall(t *testing.T, a X1{{.UType}}Collection, kind string) {
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

func Test{{.UType}}Find(t *testing.T) {
	test{{.UType}}Find(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}Find(t, NewX1{{.UType}}List(1, 2, 3, 4), "List")
{{- if .Mutable}}
	test{{.UType}}Find(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}Find(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}Find(t *testing.T, a X1{{.UType}}Collection, kind string) {
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

func Test{{.UType}}CountBy(t *testing.T) {
	test{{.UType}}CountBy(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}CountBy(t, NewX1{{.UType}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.UType}}CountBy(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}CountBy(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}CountBy(t *testing.T, a X1{{.UType}}Collection, kind string) {
	g := NewGomegaWithT(t)

	n := a.CountBy(func(v int) bool {
		return v >= 2
	})

	g.Expect(n).To(Equal(2))
}

func Test{{.UType}}Foreach(t *testing.T) {
	test{{.UType}}Foreach(t, NewX1{{.UType}}Set(1, 2, 3), "Set")
	test{{.UType}}Foreach(t, NewX1{{.UType}}List(1, 2, 3), "List")
{{- if .Mutable}}
	test{{.UType}}Foreach(t, interestingFullQueue(1, 2, 3), "Queue")
	test{{.UType}}Foreach(t, interestingPartialQueue(1, 2, 3), "Queue")
{{- end}}
}

func test{{.UType}}Foreach(t *testing.T, a X1{{.UType}}Collection, kind string) {
	g := NewGomegaWithT(t)

	sum1 := int(0)
	a.Foreach(func(v int) {
		sum1 += v
	})

	g.Expect(sum1).To(Equal(6))
}

func Test{{.UType}}Contains(t *testing.T) {
	test{{.UType}}Contains(t, NewX1{{.UType}}Set(71, 1, 7, 13), "Set")
	test{{.UType}}Contains(t, NewX1{{.UType}}List(71, 1, 7, 13), "List")
{{- if .Mutable}}
	test{{.UType}}Contains(t, interestingFullQueue(71, 1, 7, 13), "Queue")
	test{{.UType}}Contains(t, interestingPartialQueue(71, 1, 7, 13), "Queue")
{{- end}}
}

func test{{.UType}}Contains(t *testing.T, a X1{{.UType}}Collection, kind string) {
	g := NewGomegaWithT(t)

	g.Expect(a.Contains(71)).To(BeTrue())
	g.Expect(a.Contains(72)).To(BeFalse())
	g.Expect((a.Contains(13) && a.Contains(7) && a.Contains(1))).To(BeTrue())
	g.Expect(a.ContainsAll(1, 7, 13)).To(BeTrue())
	g.Expect(a.ContainsAll(1, 3, 5, 7, 9, 11, 13)).To(BeFalse())
}

func Test{{.UType}}MinMaxSum(t *testing.T) {
	test{{.UType}}MinMaxSum(t, NewX1{{.UType}}Set(10, 71, 3, 7, 13), "Set")
	test{{.UType}}MinMaxSum(t, NewX1{{.UType}}List(10, 71, 3, 7, 13), "List")
{{- if .Mutable}}
	test{{.UType}}MinMaxSum(t, interestingFullQueue(10, 71, 3, 7, 13), "Queue")
	test{{.UType}}MinMaxSum(t, interestingPartialQueue(10, 71, 3, 7, 13), "Queue")
{{- end}}
}

func test{{.UType}}MinMaxSum(t *testing.T, a X1{{.UType}}Collection, kind string) {
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

func Test{{.UType}}Send(t *testing.T) {
	//test{{.UType}}Send(t, NewX1{{.UType}}Set(10, 71, 3, 7, 13), "Set")
	test{{.UType}}Send(t, NewX1{{.UType}}List(10, 71, 3, 7, 13), "List")
{{- if .Mutable}}
	test{{.UType}}Send(t, interestingFullQueue(10, 71, 3, 7, 13), "Queue")
	test{{.UType}}Send(t, interestingPartialQueue(10, 71, 3, 7, 13), "Queue")
{{- end}}
}

func test{{.UType}}Send(t *testing.T, a X1{{.UType}}Collection, kind string) {
	g := NewGomegaWithT(t)

    s := make([]{{.PType}}, 0)
    for v := range a.Send() {
        s = append(s, v)
    }
	g.Expect(s).To(Equal([]{{.PType}}{10, 71, 3, 7, 13}))
}

func Test{{.UType}}Stringer(t *testing.T) {
	test{{.UType}}Stringer(t, NewX1{{.UType}}Set(10, 71, 3, 7, 13), false, "Set")
	test{{.UType}}Stringer(t, NewX1{{.UType}}List(10, 71, 3, 7, 13), true, "List")
{{- if .Mutable}}
	test{{.UType}}Stringer(t, interestingFullQueue(10, 71, 3, 7, 13), true, "Queue")
	test{{.UType}}Stringer(t, interestingPartialQueue(10, 71, 3, 7, 13), true, "Queue")
{{- end}}
}

func test{{.UType}}Stringer(t *testing.T, a X1{{.UType}}Collection, ordered bool, kind string) {
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

func interestingFullQueue(values ...{{.PType}}) *X1{{.UType}}Queue {
    // need to use Pop in order to wrap the read & write indexes of the queue
    q := NewX1{{.UType}}Queue(len(values), false).Push(0, 1, 2, values[0])
    q.Pop(3)
    q.Push(values[1:]...)
	return q
}

func interestingPartialQueue(values ...{{.PType}}) *X1{{.UType}}Queue {
    // need to use Pop in order to wrap the read & write indexes of the queue
    q := NewX1{{.UType}}Queue(len(values)+2, false).Push(0, 1, 2, values[0])
    q.Pop(3)
    q.Push(values[1:]...)
	return q
}
{{- end}}