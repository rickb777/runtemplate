// Generated from {{.TemplateFile}} with Type={{.Type}}
// options: Append:{{.Append}} Find:{{.Find}} Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
	"testing"
	. "github.com/onsi/gomega"
)

func Test{{.Type.U}}Queue_withEquals(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(10, false)
	b := NewX1{{.Type.U}}Queue(12, true)

	g.Expect(a.IsSequence()).To(BeTrue())
	g.Expect(a.IsSet()).To(BeFalse())
	g.Expect(a.Equals(b)).To(BeTrue())
	g.Expect(a.Equals(nil)).To(BeTrue())

	a.Push(1)
	g.Expect(a.Equals(b)).To(BeFalse())
	g.Expect(a.Equals(nil)).To(BeFalse())

	b.Push(1)
	g.Expect(a.Equals(b)).To(BeTrue())

	a.Push(2)
	b.Push(3)
	g.Expect(a.Equals(b)).To(BeFalse())

	a = nil
	g.Expect(a.Equals(NewX1{{.Type.U}}Queue(2, true))).To(BeTrue())
	g.Expect(a.Equals(b)).To(BeFalse())
	g.Expect(a.Equals(nil)).To(BeTrue())
}

func Test{{.Type.U}}OverwritingQueueOffer(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(10, true)

	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.Space()).To(Equal(10))
	g.Expect(a.ToSlice()).To(HaveLen(0))

	r1 := a.Offer(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(a.Space()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6}))
	g.Expect(r1).To(HaveLen(0))

	r2 := a.Offer(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}))
	g.Expect(r2).To(HaveLen(0))

	r3 := a.Offer(20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}))
	g.Expect(r3).To(Equal([]{{.Type.Name}}{30}))
}

func Test{{.Type.U}}RefusingQueueOffer(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(10, false)

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.Space()).To(Equal(10))
	g.Expect(a.ToSlice()).To(HaveLen(0))

	r1 := a.Offer(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(a.Space()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6}))
	g.Expect(r1).To(HaveLen(0))

	r2 := a.Offer(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	g.Expect(r2).To(Equal([]{{.Type.Name}}{11, 12}))

	r3 := a.Offer(20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	g.Expect(r3).To(Equal([]{{.Type.Name}}{20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}))
}

func Test{{.Type.U}}OverwritingQueuePush(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(5, true)

	a.Push(1, 2, 3, 4)

	g.Expect(a.Cap()).To(Equal(5))
	g.Expect(a.Len()).To(Equal(4))
	g.Expect(a.Space()).To(Equal(1))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4}))

	a.Push(5, 6, 7, 8, 9, 10, 11)

	g.Expect(a.Cap()).To(Equal(7))
	g.Expect(a.Len()).To(Equal(7))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{5, 6, 7, 8, 9, 10, 11}))
}

func Test{{.Type.U}}RefusingQueuePush(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(5, false)

	a.Push(1, 2, 3, 4)

	g.Expect(a.Cap()).To(Equal(5))
	g.Expect(a.Len()).To(Equal(4))
	g.Expect(a.Space()).To(Equal(1))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4}))

	a.Push(5, 6, 7, 8, 9, 10, 11)

	g.Expect(a.Cap()).To(Equal(128))
	g.Expect(a.Len()).To(Equal(11))
	g.Expect(a.Space()).To(Equal(117))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
}

func Test{{.Type.U}}SortedQueue(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}SortedQueue(10, false, func(i, j int) bool {return i < j})

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.Space()).To(Equal(10))
	g.Expect(a.ToSlice()).To(HaveLen(0))

	r1 := a.Offer(2, 5, 1, 4, 6, 3, 9, 8, 7, 10)
	a.Sort()

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type.Name}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	g.Expect(r1).To(HaveLen(0))
}

func Test{{.Type.U}}QueuePop(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(10, false)

	g.Expect(a.IsEmpty()).To(BeTrue())
	g.Expect(a.NonEmpty()).To(BeFalse())
	g.Expect(a.IsFull()).To(BeFalse())
	g.Expect(a.Space()).To(Equal(10))

	_, ok := a.Pop1()
	g.Expect(ok).To(BeFalse())

	g.Expect(a.Pop(1)).To(HaveLen(0))

	a.Offer(1, 2, 3, 4, 5, 6)

	g.Expect(a.IsEmpty()).To(BeFalse())
	g.Expect(a.NonEmpty()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))

	v1, ok := a.Pop1()
	g.Expect(ok).To(BeTrue())
	g.Expect(v1).To(Equal(1))
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(5))

	v2, ok := a.Pop1()
	g.Expect(ok).To(BeTrue())
	g.Expect(v2).To(Equal(2))
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(4))

	a.Offer(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.IsFull()).To(BeTrue())
	g.Expect(a.Space()).To(Equal(0))

	s1 := a.Pop(3)
	g.Expect(s1).To(Equal([]int{3, 4, 5}))
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(7))

	s2 := a.Pop(8)
	g.Expect(s2).To(Equal([]int{6, 7, 8, 9, 10, 11, 12}))
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.IsFull()).To(BeFalse())
	g.Expect(a.Space()).To(Equal(10))
}

