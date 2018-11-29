// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Append:{{.Append}} Find:{{.Find}} Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
	"testing"
	. "github.com/onsi/gomega"
)

func Test{{.UType}}QueueBasics(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.IsSequence()).To(BeTrue())
	g.Expect(a.IsSet()).To(BeFalse())
}

func Test{{.UType}}OverwritingQueue(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, true)

	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.Space()).To(Equal(10))
	g.Expect(a.ToSlice()).To(HaveLen(0))

    r1 := a.Push(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(a.Space()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{1, 2, 3, 4, 5, 6}))
	g.Expect(r1).To(HaveLen(0))

    r2 := a.Push(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{3, 4, 5, 6, 7, 8, 9, 10, 11, 12}))
	g.Expect(r2).To(HaveLen(0))

    r3 := a.Push(20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}))
	g.Expect(r3).To(Equal([]{{.Type}}{30}))
}

func Test{{.UType}}RefusingQueue(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))
	g.Expect(a.Space()).To(Equal(10))
	g.Expect(a.ToSlice()).To(HaveLen(0))

    r1 := a.Push(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(a.Space()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{1, 2, 3, 4, 5, 6}))
	g.Expect(r1).To(HaveLen(0))

    r2 := a.Push(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	g.Expect(r2).To(Equal([]{{.Type}}{11, 12}))

    r3 := a.Push(20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(a.Space()).To(Equal(0))
	g.Expect(a.ToSlice()).To(Equal([]{{.Type}}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	g.Expect(r3).To(Equal([]{{.Type}}{20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}))
}

func Test{{.UType}}QueuePop(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.IsEmpty()).To(BeTrue())
	g.Expect(a.NonEmpty()).To(BeFalse())
	g.Expect(a.IsFull()).To(BeFalse())
	g.Expect(a.Space()).To(Equal(10))

    _, ok := a.Pop1()
	g.Expect(ok).To(BeFalse())

	g.Expect(a.Pop(1)).To(HaveLen(0))

    a.Push(1, 2, 3, 4, 5, 6)

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

    a.Push(7, 8, 9, 10, 11, 12)

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

func Test{{.UType}}QueueHeadLast(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.HeadOption()).To(Equal(0))
	g.Expect(a.LastOption()).To(Equal(0))

    a.Push(1)

	g.Expect(a.HeadOption()).To(Equal(1))
	g.Expect(a.LastOption()).To(Equal(1))

    a.Push(2, 3, 4, 5, 6)
    a.Pop(2)

	g.Expect(a.HeadOption()).To(Equal(3))
	g.Expect(a.LastOption()).To(Equal(6))
}

func Test{{.UType}}QueueClone(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.Clone().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToList().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToInterfaceSlice()).To(HaveLen(0))

    a.Push(1, 2, 3, 4, 5, 6)
    a.Pop1()

	g.Expect(a.Clone().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6}))
	g.Expect(a.ToList().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6}))
	g.Expect(a.ToInterfaceSlice()).To(Equal([]interface{}{2, 3, 4, 5, 6}))

    a.Push(7, 8, 9, 10, 11)

	g.Expect(a.Clone().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	g.Expect(a.ToList().ToSlice()).To(Equal([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))
	g.Expect(a.ToInterfaceSlice()).To(Equal([]interface{}{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}))

    // check correct nil handling
    a = nil

	g.Expect(a.Clone().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToList().ToSlice()).To(HaveLen(0))
	g.Expect(a.ToInterfaceSlice()).To(HaveLen(0))
}

func Test{{.UType}}QueueResize(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(4, false)

	g.Expect(a.IsOverwriting()).To(BeFalse())
	g.Expect(a.ToSlice()).To(HaveLen(0))
	g.Expect(a.Cap()).To(Equal(4))

    a.Push(1, 2, 3, 4)

	g.Expect(a.IsFull()).To(BeTrue())
	g.Expect(a.Size()).To(Equal(4))
	g.Expect(a.ToSlice()).To(Equal([]int{1, 2, 3, 4}))

    a.Pop1()
	b := a.Reallocate(6, true)

	g.Expect(a).To(Equal(b))
	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(6))

	a.Push(5, 6, 7, 8)

	g.Expect(a.Size()).To(Equal(6))
	g.Expect(b.Size()).To(Equal(6))
	g.Expect(a.ToSlice()).To(Equal([]int{3, 4, 5, 6, 7, 8}))

    a.Pop1()
	a.Reallocate(3, true)

	g.Expect(a.IsOverwriting()).To(BeTrue())
	g.Expect(a.Cap()).To(Equal(3))
	g.Expect(a.Size()).To(Equal(3))
	g.Expect(a.ToSlice()).To(Equal([]int{6, 7, 8}))

    // check correct nil handling
    a = nil
	b = a.Reallocate(7, false)

	g.Expect(b.IsOverwriting()).To(BeFalse())
	g.Expect(b.IsEmpty()).To(BeTrue())
	g.Expect(b.IsFull()).To(BeFalse())
	g.Expect(b.Cap()).To(Equal(7))
}
