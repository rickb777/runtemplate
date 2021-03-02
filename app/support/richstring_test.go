package support

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRichStringNoDots(t *testing.T) {
	g := NewGomegaWithT(t)

	a := RichString(".foo.Bar.Baz.").NoDots()
	g.Expect(a).To(BeEquivalentTo(RichString("fooBarBaz")))
}

func TestRichStringDivideOr0A(t *testing.T) {
	g := NewGomegaWithT(t)

	a, b := RichString("foo,bar").DivideLastOr0(',')
	g.Expect(a).To(BeEquivalentTo("foo"))
	g.Expect(b).To(BeEquivalentTo("bar"))
}

func TestRichStringDivideOr0B(t *testing.T) {
	g := NewGomegaWithT(t)

	a, b := RichString("foo bar").DivideLastOr0(',')
	g.Expect(a).To(BeEquivalentTo("foo bar"))
	g.Expect(b).To(BeEquivalentTo(""))
}

func TestRichStringDivideOr1A(t *testing.T) {
	g := NewGomegaWithT(t)

	a, b := RichString("foo,bar").DivideLastOr1(',')
	g.Expect(a).To(BeEquivalentTo("foo"))
	g.Expect(b).To(BeEquivalentTo("bar"))
}

func TestRichStringDivideOr1B(t *testing.T) {
	g := NewGomegaWithT(t)

	a, b := RichString("foo bar").DivideLastOr1(',')
	g.Expect(a).To(BeEquivalentTo(""))
	g.Expect(b).To(BeEquivalentTo("foo bar"))
}

func TestRichStringRemoveBeforeA(t *testing.T) {
	g := NewGomegaWithT(t)

	a := RichString("foo/bar/baz").RemoveBeforeLast('/')
	g.Expect(a).To(BeEquivalentTo("baz"))
}

func TestRichStringRemoveBeforeB(t *testing.T) {
	g := NewGomegaWithT(t)

	a := RichString("foo/bar/baz").RemoveBeforeLast(',')
	g.Expect(a).To(BeEquivalentTo("foo/bar/baz"))
}
