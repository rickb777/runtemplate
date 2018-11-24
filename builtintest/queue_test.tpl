// Generated from {{.TemplateFile}} with Type={{.PType}}
// options: Append:{{.Append}} Find:{{.Find}} Mutable:{{.Mutable}} M:{{.M}}

package {{.Package}}

import (
	"testing"
	. "github.com/onsi/gomega"
)

func Test{{.UType}}OverwritingQueue(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, true)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))

    r1 := a.Push(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(r1).To(HaveLen(0))

    r2 := a.Push(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(r2).To(HaveLen(0))

    r3 := a.Push(20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(10))
	g.Expect(r3).To(Equal([]{{.Type}}{30}))
}

func Test{{.UType}}RefusingQueue(t *testing.T) {
	g := NewGomegaWithT(t)

	a := NewX1{{.UType}}Queue(10, false)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(0))

    r1 := a.Push(1, 2, 3, 4, 5, 6)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(a.Len()).To(Equal(6))
	g.Expect(r1).To(HaveLen(0))

    r2 := a.Push(7, 8, 9, 10, 11, 12)

	g.Expect(a.Cap()).To(Equal(10))
	g.Expect(r2).To(Equal([]{{.Type}}{11, 12}))
}