func Test{{.Type.U}}QueueGetHeadLast(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(8, false)

	g.Expect(a.HeadOption()).To(Equal(0))
	g.Expect(a.LastOption()).To(Equal(0))

	a.Offer(1)

	g.Expect(a.Get(0)).To(Equal(1))
	g.Expect(a.Head()).To(Equal(1))
	g.Expect(a.HeadOption()).To(Equal(1))
	g.Expect(a.Last()).To(Equal(1))
	g.Expect(a.LastOption()).To(Equal(1))

	a.Offer(2, 3, 4, 5, 6)
	a.Pop(2)

	g.Expect(a.Get(0)).To(Equal(3))
	g.Expect(a.Get(3)).To(Equal(6))
	g.Expect(a.Head()).To(Equal(3))
	g.Expect(a.HeadOption()).To(Equal(3))
	g.Expect(a.Last()).To(Equal(6))
	g.Expect(a.LastOption()).To(Equal(6))

	a.Offer(7, 8, 9, 10)
	a.Pop(4)

	g.Expect(a.Get(0)).To(Equal(7))
	g.Expect(a.Get(3)).To(Equal(10))
	g.Expect(a.Head()).To(Equal(7))
	g.Expect(a.HeadOption()).To(Equal(7))
	g.Expect(a.Last()).To(Equal(10))
	g.Expect(a.LastOption()).To(Equal(10))

	// check correct nil handling
	a = nil

	g.Expect(a.HeadOption()).To(Equal(0))
	g.Expect(a.LastOption()).To(Equal(0))
}

func Test{{.Type.U}}QueueSend(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingFullQueue(1, 2, 3, 4)
	b := BuildX1{{.Type.U}}QueueFromChan(a.Send())
	g.Expect(a.Equals(b)).To(BeTrue())

	a = interestingPartialQueue(1, 2, 3, 4)
	b = BuildX1{{.Type.U}}QueueFromChan(a.Send())
	g.Expect(a.Equals(b)).To(BeTrue())
{{- if .Mutable}}

    // check correct nil handling
	a = nil
	b = BuildX1{{.Type.U}}QueueFromChan(a.Send())

	g.Expect(b.IsEmpty()).To(BeTrue())
{{- end}}
}
{{- if and .Mutable .Numeric}}

