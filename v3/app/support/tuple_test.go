package support

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestPairsTValues(t *testing.T) {
	g := NewGomegaWithT(t)

	triples := Tuples([]Tuple{NewTuple("a=x/z"), NewTuple("b=*big.Int"), NewTuple("c=interface{}/Any/nil"), NewTuple("d=3/Cho")})

	g.Expect(triples.TValues()).To(Equal([]string{"z", "bigInt", "Any", "Cho"}))

	g.Expect(triples[0].IsPtr()).To(BeFalse())
	g.Expect(triples[0].String()).To(Equal("x"))
	g.Expect(triples[0].Name()).To(Equal("x"))
	g.Expect(triples[0].Ident()).To(Equal(RichString("z")))
	g.Expect(triples[0].Zero()).To(Equal("*(new(x))"))

	g.Expect(triples[1].IsPtr()).To(BeTrue())
	g.Expect(triples[1].String()).To(Equal("*big.Int"))
	g.Expect(triples[1].Name()).To(Equal("big.Int"))
	g.Expect(triples[1].Ident()).To(Equal(RichString("bigInt")))
	g.Expect(triples[1].Zero()).To(Equal("nil"))

	g.Expect(triples[2].IsPtr()).To(BeFalse())
	g.Expect(triples[2].String()).To(Equal("interface{}"))
	g.Expect(triples[2].Name()).To(Equal("interface{}"))
	g.Expect(triples[2].Ident()).To(Equal(RichString("Any")))
	g.Expect(triples[2].Zero()).To(Equal("nil"))
}

func TestNewType(t *testing.T) {
	g := NewGomegaWithT(t)

	cases := []struct{ in, exp Type }{
		{in: NewType("/x/y/z"), exp: Type{}},
		{in: NewType(""), exp: Type{}},
		{in: NewType("foo"), exp: Type{s: "foo", ident: "foo"}},
		{in: NewType("*foo"), exp: Type{s: "*foo", ident: "foo"}},
		{in: NewType("*big.Int"), exp: Type{s: "*big.Int", ident: "bigInt"}},
		{in: NewType("*big.Int/Integer"), exp: Type{s: "*big.Int", ident: "Integer"}},
		{in: NewType("*big.Int/Integer/nil"), exp: Type{s: "*big.Int", ident: "Integer", zero: "nil"}},
	}

	for _, c := range cases {
		g.Expect(c.in).To(Equal(c.exp))
	}
}
