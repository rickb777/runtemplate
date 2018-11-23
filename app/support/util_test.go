package support

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFindTemplateArg1(t *testing.T) {
	g := NewGomegaWithT(t)

	tpl, args := FindTemplateArg("", nil)
	g.Expect(tpl).To(Equal(""))
	g.Expect(len(args)).To(Equal(0))
}

func TestFindTemplateArg2(t *testing.T) {
	g := NewGomegaWithT(t)

	tpl, args := FindTemplateArg("", []string{"x=1", "y=2", "f.tpl", "z=3"})
	g.Expect(tpl).To(Equal("f.tpl"))
	g.Expect(args).To(Equal([]string{"x=1", "y=2", "z=3"}))
}

func TestSplitKeyValArgs1(t *testing.T) {
	g := NewGomegaWithT(t)

	types, others, left := SplitKeyValArgs(nil)
	g.Expect(types).To(HaveLen(0))
	g.Expect(others).To(HaveLen(0))
	g.Expect(left).To(HaveLen(0))
}

func TestSplitKeyValArgs2(t *testing.T) {
	g := NewGomegaWithT(t)

	types, others, left := SplitKeyValArgs([]string{"w=t1", "x=t2", "foo", "y=xyz/abc/123", "a:v1", "z=t3", "b:v2", `c:a\n\tb`, "q:", ""})

	g.Expect(types).To(Equal(Tuples([]Tuple{NewTuple("w=t1"), NewTuple("x=t2"), NewTuple("y=xyz/abc/123"), NewTuple("z=t3")})))
	g.Expect(others).To(Equal(Pairs([]Pair{NewPair("a:v1"), NewPair("b:v2"), NewPair("c:a\n\tb"), NewPair("q:")})))
	g.Expect(left).To(Equal([]string{"foo", ""}))
}