func Test{{.Type.U}}QueueDoKeepWhere(t *testing.T) {
	g := NewGomegaWithT(t)

    cases := []struct{
        act, exp *X1{{.Type.U}}Queue
        m int
    }{
        {
            act: nil,
            exp: nil,
        },
        {
            act: NewX1{{.Type.U}}Queue(7, false).Push(1, 3, 5, 2, 4, 6),
            exp: NewX1{{.Type.U}}Queue(7, false).Push(5, 4, 6),
            m: 1,
        },
        {
            act: NewX1{{.Type.U}}Queue(6, false).Push(1, 3, 5, 2, 4, 6),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(5, 4, 6),
            m: 1,
        },
        {
            act: NewX1{{.Type.U}}Queue(7, false).Push(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(7, false).Push(4, 6, 5, 7),
            m: 1,
        },
        {
            act: NewX1{{.Type.U}}Queue(6, false).Push(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(4, 6, 5, 7),
            m: 1,
        },
        {
            act: NewX1{{.Type.U}}Queue(7, false).Push(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(4, 6, 5, 7),
            m: 1,
        },
        {
            act: NewX1{{.Type.U}}Queue(7, false).Push(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(),
            m: 0,
        },
        {
            act: interestingFullQueue(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(4, 6, 5, 7),
            m: 1,
        },
        {
            act: interestingPartialQueue(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(4, 6, 5, 7),
            m: 1,
        },
        {
            act: interestingPartialQueue(2, 4, 6, 3, 5, 7),
            exp: NewX1{{.Type.U}}Queue(6, false).Push(2, 4, 6, 3, 5, 7),
            m: 3,
        },
    }

    for i, c := range cases {
        r := c.act.DoKeepWhere(func (v {{.Type}}) bool {
            return (v * c.m) > 3
        })

        g.Expect(c.act.Equals(c.exp)).To(BeTrue(), "%d %v", i, c.act)
        g.Expect(c.act.Equals(r)).To(BeTrue(), "%d %v", i, r)
    }
}
{{- end}}

func Test{{.Type.U}}QueueClone(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(10, false)

	g.Expect(a.Clone().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToList().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToInterfaceSlice()).To(HaveLen(0))

	a.Offer(1, 2, 3, 4, 5, 6)
	a.Pop1()

	g.Expect(a.Clone().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6}))
	g.Expect(a.ToList().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6}))
	g.Expect(a.ToInterfaceSlice()).To(Equal([]interface{}{2, 3, 4, 5, 6}))

	a.Offer(7, 8, 9, 10, 11)

	g.Expect(a.Clone().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	g.Expect(a.ToList().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	g.Expect(a.ToInterfaceSlice()).To(Equal([]interface{}{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))

	// check correct nil handling
	a = nil

	g.Expect(a.Clone().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToList().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToInterfaceSlice()).To(HaveLen(0))
}

func Test{{.Type.U}}QueueToList(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(5, false)
	a.Push(1, 2, 3, 4)
	b := a.ToList()

	g.Expect(b.ToSlice()).To(Equal([]int{1, 2, 3, 4}))
{{- if .Mutable}}

	a = nil
	b = a.ToList()

	g.Expect(b.IsEmpty()).To(BeTrue())
{{- end}}
}

func Test{{.Type.U}}QueueToSet(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(5, false)
	a.Push(1, 2, 3, 4)
	b := a.ToSet()

	g.Expect(b.Size()).To(Equal(4))
	g.Expect(b.Contains(1)).To(BeTrue())
	g.Expect(b.Contains(2)).To(BeTrue())
	g.Expect(b.Contains(3)).To(BeTrue())
	g.Expect(b.Contains(4)).To(BeTrue())
{{- if .Mutable}}

	a = nil
	b = a.ToSet()

	g.Expect(b.IsEmpty()).To(BeTrue())
{{- end}}
}

func Test{{.Type.U}}QueueResize(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(4, false)

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.ToSlice()).To(HaveLen(0))
	g.Expect(a.Cap()).To(Equal(4))

	a.Offer(1, 2, 3, 4)

	g.Expect(a.IsFull()).To(BeTrue())
	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]int{1, 2, 3, 4}))

	a.Pop1()
	b := a.Reallocate(6, true)

	g.Expect(a).To(Equal(b))
	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(6))

	a.Offer(5, 6, 7, 8)

	g.Expect(a.Size()).To(Equal(6))
	g.Expect(b.Size()).To(Equal(6))
	g.Expect(a.ToSlice()).To(Equal([]int{3, 4, 5, 6, 7, 8}))

	a.Pop1()
	a.Reallocate(5, false)

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.Cap()).To(Equal(5))
	g.Expect(a.Size()).To(Equal(5))
	g.Expect(a.ToSlice()).To(Equal([]int{4, 5, 6, 7, 8}))

	a.Reallocate(4, true)

	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(4))
	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]int{5, 6, 7, 8}))
}

func Test{{.Type.U}}QueueContains(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(4, false).Push(1, 2, 3, 4)

	found := a.Contains(3)

	g.Expect(found).To(BeTrue())

	found = a.Contains(5)

	g.Expect(found).To(BeFalse())

	// check correct nil handling
	a = nil
	a.Contains(3)
}

func Test{{.Type.U}}QueueContainsAll(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(4, false).Push(1, 2, 3, 4)

	found := a.ContainsAll(3)

	g.Expect(found).To(BeTrue())

	found = a.ContainsAll(1, 3, 5, 7)

	g.Expect(found).To(BeFalse())

	// check correct nil handling
	a = nil
	a.ContainsAll(3)
}

func Test{{.Type.U}}QueueFind(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.Type.U}}Queue(4, false).Push(1, 2, 3, 4)

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

func Test{{.Type.U}}QueueForall(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingFullQueue(1, 2, 3, 4)

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

func Test{{.Type.U}}QueueExists(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingFullQueue(1, 2, 3, 4)
    a.Pop(2)
    a.Push(5, 6)

	found := a.Exists(func(v int) bool {
		return v > 3
	})

	g.Expect(found).To(BeTrue())

	found = a.Exists(func(v int) bool {
		return v > 5
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

func Test{{.Type.U}}QueueForeach(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingFullQueue(1, 2, 3, 4)

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

func Test{{.Type.U}}QueueFilter(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingPartialQueue(1, 2, 3, 4)

	b := a.Filter(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1{{.Type.U}}Queue(2, false).Push(3, 4))).To(BeTrue(), "%v", b)

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

func Test{{.Type.U}}QueuePartition(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingPartialQueue(1, 2, 3, 4)

	b, c := a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.Equals(NewX1{{.Type.U}}Queue(4, false).Push(3, 4))).To(BeTrue(), "%v", b)
	g.Expect(c.Equals(NewX1{{.Type.U}}Queue(4, false).Push(1, 2))).To(BeTrue(), "%v", c)

	a = nil
	b, c = a.Partition(func(v int) bool {
		return v > 2
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
	g.Expect(c.IsEmpty()).To(BeTrue())
}

func Test{{.Type.U}}QueueTransform(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingFullQueue(1, 2, 3, 4)

	b := a.Map(func(v int) int {
		return v * v
	})

	exp := NewX1{{.Type.U}}Queue(4, false).Push(1, 4, 9, 16)

	g.Expect(b.Equals(exp)).To(BeTrue(), "%v %v", b, exp)

	a = nil
	b = a.Map(func(v int) int {
		return v * v
	})

	g.Expect(b.IsEmpty()).To(BeTrue())
}

func Test{{.Type.U}}QueueMkString(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingPartialQueue(13, 4, 7, -2, 9)

	c := a.MkString("|")

	g.Expect(c).To(Equal("13|4|7|-2|9"))

	// check correct nil handling
	a = nil
	a.MkString("|")
}

func Test{{.Type.U}}QueueMkString3(t *testing.T) {
	g := NewGomegaWithT(t)

	a := interestingPartialQueue(13, 4, 7, -2, 9)

	c := a.MkString3("<", ", ", ">")

	g.Expect(c).To(Equal("<13, 4, 7, -2, 9>"))

	// check correct nil handling
	a = nil
	a.MkString3("<", ", ", ">")
}
//{{- if .GobEncode}}
//
//func Test{{.Type.U}}QueueGobEncode(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	a := NewX1{{.Type.U}}Queue(10, false).Push(13, 4, 7, -2, 9)
//	b := NewX1{{.Type.U}}Queue(10, false)
//
//    buf := &bytes.Buffer{}
//    err := gob.NewEncoder(buf).Encode(a)
//
//    g.Expect(err).NotTo(HaveOccurred())
//
//    err = gob.NewDecoder(buf).Decode(&b)
//
//    g.Expect(err).NotTo(HaveOccurred())
//
//	g.Expect(a).To(Equal(b))
//}
//{{- end}}
//
//func Test{{.Type.U}}QueueJsonEncode(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	a := NewX1{{.Type.U}}Queue(10, false).Push(13, 4, 7, -2, 9)
//	b := NewX1{{.Type.U}}Queue(10, false)
//
//    buf, err := json.Marshal(a)
//
//    g.Expect(err).NotTo(HaveOccurred())
//
//    err = json.Unmarshal(buf, &b)
//
//    g.Expect(err).NotTo(HaveOccurred())
//
//	g.Expect(a).To(Equal(b))
//}
//

{{- if .Mutable}}

func Benchmark{{.Type.U}}QueuePopPush(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    stuff := a.Pop(100)
	    a.Push(stuff...)
    }
}

func Benchmark{{.Type.U}}QueueCountBy(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.CountBy(func(v int) bool {
		    return v > 100
	    })
    }
}

func Benchmark{{.Type.U}}QueueFilter(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Filter(func(v int) bool {
		    return v > 100
	    })
    }
}

func Benchmark{{.Type.U}}QueueDoKeepWhere(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.DoKeepWhere(func(v int) bool {
		    return v > 100
	    })
    }
}

func Benchmark{{.Type.U}}QueueMap(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Map(func(v int) int {
		    return v
	    })
    }
}

func Benchmark{{.Type.U}}QueueSum(b *testing.B) {
    b.StopTimer()
    a := fibonacciX1{{.Type.U}}Queue(1000)
    b.StartTimer()

    for i := 0; i < b.N; i++ {
	    a.Sum()
    }
}

func fibonacciX1{{.Type.U}}Queue(n int) *X1{{.Type.U}}Queue {
    a := NewX1{{.Type.U}}Queue(n + 10, false)
    i0 := 1
    i1 := 1
    for j := 0; j < n; j++ {
        i2 := i0 + i1
        a.Push(i2)
        i1, i0 = i2, i1
    }
    return a
}
{{- end}}